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
	// scatterGetBillURL 对账单下载
	scatterGetBillURL = "https://api.weixin.qq.com/cgi-bin/express/delivery/single_waybill/get_bill?access_token=%s"
	// updateBusinessURL 更新商户审核结果
	updateBusinessURL = "https://api.weixin.qq.com/cgi-bin/express/delivery/service/business/update?access_token=%s"
	// updatePathURL 更新运单轨迹
	updatePathURL = "https://api.weixin.qq.com/cgi-bin/express/delivery/path/update?access_token=%s"
	// getContactURL 获取面单联系人信息
	getContactURL = "https://api.weixin.qq.com/cgi-bin/express/delivery/contact/get?access_token=%s"
	// scatterDeliveryRefundURL 运力公司退款
	scatterDeliveryRefundURL = "https://api.weixin.qq.com/cgi-bin/express/delivery/single_waybill/refund_order?access_token=%s"
	// scatterDeliveryCancelURL 运力取消订单
	scatterDeliveryCancelURL = "https://api.weixin.qq.com/cgi-bin/express/delivery/single_waybill/cancel_order?access_token=%s"
	// scatterUpdateOrderStatusURL 运力更新订单状态
	scatterUpdateOrderStatusURL = "https://api.weixin.qq.com/cgi-bin/express/delivery/single_waybill/update?access_token=%s"
	// scatterUpdateOrderFeeURL 运力更新需支付的运费
	scatterUpdateOrderFeeURL = "https://api.weixin.qq.com/cgi-bin/express/delivery/single_waybill/fee?access_token=%s"
	// scatterUpdateComplainResultURL 运力返回用户投诉处理结果
	scatterUpdateComplainResultURL = "https://api.weixin.qq.com/cgi-bin/express/delivery/scatter/update_complaint_result?access_token=%s"
	// previewTemplateURL 预览面单模板
	previewTemplateURL = "https://api.weixin.qq.com/cgi-bin/express/delivery/template/preview?access_token=%s"
)

// ScatterGetBillRequest 对账单下载请求参数
type ScatterGetBillRequest struct {
	Date string `json:"date"` // 下载对账单的时间，格式为YYYYMMDD
	Type string `json:"type"` // 下载对账单的类型 ALL-所以账单 SUCCESS-成功支付账单 REFUND-退款账单
}

// ScatterGetBillResponse 对账单下载返回参数
type ScatterGetBillResponse struct {
	util.CommonError
}

// ScatterGetBill 对账单下载
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/express/express-by-provider/api_scatter_get_bill
func (s *Logistics) ScatterGetBill(in *ScatterGetBillRequest) (out ScatterGetBillResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(scatterGetBillURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "ScatterGetBill")
	return
}

// UpdateBusinessRequest 更新商户审核结果请求参数
type UpdateBusinessRequest struct {
	ShopAppID  string `json:"shop_app_id"` // 商户的小程序AppID，即审核商户事件中的 ShopAppID
	BizID      string `json:"biz_id"`      // 商户账户
	ResultCode int64  `json:"result_code"` // 审核结果，0 表示审核通过，其他表示审核失败
	ResultMsg  string `json:"result_msg"`  // 审核错误原因，仅 result_code 不等于 0 时需要设置
}

// UpdateBusinessResponse 更新商户审核结果返回参数
type UpdateBusinessResponse struct {
	util.CommonError
}

// UpdateBusiness 更新商户审核结果
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/express/express-by-provider/api_updatebusiness
func (s *Logistics) UpdateBusiness(in *UpdateBusinessRequest) (out UpdateBusinessResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(updateBusinessURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "UpdateBusiness")
	return
}

// UpdatePathRequest 更新运单轨迹请求参数
type UpdatePathRequest struct {
	Token      string `json:"token"`       // 商户侧下单事件中推送的 Token 字段
	WaybillID  string `json:"waybill_id"`  // 运单 ID
	ActionTime int64  `json:"action_time"` // 轨迹变化 Unix 时间戳
	ActionType int64  `json:"action_type"` // 轨迹变化类型
	ActionMsg  string `json:"action_msg"`  // 轨迹变化具体信息说明，展示在快递轨迹详情页中，使用UTF-8编码
}

// UpdatePathResponse 更新运单轨迹返回参数
type UpdatePathResponse struct {
	util.CommonError
}

// UpdatePath 更新运单轨迹
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/express/express-by-provider/api_updatepath
func (s *Logistics) UpdatePath(in *UpdatePathRequest) (out UpdatePathResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(updatePathURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "UpdatePath")
	return
}

// Contact 面单联系人信息
type Contact struct {
	Name    string `json:"name"`    // 用户姓名
	Tel     string `json:"tel"`     // 座机号码
	Mobile  string `json:"mobile"`  // 手机号码
	Address string `json:"address"` // 地址，已经将省市区信息合并
}

// GetContactRequest 获取面单联系人信息请求参数
type GetContactRequest struct {
	Token     string `json:"token"`      // 商户侧下单事件中推送的 Token 字段
	WaybillID string `json:"waybill_id"` // 运单 ID
}

// GetContactResponse 获取面单联系人信息返回参数
type GetContactResponse struct {
	util.CommonError
	WaybillID string  `json:"waybill_id"` // 运单 ID
	Sender    Contact `json:"sender"`     // 发件人信息
	Receiver  Contact `json:"receiver"`   // 收件人信息
}

// GetContact 获取面单联系人信息
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/express/express-by-provider/api_getcontact
func (s *Logistics) GetContact(in *GetContactRequest) (out GetContactResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getContactURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetContact")
	return
}

// ScatterDeliveryRefundRequest 运力公司退款请求参数
type ScatterDeliveryRefundRequest struct {
	Token string `json:"token"` // 下单时的Token字段
	Fee   int64  `json:"fee"`   // 退款金额，单位分
}

// ScatterDeliveryRefundResponse 运力公司退款返回参数
type ScatterDeliveryRefundResponse struct {
	util.CommonError
}

// ScatterDeliveryRefund 运力公司退款
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/express/express-by-provider/api_scatterdeliveryrefund
func (s *Logistics) ScatterDeliveryRefund(in *ScatterDeliveryRefundRequest) (out ScatterDeliveryRefundResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(scatterDeliveryRefundURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "ScatterDeliveryRefund")
	return
}

// ScatterDeliveryCancelRequest 运力取消订单请求参数
type ScatterDeliveryCancelRequest struct {
	Token  string `json:"token"`  // 商户侧下单事件中推送的 Token 字段
	Reason string `json:"reason"` // 取消的原因
}

// ScatterDeliveryCancelResponse 运力取消订单返回参数
type ScatterDeliveryCancelResponse struct {
	util.CommonError
}

// ScatterDeliveryCancel 运力取消订单
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/express/express-by-provider/api_scatterdeliverycancel
func (s *Logistics) ScatterDeliveryCancel(in *ScatterDeliveryCancelRequest) (out ScatterDeliveryCancelResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(scatterDeliveryCancelURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "ScatterDeliveryCancel")
	return
}

// ScatterUpdateOrderStatusRequest 运力更新订单状态请求参数
type ScatterUpdateOrderStatusRequest struct {
	Token                string `json:"token"`                  // 商户侧下单事件中推送的Token字段
	WaybillID            string `json:"waybill_id"`             // 运单ID
	ActionTime           int64  `json:"action_time"`            // 轨迹变化Unix时间戳
	ActionType           int64  `json:"action_type"`            // 轨迹变化类型，与普通单保持一致
	ActionMsg            string `json:"action_msg"`             // 轨迹变化具体信息说明，展示在快递轨迹详情页中，使用UTF-8编码
	PickupCourierName    string `json:"pickup_courier_name"`    // 取件员姓名
	PickupCourierPhone   string `json:"pickup_courier_phone"`   // 取件员电话
	DeliveryCourierName  string `json:"delivery_courier_name"`  // 派件员姓名
	DeliveryCourierPhone string `json:"delivery_courier_phone"` // 派件员电话
}

// ScatterUpdateOrderStatusResponse 运力更新订单状态返回参数
type ScatterUpdateOrderStatusResponse struct {
	util.CommonError
}

// ScatterUpdateOrderStatus 运力更新订单状态
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/express/express-by-provider/api_scatterupdateorderstatus
func (s *Logistics) ScatterUpdateOrderStatus(in *ScatterUpdateOrderStatusRequest) (out ScatterUpdateOrderStatusResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(scatterUpdateOrderStatusURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "ScatterUpdateOrderStatus")
	return
}

// ScatterUpdateOrderFeeRequest 运力更新需支付的运费请求参数
type ScatterUpdateOrderFeeRequest struct {
	Token        string `json:"token"`          // 商户侧下单事件中推送的 Token 字段
	WaybillID    string `json:"waybill_id"`     // 运单 ID
	NeedPay      int64  `json:"need_pay"`       // 是否需要用户在线支付，0不需要，1需要，2需要(支付分)
	Fee          int64  `json:"fee"`            // 需要支付的金额（一般等于original_fee，如有优惠则填最终需要支付的金额），单位分
	OriginalFee  int64  `json:"original_fee"`   // 原价（base_fee+insured_fee+other_fee），单位分
	BaseFee      int64  `json:"base_fee"`       // 运费，单位分
	InsuredFee   int64  `json:"insured_fee"`    // 保价费，单位分
	OtherFee     int64  `json:"other_fee"`      // 其他费用，单位分
	Remark       int64  `json:"remark"`         // 其他费用备注
	PayGoodsName string `json:"pay_goods_name"` // 商品名称（对应微信支付商品详情页上的商品名称）
}

// ScatterUpdateOrderFeeResponse 运力更新需支付的运费返回参数
type ScatterUpdateOrderFeeResponse struct {
	util.CommonError
}

// ScatterUpdateOrderFee 运力更新需支付的运费
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/express/express-by-provider/api_scatterupdateorderfee
func (s *Logistics) ScatterUpdateOrderFee(in *ScatterUpdateOrderFeeRequest) (out ScatterUpdateOrderFeeResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(scatterUpdateOrderFeeURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "ScatterUpdateOrderFee")
	return
}

// ScatterUpdateComplainResultRequest 运力返回用户投诉处理结果请求参数
type ScatterUpdateComplainResultRequest struct {
	Token     string `json:"token"`      // 商户侧下单事件中推送的 Token 字段
	WaybillID string `json:"waybill_id"` // 运单 ID
	Result    string `json:"result"`     // 处理结果
	Desc      string `json:"desc"`       // 处理结果说明
}

// ScatterUpdateComplainResultResponse 运力返回用户投诉处理结果返回参数
type ScatterUpdateComplainResultResponse struct {
	util.CommonError
}

// ScatterUpdateComplainResult 运力返回用户投诉处理结果
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/express/express-by-provider/api_scatterupdatecomplainresult
func (s *Logistics) ScatterUpdateComplainResult(in *ScatterUpdateComplainResultRequest) (out ScatterUpdateComplainResultResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(scatterUpdateComplainResultURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "ScatterUpdateComplainResult")
	return
}

// PreviewTemplateRequest 预览面单模板请求参数
type PreviewTemplateRequest struct {
	WaybillID       string          `json:"waybill_id"`       // 运单 ID
	WaybillTemplate string          `json:"waybill_template"` // 面单 HTML 模板内容（需经 Base64 编码）
	WaybillData     string          `json:"waybill_data"`     // 面单数据，详情参考下单事件返回值中的 WaybillData
	Custom          AddOrderRequest `json:"custom"`           // 商户下单数据，格式是商户侧下单 AddOrder 接口中的请求体
}

// PreviewTemplateResponse 预览面单模板返回参数
type PreviewTemplateResponse struct {
	util.CommonError
	WaybillID               string `json:"waybill_id"`                // 运单 ID
	RenderedWaybillTemplate string `json:"rendered_waybill_template"` // 渲染后的面单 HTML 文件（已经过 Base64 编码）
}

// PreviewTemplate 预览面单模板
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/express/express-by-provider/api_previewtemplate
func (s *Logistics) PreviewTemplate(in *PreviewTemplateRequest) (out PreviewTemplateResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(previewTemplateURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "PreviewTemplate")
	return
}
