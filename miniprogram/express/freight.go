package express

import (
	"context"
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// 运费险（无忧退货）API URL
const (
	// 充值保费退款
	insuranceFreightRefundURL = "https://api.weixin.qq.com/wxa/business/insurance_freight/refund?access_token=%s"
	// 发货时投保
	insuranceFreightCreateOrderURL = "https://api.weixin.qq.com/wxa/business/insurance_freight/createorder?access_token=%s"
	// 开通无忧退货
	insuranceFreightOpenURL = "https://api.weixin.qq.com/wxa/business/insurance_freight/open?access_token=%s"
	// 拉取保单信息
	insuranceFreightGetOrderListURL = "https://api.weixin.qq.com/wxa/business/insurance_freight/getorderlist?access_token=%s"
	// 拉取充值订单信息
	insuranceFreightGetPayOrderListURL = "https://api.weixin.qq.com/wxa/business/insurance_freight/getpayorderlist?access_token=%s"
	// 拉取理赔摘要
	insuranceFreightGetSummaryURL = "https://api.weixin.qq.com/wxa/business/insurance_freight/getsummary?access_token=%s"
	// 无忧退理赔
	insuranceFreightClaimURL = "https://api.weixin.qq.com/wxa/business/insurance_freight/claim?access_token=%s"
	// 查询开通状态
	insuranceFreightQueryOpenURL = "https://api.weixin.qq.com/wxa/business/insurance_freight/query_open?access_token=%s"
	// 申请充值订单号
	insuranceFreightCreateChargeIDURL = "https://api.weixin.qq.com/wxa/business/insurance_freight/createchargeid?access_token=%s"
	// 申请支付
	insuranceFreightApplyPayURL = "https://api.weixin.qq.com/wxa/business/insurance_freight/applypay?access_token=%s"
	// 设置保费告警余额
	insuranceFreightUpdateNotifyFundsURL = "https://api.weixin.qq.com/wxa/business/insurance_freight/update_notify_funds?access_token=%s"
)

// InsuranceFreightRefund 充值保费退款
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/express/business/insurance_freight.html
func (express *Express) InsuranceFreightRefund(ctx context.Context) (res InsuranceFreightRefundResponse, err error) {
	accessToken, err := express.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(insuranceFreightRefundURL, accessToken)
	response, err := util.PostJSONContext(ctx, uri, map[string]interface{}{})
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &res, "InsuranceFreightRefund")
	return
}

// InsuranceFreightCreateOrder 发货时投保
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/express/business/insurance_freight.html
func (express *Express) InsuranceFreightCreateOrder(ctx context.Context, req *InsuranceFreightCreateOrderRequest) (res InsuranceFreightCreateOrderResponse, err error) {
	accessToken, err := express.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(insuranceFreightCreateOrderURL, accessToken)
	response, err := util.PostJSONContext(ctx, uri, req)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &res, "InsuranceFreightCreateOrder")
	return
}

// InsuranceFreightOpen 开通无忧退货
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/express/business/insurance_freight.html
func (express *Express) InsuranceFreightOpen(ctx context.Context) (res InsuranceFreightOpenResponse, err error) {
	accessToken, err := express.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(insuranceFreightOpenURL, accessToken)
	response, err := util.HTTPGetContext(ctx, uri)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &res, "InsuranceFreightOpen")
	return
}

// InsuranceFreightGetOrderList 拉取保单信息
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/express/business/insurance_freight.html
func (express *Express) InsuranceFreightGetOrderList(ctx context.Context, req *InsuranceFreightGetOrderListRequest) (res InsuranceFreightGetOrderListResponse, err error) {
	accessToken, err := express.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(insuranceFreightGetOrderListURL, accessToken)
	response, err := util.PostJSONContext(ctx, uri, req)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &res, "InsuranceFreightGetOrderList")
	return
}

// InsuranceFreightGetPayOrderList 拉取充值订单信息
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/express/business/insurance_freight.html
func (express *Express) InsuranceFreightGetPayOrderList(ctx context.Context, req *InsuranceFreightGetPayOrderListRequest) (res InsuranceFreightGetPayOrderListResponse, err error) {
	accessToken, err := express.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(insuranceFreightGetPayOrderListURL, accessToken)
	response, err := util.PostJSONContext(ctx, uri, req)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &res, "InsuranceFreightGetPayOrderList")
	return
}

// InsuranceFreightGetSummary 拉取理赔摘要
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/express/business/insurance_freight.html
func (express *Express) InsuranceFreightGetSummary(ctx context.Context, req *InsuranceFreightGetSummaryRequest) (res InsuranceFreightGetSummaryResponse, err error) {
	accessToken, err := express.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(insuranceFreightGetSummaryURL, accessToken)
	response, err := util.PostJSONContext(ctx, uri, req)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &res, "InsuranceFreightGetSummary")
	return
}

// InsuranceFreightClaim 无忧退理赔
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/express/business/insurance_freight.html
func (express *Express) InsuranceFreightClaim(ctx context.Context, req *InsuranceFreightClaimRequest) (res InsuranceFreightClaimResponse, err error) {
	accessToken, err := express.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(insuranceFreightClaimURL, accessToken)
	response, err := util.PostJSONContext(ctx, uri, req)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &res, "InsuranceFreightClaim")
	return
}

// InsuranceFreightQueryOpen 查询开通状态
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/express/business/insurance_freight.html
func (express *Express) InsuranceFreightQueryOpen(ctx context.Context) (res InsuranceFreightQueryOpenResponse, err error) {
	accessToken, err := express.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(insuranceFreightQueryOpenURL, accessToken)
	response, err := util.HTTPGetContext(ctx, uri)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &res, "InsuranceFreightQueryOpen")
	return
}

// InsuranceFreightCreateChargeID 申请充值订单号
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/express/business/insurance_freight.html
func (express *Express) InsuranceFreightCreateChargeID(ctx context.Context, req *InsuranceFreightCreateChargeIDRequest) (res InsuranceFreightCreateChargeIDResponse, err error) {
	accessToken, err := express.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(insuranceFreightCreateChargeIDURL, accessToken)
	response, err := util.PostJSONContext(ctx, uri, req)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &res, "InsuranceFreightCreateChargeID")
	return
}

// InsuranceFreightApplyPay 申请支付
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/express/business/insurance_freight.html
func (express *Express) InsuranceFreightApplyPay(ctx context.Context, req *InsuranceFreightApplyPayRequest) (res InsuranceFreightApplyPayResponse, err error) {
	accessToken, err := express.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(insuranceFreightApplyPayURL, accessToken)
	response, err := util.PostJSONContext(ctx, uri, req)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &res, "InsuranceFreightApplyPay")
	return
}

// InsuranceFreightUpdateNotifyFunds 设置保费告警余额
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/express/business/insurance_freight.html
func (express *Express) InsuranceFreightUpdateNotifyFunds(ctx context.Context, req *InsuranceFreightUpdateNotifyFundsRequest) (res InsuranceFreightUpdateNotifyFundsResponse, err error) {
	accessToken, err := express.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(insuranceFreightUpdateNotifyFundsURL, accessToken)
	response, err := util.PostJSONContext(ctx, uri, req)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &res, "InsuranceFreightUpdateNotifyFunds")
	return
}

// InsuranceFreightRefundResponse 充值保费退款返回参数
type InsuranceFreightRefundResponse struct {
	util.CommonError
}

// InsuranceFreightCreateOrderRequest 发货时投保请求参数
type InsuranceFreightCreateOrderRequest struct {
	OpenID        string                       `json:"openid"`         // 买家openid，必须和理赔openid一致
	OrderNo       string                       `json:"order_no"`       // 微信支付单号，一个微信支付单号只能投保一次
	PayTime       int64                        `json:"pay_time"`       // 微信支付时间，秒级时间戳，时间误差3天内
	PayAmount     int64                        `json:"pay_amount"`     // 微信支付金额（单位：分）
	DeliveryNo    string                       `json:"delivery_no"`    // 发货运单号
	DeliveryPlace *InsuranceFreightPlace       `json:"delivery_place"` // 发货地址
	ReceiptPlace  *InsuranceFreightPlace       `json:"receipt_place"`  // 收货地址
	ProductInfo   *InsuranceFreightProductInfo `json:"product_info"`   // 投保订单信息，用于微信下发投保和理赔通知给用户，用户点击可查看投保订单，点击订单可跳回商家小程序
}

// InsuranceFreightPlace 运费险地址信息
type InsuranceFreightPlace struct {
	Province string `json:"province"` // 省
	City     string `json:"city"`     // 市
	County   string `json:"county"`   // 区
	Address  string `json:"address"`  // 详细地址
}

// InsuranceFreightProductInfo 运费险投保订单信息
type InsuranceFreightProductInfo struct {
	OrderPath string                       `json:"order_path"` // 投保订单在商家小程序的path
	GoodsList []*InsuranceFreightGoodsInfo `json:"goods_list"` // 投保商品list，一个元素为对象的数组
}

// InsuranceFreightGoodsInfo 运费险投保商品信息
type InsuranceFreightGoodsInfo struct {
	Name string `json:"name"` // 投保商品名称
	URL  string `json:"url"`  // 投保商品图片url
}

// InsuranceFreightCreateOrderResponse 发货时投保返回参数
type InsuranceFreightCreateOrderResponse struct {
	util.CommonError
	PolicyNo         string `json:"policy_no"`          // 保单号
	InsuranceEndDate string `json:"insurance_end_date"` // 保险止期，格式: yyyy-mm-dd hh24:mi:ss
	EstimateAmount   int64  `json:"estimate_amount"`    // 保险公司预估理赔金额，单位： 分
	Premium          int64  `json:"premium"`            // 保费，单位： 分
}

// InsuranceFreightOpenResponse 开通无忧退货返回参数
type InsuranceFreightOpenResponse struct {
	util.CommonError
}

// InsuranceFreightGetOrderListRequest 拉取保单信息请求参数
type InsuranceFreightGetOrderListRequest struct {
	OpenID           string `json:"openid,omitempty"`             // 买家openid，与投保理赔保持一致
	OrderNo          string `json:"order_no,omitempty"`           // 微信支付单号，与投保理赔保持一致
	PolicyNo         string `json:"policy_no,omitempty"`          // 保单号
	ReportNo         string `json:"report_no,omitempty"`          // 理赔报案号
	DeliveryNo       string `json:"delivery_no,omitempty"`        // 发货运单号
	RefundDeliveryNo string `json:"refund_delivery_no,omitempty"` // 退款运单号
	BeginTime        int64  `json:"begin_time,omitempty"`         // 查询开始时间，秒级时间戳
	EndTime          int64  `json:"end_time,omitempty"`           // 查询结束时间，秒级时间戳
	StatusList       []int  `json:"status_list,omitempty"`        // 保单状态
	Offset           int    `json:"offset,omitempty"`             // 分页offset
	Limit            int    `json:"limit,omitempty"`              // 分页limit（默认为100，最大为100）
	SortDirect       int    `json:"sort_direct,omitempty"`        // 排序方式
}

// InsuranceFreightGetOrderListResponse 拉取保单信息返回参数
type InsuranceFreightGetOrderListResponse struct {
	util.CommonError
	Total int                      `json:"total"` // 总数
	List  []*InsuranceFreightOrder `json:"list"`  // 保单列表
}

// InsuranceFreightOrder 保单信息
type InsuranceFreightOrder struct {
	OrderNo          string `json:"order_no"`           // 微信支付单号
	PolicyNo         string `json:"policy_no"`          // 保单号
	ReportNo         string `json:"report_no"`          // 理赔报案号
	DeliveryNo       string `json:"delivery_no"`        // 发货运单号
	RefundDeliveryNo string `json:"refund_delivery_no"` // 退款运单号
	Premium          int64  `json:"premium"`            // 保费（单位：分 ）
	EstimateAmount   int64  `json:"estimate_amount"`    // 预估理赔金额（单位：分 ）
	Status           int    `json:"status"`             // 保单状态
	PayFailReason    string `json:"pay_fail_reason"`    // 理赔打款失败原因
	PayFinishTime    int64  `json:"pay_finish_time"`    // 理赔款打给用户的时间
	IsHomePickUp     int    `json:"is_home_pick_up"`    // 是否上门取件
	InsuranceEndDate string `json:"insurance_end_date"` // 保险止期，格式: yyyy-mm-dd hh24:mi:ss
}

// InsuranceFreightGetPayOrderListRequest 拉取充值订单信息请求参数
type InsuranceFreightGetPayOrderListRequest struct {
	StatusList []int `json:"status_list"`      // 订单状态
	Offset     int   `json:"offset,omitempty"` // 分页offset
	Limit      int   `json:"limit,omitempty"`  // 分页limit
}

// InsuranceFreightGetPayOrderListResponse 拉取充值订单信息返回参数
type InsuranceFreightGetPayOrderListResponse struct {
	util.CommonError
	Total int                         `json:"total"` // 总数
	List  []*InsuranceFreightPayOrder `json:"list"`  // 充值订单列表
}

// InsuranceFreightPayOrder 充值订单信息
type InsuranceFreightPayOrder struct {
	OrderID      int64 `json:"order_id"`      // 充值单号
	OrderStatus  int   `json:"order_status"`  // 订单状态
	TotalPrice   int64 `json:"total_price"`   // 充值金额
	CreateTime   int64 `json:"create_time"`   // 订单创建时间
	PayTime      int64 `json:"pay_time"`      // 支付时间
	CanRefund    bool  `json:"can_refund"`    // 是否可以退款（官方参数表为 number、返回示例为布尔值）
	RefundTime   int64 `json:"refund_time"`   // 退款时间
	RefundStatus int   `json:"refund_status"` // 退款状态
	RefundAmt    int64 `json:"refund_amt"`    // 退款金额
}

// InsuranceFreightGetSummaryRequest 拉取理赔摘要请求参数
type InsuranceFreightGetSummaryRequest struct {
	BeginTime int64 `json:"begin_time"` // 查询开始时间，
	EndTime   int64 `json:"end_time"`   // 查询结束时间戳，秒级时间戳
}

// InsuranceFreightGetSummaryResponse 拉取理赔摘要返回参数
type InsuranceFreightGetSummaryResponse struct {
	util.CommonError
	Total        int   `json:"total"`          // 投保总数
	ClaimNum     int   `json:"claim_num"`      // 理赔总数
	ClaimSuccNum int   `json:"claim_succ_num"` // 理赔成功数
	Premium      int64 `json:"premium"`        // 当前保费，单位: 分
	Funds        int64 `json:"funds"`          // 当前账号余额，单位: 分
	NeedClose    bool  `json:"need_close"`     // 是否不能投保，系统安全原因不能投保
}

// InsuranceFreightClaimRequest 无忧退理赔请求参数
type InsuranceFreightClaimRequest struct {
	OpenID           string `json:"openid"`             // 买家openid，与投保保持一致
	OrderNo          string `json:"order_no"`           // 微信支付单号，与投保保持一致
	RefundDeliveryNo string `json:"refund_delivery_no"` // 退款运单号，理赔退款运单号唯一
	RefundCompany    string `json:"refund_company"`     // 退款快递公司
}

// InsuranceFreightClaimResponse 无忧退理赔返回参数
type InsuranceFreightClaimResponse struct {
	util.CommonError
	ReportNo     string `json:"report_no"`       // 理赔报案号，成功申请理赔时返回
	IsHomePickUp int    `json:"is_home_pick_up"` // 是否上门取件
}

// InsuranceFreightQueryOpenResponse 查询开通状态返回参数
type InsuranceFreightQueryOpenResponse struct {
	util.CommonError
	IsOpen int `json:"is_open"` // 是否开通
}

// InsuranceFreightCreateChargeIDRequest 申请充值订单号请求参数
type InsuranceFreightCreateChargeIDRequest struct {
	Quota int64 `json:"quota"` // 充值金额，单位：分
}

// InsuranceFreightCreateChargeIDResponse 申请充值订单号返回参数
type InsuranceFreightCreateChargeIDResponse struct {
	util.CommonError
	OrderID int64 `json:"order_id"` // 充值订单id
}

// InsuranceFreightApplyPayRequest 申请支付请求参数
type InsuranceFreightApplyPayRequest struct {
	OrderID int64 `json:"order_id"` // 充值订单id，与 申请充值订单号 保持一致，js等语言防止精度错误传string
}

// InsuranceFreightApplyPayResponse 申请支付返回参数
type InsuranceFreightApplyPayResponse struct {
	util.CommonError
	PayURL string `json:"pay_url"` // 充值链接
}

// InsuranceFreightUpdateNotifyFundsRequest 设置保费告警余额请求参数
type InsuranceFreightUpdateNotifyFundsRequest struct {
	NotifyFunds int64 `json:"notify_funds"` // 通知的金额（单位：分；设置为0不通知）
}

// InsuranceFreightUpdateNotifyFundsResponse 设置保费告警余额返回参数
type InsuranceFreightUpdateNotifyFundsResponse struct {
	util.CommonError
}
