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
	// createChapterURL 上传章节
	createChapterURL = "https://api.weixin.qq.com/wxa/book/createchapter?access_token=%s"
	// auditBookURL 作品提审
	auditBookURL = "https://api.weixin.qq.com/wxa/book/auditbook?access_token=%s"
	// createBookURL 创建作品
	createBookURL = "https://api.weixin.qq.com/wxa/book/createbook?access_token=%s"
	// deleteBookURL 删除作品
	deleteBookURL = "https://api.weixin.qq.com/wxa/book/deletebook?access_token=%s"
	// deleteChapterURL 删除章节
	deleteChapterURL = "https://api.weixin.qq.com/wxa/book/deletechapter?access_token=%s"
	// batchCreateChapterURL 批量上传章节
	batchCreateChapterURL = "https://api.weixin.qq.com/wxa/book/batchcreatechapter?access_token=%s"
	// replaceChapterURL 替换章节
	replaceChapterURL = "https://api.weixin.qq.com/wxa/book/replacechapter?access_token=%s"
	// updateBookURL 编辑作品
	updateBookURL = "https://api.weixin.qq.com/wxa/book/updatebook?access_token=%s"
	// getBookURL 获取作品信息
	getBookURL = "https://api.weixin.qq.com/wxa/book/getbook?access_token=%s"
	// listBookURL 获取作品列表
	listBookURL = "https://api.weixin.qq.com/wxa/book/listbook?access_token=%s"
	// getChapterURL 获取章节信息
	getChapterURL = "https://api.weixin.qq.com/wxa/book/getchapter?access_token=%s"
	// listChapterURL 获取章节列表
	listChapterURL = "https://api.weixin.qq.com/wxa/book/listchapter?access_token=%s"
	// updateChapterSeqURL 调整章节相对顺序
	updateChapterSeqURL = "https://api.weixin.qq.com/wxa/book/updatechapterseq?access_token=%s"
	// reorderChapterURL 调整章节顺序
	reorderChapterURL = "https://api.weixin.qq.com/wxa/book/reorderchapter?access_token=%s"
)

// ChapterInfo 章节信息
type ChapterInfo struct {
	ChapterTitle string `json:"chapter_title"` // 章节标题，长度限制 1-80 字
	Content      string `json:"content"`       // 章节内容，长度限制 1-20000 字
	OriginalID   string `json:"original_id"`   // 提供方章节主键，可用于去重，长度限制 0-255 字节
	Seq          int64  `json:"seq"`           // 章节相对顺序，可非连续递增，默认值 0
	CustomInfo   string `json:"custom_info"`   // 自定义信息，长度限制 0-128 字节
}

// CreateChapterRequest 上传章节请求
type CreateChapterRequest struct {
	BookID  string       `json:"book_id"` // 作品 id
	Chapter *ChapterInfo `json:"chapter"` // 章节信息
}

// CreateChapterResponse 上传章节响应
type CreateChapterResponse struct {
	util.CommonError
	ChapterID string `json:"chapter_id"` // 章节 id
}

// CreateChapter 上传章节
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/novel/business/api_createchapter
func (s *Novel) CreateChapter(in *CreateChapterRequest) (out CreateChapterResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(createChapterURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "CreateChapter")
	return
}

// AuditBookRequest 作品提审请求
type AuditBookRequest struct {
	BookID string `json:"book_id"` // 作品 id
}

// AuditBookResponse 作品提审响应
type AuditBookResponse struct {
	util.CommonError
}

// AuditBook 作品提审
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/novel/business/api_auditbook
func (s *Novel) AuditBook(in *AuditBookRequest) (out AuditBookResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(auditBookURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "AuditBook")
	return
}

// CreateBookRequest 创建作品请求
type CreateBookRequest struct {
	Title              string   `json:"title"`                // 作品名，长度限制 1-30 字
	Intro              string   `json:"intro"`                // 作品简介，长度限制 1-500 字
	CoverMediaID       string   `json:"cover_media_id"`       // 封面图 media_id，通过新增临时素材接口上传得到
	Author             string   `json:"author"`               // 作者名，长度限制 1-100 字
	FirstCategoryID    int64    `json:"first_category_id"`    // 一级类型 id
	SecondCategoryID   int64    `json:"second_category_id"`   // 二级类型 id
	ThirdCategoryID    int64    `json:"third_category_id"`    // 三级类型 id
	CompleteStatus     int64    `json:"complete_status"`      // 完结状态，1 连载中，2 已完结
	OriginalID         string   `json:"original_id"`          // 提供方作品主键，可用于去重，长度限制 0-255 字节
	ChapterOrderMethod int64    `json:"chapter_order_method"` // 章节排序方式，0 追加，1 seq 递增，默认值 0
	CustomInfo         string   `json:"custom_info"`          // 自定义信息，长度限制 0-128 字节
	KeywordList        []string `json:"keyword_list"`         // 题材关键词，最多 3 个，每个长度限制 1-4 字
	AwesomeParagraph   string   `json:"awesome_paragraph"`    // 精彩片段，长度限制 400-1000 字
}

// CreateBookResponse 创建作品响应
type CreateBookResponse struct {
	util.CommonError
	BookID string `json:"book_id"` // 作品 id，非定长，不超过 64 字节
}

// CreateBook 创建作品
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/novel/business/api_createbook
func (s *Novel) CreateBook(in *CreateBookRequest) (out CreateBookResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(createBookURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "CreateBook")
	return
}

// DeleteBookRequest 删除作品请求
type DeleteBookRequest struct {
	BookID string `json:"book_id"` // 作品 id
}

// DeleteBookResponse 删除作品响应
type DeleteBookResponse struct {
	util.CommonError
}

// DeleteBook 删除作品
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/novel/business/api_deletebook
func (s *Novel) DeleteBook(in *DeleteBookRequest) (out DeleteBookResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(deleteBookURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "DeleteBook")
	return
}

// DeleteChapterRequest 删除章节请求
type DeleteChapterRequest struct {
	BookID    string `json:"book_id"`    // 作品 id
	ChapterID string `json:"chapter_id"` // 章节 id
}

// DeleteChapterResponse 删除章节响应
type DeleteChapterResponse struct {
	util.CommonError
}

// DeleteChapter 删除章节
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/novel/business/api_deletechapter
func (s *Novel) DeleteChapter(in *DeleteChapterRequest) (out DeleteChapterResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(deleteChapterURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "DeleteChapter")
	return
}

// BatchCreateChapterRequest 批量上传章节请求
type BatchCreateChapterRequest struct {
	BookID      string        `json:"book_id"`      // 作品 id
	ChapterList []ChapterInfo `json:"chapter_list"` // 章节信息，单次最多上传 10 章
}

// BatchCreateChapterResponse 批量上传章节响应
type BatchCreateChapterResponse struct {
	util.CommonError
	ChapterID              []string `json:"chapter_id"`                // 章节 id 列表
	ConflictOriginalIDList []string `json:"conflict_original_id_list"` // 冲突的提供方章节主键列表
}

// BatchCreateChapter 批量上传章节
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/novel/business/api_batchcreatechapter
func (s *Novel) BatchCreateChapter(in *BatchCreateChapterRequest) (out BatchCreateChapterResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(batchCreateChapterURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "BatchCreateChapter")
	return
}

// ReplaceChapterRequest 替换章节请求
type ReplaceChapterRequest struct {
	BookID          string `json:"book_id"`           // 作品 id
	ChapterID       string `json:"chapter_id"`        // 章节 id
	NewChapterTitle string `json:"new_chapter_title"` // 新章节标题，长度限制 1-80 字
	NewContent      string `json:"new_content"`       // 新章节内容，长度限制 1-20000 字
}

// ReplaceChapterResponse 替换章节响应
type ReplaceChapterResponse struct {
	util.CommonError
	NewChapterID string `json:"new_chapter_id"` // 替换后的新章节 id
}

// ReplaceChapter 替换章节
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/novel/business/api_replacechapter
func (s *Novel) ReplaceChapter(in *ReplaceChapterRequest) (out ReplaceChapterResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(replaceChapterURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "ReplaceChapter")
	return
}

// VolumeInfo 分卷信息
type VolumeInfo struct {
	VolumeTitle string `json:"volume_title"` // 分卷名，长度限制 1-100 字
	StartIndex  int64  `json:"start_index"`  // 分卷起始章节下标，取值范围 [0, 章节总数)
	EndIndex    int64  `json:"end_index"`    // 分卷截止章节下标，取值范围 [起始章节下标, 章节总数)
}

// UpdateBookRequest 编辑作品请求
type UpdateBookRequest struct {
	BookID             string       `json:"book_id"`              // 作品 id
	Title              string       `json:"title"`                // 作品名，长度限制 1-30 字
	Intro              string       `json:"intro"`                // 作品简介，长度限制 1-500 字
	CoverMediaID       string       `json:"cover_media_id"`       // 封面图 media_id，通过新增临时素材上传得到
	Author             string       `json:"author"`               // 作者名，长度限制 1-100 字
	FirstCategoryID    int64        `json:"first_category_id"`    // 一级类型 id
	SecondCategoryID   int64        `json:"second_category_id"`   // 二级类型 id
	ThirdCategoryID    int64        `json:"third_category_id"`    // 三级类型 id
	CompleteStatus     int64        `json:"complete_status"`      // 完结状态，1 连载中，2 已完结
	ChapterIDList      []string     `json:"chapter_id_list"`      // 按照预期顺序传入章节 id
	NeedVolume         bool         `json:"need_volume"`          // 是否需要分卷
	VolumeList         []VolumeInfo `json:"volume_list"`          // 分卷信息
	ChapterOrderMethod int64        `json:"chapter_order_method"` // 章节排序方式，0 追加，1 seq 递增
	CustomInfo         string       `json:"custom_info"`          // 自定义信息，长度限制 0-128 字节
	UpdateKeyword      bool         `json:"update_keyword"`       // 是否更新题材关键词
	KeywordList        []string     `json:"keyword_list"`         // 题材关键词，最多 3 个，每个长度限制 1-4 字
	AwesomeParagraph   string       `json:"awesome_paragraph"`    // 精彩片段，长度限制 400-1000 字
}

// UpdateBookResponse 编辑作品响应
type UpdateBookResponse struct {
	util.CommonError
}

// UpdateBook 编辑作品
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/novel/business/api_updatebook
func (s *Novel) UpdateBook(in *UpdateBookRequest) (out UpdateBookResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(updateBookURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "UpdateBook")
	return
}

// AuditInfo 审核信息
type AuditInfo struct {
	AuditStatus int64  `json:"audit_status"` // 0 未提审，1 审核中，2 审核不通过，3 审核通过
	CreateTime  int64  `json:"create_time"`  // 提审时间戳
	AuditTime   int64  `json:"audit_time"`   // 审核时间戳
	Reason      string `json:"reason"`       // 审核原因
	Suggestion  string `json:"suggestion"`   // 修改建议，审核不通过时才会有该字段
}

// BookInfo 作品信息
type BookInfo struct {
	BookID             string       `json:"book_id"`              // 作品 id
	Title              string       `json:"title"`                // 作品名
	Intro              string       `json:"intro"`                // 作品简介
	CoverURL           string       `json:"cover_url"`            // 封面图 url
	Author             string       `json:"author"`               // 作者名
	FirstCategoryID    int64        `json:"first_category_id"`    // 一级类型 id
	FirstCategoryName  string       `json:"first_category_name"`  // 一级类型名
	SecondCategoryID   int64        `json:"second_category_id"`   // 二级类型 id
	SecondCategoryName string       `json:"second_category_name"` // 二级类型名
	ThirdCategoryID    int64        `json:"third_category_id"`    // 三级类型 id
	ThirdCategoryName  string       `json:"third_category_name"`  // 三级类型名
	CompleteStatus     int64        `json:"complete_status"`      // 完结状态，1 连载中，2 已完结
	UploadScene        int64        `json:"upload_scene"`         // 上传场景，1 本地上传，2 API 上传
	ChapterCnt         int64        `json:"chapter_cnt"`          // 章节数
	VolumeCnt          int64        `json:"volume_cnt"`           // 分卷数
	VolumeList         []VolumeInfo `json:"volume_list"`          // 分卷信息
	TotalWordCnt       int64        `json:"total_word_cnt"`       // 作品总字数
	AuditInfo          *AuditInfo   `json:"audit_info"`           // 审核信息，未发起审核不返回该字段
	CreateTime         int64        `json:"create_time"`          // 创建时间戳
	OriginalID         string       `json:"original_id"`          // 提供方作品主键
	ChapterOrderMethod int64        `json:"chapter_order_method"` // 章节排序方式，0 追加，1 seq 递增
	CustomInfo         string       `json:"custom_info"`          // 自定义信息
	BanStatus          int64        `json:"ban_status"`           // 管控状态，0 正常，1 下架
}

// GetBookRequest 获取作品信息请求
type GetBookRequest struct {
	BookID         string `json:"book_id"`          // 作品 id，与提供方作品主键二选一，优先使用该字段
	NeedEditedData bool   `json:"need_edited_data"` // true 编辑版信息，false 发布版信息，默认值 false
	OriginalID     string `json:"original_id"`      // 提供方作品主键，与作品 id 二选一
}

// GetBookResponse 获取作品信息响应
type GetBookResponse struct {
	util.CommonError
	Book []BookInfo `json:"book"` // 作品信息
}

// GetBook 获取作品信息
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/novel/business/api_getbook
func (s *Novel) GetBook(in *GetBookRequest) (out GetBookResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getBookURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetBook")
	return
}

// ListBookRequest 获取作品列表请求
type ListBookRequest struct {
	Limit          int64 `json:"limit"`            // 分页拉取的最大返回结果数，取值范围 1-100，默认值 100
	Offset         int64 `json:"offset"`           // 分页拉取的起始偏移量，默认值 0，与 last_id 二选一，优先使用 last_id
	LastID         int64 `json:"last_id"`          // 分页 id，首次调用填 0，后续填上次返回的 last_id
	NeedEditedData bool  `json:"need_edited_data"` // true 编辑版信息，false 发布版信息，默认值 false
}

// ListBookResponse 获取作品列表响应
type ListBookResponse struct {
	util.CommonError
	BookList []BookInfo `json:"book_list"` // 作品列表
	TotalCnt int64      `json:"total_cnt"` // 作品总数
	LastID   int64      `json:"last_id"`   // 分页 id
}

// ListBook 获取作品列表
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/novel/business/api_listbook
func (s *Novel) ListBook(in *ListBookRequest) (out ListBookResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(listBookURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "ListBook")
	return
}

// ChapterDetail 章节详情
type ChapterDetail struct {
	ChapterID    string     `json:"chapter_id"`    // 章节 id
	ChapterTitle string     `json:"chapter_title"` // 章节标题
	WordCnt      int64      `json:"word_cnt"`      // 字数
	CreateTime   int64      `json:"create_time"`   // 创建时间戳
	AuditInfo    *AuditInfo `json:"audit_info"`    // 审核信息
	VolumeIndex  int64      `json:"volume_index"`  // 所属分卷下标，-1 表示不属于任何分卷
	OriginalID   string     `json:"original_id"`   // 提供方章节主键
	Seq          int64      `json:"seq"`           // 章节相对顺序
	CustomInfo   string     `json:"custom_info"`   // 自定义信息
	BanStatus    int64      `json:"ban_status"`    // 管控状态，0 正常，1 下架
}

// GetChapterRequest 获取章节信息请求
type GetChapterRequest struct {
	BookID         string `json:"book_id"`          // 作品 id
	ChapterID      string `json:"chapter_id"`       // 章节 id
	NeedEditedData bool   `json:"need_edited_data"` // true 编辑版信息，false 发布版信息，默认值 false
}

// GetChapterResponse 获取章节信息响应
type GetChapterResponse struct {
	util.CommonError
	Chapter *ChapterDetail `json:"chapter"` // 章节信息
}

// GetChapter 获取章节信息
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/novel/business/api_getchapter
func (s *Novel) GetChapter(in *GetChapterRequest) (out GetChapterResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getChapterURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetChapter")
	return
}

// ListChapterRequest 获取章节列表请求
type ListChapterRequest struct {
	BookID         string `json:"book_id"`          // 作品 id
	NeedEditedData bool   `json:"need_edited_data"` // true 编辑版信息，false 发布版信息，默认值 false
	Limit          int64  `json:"limit"`            // 分页拉取的最大返回结果数，默认值 10，最大值 100
	Offset         int64  `json:"offset"`           // 分页拉取的起始偏移量，默认值 0
	VolumeIndex    int64  `json:"volume_index"`     // 分卷下标，若设置则只返回指定卷的章节列表
}

// ListChapterResponse 获取章节列表响应
type ListChapterResponse struct {
	util.CommonError
	ChapterList []ChapterDetail `json:"chapter_list"` // 章节列表
	TotalCnt    int64           `json:"total_cnt"`    // 章节总数
}

// ListChapter 获取章节列表
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/novel/business/api_listchapter
func (s *Novel) ListChapter(in *ListChapterRequest) (out ListChapterResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(listChapterURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "ListChapter")
	return
}

// ChapterSeqInfo 章节顺序信息
type ChapterSeqInfo struct {
	ChapterID string `json:"chapter_id"` // 章节 id
	Seq       int64  `json:"seq"`        // 章节相对顺序，可非连续递增
}

// UpdateChapterSeqRequest 调整章节相对顺序请求
type UpdateChapterSeqRequest struct {
	BookID         string           `json:"book_id"`          // 作品 id
	ChapterSeqList []ChapterSeqInfo `json:"chapter_seq_list"` // 章节 seq 列表
}

// UpdateChapterSeqResponse 调整章节相对顺序响应
type UpdateChapterSeqResponse struct {
	util.CommonError
}

// UpdateChapterSeq 调整章节相对顺序
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/novel/business/api_updatechapterseq
func (s *Novel) UpdateChapterSeq(in *UpdateChapterSeqRequest) (out UpdateChapterSeqResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(updateChapterSeqURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "UpdateChapterSeq")
	return
}

// ReorderChapterRequest 调整章节顺序请求
type ReorderChapterRequest struct {
	BookID          string `json:"book_id"`           // 作品 id
	ChapterID       string `json:"chapter_id"`        // 章节 id
	TargetChapterID string `json:"target_chapter_id"` // 目标章节 id
	Operation       int64  `json:"operation"`         // 排序操作，1 交换，2 插入到目标章节之前，3 插入到目标章节之后
}

// ReorderChapterResponse 调整章节顺序响应
type ReorderChapterResponse struct {
	util.CommonError
}

// ReorderChapter 调整章节顺序
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/novel/business/api_reorderchapter
func (s *Novel) ReorderChapter(in *ReorderChapterRequest) (out ReorderChapterResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(reorderChapterURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "ReorderChapter")
	return
}
