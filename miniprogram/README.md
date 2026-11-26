# 微信小程序

[官方文档](https://developers.weixin.qq.com/miniprogram/dev/framework/)

## 包说明

- analysis 数据分析相关 API
- auth 登录/用户信息相关 API
- b2b B2b 门店助手 API
- business 业务接口（手机号等）
- charge 付费管理 API
- cityservice 城市服务/人脸核身 API
- content 内容安全接口
- delivery 即时配送 API
- dramaplayer 短剧播放器 API
- encryptor 小程序加解密
- express 微信物流服务 API
- generic 通用 API 调用（访问 SDK 尚未封装的接口）
- hardware 硬件设备 API
- kf 小程序客服/微信客服 API
- laboruse 用工关系 API
- livebroadcast 小程序直播 API
- logistics 物流助手 API
- message 客服消息/动态消息/消息推送接收器
- minidrama 短剧媒资管理 API
- nearby 附近小程序 API
- novel 小说作品管理 API
- ocr 图像处理与文字识别 API
- operation 运维中心 API
- order 订单发货 API
- plugin 插件管理 API
- privacy 小程序隐私协议 API
- qrcode 小程序码 API
- redpacketcover 微信红包封面 API
- riskcontrol 安全风控 API
- search 微信搜一搜 API
- security 内容安全（异步检测） API
- servicemarket 微信服务市场 API
- shortlink 小程序短链 API
- soter 小程序 Soter API
- student 学生身份核验 API
- subscribe 订阅消息/服务卡片 API
- tcb 云开发 API
- transaction 交易保障 API
- urllink URL Link API
- urlscheme URL Scheme API
- virtualpayment 小程序虚拟支付 API
- werun 微信运动 API

## 快速入门

```go
wc := wechat.NewWechat()
memory := cache.NewMemory()
cfg := &miniConfig.Config{
    AppID:     "xxx",
    AppSecret: "xxx",
    Cache: memory,
}
miniprogram := wc.GetMiniProgram(cfg)
miniprogram.GetAnalysis().GetAnalysisDailyRetain()
```

### 小程序虚拟支付 
#### `注意：需要传入 Appkey、OfferID 的值`
相关文档：[小程序虚拟支付](https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/virtual-payment.html)
```go
wc := wechat.NewWechat()
miniprogram := wc.GetMiniProgram(&miniConfig.Config{
    AppID:     "xxx",
    AppSecret: "xxx",
    AppKey:    "xxx",
    OfferID:   "xxx",
    Cache: cache.NewRedis(&redis.Options{
        Addr: "",
    }),
})
virtualPayment := miniprogram.GetVirtualPayment()
virtualPayment.SetSessionKey("xxx")
// 查询用户余额
var (
    res *virtualPayment.QueryUserBalanceResponse
    err error
)

if res, err = virtualPayment.QueryUserBalance(context.TODO(), &virtualPayment.QueryUserBalanceRequest{
    OpenID: "xxx",
    Env: virtualPayment.EnvProduction,
    UserIP: "xxx",
}); err != nil {
    panic(err)
}

```