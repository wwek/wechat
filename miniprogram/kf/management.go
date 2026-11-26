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
	// registerBusinessURL 客服子商户创建商户
	registerBusinessURL = "https://api.weixin.qq.com/cgi-bin/business/register"
	// getBusinessURL 客服子商户拉取单个商户信息
	getBusinessURL = "https://api.weixin.qq.com/cgi-bin/business/get"
	// listBusinessURL 客服子商户拉取多个商户信息
	listBusinessURL = "https://api.weixin.qq.com/cgi-bin/business/list"
	// updateBusinessURL 客服子商户更新商户信息
	updateBusinessURL = "https://api.weixin.qq.com/cgi-bin/business/update"
)

// RegisterBusinessRequest 客服子商户创建商户请求
type RegisterBusinessRequest struct {
	AccountName      string `json:"account_name"`       // 创建商户时用到，一个account_name只能创建一次，account_name为6-30字符，必须为英文、数字、或者下划线，区分大小写。
	Nickname         string `json:"nickname"`           // 商户昵称，会在客户端会话里展示，4-30字符(中文视为2字符)，由中文、英文、数字组成
	IconMediaID      string `json:"icon_media_id"`      // 头像，图片类型，需要用临时素材接口得到，使用 新增临时素材 ，为空则不更新头像
	TransferToCommKF bool   `json:"transfer_to_commkf"` // 是否将消息转发到通用客服，false为不转发，true为转发，如果transfer_to_commkf=true时，可以使用客服管理相关api，请求参数新增business_id
}

// RegisterBusinessResponse 客服子商户创建商户响应
type RegisterBusinessResponse struct {
	util.CommonError
	BusinessID int64 `json:"business_id"` // 商户ID（官方请求与返回示例中均为数字）
}

// RegisterBusiness 客服子商户创建商户
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/kf-mgnt/kf-management/api_registerbusiness
func (kf *KF) RegisterBusiness(req *RegisterBusinessRequest) (res RegisterBusinessResponse, err error) {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return
	}

	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf("%s?access_token=%s", registerBusinessURL, accessToken), req); err != nil {
		return
	}

	err = util.DecodeWithError(response, &res, "RegisterBusiness")
	return
}

// GetBusinessRequest 客服子商户拉取单个商户信息请求
type GetBusinessRequest struct {
	BusinessID  int64  `json:"business_id"`  // 创建商户时得到的商户id,与account_name选填一个（官方请求示例中为数字）
	AccountName string `json:"account_name"` // 创建商户时用到的account_name,与business_id选填一个
}

// GetBusinessResponse 客服子商户拉取单个商户信息响应
type GetBusinessResponse struct {
	util.CommonError
	BusinessInfo BusinessInfo `json:"business_info"` // 商户信息
}

// BusinessInfo 商户信息
type BusinessInfo struct {
	BusinessID  int64  `json:"business_id"`   // 商户ID（官方请求与返回示例中均为数字）
	AccountName string `json:"account_name"`  // 商户账号名称
	Nickname    string `json:"nickname"`      // 商户昵称，会在客户端会话里展示
	IconMediaID string `json:"icon_media_id"` // 商户头像
}

// GetBusiness 客服子商户拉取单个商户信息
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/kf-mgnt/kf-management/api_getbusiness
func (kf *KF) GetBusiness(req *GetBusinessRequest) (res GetBusinessResponse, err error) {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return
	}

	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf("%s?access_token=%s", getBusinessURL, accessToken), req); err != nil {
		return
	}

	err = util.DecodeWithError(response, &res, "GetBusiness")
	return
}

// ListBusinessRequest 客服子商户拉取多个商户信息请求
type ListBusinessRequest struct {
	Offset int64 `json:"offset"` // 用于分页拉取，从0开始
	Count  int64 `json:"count"`  // 一次拉取的商户个数，最多为200
}

// ListBusinessResponse 客服子商户拉取多个商户信息响应
type ListBusinessResponse struct {
	util.CommonError
	List []BusinessInfo `json:"list"` // 商户信息列表
}

// ListBusiness 客服子商户拉取多个商户信息
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/kf-mgnt/kf-management/api_listbusiness
func (kf *KF) ListBusiness(req *ListBusinessRequest) (res ListBusinessResponse, err error) {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return
	}

	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf("%s?access_token=%s", listBusinessURL, accessToken), req); err != nil {
		return
	}

	err = util.DecodeWithError(response, &res, "ListBusiness")
	return
}

// UpdateBusinessRequest 客服子商户更新商户信息请求
type UpdateBusinessRequest struct {
	BusinessID  string `json:"business_id"`   // 创建商户时得到的商户id
	Nickname    string `json:"nickname"`      // 商户昵称，会在客户端会话里展示，4-30字符(中文视为2字符)，由中文、英文、数字组成
	IconMediaID string `json:"icon_media_id"` // 头像，图片类型，需要用临时素材接口得到，使用 新增临时素材 ，为空则不更新头像
}

// UpdateBusinessResponse 客服子商户更新商户信息响应
type UpdateBusinessResponse struct {
	util.CommonError
}

// UpdateBusiness 客服子商户更新商户信息
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/kf-mgnt/kf-management/api_updatebusiness
func (kf *KF) UpdateBusiness(req *UpdateBusinessRequest) (res UpdateBusinessResponse, err error) {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return
	}

	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf("%s?access_token=%s", updateBusinessURL, accessToken), req); err != nil {
		return
	}

	err = util.DecodeWithError(response, &res, "UpdateBusiness")
	return
}
