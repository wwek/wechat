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

package plugin

import (
	"fmt"

	"github.com/silenceper/wechat/v2/miniprogram/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	// managePluginURL 插件管理
	managePluginURL = "https://api.weixin.qq.com/wxa/plugin?access_token=%s"
	// managePluginApplicationURL 插件申请管理
	managePluginApplicationURL = "https://api.weixin.qq.com/wxa/devplugin?access_token=%s"
)

// Plugin 插件管理
type Plugin struct {
	*context.Context
}

// NewPlugin 实例化
func NewPlugin(ctx *context.Context) *Plugin {
	return &Plugin{ctx}
}

// ManagePluginRequest 插件管理请求
type ManagePluginRequest struct {
	Action      string `json:"action"`       // action 可取值为"list"、"apply"、"unbind"、"update"，分别表示“查询已添加的插件列表”、“申请使用插件”、“删除已添加的插件”、“快速更新插件版本号”
	PluginAppID string `json:"plugin_appid"` // 插件的 appid
	Reason      string `json:"reason"`       // 当action是"apply"时必填，申请原因
	UserVersion string `json:"user_version"` // 当action是"update"时使用。升级至版本号，要求此插件版本支持快速更新
}

// ManagePluginListItem 插件信息列表项
type ManagePluginListItem struct {
	AppID      string `json:"appid"`      // 插件 appId
	Status     int64  `json:"status"`     // 插件状态。1表示申请中；2表示申请通过；3表示被拒绝；4表示申请已超时
	Nickname   string `json:"nickname"`   // 插件昵称
	HeadImgURL string `json:"headimgurl"` // 插件头像
}

// ManagePluginResponse 插件管理响应
type ManagePluginResponse struct {
	util.CommonError
	PluginList []ManagePluginListItem `json:"plugin_list"` // 申请或使用中的插件信息列表，当 action == 'list' 时返回
}

// ManagePlugin 插件管理
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/plugin-management/api_manageplugin
func (s *Plugin) ManagePlugin(in *ManagePluginRequest) (out ManagePluginResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(managePluginURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "ManagePlugin")
	return
}

// ManagePluginApplicationRequest 插件申请管理请求
type ManagePluginApplicationRequest struct {
	Action string `json:"action"` // dev_agree表示同意申请、dev_refuse表示拒绝申请、dev_delete表示删除已拒绝的申请者、dev_apply_list表示获取当前所有插件使用方信息
	AppID  string `json:"appid"`  // action为"dev_agree"时填写，使用者的 appid，同意申请时填写
	Page   int64  `json:"page"`   // action为"dev_apply_list"时填写，要拉取第几页的数据
	Num    int64  `json:"num"`    // action为"dev_apply_list"时填写，每页的记录数
	Reason string `json:"reason"` // action为"dev_refuse"时填写，拒绝理由
}

// ManagePluginApplicationCategory 使用者的类目
type ManagePluginApplicationCategory struct {
	First  string `json:"first"`  // 一级类目名称
	Second string `json:"second"` // 二级类目名称
}

// ManagePluginApplicationItem 插件使用方信息
type ManagePluginApplicationItem struct {
	AppID      string                            `json:"appid"`       // 使用者的appid
	Status     int64                             `json:"status"`      // 插件状态
	Nickname   string                            `json:"nickname"`    // 使用者的昵称
	HeadImgURL string                            `json:"headimgurl"`  // 使用者的头像
	Reason     string                            `json:"reason"`      // 使用者的申请说明
	ApplyURL   string                            `json:"apply_url"`   // 使用者的小程序码
	CreateTime string                            `json:"create_time"` // 使用者的申请时间
	Categories []ManagePluginApplicationCategory `json:"categories"`  // 使用者的类目
}

// ManagePluginApplicationResponse 插件申请管理响应
type ManagePluginApplicationResponse struct {
	util.CommonError
	ApplyList []ManagePluginApplicationItem `json:"apply_list"` // 插件使用方列表，action为"dev_apply_list"时返回
}

// ManagePluginApplication 插件申请管理
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/plugin-management/api_managepluginapplication
func (s *Plugin) ManagePluginApplication(in *ManagePluginApplicationRequest) (out ManagePluginApplicationResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(managePluginApplicationURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "ManagePluginApplication")
	return
}
