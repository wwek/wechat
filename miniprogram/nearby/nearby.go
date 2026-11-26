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

package nearby

import (
	"fmt"

	"github.com/silenceper/wechat/v2/miniprogram/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	// addNearbyPoiURL 添加地点
	addNearbyPoiURL = "https://api.weixin.qq.com/wxa/addnearbypoi?access_token=%s"
	// deleteNearbyPoiURL 删除地点
	deleteNearbyPoiURL = "https://api.weixin.qq.com/wxa/delnearbypoi?access_token=%s"
	// getNearbyPoiListURL 查看地点
	getNearbyPoiListURL = "https://api.weixin.qq.com/wxa/getnearbypoilist?access_token=%s"
	// setShowStatusURL 设置展示状态
	setShowStatusURL = "https://api.weixin.qq.com/wxa/setnearbypoishowstatus?access_token=%s"
)

// Nearby 附近小程序
type Nearby struct {
	*context.Context
}

// NewNearby 实例化
func NewNearby(ctx *context.Context) *Nearby {
	return &Nearby{ctx}
}

// AddNearbyPoiRequest 添加地点请求
type AddNearbyPoiRequest struct {
	KfInfo            string `json:"kf_info"`            // 客服信息，选填，可自定义服务头像与昵称，内容是一个 json
	PicList           string `json:"pic_list"`           // 门店图片，最多9张，最少1张，内容是一个 json
	ServiceInfos      string `json:"service_infos"`      // 服务标签列表，内容是一个 json
	StoreName         string `json:"store_name"`         // 门店名字，需按照所选地理位置自动拉取腾讯地图门店名称，不可修改
	ContractPhone     string `json:"contract_phone"`     // 门店电话
	Hour              string `json:"hour"`               // 营业时间，格式11:11-12:12
	CompanyName       string `json:"company_name"`       // 主体名字
	Credential        string `json:"credential"`         // 资质号，15位营业执照注册号或9位组织机构代码
	Address           string `json:"address"`            // 地址
	QualificationList string `json:"qualification_list"` // 证明材料，company_name 和该小程序主体不一致时需要填写
	IsCommNearby      string `json:"is_comm_nearby"`     // 必填，写死为"1"
	PoiID             string `json:"poi_id"`             // 附近地点 ID，创建新门店时为空，更新门店时填写对应门店的 poi_id，选填
	MapPoiID          string `json:"map_poi_id"`         // 对应《在腾讯地图中搜索门店》中的 sosomap_poi_uid 字段
}

// AddNearbyPoiData 添加地点返回的业务数据
type AddNearbyPoiData struct {
	AuditID           int64  `json:"audit_id"`           // 审核单 ID
	PoiID             int64  `json:"poi_id"`             // 附近地点 ID
	RelatedCredential string `json:"related_credential"` // 经营资质证件号
}

// AddNearbyPoiResponse 添加地点响应
type AddNearbyPoiResponse struct {
	util.CommonError
	Data AddNearbyPoiData `json:"data"` // 附近小程序的数据
}

// AddNearbyPoi 添加地点
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/nearby-poi/api_addnearbypoi
func (s *Nearby) AddNearbyPoi(in *AddNearbyPoiRequest) (out AddNearbyPoiResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(addNearbyPoiURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "AddNearbyPoi")
	return
}

// DeleteNearbyPoiRequest 删除地点请求
type DeleteNearbyPoiRequest struct {
	PoiID string `json:"poi_id"` // 附近地点 ID
}

// DeleteNearbyPoiResponse 删除地点响应
type DeleteNearbyPoiResponse struct {
	util.CommonError
}

// DeleteNearbyPoi 删除地点
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/nearby-poi/api_deletenearbypoi
func (s *Nearby) DeleteNearbyPoi(in *DeleteNearbyPoiRequest) (out DeleteNearbyPoiResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(deleteNearbyPoiURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "DeleteNearbyPoi")
	return
}

// GetNearbyPoiListRequest 查看地点请求
type GetNearbyPoiListRequest struct {
	Page     int64 `json:"page"`      // 起始页 id（从1开始计数）
	PageRows int64 `json:"page_rows"` // 每页展示个数（最多1000个）
}

// GetNearbyPoiListData 查看地点返回的业务数据
type GetNearbyPoiListData struct {
	MaxApplyNum  int64  `json:"max_apply_num"`  // 最大可添加地点个数
	LeftApplyNum int64  `json:"left_apply_num"` // 剩余可添加地点个数
	Data         string `json:"data"`           // 地址列表的 JSON 格式字符串
}

// GetNearbyPoiListResponse 查看地点响应
type GetNearbyPoiListResponse struct {
	util.CommonError
	Data GetNearbyPoiListData `json:"data"` // 附近小程序的数据
}

// GetNearbyPoiList 查看地点
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/nearby-poi/api_getnearbypoilist
func (s *Nearby) GetNearbyPoiList(in *GetNearbyPoiListRequest) (out GetNearbyPoiListResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(getNearbyPoiListURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetNearbyPoiList")
	return
}

// SetShowStatusRequest 设置展示状态请求
type SetShowStatusRequest struct {
	PoiID  string `json:"poi_id"` // 附近地点 ID
	Status int64  `json:"status"` // 是否展示，0表示不展示，1表示展示
}

// SetShowStatusResponse 设置展示状态响应
type SetShowStatusResponse struct {
	util.CommonError
}

// SetShowStatus 设置展示状态
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/nearby-poi/api_setshowstatus
func (s *Nearby) SetShowStatus(in *SetShowStatusRequest) (out SetShowStatusResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(setShowStatusURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "SetShowStatus")
	return
}
