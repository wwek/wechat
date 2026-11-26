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
	// bindKfWorkURL 绑定微信客服
	bindKfWorkURL = "https://api.weixin.qq.com/customservice/work/bind"
	// unbindKfWorkURL 解除绑定微信客服
	unbindKfWorkURL = "https://api.weixin.qq.com/customservice/work/unbind"
	// getKfWorkBoundURL 查询绑定情况
	getKfWorkBoundURL = "https://api.weixin.qq.com/customservice/work/get"
)

// BindKfWorkRequest 绑定微信客服请求
type BindKfWorkRequest struct {
	CorpID string `json:"corpid"` // 企业ID，获取方式参考：术语说明-corpid
}

// BindKfWorkResponse 绑定微信客服响应
type BindKfWorkResponse struct {
	util.CommonError
}

// BindKfWork 绑定微信客服
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/kf-work/api_bindkfwork
func (kf *KF) BindKfWork(req *BindKfWorkRequest) (res BindKfWorkResponse, err error) {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return
	}

	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf("%s?access_token=%s", bindKfWorkURL, accessToken), req); err != nil {
		return
	}

	err = util.DecodeWithError(response, &res, "BindKfWork")
	return
}

// UnbindKfWorkRequest 解除绑定微信客服请求
type UnbindKfWorkRequest struct {
	CorpID string `json:"corpid"` // 企业ID，获取方式参考：术语说明-corpid
}

// UnbindKfWorkResponse 解除绑定微信客服响应
type UnbindKfWorkResponse struct {
	util.CommonError
}

// UnbindKfWork 解除绑定微信客服
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/kf-work/api_unbindkfwork
func (kf *KF) UnbindKfWork(req *UnbindKfWorkRequest) (res UnbindKfWorkResponse, err error) {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return
	}

	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf("%s?access_token=%s", unbindKfWorkURL, accessToken), req); err != nil {
		return
	}

	err = util.DecodeWithError(response, &res, "UnbindKfWork")
	return
}

// GetKfWorkBoundResponse 查询绑定情况响应
type GetKfWorkBoundResponse struct {
	util.CommonError
	EntityName string `json:"entityName"` // 该小程序的主体名称，未绑定时不返回
	CorpID     string `json:"corpid"`     // 企业ID，未绑定时不返回
	BindTime   int64  `json:"bindTime"`   // 绑定时间戳（单位：秒），未绑定时不返回
}

// GetKfWorkBound 查询绑定情况
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/kf-work/api_getkfworkbound
func (kf *KF) GetKfWorkBound() (res GetKfWorkBoundResponse, err error) {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return
	}

	params := map[string]interface{}{
		"access_token": accessToken,
	}
	query := util.Query(params)

	var response []byte
	if response, err = util.HTTPGet(fmt.Sprintf("%s?%s", getKfWorkBoundURL, query)); err != nil {
		return
	}

	err = util.DecodeWithError(response, &res, "GetKfWorkBound")
	return
}
