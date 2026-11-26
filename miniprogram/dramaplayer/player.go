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

package dramaplayer

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// developerSetFlushDramaURL 刷剧剧目设置
	developerSetFlushDramaURL = "https://api.weixin.qq.com/wxadrama/developersetflushdrama?access_token=%s"
	// batchProcessDramaPromotionURL 批处理短剧合作推广计划
	batchProcessDramaPromotionURL = "https://api.weixin.qq.com/wxadrama/batchprocessdramapromotion?access_token=%s"
	// developerSetRecmdDramaURL 推荐剧目
	developerSetRecmdDramaURL = "https://api.weixin.qq.com/wxadrama/developersetrecmddrama?access_token=%s"
	// setPlayerDramaRecmdSwitchURL 播放原始视频-推荐位控制接口
	setPlayerDramaRecmdSwitchURL = "https://api.weixin.qq.com/wxadrama/setplayerdramarecmdswitch?access_token=%s"
	// developerPublishDramaURL 短剧上架
	developerPublishDramaURL = "https://api.weixin.qq.com/wxadrama/developerpublishdrama?access_token=%s"
	// developerGetPublishedDramaURL 获取已上架短剧
	developerGetPublishedDramaURL = "https://api.weixin.qq.com/wxadrama/developergetpublisheddrama?access_token=%s"
	// developerGetIaaDramaURL 获取短剧变现类型
	developerGetIaaDramaURL = "https://api.weixin.qq.com/wxadrama/developergetiaadrama?access_token=%s"
	// getFinderEventURL 获取短剧合作推广活动
	getFinderEventURL = "https://api.weixin.qq.com/wxadrama/getfinderevent?access_token=%s"
	// developerSetIaaDramaURL 设置短剧变现类型
	developerSetIaaDramaURL = "https://api.weixin.qq.com/wxadrama/developersetiaadrama?access_token=%s"
)

// FlushDramaItem 刷剧剧目
type FlushDramaItem struct {
	SrcAppID  string `json:"src_appid"`  // 提审方小程序的appid
	DramaID   string `json:"drama_id"`   // 提审的剧目id
	DramaName string `json:"drama_name"` // 短剧名称
}

// PromotionDramaItem 合作推广短剧
type PromotionDramaItem struct {
	SrcAppID string `json:"src_appid"` // 提审方小程序的appid
	DramaID  string `json:"drama_id"`  // 提审的剧目id
}

// DramaPromotionStatus 短剧合作推广状态
type DramaPromotionStatus struct {
	SrcAppID string `json:"src_appid"` // 提审方小程序的appid
	DramaID  string `json:"drama_id"`  // 提审的剧目id
	Status   int64  `json:"status"`    // 短剧状态。 0-未加入计划；1-审核中；2-审核通过；3-审核拒绝；4-退出计划中；5-已被下架
}

// RecmdDramaItem 推荐剧目
type RecmdDramaItem struct {
	SrcAppID string `json:"src_appid"` // 提审方小程序的appid
	DramaID  string `json:"drama_id"`  // 提审的剧目id
}

// PublishDramaItem 上架短剧
type PublishDramaItem struct {
	SrcAppID    string `json:"src_appid"`    // 提审方小程序的appid
	DramaID     string `json:"drama_id"`     // 提审的剧目id
	DramaName   string `json:"drama_name"`   // 短剧名称
	PublishTime int64  `json:"publish_time"` // 上架时间戳（秒级别）
}

// PublishedDrama 已上架短剧
type PublishedDrama struct {
	SrcAppID    string `json:"src_appid"`    // 提审方小程序的appid
	DramaID     string `json:"drama_id"`     // 提审的剧目id
	PublishTime int64  `json:"publish_time"` // 上架时间戳（秒级别）
}

// IaaDramaItem 短剧变现类型查询项
type IaaDramaItem struct {
	SrcAppID string `json:"src_appid"` // 提审方小程序的appid
	DramaID  string `json:"drama_id"`  // 提审的剧目id
}

// IaaDramaInfo 短剧变现类型
type IaaDramaInfo struct {
	SrcAppID string `json:"src_appid"` // 提审方小程序的appid
	DramaID  string `json:"drama_id"`  // 提审的剧目id
	IaaType  int64  `json:"iaa_type"`  // 短剧变现类型。 1-纯IAA；2-纯IAP；3-IAAP混合变现
	VipFlag  int64  `json:"vip_flag"`  // 是否存在会员功能。 1-有；2-无
}

// SetIaaDramaItem 设置短剧变现类型项
type SetIaaDramaItem struct {
	SrcAppID string `json:"src_appid"` // 提审方小程序的appid
	DramaID  string `json:"drama_id"`  // 提审的剧目id
	IaaType  int64  `json:"iaa_type"`  // 短剧变现类型。 1-纯IAA；2-纯IAP；3-IAAP混合变现
	VipFlag  int64  `json:"vip_flag"`  // 是否存在会员功能。 1-有；2-无
}

// FinderEvent 短剧合作推广活动
type FinderEvent struct {
	EncryptedEventID string `json:"encrypted_event_id"` // 活动id
	EventName        string `json:"event_name"`         // 活动名
	EventURL         string `json:"event_url"`          // 活动链接
	SrcAppID         string `json:"src_appid"`          // 提审小程序id
	DramaID          string `json:"drama_id"`           // 提审剧目id
}

// DeveloperSetFlushDramaRequest 刷剧剧目设置请求
type DeveloperSetFlushDramaRequest struct {
	List []FlushDramaItem `json:"list"` // 剧目列表
}

// DeveloperSetFlushDramaResponse 刷剧剧目设置响应
type DeveloperSetFlushDramaResponse struct {
	util.CommonError
}

// DeveloperSetFlushDrama 刷剧剧目设置
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/dramaOthersAPI/api_developersetflushdrama
func (s *DramaPlayer) DeveloperSetFlushDrama(in *DeveloperSetFlushDramaRequest) (out DeveloperSetFlushDramaResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(developerSetFlushDramaURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "DeveloperSetFlushDrama")
	return
}

// BatchProcessDramaPromotionRequest 批处理短剧合作推广计划请求
type BatchProcessDramaPromotionRequest struct {
	ActionType int64                `json:"action_type"` // 操作类型。1-加入计划；2-查询计划；3-退出计划
	List       []PromotionDramaItem `json:"list"`        // 短剧数组
}

// BatchProcessDramaPromotionResponse 批处理短剧合作推广计划响应
type BatchProcessDramaPromotionResponse struct {
	util.CommonError
	List []DramaPromotionStatus `json:"list"` // 短剧数组
}

// BatchProcessDramaPromotion 批处理短剧合作推广计划
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/dramaOthersAPI/api_batchprocessdramapromotion
func (s *DramaPlayer) BatchProcessDramaPromotion(in *BatchProcessDramaPromotionRequest) (out BatchProcessDramaPromotionResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(batchProcessDramaPromotionURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "BatchProcessDramaPromotion")
	return
}

// DeveloperSetRecmdDramaRequest 推荐剧目请求
type DeveloperSetRecmdDramaRequest struct {
	EntryType int64            `json:"entry_type"` // 1-剧结束 2-选集最右侧推荐 3-剧集profile页相关推荐
	SrcAppID  string           `json:"src_appid"`  // 提审方小程序的appid
	DramaID   string           `json:"drama_id"`   // 提审的剧目id
	List      []RecmdDramaItem `json:"list"`       // list
}

// DeveloperSetRecmdDramaResponse 推荐剧目响应
type DeveloperSetRecmdDramaResponse struct {
	util.CommonError
}

// DeveloperSetRecmdDrama 推荐剧目
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/dramaOthersAPI/api_developersetrecmddrama
func (s *DramaPlayer) DeveloperSetRecmdDrama(in *DeveloperSetRecmdDramaRequest) (out DeveloperSetRecmdDramaResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(developerSetRecmdDramaURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "DeveloperSetRecmdDrama")
	return
}

// SetPlayerDramaRecmdSwitchRequest 播放原始视频-推荐位控制接口请求
type SetPlayerDramaRecmdSwitchRequest struct {
	EntryType    int64 `json:"entry_type"`    // 入口类型
	SwitchStatus bool  `json:"switch_status"` // 是否开启推荐位
}

// SetPlayerDramaRecmdSwitchResponse 播放原始视频-推荐位控制接口响应
type SetPlayerDramaRecmdSwitchResponse struct {
	util.CommonError
}

// SetPlayerDramaRecmdSwitch 播放原始视频-推荐位控制接口
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/dramaOthersAPI/api_setplayerdramarecmdswitch
func (s *DramaPlayer) SetPlayerDramaRecmdSwitch(in *SetPlayerDramaRecmdSwitchRequest) (out SetPlayerDramaRecmdSwitchResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(setPlayerDramaRecmdSwitchURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "SetPlayerDramaRecmdSwitch")
	return
}

// DeveloperPublishDramaRequest 短剧上架请求
type DeveloperPublishDramaRequest struct {
	List []PublishDramaItem `json:"list"` // 剧目列表
}

// DeveloperPublishDramaResponse 短剧上架响应
type DeveloperPublishDramaResponse struct {
	util.CommonError
}

// DeveloperPublishDrama 短剧上架
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/dramaOthersAPI/api_developerpublishdrama
func (s *DramaPlayer) DeveloperPublishDrama(in *DeveloperPublishDramaRequest) (out DeveloperPublishDramaResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(developerPublishDramaURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "DeveloperPublishDrama")
	return
}

// DeveloperGetPublishedDramaRequest 获取已上架短剧请求
type DeveloperGetPublishedDramaRequest struct{}

// DeveloperGetPublishedDramaResponse 获取已上架短剧响应
type DeveloperGetPublishedDramaResponse struct {
	util.CommonError
	List []PublishedDrama `json:"list"` // 已上架短剧列表
}

// DeveloperGetPublishedDrama 获取已上架短剧
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/dramaOthersAPI/api_developergetpublisheddrama
func (s *DramaPlayer) DeveloperGetPublishedDrama(in *DeveloperGetPublishedDramaRequest) (out DeveloperGetPublishedDramaResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(developerGetPublishedDramaURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "DeveloperGetPublishedDrama")
	return
}

// DeveloperGetIaaDramaRequest 获取短剧变现类型请求
type DeveloperGetIaaDramaRequest struct {
	List []IaaDramaItem `json:"list"` // 剧目列表
}

// DeveloperGetIaaDramaResponse 获取短剧变现类型响应
type DeveloperGetIaaDramaResponse struct {
	util.CommonError
	List []IaaDramaInfo `json:"list"` // 短剧变现类型列表
}

// DeveloperGetIaaDrama 获取短剧变现类型
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/dramaOthersAPI/api_developergetiaadrama
func (s *DramaPlayer) DeveloperGetIaaDrama(in *DeveloperGetIaaDramaRequest) (out DeveloperGetIaaDramaResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(developerGetIaaDramaURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "DeveloperGetIaaDrama")
	return
}

// GetFinderEventRequest 获取短剧合作推广活动请求
type GetFinderEventRequest struct {
	EventIDList []string `json:"event_id_list"` // 筛选短剧合作推广活动
}

// GetFinderEventResponse 获取短剧合作推广活动响应
type GetFinderEventResponse struct {
	util.CommonError
	FinderEventList []FinderEvent `json:"finder_event_list"` // 推广计划数组
}

// GetFinderEvent 获取短剧合作推广活动
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/dramaOthersAPI/api_getfinderevent
func (s *DramaPlayer) GetFinderEvent(in *GetFinderEventRequest) (out GetFinderEventResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(getFinderEventURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetFinderEvent")
	return
}

// DeveloperSetIaaDramaRequest 设置短剧变现类型请求
type DeveloperSetIaaDramaRequest struct {
	List []SetIaaDramaItem `json:"list"` // 剧目列表
}

// DeveloperSetIaaDramaResponse 设置短剧变现类型响应
type DeveloperSetIaaDramaResponse struct {
	util.CommonError
}

// DeveloperSetIaaDrama 设置短剧变现类型
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/dramaOthersAPI/api_developersetiaadrama
func (s *DramaPlayer) DeveloperSetIaaDrama(in *DeveloperSetIaaDramaRequest) (out DeveloperSetIaaDramaResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(developerSetIaaDramaURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "DeveloperSetIaaDrama")
	return
}
