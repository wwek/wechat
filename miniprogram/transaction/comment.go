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

package transaction

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// addCommentReplyURL 创建回复
	addCommentReplyURL = "https://api.weixin.qq.com/wxaapi/comment/commentreply/add?access_token=%s"
	// addReplyURL 创建评论
	addReplyURL = "https://api.weixin.qq.com/wxaapi/comment/reply/add?access_token=%s"
	// deleteCommentReplyURL 删除回复
	deleteCommentReplyURL = "https://api.weixin.qq.com/wxaapi/comment/commentreply/delete?access_token=%s"
	// deleteReplyURL 删除评论
	deleteReplyURL = "https://api.weixin.qq.com/wxaapi/comment/reply/delete?access_token=%s"
	// getCommentListURL 查询评价列表
	getCommentListURL = "https://api.weixin.qq.com/wxaapi/comment/mpcommentlist/get?%s"
	// getCommentInfoURL 查询评价详情
	getCommentInfoURL = "https://api.weixin.qq.com/wxaapi/comment/commentinfo/get?%s"
	// getCommentReplyListURL 查询评论列表
	getCommentReplyListURL = "https://api.weixin.qq.com/wxaapi/comment/replyandcommentreplylist/get?%s"
	// confirmCompromiseURL 确认和解
	confirmCompromiseURL = "https://api.weixin.qq.com/wxaapi/comment/confirmcompromise?access_token=%s"
	// resetApiKfQuotaURL 重置Api客服quota
	resetApiKfQuotaURL = "https://api.weixin.qq.com/wxaapi/comment/apikfquota/reset?access_token=%s"
)

// AddCommentReplyRequest 创建回复请求
type AddCommentReplyRequest struct {
	CommentID string `json:"commentId"` // 评价的 id
	ReplyID   string `json:"replyId"`   // 评论的 id
	Content   string `json:"content"`   // 回复评论的内容
}

// AddCommentReplyResponse 创建回复响应
type AddCommentReplyResponse struct {
	util.CommonError
	Success util.FlexBool `json:"success"` // 请求是否成功状态（官方返回示例中可能是字符串）
}

// AddCommentReply 创建回复
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/transaction-guarantee/comment/api_addcommentreply
func (s *Transaction) AddCommentReply(in *AddCommentReplyRequest) (out AddCommentReplyResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(addCommentReplyURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "AddCommentReply")
	return
}

// AddReplyRequest 创建评论请求
type AddReplyRequest struct {
	CommentID string `json:"commentId"` // 评价的 id
	Content   string `json:"content"`   // 评论的内容
}

// AddReplyResponse 创建评论响应
type AddReplyResponse struct {
	util.CommonError
	Success util.FlexBool `json:"success"` // 请求是否成功状态（官方返回示例中可能是字符串）
}

// AddReply 创建评论
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/transaction-guarantee/comment/api_addreply
func (s *Transaction) AddReply(in *AddReplyRequest) (out AddReplyResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(addReplyURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "AddReply")
	return
}

// DeleteCommentReplyRequest 删除回复请求
type DeleteCommentReplyRequest struct {
	CommentID      string `json:"commentId"`      // 评价的 id
	ReplyID        string `json:"replyId"`        // 评论的 id
	CommentReplyID string `json:"commentReplyId"` // 回复评论的 id
}

// DeleteCommentReplyResponse 删除回复响应
type DeleteCommentReplyResponse struct {
	util.CommonError
	Success util.FlexBool `json:"success"` // 请求是否成功状态（官方返回示例中可能是字符串）
}

// DeleteCommentReply 删除回复
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/transaction-guarantee/comment/api_deletecommentreply
func (s *Transaction) DeleteCommentReply(in *DeleteCommentReplyRequest) (out DeleteCommentReplyResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(deleteCommentReplyURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "DeleteCommentReply")
	return
}

// DeleteReplyRequest 删除评论请求
type DeleteReplyRequest struct {
	CommentID string `json:"commentId"` // 评价的 id
	ReplyID   string `json:"replyId"`   // 评论的 id
}

// DeleteReplyResponse 删除评论响应
type DeleteReplyResponse struct {
	util.CommonError
	Success util.FlexBool `json:"success"` // 请求是否成功状态（官方返回示例中可能是字符串）
}

// DeleteReply 删除评论
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/transaction-guarantee/comment/api_deletereply
func (s *Transaction) DeleteReply(in *DeleteReplyRequest) (out DeleteReplyResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(deleteReplyURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "DeleteReply")
	return
}

// GetCommentListRequest 查询评价列表请求
type GetCommentListRequest struct {
	StartTime  string `json:"startTime"`  // 查询时间段的开始时间
	EndTime    string `json:"endTime"`    // 查询时间段的结束时间
	FilterType int64  `json:"filterType"` // 过滤的数据类型
	Offset     int64  `json:"offset"`     // 查询的偏移数（从offset开始计数拉取)，默认值为 0 代表首页拉取
	Limit      int64  `json:"limit"`      // 查询每页中的数量，默认值为 8
}

// GetCommentListResponse 查询评价列表响应
type GetCommentListResponse struct {
	util.CommonError
	Offset      int64         `json:"offset"`      // 查询的偏移数（从offset开始计数拉取)
	Total       int64         `json:"total"`       // 评价总数
	Success     util.FlexBool `json:"success"`     // 请求是否成功（官方返回示例中可能是字符串）
	CommentList []CommentItem `json:"commentList"` // 评价的列表
}

// GetCommentList 查询评价列表
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/transaction-guarantee/comment/api_getccommentlist
func (s *Transaction) GetCommentList(in *GetCommentListRequest) (out GetCommentListResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	params := map[string]interface{}{
		"access_token": accessToken,
		"startTime":    in.StartTime,
		"endTime":      in.EndTime,
	}
	if in.FilterType > 0 {
		params["filterType"] = in.FilterType
	}
	if in.Offset > 0 {
		params["offset"] = in.Offset
	}
	if in.Limit > 0 {
		params["limit"] = in.Limit
	}

	response, err := util.HTTPGet(fmt.Sprintf(getCommentListURL, util.Query(params)))
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetCommentList")
	return
}

// GetCommentInfoRequest 查询评价详情请求
type GetCommentInfoRequest struct {
	CommentID string `json:"commentId"` // 评价的 id
}

// GetCommentInfoResponse 查询评价详情响应
type GetCommentInfoResponse struct {
	util.CommonError
	Info        CommentInfo       `json:"info"`        // 评价的信息
	ProcessInfo CommentProcess    `json:"processInfo"` // 进度条的信息（只有差评才会有）
	OldComment  CommentOldComment `json:"oldComment"`  // 旧评价的信息(只有改评的新评价才有)
}

// GetCommentInfo 查询评价详情
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/transaction-guarantee/comment/api_getcommentinfo
func (s *Transaction) GetCommentInfo(in *GetCommentInfoRequest) (out GetCommentInfoResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	params := map[string]interface{}{
		"access_token": accessToken,
		"commentId":    in.CommentID,
	}

	response, err := util.HTTPGet(fmt.Sprintf(getCommentInfoURL, util.Query(params)))
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetCommentInfo")
	return
}

// GetCommentReplyListRequest 查询评论列表请求
type GetCommentReplyListRequest struct {
	CommentID string `json:"commentId"` // 评价的 id
}

// GetCommentReplyListResponse 查询评论列表响应
type GetCommentReplyListResponse struct {
	util.CommonError
	List CommentReplyList `json:"list"` // 评论（评论回复第一条内容）
}

// GetCommentReplyList 查询评论列表
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/transaction-guarantee/comment/api_getcommentreplylist
func (s *Transaction) GetCommentReplyList(in *GetCommentReplyListRequest) (out GetCommentReplyListResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	params := map[string]interface{}{
		"access_token": accessToken,
		"commentId":    in.CommentID,
	}

	response, err := util.HTTPGet(fmt.Sprintf(getCommentReplyListURL, util.Query(params)))
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetCommentReplyList")
	return
}

// ConfirmCompromiseRequest 确认和解请求
type ConfirmCompromiseRequest struct {
	CommentID string   `json:"commentId"` // 评价的 id
	PicList   []string `json:"picList"`   // 和解的图片mediaId集合,可参考 素材管理
	Content   string   `json:"content"`   // 和解的文本内容
}

// ConfirmCompromiseResponse 确认和解响应
type ConfirmCompromiseResponse struct {
	util.CommonError
	Success util.FlexBool `json:"success"` // 请求是否成功状态（官方返回示例中可能是字符串）
}

// ConfirmCompromise 确认和解
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/transaction-guarantee/comment/api_confirmcompromise
func (s *Transaction) ConfirmCompromise(in *ConfirmCompromiseRequest) (out ConfirmCompromiseResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(confirmCompromiseURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "ConfirmCompromise")
	return
}

// ResetApiKfQuotaRequest 重置Api客服quota请求
type ResetApiKfQuotaRequest struct {
	CommentID string `json:"commentId"` // 评价的 id
}

// ResetApiKfQuotaResponse 重置Api客服quota响应
type ResetApiKfQuotaResponse struct {
	util.CommonError
	Success util.FlexBool `json:"success"` // 请求是否成功状态（官方返回示例中可能是字符串）
}

// ResetApiKfQuota 重置Api客服quota
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/transaction-guarantee/comment/api_resetapikfquota
func (s *Transaction) ResetApiKfQuota(in *ResetApiKfQuotaRequest) (out ResetApiKfQuotaResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(resetApiKfQuotaURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "ResetApiKfQuota")
	return
}

// CommentItem 评价信息
type CommentItem struct {
	CommentID   string             `json:"commentId"`   // 评价id
	Amount      util.FlexInt64     `json:"amount"`      // 金额,单位是分（官方参数表为 string、返回示例为数字）
	OrderID     string             `json:"orderId"`     // 订单id
	CreateTime  string             `json:"createTime"`  // 创建时间，单位是秒
	PayTime     string             `json:"payTime"`     // 支付时间，单位是秒
	WxPayID     string             `json:"wxPayId"`     // 微信支付交易单号，一般以420开头，是微信支付接口文档中的transaction_id
	OrderInfo   CommentOrderInfo   `json:"orderInfo"`   // 商家订单信息
	UserInfo    CommentUserInfo    `json:"userInfo"`    // 评价用户信息
	BizInfo     CommentBizInfo     `json:"bizInfo"`     // 商家小程序信息
	Score       util.FlexInt64     `json:"score"`       // 评价分数（每100分对应1星，官方参数表为 string、返回示例为数字）
	Content     CommentContent     `json:"content"`     // 评价内容
	ExtInfo     CommentExtInfo     `json:"extInfo"`     // 评价额外信息
	ProductInfo CommentProductInfo `json:"productInfo"` // 评价商品信息
}

// CommentOrderInfo 商家订单信息
type CommentOrderInfo struct {
	BusiOrderID string `json:"busiOrderId"` // 商户单号：商户系统内部订单号，只能是数字、大小写字母_-*且在同一个商户号下唯一；是微信支付接口文档中的 out_trade_no
}

// CommentUserInfo 评价用户信息
type CommentUserInfo struct {
	OpenID   string `json:"openid"`   // 评价用户openid
	HeadImg  string `json:"headImg"`  // 评价用户的头像
	NickName string `json:"nickName"` // 评价用户的昵称
}

// CommentBizInfo 商家小程序信息
type CommentBizInfo struct {
	AppID    string `json:"appid"`    // 商家小程序的appid
	HeadImg  string `json:"headImg"`  // 商家小程序的头像
	NickName string `json:"nickName"` // 商家小程序的昵称
}

// CommentContent 评价内容
type CommentContent struct {
	Txt   string         `json:"txt"`   // 评价内容
	Media []CommentMedia `json:"media"` // 评价的媒体文件，如图片、视频
}

// CommentMedia 评价的媒体文件
type CommentMedia struct {
	Img           string `json:"img"`           // 图片cdn
	ThumbImg      string `json:"thumbImg"`      // 缩略图cdn
	Video         string `json:"video"`         // 视频资源cdn（有img的时候没有video）
	VideoCover    string `json:"videoCover"`    // 缩略图cdn（有img的时候没有videoCover）
	VideoDuration int64  `json:"videoDuration"` // 视频时长，单位秒（有img的时候没有videoDuration）
}

// CommentExtInfo 评价额外信息
type CommentExtInfo struct {
	IsAlreadySendTmpl bool `json:"isAlreadySendTmpl"` // 是否已发过「差评客服会话」
}

// CommentProductInfo 评价商品信息
type CommentProductInfo struct {
	ProductList []CommentProduct `json:"productList"` // 商品列表
}

// CommentProduct 评价商品
type CommentProduct struct {
	Name   string `json:"name"`   // 商品名
	PicURL string `json:"picUrl"` // 商品图片
}

// CommentInfo 评价的信息
type CommentInfo struct {
	Content CommentItem `json:"content"` // 评价
}

// CommentProcess 进度条的信息（只有差评才会有）
type CommentProcess struct {
	ActionList []CommentAction `json:"actionList"` // 进度的具体状态,数组类型，从数组最后往前数有updateTime的就是当前状态
	CommentID  string          `json:"commentId"`  // 评价进度对应的评价id
}

// CommentAction 进度的具体状态
type CommentAction struct {
	Type       int64 `json:"type"`       // 进度的类型
	UpdateTime int64 `json:"updateTime"` // 更新时间，processInfo.actionList里从后往前数第一个有更新时间的就是当前状态
}

// CommentOldComment 旧评价的信息(只有改评的新评价才有)
type CommentOldComment struct {
	CommentID  string            `json:"commentId"`  // 评价id
	CreateTime util.FlexString   `json:"createTime"` // 创建时间，单位是秒（官方返回示例中为数字）
	Score      util.FlexInt64    `json:"score"`      // 用户评价的打分，每100对应1分以及1颗星（官方参数表为 string、返回示例为数字）
	Content    CommentOldContent `json:"content"`    // 评论的内容
}

// CommentOldContent 旧评价的内容
type CommentOldContent struct {
	Ext   string         `json:"ext"`   // 评价的文本内容
	Media []CommentMedia `json:"media"` // 评价的多媒体内容,跟上面提到的媒体结构一致
}

// CommentReplyList 评论及回复列表
type CommentReplyList struct {
	Reply            CommentReply       `json:"reply"`            // 评论
	CommentReplyList []CommentReplyItem `json:"commentReplyList"` // 回复（评论回复第二条及之后的所有集合）
}

// CommentReply 评论
type CommentReply struct {
	CommentID    string              `json:"commentId"`    // 评价id
	ReplyID      string              `json:"replyId"`      // 评论id
	CreateTime   string              `json:"createTime"`   // 创建时间，单位是秒
	UpdateTime   string              `json:"updateTime"`   // 更新时间，单位是秒
	ReplyContent CommentReplyContent `json:"replyContent"` // 回复
	ReplyObject  CommentReplyObject  `json:"replyObject"`  // 评论的内容
}

// CommentReplyItem 回复
type CommentReplyItem struct {
	CommentID           string              `json:"commentId"`           // 评价id
	CommentReplyID      string              `json:"commentReplyId"`      // 回复id
	CreateTime          string              `json:"createTime"`          // 评价创建时间，单位是秒
	UpdateTime          string              `json:"updateTime"`          // 更新时间，单位是秒
	CommentReplyContent CommentReplyContent `json:"commentReplyContent"` // 回复
	CommentReplyObject  CommentReplyObject  `json:"commentReplyObject"`  // 评论的内容
}

// CommentReplyContent 回复的内容
type CommentReplyContent struct {
	Content string `json:"content"` // 回复的内容
}

// CommentReplyObject 评论的内容
type CommentReplyObject struct {
	Nickname string `json:"nickname"` // 用户的昵称
	ImgURL   string `json:"imgUrl"`   // 用户的头像
}
