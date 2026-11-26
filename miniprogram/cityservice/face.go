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
	// queryVerifyInfoURL 查询用户人脸核身真实验证结果
	queryVerifyInfoURL = "https://api.weixin.qq.com/cityservice/face/identify/queryverifyinfo?access_token=%s"
	// getVerifyIdURL 获取用户人脸核身会话唯一标识
	getVerifyIdURL = "https://api.weixin.qq.com/cityservice/face/identify/getverifyid?access_token=%s"
)

// QueryVerifyInfoRequest 查询用户人脸核身真实验证结果请求
type QueryVerifyInfoRequest struct {
	VerifyID string `json:"verify_id"`  // getVerifyId 接口返回的人脸核身会话唯一标识
	OutSeqNo string `json:"out_seq_no"` // 业务方系统外部流水号，必须和 getVerifyId 接口传入的一致
	CertHash string `json:"cert_hash"`  // 根据 getVerifyId 中传入的证件信息生成的信息摘要，计算方式见注意事项。
	OpenID   string `json:"openid"`     // 用户身份标识，必须和 getVerifyId 接口传入的一致
}

// QueryVerifyInfoResponse 查询用户人脸核身真实验证结果响应
type QueryVerifyInfoResponse struct {
	util.CommonError
	VerifyRet int64 `json:"verify_ret"` // 人脸核身验证结果
}

// QueryVerifyInfo 查询用户人脸核身真实验证结果
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/face/api_queryverifyinfo
func (s *CityService) QueryVerifyInfo(in *QueryVerifyInfoRequest) (out QueryVerifyInfoResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(queryVerifyInfoURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "QueryVerifyInfo")
	return
}

// GetVerifyIdRequest 获取用户人脸核身会话唯一标识请求
type GetVerifyIdRequest struct {
	OutSeqNo string   `json:"out_seq_no"` // 业务方系统内部流水号，要求5-32个字符内，只能包含数字、大小写字母和_-字符，且在同一个appid下唯一。
	CertInfo CertInfo `json:"cert_info"`  // 用户身份信息
	OpenID   string   `json:"openid"`     // 用户身份标识
}

// GetVerifyIdResponse 获取用户人脸核身会话唯一标识响应
type GetVerifyIdResponse struct {
	util.CommonError
	VerifyID  string `json:"verify_id"`  // 微信侧生成的人脸核身会话唯一标识，用于后续接口调用。长度不超过 256 字符
	ExpiresIn int64  `json:"expires_in"` // verify_id 有效期，过期后无法发起核身，默认值3600，单位：秒
}

// GetVerifyId 获取用户人脸核身会话唯一标识
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/face/api_getverifyid
func (s *CityService) GetVerifyId(in *GetVerifyIdRequest) (out GetVerifyIdResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getVerifyIdURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetVerifyId")
	return
}

// CertInfo 用户身份信息
type CertInfo struct {
	CertType string `json:"cert_type"` // 证件类型，身份证填 IDENTITY_CARD
	CertName string `json:"cert_name"` // 证件姓名
	CertNo   string `json:"cert_no"`   // 证件号码
}
