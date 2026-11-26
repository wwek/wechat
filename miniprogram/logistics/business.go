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

package logistics

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// cancelOrderURL 取消运单
	cancelOrderURL = "https://api.weixin.qq.com/cgi-bin/express/business/order/cancel?access_token=%s"
	// batchGetOrderURL 批量获取运单数据
	batchGetOrderURL = "https://api.weixin.qq.com/cgi-bin/express/business/order/batchget?access_token=%s"
	// getPathURL 查询运单轨迹
	getPathURL = "https://api.weixin.qq.com/cgi-bin/express/business/path/get?access_token=%s"
	// testUpdateOrderURL 模拟更新订单状态
	testUpdateOrderURL = "https://api.weixin.qq.com/cgi-bin/express/business/test_update_order?access_token=%s"
	// addOrderURL 生成运单
	addOrderURL = "https://api.weixin.qq.com/cgi-bin/express/business/order/add?access_token=%s"
	// bindAccountURL 绑定/解绑物流账号
	bindAccountURL = "https://api.weixin.qq.com/cgi-bin/express/business/account/bind?access_token=%s"
	// getAllAccountURL 获取所有绑定的物流账号
	getAllAccountURL = "https://api.weixin.qq.com/cgi-bin/express/business/account/getall?access_token=%s"
	// getPrinterURL 获取打印员
	getPrinterURL = "https://api.weixin.qq.com/cgi-bin/express/business/printer/getall?access_token=%s"
	// getAllDeliveryURL 获取支持的快递公司列表
	getAllDeliveryURL = "https://api.weixin.qq.com/cgi-bin/express/business/delivery/getall?access_token=%s"
	// getQuotaURL 获取电子面单余额
	getQuotaURL = "https://api.weixin.qq.com/cgi-bin/express/business/quota/get?access_token=%s"
	// getOrderURL 获取运单数据
	getOrderURL = "https://api.weixin.qq.com/cgi-bin/express/business/order/get?access_token=%s"
	// updatePrinterURL 配置面单打印员
	updatePrinterURL = "https://api.weixin.qq.com/cgi-bin/express/business/printer/update?access_token=%s"
)

// AddOrderContact 发/收件人信息
type AddOrderContact struct {
	Name     string `json:"name"`      // 发/收件人姓名，不超过64字节
	Tel      string `json:"tel"`       // 发/收件人座机号码，若不填写则必须填写 mobile，不超过32字节
	Mobile   string `json:"mobile"`    // 发/收件人手机号码，若不填写则必须填写 tel，不超过32字节
	Company  string `json:"company"`   // 发/收件人公司名称，不超过64字节
	PostCode string `json:"post_code"` // 发/收件人邮编，不超过10字节
	Country  string `json:"country"`   // 发/收件人国家，不超过64字节
	Province string `json:"province"`  // 发/收件人省份，比如："广东省"，不超过64字节
	City     string `json:"city"`      // 发/收件人市/地区，比如："广州市"，不超过64字节
	Area     string `json:"area"`      // 发/收件人区/县，比如："海珠区"，不超过64字节
	Address  string `json:"address"`   // 发/收件人详细地址，比如："XX路XX号XX大厦XX"，不超过512字节
}

// CargoDetail 包裹商品明细
type CargoDetail struct {
	Name  string `json:"name"`  // 商品名，不超过128字节
	Count int64  `json:"count"` // 商品数量
}

// Cargo 包裹信息，将传递给快递公司
type Cargo struct {
	Count      int64         `json:"count"`       // 包裹数量, 默认为1
	Weight     float64       `json:"weight"`      // 货物总重量，比如1.2，单位是千克(kg)
	SpaceX     float64       `json:"space_x"`     // 货物长度，比如20.0，单位是厘米(cm)
	SpaceY     float64       `json:"space_y"`     // 货物宽度，比如15.0，单位是厘米(cm)
	SpaceZ     float64       `json:"space_z"`     // 货物高度，比如10.0，单位是厘米(cm)
	DetailList []CargoDetail `json:"detail_list"` // 货物明细列表
}

// ShopDetail 商品详情
type ShopDetail struct {
	GoodsName   string `json:"goods_name"`    // 商品名称
	GoodsImgURL string `json:"goods_img_url"` // 商品图片url
	GoodsDesc   string `json:"goods_desc"`    // 商品详情
}

// Shop 商品信息，会展示到物流服务通知和电子面单中
type Shop struct {
	WxaPath    string       `json:"wxa_path"`    // 商家小程序的路径，建议为订单页面
	ImgURL     string       `json:"img_url"`     // 商品缩略图 url；shop.detail_list为空则必传，shop.detail_list非空可不传
	GoodsName  string       `json:"goods_name"`  // 商品名称, 不超过128字节；shop.detail_list为空则必传，shop.detail_list非空可不传
	GoodsCount int64        `json:"goods_count"` // 商品数量；shop.detail_list为空则必传，默认取shop.detail_list的size
	DetailList []ShopDetail `json:"detail_list"` // 商品详情列表，适配多商品场景，用以消息落地页展示
	WxaAppID   string       `json:"wxa_appid"`   // 商家小程序 appid，模板消息跳转用
}

// Insured 保价信息
type Insured struct {
	UseInsured   int64 `json:"use_insured"`   // 是否保价，0 表示不保价，1 表示保价
	InsuredValue int64 `json:"insured_value"` // 保价金额，单位是分，比如: 10000 表示 100 元
}

// ServiceType 服务类型
type ServiceType struct {
	ServiceType int64  `json:"service_type"` // 服务类型ID，详见已经支持的快递公司基本信息
	ServiceName string `json:"service_name"` // 服务名称，详见已经支持的快递公司基本信息
}

// WaybillData 运单信息
type WaybillData struct {
	Key   string `json:"key"`   // 运单信息 key
	Value string `json:"value"` // 运单信息 value
}

// AddOrderRequest 生成运单请求参数
type AddOrderRequest struct {
	OrderID      string          `json:"order_id"`      // 订单ID，须保证全局唯一，不超过512字节
	OpenID       string          `json:"openid"`        // 用户openid，当add_source=2时无需填写（不发送物流服务通知）
	DeliveryID   string          `json:"delivery_id"`   // 快递公司ID，参见 GetAllDelivery
	BizID        string          `json:"biz_id"`        // 快递客户编码或者现付编码
	CustomRemark string          `json:"custom_remark"` // 快递备注信息，比如"易碎物品"，不超过1024字节
	TagID        int64           `json:"tagid"`         // 订单标签id，用于平台型小程序区分平台上的入驻方，非平台型小程序无需填写该字段
	AddSource    int64           `json:"add_source"`    // 订单来源，0为小程序订单，2为App或H5订单，填2则不发送物流服务通知
	WxAppID      string          `json:"wx_appid"`      // App或H5的appid，add_source=2时必填，需和开通了物流助手的小程序绑定同一open帐号
	Sender       AddOrderContact `json:"sender"`        // 发件人信息
	Receiver     AddOrderContact `json:"receiver"`      // 收件人信息
	Cargo        Cargo           `json:"cargo"`         // 包裹信息，将传递给快递公司
	Shop         Shop            `json:"shop"`          // 商品信息，会展示到物流服务通知和电子面单中
	Insured      Insured         `json:"insured"`       // 保价信息
	Service      ServiceType     `json:"service"`       // 服务类型
	ExpectTime   int64           `json:"expect_time"`   // Unix 时间戳, 单位秒，顺丰必须传，预期的上门揽件时间
	TakeMode     int64           `json:"take_mode"`     // 分单策略，0：线下网点签约，1：总部签约结算，目前支持圆通
}

// AddOrderResponse 生成运单返回参数
type AddOrderResponse struct {
	util.CommonError
	OrderID            string        `json:"order_id"`            // 订单ID，下单成功时返回
	WaybillID          string        `json:"waybill_id"`          // 运单ID，下单成功时返回
	DeliveryResultCode int64         `json:"delivery_resultcode"` // 快递侧错误码，下单失败时返回
	DeliveryResultMsg  string        `json:"delivery_resultmsg"`  // 快递侧错误信息，下单失败时返回
	WaybillData        []WaybillData `json:"waybill_data"`        // 运单信息，下单成功时返回
}

// AddOrder 生成运单
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/express/express-by-business/api_addorder
func (s *Logistics) AddOrder(in *AddOrderRequest) (out AddOrderResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(addOrderURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "AddOrder")
	return
}

// CancelOrderRequest 取消运单请求参数
type CancelOrderRequest struct {
	OpenID     string `json:"openid"`      // 用户openid，当add_source=2时无需填写（不发送物流服务通知）
	DeliveryID string `json:"delivery_id"` // 快递公司ID，参见 GetAllDelivery
	WaybillID  string `json:"waybill_id"`  // 运单ID
	OrderID    string `json:"order_id"`    // 订单 ID，需保证全局唯一
}

// CancelOrderResponse 取消运单返回参数
type CancelOrderResponse struct {
	util.CommonError
	DeliveryResultCode int64  `json:"delivery_resultcode"` // 运力返回的错误码
	DeliveryResultMsg  string `json:"delivery_resultmsg"`  // 运力返回的错误信息
}

// CancelOrder 取消运单
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/express/express-by-business/api_cancelorder
func (s *Logistics) CancelOrder(in *CancelOrderRequest) (out CancelOrderResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(cancelOrderURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "CancelOrder")
	return
}

// BatchGetOrderItem 批量获取运单数据的订单项
type BatchGetOrderItem struct {
	OrderID    string `json:"order_id"`    // 订单 ID，需保证全局唯一
	DeliveryID string `json:"delivery_id"` // 快递公司ID，参见 GetAllDelivery, 必须和waybill_id对应
	WaybillID  string `json:"waybill_id"`  // 运单ID
}

// BatchGetOrderRequest 批量获取运单数据请求参数
type BatchGetOrderRequest struct {
	OrderList []BatchGetOrderItem `json:"order_list"` // 订单列表, 最多不能超过100个
}

// BatchGetOrderItemResponse 批量获取运单数据的运单项
type BatchGetOrderItemResponse struct {
	util.CommonError
	PrintHTML   string        `json:"print_html"`   // 运单 html 的 BASE64 结果
	WaybillData []WaybillData `json:"waybill_data"` // 运单信息
	OrderID     string        `json:"order_id"`     // 订单ID
	DeliveryID  string        `json:"delivery_id"`  // 快递公司ID
	WaybillID   string        `json:"waybill_id"`   // 运单号
	OrderStatus int64         `json:"order_status"` // 运单状态, 0正常，1取消
}

// BatchGetOrderResponse 批量获取运单数据返回参数
type BatchGetOrderResponse struct {
	util.CommonError
	OrderList []BatchGetOrderItemResponse `json:"order_list"` // 运单列表
}

// BatchGetOrder 批量获取运单数据
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/express/express-by-business/api_batchgetorder
func (s *Logistics) BatchGetOrder(in *BatchGetOrderRequest) (out BatchGetOrderResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(batchGetOrderURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "BatchGetOrder")
	return
}

// GetPathRequest 查询运单轨迹请求参数
type GetPathRequest struct {
	OpenID     string `json:"openid"`      // 用户openid，当add_source=2时无需填写（不发送物流服务通知）
	DeliveryID string `json:"delivery_id"` // 快递公司ID，参见 GetAllDelivery
	WaybillID  string `json:"waybill_id"`  // 运单ID
	OrderID    string `json:"order_id"`    // 订单号
}

// PathItem 运单轨迹节点
type PathItem struct {
	ActionTime int64  `json:"action_time"` // 轨迹节点 Unix 时间戳
	ActionType int64  `json:"action_type"` // 轨迹节点类型
	ActionMsg  string `json:"action_msg"`  // 轨迹节点详情
}

// GetPathResponse 查询运单轨迹返回参数
type GetPathResponse struct {
	util.CommonError
	OpenID       string     `json:"openid"`         // 用户openid
	DeliveryID   string     `json:"delivery_id"`    // 快递公司 ID
	WaybillID    string     `json:"waybill_id"`     // 运单 ID
	PathItemNum  int64      `json:"path_item_num"`  // 轨迹节点数量
	PathItemList []PathItem `json:"path_item_list"` // 轨迹节点列表
}

// GetPath 查询运单轨迹
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/express/express-by-business/api_getpath
func (s *Logistics) GetPath(in *GetPathRequest) (out GetPathResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getPathURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetPath")
	return
}

// TestUpdateOrderRequest 模拟更新订单状态请求参数
type TestUpdateOrderRequest struct {
	BizID      string `json:"biz_id"`      // 商户id,需填test_biz_id
	OrderID    string `json:"order_id"`    // 订单号
	DeliveryID string `json:"delivery_id"` // 快递公司id,需填TEST
	WaybillID  string `json:"waybill_id"`  // 运单号
	ActionTime int64  `json:"action_time"` // 轨迹变化 Unix 时间戳
	ActionType int64  `json:"action_type"` // 轨迹变化类型
	ActionMsg  string `json:"action_msg"`  // 轨迹变化具体信息说明,使用UTF-8编码
}

// TestUpdateOrderResponse 模拟更新订单状态返回参数
type TestUpdateOrderResponse struct {
	util.CommonError
}

// TestUpdateOrder 模拟更新订单状态
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/express/express-by-business/api_testupdateorder
func (s *Logistics) TestUpdateOrder(in *TestUpdateOrderRequest) (out TestUpdateOrderResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(testUpdateOrderURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "TestUpdateOrder")
	return
}

// BindAccountRequest 绑定/解绑物流账号请求参数
type BindAccountRequest struct {
	Type          string `json:"type"`           // bind表示绑定，unbind表示解除绑定
	BizID         string `json:"biz_id"`         // 快递公司客户编码
	DeliveryID    string `json:"delivery_id"`    // 快递公司ID
	Password      string `json:"password"`       // 快递公司客户密码
	RemarkContent string `json:"remark_content"` // 备注内容（提交EMS审核需要）
}

// BindAccountResponse 绑定/解绑物流账号返回参数
type BindAccountResponse struct {
	util.CommonError
}

// BindAccount 绑定/解绑物流账号
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/express/express-by-business/api_bindaccount
func (s *Logistics) BindAccount(in *BindAccountRequest) (out BindAccountResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(bindAccountURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "BindAccount")
	return
}

// LogisticsAccount 已绑定的物流账号
type LogisticsAccount struct {
	BizID           string        `json:"biz_id"`            // 快递公司客户编码
	DeliveryID      string        `json:"delivery_id"`       // 快递公司ID
	CreateTime      int64         `json:"create_time"`       // 账号绑定时间
	UpdateTime      int64         `json:"update_time"`       // 账号更新时间
	StatusCode      int64         `json:"status_code"`       // 绑定状态
	Alias           string        `json:"alias"`             // 账号别名
	RemarkWrongMsg  string        `json:"remark_wrong_msg"`  // 账号绑定失败的错误信息（EMS审核结果）
	RemarkContent   string        `json:"remark_content"`    // 账号绑定时的备注内容（提交EMS审核需要）
	QuotaNum        int64         `json:"quota_num"`         // 电子面单余额
	QuotaUpdateTime int64         `json:"quota_update_time"` // 电子面单余额更新时间
	ServiceType     []ServiceType `json:"service_type"`      // 该绑定帐号支持的服务类型
}

// GetAllAccountResponse 获取所有绑定的物流账号返回参数
type GetAllAccountResponse struct {
	util.CommonError
	Count int64              `json:"count"` // 账号数量
	List  []LogisticsAccount `json:"list"`  // 账号列表
}

// GetAllAccount 获取所有绑定的物流账号
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/express/express-by-business/api_getallaccount
func (s *Logistics) GetAllAccount() (out GetAllAccountResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getAllAccountURL, accessToken)
	response, err := util.HTTPGet(uri)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetAllAccount")
	return
}

// GetPrinterResponse 获取打印员返回参数
type GetPrinterResponse struct {
	util.CommonError
	Count     int64             `json:"count"`      // 打印员数量
	OpenID    []string          `json:"openid"`     // 打印员openid
	TagIDList []util.FlexString `json:"tagid_list"` // tagid列表（官方返回示例中为字符串）
}

// GetPrinter 获取打印员
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/express/express-by-business/api_getprinter
func (s *Logistics) GetPrinter() (out GetPrinterResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getPrinterURL, accessToken)
	response, err := util.HTTPGet(uri)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetPrinter")
	return
}

// DeliveryInfo 快递公司信息
type DeliveryInfo struct {
	DeliveryID   string        `json:"delivery_id"`   // 快递公司 ID
	DeliveryName string        `json:"delivery_name"` // 快递公司名称
	CanUseCash   int64         `json:"can_use_cash"`  // 是否支持散单, 1表示支持
	CanGetQuota  int64         `json:"can_get_quota"` // 是否支持查询面单余额, 1表示支持
	ServiceType  []ServiceType `json:"service_type"`  // 支持的服务类型
	CashBizID    string        `json:"cash_biz_id"`   // 散单对应的bizid，当can_use_cash=1时有效
}

// GetAllDeliveryResponse 获取支持的快递公司列表返回参数
type GetAllDeliveryResponse struct {
	util.CommonError
	Count int64          `json:"count"` // 快递公司数量
	Data  []DeliveryInfo `json:"data"`  // 快递公司信息列表
}

// GetAllDelivery 获取支持的快递公司列表
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/express/express-by-business/api_getalldelivery
func (s *Logistics) GetAllDelivery() (out GetAllDeliveryResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getAllDeliveryURL, accessToken)
	response, err := util.HTTPGet(uri)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetAllDelivery")
	return
}

// GetQuotaRequest 获取电子面单余额请求参数
type GetQuotaRequest struct {
	DeliveryID string `json:"delivery_id"` // 快递公司ID，参见 GetAllDelivery
	BizID      string `json:"biz_id"`      // 快递公司客户编码
}

// GetQuotaResponse 获取电子面单余额返回参数
type GetQuotaResponse struct {
	util.CommonError
	QuotaNum int64 `json:"quota_num"` // 电子面单余额
}

// GetQuota 获取电子面单余额
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/express/express-by-business/api_getquota
func (s *Logistics) GetQuota(in *GetQuotaRequest) (out GetQuotaResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getQuotaURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetQuota")
	return
}

// GetOrderRequest 获取运单数据请求参数
type GetOrderRequest struct {
	OrderID    string `json:"order_id"`    // 订单 ID，需保证全局唯一
	OpenID     string `json:"openid"`      // 用户openid，当add_source=2时无需填写（不发送物流服务通知）
	DeliveryID string `json:"delivery_id"` // 快递公司ID，参见 GetAllDelivery, 必须和waybill_id对应
	WaybillID  string `json:"waybill_id"`  // 运单ID
	PrintType  int64  `json:"print_type"`  // 获取打印面单类型，1：一联单，0：二联单，默认获取二联单
}

// GetOrderResponse 获取运单数据返回参数
type GetOrderResponse struct {
	util.CommonError
	PrintHTML   string        `json:"print_html"`   // 运单 html 的 BASE64 结果
	WaybillData []WaybillData `json:"waybill_data"` // 运单信息
	OrderID     string        `json:"order_id"`     // 订单ID
	DeliveryID  string        `json:"delivery_id"`  // 快递公司ID
	WaybillID   string        `json:"waybill_id"`   // 运单号
	OrderStatus int64         `json:"order_status"` // 运单状态, 0正常，1取消
}

// GetOrder 获取运单数据
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/express/express-by-business/api_getorder
func (s *Logistics) GetOrder(in *GetOrderRequest) (out GetOrderResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getOrderURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetOrder")
	return
}

// UpdatePrinterRequest 配置面单打印员请求参数
type UpdatePrinterRequest struct {
	OpenID     string `json:"openid"`      // 打印员 openid
	UpdateType string `json:"update_type"` // 更新类型。bind表示绑定；unbind表示解除绑定
	TagIDList  string `json:"tagid_list"`  // 用于平台型小程序设置入驻方的打印员面单打印权限，同一打印员最多支持10个tagid，使用半角逗号分隔
}

// UpdatePrinterResponse 配置面单打印员返回参数
type UpdatePrinterResponse struct {
	util.CommonError
}

// UpdatePrinter 配置面单打印员
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/express/express-by-business/api_updateprinter
func (s *Logistics) UpdatePrinter(in *UpdatePrinterRequest) (out UpdatePrinterResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(updatePrinterURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "UpdatePrinter")
	return
}
