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

package cityservice

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// previewHopsNoticeURL 公告草稿预览设置
	previewHopsNoticeURL = "https://api.weixin.qq.com/intp/eldermedical/previewhopsnotice?access_token=%s"
	// getHospNoticeListURL 查询所有公告
	getHospNoticeListURL = "https://api.weixin.qq.com/intp/eldermedical/gethospnoticelist?access_token=%s"
	// getMedRealNameURL 查询用户实名API
	getMedRealNameURL = "https://api.weixin.qq.com/cityservice/getmedrealname?access_token=%s"
	// getMsgRelationURL 查询用户是否开通
	getMsgRelationURL = "https://api.weixin.qq.com/cityservice/getmsgrelation?access_token=%s"
	// publishHopsNoticeURL 正式发布公告
	publishHopsNoticeURL = "https://api.weixin.qq.com/intp/eldermedical/publichopsnotice?access_token=%s"
	// setHopsNoticeURL 添加公告草稿
	setHopsNoticeURL = "https://api.weixin.qq.com/intp/eldermedical/sethopsnotice?access_token=%s"
)

// PreviewHopsNoticeRequest 公告草稿预览设置请求
type PreviewHopsNoticeRequest struct {
	AppID           string `json:"app_id"`           // 业务方公众号的appid
	NoticeType      int64  `json:"notice_type"`      // 公告类型
	NoticeID        string `json:"notice_id"`        // 公告ID
	PreviewUsername string `json:"preview_username"` // 可以预览公告的微信号
	Operation       int64  `json:"operation"`        // 操作类型
}

// PreviewHopsNoticeResponse 公告草稿预览设置响应
type PreviewHopsNoticeResponse struct {
	util.CommonError
	NoticeID int64 `json:"notice_id"` // 草稿ID
}

// PreviewHopsNotice 公告草稿预览设置
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/cityservice/elderMedical/api_previewhopsnotice
func (s *CityService) PreviewHopsNotice(in *PreviewHopsNoticeRequest) (out PreviewHopsNoticeResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(previewHopsNoticeURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "PreviewHopsNotice")
	return
}

// GetHospNoticeListRequest 查询所有公告请求
type GetHospNoticeListRequest struct {
	AppID      string `json:"app_id"`      // 业务方公众号的appid
	NoticeType int64  `json:"notice_type"` // 公告类型
}

// GetHospNoticeListResponse 查询所有公告响应
type GetHospNoticeListResponse struct {
	util.CommonError
	NoticeList []HospNotice `json:"notice_list"` // 公告列表
}

// GetHospNoticeList 查询所有公告
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/cityservice/elderMedical/api_intp_eldermed_gethospnoticelist
func (s *CityService) GetHospNoticeList(in *GetHospNoticeListRequest) (out GetHospNoticeListResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getHospNoticeListURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetHospNoticeList")
	return
}

// HospNotice 公告
type HospNotice struct {
	NoticeID      int64    `json:"notice_id"`      // 公告ID
	Content       string   `json:"content"`        // 公告内容
	Status        string   `json:"status"`         // 公告状态
	PreviewOpenID []string `json:"preview_openid"` // 可以预览草稿的用户 openid
}

// GetMedRealNameRequest 查询用户实名API请求
type GetMedRealNameRequest struct {
	AppID         string `json:"app_id"`         // 业务方appid
	OpenID        string `json:"open_id"`        // 微信用户openid
	WxmedAuthCode string `json:"wxmed_authcode"` // 实名信息code，对应url中的wxmed_authcode，有效期10分钟
}

// GetMedRealNameResponse 查询用户实名API响应
type GetMedRealNameResponse struct {
	util.CommonError
	CipherRealName  string `json:"cipher_real_name"` // 加密后的实名信息，base64后的加密信息
	CipherAlgorithm string `json:"cipher_algorithm"` // 加密算法，默认：AES_256_ECB_PKCS7Padding
	KeyVersion      int64  `json:"key_version"`      // 实名信息加密key版本号，初始版本号为0，用于识别后续密钥更换升级
	AppID           string `json:"app_id"`           // 业务方appid
	OpenID          string `json:"open_id"`          // 微信用户openid
}

// GetMedRealName 查询用户实名API
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/cityservice/elderMedical/api_cityservice_getmedrealname
func (s *CityService) GetMedRealName(in *GetMedRealNameRequest) (out GetMedRealNameResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getMedRealNameURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetMedRealName")
	return
}

// GetMsgRelationRequest 查询用户是否开通请求
type GetMsgRelationRequest struct {
	BusinessID string `json:"business_id"` // 业务id，如130
	OpenID     string `json:"open_id"`     // 微信用户openid
}

// GetMsgRelationResponse 查询用户是否开通响应
type GetMsgRelationResponse struct {
	util.CommonError
	IsSubscribed int64 `json:"is_subscribed"` // 是否订阅（官方参数表为 boolean、返回示例为数字）
}

// UnmarshalJSON 该接口返回的错误码字段是 err_code/err_msg（与通用的 errcode/errmsg 不同），
// 这里做一次转换，保证 util.DecodeWithError 能正确识别错误。
func (r *GetMsgRelationResponse) UnmarshalJSON(data []byte) error {
	type alias GetMsgRelationResponse
	aux := &struct {
		ErrCode *int64  `json:"err_code"`
		ErrMsg  *string `json:"err_msg"`
		*alias
	}{alias: (*alias)(r)}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	if aux.ErrCode != nil {
		r.ErrCode = *aux.ErrCode
	}
	if aux.ErrMsg != nil {
		r.ErrMsg = *aux.ErrMsg
	}
	return nil
}

// GetMsgRelation 查询用户是否开通
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/cityservice/elderMedical/api_cityservice_getmsgrelation
func (s *CityService) GetMsgRelation(in *GetMsgRelationRequest) (out GetMsgRelationResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getMsgRelationURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetMsgRelation")
	return
}

// PublishHopsNoticeRequest 正式发布公告请求
type PublishHopsNoticeRequest struct {
	AppID      string `json:"app_id"`      // 业务方公众号的appid
	NoticeType int64  `json:"notice_type"` // 公告类型
	NoticeID   int64  `json:"notice_id"`   // 公告ID
}

// PublishHopsNoticeResponse 正式发布公告响应
type PublishHopsNoticeResponse struct {
	util.CommonError
	NoticeID int64 `json:"notice_id"` // 公告ID
}

// PublishHopsNotice 正式发布公告
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/cityservice/elderMedical/api_intp_eldermed_publichopsnotice
func (s *CityService) PublishHopsNotice(in *PublishHopsNoticeRequest) (out PublishHopsNoticeResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(publishHopsNoticeURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "PublishHopsNotice")
	return
}

// SetHopsNoticeRequest 添加公告草稿请求
type SetHopsNoticeRequest struct {
	AppID         string `json:"app_id"`         // 业务方公众号的appid
	NoticeType    int64  `json:"notice_type"`    // 公告类型
	NoticeContent string `json:"notice_content"` // 公告内容，长度限制3000个字符，**支持富文本**，图片用标签的方式嵌入富文本内，具体可看下文
	NoticeID      int64  `json:"notice_id"`      // 公告ID，不传notice_id则新增一条草稿，传notice_id不新增，会覆盖指定草稿的内容
}

// SetHopsNoticeResponse 添加公告草稿响应
type SetHopsNoticeResponse struct {
	util.CommonError
	NoticeID int64 `json:"notice_id"` // 草稿ID
}

// SetHopsNotice 添加公告草稿
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/cityservice/elderMedical/api_sethopsnotice
func (s *CityService) SetHopsNotice(in *SetHopsNoticeRequest) (out SetHopsNoticeResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(setHopsNoticeURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "SetHopsNotice")
	return
}
