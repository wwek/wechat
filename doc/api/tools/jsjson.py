"""把官方文档里的“返回示例”（多为 JS 字面量：单引号、无引号键、注释）转成合法 JSON。"""
import re


def js_to_json(src: str) -> str:
    # 官方示例里混有零宽空格、不换行空格
    src = src.replace("\u200b", "").replace("\u00a0", " ")
    out = []
    i = 0
    n = len(src)
    while i < n:
        c = src[i]
        # 行注释
        if c == "/" and i + 1 < n and src[i + 1] == "/":
            j = src.find("\n", i)
            i = n if j < 0 else j
            continue
        # 块注释
        if c == "/" and i + 1 < n and src[i + 1] == "*":
            j = src.find("*/", i)
            i = n if j < 0 else j + 2
            continue
        # 字符串
        if c in "\"'":
            quote = c
            out.append('"')
            i += 1
            while i < n:
                ch = src[i]
                if ch == "\\" and i + 1 < n:
                    nxt = src[i + 1]
                    if quote == "'":
                        out.append({'"': '\\"', "'": "'", "n": "\\n", "t": "\\t", "r": "\\r"}.get(nxt, "\\" + nxt))
                    else:
                        out.append(ch + nxt)
                    i += 2
                    continue
                if ch == quote:
                    out.append('"')
                    i += 1
                    break
                if ch == '"':
                    out.append('\\"')
                    i += 1
                    continue
                out.append(ch)
                i += 1
            continue
        # 无引号的键
        if c.isalpha() or c == "_":
            m = re.match(r"[A-Za-z_$][A-Za-z0-9_$]*", src[i:])
            if m:
                word = m.group(0)
                j = i + len(word)
                while j < n and src[j] in " \t":
                    j += 1
                if j < n and src[j] == ":":
                    if word in ("true", "false", "null"):
                        out.append(word)
                    else:
                        out.append('"%s"' % word)
                    i = i + len(word)
                    continue
                if word in ("true", "false", "null"):
                    out.append(word)
                else:
                    out.append('"%s"' % word)
                i += len(word)
                continue
        # 数字（示例里偶尔在数字后多写一个引号，如 payTime: 123123",）
        if c.isdigit() or (c == "-" and i + 1 < n and src[i + 1].isdigit()):
            m = re.match(r"-?\d+(?:\.\d+)?", src[i:])
            if m:
                out.append(m.group(0))
                i += len(m.group(0))
                if i < n and src[i] in "\"'":
                    i += 1  # 跳过多余的引号
                continue
        # 尾逗号
        if c == ",":
            j = i + 1
            while j < n and src[j] in " \t\r\n":
                j += 1
            if j < n and src[j] in "}]":
                i += 1
                continue
        out.append(c)
        i += 1
    txt = "".join(out)
    # 示例里常有裸控制字符（换行等）出现在字符串内部，JSON 不允许，去掉即可
    txt = "".join(ch for ch in txt if ch >= " " or ch == "\t")
    txt = re.sub(r",\s*([}\]])", r"\1", txt)
    # 修复对象/数组里被注释截断后残留的裸标识符（如 elide: true）
    txt = re.sub(r'([{,]\s*)([A-Za-z_$][A-Za-z0-9_$]*)\s*:', r'\1"\2":', txt)
    return txt


def extract_example(text: str) -> str:
    """取“返回示例”后的 JSON/JS 片段。"""
    i = text.find("返回示例")
    if i < 0:
        return ""
    seg = text[i + 4:]
    for stop in ("错误码", "适用范围", "请求示例", "注意事项"):
        j = seg.find(stop)
        if j > 0:
            seg = seg[:j]
    seg = seg.strip().strip("`").strip()
    return seg


def first_json_value(src: str) -> str:
    """从文本中取出第一个完整的 JSON 对象/数组（括号配平，跳过字符串内的括号）。"""
    start = None
    for i, ch in enumerate(src):
        if ch in "{[":
            start = i
            break
    if start is None:
        return ""
    open_ch = src[start]
    close_ch = "}" if open_ch == "{" else "]"
    depth = 0
    in_str = False
    quote = ""
    for i in range(start, len(src)):
        ch = src[i]
        if in_str:
            if ch == "\\":
                continue
            if ch == quote:
                in_str = False
            continue
        if ch in "\"'":
            in_str = True
            quote = ch
            continue
        if ch == open_ch:
            depth += 1
        elif ch == close_ch:
            depth -= 1
            if depth == 0:
                return src[start:i + 1]
    return src[start:]
