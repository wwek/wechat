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

// AuthorizeApp 增加账号授权
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/minidrama/authorizeapp/api_authorizeapp
func (s *MiniDrama) AuthorizeApp(ctx context.Context, in *AuthorizeAppRequest) (out AuthorizeAppResponse, err error) {
	var address string
	if address, err = s.requestAddress(ctx, authorizeApp); err != nil {
		return
	}

	var response []byte
	if response, err = util.PostJSONContext(ctx, address, in); err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "AuthorizeApp")
	return
}

// GetAuthorizedObjects 查询被授权信息
// 该接口在账号授权和剧目授权两个场景下均可用，参数与返回均一致。
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/minidrama/authorizeapp/api_getauthorizedobjects
func (s *MiniDrama) GetAuthorizedObjects(ctx context.Context, in *GetAuthorizedObjectsRequest) (out GetAuthorizedObjectsResponse, err error) {
	var address string
	if address, err = s.requestAddress(ctx, getAuthorizedObjects); err != nil {
		return
	}

	var response []byte
	if response, err = util.PostJSONContext(ctx, address, in); err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetAuthorizedObjects")
	return
}

// GetAuthorizeApps 查询账号授权信息
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/minidrama/authorizeapp/api_getauthorizeapps
func (s *MiniDrama) GetAuthorizeApps(ctx context.Context) (out GetAuthorizeAppsResponse, err error) {
	var address string
	if address, err = s.requestAddress(ctx, getAuthorizeApps); err != nil {
		return
	}

	var response []byte
	if response, err = util.PostJSONContext(ctx, address, map[string]interface{}{}); err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetAuthorizeApps")
	return
}

// DeauthorizeApp 解除账号授权
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/minidrama/authorizeapp/api_deauthorizeapp
func (s *MiniDrama) DeauthorizeApp(ctx context.Context, in *DeauthorizeAppRequest) (out DeauthorizeAppResponse, err error) {
	var address string
	if address, err = s.requestAddress(ctx, deauthorizeApp); err != nil {
		return
	}

	var response []byte
	if response, err = util.PostJSONContext(ctx, address, in); err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "DeauthorizeApp")
	return
}

// AuthorizeAppRequest 增加账号授权请求
type AuthorizeAppRequest struct {
	AuthorizedAppID string `json:"authorized_appid"`            // 被授权方小程序id。
	AuthzExpireTime int64  `json:"authz_expire_time,omitempty"` // 不传或者传0，表示永久有效；否则表示授权到期的时间戳。
}

// AuthorizeAppResponse 增加账号授权响应
type AuthorizeAppResponse struct {
	util.CommonError
}

// GetAuthorizedObjectsRequest 查询被授权信息请求
type GetAuthorizedObjectsRequest struct {
	AuthorizerAppID string `json:"authorizer_appid,omitempty"` // 授权方小程序id。不填表示查询所有授权方。
	Offset          int    `json:"offset,omitempty"`           // 偏移值，分页请求用。默认0。
	Limit           int    `json:"limit,omitempty"`            // 分页大小，默认100，最大值为1000。
}

// GetAuthorizedObjectsResponse 查询被授权信息响应
type GetAuthorizedObjectsResponse struct {
	util.CommonError
	TotalCount int                 `json:"total_count"` // 总数
	Objects    []*AuthorizedObject `json:"objects"`     // 授权对象列表
}

// AuthorizedObject 被授权对象
type AuthorizedObject struct {
	DramaID         int64  `json:"drama_id"`          // 授权的剧目ID
	AuthorizerAppID string `json:"authorizer_appid"`  // 授权方小程序id
	AuthorizedTime  int64  `json:"authorized_time"`   // 授权时间戳
	AuthzExpireTime int64  `json:"authz_expire_time"` // 授权到期时间戳。0表示长期有效。
}

// GetAuthorizeAppsResponse 查询账号授权信息响应
type GetAuthorizeAppsResponse struct {
	util.CommonError
	Objects []*AuthorizedApp `json:"objects"` // 被授权账号列表
}

// AuthorizedApp 被授权账号
type AuthorizedApp struct {
	AuthorizedAppID string `json:"authorized_appid"`  // 被授权方小程序id
	AuthorizedTime  int64  `json:"authorized_time"`   // 授权时间戳
	AuthzExpireTime int64  `json:"authz_expire_time"` // 授权到期时间戳。0表示长期有效。
}

// DeauthorizeAppRequest 解除账号授权请求
type DeauthorizeAppRequest struct {
	AuthorizedAppID string `json:"authorized_appid"` // 被授权方小程序id。
}

// DeauthorizeAppResponse 解除账号授权响应
type DeauthorizeAppResponse struct {
	util.CommonError
}
