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

package kf

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// cancelKfAdminURL 取消客服管理员
	cancelKfAdminURL = "https://api.weixin.qq.com/customservice/kfaccount/canceladmin"
	// setKfAdminURL 设置客服管理员
	setKfAdminURL = "https://api.weixin.qq.com/customservice/kfaccount/setadmin"
	// typingURL 客服输入状态
	typingURL = "https://api.weixin.qq.com/cgi-bin/message/custom/business/typing"
)

// CancelKfAdminRequest 取消客服管理员请求
type CancelKfAdminRequest struct {
	KfOpenID   string `json:"kf_openid"`   // 客服人员的openid
	BusinessID string `json:"business_id"` // 子商户id，普通账号无需填写
}

// CancelKfAdminResponse 取消客服管理员响应
type CancelKfAdminResponse struct {
	util.CommonError
}

// CancelKfAdmin 取消客服管理员
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/kf-mgnt/kf-message/api_cancelkfadmin
func (kf *KF) CancelKfAdmin(req *CancelKfAdminRequest) (res CancelKfAdminResponse, err error) {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return
	}

	params := map[string]interface{}{
		"access_token": accessToken,
		"kf_openid":    req.KfOpenID,
	}
	if req.BusinessID != "" {
		params["business_id"] = req.BusinessID
	}
	query := util.Query(params)

	var response []byte
	if response, err = util.HTTPGet(fmt.Sprintf("%s?%s", cancelKfAdminURL, query)); err != nil {
		return
	}

	err = util.DecodeWithError(response, &res, "CancelKfAdmin")
	return
}

// SetKfAdminRequest 设置客服管理员请求
type SetKfAdminRequest struct {
	KfOpenID   string `json:"kf_openid"`   // 客服人员的openid
	BusinessID string `json:"business_id"` // 子商户id，普通账号无需填写
}

// SetKfAdminResponse 设置客服管理员响应
type SetKfAdminResponse struct {
	util.CommonError
}

// SetKfAdmin 设置客服管理员
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/kf-mgnt/kf-message/api_setkfadmin
func (kf *KF) SetKfAdmin(req *SetKfAdminRequest) (res SetKfAdminResponse, err error) {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return
	}

	params := map[string]interface{}{
		"access_token": accessToken,
		"kf_openid":    req.KfOpenID,
	}
	if req.BusinessID != "" {
		params["business_id"] = req.BusinessID
	}
	query := util.Query(params)

	var response []byte
	if response, err = util.HTTPGet(fmt.Sprintf("%s?%s", setKfAdminURL, query)); err != nil {
		return
	}

	err = util.DecodeWithError(response, &res, "SetKfAdmin")
	return
}

// TypingRequest 客服输入状态请求
type TypingRequest struct {
	ToUser     string `json:"touser"`     // 用户的 OpenID
	Command    string `json:"command"`    // 命令。Typing表示对用户下发"正在输入"状态；CancelTyping表示取消对用户的"正在输入"状态
	BusinessID string `json:"businessid"` // 子商户 ID，普通账号无需填写
}

// TypingResponse 客服输入状态响应
type TypingResponse struct {
	util.CommonError
}

// Typing 客服输入状态
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/kf-mgnt/kf-message/api_typing
func (kf *KF) Typing(req *TypingRequest) (res TypingResponse, err error) {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return
	}

	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf("%s?access_token=%s", typingURL, accessToken), req); err != nil {
		return
	}

	err = util.DecodeWithError(response, &res, "Typing")
	return
}
