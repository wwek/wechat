"""解码回归：把官方文档的返回示例喂进本仓库真实的响应结构体。

前置条件：先用 sync_miniprogram_api.py / extract_api_fields.py 抓取官方文档页面，
缓存在 .local/apidoc/pages 下（该目录不入库）。

用法:
    python3 doc/api/tools/decode_regress.py     # 生成并运行 .local/regress/main.go
    go run ./.local/regress/                    # 单独跑回归

输出：每个官方返回示例的解码结果，任何类型不匹配都会导致该接口失败。
"""
import html
import json
import os
import re
import subprocess
import sys

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
from jsjson import extract_example, first_json_value, js_to_json  # noqa: E402

ROOT = os.path.abspath(os.path.join(os.path.dirname(__file__), "..", ".."))
MODULE = "github.com/silenceper/wechat/v2"

METHOD_RE = re.compile(
    r"func \([^)]*\) ([A-Z]\w*)\([^)]*\)\s*\((?:out\s+)?(\*?)([A-Z]\w*(?:Response|Res))\b")


def collect():
    gofiles = [f for f in subprocess.check_output(
        ["git", "ls-files", "*.go"], text=True).split()
        if f.startswith("miniprogram/") or f.startswith("internal/")]
    const_url = {}
    method_src = {}
    for f in gofiles:
        src = open(f, encoding="utf-8", errors="ignore").read()
        pkg = f.rsplit("/", 1)[0] if "/" in f else "."
        for m in re.finditer(r'(\w+URL)\s*=\s*"(https://api\.weixin\.qq\.com[^"]*)"', src):
            u = re.sub(r"\?.*$", "", m.group(2)).replace(
                "https://api.weixin.qq.com", "").rstrip("/")
            const_url[(pkg, m.group(1))] = u
        method_src.setdefault(pkg, []).append(src)

    methods = []
    for pkg, srcs in method_src.items():
        for src in srcs:
            for m in METHOD_RE.finditer(src):
                methods.append((pkg, m.group(1), m.group(3), src))

    bound = []
    for pkg, name, rettype, src in methods:
        mm = re.search(r"func \([^)]*\) %s\(.*?\n\}\n" % re.escape(name), src, re.S)
        if not mm:
            continue
        body = mm.group(0)
        urls = [const_url[(pkg, c)] for c in re.findall(r"\b(\w+URL)\b", body)
                if (pkg, c) in const_url]
        if urls:
            bound.append((pkg, name, rettype, urls[0]))
    return bound


def main():
    bound = collect()
    doc = json.load(open(os.path.join(ROOT, ".local/apidoc/apis_final.json")))
    url2page = {r["url"]: "%s__%s" % (r["cat"].replace("/", "__"), r["slug"])
                for r in doc if r["url"]}
    pages = {}
    pdir = os.path.join(ROOT, ".local/apidoc/pages")
    for f in sorted(os.listdir(pdir)):
        if not f.endswith(".html"):
            continue
        raw = open(os.path.join(pdir, f), encoding="utf-8", errors="ignore").read()
        raw_text = re.sub(r"<[^>]+>", "", raw)
        # 两种预处理都试一遍：一种还原实体、一种保留原样，谁先解析成功用谁
        text_a = raw_text
        for ent, ch in (("&quot;", '"'), ("&#39;", "'"), ("&amp;", "&")):
            text_a = text_a.replace(ent, ch)
        seg = extract_example(text_a) or extract_example(raw_text)
        val = first_json_value(seg) if seg else ""
        if val:
            pages[f[:-5]] = val
            raw_seg = extract_example(raw_text)
            if raw_seg:
                pages[f[:-5] + "@raw"] = first_json_value(raw_seg)

    imports, cases, skipped = set(), [], []
    for pkg, name, rettype, url in bound:
        page = url2page.get(url)
        if not page or page not in pages:
            skipped.append((name, "官方页面无返回示例"))
            continue
        payload = js_to_json(pages[page])
        try:
            json.loads(payload)
        except Exception:  # noqa: BLE001
            # 示例里可能有多段 JSON，只取第一段完整的
            val = first_json_value(payload)
            if val:
                try:
                    json.loads(val)
                    payload = val
                except Exception:  # noqa: BLE001
                    pass
        try:
            json.loads(payload)
        except Exception as e:  # noqa: BLE001
            alt = js_to_json(pages.get(page + "@raw", ""))
            try:
                json.loads(alt)
                payload = alt
            except Exception:  # noqa: BLE001
                skipped.append((name, "示例非合法 JSON: %s" % str(e)[:50]))
                continue
        alias = "pkg" + re.sub(r"[^A-Za-z0-9]", "", pkg.replace("/", "_"))
        imports.add((alias, pkg))
        cases.append((name, alias, rettype, payload))

    lines = ["package main", "", "import (", '\t"encoding/json"', '\t"fmt"', ""]
    for alias, pkg in sorted(imports):
        lines.append('\t%s "%s/%s"' % (alias, MODULE, pkg))
    lines += [")", "", "type checkCase struct {", "\tname   string",
              "\tdecode func([]byte) error", "}", "", "var cases = []checkCase{"]
    for name, alias, rettype, payload in cases:
        lines += ["\t{", '\t\tname: "%s",' % name,
                  "\t\tdecode: func(b []byte) error {",
                  "\t\t\tvar v %s.%s" % (alias, rettype),
                  "\t\t\tif err := json.Unmarshal(b, &v); err != nil {",
                  "\t\t\t\treturn err", "\t\t\t}", "\t\t\treturn nil", "\t\t},",
                  "\t},"]
    lines += ["}", "", "func main() {", "\tdata := map[string]string{"]
    for name, alias, rettype, payload in cases:
        lines.append('\t\t"%s": %s,' % (name, json.dumps(payload)))
    lines += ["\t}", "\tfail := 0", "\tfor _, c := range cases {",
              "\t\tif err := c.decode([]byte(data[c.name])); err != nil {",
              '\t\t\tfmt.Printf("FAIL %-38s %v\\n", c.name, err)',
              "\t\t\tfail++", "\t\t}", "\t}",
              '\tfmt.Printf("总计 %d 个接口，失败 %d 个\\n", len(cases), fail)', "}"]
    out = os.path.join(ROOT, ".local/regress/main.go")
    open(out, "w").write("\n".join(lines))
    json.dump(skipped, open(os.path.join(ROOT, ".local/regress/skipped.json"), "w"),
              ensure_ascii=False, indent=1)
    print("可回归 %d 个，跳过 %d 个" % (len(cases), len(skipped)))


if __name__ == "__main__":
    main()
