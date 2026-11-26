# 小程序服务端 API 覆盖清单
> 数据源：微信官方文档「服务端 API」全量目录（/miniprogram/dev/server/API/），抓取日期 2026-11-26。
> 匹配口径：官方文档「调用方式」里的请求地址，去掉 query 后与仓库 Go 源码里的接口路径字面量精确比对。

## 一、结论
| 指标 | 数量 |
| --- | ---: |
| 官方服务端接口总数 | 433 |
| 本仓库已实现 | 433 |
| **尚未实现（缺口）** | **0** |

完全覆盖的分类（36 个）：B2b 门店助手、openApi 管理、云开发、交易保障、付费管理、即时配送、图像处理与文字识别、城市服务、小程序 Soter、小程序安全、小程序客服、小程序登录、小程序直播、小程序码与小程序链接、小说作品管理、微信学生身份快速验证、微信客服、微信搜一搜、微信服务市场、微信物流服务、微信红包封面、接口调用凭证、插件管理、数据分析、消息相关、物流助手、生物认证 / 人脸核身、用工关系、用户信息、短剧媒资管理、短剧播放器、硬件设备、虚拟支付、订单发货、运维中心、附近小程序

### 按优先级
| 优先级 | 含义 | 缺口数 |
| --- | --- | ---: |
| P0 | 通用能力，谁都用得上，建议先补 | 0 |
| P1 | 垂直业务能力，按接入场景选做 | 0 |
| P2 | 企业级/大垂类，需求驱动 | 0 |

## 二、顶级分类覆盖总览
| 分类 | 官方接口数 | 已实现 | 缺口 | 优先级 |
| --- | ---: | ---: | ---: | --- |
| B2b 门店助手 | 34 | 34 | 0 | P2 |
| 虚拟支付 | 38 | 38 | 0 | P1 |
| 付费管理 | 2 | 2 | 0 | P0 |
| 城市服务 | 11 | 11 | 0 | P1 |
| 云开发 | 26 | 26 | 0 | P1 |
| 数据分析 | 11 | 11 | 0 | P2 |
| 短剧播放器 | 9 | 9 | 0 | P1 |
| 物流助手 | 22 | 22 | 0 | P1 |
| 生物认证 / 人脸核身 | 2 | 2 | 0 | P0 |
| 硬件设备 | 9 | 9 | 0 | P1 |
| 图像处理与文字识别 | 9 | 9 | 0 | P0 |
| 即时配送 | 15 | 15 | 0 | P1 |
| 小程序客服 | 14 | 14 | 0 | P0 |
| 微信客服 | 3 | 3 | 0 | P2 |
| 用工关系 | 2 | 2 | 0 | P2 |
| 小程序直播 | 38 | 38 | 0 | P2 |
| 短剧媒资管理 | 32 | 32 | 0 | P1 |
| 接口调用凭证 | 2 | 2 | 0 | P2 |
| 消息相关 | 13 | 13 | 0 | P1 |
| 附近小程序 | 4 | 4 | 0 | P0 |
| 小说作品管理 | 23 | 23 | 0 | P2 |
| openApi 管理 | 8 | 8 | 0 | P0 |
| 运维中心 | 10 | 10 | 0 | P0 |
| 订单发货 | 12 | 12 | 0 | P1 |
| 插件管理 | 2 | 2 | 0 | P0 |
| 小程序码与小程序链接 | 9 | 9 | 0 | P2 |
| 微信红包封面 | 1 | 1 | 0 | P2 |
| 小程序安全 | 3 | 3 | 0 | P2 |
| 小程序 Soter | 1 | 1 | 0 | P0 |
| 微信学生身份快速验证 | 1 | 1 | 0 | P1 |
| 交易保障 | 16 | 16 | 0 | P2 |
| 用户信息 | 5 | 5 | 0 | P0 |
| 小程序登录 | 3 | 3 | 0 | P0 |
| 微信物流服务 | 40 | 40 | 0 | P1 |
| 微信服务市场 | 2 | 2 | 0 | P0 |
| 微信搜一搜 | 1 | 1 | 0 | P0 |

## 三、二级分类覆盖总览
| 二级分类 | 官方接口数 | 已实现 | 缺口 | 优先级 |
| --- | ---: | ---: | ---: | --- |
| B2b/bill | 28 | 28 | 0 | P2 |
| B2b/notify | 2 | 2 | 0 | P2 |
| B2b/store_assistant | 4 | 4 | 0 | P2 |
| VirtualPayment | 38 | 38 | 0 | P1 |
| charge | 2 | 2 | 0 | P0 |
| cityservice/basic | 4 | 4 | 0 | P1 |
| cityservice/elderMedical | 6 | 6 | 0 | P1 |
| cityservice/medicalassistant | 1 | 1 | 0 | P1 |
| cloudbase/database | 13 | 13 | 0 | P1 |
| cloudbase/functions | 2 | 2 | 0 | P1 |
| cloudbase/others | 8 | 8 | 0 | P1 |
| cloudbase/storage | 3 | 3 | 0 | P1 |
| data-analysis/others | 5 | 5 | 0 | P2 |
| data-analysis/visit-retain | 3 | 3 | 0 | P2 |
| data-analysis/visit-trend | 3 | 3 | 0 | P2 |
| dramaOthersAPI | 9 | 9 | 0 | P1 |
| express/express-by-business | 12 | 12 | 0 | P1 |
| express/express-by-provider | 10 | 10 | 0 | P1 |
| face | 2 | 2 | 0 | P0 |
| hardware-device | 9 | 9 | 0 | P1 |
| img-ocr/img | 3 | 3 | 0 | P0 |
| img-ocr/ocr | 6 | 6 | 0 | P0 |
| immediate-delivery/deliver-by-business | 14 | 14 | 0 | P1 |
| immediate-delivery/deliver-by-provider | 1 | 1 | 0 | P1 |
| kf-mgnt/kf-management | 4 | 4 | 0 | P0 |
| kf-mgnt/kf-message | 10 | 10 | 0 | P0 |
| kf-work | 3 | 3 | 0 | P2 |
| laboruse | 2 | 2 | 0 | P2 |
| livebroadcast/commodity-management | 7 | 7 | 0 | P2 |
| livebroadcast/role-management | 3 | 3 | 0 | P2 |
| livebroadcast/studio-management | 26 | 26 | 0 | P2 |
| livebroadcast/subscribe-management | 2 | 2 | 0 | P2 |
| minidrama/auditdrama | 7 | 7 | 0 | P1 |
| minidrama/authorizeapp | 4 | 4 | 0 | P1 |
| minidrama/authorizecopyright | 4 | 4 | 0 | P1 |
| minidrama/authorizedrama | 4 | 4 | 0 | P1 |
| minidrama/usagedata | 3 | 3 | 0 | P1 |
| minidrama/vod_fileupload | 6 | 6 | 0 | P1 |
| minidrama/vod_media | 4 | 4 | 0 | P1 |
| mp-access-token | 2 | 2 | 0 | P2 |
| mp-message-management/subscribe-message | 10 | 10 | 0 | P1 |
| mp-message-management/updatable-message | 3 | 3 | 0 | P1 |
| nearby-poi | 4 | 4 | 0 | P0 |
| novel/auth | 6 | 6 | 0 | P2 |
| novel/business | 14 | 14 | 0 | P2 |
| novel/other | 1 | 1 | 0 | P2 |
| novel/preview | 2 | 2 | 0 | P2 |
| openApi-mgnt | 8 | 8 | 0 | P0 |
| operation | 10 | 10 | 0 | P0 |
| order_shipping | 12 | 12 | 0 | P1 |
| plugin-management | 2 | 2 | 0 | P0 |
| qrcode-link/qr-code | 3 | 3 | 0 | P2 |
| qrcode-link/short-link | 1 | 1 | 0 | P2 |
| qrcode-link/url-link | 2 | 2 | 0 | P2 |
| qrcode-link/url-scheme | 3 | 3 | 0 | P2 |
| red-packet-cover | 1 | 1 | 0 | P2 |
| sec-center/safety-control-capability | 1 | 1 | 0 | P2 |
| sec-center/sec-check | 2 | 2 | 0 | P2 |
| soter | 1 | 1 | 0 | P0 |
| student | 1 | 1 | 0 | P1 |
| transaction-guarantee/basic | 2 | 2 | 0 | P2 |
| transaction-guarantee/comment | 9 | 9 | 0 | P2 |
| transaction-guarantee/complaint | 5 | 5 | 0 | P2 |
| user-info/basic-info | 3 | 3 | 0 | P0 |
| user-info/internet | 1 | 1 | 0 | P0 |
| user-info/phone-number | 1 | 1 | 0 | P0 |
| user-login | 3 | 3 | 0 | P0 |
| weixin-express/express-msg | 4 | 4 | 0 | P1 |
| weixin-express/express-return | 3 | 3 | 0 | P1 |
| weixin-express/express-search | 4 | 4 | 0 | P1 |
| weixin-express/freight | 11 | 11 | 0 | P1 |
| weixin-express/msgpush | 2 | 2 | 0 | P1 |
| weixin-express/same_city_distribution | 16 | 16 | 0 | P1 |
| wx-service-market | 2 | 2 | 0 | P0 |
| wxsearch | 1 | 1 | 0 | P0 |

## 四、缺口明细
无。当前官方文档中列出的 433 个服务端接口，本仓库均已实现。

## 五、已实现接口在 SDK 中的位置
共 433 条已实现，按 SDK 包归类——新增接口时对着已有包写即可。
| SDK 包 | 已实现接口数 | 接口 |
| --- | ---: | --- |
| `miniprogram/express` | 38 | 传运单、充值保费退款、创建退货ID、创建配送单、创建门店、发货时投保、取消配送单、开通无忧退货、开通门店权限、拉取保单信息、拉取充值订单信息、拉取理赔摘要、推送已绑定物流轨迹信息、无忧退理赔、更新物品信息接口、更新物流信息、更新门店、查询开通状态、查询扣费主体、查询支持同城配送的城市、查询运单、查询运费、查询退货ID状态、查询配送单、查询门店、查运单、模拟回调接口、用户手机状态查询、申请充值订单号、申请支付、获取运力id列表、解绑退货ID、设置保费告警余额、设置扣费主体、门店余额查询、门店运费充值、门店运费流水查询、门店运费退款 |
| `miniprogram/livebroadcast` | 38 | 上下架商品、下载商品讲解视频、修改主播副号、修改直播间小助手、创建直播间、删除主播副号、删除商品、删除直播间、删除直播间商品、删除直播间小助手、发送直播开始事件、回放功能管理、官方收录管理、客服功能管理、导入商品、推送商品、撤回商品审核、更新商品、查询成员列表、查询直播间小助手、添加主播副号、添加并提审商品、添加管理直播间小助手、直播挂件获取全局KEY、直播挂件设置全局KEY、直播间商品排序、禁言管理、移除成员角色、编辑直播间、获取主播副号、获取商品列表、获取商品的信息与审核状态、获取直播间分享二维码、获取直播间列表和回放、获取直播间推流地址、获取长期订阅用户、设置成员角色、重新提交商品审核 |
| `miniprogram/virtualpayment` | 38 | 上传媒体文件、下载小程序账单、下载广告金对应的商户订单信息、下载支付订单、下载虚拟支付ios月结账单、代币支付退款、代币赠送、充值广告金、创建提现单、发起订阅扣款、启动批量发布道具任务、启动订单退款任务、商家解约、商户被管控原因查询、回复用户、完成投诉处理、扣减代币、批量上传道具、查询下载订单任务、查询代币余额、查询创建的订单、查询商家账户可提现余额、查询广告金充值记录、查询广告金充值账户、查询广告金发放记录、查询广告金回收记录、查询批量上传道具任务、查询批量发布道具任务、查询提现单、查询签约关系、查询苹果订单对应单号、绑定广告金充值账户、获取协商历史、获取微信支付反馈投诉图片的签名头部、获取投诉列表、获取投诉详情、通知已发货完成、预通知扣款 |
| `miniprogram/b2b` | 34 | 上传商户图片、全量授权门店查询、关闭订单、删除分账方、发起手动提现、商户号进件、完成分账、开通流程、微信支付自动提现接口、报名微信支付技术服务费优惠活动、报名银行转账技术服务费优惠活动、接口下载交易账单与资金账单、查询分账剩余金额、查询分账回退结果、查询分账方、查询分账结果、查询商户号开通状态、查询提现状态、查询订单、查询账户余额、查询退款、查询银行转账的技术服务费率、模板消息列表及下发、消息效果数据、添加分账方、申请开通银行转账、获取密钥AppKey、获取小程序下所有商户的信息、请求分账、请求分账回退、跳转银行转账页面、退款、门店信息查询、预录入门店信息 |
| `miniprogram/minidrama` | 31 | 上传分片、修改剧目基本信息、删除媒资、剧目提审、单个文件上传、增加剧目授权、增加版权授权、增加账号授权、拉取上传、提交替换剧集审核、替换审核通过的剧集、查询CDN日志下载链接列表、查询CDN用量数据、查询任务、查询剧目审核信息、查询授权信息、查询流量包详情、查询版权授权信息、查询被授权信息、查询被版权授权信息、查询账号授权信息、申请分片上传、确认上传、获取剧目信息、获取剧目列表、获取媒资列表、获取媒资播放链接、获取媒资详细信息、解除剧目授权、解除版权授权、解除账号授权 |
| `miniprogram/tcb` | 26 | 云开发上报接口、创建发短信任务、删除文件、删除集合、延时调用云函数、描述扩展上传文件信息、数据库删除记录、数据库导入、数据库导出、数据库插入记录、数据库更新记录、数据库查询记录、数据库聚合、数据库迁移状态查询、新增集合、更新数据库索引、查询短信记录、统计集合记录数、获取cloudID对应的数据、获取云开发数据、获取实时语音签名、获取文件上传链接、获取文件下载链接、获取腾讯云API调用凭证、获取集合信息、触发云函数 |
| `miniprogram/novel` | 23 | 上传章节、作品提审、创建作品、删除作品、删除指定的授权关系、删除指定的账号授权关系、删除章节、批量上传章节、新增账号-小说授权、新增账号-账号授权、替换章节、查看主授权关系列表-查看被授权关系列表、查看账号主授权关系列表、编辑作品、获取作品信息、获取作品列表、获取章节信息、获取章节列表、读后推荐、调整章节相对顺序、调整章节顺序、预览设置修改、预览设置获取 |
| `miniprogram/logistics` | 22 | 取消运单、对账单下载、批量获取运单数据、更新商户审核结果、更新运单轨迹、查询运单轨迹、模拟更新订单状态、生成运单、绑定/解绑物流账号、获取所有绑定的物流账号、获取打印员、获取支持的快递公司列表、获取电子面单余额、获取运单数据、获取面单联系人信息、运力公司退款、运力取消订单、运力更新订单状态、运力更新需支付的运费、运力返回用户投诉处理结果、配置面单打印员、预览面单模板 |
| `miniprogram/transaction` | 16 | 创建回复、创建评论、删除回复、删除评论、商家回应投诉、商家提交退款凭证、商家申诉、商家补充凭证、查询投诉单详情、查询评价列表、查询评价详情、查询评论列表、确认和解、获取交易保障标状态、获取小程序交易体验分违规记录、重置Api客服quota |
| `miniprogram/delivery` | 15 | 发起绑定请求、取消配送单、异常件退回商家确认、拉取已绑定账号、拉取配送单信息、更新配送单状态、模拟更新配送单状态、模拟配送公司更新配送单状态、添加小费、添加配送单、申请开通即时配送、获取已支持的配送公司列表、重新下单、预下配送单、预取消配送单 |
| `miniprogram/cityservice` | 13 | 仿原生跳转、公告草稿预览设置、查询所有公告、查询用户人脸核身真实验证结果、查询用户实名API、查询用户是否开通、校验实名信息、正式发布公告、消息推送接口、消息通路发消息、添加公告草稿、获取城市服务限定页面链接、获取用户人脸核身会话唯一标识 |
| `miniprogram/order` | 12 | 发货信息合单录入、发货信息录入、品牌申请、小程序交易类型变更申请、小程序品牌申请状态查询、查询小程序是否已完成交易结算管理确认、查询小程序是否已开通发货信息管理服务、查询订单列表、查询订单发货状态、消息跳转路径设置、特殊发货报备、确认收货提醒 |
| `miniprogram/analysis` | 11 | 获取小程序性能数据、获取小程序用户画像分布、获取用户小程序访问分布数据、获取用户访问小程序周留存、获取用户访问小程序数据周趋势、获取用户访问小程序数据日趋势、获取用户访问小程序数据月趋势、获取用户访问小程序数据概况、获取用户访问小程序日留存、获取用户访问小程序月留存、获取访问页面数据 |
| `miniprogram/kf` | 10 | 取消客服管理员、客服子商户创建商户、客服子商户拉取单个商户信息、客服子商户拉取多个商户信息、客服子商户更新商户信息、客服输入状态、查询绑定情况、绑定微信客服、解除绑定微信客服、设置客服管理员 |
| `miniprogram/operation` | 10 | 查询js错误详情、查询域名配置、查询实时日志、查询错误列表、获取mediaId图片、获取分阶段发布详情、获取客户端版本、获取性能数据、获取用户反馈列表、获取访问来源 |
| `miniprogram/subscribe` | 10 | 删除模板、发送订阅消息、更新服务卡片扩展信息、查询服务卡片状态、激活与更新服务卡片、获取已有模板列表、获取模板中的关键词、获取类目、获取类目下的公共模板、选用模板 |
| `miniprogram/dramaplayer` | 9 | 刷剧剧目设置、批处理短剧合作推广计划、推荐剧目、播放原始视频-推荐位控制接口、短剧上架、获取已上架短剧、获取短剧变现类型、获取短剧合作推广活动、设置短剧变现类型 |
| `miniprogram/hardware` | 9 | 创建设备组、发送设备消息、查询license资源包列表、查询设备激活详情、查询设备组信息、激活设备license、获取设备票据、设备组删除设备、设备组添加设备 |
| `miniprogram/ocr` | 9 | 二维码/条码识别、图片智能裁剪、图片高清化、营业执照识别、行驶证识别、身份证识别、通用印刷体识别、银行卡识别、驾驶证识别 |
| `miniprogram/auth` | 8 | 小程序登录凭证校验、支付后获取Unionid、检查加密信息、检验登录态、获取手机号、获取插件用户openpid、获取用户encryptKey、重置登录态 |
| `internal/openapi` | 6 | 使用AppSecret重置API调用次数、查询API调用额度、查询rid信息、网络通信检测、重置API调用次数、重置指定API调用次数 |
| `miniprogram/message` | 4 | 修改动态消息、修改小程序聊天工具的动态卡片消息、创建activity_id、发送客服消息 |
| `miniprogram/nearby` | 4 | 删除地点、查看地点、添加地点、设置展示状态 |
| `officialaccount/customerservice` | 4 | 删除客服账号、添加客服账号、获取在线客服列表、获取所有客服账号 |
| `miniprogram/qrcode` | 3 | 获取不限制的小程序码、获取小程序二维码、获取小程序码 |
| `miniprogram/urlscheme` | 3 | 查询scheme码、获取NFC的小程序scheme、获取加密scheme码 |
| `officialaccount/basic` | 3 | 获取微信API服务器IP、获取微信推送服务器IP、重置API调用次数 |
| `credential` | 2 | 获取接口调用凭据、获取稳定版接口调用凭据 |
| `miniprogram/charge` | 2 | 查询购买资源包的用量情况、获取小程序某个付费能力的最近用量数据 |
| `miniprogram/laboruse` | 2 | 推送用工消息、解绑用工关系 |
| `miniprogram/plugin` | 2 | 插件申请管理、插件管理 |
| `miniprogram/security` | 2 | 多媒体内容安全识别、文本内容安全识别 |
| `miniprogram/servicemarket` | 2 | 异步获取处理数据、调用服务市场接口 |
| `miniprogram/urllink` | 2 | 查询加密URLLink、获取加密URLLink |
| `officialaccount/material` | 2 | 新增临时素材、获取临时素材 |
| `miniprogram/business` | 1 | 获取手机号 |
| `miniprogram/content` | 1 | 文本内容安全识别 |
| `miniprogram/redpacketcover` | 1 | 获取微信红包封面 |
| `miniprogram/riskcontrol` | 1 | 获取用户安全等级 |
| `miniprogram/search` | 1 | 搜一搜数据推送 |
| `miniprogram/shortlink` | 1 | 获取ShortLink |
| `miniprogram/soter` | 1 | 生物认证秘钥签名验证 |
| `miniprogram/student` | 1 | 快速获取学生身份 |

## 六、复跑方式
```bash
python3 doc/api/tools/sync_miniprogram_api.py
```
产出：本文档 + `miniprogram-server-api-gap.csv`（缺口）+ `miniprogram-server-api-all.csv`（全量）。

> 注：官方文档「事件通知」（event_push，32 页）属于回调而非主动调用接口，未计入本次缺口；需要可另行整理。
