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

package minidrama

import (
	"context"

	"github.com/silenceper/wechat/v2/util"
)

// AuthorizeCopyright 增加版权授权
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/minidrama/authorizecopyright/api_authorizecopyright
func (s *MiniDrama) AuthorizeCopyright(ctx context.Context, in *AuthorizeCopyrightRequest) (out AuthorizeCopyrightResponse, err error) {
	var address string
	if address, err = s.requestAddress(ctx, authorizeCopyright); err != nil {
		return
	}

	var response []byte
	if response, err = util.PostJSONContext(ctx, address, in); err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "AuthorizeCopyright")
	return
}

// GetCopyrightAuthorizationList 查询版权授权信息
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/minidrama/authorizecopyright/api_getcopyrightauthorizationlist
func (s *MiniDrama) GetCopyrightAuthorizationList(ctx context.Context, in *GetCopyrightAuthorizationListRequest) (out GetCopyrightAuthorizationListResponse, err error) {
	var address string
	if address, err = s.requestAddress(ctx, getCopyrightAuthorizationList); err != nil {
		return
	}

	var response []byte
	if response, err = util.PostJSONContext(ctx, address, in); err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetCopyrightAuthorizationList")
	return
}

// GetCopyrightAuthorizedList 查询被版权授权信息
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/minidrama/authorizecopyright/api_getcopyrightauthorizedlist
func (s *MiniDrama) GetCopyrightAuthorizedList(ctx context.Context, in *GetCopyrightAuthorizedListRequest) (out GetCopyrightAuthorizedListResponse, err error) {
	var address string
	if address, err = s.requestAddress(ctx, getCopyrightAuthorizedList); err != nil {
		return
	}

	var response []byte
	if response, err = util.PostJSONContext(ctx, address, in); err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetCopyrightAuthorizedList")
	return
}

// DeauthorizeCopyright 解除版权授权
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/minidrama/authorizecopyright/api_deauthorizecopyright
func (s *MiniDrama) DeauthorizeCopyright(ctx context.Context, in *DeauthorizeCopyrightRequest) (out DeauthorizeCopyrightResponse, err error) {
	var address string
	if address, err = s.requestAddress(ctx, deauthorizeCopyright); err != nil {
		return
	}

	var response []byte
	if response, err = util.PostJSONContext(ctx, address, in); err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "DeauthorizeCopyright")
	return
}

// AuthorizeCopyrightRequest 增加版权授权请求
type AuthorizeCopyrightRequest struct {
	AuthorizationType       int     `json:"authorization_type"`                   // 授权类型：1. 授权给主体；2. 授权给小程序；
	AuthorizedAppID         string  `json:"authorized_appid,omitempty"`           // 被授权方小程序id，`authorization_type` 为 2 时必填。
	AuthorizedSubjectCertNo string  `json:"authorized_subject_cert_no,omitempty"` // 被授权方主体证件号，必须为统一社会信用代码，`authorization_type` 为 1 时必填。
	DramaIDs                []int64 `json:"drama_ids,omitempty"`                  // 要授权的受版权保护的剧目 id，最多 100 个。
	ExpireTime              int64   `json:"expire_time,omitempty"`                // 不传或者传 0，表示永久有效；否则表示授权到期的时间戳。默认值为 0。
}

// AuthorizeCopyrightResponse 增加版权授权响应
type AuthorizeCopyrightResponse struct {
	util.CommonError
	Result []*CopyrightAuthResult `json:"result"` // 剧目授权结果
}

// CopyrightAuthResult 剧目版权授权结果
type CopyrightAuthResult struct {
	DramaID int64  `json:"drama_id"` // 授权剧目 id
	ErrCode int64  `json:"errcode"`  // 该剧目授权时的错误码
	ErrMsg  string `json:"errmsg"`   // 该剧目授权时的错误原因
}

// GetCopyrightAuthorizationListRequest 查询版权授权信息请求
type GetCopyrightAuthorizationListRequest struct {
	AuthorizationType       int    `json:"authorization_type,omitempty"`         // 授权类型：1. 授权给主体；2. 授权给小程序。不填默认值为 0，表示查询全部。
	AuthorizedAppID         string `json:"authorized_appid,omitempty"`           // 被授权方小程序id，authorization_type 为 2 时必填。
	AuthorizedSubjectCertNo string `json:"authorized_subject_cert_no,omitempty"` // 被授权方主体证件号，必须为统一社会信用代码，authorization_type 为 1 时必填。
	DramaID                 int64  `json:"drama_id,omitempty"`                   // 授权剧目 id。不填默认值为 0，表示查询全部。
	Offset                  int    `json:"offset,omitempty"`                     // 偏移值，分页请求用。默认 0。
	Limit                   int    `json:"limit,omitempty"`                      // 分页大小，默认 100，最大值为 1000。
}

// GetCopyrightAuthorizationListResponse 查询版权授权信息响应
type GetCopyrightAuthorizationListResponse struct {
	util.CommonError
	TotalCount int                       `json:"total_count"` // 总数
	List       []*CopyrightAuthorization `json:"list"`        // 版权授权列表
}

// CopyrightAuthorization 版权授权信息
type CopyrightAuthorization struct {
	AuthorizerAppID         string `json:"authorizer_appid"`           // 授权方小程序 id
	AuthorizationType       int    `json:"authorization_type"`         // 授权类型
	AuthorizedAppID         string `json:"authorized_appid"`           // 被授权方小程序 id
	AuthorizedSubjectCertNo string `json:"authorized_subject_cert_no"` // 被授权方主体证件号
	DramaID                 int64  `json:"drama_id"`                   // 授权剧目 id
	AuthorizedTime          int64  `json:"authorized_time"`            // 授权时间戳
	ExpireTime              int64  `json:"expire_time"`                // 授权到期时间戳，0 表示长期有效
}

// GetCopyrightAuthorizedListRequest 查询被版权授权信息请求
type GetCopyrightAuthorizedListRequest struct {
	AuthorizerAppID string `json:"authorizer_appid,omitempty"` // 授权方小程序id，不填表示查全部。
	Offset          int    `json:"offset,omitempty"`           // 偏移值，分页请求用。默认 0。
	Limit           int    `json:"limit,omitempty"`            // 分页大小，默认 100，最大值为 1000。
}

// GetCopyrightAuthorizedListResponse 查询被版权授权信息响应
type GetCopyrightAuthorizedListResponse struct {
	util.CommonError
	TotalCount int                       `json:"total_count"` // 总数
	List       []*CopyrightAuthorization `json:"list"`        // 被版权授权列表
}

// DeauthorizeCopyrightRequest 解除版权授权请求
type DeauthorizeCopyrightRequest struct {
	AuthorizationType       int     `json:"authorization_type"`                   // 被授权方授权类型：1. 主体；2. 小程序；
	AuthorizedAppID         string  `json:"authorized_appid,omitempty"`           // 被授权方小程序id，`authorization_type` 为 2 时必填。
	AuthorizedSubjectCertNo string  `json:"authorized_subject_cert_no,omitempty"` // 被授权方主体证件号，必须为统一社会信用代码，`authorization_type` 为 1 时必填。
	DramaIDs                []int64 `json:"drama_ids,omitempty"`                  // 要解除授权的受版权保护的剧目 id，最多 100 个
}

// DeauthorizeCopyrightResponse 解除版权授权响应
type DeauthorizeCopyrightResponse struct {
	util.CommonError
	Result []*CopyrightAuthResult `json:"result"` // 剧目解授权结果
}
