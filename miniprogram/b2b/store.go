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

package b2b

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// getRetailOpenidListURL 全量授权门店查询
	getRetailOpenidListURL = "https://api.weixin.qq.com/wxa/business/getretailopenidlist?access_token=%s"
	// retailBusinessApplyURL 开通流程
	retailBusinessApplyURL = "https://api.weixin.qq.com/wxa/business/retailbusinessapply?access_token=%s"
	// getRetailInfoURL 门店信息查询
	getRetailInfoURL = "https://api.weixin.qq.com/wxa/business/getretailinfo?access_token=%s"
	// batchCreateRetailURL 预录入门店信息
	batchCreateRetailURL = "https://api.weixin.qq.com/wxa/business/batchcreateretail?access_token=%s"
)

// GetRetailOpenidListRequest 全量授权门店查询请求
type GetRetailOpenidListRequest struct {
	Limit       int64  `json:"limit"`        // 分页拉取最大返回结果数，取值范围 1-100
	PageContext string `json:"page_context"` // 分页上下文，首次调用传空值，后续填上次返回的 page_context
}

// GetRetailOpenidListResponse 全量授权门店查询响应
type GetRetailOpenidListResponse struct {
	util.CommonError
	OpenidList  []string `json:"openid_list"`  // openid 列表
	PageContext string   `json:"page_context"` // 分页上下文
}

// GetRetailOpenidList 全量授权门店查询
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/store_assistant/api_getretailopenidlist
func (s *B2B) GetRetailOpenidList(in *GetRetailOpenidListRequest) (out GetRetailOpenidListResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getRetailOpenidListURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetRetailOpenidList")
	return
}

// RetailBusinessApplyRequest 开通流程请求
type RetailBusinessApplyRequest struct {
	GoodsTypeList []string `json:"goods_type_list"` // 主营商品类型
	GoodsSaleList []string `json:"goods_sale_list"` // 主要线下销售渠道
	CoverNum      string   `json:"cover_num"`       // 门店覆盖数，如 "0-5千"、"5千-1万"、"1万-10万"、"10万-50万"、"50万以上"
	ServiceList   []string `json:"service_list"`    // 所需服务类型
	Description   string   `json:"description"`     // 小程序方案概述，长度限制 21-100 字
	ContactName   string   `json:"contact_name"`    // 联系人姓名，长度限制 1-7 字
	ContactPhone  string   `json:"contact_phone"`   // 联系人手机号
	ContactEmail  string   `json:"contact_email"`   // 联系人邮箱
}

// RetailBusinessApplyResponse 开通流程响应
type RetailBusinessApplyResponse struct {
	util.CommonError
}

// RetailBusinessApply 开通流程
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/store_assistant/api_retailbusinessapply
func (s *B2B) RetailBusinessApply(in *RetailBusinessApplyRequest) (out RetailBusinessApplyResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(retailBusinessApplyURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "RetailBusinessApply")
	return
}

// GetRetailInfoRequest 门店信息查询请求
type GetRetailInfoRequest struct {
	OpenID      string `json:"openid"`       // 门店管理员/员工 openid，与手机号二选一，都填写时优先使用 openid
	MobilePhone string `json:"mobile_phone"` // 手机号
}

// RetailStaff 门店员工信息
type RetailStaff struct {
	OpenID     string `json:"openid"`      // 管理员/员工 openid
	Role       int64  `json:"role"`        // 角色，1 管理员，2 员工
	CreateTime int64  `json:"create_time"` // 员工加入时间，管理员无该字段
}

// GetRetailInfoResponse 门店信息查询响应
type GetRetailInfoResponse struct {
	util.CommonError
	MobilePhone       string        `json:"mobile_phone"`        // 手机号
	RetailType        string        `json:"retail_type"`         // 一级门店类型
	SubRetailType     string        `json:"sub_retail_type"`     // 二级门店类型
	RetailAddress     string        `json:"retail_address"`      // 门店地址
	RetailName        string        `json:"retail_name"`         // 门店名称
	Identification    string        `json:"identification"`      // 营业执照注册号
	Principal         string        `json:"principal"`           // 企业名称
	LegalPersonName   string        `json:"legal_person_name"`   // 法人姓名
	OpenID            string        `json:"openid"`              // 门店管理员/员工 openid，手机号查询时返回管理员 openid
	Role              string        `json:"role"`                // 角色，1 管理员，2 员工，手机号查询时返回管理员
	Status            int64         `json:"status"`              // 认证状态，1 已完成认证
	AuthTime          int64         `json:"auth_time"`           // 认证时间戳
	GrantTime         int64         `json:"grant_time"`          // 授权时间戳
	Longitude         float64       `json:"longitude"`           // 经度，门店未提交定位信息时无该字段
	Latitude          float64       `json:"latitude"`            // 纬度，门店未提交定位信息时无该字段
	BusinessType      []string      `json:"business_type"`       // 一级主营商品
	OtherBusinessType string        `json:"other_business_type"` // 二级主营商品
	StaffList         []RetailStaff `json:"staff_list"`          // 员工列表，包含管理员
}

// GetRetailInfo 门店信息查询
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/store_assistant/api_getretailinfo
func (s *B2B) GetRetailInfo(in *GetRetailInfoRequest) (out GetRetailInfoResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getRetailInfoURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetRetailInfo")
	return
}

// RetailInfoItem 预录入门店信息
type RetailInfoItem struct {
	MobilePhone        string   `json:"mobile_phone"`        // 手机号
	RetailName         string   `json:"retail_name"`         // 门店名称，长度限制 1-100 个字符
	RetailType         string   `json:"retail_type"`         // 一级门店类型
	SubRetailType      string   `json:"sub_retail_type"`     // 二级门店类型，一级类型为"其他"时必填
	AddressProvince    string   `json:"address_province"`    // 门店地址，省
	AddressCity        string   `json:"address_city"`        // 门店地址，市
	AddressRegion      string   `json:"address_region"`      // 门店地址，区县
	AddressStreet      string   `json:"address_street"`      // 门店地址，街道详细地址
	RegistrationNumber string   `json:"registration_number"` // 营业执照注册号
	BizName            string   `json:"biz_name"`            // 企业名称
	CorporationName    string   `json:"corporation_name"`    // 法人姓名
	Latitude           float64  `json:"latitude"`            // 纬度
	Longitude          float64  `json:"longitude"`           // 经度
	BusinessType       []string `json:"business_type"`       // 一级主营商品
	OtherBusinessType  string   `json:"other_business_type"` // 二级主营商品，一级主营商品包含"其他"时必填
}

// BatchCreateRetailRequest 预录入门店信息请求
type BatchCreateRetailRequest struct {
	RetailInfoList []RetailInfoItem `json:"retail_info_list"` // 门店信息列表，每次调用最多可导入 100 个门店
}

// RetailFailureRecord 门店信息导入失败记录
type RetailFailureRecord struct {
	MobilePhone        string `json:"mobile_phone"`        // 手机号
	RegistrationNumber string `json:"registration_number"` // 营业执照注册号
	FailureCode        int64  `json:"failure_code"`        // 失败原因码
}

// BatchCreateRetailResponse 预录入门店信息响应
type BatchCreateRetailResponse struct {
	util.CommonError
	NumSuccess        int64                 `json:"num_success"`         // 导入成功数量
	NumFailure        int64                 `json:"num_failure"`         // 导入失败数量
	FailureRecordList []RetailFailureRecord `json:"failure_record_list"` // 失败记录列表
}

// BatchCreateRetail 预录入门店信息
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/store_assistant/api_batchcreateretail
func (s *B2B) BatchCreateRetail(in *BatchCreateRetailRequest) (out BatchCreateRetailResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(batchCreateRetailURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "BatchCreateRetail")
	return
}
