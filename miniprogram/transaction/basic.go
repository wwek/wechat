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
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// getGuaranteeStatusURL 获取交易保障标状态
	getGuaranteeStatusURL = "https://api.weixin.qq.com/wxaapi/wxamptrade/get_guarantee_status?%s"
	// getPenaltyListURL 获取小程序交易体验分违规记录
	getPenaltyListURL = "https://api.weixin.qq.com/wxaapi/wxamptrade/get_penalty_list?%s"
)

// GetGuaranteeStatusResponse 获取交易保障标状态响应
type GetGuaranteeStatusResponse struct {
	util.CommonError
	IsActived util.FlexBool `json:"isActived"` // 是否激活交易保障标（官方返回示例中为字符串）
	Msg       string        `json:"msg"`       // 保障标状态相关信息
	Reasons   []string      `json:"reasons"`   // 未能开通保障标的原因
}

// GetGuaranteeStatus 获取交易保障标状态
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/transaction-guarantee/basic/api_getguaranteestatus
func (s *Transaction) GetGuaranteeStatus() (out GetGuaranteeStatusResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	params := map[string]interface{}{
		"access_token": accessToken,
	}

	response, err := util.HTTPGet(fmt.Sprintf(getGuaranteeStatusURL, util.Query(params)))
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetGuaranteeStatus")
	return
}

// GetPenaltyListRequest 获取小程序交易体验分违规记录请求
type GetPenaltyListRequest struct {
	Offset int64 `json:"offset"` // 获取从第offset条开始的limit条记录（序号从 0 开始），最大不超过总记录数
	Limit  int64 `json:"limit"`  // 获取从第offset条开始的limit条记录（序号从 0 开始），最大不超过 100
}

// GetPenaltyListResponse 获取小程序交易体验分违规记录响应
type GetPenaltyListResponse struct {
	util.CommonError
	AppealList   []PenaltyItem `json:"appealList"`   // 记录列表
	CurrentScore int64         `json:"currentScore"` // 当前小程序的交易体验分
	TotalNum     int64         `json:"totalNum"`     // 当前小程序的总扣分记录数
}

// GetPenaltyList 获取小程序交易体验分违规记录
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/transaction-guarantee/basic/api_getpenaltylist
func (s *Transaction) GetPenaltyList(in *GetPenaltyListRequest) (out GetPenaltyListResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	params := map[string]interface{}{
		"access_token": accessToken,
		"offset":       in.Offset,
		"limit":        in.Limit,
	}

	response, err := util.HTTPGet(fmt.Sprintf(getPenaltyListURL, util.Query(params)))
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetPenaltyList")
	return
}

// PenaltyItem 扣分记录
type PenaltyItem struct {
	IllegalOrderID   util.FlexString `json:"illegalOrderId"`   // 扣分记录ID（官方返回示例中为字符串）
	ComplaintOrderID util.FlexString `json:"complaintOrderId"` // 投诉单ID（官方返回示例中为字符串）
	IllegalWording   string          `json:"illegalWording"`   // 违规行为
	Status           int64           `json:"status"`           // 扣分记录状态
	MinusScore       int64           `json:"minusScore"`       // 扣除分数
	OrderID          util.FlexString `json:"orderId"`          // 订单号（官方返回示例中为字符串，如 payorder@xxx）
	IllegalTime      int64           `json:"illegalTime"`      // 扣分记录创建时间
	UpdateTime       int64           `json:"updateTime"`       // 更新时间
}
