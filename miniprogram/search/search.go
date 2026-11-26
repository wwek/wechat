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

package search

import (
	"fmt"

	"github.com/silenceper/wechat/v2/miniprogram/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	// submitPagesURL 搜一搜数据推送
	submitPagesURL = "https://api.weixin.qq.com/wxa/search/wxaapi_submitpages?access_token=%s"
)

// Search 微信搜一搜
type Search struct {
	*context.Context
}

// NewSearch 实例化
func NewSearch(ctx *context.Context) *Search {
	return &Search{ctx}
}

// SubmitPagesAuthor 作者信息
type SubmitPagesAuthor struct {
	AuthorName     string `json:"author_name"`     // 作者名字
	AuthorTitle    string `json:"author_title"`    // 作者职务
	AuthorPortrait string `json:"author_portrait"` // 作者头像URL，头像尺寸不低于36px*36px。图片url访问如果有有效期，建议设置为15天以上
}

// SubmitPagesVideo 视频信息
type SubmitPagesVideo struct {
	VideoTitle  string `json:"video_title"`  // 视频标题，如不填，则视为与页面标题一致
	VideoLength int64  `json:"video_length"` // 视频时长，单位为秒，优先五分钟内短视频
	VideoImg    string `json:"video_img"`    // 视频封面图URL，尺寸不低于686px*288px。图片url访问如果有有效期，建议设置为15天以上
}

// SubmitPagesDataList 小程序页面的数据
type SubmitPagesDataList struct {
	Type        string                 `json:"@type"`         // 数据结构类型，用于标识目标业务系统
	Update      int64                  `json:"update"`        // 更新字段；内容更新按照新增处理，如果页面路径（path+query）相同，微信会做覆盖更新
	ContentID   string                 `json:"content_id"`    // 数据方自定义 id
	PageType    int64                  `json:"page_type"`     // 页面类型，固定填2
	H5URL       string                 `json:"h5_url"`        // H5链接，推荐填写，如果该页面有对应的H5链接，则填上
	Title       string                 `json:"title"`         // 标题，长度建议在20个字
	Abstract    []string               `json:"abstract"`      // 摘要，添加摘要有利于召回
	Referer     string                 `json:"referer"`       // HTTP Referer，如果图片有防盗链逻辑，需要设置referer头，用于图片下载
	CoverImgURL string                 `json:"cover_img_url"` // 封面图URL，图片url访问如果有有效期，建议设置为15天以上
	MainBody    string                 `json:"mainbody"`      // 正文，不可带有html标签
	Author      SubmitPagesAuthor      `json:"author"`        // 作者信息，推荐医疗类填写医生信息
	Video       []SubmitPagesVideo     `json:"video"`         // 视频
	TimePublish int64                  `json:"time_publish"`  // 发布时间，unix时间戳，单位秒
	TimeModify  int64                  `json:"time_modify"`   // 更新时间，unix时间戳，单位秒
	ExtraInfo   map[string]interface{} `json:"extra_info"`    // 补充字段，通用字段无法满足要求时，需要额外补充的字段，具体字段内容需要与微信协商
}

// SubmitPagesPage 小程序页面信息
type SubmitPagesPage struct {
	Path     string                `json:"path"`      // 以pages/开头的小程序页面路径
	Query    string                `json:"query"`     // 小程序页面请求参数
	DataList []SubmitPagesDataList `json:"data_list"` // 小程序页面的数据，一个页面可以同时提交多个结构化信息
}

// SubmitPagesRequest 搜一搜数据推送请求
type SubmitPagesRequest struct {
	Pages []SubmitPagesPage `json:"pages"` // 请求提交的小程序页面信息数组，一次可提交多个页面的信息。（注意：path+query标识唯一一个页面，微信侧会使用这个信息构造唯一id）
}

// SubmitPagesResponse 搜一搜数据推送响应
type SubmitPagesResponse struct {
	util.CommonError
}

// SubmitPages 搜一搜数据推送
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/wxsearch/api_submitpages
func (s *Search) SubmitPages(in *SubmitPagesRequest) (out SubmitPagesResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(submitPagesURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "SubmitPages")
	return
}
