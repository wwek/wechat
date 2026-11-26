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

package soter

import (
	"fmt"

	"github.com/silenceper/wechat/v2/miniprogram/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	// verifySignatureURL 生物认证秘钥签名验证
	verifySignatureURL = "https://api.weixin.qq.com/cgi-bin/soter/verify_signature?access_token=%s"
)

// Soter 小程序 Soter
type Soter struct {
	*context.Context
}

// NewSoter 实例化
func NewSoter(ctx *context.Context) *Soter {
	return &Soter{ctx}
}

// VerifySignatureRequest 生物认证秘钥签名验证请求
type VerifySignatureRequest struct {
	OpenID        string `json:"openid"`         // 用户 openid
	JSONString    string `json:"json_string"`    // 通过 wx.startSoterAuthentication 成功回调获得的 resultJSON 字段
	JSONSignature string `json:"json_signature"` // 通过 wx.startSoterAuthentication 成功回调获得的 resultJSONSignature 字段
}

// VerifySignatureResponse 生物认证秘钥签名验证响应
type VerifySignatureResponse struct {
	util.CommonError
	IsOK util.FlexBool `json:"is_ok"` // 验证结果（官方返回示例中为字符串）
}

// VerifySignature 生物认证秘钥签名验证
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/soter/api_verifysignature
func (s *Soter) VerifySignature(in *VerifySignatureRequest) (out VerifySignatureResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(verifySignatureURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "VerifySignature")
	return
}
