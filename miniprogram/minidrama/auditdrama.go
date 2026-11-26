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

// SubmitModifyDramaBasicInfo 修改剧目基本信息
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/minidrama/auditdrama/api_submitmodifydramabasicinforeq
func (s *MiniDrama) SubmitModifyDramaBasicInfo(ctx context.Context, in *SubmitModifyDramaBasicInfoRequest) (out SubmitModifyDramaBasicInfoResponse, err error) {
	var address string
	if address, err = s.requestAddress(ctx, submitModifyDramaBasicInfo); err != nil {
		return
	}

	var response []byte
	if response, err = util.PostJSONContext(ctx, address, in); err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "SubmitModifyDramaBasicInfo")
	return
}

// SubmitReplaceMedias 提交替换剧集审核
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/minidrama/auditdrama/api_submitreplacedramamedias
func (s *MiniDrama) SubmitReplaceMedias(ctx context.Context, in *SubmitReplaceMediasRequest) (out SubmitReplaceMediasResponse, err error) {
	var address string
	if address, err = s.requestAddress(ctx, submitReplaceMedias); err != nil {
		return
	}

	var response []byte
	if response, err = util.PostJSONContext(ctx, address, in); err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "SubmitReplaceMedias")
	return
}

// ReplaceDramaMedia 替换审核通过的剧集
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/minidrama/auditdrama/api_replacedramamedia
func (s *MiniDrama) ReplaceDramaMedia(ctx context.Context, in *ReplaceDramaMediaRequest) (out ReplaceDramaMediaResponse, err error) {
	var address string
	if address, err = s.requestAddress(ctx, replaceDramaMedia); err != nil {
		return
	}

	var response []byte
	if response, err = util.PostJSONContext(ctx, address, in); err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "ReplaceDramaMedia")
	return
}

// GetDramaLatestAuditInfo 查询剧目审核信息
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/minidrama/auditdrama/api_getdramalatestauditinfo
func (s *MiniDrama) GetDramaLatestAuditInfo(ctx context.Context, in *GetDramaLatestAuditInfoRequest) (out GetDramaLatestAuditInfoResponse, err error) {
	var address string
	if address, err = s.requestAddress(ctx, getDramaLatestAuditInfo); err != nil {
		return
	}

	var response []byte
	if response, err = util.PostJSONContext(ctx, address, in); err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetDramaLatestAuditInfo")
	return
}

// SubmitModifyDramaBasicInfoRequest 修改剧目基本信息请求
type SubmitModifyDramaBasicInfoRequest struct {
	DramaID                            int64           `json:"drama_id"`                                        // 剧目id
	Description                        string          `json:"description,omitempty"`                           // 剧目简介，支持填写200个字符。
	CoverMaterialID                    string          `json:"cover_material_id,omitempty"`                     // 剧目海报临时material_id。
	Recommendations                    string          `json:"recommendations,omitempty"`                       // 剧目推荐语。
	PromotionPosterMaterialID          string          `json:"promotion_poster_material_id,omitempty"`          // 推广海报临时material_id。
	AlternateName                      string          `json:"alternate_name,omitempty"`                        // 备用剧名。
	ActorList                          *DramaActorList `json:"actor_list,omitempty"`                            // 演员信息。如果需要修改，请填写所有的演员信息。
	QualificationType                  int             `json:"qualification_type,omitempty"`                    // 剧目资质：1-取得《网络剧片发行许可证》或重点节目备案号；2-未取得《网络剧片发行许可证》或重点节目备案，且制作成本小于30万元
	RegistrationNumber                 string          `json:"registration_number,omitempty"`                   // 剧目备案号，当qualification_type=1时必填。
	QualificationCertificateMaterialID string          `json:"qualification_certificate_material_id,omitempty"` // 剧目资质证明文件，当qualification_type=1时必填。
	CostOfProduction                   int64           `json:"cost_of_production,omitempty"`                    // 剧目制作成本（单位：万元），当qualification_type=2时必填。
	CostCommitmentLetterMaterialID     string          `json:"cost_commitment_letter_material_id,omitempty"`    // 《成本配置比例情况报告》material_id，当qualification_type=2时必填。
	OtherMaterialMaterialID            string          `json:"other_material_material_id,omitempty"`            // 其他材料material_id，如涉及互动短剧，请上传完整剧情走线示意图。
	Producer                           string          `json:"producer,omitempty"`                              // 剧目制作方。apply_for_copyright_protection =1时必填，否则无需填写。
	Copyright                          *DramaCopyright `json:"copyright,omitempty"`                             // 版权保护相关。apply_for_copyright_protection =1时必填，否则无需填写 。
}

// SubmitModifyDramaBasicInfoResponse 修改剧目基本信息响应
type SubmitModifyDramaBasicInfoResponse struct {
	util.CommonError
}

// DramaActorList 演员信息
type DramaActorList struct {
	Actor []*DramaActor `json:"actor"` // 演员列表
}

// DramaActor 演员信息
type DramaActor struct {
	Name            string `json:"name"`              // 演员姓名，格式要求：可填写30个字符，支持输入中文、英文和·
	PhotoMaterialID string `json:"photo_material_id"` // 演员照片临时material_id
	Role            string `json:"role"`              // 饰演角色，格式要求：可填写30个字符，支持输入中文、英文及，。；”“？、—《》！（）-+·
	Profile         string `json:"profile"`           // 演员简介，填写内容可参考：生日、星座、籍贯、身高、毕业院校、历史/代表作品、获奖信息、成就等(以上要素需至少满足3项或涵盖任一要素且不少于20字）。
}

// DramaCopyright 版权保护相关
type DramaCopyright struct {
	CopyrightRole                               int      `json:"copyright_role"`                                            // 提审主体身份，1. 剧目制作方; 2. 授权播出方或版权方。
	ApplyForCopyrightProtection                 int      `json:"apply_for_copyright_protection"`                            // 是否申请版权保护，0. 不申请; 1. 申请。不申请时必须上传权利声明/播放授权材料。
	CopyrightVerification                       string   `json:"copyright_verification,omitempty"`                          // 版权验证方式：1. 基于版权证明材料; 2. 基于版权授权关系。
	ProofOfProduction                           []string `json:"proof_of_production,omitempty"`                             // 剧目制作证明材料临时 material_id，apply_for_copyright_protection=1时必填，最多支持上传10个。
	PurchaseOrBroadcastAuthorizationCertificate []string `json:"purchase_or_broadcast_authorization_certificate,omitempty"` // 版权采买/播出授权证明材料临时 material_id，最多支持上传 4 个。
}

// SubmitReplaceMediasRequest 提交替换剧集审核请求
type SubmitReplaceMediasRequest struct {
	DramaID          int64          `json:"drama_id"`           // 剧目id
	ReplaceMediaList []*ReplaceInfo `json:"replace_media_list"` // 替换的剧集信息。
}

// SubmitReplaceMediasResponse 提交替换剧集审核响应
type SubmitReplaceMediasResponse struct {
	util.CommonError
}

// ReplaceDramaMediaRequest 替换审核通过的剧集请求
type ReplaceDramaMediaRequest struct {
	DramaID    int64 `json:"drama_id"`     // 剧目id
	OldMediaID int64 `json:"old_media_id"` // 旧剧集meida_id
	NewMediaID int64 `json:"new_media_id"` // 新剧集meida_id
}

// ReplaceDramaMediaResponse 替换审核通过的剧集响应
type ReplaceDramaMediaResponse struct {
	util.CommonError
}

// GetDramaLatestAuditInfoRequest 查询剧目审核信息请求
type GetDramaLatestAuditInfoRequest struct {
	DramaID   int64 `json:"drama_id"`   // 剧目id
	AuditType int   `json:"audit_type"` // 审核类型。0为首次提审；1为再次提审；2为替换剧集提审；3为修改剧目基本信息审核。
}

// GetDramaLatestAuditInfoResponse 查询剧目审核信息响应
type GetDramaLatestAuditInfoResponse struct {
	util.CommonError
	AuditDetail *DramaLatestAuditDetail `json:"audit_detail"` // 审核详情
}

// DramaLatestAuditDetail 剧目最新审核详情
type DramaLatestAuditDetail struct {
	Status     int   `json:"status"`      // 0为无效值；1为审核中；2为最终失败；3为审核通过；4为驳回重填
	AuditType  int   `json:"audit_type"`  // 事件通知接口才返回：0为首次提审；1为再次提审；2为替换剧集提审；3为修改剧目基本信息审核。
	CreateTime int64 `json:"create_time"` // 提审时间戳。
	AuditTime  int64 `json:"audit_time"`  // 审核时间戳。
}
