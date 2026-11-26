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

	"github.com/silenceper/wechat/v2/util"
)

// ListPackages 查询流量包详情
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/minidrama/usagedata/api_listpackages
func (s *MiniDrama) ListPackages(ctx context.Context, in *ListPackagesRequest) (out ListPackagesResponse, err error) {
	var address string
	if address, err = s.requestAddress(ctx, listPackages); err != nil {
		return
	}

	var response []byte
	if response, err = util.PostJSONContext(ctx, address, in); err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "ListPackages")
	return
}

// ListPackagesRequest 查询流量包详情请求
type ListPackagesRequest struct {
	Status    int `json:"status"`               // 按流量包状态过滤，取值有：0-不过滤，1-有效的，2-无效的，3-过期的
	Offset    int `json:"offset"`               // 偏移值，分页请求用
	Limit     int `json:"limit"`                // 分页大小，最大值为100
	QueryType int `json:"query_type,omitempty"` // 查询类型，0：通用流量；1：定向流量。默认值为0
}

// ListPackagesResponse 查询流量包详情响应
type ListPackagesResponse struct {
	util.CommonError
	TotalCount  int               `json:"total_count"`  // 总数
	PackageList []*TrafficPackage `json:"package_list"` // 流量包列表
}

// TrafficPackage 流量包
type TrafficPackage struct {
	StartTime int64           `json:"start_time"` // 有效期起始时间
	EndTime   int64           `json:"end_time"`   // 有效期截止时间
	Used      int64           `json:"used"`       // 流量包已消耗流量的数值（MB）
	All       int64           `json:"all"`        // 流量包总额（MB），剩余流量=流量包总额-已消耗流量
	OrderID   util.FlexString `json:"order_id"`   // 订单号（官方参数表为 string、返回示例为数字）
	Status    int             `json:"status"`     // 流量包状态，参考请求参数status中的说明
	IsDeleted int             `json:"is_deleted"` // 是否已删除（失效），例如已退款
	PackageID string          `json:"package_id"` // 流量包编号
}
