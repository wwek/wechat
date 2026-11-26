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

package novel

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// delBookAuthURL 删除指定的授权关系
	delBookAuthURL = "https://api.weixin.qq.com/wxa/book/delbookauth?access_token=%s"
	// delBookAuthByAppIDURL 删除指定的账号授权关系
	delBookAuthByAppIDURL = "https://api.weixin.qq.com/wxa/book/delbookauthbyappid?access_token=%s"
	// addBookAuthURL 新增账号-小说授权
	addBookAuthURL = "https://api.weixin.qq.com/wxa/book/addbookauth?access_token=%s"
	// addBookAuthByAppIDURL 新增账号-账号授权
	addBookAuthByAppIDURL = "https://api.weixin.qq.com/wxa/book/addbookauthbyappid?access_token=%s"
	// queryBookAuthURL 查看主授权关系列表-查看被授权关系列表
	queryBookAuthURL = "https://api.weixin.qq.com/wxa/book/querybookauth?access_token=%s"
	// queryBookAuthV2URL 查看账号主授权关系列表
	queryBookAuthV2URL = "https://api.weixin.qq.com/wxa/book/querybookauthv2?access_token=%s"
)

// DelBookAuthRequest 删除指定的授权关系请求
type DelBookAuthRequest struct {
	BookID       string `json:"book_id"`       // 作品 id
	GranteeAppID string `json:"grantee_appid"` // 被授权账号 appid
}

// DelBookAuthResponse 删除指定的授权关系响应
type DelBookAuthResponse struct {
	util.CommonError
}

// DelBookAuth 删除指定的授权关系
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/novel/auth/api_delbookauth
func (s *Novel) DelBookAuth(in *DelBookAuthRequest) (out DelBookAuthResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(delBookAuthURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "DelBookAuth")
	return
}

// DelBookAuthByAppIDRequest 删除指定的账号授权关系请求
type DelBookAuthByAppIDRequest struct {
	GranteeAppID string `json:"grantee_appid"` // 被授权账号 appid
}

// DelBookAuthByAppIDResponse 删除指定的账号授权关系响应
type DelBookAuthByAppIDResponse struct {
	util.CommonError
}

// DelBookAuthByAppID 删除指定的账号授权关系
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/novel/auth/api_delbookauthbyappid
func (s *Novel) DelBookAuthByAppID(in *DelBookAuthByAppIDRequest) (out DelBookAuthByAppIDResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(delBookAuthByAppIDURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "DelBookAuthByAppID")
	return
}

// BookAuthInfo 账号-小说授权信息
type BookAuthInfo struct {
	BookID       string `json:"book_id"`       // 作品 id
	GranteeAppID string `json:"grantee_appid"` // 被授权账号 appid
	ExpireTime   int64  `json:"expire_time"`   // 授权到期时间，unix 时间戳，单位秒，按 CST 时区
}

// AddBookAuthRequest 新增账号-小说授权请求
type AddBookAuthRequest struct {
	Books []BookAuthInfo `json:"books"` // 授权关系列表
}

// AddBookAuthResult 授权结果
type AddBookAuthResult struct {
	ErrCode int64 `json:"errcode"` // 内层 errcode，表示每条授权关系是成功/失败
}

// AddBookAuthResponse 新增账号-小说授权响应
type AddBookAuthResponse struct {
	util.CommonError
	Results []AddBookAuthResult `json:"results"` // 每条授权关系的结果
}

// AddBookAuth 新增账号-小说授权
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/novel/auth/api_addbookauth
func (s *Novel) AddBookAuth(in *AddBookAuthRequest) (out AddBookAuthResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(addBookAuthURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "AddBookAuth")
	return
}

// BookAuthByAppIDInfo 账号-账号授权信息
type BookAuthByAppIDInfo struct {
	GranteeAppID string `json:"grantee_appid"` // 被授权账号 appid
	ExpireTime   int64  `json:"expire_time"`   // 授权到期时间，unix 时间戳，单位秒，按 CST 时区
}

// AddBookAuthByAppIDRequest 新增账号-账号授权请求
type AddBookAuthByAppIDRequest struct {
	Infos []BookAuthByAppIDInfo `json:"infos"` // 被授权账号列表，上限一次 20 条
}

// AddBookAuthByAppIDResponse 新增账号-账号授权响应
type AddBookAuthByAppIDResponse struct {
	util.CommonError
	Results []AddBookAuthResult `json:"results"` // 授权结果列表
}

// AddBookAuthByAppID 新增账号-账号授权
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/novel/auth/api_addbookauthbyappid
func (s *Novel) AddBookAuthByAppID(in *AddBookAuthByAppIDRequest) (out AddBookAuthByAppIDResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(addBookAuthByAppIDURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "AddBookAuthByAppID")
	return
}

// QueryBookAuthRequest 查看主授权关系列表-查看被授权关系列表请求
type QueryBookAuthRequest struct {
	Type   int64  `json:"type"`    // 填 0 或不设置查授权列表，填 1 查被授权列表
	Offset int64  `json:"offset"`  // 结果数据偏移位置，从 0 开始，用于翻页
	Count  int64  `json:"count"`   // 获取记录的数量，上限一次 30 条，用于翻页
	IsSum  bool   `json:"is_sum"`  // 填 true 返回总数，填 false 不返回
	BookID string `json:"book_id"` // 过滤出等于该 book_id 的授权结果
}

// BookAuthResult 授权关系
type BookAuthResult struct {
	BookID       string `json:"book_id"`       // 作品 id
	GrantorAppID string `json:"grantor_appid"` // 主授权账号 appid
	GranteeAppID string `json:"grantee_appid"` // 被授权账号 appid
	ExpireTime   int64  `json:"expire_time"`   // 授权到期时间，unix 时间戳，单位秒，按 CST 时区
	Sum          int64  `json:"sum"`           // 总数，is_sum 为 true 时返回
}

// QueryBookAuthResponse 查看主授权关系列表-查看被授权关系列表响应
type QueryBookAuthResponse struct {
	util.CommonError
	Results []BookAuthResult `json:"results"` // 授权关系列表
}

// QueryBookAuth 查看主授权关系列表-查看被授权关系列表
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/novel/auth/api_querybookauth
func (s *Novel) QueryBookAuth(in *QueryBookAuthRequest) (out QueryBookAuthResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(queryBookAuthURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "QueryBookAuth")
	return
}

// QueryBookAuthV2Request 查看账号主授权关系列表请求
type QueryBookAuthV2Request struct {
	Type         int64    `json:"type"`          // 填 0 或不设置查授权列表，填 1 查被授权列表
	Count        int64    `json:"count"`         // 获取记录的数量，上限一次 100 条，用于分页
	Cursor       string   `json:"cursor"`        // 分页游标，不填或填空字符串表示从第一页开始，后一页填充上一页响应的 next_cursor
	GrantorAppID string   `json:"grantor_appid"` // 用于查询指定授权方授权给当前账户的小说列表
	BookIDs      []string `json:"book_ids"`      // 用于查询指定小说的授权信息，上限一次 30 条，优先级高于 grantor_appid
}

// BookAuthAppIDResult 被授权方列表项
type BookAuthAppIDResult struct {
	GrantorAppID string `json:"grantor_appid"` // 主授权账号 appid
	GranteeAppID string `json:"grantee_appid"` // 被授权账号 appid
	ExpireTime   int64  `json:"expire_time"`   // 授权到期时间，unix 时间戳，单位秒，按 CST 时区
	BookID       string `json:"book_id"`       // 作品 id，查询指定授权方授权给当前账户的小说列表时会返回
}

// QueryBookAuthV2Response 查看账号主授权关系列表响应
type QueryBookAuthV2Response struct {
	util.CommonError
	AppIDResults []BookAuthAppIDResult `json:"appid_results"` // 被授权方列表，按授权过期时间从早到晚排序；type=1 且不传 grantor_appid/book_ids 时返回
	BookResults  []BookAuthAppIDResult `json:"book_results"`  // 作品授权列表，type=1 且传了 grantor_appid 或 book_ids 时返回
	NextCursor   string                `json:"next_cursor"`   // 下一页游标，请求时填入 cursor 参数
}

// QueryBookAuthV2 查看账号主授权关系列表
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/novel/auth/api_querybookauthv2
func (s *Novel) QueryBookAuthV2(in *QueryBookAuthV2Request) (out QueryBookAuthV2Response, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(queryBookAuthV2URL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "QueryBookAuthV2")
	return
}
