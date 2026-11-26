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

package cityservice

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// sendChannelMsgURL 消息推送接口
	sendChannelMsgURL = "https://api.weixin.qq.com/cityservice/sendchannelmsg?access_token=%s"
)

// SendChannelMsgRequest 消息推送接口请求
type SendChannelMsgRequest struct {
	Status       int64        `json:"status"`        // 消息状态
	OpenID       string       `json:"open_id"`       // 用户openid
	OrderID      string       `json:"order_id"`      // 订单号
	MsgID        string       `json:"msg_id"`        // 消息id
	AppID        string       `json:"app_id"`        // 小程序appid
	BusinessID   int64        `json:"business_id"`   // 业务id
	BusinessInfo BusinessInfo `json:"business_info"` // 业务信息
}

// SendChannelMsgResponse 消息推送接口响应
type SendChannelMsgResponse struct {
	util.CommonError
}

// SendChannelMsg 消息推送接口
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/cityservice/medicalassistant/api_cityservice_sendchannelmsg
func (s *CityService) SendChannelMsg(in *SendChannelMsgRequest) (out SendChannelMsgResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(sendChannelMsgURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "SendChannelMsg")
	return
}

// BusinessInfo 业务信息
type BusinessInfo struct {
	PatHospitalID      string       `json:"pat_hospital_id"`     // 医院id
	PatName            string       `json:"pat_name"`            // 患者姓名
	DocName            string       `json:"doc_name"`            // 医生姓名
	DepartmentName     string       `json:"department_name"`     // 科室名称
	DepartmentLocation string       `json:"department_location"` // 科室位置
	AppointmentTime    string       `json:"appointment_time"`    // 预约时间，如2023-06-07 10:30-11:00
	Memo               string       `json:"memo"`                // 备注
	RedirectPage       RedirectPage `json:"redirect_page"`       // 跳转页面
	ElderRedirectPage  RedirectPage `json:"elder_redirect_page"` // 长者跳转页面
}

// RedirectPage 跳转页面
type RedirectPage struct {
	PageType string `json:"page_type"` // 页面路径类型，如果有redirect_page字段需要填，则必填page_type; 否则不用填
	URL      string `json:"url"`       // 网页网址，page_type为web类型时必填，url需为业务域名
	AppID    string `json:"app_id"`    // 小程序appid，page_type为小程序时必填，小程序类型时需与当前账号关联主体
	FullPath string `json:"fullpath"`  // 小程序路径，page_type为小程序时必填
}
