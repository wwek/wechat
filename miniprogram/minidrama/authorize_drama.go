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
	"strconv"
	"strings"

	"github.com/silenceper/wechat/v2/util"
)

// 本文件为剧目授权相关接口。
// 说明：剧目授权分类下的“查询被授权信息”与账号授权分类下的“查询被授权信息”
// 为同一个接口（/wxa/sec/vod/getauthorizedobjects），参数与返回一致，统一在 authorize_app.go 中实现。

// AuthorizeDrama 增加剧目授权
// 该接口为 GET 请求，参数通过 query 传递，剧目 id 列表以英文逗号分隔。
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/minidrama/authorizedrama/api_authorizedrama
func (s *MiniDrama) AuthorizeDrama(ctx context.Context, in *AuthorizeDramaRequest) (out AuthorizeDramaResponse, err error) {
	var address string
	if address, err = s.requestAddress(ctx, authorizeDrama); err != nil {
		return
	}

	dramaIDs := make([]string, 0, len(in.DramaIDs))
	for _, dramaID := range in.DramaIDs {
		dramaIDs = append(dramaIDs, strconv.FormatInt(dramaID, 10))
	}
	params := map[string]interface{}{
		"drama_id":   strings.Join(dramaIDs, ","),
		"authorized": in.Authorized,
	}
	if in.AuthzExpireTime != 0 {
		params["authz_expire_time"] = in.AuthzExpireTime
	}
	address = address + "&" + util.Query(params)

	var response []byte
	if response, err = util.HTTPGetContext(ctx, address); err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "AuthorizeDrama")
	return
}

// GetAuthorizeObjects 查询剧目授权信息
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/minidrama/authorizedrama/api_getauthorizeobjects
func (s *MiniDrama) GetAuthorizeObjects(ctx context.Context, in *GetAuthorizeObjectsRequest) (out GetAuthorizeObjectsResponse, err error) {
	var address string
	if address, err = s.requestAddress(ctx, getAuthorizeObjects); err != nil {
		return
	}

	var response []byte
	if response, err = util.PostJSONContext(ctx, address, in); err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetAuthorizeObjects")
	return
}

// DeauthorizeDrama 解除剧目授权
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/minidrama/authorizedrama/api_deauthorizedrama
func (s *MiniDrama) DeauthorizeDrama(ctx context.Context, in *DeauthorizeDramaRequest) (out DeauthorizeDramaResponse, err error) {
	var address string
	if address, err = s.requestAddress(ctx, deauthorizeDrama); err != nil {
		return
	}

	var response []byte
	if response, err = util.PostJSONContext(ctx, address, in); err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "DeauthorizeDrama")
	return
}

// AuthorizeDramaRequest 增加剧目授权请求
type AuthorizeDramaRequest struct {
	DramaIDs        []int64 `json:"drama_id"`                    // 授权的剧目ID
	Authorized      string  `json:"authorized"`                  // 被授权方小程序id。
	AuthzExpireTime int64   `json:"authz_expire_time,omitempty"` // 不传或者传0，表示永久有效；否则表示授权到期的时间戳
}

// AuthorizeDramaResponse 增加剧目授权响应
type AuthorizeDramaResponse struct {
	util.CommonError
	Result []*DramaAuthResult `json:"result"` // 剧目授权结果
}

// DramaAuthResult 剧目授权结果
type DramaAuthResult struct {
	DramaID int64  `json:"drama_id"` // 剧目ID。
	ErrCode int64  `json:"errcode"`  // 错误码。
	ErrMsg  string `json:"errmsg"`   // 错误原因。
}

// GetAuthorizeObjectsRequest 查询剧目授权信息请求
type GetAuthorizeObjectsRequest struct {
	DramaID         int64  `json:"drama_id,omitempty"`         // 查询的剧目ID。不填表示查询所有剧目。
	AuthorizedAppID string `json:"authorized_appid,omitempty"` // 查询的被授权方小程序id。不填表示查询所有被授权方。
	Offset          int    `json:"offset,omitempty"`           // 偏移值，分页请求用。默认0。
	Limit           int    `json:"limit,omitempty"`            // 分页大小，默认100，最大值为1000。
}

// GetAuthorizeObjectsResponse 查询剧目授权信息响应
type GetAuthorizeObjectsResponse struct {
	util.CommonError
	TotalCount int                `json:"total_count"` // 总数
	Objects    []*AuthorizeObject `json:"objects"`     // 剧目授权对象列表
}

// AuthorizeObject 剧目授权对象
type AuthorizeObject struct {
	DramaID         int64  `json:"drama_id"`          // 授权的剧目ID
	AuthorizedAppID string `json:"authorized_appid"`  // 被授权方小程序id
	AuthorizedTime  int64  `json:"authorized_time"`   // 授权时间戳
	AuthzExpireTime int64  `json:"authz_expire_time"` // 授权到期时间戳。0表示长期有效。
}

// DeauthorizeDramaRequest 解除剧目授权请求
type DeauthorizeDramaRequest struct {
	DramaIDs        []int64 `json:"drama_id"`         // 解除授权剧目ID。
	AuthorizedAppID string  `json:"authorized_appid"` // 被授权方小程序id。
}

// DeauthorizeDramaResponse 解除剧目授权响应
type DeauthorizeDramaResponse struct {
	util.CommonError
	Result []*DramaAuthResult `json:"result"` // 剧目解授权结果
}
