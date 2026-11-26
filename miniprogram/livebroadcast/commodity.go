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
	// deleteGoodsInfoURL 删除商品
	deleteGoodsInfoURL = "https://api.weixin.qq.com/wxaapi/broadcast/goods/delete?access_token=%s"
	// resetAuditURL 撤回商品审核
	resetAuditURL = "https://api.weixin.qq.com/wxaapi/broadcast/goods/resetaudit?access_token=%s"
	// updateGoodsinfoURL 更新商品
	updateGoodsinfoURL = "https://api.weixin.qq.com/wxaapi/broadcast/goods/update?access_token=%s"
	// addGoodsURL 添加并提审商品
	addGoodsURL = "https://api.weixin.qq.com/wxaapi/broadcast/goods/add?access_token=%s"
	// getGoodsInfoURL 获取商品列表
	getGoodsInfoURL = "https://api.weixin.qq.com/wxaapi/broadcast/goods/getapproved?%s"
	// getGoodsAuditInfoURL 获取商品的信息与审核状态
	getGoodsAuditInfoURL = "https://api.weixin.qq.com/wxa/business/getgoodswarehouse?access_token=%s"
	// resubmitAuditURL 重新提交商品审核
	resubmitAuditURL = "https://api.weixin.qq.com/wxaapi/broadcast/goods/audit?access_token=%s"
)

// GoodsInfo 商品信息
type GoodsInfo struct {
	CoverImgUrl     string  `json:"coverImgUrl"`     // 填入mediaID（mediaID获取后，三天内有效）；图片mediaID的获取，请参考 新增临时素材 ；图片规则：图片尺寸最大300像素*300像素
	Name            string  `json:"name"`            // 商品名称，最长14个汉字，1个汉字相当于2个字符
	PriceType       int64   `json:"priceType"`       // 价格类型，1：一口价（只需要传入price，price2不传） 2：价格区间（price字段为左边界，price2字段为右边界，price和price2必传） 3：显示折扣价（price字段为原价，price2字段为现价，price和price2必传）
	Price           float64 `json:"price"`           // 数字，最多保留两位小数，单位元
	Price2          float64 `json:"price2"`          // 数字，最多保留两位小数，单位元
	URL             string  `json:"url"`             // 商品详情页的小程序路径，路径参数存在 url 的，该参数的值需要进行 encode 处理再填入
	ThirdPartyAppid string  `json:"thirdPartyAppid"` // 当商品为第三方小程序的商品则填写为对应第三方小程序的appid，自身小程序商品则填空
	GoodsID         int64   `json:"goodsId"`         // 商品ID（更新商品时必填）
}

// Goods 商品信息列表项
type Goods struct {
	GoodsID         string  `json:"goodsId"`         // 商品ID
	Name            string  `json:"name"`            // 商品名称，最长14个汉字，1个汉字相当于2个字符
	CoverImgUrl     string  `json:"coverImgUrl"`     // 填入mediaID（mediaID获取后，三天内有效）；图片mediaID的获取，请参考 新增临时素材 ；图片规则：图片尺寸最大300像素*300像素
	URL             string  `json:"url"`             // 商品详情页的小程序路径，路径参数存在 url 的，该参数的值需要进行 encode 处理再填入
	PriceType       int64   `json:"priceType"`       // 价格类型，1：一口价（只需要传入price，price2不传） 2：价格区间（price字段为左边界，price2字段为右边界，price和price2必传） 3：显示折扣价（price字段为原价，price2字段为现价，price和price2必传）
	Price           float64 `json:"price"`           // 数字，最多保留两位小数，单位元
	Price2          float64 `json:"price2"`          // 数字，最多保留两位小数，单位元
	ThirdPartyTag   int64   `json:"third_party_tag"` // 1、2：表示是为 API 添加商品，否则是直播控制台添加的商品
	ThirdPartyAppid string  `json:"thirdPartyAppid"` // 当商品为第三方小程序的商品则填写为对应第三方小程序的appid，自身小程序商品则填空
}

// GoodsAuditInfo 商品审核信息列表项
type GoodsAuditInfo struct {
	GoodsID         string  `json:"goods_id"`        // 商品ID
	Name            string  `json:"name"`            // 商品名称，最长14个汉字，1个汉字相当于2个字符
	CoverImgURL     string  `json:"cover_img_url"`   // 填入mediaID（mediaID获取后，三天内有效）；图片mediaID的获取，请参考 新增临时素材 ；图片规则：图片尺寸最大300像素*300像素
	URL             string  `json:"url"`             // 商品详情页的小程序路径，路径参数存在 url 的，该参数的值需要进行 encode 处理再填入
	PriceType       int64   `json:"priceType"`       // 价格类型，1：一口价（只需要传入price，price2不传） 2：价格区间（price字段为左边界，price2字段为右边界，price和price2必传） 3：显示折扣价（price字段为原价，price2字段为现价，price和price2必传）
	Price           float64 `json:"price"`           // 数字，最多保留两位小数，单位元
	Price2          float64 `json:"price2"`          // 数字，最多保留两位小数，单位元
	AuditStatus     string  `json:"audit_status"`    // 审核状态（0：未审核，1：审核中，2:审核通过，3审核失败）
	ThirdPartyTag   int64   `json:"third_party_tag"` // 1、2：表示是为 API 添加商品，否则是直播控制台添加的商品
	ThirdPartyAppid string  `json:"thirdPartyAppid"` // 当商品为第三方小程序的商品则填写为对应第三方小程序的appid，自身小程序商品则填空
}

// DeleteGoodsInfoRequest 删除商品请求
type DeleteGoodsInfoRequest struct {
	GoodsID int64 `json:"goodsId"` // 商品ID
}

// DeleteGoodsInfoResponse 删除商品响应
type DeleteGoodsInfoResponse struct {
	util.CommonError
}

// DeleteGoodsInfo 删除商品
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/commodity-management/api_deletegoodsinfo
func (s *LiveBroadcast) DeleteGoodsInfo(in *DeleteGoodsInfoRequest) (out DeleteGoodsInfoResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(deleteGoodsInfoURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "DeleteGoodsInfo")
	return
}

// ResetAuditRequest 撤回商品审核请求
type ResetAuditRequest struct {
	GoodsID int64 `json:"goodsId"` // 商品ID
	AuditID int64 `json:"auditId"` // 审核单ID
}

// ResetAuditResponse 撤回商品审核响应
type ResetAuditResponse struct {
	util.CommonError
}

// ResetAudit 撤回商品审核
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/commodity-management/api_resetaudit
func (s *LiveBroadcast) ResetAudit(in *ResetAuditRequest) (out ResetAuditResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(resetAuditURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "ResetAudit")
	return
}

// UpdateGoodsinfoRequest 更新商品请求
type UpdateGoodsinfoRequest struct {
	GoodsInfo GoodsInfo `json:"goodsInfo"` // 商品信息
}

// UpdateGoodsinfoResponse 更新商品响应
type UpdateGoodsinfoResponse struct {
	util.CommonError
}

// UpdateGoodsinfo 更新商品
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/commodity-management/api_updategoodsinfo
func (s *LiveBroadcast) UpdateGoodsinfo(in *UpdateGoodsinfoRequest) (out UpdateGoodsinfoResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(updateGoodsinfoURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "UpdateGoodsinfo")
	return
}

// AddGoodsRequest 添加并提审商品请求
type AddGoodsRequest struct {
	GoodsInfo GoodsInfo `json:"goodsInfo"` // 商品信息
}

// AddGoodsResponse 添加并提审商品响应
type AddGoodsResponse struct {
	util.CommonError
	GoodsID int64 `json:"goodsId"` // 商品ID
	AuditID int64 `json:"auditId"` // 审核单ID
}

// AddGoods 添加并提审商品
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/commodity-management/api_addgoods
func (s *LiveBroadcast) AddGoods(in *AddGoodsRequest) (out AddGoodsResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(addGoodsURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "AddGoods")
	return
}

// GetGoodsInfoRequest 获取商品列表请求
type GetGoodsInfoRequest struct {
	Offset int64 `json:"offset"` // 分页偏移量
	Limit  int64 `json:"limit"`  // 分页大小
	Status int64 `json:"status"` // 商品状态
}

// GetGoodsInfoResponse 获取商品列表响应
type GetGoodsInfoResponse struct {
	util.CommonError
	Goods []Goods `json:"goods"` // 商品列表
	Total int64   `json:"total"` // 商品个数
}

// GetGoodsInfo 获取商品列表
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/commodity-management/api_getgoodsinfo
func (s *LiveBroadcast) GetGoodsInfo(in *GetGoodsInfoRequest) (out GetGoodsInfoResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	params := map[string]interface{}{
		"access_token": accessToken,
		"offset":       in.Offset,
	}
	if in.Limit > 0 {
		params["limit"] = in.Limit
	}
	if in.Status > 0 {
		params["status"] = in.Status
	}
	query := util.Query(params)

	response, err := util.HTTPGet(fmt.Sprintf(getGoodsInfoURL, query))
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetGoodsInfo")
	return
}

// GetGoodsAuditInfoRequest 获取商品的信息与审核状态请求
type GetGoodsAuditInfoRequest struct {
	GoodsIDs []int64 `json:"goods_ids"` // 商品ID
}

// GetGoodsAuditInfoResponse 获取商品的信息与审核状态响应
type GetGoodsAuditInfoResponse struct {
	util.CommonError
	Goods []GoodsAuditInfo `json:"goods"` // 商品列表
	Total int64            `json:"total"` // 商品个数
}

// GetGoodsAuditInfo 获取商品的信息与审核状态
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/commodity-management/api_getgoodsauditinfo
func (s *LiveBroadcast) GetGoodsAuditInfo(in *GetGoodsAuditInfoRequest) (out GetGoodsAuditInfoResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(getGoodsAuditInfoURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetGoodsAuditInfo")
	return
}

// ResubmitAuditRequest 重新提交商品审核请求
type ResubmitAuditRequest struct {
	GoodsID int64 `json:"goodsId"` // 商品ID
}

// ResubmitAuditResponse 重新提交商品审核响应
type ResubmitAuditResponse struct {
	util.CommonError
	AuditID int64 `json:"auditId"` // 审核单ID
}

// ResubmitAudit 重新提交商品审核
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/livebroadcast/commodity-management/api_resubmitaudit
func (s *LiveBroadcast) ResubmitAudit(in *ResubmitAuditRequest) (out ResubmitAuditResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(resubmitAuditURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "ResubmitAudit")
	return
}
