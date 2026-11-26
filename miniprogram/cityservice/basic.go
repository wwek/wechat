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
	// getBusinessViewURL 仿原生跳转
	getBusinessViewURL = "https://api.weixin.qq.com/intp/transportcode/getbusinessview?access_token=%s"
	// checkRealNameInfoURL 校验实名信息
	checkRealNameInfoURL = "https://api.weixin.qq.com/intp/realname/checkrealnameinfo?access_token=%s"
	// sendMsgDataURL 消息通路发消息
	sendMsgDataURL = "https://api.weixin.qq.com/cityservice/sendmsgdata?access_token=%s"
	// getServicePathURL 获取城市服务限定页面链接
	getServicePathURL = "https://api.weixin.qq.com/cityservice/getservicepath?access_token=%s"
)

// GetBusinessViewRequest 仿原生跳转请求
type GetBusinessViewRequest struct {
	PathType int64 `json:"path_type"` // 需要跳转的页面
}

// GetBusinessViewResponse 仿原生跳转响应
type GetBusinessViewResponse struct {
	util.CommonError
	BusinessType string `json:"business_type"` // 业务类型
	QueryString  string `json:"query_string"`  // 调用仿原生小程序时使用的参数
	ExpireAt     int64  `json:"expire_at"`     // 返回 query_string 的到期时间（uinx时间戳）
}

// GetBusinessView 仿原生跳转
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/cityservice/basic/api_transportcode_getbusinessview
func (s *CityService) GetBusinessView(in *GetBusinessViewRequest) (out GetBusinessViewResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getBusinessViewURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetBusinessView")
	return
}

// CheckRealNameInfoRequest 校验实名信息请求
type CheckRealNameInfoRequest struct {
	OpenID   string `json:"openid"`    // 用户在业务方下的openid，与申请权限时提供的业务方的小程序appid保持一致
	RealName string `json:"real_name"` // 需要校验的姓名
	CredID   string `json:"cred_id"`   // 需要校验的证件号
	CredType string `json:"cred_type"` // 默认为1，即身份证，目前暂只支持身份证
	Code     string `json:"code"`      // 通过小程序回跳获取的code参数
}

// CheckRealNameInfoResponse 校验实名信息响应
type CheckRealNameInfoResponse struct {
	util.CommonError
	VerifyOpenID   string `json:"verify_openid"`    // openid 验证结果，有多个结果时用分号;连接
	VerifyRealName string `json:"verify_real_name"` // real_name 校验结果，当verify_openid 为V_OP_NM_MA 时返回
}

// CheckRealNameInfo 校验实名信息
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/cityservice/basic/api_checkrealnameinfo
func (s *CityService) CheckRealNameInfo(in *CheckRealNameInfoRequest) (out CheckRealNameInfoResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(checkRealNameInfoURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "CheckRealNameInfo")
	return
}

// SendMsgDataRequest 消息通路发消息请求
type SendMsgDataRequest struct {
	OpenID            string      `json:"openid"`               // 用户唯一标识
	BizTemplateID     string      `json:"biz_template_id"`      // 城市服务分配给公众号的模板id
	ResultPageStyleID string      `json:"result_page_style_id"` // 结果页样式id，含结果页必填
	DealMsgStyleID    string      `json:"deal_msg_style_id"`    // 办事记录样式id，含办事记录必填
	CardStyleID       string      `json:"card_style_id"`        // 页卡样式id，含页卡必填
	OrderNo           string      `json:"order_no"`             // 订单号，同一订单号的办事记录会合并
	URL               string      `json:"url"`                  // 跳转链接，用于服务通知、结果页、待办提醒，。含结果页必填
	Data              interface{} `json:"data"`                 // 模板json数据，对象信息请按照模板要求，其中color字段只对服务通知有效。如为数组时可用[ ]括起字段内数据。
}

// SendMsgDataResponse 消息通路发消息响应
type SendMsgDataResponse struct {
	util.CommonError
	ResultPageURL string `json:"result_page_url"` // 需跳转至该url，替代原有的服务结果页面。如未传入result_page_style_id，则调用后result_page_url返回为空。
}

// SendMsgData 消息通路发消息
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/cityservice/basic/api_cityservice_sendmsgdata
func (s *CityService) SendMsgData(in *SendMsgDataRequest) (out SendMsgDataResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(sendMsgDataURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "SendMsgData")
	return
}

// GetServicePathRequest 获取城市服务限定页面链接请求
type GetServicePathRequest struct {
	PageType     int64       `json:"page_type"`      // 获取城市服务路径类型
	SrcChannel   int64       `json:"src_channel"`    // 跳转来源渠道
	NeedPathType int64       `json:"need_path_type"` // 获取h5 URL需填写1（OpenSDK不需要填写）
	DeviceType   int64       `json:"device_type"`    // 获取h5 URL需填写2（OpenSDK不需要填写）
	CityName     string      `json:"city_name"`      // 城市名称，page_type为1、3、5时必填（城市名称以地级市为准）
	ContentName  string      `json:"content_name"`   // 专题页名称，page_type为3时必填
	ExtParams    []ExtParams `json:"ext_params"`     // 附加参数，包括关键字等其他参数，page_type为5时必填
	ServiceID    int64       `json:"service_id"`     // 服务id，page_type为0时必填
	Params       string      `json:"params"`         // 透传参数，page_type为0时必填 必须为一个json数组，每个object包括两个成员key和value
	CityID       string      `json:"city_id"`        // 用户所在城市id(需与腾讯内部要是同一套编码)，page_type为0时可传
}

// GetServicePathResponse 获取城市服务限定页面链接响应
type GetServicePathResponse struct {
	util.CommonError
	Path         string `json:"path"`          // 结果路径
	BusinessType string `json:"business_type"` // 类型
	AppID        string `json:"app_id"`        // 小程序 id，page_type为5时返回
	Username     string `json:"username"`      // 原始id，page_type为5时返回
	QueryString  string `json:"query_string"`  // path的参数，page_type为0时返回
}

// GetServicePath 获取城市服务限定页面链接
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/cityservice/basic/api_cityserviceservicehomepath
func (s *CityService) GetServicePath(in *GetServicePathRequest) (out GetServicePathResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getServicePathURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetServicePath")
	return
}

// ExtParams 附加参数
type ExtParams struct {
	Key   string `json:"key"`   // 关键字的key名称置顶位keyword
	Value string `json:"value"` // 服务列表名称（以城市服务的标签名称为准）
}
