#!/usr/bin/env python3
"""从缓存的微信官方文档页面里抽取请求/返回参数表。

用法:
    python3 doc/api/tools/extract_api_fields.py          # 解析 .local/apidoc/pages 下所有页面

产物:
    .local/apidoc/fields/<page>.json    单接口的请求/返回字段
    .local/apidoc/fields.json           全量合并（按 cat 分组）
    .local/apidoc/fields.md             人读速查（每个接口一节）
"""

import html
import json
import os
import re

ROOT = os.path.abspath(os.path.join(os.path.dirname(__file__), "..", "..", ".."))
CACHE = os.path.join(ROOT, ".local", "apidoc")
PAGES = os.path.join(CACHE, "pages")
FIELDS = os.path.join(CACHE, "fields")

TAG_RE = re.compile(r"<[^>]+>")
WS_RE = re.compile(r"\s+")
ROW_RE = re.compile(r"<tr>(.*?)</tr>", re.S)
CELL_RE = re.compile(r"<t[hd][^>]*>(.*?)</t[hd]>", re.S)
TABLE_RE = re.compile(r"<table>(.*?)</table>", re.S)
H3_RE = re.compile(r"<h3[^>]*>(.*?)</h3>", re.S)
H2_RE = re.compile(r'<h2 id="([^"]*)"[^>]*>(.*?)</h2>', re.S)
CALL_RE = re.compile(r'(GET|POST)\s+(?:[A-Za-z]+\s+)?(https://api\.weixin\.qq\.com(/[A-Za-z0-9_./-]*))')
EN_RE = re.compile(r'接口英文名[：:]\s*([A-Za-z0-9_]+)')
DATE_RE = re.compile(r'最后一次编辑于\s*(\d{4}-\d{2}-\d{2})')
TITLE_RE = re.compile(r'<title>(.*?)</title>', re.S)


def clean(s):
    s = TAG_RE.sub(" ", s)
    s = html.unescape(s)
    return WS_RE.sub(" ", s).strip()


def section(raw, key):
    """按二级标题切段。"""
    for m in H2_RE.finditer(raw):
        if key in m.group(1):
            start = m.end()
            nxt = raw.find("<h2 ", start)
            return raw[start: nxt if nxt > 0 else len(raw)]
    return ""


def parse_tables(block):
    """返回 [(表名, [行...])]，行是 (参数名, 类型, 必填, 说明)。"""
    out = []
    pos = 0
    for tm in TABLE_RE.finditer(block):
        head = block[: tm.start()]
        names = H3_RE.findall(head)
        group = clean(names[-1]) if names else ""
        rows = []
        for rm in ROW_RE.finditer(tm.group(1)):
            cells = [clean(c) for c in CELL_RE.findall(rm.group(1))]
            if not cells:
                continue
            rows.append(cells)
        if not rows:
            continue
        header = rows[0]
        body = rows[1:] if header and header[0] in ("参数名", "字段名", "参数") else rows
        out.append((group, header, body))
    return out


def parse_page(path):
    raw = open(path, encoding="utf-8", errors="ignore").read()
    fn = os.path.basename(path)
    seg = fn[:-5].split("__")

    call = CALL_RE.search(raw)
    en = EN_RE.search(raw)
    dt = DATE_RE.search(raw)

    req_block = section(raw, "请求参数")
    resp_block = section(raw, "返回参数")

    # 官方表格有多种列布局：参数名|类型|必填|说明、参数名|类型|必填|示例|说明、
    # 参数名|类型|示例|说明 等，必须按表头名定位列，写死下标会把「说明」列丢掉、
    # 把「示例」当成说明。
    COLUMN_ALIASES = {
        "name": ("参数名", "字段名", "参数"),
        "type": ("类型",),
        "required": ("必填",),
        "example": ("示例",),
        "desc": ("说明", "描述"),
    }

    def col_index(header):
        idx = {}
        for i, h in enumerate(header):
            h = h.strip()
            for key, aliases in COLUMN_ALIASES.items():
                if h in aliases and key not in idx:
                    idx[key] = i
        return idx

    def cell(cells, idx, key, fallback):
        if key in idx and len(cells) > idx[key]:
            return cells[idx[key]]
        return fallback

    def norm(tables, with_required):
        items = []
        for group, header, body in tables:
            idx = col_index(header)
            for cells in body:
                name = cell(cells, idx, "name", cells[0] if cells else "")
                typ = cell(cells, idx, "type", cells[1] if len(cells) > 1 else "")
                if with_required:
                    req = cell(cells, idx, "required", cells[2] if len(cells) > 2 else "")
                else:
                    req = ""
                example = cell(cells, idx, "example", "")
                desc = cell(cells, idx, "desc", "")
                if not desc and example:
                    desc = example
                if not desc:
                    desc = cells[3] if (with_required and len(cells) > 3) else (cells[2] if len(cells) > 2 else "")
                if name in ("参数名", "字段名", "参数"):
                    continue
                items.append({"group": group, "name": name, "type": typ,
                              "required": req, "example": example, "desc": desc})
        return items

    return {
        "cat": "/".join(seg[:-1]),
        "slug": seg[-1].replace(".html", ""),
        "title": clean(TITLE_RE.search(raw).group(1)).split(" - ")[0] if TITLE_RE.search(raw) else "",
        "en": en.group(1) if en else "",
        "method": call.group(1) if call else "",
        "url": re.sub(r"\?.*$", "", call.group(3)).rstrip("/") if call else "",
        "updated": dt.group(1) if dt else "",
        "doc": "https://developers.weixin.qq.com/miniprogram/dev/server/API/%s/%s" % ("/".join(seg[:-1]), seg[-1]),
        "request": norm(parse_tables(req_block), True),
        "response": norm(parse_tables(resp_block), False),
    }


def main():
    os.makedirs(FIELDS, exist_ok=True)
    all_rows = []
    for fn in sorted(os.listdir(PAGES)):
        path = os.path.join(PAGES, fn)
        if not os.path.getsize(path):
            continue
        d = parse_page(path)
        json.dump(d, open(os.path.join(FIELDS, fn[:-5] + ".json"), "w"),
                  ensure_ascii=False, indent=1)
        all_rows.append(d)
    json.dump(all_rows, open(os.path.join(CACHE, "fields.json"), "w"), ensure_ascii=False, indent=1)

    md = ["# 官方接口字段速查\n"]
    for d in all_rows:
        md.append("\n## %s / %s  `%s`\n" % (d["cat"], d["title"], d["url"]))
        md.append("- 英文名：`%s`  方法：%s  文档：%s\n" % (d["en"], d["method"], d["doc"]))
        for label, key in (("请求参数", "request"), ("返回参数", "response")):
            items = d[key]
            if not items:
                continue
            md.append("\n### %s\n\n| 参数名 | 类型 | 必填 | 说明 |\n| --- | --- | --- | --- |\n" % label)
            for it in items:
                md.append("| %s | %s | %s | %s |\n" % (it["name"], it["type"], it["required"], it["desc"]))
    open(os.path.join(CACHE, "fields.md"), "w", encoding="utf-8").write("".join(md))
    print("解析接口 %d 个，输出到 %s" % (len(all_rows), os.path.relpath(FIELDS, ROOT)))


if __name__ == "__main__":
    main()
