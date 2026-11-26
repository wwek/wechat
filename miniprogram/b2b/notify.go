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

package b2b

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// retailNotifyBusinessURL 模板消息列表及下发
	retailNotifyBusinessURL = "https://api.weixin.qq.com/wxa/business/retailnotifybusiness?access_token=%s"
	// getRetailMessageListURL 消息效果数据
	getRetailMessageListURL = "https://api.weixin.qq.com/wxa/business/getretailmessagelist?access_token=%s"
)

// RetailNotifyBusinessRequest 模板消息列表及下发请求
type RetailNotifyBusinessRequest struct {
	Type       int64    `json:"type"`         // 消息类型，枚举值参考 B2b 门店助手模板消息汇总列表
	ToUserList []string `json:"to_user_list"` // 门店负责人 openid 列表，最多一次发送 200 个 openid
	Content    string   `json:"content"`      // 消息内容，为 json 格式的字符串
}

// RetailNotifyBusinessResponse 模板消息列表及下发响应
type RetailNotifyBusinessResponse struct {
	util.CommonError
}

// RetailNotifyBusiness 模板消息列表及下发
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/notify/api_retailnotifybusiness
func (s *B2B) RetailNotifyBusiness(in *RetailNotifyBusinessRequest) (out RetailNotifyBusinessResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(retailNotifyBusinessURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "RetailNotifyBusiness")
	return
}

// GetRetailMessageListRequest 消息效果数据请求
type GetRetailMessageListRequest struct {
	Start     int64  `json:"start"`      // 分页起始位置，从 0 开始计数，默认值为 0
	Offset    int64  `json:"offset"`     // 每页返回的数据条数上限，最大值为 1000，默认值为 20
	BeginDate string `json:"begin_date"` // 开始时间，格式为 "xxxx-yy-zz"，如 "2024-08-01"
	EndDate   string `json:"end_date"`   // 结束时间，格式为 "xxxx-yy-zz"，如 "2024-08-10"
}

// RetailMessageData 消息效果数据
type RetailMessageData struct {
	MsgID         int64  `json:"msg_id"`          // 消息 id
	MsgType       int64  `json:"msg_type"`        // 消息类型
	Date          string `json:"date"`            // 消息日期
	MsgTime       string `json:"msg_time"`        // 消息发送时间
	SendUV        int64  `json:"send_uv"`         // 发送人数
	EntryUV       int64  `json:"entry_uv"`        // 进入人数
	BusinessMsgID string `json:"business_msg_id"` // 自定义 msg_id
}

// GetRetailMessageListResponse 消息效果数据响应
type GetRetailMessageListResponse struct {
	util.CommonError
	TotalNum int64               `json:"total_num"` // 查询到数据总数
	DataLine []RetailMessageData `json:"data_line"` // 数据列表
}

// GetRetailMessageList 消息效果数据
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/notify/api_getretailmessagelist
func (s *B2B) GetRetailMessageList(in *GetRetailMessageListRequest) (out GetRetailMessageListResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getRetailMessageListURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetRetailMessageList")
	return
}
