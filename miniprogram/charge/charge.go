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

package charge

import (
	"fmt"

	"github.com/silenceper/wechat/v2/miniprogram/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	// getUsageDetailURL 查询购买资源包的用量情况
	getUsageDetailURL = "https://api.weixin.qq.com/wxa/charge/usage/get?%s"
	// getRecentAverageUsageURL 获取小程序某个付费能力的最近用量数据
	getRecentAverageUsageURL = "https://api.weixin.qq.com/wxa/charge/usage/get_recent_average?%s"
)

// Charge 付费管理
type Charge struct {
	*context.Context
}

// NewCharge 实例化
func NewCharge(ctx *context.Context) *Charge {
	return &Charge{ctx}
}

// GetUsageDetailRequest 查询购买资源包的用量情况请求
type GetUsageDetailRequest struct {
	SpuID  string `json:"spuId"`  // 商品SPU ID
	Offset int64  `json:"offset"` // 分页偏移量，从0开始
	Limit  int64  `json:"limit"`  // 每页个数，最大20
}

// GetUsageDetailDetail 用量详情列表项
type GetUsageDetailDetail struct {
	PkgID     string `json:"pkgId"`     // 资源包ID
	Status    int64  `json:"status"`    // 资源包状态
	StartTime int64  `json:"startTime"` // 额度有效期开始时间戳（单位：秒）
	EndTime   int64  `json:"endTime"`   // 额度有效期至结束时间戳（单位：秒）
	All       string `json:"all"`       // 额度容量（64位数字）
	Used      string `json:"used"`      // 使用额度（64位数字）
	SpuID     string `json:"spuId"`     // 额度来源的商品SPU ID（64位数字）
	SkuID     string `json:"skuId"`     // 额度来源的商品SKU ID（64位数字）
	Source    int64  `json:"source"`    // 额度来源
}

// GetUsageDetailResponse 查询购买资源包的用量情况响应
type GetUsageDetailResponse struct {
	util.CommonError
	EffectiveUse     string                 `json:"effectiveUse"`     // 累计用量（64位数字），用于资源包类商品
	EffectiveAll     string                 `json:"effectiveAll"`     // 资源总量（64位数字），用于资源包类商品（已废弃）
	All              string                 `json:"all"`              // 资源可用总量（64位数字），用于资源包类商品
	StartServiceTime int64                  `json:"startServiceTime"` // 订阅开始时间戳（单位：秒），用于订阅类商品
	EndServiceTime   int64                  `json:"endServiceTime"`   // 订阅结束时间戳（单位：秒），用于订阅类商品
	Total            int64                  `json:"total"`            // 用量详情列表总数
	DetailList       []GetUsageDetailDetail `json:"detailList"`       // 用量详情列表
}

// GetUsageDetail 查询购买资源包的用量情况
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/charge/api_getusagedetail
func (s *Charge) GetUsageDetail(in *GetUsageDetailRequest) (out GetUsageDetailResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(getUsageDetailURL, util.Query(map[string]interface{}{
		"access_token": accessToken,
		"spuId":        in.SpuID,
		"offset":       in.Offset,
		"limit":        in.Limit,
	}))
	response, err := util.HTTPGet(uri)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetUsageDetail")
	return
}

// GetRecentAverageUsageRequest 获取小程序某个付费能力的最近用量数据请求
type GetRecentAverageUsageRequest struct {
	SpuID string `json:"spuId"` // 商品SPU ID
}

// GetRecentAverageUsageResponse 获取小程序某个付费能力的最近用量数据响应
type GetRecentAverageUsageResponse struct {
	util.CommonError
	AverageData int64 `json:"averageData"` // 最近月平均用量，经模糊化处理，非精确值。注意：当averageData返回值为50时，语义为小程序最近平均用量小于等于50次/月，并不是特指精确等于每月50次
}

// GetRecentAverageUsage 获取小程序某个付费能力的最近用量数据
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/charge/api_getrecentaverageusage
func (s *Charge) GetRecentAverageUsage(in *GetRecentAverageUsageRequest) (out GetRecentAverageUsageResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(getRecentAverageUsageURL, util.Query(map[string]interface{}{
		"access_token": accessToken,
		"spuId":        in.SpuID,
	}))
	response, err := util.HTTPGet(uri)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetRecentAverageUsage")
	return
}
