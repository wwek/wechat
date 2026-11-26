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

package livebroadcast

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// getRoleListURL 查询成员列表
	getRoleListURL = "https://api.weixin.qq.com/wxaapi/broadcast/role/getrolelist?%s"
	// deleteRoleURL 移除成员角色
	deleteRoleURL = "https://api.weixin.qq.com/wxaapi/broadcast/role/deleterole?access_token=%s"
	// addRoleURL 设置成员角色
	addRoleURL = "https://api.weixin.qq.com/wxaapi/broadcast/role/addrole?access_token=%s"
)

// RoleInfo 成员角色信息
type RoleInfo struct {
	Headingimg      string  `json:"headingimg"`      // 微信用户头像url
	Nickname        string  `json:"nickname"`        // 微信用户昵称
	OpenID          string  `json:"openid"`          // openid
	RoleList        []int64 `json:"roleList"`        // 具有的身份，[0-超级管理员，1-管理员，2-主播，3-运营者]
	UpdateTimestamp string  `json:"updateTimestamp"` // 更新时间
	Username        string  `json:"username"`        // 微信号
}

// GetRoleListRequest 查询成员列表请求
type GetRoleListRequest struct {
	Offset  int64  `json:"offset"`  // 分页偏移量
	Limit   int64  `json:"limit"`   // 分页大小
	Keyword string `json:"keyword"` // 搜索的微信昵称或微信号
}

// GetRoleListResponse 查询成员列表响应
type GetRoleListResponse struct {
	util.CommonError
	Total int64      `json:"total"` // 总个数
	List  []RoleInfo `json:"list"`  // 角色列表
}

// GetRoleList 查询成员列表
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/role-management/api_getrolelist
func (s *LiveBroadcast) GetRoleList(in *GetRoleListRequest) (out GetRoleListResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	params := map[string]interface{}{
		"access_token": accessToken,
	}
	if in.Offset > 0 {
		params["offset"] = in.Offset
	}
	if in.Limit > 0 {
		params["limit"] = in.Limit
	}
	if in.Keyword != "" {
		params["keyword"] = in.Keyword
	}
	query := util.Query(params)

	response, err := util.HTTPGet(fmt.Sprintf(getRoleListURL, query))
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetRoleList")
	return
}

// DeleteRoleRequest 移除成员角色请求
type DeleteRoleRequest struct {
	Username string `json:"username"` // 用户的微信号
	Role     int64  `json:"role"`     // 删除用户的角色
}

// DeleteRoleResponse 移除成员角色响应
type DeleteRoleResponse struct {
	util.CommonError
}

// DeleteRole 移除成员角色
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/role-management/api_deleterole
func (s *LiveBroadcast) DeleteRole(in *DeleteRoleRequest) (out DeleteRoleResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(deleteRoleURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "DeleteRole")
	return
}

// SetRoleRequest 设置成员角色请求
type SetRoleRequest struct {
	Username string `json:"username"` // 用户的微信号
	Role     int64  `json:"role"`     // 设置用户的角色
}

// SetRoleResponse 设置成员角色响应
type SetRoleResponse struct {
	util.CommonError
	CodeURL string `json:"codeurl"` // 如果主播未实名认证，需要先前往“ 小程序直播 ”小程序进行实名验证
}

// SetRole 设置成员角色
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/role-management/api_getrolelistdw
func (s *LiveBroadcast) SetRole(in *SetRoleRequest) (out SetRoleResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(addRoleURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "SetRole")
	return
}
