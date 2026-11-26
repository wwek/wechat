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
	"net/url"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// saleGoodsURL 上下架商品
	saleGoodsURL = "https://api.weixin.qq.com/wxaapi/broadcast/goods/onsale?access_token=%s"
	// downloadGoodsVideoURL 下载商品讲解视频
	downloadGoodsVideoURL = "https://api.weixin.qq.com/wxaapi/broadcast/goods/getVideo?access_token=%s"
	// modifySubAnchorURL 修改主播副号
	modifySubAnchorURL = "https://api.weixin.qq.com/wxaapi/broadcast/room/modifysubanchor?access_token=%s"
	// modifyAssistantURL 修改直播间小助手
	modifyAssistantURL = "https://api.weixin.qq.com/wxaapi/broadcast/room/modifyassistant?access_token=%s"
	// createRoomURL 创建直播间
	createRoomURL = "https://api.weixin.qq.com/wxaapi/broadcast/room/create?access_token=%s"
	// deleteSubAnchorURL 删除主播副号
	deleteSubAnchorURL = "https://api.weixin.qq.com/wxaapi/broadcast/room/deletesubanchor?access_token=%s"
	// deleteRoomURL 删除直播间
	deleteRoomURL = "https://api.weixin.qq.com/wxaapi/broadcast/room/deleteroom?access_token=%s"
	// deleteRoomGoodsURL 删除直播间商品
	deleteRoomGoodsURL = "https://api.weixin.qq.com/wxaapi/broadcast/goods/deleteInRoom?access_token=%s"
	// removeAssistantURL 删除直播间小助手
	removeAssistantURL = "https://api.weixin.qq.com/wxaapi/broadcast/room/removeassistant?access_token=%s"
	// updateReplayURL 回放功能管理
	updateReplayURL = "https://api.weixin.qq.com/wxaapi/broadcast/room/updatereplay?access_token=%s"
	// updateFeedPublicURL 官方收录管理
	updateFeedPublicURL = "https://api.weixin.qq.com/wxaapi/broadcast/room/updatefeedpublic?access_token=%s"
	// updateKFURL 客服功能管理
	updateKFURL = "https://api.weixin.qq.com/wxaapi/broadcast/room/updatekf?access_token=%s"
	// importGoodsURL 导入商品
	importGoodsURL = "https://api.weixin.qq.com/wxaapi/broadcast/room/addgoods?access_token=%s"
	// pushGoodsURL 推送商品
	pushGoodsURL = "https://api.weixin.qq.com/wxaapi/broadcast/goods/push?access_token=%s"
	// getAssistantListURL 查询直播间小助手
	getAssistantListURL = "https://api.weixin.qq.com/wxaapi/broadcast/room/getassistantlist?%s"
	// addSubAnchorURL 添加主播副号
	addSubAnchorURL = "https://api.weixin.qq.com/wxaapi/broadcast/room/addsubanchor?access_token=%s"
	// addAssistantURL 添加管理直播间小助手
	addAssistantURL = "https://api.weixin.qq.com/wxaapi/broadcast/room/addassistant?access_token=%s"
	// getDefaultGoodsKeyURL 直播挂件获取全局KEY
	getDefaultGoodsKeyURL = "https://api.weixin.qq.com/wxaapi/broadcast/goods/getkey?%s"
	// setDefaultGoodsKeyURL 直播挂件设置全局KEY
	setDefaultGoodsKeyURL = "https://api.weixin.qq.com/wxaapi/broadcast/goods/setkey?access_token=%s"
	// sortGoodsURL 直播间商品排序
	sortGoodsURL = "https://api.weixin.qq.com/wxaapi/broadcast/goods/sort?access_token=%s"
	// updateCommentURL 禁言管理
	updateCommentURL = "https://api.weixin.qq.com/wxaapi/broadcast/room/updatecomment?access_token=%s"
	// editRoomURL 编辑直播间
	editRoomURL = "https://api.weixin.qq.com/wxaapi/broadcast/room/editroom?access_token=%s"
	// getSubAnchorURL 获取主播副号
	getSubAnchorURL = "https://api.weixin.qq.com/wxaapi/broadcast/room/getsubanchor?%s"
	// getSharedCodeURL 获取直播间分享二维码
	getSharedCodeURL = "https://api.weixin.qq.com/wxaapi/broadcast/room/getsharedcode?%s"
	// getLiveInfoURL 获取直播间列表和回放
	getLiveInfoURL = "https://api.weixin.qq.com/wxa/business/getliveinfo?access_token=%s"
	// getPushUrlURL 获取直播间推流地址
	getPushUrlURL = "https://api.weixin.qq.com/wxaapi/broadcast/room/getpushurl?%s"
)

// Assistant 直播间小助手
type Assistant struct {
	Timestamp int64  `json:"timestamp"` // 修改时间
	Headimg   string `json:"headimg"`   // 头像
	Nickname  string `json:"nickname"`  // 昵称
	Alias     string `json:"alias"`     // 微信号
	OpenID    string `json:"openid"`    // openid
}

// AssistantUser 小助手用户
type AssistantUser struct {
	Username string `json:"username"` // 用户微信号
	Nickname string `json:"nickname"` // 用户昵称
}

// GoodsIDItem 商品ID列表项
type GoodsIDItem struct {
	GoodsID int64 `json:"goodsId"` // 商品id
}

// RoomGoods 直播间商品
type RoomGoods struct {
	Name            string  `json:"name"`              // 商品名称
	CoverImg        string  `json:"cover_img"`         // 商品封面图链接
	URL             string  `json:"url"`               // 商品小程序路径
	Price           float64 `json:"price"`             // 商品价格（分）
	Price2          float64 `json:"price2"`            // 商品价格，使用方式看price_type
	PriceType       int64   `json:"price_type"`        // 价格类型，1：一口价（只需要传入price，price2不传） 2：价格区间（price字段为左边界，price2字段为右边界，price和price2必传） 3：显示折扣价（price字段为原价，price2字段为现价，price和price2必传）
	GoodsID         int64   `json:"goods_id"`          // 商品id
	ThirdPartyAppid string  `json:"third_party_appid"` // 第三方商品appid，当前小程序商品则为空
}

// RoomInfo 直播间信息
type RoomInfo struct {
	Name          string      `json:"name"`            // 直播间名称
	CoverImg      string      `json:"cover_img"`       // 直播间背景图链接
	StartTime     int64       `json:"start_time"`      // 直播间开始时间，列表按照start_time降序排列
	EndTime       int64       `json:"end_time"`        // 直播计划结束时间
	AnchorName    string      `json:"anchor_name"`     // 主播名
	RoomID        int64       `json:"roomid"`          // 直播间ID
	Goods         []RoomGoods `json:"goods"`           // 商品
	LiveStatus    int64       `json:"live_status"`     // 直播间状态。101：直播中，102：未开始，103已结束，104禁播，105：暂停，106：异常，107：已过期
	ShareImg      string      `json:"share_img"`       // 直播间分享图链接
	LiveType      int64       `json:"live_type"`       // 直播类型，1 推流 0 手机直播
	CloseLike     int64       `json:"close_like"`      // 是否关闭点赞 【0：开启，1：关闭】
	CloseGoods    int64       `json:"close_goods"`     // 是否关闭货架 【0：开启，1：关闭】
	CloseComment  int64       `json:"close_comment"`   // 是否关闭评论 【0：开启，1：关闭】
	CloseKf       int64       `json:"close_kf"`        // 是否关闭客服 【0：开启，1：关闭】
	CloseReplay   int64       `json:"close_replay"`    // 是否关闭回放 【0：开启，1：关闭】
	IsFeedsPublic int64       `json:"is_feeds_public"` // 是否开启官方收录，1 开启，0 关闭
	CreaterOpenid string      `json:"creater_openid"`  // 创建者openid
	FeedsImg      string      `json:"feeds_img"`       // 官方收录封面
}

// LiveReplay 直播回放
type LiveReplay struct {
	CreateTime string `json:"create_time"` // 回放视频创建时间
	ExpireTime string `json:"expire_time"` // 回放视频url过期时间
	MediaURL   string `json:"media_url"`   // 回放视频链接
}

// SaleGoodsRequest 上下架商品请求
type SaleGoodsRequest struct {
	RoomID  int64 `json:"roomId"`  // 房间ID
	GoodsID int64 `json:"goodsId"` // 商品ID
	OnSale  int64 `json:"onSale"`  // 上下架 【0：下架，1：上架】
}

// SaleGoodsResponse 上下架商品响应
type SaleGoodsResponse struct {
	util.CommonError
}

// SaleGoods 上下架商品
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/studio-management/api_salegoods
func (s *LiveBroadcast) SaleGoods(in *SaleGoodsRequest) (out SaleGoodsResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(saleGoodsURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "SaleGoods")
	return
}

// DownloadGoodsVideoRequest 下载商品讲解视频请求
type DownloadGoodsVideoRequest struct {
	RoomID  int64 `json:"roomId"`  // 房间ID
	GoodsID int64 `json:"goodsId"` // 商品ID
}

// DownloadGoodsVideoResponse 下载商品讲解视频响应
type DownloadGoodsVideoResponse struct {
	util.CommonError
	URL string `json:"url"` // 讲解链接
}

// DownloadGoodsVideo 下载商品讲解视频
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/studio-management/api_downloadgoodsvideo
func (s *LiveBroadcast) DownloadGoodsVideo(in *DownloadGoodsVideoRequest) (out DownloadGoodsVideoResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(downloadGoodsVideoURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "DownloadGoodsVideo")
	return
}

// ModifySubAnchorRequest 修改主播副号请求
type ModifySubAnchorRequest struct {
	RoomID   int64  `json:"roomId"`   // 房间ID
	Username string `json:"username"` // 微信号
}

// ModifySubAnchorResponse 修改主播副号响应
type ModifySubAnchorResponse struct {
	util.CommonError
}

// ModifySubAnchor 修改主播副号
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/studio-management/api_modifysubanchor
func (s *LiveBroadcast) ModifySubAnchor(in *ModifySubAnchorRequest) (out ModifySubAnchorResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(modifySubAnchorURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "ModifySubAnchor")
	return
}

// ModifyAssistantRequest 修改直播间小助手请求
type ModifyAssistantRequest struct {
	RoomID   int64  `json:"roomId"`   // 房间ID
	Username string `json:"username"` // 用户微信号
	Nickname string `json:"nickname"` // 用户微信昵称
}

// ModifyAssistantResponse 修改直播间小助手响应
type ModifyAssistantResponse struct {
	util.CommonError
}

// ModifyAssistant 修改直播间小助手
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/studio-management/api_modifyassistant
func (s *LiveBroadcast) ModifyAssistant(in *ModifyAssistantRequest) (out ModifyAssistantResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(modifyAssistantURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "ModifyAssistant")
	return
}

// CreateRoomRequest 创建直播间请求
type CreateRoomRequest struct {
	Name            string `json:"name"`            // 直播间名字，最短3个汉字，最长17个汉字，1个汉字相当于2个字符
	CoverImg        string `json:"coverImg"`        // 背景图，填入mediaID（mediaID获取后，三天内有效）；直播间背景图，图片规则：建议像素1080*1920，大小不超过2M
	StartTime       int64  `json:"startTime"`       // 直播计划开始时间（开播时间需要在当前时间的10分钟后 并且 开始时间不能在 6 个月后）
	EndTime         int64  `json:"endTime"`         // 直播计划结束时间（开播时间和结束时间间隔不得短于30分钟，不得超过24小时）
	AnchorName      string `json:"anchorName"`      // 主播昵称，最短2个汉字，最长15个汉字，1个汉字相当于2个字符
	AnchorWechat    string `json:"anchorWechat"`    // 主播微信号，如果未实名认证，需要先前往 小程序直播 小程序进行实名验证
	SubAnchorWechat string `json:"subAnchorWechat"` // 主播副号微信号，如果未实名认证，需要先前往 小程序直播 小程序进行实名验证
	CreaterWechat   string `json:"createrWechat"`   // 创建者微信号，不传入则此直播间所有成员可见。传入则此房间仅创建者、管理员、超管、直播间主播可见
	ShareImg        string `json:"shareImg"`        // 分享图，填入mediaID（mediaID获取后，三天内有效）；直播间分享图，图片规则：建议像素800*640，大小不超过1M
	FeedsImg        string `json:"feedsImg"`        // 购物直播频道封面图，填入mediaID（mediaID获取后，三天内有效）；购物直播频道封面图，图片规则：建议像素800*800，大小不超过100KB
	IsFeedsPublic   int64  `json:"isFeedsPublic"`   // 是否开启官方收录 【1: 开启，0：关闭】，默认开启收录
	Type            int64  `json:"type"`            // 直播间类型 【1: 推流，0：手机直播】
	CloseLike       int64  `json:"closeLike"`       // 是否关闭点赞 【0：开启，1：关闭】（若关闭，观众端将隐藏点赞按钮，直播开始后不允许开启）
	CloseGoods      int64  `json:"closeGoods"`      // 是否关闭货架 【0：开启，1：关闭】（若关闭，观众端将隐藏商品货架，直播开始后不允许开启）
	CloseComment    int64  `json:"closeComment"`    // 是否关闭评论 【0：开启，1：关闭】（若关闭，观众端将隐藏评论入口，直播开始后不允许开启）
	CloseReplay     int64  `json:"closeReplay"`     // 是否关闭回放 【0：开启，1：关闭】默认关闭回放（直播开始后允许开启）
	CloseShare      int64  `json:"closeShare"`      // 是否关闭分享 【0：开启，1：关闭】默认开启分享（直播开始后不允许修改）
	CloseKf         int64  `json:"closeKf"`         // 是否关闭客服 【0：开启，1：关闭】 默认关闭客服（直播开始后允许开启）
}

// CreateRoomResponse 创建直播间响应
type CreateRoomResponse struct {
	util.CommonError
	RoomID    int64  `json:"roomId"`     // 房间ID
	QrcodeURL string `json:"qrcode_url"` // 小程序直播 小程序码
}

// CreateRoom 创建直播间
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/studio-management/api_createroom
func (s *LiveBroadcast) CreateRoom(in *CreateRoomRequest) (out CreateRoomResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(createRoomURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "CreateRoom")
	return
}

// DeleteSubAnchorRequest 删除主播副号请求
type DeleteSubAnchorRequest struct {
	RoomID int64 `json:"roomId"` // 房间ID
}

// DeleteSubAnchorResponse 删除主播副号响应
type DeleteSubAnchorResponse struct {
	util.CommonError
}

// DeleteSubAnchor 删除主播副号
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/studio-management/api_deletesubanchor
func (s *LiveBroadcast) DeleteSubAnchor(in *DeleteSubAnchorRequest) (out DeleteSubAnchorResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(deleteSubAnchorURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "DeleteSubAnchor")
	return
}

// DeleteRoomRequest 删除直播间请求
type DeleteRoomRequest struct {
	ID int64 `json:"id"` // 房间ID
}

// DeleteRoomResponse 删除直播间响应
type DeleteRoomResponse struct {
	util.CommonError
}

// DeleteRoom 删除直播间
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/studio-management/api_deleteroom
func (s *LiveBroadcast) DeleteRoom(in *DeleteRoomRequest) (out DeleteRoomResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(deleteRoomURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "DeleteRoom")
	return
}

// DeleteRoomGoodsRequest 删除直播间商品请求
type DeleteRoomGoodsRequest struct {
	RoomID  int64 `json:"roomId"`  // 房间ID
	GoodsID int64 `json:"goodsId"` // 商品ID
}

// DeleteRoomGoodsResponse 删除直播间商品响应
type DeleteRoomGoodsResponse struct {
	util.CommonError
}

// DeleteRoomGoods 删除直播间商品
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/studio-management/api_deletedoods
func (s *LiveBroadcast) DeleteRoomGoods(in *DeleteRoomGoodsRequest) (out DeleteRoomGoodsResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(deleteRoomGoodsURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "DeleteRoomGoods")
	return
}

// RemoveAssistantRequest 删除直播间小助手请求
type RemoveAssistantRequest struct {
	RoomID   int64  `json:"roomId"`   // 房间ID
	Username string `json:"username"` // 用户微信号
}

// RemoveAssistantResponse 删除直播间小助手响应
type RemoveAssistantResponse struct {
	util.CommonError
}

// RemoveAssistant 删除直播间小助手
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/studio-management/api_removeassistant
func (s *LiveBroadcast) RemoveAssistant(in *RemoveAssistantRequest) (out RemoveAssistantResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(removeAssistantURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "RemoveAssistant")
	return
}

// UpdateReplayRequest 回放功能管理请求
type UpdateReplayRequest struct {
	RoomID      int64 `json:"roomId"`      // 房间ID
	CloseReplay int64 `json:"closeReplay"` // 是否关闭回放 【0：开启，1：关闭】
}

// UpdateReplayResponse 回放功能管理响应
type UpdateReplayResponse struct {
	util.CommonError
}

// UpdateReplay 回放功能管理
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/studio-management/api_updatereplay
func (s *LiveBroadcast) UpdateReplay(in *UpdateReplayRequest) (out UpdateReplayResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(updateReplayURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "UpdateReplay")
	return
}

// UpdateFeedPublicRequest 官方收录管理请求
type UpdateFeedPublicRequest struct {
	RoomID        int64 `json:"roomId"`        // 房间ID
	IsFeedsPublic int64 `json:"isFeedsPublic"` // 是否开启官方收录 【1: 开启，0：关闭】
}

// UpdateFeedPublicResponse 官方收录管理响应
type UpdateFeedPublicResponse struct {
	util.CommonError
}

// UpdateFeedPublic 官方收录管理
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/studio-management/api_updatefeedpublic
func (s *LiveBroadcast) UpdateFeedPublic(in *UpdateFeedPublicRequest) (out UpdateFeedPublicResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(updateFeedPublicURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "UpdateFeedPublic")
	return
}

// UpdateKFRequest 客服功能管理请求
type UpdateKFRequest struct {
	RoomID  int64 `json:"roomId"`  // 房间ID
	CloseKf int64 `json:"closeKf"` // 是否关闭客服 【0：开启，1：关闭】
}

// UpdateKFResponse 客服功能管理响应
type UpdateKFResponse struct {
	util.CommonError
}

// UpdateKF 客服功能管理
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/studio-management/api_updatekf
func (s *LiveBroadcast) UpdateKF(in *UpdateKFRequest) (out UpdateKFResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(updateKFURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "UpdateKF")
	return
}

// ImportGoodsRequest 导入商品请求
type ImportGoodsRequest struct {
	IDs    []int64 `json:"ids"`    // 数组列表，可传入多个，里面填写 商品 ID
	RoomID int64   `json:"roomId"` // 房间ID
}

// ImportGoodsResponse 导入商品响应
type ImportGoodsResponse struct {
	util.CommonError
}

// ImportGoods 导入商品
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/studio-management/api_importgoods
func (s *LiveBroadcast) ImportGoods(in *ImportGoodsRequest) (out ImportGoodsResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(importGoodsURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "ImportGoods")
	return
}

// PushGoodsRequest 推送商品请求
type PushGoodsRequest struct {
	RoomID  int64 `json:"roomId"`  // 房间ID
	GoodsID int64 `json:"goodsId"` // 商品ID
}

// PushGoodsResponse 推送商品响应
type PushGoodsResponse struct {
	util.CommonError
}

// PushGoods 推送商品
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/studio-management/api_pushgoods
func (s *LiveBroadcast) PushGoods(in *PushGoodsRequest) (out PushGoodsResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(pushGoodsURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "PushGoods")
	return
}

// GetAssistantListRequest 查询直播间小助手请求
type GetAssistantListRequest struct {
	RoomID int64 `json:"roomId"` // 房间ID
}

// GetAssistantListResponse 查询直播间小助手响应
type GetAssistantListResponse struct {
	util.CommonError
	List     []Assistant `json:"list"`     // 小助手列表
	Count    int64       `json:"count"`    // 小助手个数
	MaxCount int64       `json:"maxCount"` // 小助手最大个数
}

// GetAssistantList 查询直播间小助手
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/studio-management/api_getassistantlist
func (s *LiveBroadcast) GetAssistantList(in *GetAssistantListRequest) (out GetAssistantListResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	params := map[string]interface{}{
		"access_token": accessToken,
		"roomId":       in.RoomID,
	}
	query := util.Query(params)

	response, err := util.HTTPGet(fmt.Sprintf(getAssistantListURL, query))
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetAssistantList")
	return
}

// AddSubAnchorRequest 添加主播副号请求
type AddSubAnchorRequest struct {
	RoomID   int64  `json:"roomId"`   // 房间ID
	Username string `json:"username"` // 用户微信号
}

// AddSubAnchorResponse 添加主播副号响应
type AddSubAnchorResponse struct {
	util.CommonError
}

// AddSubAnchor 添加主播副号
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/studio-management/api_addsubanchor
func (s *LiveBroadcast) AddSubAnchor(in *AddSubAnchorRequest) (out AddSubAnchorResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(addSubAnchorURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "AddSubAnchor")
	return
}

// AddAssistantRequest 添加管理直播间小助手请求
type AddAssistantRequest struct {
	RoomID int64           `json:"roomId"` // 房间ID
	Users  []AssistantUser `json:"users"`  // 用户数组
}

// AddAssistantResponse 添加管理直播间小助手响应
type AddAssistantResponse struct {
	util.CommonError
}

// AddAssistant 添加管理直播间小助手
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/studio-management/api_addveassistant
func (s *LiveBroadcast) AddAssistant(in *AddAssistantRequest) (out AddAssistantResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(addAssistantURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "AddAssistant")
	return
}

// GetDefaultGoodsKeyResponse 直播挂件获取全局KEY响应
type GetDefaultGoodsKeyResponse struct {
	util.CommonError
	VendorGoodsKey []string `json:"vendorGoodsKey"` // 全局KEY
}

// GetDefaultGoodsKey 直播挂件获取全局KEY
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/studio-management/api_getdefault_goodskey
func (s *LiveBroadcast) GetDefaultGoodsKey() (out GetDefaultGoodsKeyResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	params := map[string]interface{}{
		"access_token": accessToken,
	}
	query := util.Query(params)

	response, err := util.HTTPGet(fmt.Sprintf(getDefaultGoodsKeyURL, query))
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetDefaultGoodsKey")
	return
}

// SetDefaultGoodsKeyRequest 直播挂件设置全局KEY请求
type SetDefaultGoodsKeyRequest struct {
	GoodsKey []string `json:"goodsKey"` // 商品KEY
}

// SetDefaultGoodsKeyResponse 直播挂件设置全局KEY响应
type SetDefaultGoodsKeyResponse struct {
	util.CommonError
}

// SetDefaultGoodsKey 直播挂件设置全局KEY
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/studio-management/api_setdefault_goodskey
func (s *LiveBroadcast) SetDefaultGoodsKey(in *SetDefaultGoodsKeyRequest) (out SetDefaultGoodsKeyResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(setDefaultGoodsKeyURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "SetDefaultGoodsKey")
	return
}

// SortGoodsRequest 直播间商品排序请求
type SortGoodsRequest struct {
	RoomID int64         `json:"roomId"` // 房间ID
	Goods  []GoodsIDItem `json:"goods"`  // 商品ID列表
}

// SortGoodsResponse 直播间商品排序响应
type SortGoodsResponse struct {
	util.CommonError
}

// SortGoods 直播间商品排序
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/studio-management/api_sortgoods
func (s *LiveBroadcast) SortGoods(in *SortGoodsRequest) (out SortGoodsResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(sortGoodsURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "SortGoods")
	return
}

// UpdateCommentRequest 禁言管理请求
type UpdateCommentRequest struct {
	ID         int64 `json:"id"`         // 房间ID
	RoomID     int64 `json:"roomId"`     // 房间ID（官方参数表写 id、请求示例写 roomId，两个都传）
	BanComment int64 `json:"banComment"` // 1-禁言，0-取消禁言
}

// UpdateCommentResponse 禁言管理响应
type UpdateCommentResponse struct {
	util.CommonError
}

// UpdateComment 禁言管理
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/studio-management/api_updatecomment
func (s *LiveBroadcast) UpdateComment(in *UpdateCommentRequest) (out UpdateCommentResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	// 官方参数表写 id、请求示例写 roomId，两个字段都带上
	if in.RoomID == 0 {
		in.RoomID = in.ID
	}
	if in.ID == 0 {
		in.ID = in.RoomID
	}
	uri := fmt.Sprintf(updateCommentURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "UpdateComment")
	return
}

// EditRoomRequest 编辑直播间请求
type EditRoomRequest struct {
	ID            int64  `json:"id"`            // 直播间id
	Name          string `json:"name"`          // 直播间名字，最短3个汉字，最长17个汉字，1个汉字相当于2个字符
	CoverImg      string `json:"coverImg"`      // 背景图，填入mediaID（mediaID获取后，三天内有效）；直播间背景图，图片规则：建议像素1080*1920，大小不超过2M
	StartTime     int64  `json:"startTime"`     // 直播计划开始时间（开播时间需要在当前时间的10分钟后 并且 开始时间不能在 6 个月后）
	EndTime       int64  `json:"endTime"`       // 直播计划结束时间（开播时间和结束时间间隔不得短于30分钟，不得超过24小时）
	AnchorName    string `json:"anchorName"`    // 主播昵称，最短2个汉字，最长15个汉字，1个汉字相当于2个字符
	AnchorWechat  string `json:"anchorWechat"`  // 主播微信号，如果未实名认证，需要先前往 小程序直播 小程序进行实名验证
	ShareImg      string `json:"shareImg"`      // 分享图，填入mediaID（mediaID获取后，三天内有效）；直播间分享图，图片规则：建议像素800*640，大小不超过1M
	FeedsImg      string `json:"feedsImg"`      // 购物直播频道封面图，填入mediaID（mediaID获取后，三天内有效）；购物直播频道封面图，图片规则：建议像素800*800，大小不超过100KB
	IsFeedsPublic int64  `json:"isFeedsPublic"` // 是否开启官方收录 【1: 开启，0：关闭】，默认开启收录
	CloseLike     int64  `json:"closeLike"`     // 是否关闭点赞 【0：开启，1：关闭】（若关闭，观众端将隐藏点赞按钮，直播开始后不允许开启）
	CloseGoods    int64  `json:"closeGoods"`    // 是否关闭货架 【0：开启，1：关闭】（若关闭，观众端将隐藏商品货架，直播开始后不允许开启）
	CloseComment  int64  `json:"closeComment"`  // 是否关闭评论 【0：开启，1：关闭】（若关闭，观众端将隐藏评论入口，直播开始后不允许开启）
	CloseReplay   int64  `json:"closeReplay"`   // 是否关闭回放 【0：开启，1：关闭】默认关闭回放（直播开始后允许开启）
	CloseShare    int64  `json:"closeShare"`    // 是否关闭分享 【0：开启，1：关闭】默认开启分享（直播开始后不允许修改）
	CloseKf       int64  `json:"closeKf"`       // 是否关闭客服 【0：开启，1：关闭】 默认关闭客服（直播开始后允许开启）
}

// EditRoomResponse 编辑直播间响应
type EditRoomResponse struct {
	util.CommonError
}

// EditRoom 编辑直播间
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/studio-management/api_editroom
func (s *LiveBroadcast) EditRoom(in *EditRoomRequest) (out EditRoomResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(editRoomURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "EditRoom")
	return
}

// GetSubAnchorRequest 获取主播副号请求
type GetSubAnchorRequest struct {
	RoomID int64 `json:"roomId"` // 房间ID
}

// GetSubAnchorResponse 获取主播副号响应
type GetSubAnchorResponse struct {
	util.CommonError
	Username string `json:"username"` // 主播微信号
}

// GetSubAnchor 获取主播副号
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/studio-management/api_getsubanchor
func (s *LiveBroadcast) GetSubAnchor(in *GetSubAnchorRequest) (out GetSubAnchorResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	params := map[string]interface{}{
		"access_token": accessToken,
		"roomId":       in.RoomID,
	}
	query := util.Query(params)

	response, err := util.HTTPGet(fmt.Sprintf(getSubAnchorURL, query))
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetSubAnchor")
	return
}

// GetSharedCodeRequest 获取直播间分享二维码请求
type GetSharedCodeRequest struct {
	RoomID int64  `json:"roomId"` // 房间ID
	Params string `json:"params"` // 自定义参数
}

// GetSharedCodeResponse 获取直播间分享二维码响应
type GetSharedCodeResponse struct {
	util.CommonError
	CDNURL    string `json:"cdnUrl"`    // 分享二维码cdn url
	PagePath  string `json:"pagePath"`  // 分享路径
	PosterURL string `json:"posterUrl"` // 分享海报 url
}

// GetSharedCode 获取直播间分享二维码
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/studio-management/api_getsharedcode
func (s *LiveBroadcast) GetSharedCode(in *GetSharedCodeRequest) (out GetSharedCodeResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	params := map[string]interface{}{
		"access_token": accessToken,
		"roomId":       in.RoomID,
	}
	if in.Params != "" {
		params["params"] = url.QueryEscape(in.Params)
	}
	query := util.Query(params)

	response, err := util.HTTPGet(fmt.Sprintf(getSharedCodeURL, query))
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetSharedCode")
	return
}

// GetLiveInfoRequest 获取直播间列表和回放请求
type GetLiveInfoRequest struct {
	Start  int64  `json:"start"`   // 起始拉取视频，0表示从第一个视频片段开始拉取
	Limit  int64  `json:"limit"`   // 每次拉取的数量，建议100以内
	Action string `json:"action"`  // 只能填"get_replay"，表示获取回放
	RoomID int64  `json:"room_id"` // 当action有值时该字段必填，直播间ID
}

// GetLiveInfoResponse 获取直播间列表和回放响应
type GetLiveInfoResponse struct {
	util.CommonError
	RoomInfo   []RoomInfo   `json:"room_info"`   // action="get_replay"不返回
	Total      int64        `json:"total"`       // 拉取房间总数
	LiveReplay []LiveReplay `json:"live_replay"` // action="get_replay"才返回
}

// GetLiveInfo 获取直播间列表和回放
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/studio-management/api_getliveinfo
func (s *LiveBroadcast) GetLiveInfo(in *GetLiveInfoRequest) (out GetLiveInfoResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(getLiveInfoURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetLiveInfo")
	return
}

// GetPushUrlRequest 获取直播间推流地址请求
type GetPushUrlRequest struct {
	RoomID int64 `json:"roomId"` // 房间ID
}

// GetPushUrlResponse 获取直播间推流地址响应
type GetPushUrlResponse struct {
	util.CommonError
	PushAddr string `json:"pushAddr"` // 推流地址
}

// GetPushUrl 获取直播间推流地址
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/studio-management/api_getpushurl
func (s *LiveBroadcast) GetPushUrl(in *GetPushUrlRequest) (out GetPushUrlResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	params := map[string]interface{}{
		"access_token": accessToken,
		"roomId":       in.RoomID,
	}
	query := util.Query(params)

	response, err := util.HTTPGet(fmt.Sprintf(getPushUrlURL, query))
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetPushUrl")
	return
}
