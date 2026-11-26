# API 文档

已完成以及未完成API列表汇总

如果有兴趣参与贡献，可以在具体的API表格后面标识自己为贡献者以及完成时间，例如：

|          名称           | 请求方式 | URL                        | 是否已实现 | 使用方法 |贡献者|完成时间|
| :---------------------: | -------- | :------------------------- | ---------- | -------- |-------- |-------- |
| 获取公众号类目       | GET      | /wxaapi/newtmpl/getcategory            | NO         |   |silenceper| 2021-12-20|

- [微信公众号](./officialaccount.md)
- [小程序](./miniprogram.md)
- [小游戏](./minigame.md)
- [开放平台](./oplatform.md)
- [微信支付](./wxpay.md)
- [企业微信](./work.md)
- [智能对话](./aispeech.md)

小程序服务端 API 的官方接口覆盖情况见 [小程序服务端 API 覆盖清单](./miniprogram-server-api-gap.md)
（机读版：[全量](./miniprogram-server-api-all.csv) / [缺口](./miniprogram-server-api-gap.csv)，
可用 `doc/api/tools/sync_miniprogram_api.py` 重新抓取官方文档刷新）。
