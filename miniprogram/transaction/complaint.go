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

package transaction

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// respondComplaintURL 商家回应投诉
	respondComplaintURL = "https://api.weixin.qq.com/wxaapi/minishop/bussiRespondComplaint?access_token=%s"
	// submitRefundURL 商家提交退款凭证
	submitRefundURL = "https://api.weixin.qq.com/wxaapi/minishop/bussiSupplyRefund?access_token=%s"
	// busiAppealURL 商家申诉
	busiAppealURL = "https://api.weixin.qq.com/wxaapi/minishop/busiAppeal?access_token=%s"
	// supplyProofURL 商家补充凭证
	supplyProofURL = "https://api.weixin.qq.com/wxaapi/minishop/bussiSupplyProof?access_token=%s"
	// getOrderDetailURL 查询投诉单详情
	getOrderDetailURL = "https://api.weixin.qq.com/wxaapi/minishop/complaintOrderDetail?%s"
)

// RespondComplaintRequest 商家回应投诉请求
type RespondComplaintRequest struct {
	Content          string   `json:"content"`          // 回应的内容（图片数组mediaIdList跟内容content二选一）
	ComplaintOrderID int64    `json:"complaintOrderId"` // 单号
	MediaIDList      []string `json:"mediaIdList"`      // 图片 id 列表，可参考 新增素材 （图片数组mediaIdList跟内容content二选一）
	BussiHandle      int64    `json:"bussiHandle"`      // 操作：1是同意和解，2是拒绝和解
}

// RespondComplaintResponse 商家回应投诉响应
type RespondComplaintResponse struct {
	util.CommonError
}

// RespondComplaint 商家回应投诉
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/transaction-guarantee/complaint/api_respondcomplaint
func (s *Transaction) RespondComplaint(in *RespondComplaintRequest) (out RespondComplaintResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(respondComplaintURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "RespondComplaint")
	return
}

// SubmitRefundRequest 商家提交退款凭证请求
type SubmitRefundRequest struct {
	Content          string   `json:"content"`          // 回应的内容（图片数组mediaIdList跟内容content二选一）
	ComplaintOrderID int64    `json:"complaintOrderId"` // 单号
	MediaIDList      []string `json:"mediaIdList"`      // 图片的id list 图片 id 列表，可参考 新增素材 （图片数组mediaIdList跟内容content二选一）
	AcceptReturn     int64    `json:"acceptReturn"`     // 确认是否收货（退货状态下必填）
	ReturnID         int64    `json:"returnId"`         // 退货单号，需要通过查询投诉单详情接口获取（退货状态下必填）
}

// SubmitRefundResponse 商家提交退款凭证响应
type SubmitRefundResponse struct {
	util.CommonError
}

// SubmitRefund 商家提交退款凭证
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/transaction-guarantee/complaint/api_submitrefund
func (s *Transaction) SubmitRefund(in *SubmitRefundRequest) (out SubmitRefundResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(submitRefundURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "SubmitRefund")
	return
}

// BusiAppealRequest 商家申诉请求
type BusiAppealRequest struct {
	Content          string   `json:"content"`          // 回应的内容（图片数组mediaIdList跟内容content二选一）
	ComplaintOrderID int64    `json:"complaintOrderId"` // 单号
	MediaIDList      []string `json:"mediaIdList"`      // 图片 id 列表，可参考这 新增素材 （图片数组mediaIdList跟内容content二选一）
}

// BusiAppealResponse 商家申诉响应
type BusiAppealResponse struct {
	util.CommonError
}

// BusiAppeal 商家申诉
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/transaction-guarantee/complaint/api_busiappeal
func (s *Transaction) BusiAppeal(in *BusiAppealRequest) (out BusiAppealResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(busiAppealURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "BusiAppeal")
	return
}

// SupplyProofRequest 商家补充凭证请求
type SupplyProofRequest struct {
	Content          string   `json:"content"`          // 回应的内容（图片数组mediaIdList跟内容content二选一）
	ComplaintOrderID int64    `json:"complaintOrderId"` // 单号
	MediaIDList      []string `json:"mediaIdList"`      // 图片 id 列表，可参考这 新增素材 （图片数组mediaIdList跟内容content二选一）
}

// SupplyProofResponse 商家补充凭证响应
type SupplyProofResponse struct {
	util.CommonError
}

// SupplyProof 商家补充凭证
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/transaction-guarantee/complaint/api_supplyproof
func (s *Transaction) SupplyProof(in *SupplyProofRequest) (out SupplyProofResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(supplyProofURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "SupplyProof")
	return
}

// GetOrderDetailRequest 查询投诉单详情请求
type GetOrderDetailRequest struct {
	ComplaintOrderID string `json:"complaintOrderId"` // 订单id
}

// GetOrderDetailResponse 查询投诉单详情响应
type GetOrderDetailResponse struct {
	util.CommonError
	ComplaintOrder ComplaintOrder `json:"complaintOrder"` // 投诉单信息
	Item           ComplaintItem  `json:"item"`           // 投诉进度
	ReturnBill     ReturnBill     `json:"returnBill"`     // 退货单信息
}

// UnmarshalJSON 官方文档中 item 字段的参数表标注为对象、返回示例却是数组，
// 这里两种形态都接受，避免与线上返回不符时整个响应解码失败。
func (r *GetOrderDetailResponse) UnmarshalJSON(data []byte) error {
	type alias GetOrderDetailResponse
	aux := &struct {
		Item json.RawMessage `json:"item"`
		*alias
	}{alias: (*alias)(r)}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	if len(aux.Item) == 0 || string(aux.Item) == "null" {
		return nil
	}
	if aux.Item[0] == '[' {
		var list []ComplaintItem
		if err := json.Unmarshal(aux.Item, &list); err != nil {
			return err
		}
		if len(list) > 0 {
			r.Item = list[0]
		}
		return nil
	}
	return json.Unmarshal(aux.Item, &r.Item)
}

// GetOrderDetail 查询投诉单详情
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/transaction-guarantee/complaint/api_getorderdetail
func (s *Transaction) GetOrderDetail(in *GetOrderDetailRequest) (out GetOrderDetailResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	params := map[string]interface{}{
		"access_token":     accessToken,
		"complaintOrderId": url.QueryEscape(in.ComplaintOrderID),
	}

	response, err := util.HTTPGet(fmt.Sprintf(getOrderDetailURL, util.Query(params)))
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetOrderDetail")
	return
}

// ComplaintOrder 投诉单信息
type ComplaintOrder struct {
	ComplaintOrderID string            `json:"complaintOrderId"` // 投诉单号
	OpenID           string            `json:"openid"`           // 投诉用户openid
	CreateTime       int64             `json:"createTime"`       // 投诉单创建时间
	PhoneNumber      int64             `json:"phoneNumber"`      // 投诉用户手机号（官方返回为数字类型）
	Type             int64             `json:"type"`             // 投诉单类型
	Status           int64             `json:"status"`           // 投诉单状态
	CustomerMaterial ComplaintMaterial `json:"customerMaterial"` // 投诉凭证
	OrderID          string            `json:"orderId"`          // 订单id
	OutTradeNo       string            `json:"outTradeNo"`       // 外部交易单号
	ProductName      string            `json:"productName"`      // 商品名称
	PayTime          int64             `json:"payTime"`          // 支付时间
	TotalCost        int64             `json:"totalCost"`        // 订单总金额
	ExpireTime       int64             `json:"expireTime"`       // 投诉过期时间
	HeadImgURL       string            `json:"headImgUrl"`       // 投诉用户头像
	NickName         string            `json:"nickName"`         // 投诉用户昵称
	AppealState      int64             `json:"appealState"`      // 申诉状态
}

// ComplaintMaterial 投诉凭证
type ComplaintMaterial struct {
	Content     string   `json:"content"`     // 投诉内容
	MediaIDList []string `json:"mediaIdList"` // 投诉内容图片cdn(出于安全性考虑，目前该图片url有过期时间，如需查看图片，需每次重新调用该接口获取，防止图片过期导致无法查看)
}

// ComplaintItem 投诉进度
type ComplaintItem struct {
	ItemType       int64    `json:"itemType"`       // 投诉节点状态
	Time           int64    `json:"time"`           // 时间
	Content        string   `json:"content"`        // 内容文本
	MediaIDList    []string `json:"mediaIdList"`    // 图片cdn资源列表
	PhoneNumber    int64    `json:"phoneNumber"`    // 手机号（官方返回为数字类型）
	BlameResult    int64    `json:"blameResult"`    // itemType 为（31， 32）时，判断是否经过退货的描述， 1对应文案为待用户退货中，0则对应文案为待上传处理凭证
	NickName       string   `json:"nickName"`       // 操作者昵称
	AppealItemType int64    `json:"appealItemType"` // 处于申诉阶段节点的申诉状态
}

// ReturnBill 退货单信息
type ReturnBill struct {
	ReturnID     string `json:"returnId"`     // 退货id
	WaybillID    string `json:"waybillId"`    // 运单号
	DeliveryName string `json:"deliveryName"` // 运力公司
	OrderStatus  int64  `json:"orderStatus"`  // 运单状态
}
