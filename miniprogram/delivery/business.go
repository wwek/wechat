/*
 *   Copyright silenceper/wechat Author(https://silenceper.com/wechat/). All Rights Reserved.
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 *
 *    You can obtain one at https://github.com/silenceper/wechat.
 *
 */

package delivery

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// bindLocalAccountURL 发起绑定请求
	bindLocalAccountURL = "https://api.weixin.qq.com/cgi-bin/express/local/business/shop/add?access_token=%s"
	// cancelLocalOrderURL 取消配送单
	cancelLocalOrderURL = "https://api.weixin.qq.com/cgi-bin/express/local/business/order/cancel?access_token=%s"
	// abnormalConfirmURL 异常件退回商家确认
	abnormalConfirmURL = "https://api.weixin.qq.com/cgi-bin/express/local/business/order/confirm_return?access_token=%s"
	// getBindAccountURL 拉取已绑定账号
	getBindAccountURL = "https://api.weixin.qq.com/cgi-bin/express/local/business/shop/get?access_token=%s"
	// getLocalOrderURL 拉取配送单信息
	getLocalOrderURL = "https://api.weixin.qq.com/cgi-bin/express/local/business/order/get?access_token=%s"
	// realMockUpdateOrderURL 模拟更新配送单状态
	realMockUpdateOrderURL = "https://api.weixin.qq.com/cgi-bin/express/local/business/realmock_update_order?access_token=%s"
	// mockUpdateOrderURL 模拟配送公司更新配送单状态
	mockUpdateOrderURL = "https://api.weixin.qq.com/cgi-bin/express/local/business/test_update_order?access_token=%s"
	// addTipsURL 添加小费
	addTipsURL = "https://api.weixin.qq.com/cgi-bin/express/local/business/order/addtips?access_token=%s"
	// addLocalOrderURL 添加配送单
	addLocalOrderURL = "https://api.weixin.qq.com/cgi-bin/express/local/business/order/add?access_token=%s"
	// openDeliveryURL 申请开通即时配送
	openDeliveryURL = "https://api.weixin.qq.com/cgi-bin/express/local/business/open?access_token=%s"
	// getAllImmeDeliveryURL 获取已支持的配送公司列表
	getAllImmeDeliveryURL = "https://api.weixin.qq.com/cgi-bin/express/local/business/delivery/getall?access_token=%s"
	// reOrderURL 重新下单
	reOrderURL = "https://api.weixin.qq.com/cgi-bin/express/local/business/order/readd?access_token=%s"
	// preAddOrderURL 预下配送单
	preAddOrderURL = "https://api.weixin.qq.com/cgi-bin/express/local/business/order/pre_add?access_token=%s"
	// preCancelOrderURL 预取消配送单
	preCancelOrderURL = "https://api.weixin.qq.com/cgi-bin/express/local/business/order/precancel?access_token=%s"
)

// LocalContact 发/收件人信息
type LocalContact struct {
	Name           string  `json:"name"`            // 姓名，最长不超过256个字符
	City           string  `json:"city"`            // 城市名称，如广州市
	Address        string  `json:"address"`         // 地址(街道、小区、大厦等，用于定位)
	AddressDetail  string  `json:"address_detail"`  // 地址详情(楼号、单元号、层号)
	Phone          string  `json:"phone"`           // 电话/手机号，最长不超过64个字符
	Lng            float64 `json:"lng"`             // 经度（火星坐标或百度坐标，和 coordinate_type 字段配合使用，精确到小数点后6位）
	Lat            float64 `json:"lat"`             // 纬度（火星坐标或百度坐标，和 coordinate_type 字段配合使用，精确到小数点后6位）
	CoordinateType int64   `json:"coordinate_type"` // 坐标类型，0：火星坐标（高德，腾讯地图均采用火星坐标）1：百度坐标
}

// LocalGood 货品信息
type LocalGood struct {
	GoodCount int64   `json:"good_count"` // 货物数量
	GoodName  string  `json:"good_name"`  // 货品名称
	GoodPrice float64 `json:"good_price"` // 货品单价，精确到小数点后两位
	GoodUnit  string  `json:"good_unit"`  // 货品单位，最长不超过20个字符
}

// LocalGoodsDetail 货物详情
type LocalGoodsDetail struct {
	Goods []LocalGood `json:"goods"` // 货物列表
}

// LocalCargo 货物信息
type LocalCargo struct {
	GoodsValue       float64           `json:"goods_value"`        // 货物价格，单位为元，精确到小数点后两位，范围为(0-5000]
	GoodsHeight      float64           `json:"goods_height"`       // 货物高度，单位为cm，精确到小数点后两位，范围为(0-45]
	GoodsWidth       float64           `json:"goods_width"`        // 货物宽度，单位为cm，精确到小数点后两位，范围为(0-50]
	GoodsLength      float64           `json:"goods_length"`       // 货物长度，单位为cm，精确到小数点后两位，范围为(0-65]
	GoodsWeight      float64           `json:"goods_weight"`       // 货物重量，单位为kg，精确到小数点后两位，范围为(0-50]
	GoodsDetail      *LocalGoodsDetail `json:"goods_detail"`       // 货物详情，最长不超过10240个字符
	GoodsPickupInfo  string            `json:"goods_pickup_info"`  // 货物取货信息，用于骑手到店取货，最长不超过100个字符
	CargoFirstClass  string            `json:"cargo_first_class"`  // 品类一级类目
	CargoSecondClass string            `json:"cargo_second_class"` // 品类二级类目
}

// LocalOrderInfo 订单信息
type LocalOrderInfo struct {
	DeliveryServiceCode  string  `json:"delivery_service_code"`  // 配送服务代码，不同配送公司自定义，顺丰和达达不填
	ExpectedDeliveryTime int64   `json:"expected_delivery_time"` // 期望派单时间(达达支持)，unix-timestamp，比如1586342180
	OrderType            int64   `json:"order_type"`             // 订单类型，0: 即时单 1 预约单
	PoiSeq               string  `json:"poi_seq"`                // 门店订单流水号，建议提供，最长不超过32个字符
	Note                 string  `json:"note"`                   // 备注，最长不超过200个字符
	OrderTime            int64   `json:"order_time"`             // 用户下单付款时间，顺丰必填，比如1555220757
	IsInsured            int64   `json:"is_insured"`             // 是否保价，0，非保价，1.保价
	DeclaredValue        float64 `json:"declared_value"`         // 保价金额，单位为元，精确到分
	Tips                 float64 `json:"tips"`                   // 小费，单位为元，下单一般不加小费
	IsDirectDelivery     int64   `json:"is_direct_delivery"`     // 是否选择直拿直送（0：不需要；1：需要）
	CashOnDelivery       float64 `json:"cash_on_delivery"`       // 骑手应付金额，单位为元，精确到分
	CashOnPickup         float64 `json:"cash_on_pickup"`         // 骑手应收金额，单位为元，精确到分
	RiderPickMethod      int64   `json:"rider_pick_method"`      // 物流流向，1：从门店取件送至用户；2：从用户取件送至门店
	IsFinishCodeNeeded   int64   `json:"is_finish_code_needed"`  // 收货码（0：不需要；1：需要）
	IsPickupCodeNeeded   int64   `json:"is_pickup_code_needed"`  // 取货码（0：不需要；1：需要）
	ExpectedFinishTime   int64   `json:"expected_finish_time"`   // 期望送达时间(美团、顺丰同城急送支持），unix-timestamp，比如1586342180
	ExpectedPickTime     int64   `json:"expected_pick_time"`     // 期望取件时间（闪送、顺丰同城急送支持），unix-timestamp，比如1586342180
}

// LocalShopInfo 商品信息，会展示到物流通知消息中
type LocalShopInfo struct {
	WxaPath    string `json:"wxa_path"`    // 商家小程序的路径，建议为订单页面
	ImgURL     string `json:"img_url"`     // 商品缩略图 url
	GoodsName  string `json:"goods_name"`  // 商品名称, 不超过128字节
	GoodsCount int64  `json:"goods_count"` // 商品数量
	WxaAppID   string `json:"wxa_appid"`   // 第三方代授权商户下单时填写的授权小程序appid
}

// LocalShopAccount 已绑定的商家签约账号
type LocalShopAccount struct {
	ShopID      string `json:"shopid"`       // 商家id
	DeliveryID  string `json:"delivery_id"`  // 配送公司Id
	AuditResult int64  `json:"audit_result"` // 审核状态，0表示审核通过；1表示审核中；2表示审核不通过
}

// LocalDeliveryCompany 配送公司
type LocalDeliveryCompany struct {
	DeliveryID   string `json:"delivery_id"`   // 配送公司Id
	DeliveryName string `json:"delivery_name"` // 配送公司名称
}

// BindLocalAccountRequest 发起绑定请求参数
type BindLocalAccountRequest struct {
	DeliveryID string `json:"delivery_id"` // 配送公司ID
}

// BindLocalAccountResponse 发起绑定请求返回参数
type BindLocalAccountResponse struct {
	util.CommonError
	ResultCode int64  `json:"resultcode"` // 运力返回的错误码
	ResultMsg  string `json:"resultmsg"`  // 运力返回的错误描述
}

// BindLocalAccount 发起绑定请求
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/immediate-delivery/deliver-by-business/api_bindlocalaccount
func (s *Delivery) BindLocalAccount(in *BindLocalAccountRequest) (out BindLocalAccountResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(bindLocalAccountURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "BindLocalAccount")
	return
}

// CancelLocalOrderRequest 取消配送单请求参数
type CancelLocalOrderRequest struct {
	ShopID         string `json:"shopid"`           // 商家id，由配送公司分配的appkey
	ShopOrderID    string `json:"shop_order_id"`    // 唯一标识订单的 ID，由商户生成
	DeliveryID     string `json:"delivery_id"`      // 快递公司ID
	WaybillID      string `json:"waybill_id"`       // 配送单id（顺丰同城必填）
	CancelReasonID int64  `json:"cancel_reason_id"` // 取消原因Id，1暂时不需要邮寄；2价格不合适；3订单信息有误重新下单；4骑手取货不及时；5骑手配送不及时；6其他原因
	CancelReason   string `json:"cancel_reason"`    // 取消原因
	ShopNo         string `json:"shop_no"`          // 商家门店编号，闪送shop_no必填，值为店铺id
	DeliverySign   string `json:"delivery_sign"`    // 用配送公司提供的appSecret加密的校验串
}

// CancelLocalOrderResponse 取消配送单返回参数
type CancelLocalOrderResponse struct {
	util.CommonError
	ResultCode int64  `json:"resultcode"` // 运力返回的错误码
	ResultMsg  string `json:"resultmsg"`  // 运力返回的错误描述
	DeductFee  int64  `json:"deduct_fee"` // 扣除的违约金(单位：元)，精确到分
	Desc       string `json:"desc"`       // 说明
}

// CancelLocalOrder 取消配送单
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/immediate-delivery/deliver-by-business/api_cancellocalorder
func (s *Delivery) CancelLocalOrder(in *CancelLocalOrderRequest) (out CancelLocalOrderResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(cancelLocalOrderURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "CancelLocalOrder")
	return
}

// AbnormalConfirmRequest 异常件退回商家确认请求参数
type AbnormalConfirmRequest struct {
	ShopID       string `json:"shopid"`        // 商家id，由配送公司分配的appkey
	ShopOrderID  string `json:"shop_order_id"` // 唯一标识订单的 ID，由商户生成
	WaybillID    string `json:"waybill_id"`    // 配送单id
	DeliverySign string `json:"delivery_sign"` // 用配送公司提供的appSecret加密的校验串
	ShopNo       string `json:"shop_no"`       // 商家门店编号，在配送公司登记，闪送必填，值为店铺id
	Remark       string `json:"remark"`        // 备注
}

// AbnormalConfirmResponse 异常件退回商家确认返回参数
type AbnormalConfirmResponse struct {
	util.CommonError
	ResultCode int64  `json:"resultcode"` // 运力返回的错误码
	ResultMsg  string `json:"resultmsg"`  // 运力返回的错误描述
}

// AbnormalConfirm 异常件退回商家确认
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/immediate-delivery/deliver-by-business/api_abnormalconfirm
func (s *Delivery) AbnormalConfirm(in *AbnormalConfirmRequest) (out AbnormalConfirmResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(abnormalConfirmURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "AbnormalConfirm")
	return
}

// GetBindAccountResponse 拉取已绑定账号返回参数
type GetBindAccountResponse struct {
	util.CommonError
	ResultCode int64              `json:"resultcode"` // 运力返回的错误码
	ResultMsg  string             `json:"resultmsg"`  // 运力返回的错误描述
	ShopList   []LocalShopAccount `json:"shop_list"`  // 绑定的商家签约账号列表
}

// GetBindAccount 拉取已绑定账号
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/immediate-delivery/deliver-by-business/api_getbindaccount
func (s *Delivery) GetBindAccount() (out GetBindAccountResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getBindAccountURL, accessToken)
	response, err := util.PostJSON(uri, &struct{}{})
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetBindAccount")
	return
}

// GetLocalOrderRequest 拉取配送单信息请求参数
type GetLocalOrderRequest struct {
	ShopID       string `json:"shopid"`        // 商家id，由配送公司分配的appkey
	ShopOrderID  string `json:"shop_order_id"` // 唯一标识订单的 ID，由商户生成
	ShopNo       string `json:"shop_no"`       // 商家门店编号，在配送公司登记，如果只有一个门店，可以不填
	DeliverySign string `json:"delivery_sign"` // 用配送公司提供的appSecret加密的校验串
}

// GetLocalOrderResponse 拉取配送单信息返回参数
type GetLocalOrderResponse struct {
	util.CommonError
	ResultCode  int64   `json:"resultcode"`   // 运力返回的错误码
	ResultMsg   string  `json:"resultmsg"`    // 运力返回的错误描述
	OrderStatus int64   `json:"order_status"` // 配送状态
	WaybillID   string  `json:"waybill_id"`   // 配送单号
	RiderName   string  `json:"rider_name"`   // 骑手姓名
	RiderPhone  string  `json:"rider_phone"`  // 骑手电话
	RiderLng    float64 `json:"rider_lng"`    // 骑手位置经度, 配送中时返回
	RiderLat    float64 `json:"rider_lat"`    // 骑手位置纬度, 配送中时返回
	ReachTime   int64   `json:"reach_time"`   // 预计还剩多久送达时间，配送中时返回，单位秒，比如5分钟后送达填300
}

// GetLocalOrder 拉取配送单信息
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/immediate-delivery/deliver-by-business/api_getlocalorder
func (s *Delivery) GetLocalOrder(in *GetLocalOrderRequest) (out GetLocalOrderResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getLocalOrderURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetLocalOrder")
	return
}

// RealMockUpdateOrderRequest 模拟更新配送单状态请求参数
type RealMockUpdateOrderRequest struct {
	ShopID       string `json:"shopid"`        // 商家id
	ShopOrderID  string `json:"shop_order_id"` // 唯一标识订单的 ID，由商户生成
	OrderStatus  int64  `json:"order_status"`  // 配送状态
	ActionTime   int64  `json:"action_time"`   // 状态变更时间点，Unix秒级时间戳
	ActionMsg    string `json:"action_msg"`    // 附加信息
	DeliverySign string `json:"delivery_sign"` // 用配送公司提供的appSecret加密的校验串说明
}

// RealMockUpdateOrderResponse 模拟更新配送单状态返回参数
type RealMockUpdateOrderResponse struct {
	util.CommonError
	ResultCode int64  `json:"resultcode"` // 运力返回的错误码
	ResultMsg  string `json:"resultmsg"`  // 运力返回的错误描述
}

// RealMockUpdateOrder 模拟更新配送单状态
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/immediate-delivery/deliver-by-business/api_realmockupdateorder
func (s *Delivery) RealMockUpdateOrder(in *RealMockUpdateOrderRequest) (out RealMockUpdateOrderResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(realMockUpdateOrderURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "RealMockUpdateOrder")
	return
}

// MockUpdateOrderRequest 模拟配送公司更新配送单状态请求参数
type MockUpdateOrderRequest struct {
	ShopID      string `json:"shopid"`        // 商家id, 必须是 "test_shop_id"
	ShopOrderID string `json:"shop_order_id"` // 唯一标识订单的 ID，由商户生成
	OrderStatus int64  `json:"order_status"`  // 配送状态
	ActionTime  int64  `json:"action_time"`   // 状态变更时间点，Unix秒级时间戳
	ActionMsg   string `json:"action_msg"`    // 附加信息
}

// MockUpdateOrderResponse 模拟配送公司更新配送单状态返回参数
type MockUpdateOrderResponse struct {
	util.CommonError
	ResultCode int64  `json:"resultcode"` // 运力返回的错误码
	ResultMsg  string `json:"resultmsg"`  // 运力返回的错误描述
}

// MockUpdateOrder 模拟配送公司更新配送单状态
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/immediate-delivery/deliver-by-business/api_mockupdateorder
func (s *Delivery) MockUpdateOrder(in *MockUpdateOrderRequest) (out MockUpdateOrderResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(mockUpdateOrderURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "MockUpdateOrder")
	return
}

// AddTipsRequest 添加小费请求参数
type AddTipsRequest struct {
	ShopID       string  `json:"shopid"`        // 商家id，由配送公司分配的appkey
	ShopOrderID  string  `json:"shop_order_id"` // 唯一标识订单的 ID，由商户生成
	WaybillID    string  `json:"waybill_id"`    // 配送单id
	Tips         float64 `json:"tips"`          // 小费金额(单位：元)，各家配送公司最大值不同
	Remark       string  `json:"remark"`        // 备注
	DeliverySign string  `json:"delivery_sign"` // 用配送公司提供的appSecret加密的校验串说明
	ShopNo       string  `json:"shop_no"`       // 商家门店编号，在配送公司登记，闪送shop_no必填，值为店铺id
}

// AddTipsResponse 添加小费返回参数
type AddTipsResponse struct {
	util.CommonError
	ResultCode int64  `json:"resultcode"` // 运力返回的错误码
	ResultMsg  string `json:"resultmsg"`  // 运力返回的错误描述
}

// AddTips 添加小费
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/immediate-delivery/deliver-by-business/api_addtips
func (s *Delivery) AddTips(in *AddTipsRequest) (out AddTipsResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(addTipsURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "AddTips")
	return
}

// AddLocalOrderRequest 添加配送单请求参数
type AddLocalOrderRequest struct {
	ShopID        string         `json:"shopid"`         // 商家id，由配送公司分配的appkey
	ShopOrderID   string         `json:"shop_order_id"`  // 唯一标识订单的 ID，由商户生成, 不超过128字节
	DeliveryID    string         `json:"delivery_id"`    // 配送公司ID
	OpenID        string         `json:"openid"`         // 下单用户的openid
	Sender        LocalContact   `json:"sender"`         // 发件人信息
	Receiver      LocalContact   `json:"receiver"`       // 收件人信息
	Cargo         LocalCargo     `json:"cargo"`          // 货物信息
	OrderInfo     LocalOrderInfo `json:"order_info"`     // 订单信息
	Shop          LocalShopInfo  `json:"shop"`           // 商品信息，会展示到物流通知消息中
	DeliveryToken string         `json:"delivery_token"` // 预下单接口返回的参数，配送公司可保证在一段时间内运费不变
	DeliverySign  string         `json:"delivery_sign"`  // 用配送公司提供的appSecret加密的校验串
	ShopNo        string         `json:"shop_no"`        // 商家门店编号，在配送公司登记，美团闪送必填, 值为店铺id
	SubBizID      string         `json:"sub_biz_id"`     // 子商户id，区分小程序内部多个子商户
}

// AddLocalOrderResponse 添加配送单返回参数
type AddLocalOrderResponse struct {
	util.CommonError
	ResultCode       int64   `json:"resultcode"`        // 运力返回的错误码
	ResultMsg        string  `json:"resultmsg"`         // 运力返回的错误描述
	Fee              float64 `json:"fee"`               // 实际运费(单位：元)，运费减去优惠券费用
	DeliverFee       float64 `json:"deliverfee"`        // 运费(单位：元)
	CouponFee        float64 `json:"couponfee"`         // 优惠券费用(单位：元)
	Tips             float64 `json:"tips"`              // 小费(单位：元)
	InsuranceFee     float64 `json:"insurancfee"`       // 保价费(单位：元)
	Distance         int64   `json:"distance"`          // 配送距离(整数单位：米)
	WaybillID        string  `json:"waybill_id"`        // 配送单号
	OrderStatus      int64   `json:"order_status"`      // 配送状态
	FinishCode       int64   `json:"finish_code"`       // 收货码
	PickupCode       int64   `json:"pickup_code"`       // 取货码
	DispatchDuration int64   `json:"dispatch_duration"` // 预计骑手接单时间，单位秒，比如5分钟就填300, 无法预计填0
}

// AddLocalOrder 添加配送单
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/immediate-delivery/deliver-by-business/api_addlocalorder
func (s *Delivery) AddLocalOrder(in *AddLocalOrderRequest) (out AddLocalOrderResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(addLocalOrderURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "AddLocalOrder")
	return
}

// OpenDeliveryResponse 申请开通即时配送返回参数
type OpenDeliveryResponse struct {
	util.CommonError
	ResultCode int64  `json:"resultcode"` // 运力返回的错误码
	ResultMsg  string `json:"resultmsg"`  // 运力返回的错误描述
}

// OpenDelivery 申请开通即时配送
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/immediate-delivery/deliver-by-business/api_opendelivery
func (s *Delivery) OpenDelivery() (out OpenDeliveryResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(openDeliveryURL, accessToken)
	response, err := util.PostJSON(uri, &struct{}{})
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "OpenDelivery")
	return
}

// GetAllImmeDeliveryResponse 获取已支持的配送公司列表返回参数
type GetAllImmeDeliveryResponse struct {
	util.CommonError
	ResultCode int64                  `json:"resultcode"` // 运力返回的错误码
	ResultMsg  string                 `json:"resultmsg"`  // 运力返回的错误描述
	List       []LocalDeliveryCompany `json:"list"`       // 配送公司列表
}

// GetAllImmeDelivery 获取已支持的配送公司列表
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/immediate-delivery/deliver-by-business/api_getallimmedelivery
func (s *Delivery) GetAllImmeDelivery() (out GetAllImmeDeliveryResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getAllImmeDeliveryURL, accessToken)
	response, err := util.PostJSON(uri, &struct{}{})
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetAllImmeDelivery")
	return
}

// ReOrderRequest 重新下单请求参数
type ReOrderRequest struct {
	ShopID        string         `json:"shopid"`         // 商家id，由配送公司分配的appkey
	ShopOrderID   string         `json:"shop_order_id"`  // 唯一标识订单的 ID，由商户生成, 不超过128字节
	DeliveryID    string         `json:"delivery_id"`    // 配送公司ID
	OpenID        string         `json:"openid"`         // 下单用户的openid
	Sender        LocalContact   `json:"sender"`         // 发件人信息
	Receiver      LocalContact   `json:"receiver"`       // 收件人信息
	Cargo         LocalCargo     `json:"cargo"`          // 货物信息
	OrderInfo     LocalOrderInfo `json:"order_info"`     // 订单信息
	Shop          LocalShopInfo  `json:"shop"`           // 商品信息，会展示到物流通知消息中
	DeliveryToken string         `json:"delivery_token"` // 预下单接口返回的参数，配送公司可保证在一段时间内运费不变
	DeliverySign  string         `json:"delivery_sign"`  // 用配送公司提供的appSecret加密的校验串说明
	ShopNo        string         `json:"shop_no"`        // 商家门店编号，在配送公司登记，美团闪送必填, 值为店铺id
	SubBizID      string         `json:"sub_biz_id"`     // 子商户id，区分小程序内部多个子商户
}

// ReOrderResponse 重新下单返回参数
type ReOrderResponse struct {
	util.CommonError
	ResultCode       int64   `json:"resultcode"`        // 运力返回的错误码
	ResultMsg        string  `json:"resultmsg"`         // 运力返回的错误描述
	Fee              float64 `json:"fee"`               // 实际运费(单位：元)，运费减去优惠券费用
	DeliverFee       float64 `json:"deliverfee"`        // 运费(单位：元)
	CouponFee        float64 `json:"couponfee"`         // 优惠券费用(单位：元)
	Tips             float64 `json:"tips"`              // 小费(单位：元)
	InsuranceFee     float64 `json:"insurancfee"`       // 保价费(单位：元)
	Distance         int64   `json:"distance"`          // 配送距离(整数单位：米)
	WaybillID        string  `json:"waybill_id"`        // 配送单号
	OrderStatus      int64   `json:"order_status"`      // 配送状态
	FinishCode       int64   `json:"finish_code"`       // 收货码
	PickupCode       int64   `json:"pickup_code"`       // 取货码
	DispatchDuration int64   `json:"dispatch_duration"` // 预计骑手接单时间，单位秒，比如5分钟就填300, 无法预计填0
}

// ReOrder 重新下单
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/immediate-delivery/deliver-by-business/api_reorder
func (s *Delivery) ReOrder(in *ReOrderRequest) (out ReOrderResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(reOrderURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "ReOrder")
	return
}

// PreAddOrderRequest 预下配送单请求参数
type PreAddOrderRequest struct {
	ShopID       string         `json:"shopid"`        // 商家id，由配送公司分配的appkey
	ShopOrderID  string         `json:"shop_order_id"` // 唯一标识订单的 ID，由商户生成, 不超过128字节
	DeliveryID   string         `json:"delivery_id"`   // 配送公司ID
	OpenID       string         `json:"openid"`        // 下单用户的openid
	Sender       LocalContact   `json:"sender"`        // 发件人信息
	Receiver     LocalContact   `json:"receiver"`      // 收件人信息
	Cargo        LocalCargo     `json:"cargo"`         // 货物信息
	OrderInfo    LocalOrderInfo `json:"order_info"`    // 订单信息
	Shop         LocalShopInfo  `json:"shop"`          // 商品信息，会展示到物流通知消息中
	DeliverySign string         `json:"delivery_sign"` // 用配送公司提供的appSecret加密的校验串说明
	ShopNo       string         `json:"shop_no"`       // 商家门店编号，在配送公司登记，美团、闪送必填
	SubBizID     string         `json:"sub_biz_id"`    // 子商户id，区分小程序内部多个子商户
}

// PreAddOrderResponse 预下配送单返回参数
type PreAddOrderResponse struct {
	util.CommonError
	ResultCode       int64   `json:"resultcode"`        // 运力返回的错误码
	ResultMsg        string  `json:"resultmsg"`         // 运力返回的错误描述
	Fee              float64 `json:"fee"`               // 实际运费(单位：元)，运费减去优惠券费用
	DeliverFee       float64 `json:"deliverfee"`        // 运费(单位：元)
	CouponFee        float64 `json:"couponFee"`         // 优惠券费用(单位：元)
	Tips             float64 `json:"tips"`              // 小费(单位：元)
	InsuranceFee     float64 `json:"insurancefee"`      // 保价费(单位：元)
	Distance         int64   `json:"distance"`          // 配送距离(单位：米)
	DeliveryToken    string  `json:"delivery_token"`    // 配送公司返回的token，用户下单时带上这个字段，保证在一段时间内运费不变
	DispatchDuration int64   `json:"dispatch_duration"` // 预计骑手接单时间，单位秒，比如5分钟就填300, 无法预计填0
}

// PreAddOrder 预下配送单
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/immediate-delivery/deliver-by-business/api_preaddorder
func (s *Delivery) PreAddOrder(in *PreAddOrderRequest) (out PreAddOrderResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(preAddOrderURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "PreAddOrder")
	return
}

// PreCancelOrderRequest 预取消配送单请求参数
type PreCancelOrderRequest struct {
	ShopID         string `json:"shopid"`           // 商家id，由配送公司分配的appkey
	ShopOrderID    string `json:"shop_order_id"`    // 唯一标识订单的 ID，由商户生成
	DeliveryID     string `json:"delivery_id"`      // 快递公司ID
	WaybillID      string `json:"waybill_id"`       // 配送单id
	CancelReasonID int64  `json:"cancel_reason_id"` // 取消原因Id
	CancelReason   string `json:"cancel_reason"`    // 取消原因
	ShopNo         string `json:"shop_no"`          // 商家门店编号，在配送公司登记，闪送shop_no必填，值为店铺id
	DeliverySign   string `json:"delivery_sign"`    // 用配送公司提供的appSecret加密的校验串说明
}

// PreCancelOrderResponse 预取消配送单返回参数
type PreCancelOrderResponse struct {
	util.CommonError
	ResultCode int64  `json:"resultcode"` // 运力返回的错误码
	ResultMsg  string `json:"resultmsg"`  // 运力返回的错误描述
	DeductFee  int64  `json:"deduct_fee"` // 预计扣除的违约金(单位：元)，精确到分
	Desc       string `json:"desc"`       // 说明
}

// PreCancelOrder 预取消配送单
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/immediate-delivery/deliver-by-business/api_precancelorder
func (s *Delivery) PreCancelOrder(in *PreCancelOrderRequest) (out PreCancelOrderResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(preCancelOrderURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "PreCancelOrder")
	return
}
