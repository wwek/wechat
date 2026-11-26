# 小程序服务端 API 覆盖与校验工具

| 脚本 | 作用 |
| --- | --- |
| `sync_miniprogram_api.py` | 抓取微信官方文档「服务端 API」全量目录，与仓库实现比对，生成覆盖清单与 CSV |
| `extract_api_fields.py` | 从缓存的官方文档页面里抽取请求/返回参数表（含类型、必填、说明、示例） |
| `decode_regress.py` + `jsjson.py` | 解码回归：把官方「返回示例」喂进仓库真实的响应结构体，暴露字段类型/层级不匹配 |

前置：先跑 `sync_miniprogram_api.py` 与 `extract_api_fields.py`，页面缓存放在 `.local/apidoc/`（不入库）。

说明：微信官方文档中同一字段在「参数表」与「返回示例」里类型常有不一致（部分示例还是手写伪 JS）。
实现侧对这类字段使用 `util.FlexString` / `util.FlexInt64` / `util.FlexBool` 容错，
并对对象/数组两种形态并存的结构体单独实现 `UnmarshalJSON`，避免线上返回与文档不一致时整个响应解码失败。
