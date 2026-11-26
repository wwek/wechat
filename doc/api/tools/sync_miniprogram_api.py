#!/usr/bin/env python3
"""同步微信小程序服务端 API 官方文档，并比对本仓库的实现情况。

用法:
    python3 doc/api/tools/sync_miniprogram_api.py            # 全量刷新 + 比对 + 生成报告
    python3 doc/api/tools/sync_miniprogram_api.py --offline  # 只用 .local/apidoc/pages 下已缓存页面

产物:
    doc/api/miniprogram-server-api-gap.md    人读清单
    doc/api/miniprogram-server-api-gap.csv   机读清单（缺口）
    doc/api/miniprogram-server-api-all.csv    全量 433 个官方接口 + 实现状态

中间产物（不入库）:
    .local/apidoc/pages/                     官方文档页面缓存
"""

import argparse
import collections
import csv
import html
import json
import os
import re
import subprocess
import sys
import urllib.request
from concurrent.futures import ThreadPoolExecutor

ROOT = os.path.abspath(os.path.join(os.path.dirname(__file__), "..", "..", ".."))
CACHE = os.path.join(ROOT, ".local", "apidoc")
PAGES = os.path.join(CACHE, "pages")
BASE = "https://developers.weixin.qq.com"
SEED = BASE + "/miniprogram/dev/server/API/mp-access-token/"
API_PREFIX = "/miniprogram/dev/server/API/"

# ---------------------------------------------------------------- 抓取

def http_get(url, binary=False):
    req = urllib.request.Request(url, headers={"User-Agent": "Mozilla/5.0"})
    with urllib.request.urlopen(req, timeout=30) as r:
        data = r.read()
    return data if binary else data.decode("utf-8", "ignore")


def discover_pages():
    """从种子页的侧边栏取出服务端 API 全树。"""
    htmltext = http_get(SEED)
    links = set(re.findall(re.escape(API_PREFIX) + r'[^"#?]*\.html', htmltext))
    return sorted(links)


def download(pages):
    os.makedirs(PAGES, exist_ok=True)

    def one(p):
        name = p[len(API_PREFIX):].replace("/", "__")
        dest = os.path.join(PAGES, name)
        if os.path.exists(dest) and os.path.getsize(dest) > 0:
            return True
        try:
            with open(dest, "wb") as f:
                f.write(http_get(BASE + p, binary=True))
            return True
        except Exception as e:  # noqa: BLE001
            print("  ! 下载失败 %s: %s" % (p, e), file=sys.stderr)
            return False

    with ThreadPoolExecutor(max_workers=12) as ex:
        ok = list(ex.map(one, pages))
    print("下载 %d/%d 页" % (sum(ok), len(ok)))
    return [p for p, o in zip(pages, ok) if o]


# ---------------------------------------------------------------- 解析

CALL_RE = re.compile(r'(GET|POST)\s+(?:[A-Za-z]+\s+)?(https://api\.weixin\.qq\.com(/[A-Za-z0-9_./-]*))')
DATE_RE = re.compile(r'最后一次编辑于\s*(\d{4}-\d{2}-\d{2})')
EN_RE = re.compile(r'接口英文名[：:]\s*([A-Za-z0-9_]+)')
TITLE_RE = re.compile(r'<title>(.*?)</title>', re.S)


def parse_pages():
    rows = []
    for fn in sorted(os.listdir(PAGES)):
        path = os.path.join(PAGES, fn)
        if not os.path.getsize(path):
            continue
        raw = open(path, encoding="utf-8", errors="ignore").read()
        t = TITLE_RE.search(raw)
        title = html.unescape(t.group(1)).split(" - ")[0].strip() if t else ""
        m = CALL_RE.search(raw)
        dm = DATE_RE.search(raw)
        em = EN_RE.search(raw)
        seg = fn[:-5].split("__")
        rows.append({
            "cat": "/".join(seg[:-1]),
            "slug": seg[-1].replace(".html", ""),
            "title": title,
            "en": em.group(1) if em else "",
            "method": m.group(1) if m else "",
            "url": re.sub(r"\?.*$", "", m.group(3)).rstrip("/") if m else "",
            "updated": dm.group(1) if dm else "",
            "doc": BASE + API_PREFIX + "/".join(seg[:-1]) + "/" + seg[-1],
        })
    return rows


# ---------------------------------------------------------------- 仓库扫描

SKIP_DIR = (".git", ".worktrees", ".local", "vendor")


def repo_endpoints():
    """本仓库出现的所有微信接口路径字面量。"""
    eps = set()
    for dp, _, fs in os.walk(ROOT):
        rel = os.path.relpath(dp, ROOT)
        if rel == ".":
            rel = ""
        if any(rel == s or rel.startswith(s + os.sep) for s in SKIP_DIR):
            continue
        for f in fs:
            if not f.endswith(".go"):
                continue
            src = open(os.path.join(dp, f), encoding="utf-8", errors="ignore").read()
            for m in re.finditer(r'https://api\.weixin\.qq\.com(/[A-Za-z0-9_./-]+)', src):
                eps.add(re.sub(r"\?.*$", "", m.group(1)).rstrip("/"))
            for m in re.finditer(r'"(/[a-z0-9][A-Za-z0-9_.-]*(?:/[A-Za-z0-9_.-]+)+)"', src):
                eps.add(m.group(1).rstrip("/"))
    return eps


# ---------------------------------------------------------------- 报告

CAT_LABEL = {
    "mp-access-token": "接口调用凭证", "openApi-mgnt": "openApi 管理", "user-info": "用户信息",
    "user-login": "小程序登录", "qrcode-link": "小程序码与小程序链接", "kf-mgnt": "小程序客服",
    "kf-work": "微信客服", "mp-message-management": "消息相关", "laboruse": "用工关系",
    "sec-center": "小程序安全", "data-analysis": "数据分析", "hardware-device": "硬件设备",
    "operation": "运维中心", "plugin-management": "插件管理", "cloudbase": "云开发",
    "express": "物流助手", "immediate-delivery": "即时配送", "weixin-express": "微信物流服务",
    "order_shipping": "订单发货", "cityservice": "城市服务", "nearby-poi": "附近小程序",
    "livebroadcast": "小程序直播", "minidrama": "短剧媒资管理", "dramaOthersAPI": "短剧播放器",
    "novel": "小说作品管理", "student": "微信学生身份快速验证", "transaction-guarantee": "交易保障",
    "charge": "付费管理", "B2b": "B2b 门店助手", "wx-service-market": "微信服务市场",
    "img-ocr": "图像处理与文字识别", "red-packet-cover": "微信红包封面",
    "face": "生物认证 / 人脸核身", "wxsearch": "微信搜一搜", "VirtualPayment": "虚拟支付",
    "soter": "小程序 Soter", "dramaOthers": "短剧",
}

# 建议优先级：通用能力优先，垂直/企业级后置
P0_CATS = {"openApi-mgnt", "operation", "plugin-management", "kf-mgnt", "img-ocr",
           "face", "soter", "wxsearch", "charge", "nearby-poi", "wx-service-market",
           "user-info", "user-login"}
P1_CATS = {"express", "weixin-express", "immediate-delivery", "hardware-device",
           "minidrama", "dramaOthersAPI", "VirtualPayment", "order_shipping",
           "cityservice", "cloudbase", "student", "mp-message-management"}
P2_CATS = {"livebroadcast", "B2b", "novel", "transaction-guarantee"}


def priority(cat):
    if cat in P0_CATS:
        return "P0"
    if cat in P1_CATS:
        return "P1"
    return "P2"


def attribute_packages(rows):
    """把已实现接口归属到 SDK 包：按端点路径字面量出现在哪个包的源码里。"""
    literal_index = collections.defaultdict(set)
    for dp, _, fs in os.walk(ROOT):
        rel = os.path.relpath(dp, ROOT)
        rel = "" if rel == "." else rel
        if any(rel == s or rel.startswith(s + os.sep) for s in SKIP_DIR):
            continue
        for f in fs:
            if not f.endswith(".go"):
                continue
            pkg = rel or "."
            src = open(os.path.join(dp, f), encoding="utf-8", errors="ignore").read()
            for m in re.finditer(r'https://api\.weixin\.qq\.com(/[A-Za-z0-9_./-]+)', src):
                literal_index[re.sub(r"\?.*$", "", m.group(1)).rstrip("/")].add(pkg)
            for m in re.finditer(r'"(/[a-z0-9][A-Za-z0-9_.-]*(?:/[A-Za-z0-9_.-]+)+)"', src):
                literal_index[m.group(1).rstrip("/")].add(pkg)
    out = {}
    for r in rows:
        if r.get("status") != "implemented":
            continue
        pkgs = literal_index.get(r["url"]) or set()
        if not pkgs:
            continue
        mp = sorted(p for p in pkgs if p.startswith("miniprogram"))
        out[r["url"]] = mp if mp else sorted(pkgs)
    return out


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--offline", action="store_true", help="只用缓存页面，不联网")
    args = ap.parse_args()

    os.makedirs(CACHE, exist_ok=True)
    pages = []
    if args.offline:
        for fn in sorted(os.listdir(PAGES)):
            seg = fn[:-5].split("__")
            pages.append(API_PREFIX + "/".join(seg))
    else:
        pages = discover_pages()
        print("官方文档接口页：%d 个" % len(pages))
        download(pages)

    rows = parse_pages()
    eps = repo_endpoints()
    for r in rows:
        r["status"] = "implemented" if r["url"] and r["url"] in eps else "missing"

    missing = [r for r in rows if r["status"] == "missing"]
    implemented = [r for r in rows if r["status"] == "implemented"]

    # ---- 机读 CSV
    allcsv = os.path.join(ROOT, "doc", "api", "miniprogram-server-api-all.csv")
    with open(allcsv, "w", newline="", encoding="utf-8-sig") as f:
        w = csv.writer(f)
        w.writerow(["分类", "接口名称", "接口英文名", "方法", "请求地址", "实现状态", "文档最后编辑", "文档链接"])
        for r in sorted(rows, key=lambda x: (x["cat"], x["title"])):
            w.writerow([r["cat"], r["title"], r["en"], r["method"], r["url"],
                        "已实现" if r["status"] == "implemented" else "未实现",
                        r["updated"], r["doc"]])

    gapcsv = os.path.join(ROOT, "doc", "api", "miniprogram-server-api-gap.csv")
    with open(gapcsv, "w", newline="", encoding="utf-8-sig") as f:
        w = csv.writer(f)
        w.writerow(["优先级", "分类", "接口名称", "接口英文名", "方法", "请求地址", "文档最后编辑", "文档链接"])
        for r in sorted(missing, key=lambda x: (priority(x["cat"].split("/")[0]), x["cat"], x["title"])):
            w.writerow([priority(r["cat"].split("/")[0]), r["cat"], r["title"], r["en"], r["method"],
                        r["url"], r["updated"], r["doc"]])

    # ---- 分类覆盖总览
    total_by_cat = collections.Counter(r["cat"] for r in rows)
    miss_by_cat = collections.Counter(r["cat"] for r in missing)
    impl_by_cat = collections.Counter(r["cat"] for r in implemented)
    top = lambda c: c.split("/")[0]
    total_by_top = collections.Counter(top(r["cat"]) for r in rows)
    miss_by_top = collections.Counter(top(r["cat"]) for r in missing)
    impl_by_top = collections.Counter(top(r["cat"]) for r in implemented)

    # ---- 已实现接口归属到 SDK 包
    pkg_of = attribute_packages(rows)

    # ---- Markdown
    out = []
    title = "小程序服务端 API 覆盖清单" if not missing else "小程序服务端 API 缺口清单"
    out.append("# %s\n" % title)
    out.append("> 数据源：微信官方文档「服务端 API」全量目录（%s），抓取日期 %s。\n"
               % (API_PREFIX, subprocess.check_output(["date", "+%Y-%m-%d"], text=True).strip()))
    out.append("> 匹配口径：官方文档「调用方式」里的请求地址，去掉 query 后与仓库 Go 源码里的接口路径字面量精确比对。\n")
    out.append("\n## 一、结论\n")
    out.append("| 指标 | 数量 |\n| --- | ---: |\n")
    out.append("| 官方服务端接口总数 | %d |\n" % len(rows))
    out.append("| 本仓库已实现 | %d |\n" % len(implemented))
    out.append("| **尚未实现（缺口）** | **%d** |\n" % len(missing))
    covered = sorted({CAT_LABEL.get(top(c), top(c)) for c in total_by_cat if not miss_by_cat.get(c)})
    out.append("\n完全覆盖的分类（%d 个）：%s\n" % (len(covered), "、".join(covered)))
    out.append("\n### 按优先级\n")
    out.append("| 优先级 | 含义 | 缺口数 |\n| --- | --- | ---: |\n")
    pc = collections.Counter(priority(top(r["cat"])) for r in missing)
    out.append("| P0 | 通用能力，谁都用得上，建议先补 | %d |\n" % pc["P0"])
    out.append("| P1 | 垂直业务能力，按接入场景选做 | %d |\n" % pc["P1"])
    out.append("| P2 | 企业级/大垂类，需求驱动 | %d |\n" % pc["P2"])

    out.append("\n## 二、顶级分类覆盖总览\n")
    out.append("| 分类 | 官方接口数 | 已实现 | 缺口 | 优先级 |\n| --- | ---: | ---: | ---: | --- |\n")
    for cat in sorted(total_by_top, key=lambda c: (-miss_by_top.get(c, 0), c)):
        out.append("| %s | %d | %d | %d | %s |\n" % (
            CAT_LABEL.get(cat, cat), total_by_top[cat], impl_by_top.get(cat, 0),
            miss_by_top.get(cat, 0), priority(cat)))

    out.append("\n## 三、二级分类覆盖总览\n")
    out.append("| 二级分类 | 官方接口数 | 已实现 | 缺口 | 优先级 |\n| --- | ---: | ---: | ---: | --- |\n")
    for cat in sorted(total_by_cat, key=lambda c: (-miss_by_cat.get(c, 0), c)):
        out.append("| %s | %d | %d | %d | %s |\n" % (
            cat, total_by_cat[cat], impl_by_cat.get(cat, 0),
            miss_by_cat.get(cat, 0), priority(top(cat))))

    out.append("\n## 四、缺口明细\n")
    if not missing:
        out.append("无。当前官方文档中列出的 %d 个服务端接口，本仓库均已实现。\n" % len(rows))
    else:
        out.append("共 %d 条，按优先级、再按分类排序。\n" % len(missing))
    for p in ("P0", "P1", "P2"):
        sub = [r for r in missing if priority(top(r["cat"])) == p]
        if not sub:
            continue
        out.append("\n### %s（%d 条）\n\n" % (p, len(sub)))
        for cat in sorted({r["cat"] for r in sub}):
            out.append("\n#### %s\n\n" % cat)
            out.append("| 接口名称 | 英文名 | 方法 | 请求地址 | 文档最后编辑 |\n| --- | --- | --- | --- | --- |\n")
            for r in sorted([x for x in sub if x["cat"] == cat], key=lambda x: x["title"]):
                out.append("| %s | `%s` | %s | `%s` | %s |\n" % (
                    r["title"], r["en"], r["method"], r["url"], r["updated"]))

    out.append("\n## 五、已实现接口在 SDK 中的位置\n")
    out.append("共 %d 条已实现，按 SDK 包归类——新增接口时对着已有包写即可。\n" % len(implemented))
    bypkg = collections.defaultdict(list)
    for r in implemented:
        for p in (pkg_of.get(r["url"]) or ["(未定位)"]):
            bypkg[p].append(r["title"])
    out.append("| SDK 包 | 已实现接口数 | 接口 |\n| --- | ---: | --- |\n")
    for pkg in sorted(bypkg, key=lambda p: (-len(bypkg[p]), p)):
        titles = sorted(set(bypkg[pkg]))
        out.append("| `%s` | %d | %s |\n" % (pkg, len(titles), "、".join(titles)))

    out.append("\n## 六、复跑方式\n")
    out.append("```bash\npython3 doc/api/tools/sync_miniprogram_api.py\n```\n")
    out.append("产出：本文档 + `miniprogram-server-api-gap.csv`（缺口）+ `miniprogram-server-api-all.csv`（全量）。\n")
    out.append("\n> 注：官方文档「事件通知」（event_push，32 页）属于回调而非主动调用接口，未计入本次缺口；需要可另行整理。\n")

    md = os.path.join(ROOT, "doc", "api", "miniprogram-server-api-gap.md")
    with open(md, "w", encoding="utf-8") as f:
        f.write("".join(out))

    print("官方 %d / 已实现 %d / 缺口 %d" % (len(rows), len(implemented), len(missing)))
    print("报告:", os.path.relpath(md, ROOT))
    json.dump(rows, open(os.path.join(CACHE, "apis_final.json"), "w"), ensure_ascii=False, indent=1)


if __name__ == "__main__":
    main()
