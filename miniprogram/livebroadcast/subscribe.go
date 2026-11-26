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
	// pushMessageURL 发送直播开始事件
	pushMessageURL = "https://api.weixin.qq.com/wxa/business/push_message?access_token=%s"
	// getFollowersURL 获取长期订阅用户
	getFollowersURL = "https://api.weixin.qq.com/wxa/business/get_wxa_followers?access_token=%s"
)

// Follower 长期订阅用户
type Follower struct {
	OpenID        string `json:"openid"`         // 长期订阅用户OpenId
	SubscribeTime int64  `json:"subscribe_time"` // 长期订阅用户订阅时间
	RoomID        int64  `json:"room_id"`        // 用户订阅时所处房间
	RoomStatus    int64  `json:"room_status"`    // 用户订阅时房间状态
}

// PushMessageRequest 发送直播开始事件请求
type PushMessageRequest struct {
	RoomID     int64    `json:"room_id"`     // 直播开始事件的房间ID
	UserOpenID []string `json:"user_openid"` // 接收该群发开播事件的订阅用户OpenId列表
}

// PushMessageResponse 发送直播开始事件响应
type PushMessageResponse struct {
	util.CommonError
	MessageID string `json:"message_id"` // 此次群发消息的标识ID，用于对应【长期订阅群发结果回调】的message_id
}

// PushMessage 发送直播开始事件
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/subscribe-management/api_pushmessage
func (s *LiveBroadcast) PushMessage(in *PushMessageRequest) (out PushMessageResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(pushMessageURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "PushMessage")
	return
}

// GetFollowersRequest 获取长期订阅用户请求
type GetFollowersRequest struct {
	Limit     int64 `json:"limit"`      // 获取长期订阅用户的个数限制，默认200，最大2000
	PageBreak int64 `json:"page_break"` // 翻页标记，获取第一页时不带，第二页开始需带上上一页返回结果中的page_break
}

// GetFollowersResponse 获取长期订阅用户响应
type GetFollowersResponse struct {
	util.CommonError
	Followers []Follower `json:"followers"`  // 长期订阅用户列表
	PageBreak int64      `json:"page_break"` // 翻页标记，获取下一页时带上该值
}

// GetFollowers 获取长期订阅用户
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/subscribe-management/api_getfollowers
func (s *LiveBroadcast) GetFollowers(in *GetFollowersRequest) (out GetFollowersResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(getFollowersURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetFollowers")
	return
}
