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

package novel

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// setPreviewSettingURL 预览设置修改
	setPreviewSettingURL = "https://api.weixin.qq.com/wxa/business/novelreader/setpreviewsetting?access_token=%s"
	// getPreviewSettingURL 预览设置获取
	getPreviewSettingURL = "https://api.weixin.qq.com/wxa/business/novelreader/getpreviewsetting?access_token=%s"
)

// SetPreviewSettingRequest 预览设置修改请求
type SetPreviewSettingRequest struct {
	BookID       string `json:"book_id"`       // 作品 id
	DefaultWords int64  `json:"default_words"` // 默认预览字数
	ChapterIndex int64  `json:"chapter_index"` // 章节索引，从 0 开始，未设置的章节将使用 default_words 作为预览字数
	Words        int64  `json:"words"`         // 章节字数
}

// SetPreviewSettingResponse 预览设置修改响应
type SetPreviewSettingResponse struct {
	util.CommonError
}

// SetPreviewSetting 预览设置修改
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/novel/preview/api_setpreviewsetting
func (s *Novel) SetPreviewSetting(in *SetPreviewSettingRequest) (out SetPreviewSettingResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(setPreviewSettingURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "SetPreviewSetting")
	return
}

// ChapterPreviewSetting 章节预览设置
type ChapterPreviewSetting struct {
	ChapterIndex int64 `json:"chapter_index"` // 章节索引，从 0 开始，未设置的章节将使用 default_words 作为预览字数
	Words        int64 `json:"words"`         // 章节字数
}

// PreviewSetting 预览设置
type PreviewSetting struct {
	BookID         string                  `json:"book_id"`         // 作品 id
	DefaultWords   int64                   `json:"default_words"`   // 默认预览字数
	ChapterSetting []ChapterPreviewSetting `json:"chapter_setting"` // 章节预览设置列表
}

// GetPreviewSettingRequest 预览设置获取请求
type GetPreviewSettingRequest struct {
	BookID string `json:"book_id"` // 作品 id
}

// GetPreviewSettingResponse 预览设置获取响应
type GetPreviewSettingResponse struct {
	util.CommonError
	Setting *PreviewSetting `json:"setting"` // 预览设置
}

// GetPreviewSetting 预览设置获取
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/novel/preview/api_getpreviewsetting
func (s *Novel) GetPreviewSetting(in *GetPreviewSettingRequest) (out GetPreviewSettingResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getPreviewSettingURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetPreviewSetting")
	return
}
