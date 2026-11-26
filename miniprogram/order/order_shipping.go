package order

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// 发货信息合单录入
	uploadCombinedShippingInfoURL = "https://api.weixin.qq.com/wxa/sec/order/upload_combined_shipping_info?access_token=%s"

	// 品牌申请
	famousBrandApplyURL = "https://api.weixin.qq.com/wxa/sec/famousbrand/apply?access_token=%s"

	// 小程序交易类型变更申请
	setWxaTradeTypeCgiURL = "https://api.weixin.qq.com/wxa/sec/order/setwxatradetypecgi?access_token=%s"

	// 小程序品牌申请状态查询
	getFamousBrandApplyStatusURL = "https://api.weixin.qq.com/wxa/sec/famousbrand/get_status?access_token=%s"

	// 查询小程序是否已完成交易结算管理确认
	isTradeManagementConfirmationCompletedURL = "https://api.weixin.qq.com/wxa/sec/order/is_trade_management_confirmation_completed?access_token=%s"

	// 查询小程序是否已开通发货信息管理服务
	isTradeManagedURL = "https://api.weixin.qq.com/wxa/sec/order/is_trade_managed?access_token=%s"

	// 消息跳转路径设置
	setMsgJumpPathURL = "https://api.weixin.qq.com/wxa/sec/order/set_msg_jump_path?access_token=%s"

	// 特殊发货报备
	opSpecialOrderURL = "https://api.weixin.qq.com/wxa/sec/order/opspecialorder?access_token=%s"
)

// UploadCombinedShippingInfo 发货信息合单录入
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/order_shipping/api_uploadcombinedshippinginfo
func (s *Shipping) UploadCombinedShippingInfo(in *UploadCombinedShippingInfoRequest) (out UploadCombinedShippingInfoResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(uploadCombinedShippingInfoURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &out, "UploadCombinedShippingInfo")
	return
}

// FamousBrandApply 品牌申请
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/order_shipping/api_famousbrandapply
func (s *Shipping) FamousBrandApply(in *FamousBrandApplyRequest) (out FamousBrandApplyResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(famousBrandApplyURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &out, "FamousBrandApply")
	return
}

// SetWxaTradeTypeCgi 小程序交易类型变更申请
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/order_shipping/api_setwxatradetypecgi
func (s *Shipping) SetWxaTradeTypeCgi(in *SetWxaTradeTypeCgiRequest) (out SetWxaTradeTypeCgiResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(setWxaTradeTypeCgiURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &out, "SetWxaTradeTypeCgi")
	return
}

// GetFamousBrandApplyStatus 小程序品牌申请状态查询
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/order_shipping/api_getfamousbrandapplystatus
func (s *Shipping) GetFamousBrandApplyStatus(in *GetFamousBrandApplyStatusRequest) (out GetFamousBrandApplyStatusResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getFamousBrandApplyStatusURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &out, "GetFamousBrandApplyStatus")
	return
}

// IsTradeManagementConfirmationCompleted 查询小程序是否已完成交易结算管理确认
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/order_shipping/api_istrademanagementconfirmationcompleted
func (s *Shipping) IsTradeManagementConfirmationCompleted(in *IsTradeManagementConfirmationCompletedRequest) (out IsTradeManagementConfirmationCompletedResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(isTradeManagementConfirmationCompletedURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &out, "IsTradeManagementConfirmationCompleted")
	return
}

// IsTradeManaged 查询小程序是否已开通发货信息管理服务
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/order_shipping/api_istrademanaged
func (s *Shipping) IsTradeManaged(in *IsTradeManagedRequest) (out IsTradeManagedResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(isTradeManagedURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &out, "IsTradeManaged")
	return
}

// SetMsgJumpPath 消息跳转路径设置
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/order_shipping/api_setmsgjumppath
func (s *Shipping) SetMsgJumpPath(in *SetMsgJumpPathRequest) (out SetMsgJumpPathResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(setMsgJumpPathURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &out, "SetMsgJumpPath")
	return
}

// OpSpecialOrder 特殊发货报备
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/order_shipping/api_opspecialorder
func (s *Shipping) OpSpecialOrder(in *OpSpecialOrderRequest) (out OpSpecialOrderResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(opSpecialOrderURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &out, "OpSpecialOrder")
	return
}

// UploadCombinedShippingInfoRequest 发货信息合单录入请求参数
type UploadCombinedShippingInfoRequest struct {
	OrderKey   *ShippingOrderKey           `json:"order_key"`   // 合单订单，需要上传物流详情的合单订单，根据订单类型二选一
	SubOrders  []*CombinedShippingSubOrder `json:"sub_orders"`  // 子单物流详情
	UploadTime string                      `json:"upload_time"` // 上传时间，用于标识请求的先后顺序，示例值: 2022-12-15T13:29:35.120+08:00
	Payer      *ShippingPayer              `json:"payer"`       // 支付者，支付者信息
}

// CombinedShippingSubOrder 子单物流详情
type CombinedShippingSubOrder struct {
	OrderKey       *ShippingOrderKey `json:"order_key"`        // 需要上传物流详情的子单订单，订单类型与合单订单保持一致
	LogisticsType  LogisticsType     `json:"logistics_type"`   // 物流模式，发货方式枚举值：1、实体物流配送；2、同城配送；3、虚拟商品；4、用户自提
	DeliveryMode   DeliveryMode      `json:"delivery_mode"`    // 发货模式，枚举值：1、UNIFIED_DELIVERY（统一发货）；2、SPLIT_DELIVERY（分拆发货）
	IsAllDelivered bool              `json:"is_all_delivered"` // 分拆发货模式时必填，用于标识分拆发货模式下是否已全部发货完成，只有全部发货完成才会向用户推送发货完成通知
	ShippingList   []*ShippingInfo   `json:"shipping_list"`    // 子单物流信息列表，多重性: [1, 15]
}

// UploadCombinedShippingInfoResponse 发货信息合单录入返回参数
type UploadCombinedShippingInfoResponse struct {
	util.CommonError
}

// FamousBrandApplyRequest 品牌申请请求参数
type FamousBrandApplyRequest struct {
	Application *FamousBrandApplication `json:"Application"` // 申请品牌信息
}

// FamousBrandApplication 申请品牌信息
type FamousBrandApplication struct {
	ApplyFor  int                   `json:"apply_for"`  // 品牌申请类型枚举值
	AuditInfo *FamousBrandAuditInfo `json:"audit_info"` // 申请品牌信息，当申请类型为知名品牌时必填
}

// FamousBrandAuditInfo 申请品牌信息
type FamousBrandAuditInfo struct {
	BrandName                   string                   `json:"brand_name"`                     // 品牌名称
	BrandType                   int                      `json:"brand_type"`                     // 品牌类型
	FlagshipInWhichEcPlatform   string                   `json:"flagship_in_which_ec_platform"`  // 旗舰店所在电商平台，例如淘宝
	EcPlatformProofList         []map[string]interface{} `json:"ec_platform_proof_list"`         // 电商平台证明材料列表，官方文档未定义元素字段
	OtherMaterialList           []map[string]interface{} `json:"other_material_list"`            // 其他补充材料列表，官方文档未定义元素字段
	AuthorityCertifiedProofList []map[string]interface{} `json:"authority_certified_proof_list"` // 品牌授权证明材料列表，官方文档未定义元素字段
}

// FamousBrandApplyResponse 品牌申请返回参数
type FamousBrandApplyResponse struct {
	util.CommonError
}

// SetWxaTradeTypeCgiRequest 小程序交易类型变更申请请求参数
type SetWxaTradeTypeCgiRequest struct {
	TradeType    int                  `json:"trade_type"`    // 申请变更后的目标交易类型
	MaterialList []*TradeTypeMaterial `json:"material_list"` // 申请材料列表，最多10个，其中视频最多3个
	Reason       string               `json:"reason"`        // 申请理由
}

// TradeTypeMaterial 交易类型变更申请材料
type TradeTypeMaterial struct {
	Type    int    `json:"type"`     // 材料类型：1-图片，2-视频
	MediaID string `json:"media_id"` // 通过新增临时素材接口上传后获取的 media_id
}

// SetWxaTradeTypeCgiResponse 小程序交易类型变更申请返回参数
type SetWxaTradeTypeCgiResponse struct {
	util.CommonError
}

// GetFamousBrandApplyStatusRequest 小程序品牌申请状态查询请求参数
type GetFamousBrandApplyStatusRequest struct {
}

// GetFamousBrandApplyStatusResponse 小程序品牌申请状态查询返回参数
type GetFamousBrandApplyStatusResponse struct {
	util.CommonError
	Progress    *FamousBrandProgress          `json:"progress"`    // 品牌申请进度
	Application *FamousBrandApplicationStatus `json:"application"` // 品牌申请信息
}

// FamousBrandProgress 品牌申请进度
type FamousBrandProgress struct {
	Status int `json:"status"` // 品牌申请状态
}

// FamousBrandApplicationStatus 品牌申请状态信息
type FamousBrandApplicationStatus struct {
	ApplyFor  int                     `json:"apply_for"`  // 品牌申请类型枚举值
	AuditInfo *FamousBrandAuditStatus `json:"audit_info"` // 品牌审核信息
	Status    int                     `json:"status"`     // 品牌申请状态
}

// FamousBrandAuditStatus 品牌审核信息
type FamousBrandAuditStatus struct {
	AuditReason string `json:"audit_reason"` // 审核原因，只有在审核驳回时才有值
}

// IsTradeManagementConfirmationCompletedRequest 查询小程序是否已完成交易结算管理确认请求参数
type IsTradeManagementConfirmationCompletedRequest struct {
	Appid string `json:"appid"` // 待查询小程序的 appid，非服务商调用时仅能查询本账号
}

// IsTradeManagementConfirmationCompletedResponse 查询小程序是否已完成交易结算管理确认返回参数
type IsTradeManagementConfirmationCompletedResponse struct {
	util.CommonError
	Completed util.FlexBool `json:"completed"` // 是否已完成交易结算管理确认（官方返回示例中为字符串）
}

// IsTradeManagedRequest 查询小程序是否已开通发货信息管理服务请求参数
type IsTradeManagedRequest struct {
	Appid string `json:"appid"` // 待查询小程序的 appid，非服务商调用时仅能查询本账号
}

// IsTradeManagedResponse 查询小程序是否已开通发货信息管理服务返回参数
type IsTradeManagedResponse struct {
	util.CommonError
	IsTradeManaged util.FlexBool `json:"is_trade_managed"` // 是否已开通发货信息管理服务（官方返回示例中为字符串）
}

// SetMsgJumpPathRequest 消息跳转路径设置请求参数
type SetMsgJumpPathRequest struct {
	Path string `json:"path"` // 商户自定义跳转路径
}

// SetMsgJumpPathResponse 消息跳转路径设置返回参数
type SetMsgJumpPathResponse struct {
	util.CommonError
}

// OpSpecialOrderRequest 特殊发货报备请求参数
type OpSpecialOrderRequest struct {
	OrderID string `json:"order_id"` // 需要特殊发货报备的订单号，可传入微信支付单号或商户单号
	Type    int    `json:"type"`     // 特殊发货报备类型，1为预售商品订单，2为测试订单
	DelayTo int64  `json:"delay_to"` // 预计发货时间的unix时间戳，type为1时必填，type为2可省略
}

// OpSpecialOrderResponse 特殊发货报备返回参数
type OpSpecialOrderResponse struct {
	util.CommonError
}
