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
	// retailUploadMchFileURL 上传商户图片
	retailUploadMchFileURL = "https://api.weixin.qq.com/retail/B2b/retailuploadmchfile?access_token=%s"
	// closeB2BOrderURL 关闭订单
	closeB2BOrderURL = "https://api.weixin.qq.com/retail/B2b/closeb2border?access_token=%s&pay_sig=%s"
	// delProfitSharingAccountURL 删除分账方
	delProfitSharingAccountURL = "https://api.weixin.qq.com/retail/B2b/delprofitsharingaccount?access_token=%s&pay_sig=%s"
	// manualWithdrawURL 发起手动提现
	manualWithdrawURL = "https://api.weixin.qq.com/retail/B2b/withdraw?access_token=%s&pay_sig=%s"
	// retailRegisterMchURL 商户号进件
	retailRegisterMchURL = "https://api.weixin.qq.com/retail/B2b/retailregistermch?access_token=%s"
	// finishProfitSharingOrderURL 完成分账
	finishProfitSharingOrderURL = "https://api.weixin.qq.com/retail/B2b/finishprofitsharingorder?access_token=%s&pay_sig=%s"
	// setAutoWithdrawURL 微信支付自动提现
	setAutoWithdrawURL = "https://api.weixin.qq.com/retail/B2b/setautowithdraw?access_token=%s&pay_sig=%s"
	// setMchProfitRateURL 报名微信支付技术服务费优惠活动
	setMchProfitRateURL = "https://api.weixin.qq.com/retail/B2b/setmchprofitrate?access_token=%s"
	// updateWqfChargeFeeURL 报名银行转账技术服务费优惠活动
	updateWqfChargeFeeURL = "https://api.weixin.qq.com/retail/B2b/updatewqfchargefee?access_token=%s"
	// downLoadBillURL 下载交易账单与资金账单
	downLoadBillURL = "https://api.weixin.qq.com/retail/B2b/downloadbill?access_token=%s&pay_sig=%s"
	// queryProfitSharingRemainAmtURL 查询分账剩余金额
	queryProfitSharingRemainAmtURL = "https://api.weixin.qq.com/retail/B2b/queryprofitsharingremainamt?access_token=%s&pay_sig=%s"
	// queryRefundProfitSharingOrderURL 查询分账回退结果
	queryRefundProfitSharingOrderURL = "https://api.weixin.qq.com/retail/B2b/queryrefundprofitsharingorder?access_token=%s&pay_sig=%s"
	// queryProfitSharingAccountURL 查询分账方
	queryProfitSharingAccountURL = "https://api.weixin.qq.com/retail/B2b/queryprofitsharingaccount?access_token=%s&pay_sig=%s"
	// queryProfitSharingOrderURL 查询分账结果
	queryProfitSharingOrderURL = "https://api.weixin.qq.com/retail/B2b/queryprofitsharingorder?access_token=%s&pay_sig=%s"
	// retailGetMchOrderURL 查询商户号开通状态
	retailGetMchOrderURL = "https://api.weixin.qq.com/retail/B2b/retailgetmchorder?access_token=%s"
	// queryWithdrawURL 查询提现状态
	queryWithdrawURL = "https://api.weixin.qq.com/retail/B2b/querywithdraw?access_token=%s&pay_sig=%s"
	// getOrderURL 查询订单
	getOrderURL = "https://api.weixin.qq.com/retail/B2b/getorder?access_token=%s&pay_sig=%s"
	// getMchBalanceURL 查询账户余额
	getMchBalanceURL = "https://api.weixin.qq.com/retail/B2b/getmchbalance?access_token=%s&pay_sig=%s"
	// getRefundURL 查询退款
	getRefundURL = "https://api.weixin.qq.com/retail/B2b/getrefund?access_token=%s&pay_sig=%s"
	// getWqfChargeFeeURL 查询银行转账的技术服务费率
	getWqfChargeFeeURL = "https://api.weixin.qq.com/retail/B2b/getwqfchargefee?access_token=%s"
	// addProfitSharingAccountURL 添加分账方
	addProfitSharingAccountURL = "https://api.weixin.qq.com/retail/B2b/addprofitsharingaccount?access_token=%s&pay_sig=%s"
	// registerOnlyWqfURL 申请开通银行转账
	registerOnlyWqfURL = "https://api.weixin.qq.com/retail/B2b/registeronlywqf?access_token=%s"
	// getAppKeyURL 获取密钥 AppKey
	getAppKeyURL = "https://api.weixin.qq.com/retail/B2b/getappkey?access_token=%s"
	// getMchInfoURL 获取小程序下所有商户的信息
	getMchInfoURL = "https://api.weixin.qq.com/retail/B2b/getmchinfo?access_token=%s"
	// createProfitSharingOrderURL 请求分账
	createProfitSharingOrderURL = "https://api.weixin.qq.com/retail/B2b/createprofitsharingorder?access_token=%s&pay_sig=%s"
	// refundProfitSharingURL 请求分账回退
	refundProfitSharingURL = "https://api.weixin.qq.com/retail/B2b/refundprofitsharing?access_token=%s&pay_sig=%s"
	// createWqfLinkURL 跳转银行转账页面
	createWqfLinkURL = "https://api.weixin.qq.com/retail/B2b/createwqflink?access_token=%s"
	// refundOrderURL 退款
	refundOrderURL = "https://api.weixin.qq.com/retail/B2b/refund?access_token=%s&pay_sig=%s"
)

// RetailUploadMchFileRequest 上传商户图片请求
type RetailUploadMchFileRequest struct {
	FileName string `json:"file_name"` // 文件名
	File     string `json:"file"`      // 文件二进制流的 Base64 编码
}

// RetailUploadMchFileResponse 上传商户图片响应
type RetailUploadMchFileResponse struct {
	util.CommonError
	FileID string `json:"file_id"` // 文件 id
}

// RetailUploadMchFile 上传商户图片
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/bill/api_retailuploadmchfile
func (s *B2B) RetailUploadMchFile(in *RetailUploadMchFileRequest) (out RetailUploadMchFileResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(retailUploadMchFileURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "RetailUploadMchFile")
	return
}

// CloseB2BOrderRequest 关闭订单请求
type CloseB2BOrderRequest struct {
	PaySig     string `json:"-"`            // pay_sig，请求签名，通过查询参数传递
	MchID      string `json:"mchid"`        // 微信商户号，长度限制[1,32]
	OutTradeNo string `json:"out_trade_no"` // 商户订单号，长度限制[6,32]，与 order_id 必填其一
	OrderID    string `json:"order_id"`     // B2b 支付订单号，长度限制[1,32]，与 out_trade_no 必填其一
}

// CloseB2BOrderResponse 关闭订单响应
type CloseB2BOrderResponse struct {
	util.CommonError
}

// CloseB2BOrder 关闭订单
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/bill/api_closeb2border
func (s *B2B) CloseB2BOrder(in *CloseB2BOrderRequest) (out CloseB2BOrderResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(closeB2BOrderURL, accessToken, in.PaySig)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "CloseB2BOrder")
	return
}

// DelProfitSharingAccountRequest 删除分账方请求
type DelProfitSharingAccountRequest struct {
	PaySig    string `json:"-"`          // pay_sig，请求签名，通过查询参数传递
	PayeeType string `json:"payee_type"` // 分账接收方类型
	PayeeID   string `json:"payee_id"`   // 分账接收方标识，payee_type 为外部用户时填 openid，为外部商户时填商户号
}

// DelProfitSharingAccountResponse 删除分账方响应
type DelProfitSharingAccountResponse struct {
	util.CommonError
}

// DelProfitSharingAccount 删除分账方
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/bill/api_delprofitsharingaccount
func (s *B2B) DelProfitSharingAccount(in *DelProfitSharingAccountRequest) (out DelProfitSharingAccountResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(delProfitSharingAccountURL, accessToken, in.PaySig)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "DelProfitSharingAccount")
	return
}

// ManualWithdrawRequest 发起手动提现请求
type ManualWithdrawRequest struct {
	PaySig         string `json:"-"`               // pay_sig，请求签名，通过查询参数传递
	MchID          string `json:"mchid"`           // 微信商户号
	WithdrawAmount int64  `json:"withdraw_amount"` // 提现金额，单位：分
	OutWithdrawNo  string `json:"out_withdraw_no"` // 商户提现单号
}

// ManualWithdrawResponse 发起手动提现响应
type ManualWithdrawResponse struct {
	util.CommonError
}

// ManualWithdraw 发起手动提现
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/bill/api_manualwithdraw
func (s *B2B) ManualWithdraw(in *ManualWithdrawRequest) (out ManualWithdrawResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(manualWithdrawURL, accessToken, in.PaySig)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "ManualWithdraw")
	return
}

// RetailRegisterMchRequest 商户号进件请求
type RetailRegisterMchRequest struct {
	IDDocTypeNum         int64            `json:"id_doc_type_num"`        // 法人证件类型，0/1 大陆身份证，2 护照，3 港澳通行证，4 澳门通行证，5 台湾通行证，6 外国人居留证，7 港澳居民证，8 台湾居民证
	IDCardInfo           *IDCardInfo      `json:"id_card_info"`           // 经营者/法人身份证信息，id_doc_type_num 为 0 或 1 时必填
	IDDocInfo            *IDDocInfo       `json:"id_doc_info"`            // 经营者/法人其他类型证件信息，id_doc_type_num 不为 0 或 1 时必填
	AccountInfo          *AccountInfo     `json:"account_info"`           // 结算银行账户
	ContactInfo          *ContactInfo     `json:"contact_info"`           // 超级管理员信息
	BusinessLicense      *BusinessLicense `json:"business_license"`       // 营业执照
	MerchantShortname    string           `json:"merchant_shortname"`     // 商户名缩写
	OrganizationType     bool             `json:"organization_type"`      // 主体类型，个体户-0，企业-1
	Qualification        *Qualification   `json:"qualification"`          // 行业特殊资质资料
	BusinessAdditionDesc string           `json:"business_addition_desc"` // 补充说明
	BusinessAdditionPics string           `json:"business_addition_pics"` // 补充材料，图片 id 通过上传商户图片接口获取
	OpenType             int64            `json:"open_type"`              // 开通支付方式，只开通微信支付-0，同时开通微信支付和银行转账-1
	ExtRegisterInfo      *ExtRegisterInfo `json:"ext_register_info"`      // 补充信息
	ClientIP             string           `json:"client_ip"`              // 商户 ip 地址，支持 IPv4 和 IPv6
}

// IDCardInfo 经营者/法人身份证信息
type IDCardInfo struct {
	IDCardCopy           string `json:"id_card_copy"`             // 身份证人像面照片 id，通过上传商户图片接口获取
	IDCardNational       string `json:"id_card_national"`         // 身份证国徽面照片 id，通过上传商户图片接口获取
	IDCardName           string `json:"id_card_name"`             // 身份证姓名
	IDCardNumber         string `json:"id_card_number"`           // 身份证号码
	IDCardValidTime      string `json:"id_card_valid_time"`       // 身份证有效期限，格式如 "2026-06-06"、"长期"
	IDCardAddress        string `json:"id_card_address"`          // 身份证地址
	IDCardValidTimeBegin string `json:"id_card_valid_time_begin"` // 身份证有效期开始日期，格式如 "2026-06-06"
}

// IDDocInfo 经营者/法人其他类型证件信息
type IDDocInfo struct {
	IDDocName      string `json:"id_doc_name"`      // 证件姓名
	IDDocNumber    string `json:"id_doc_number"`    // 证件号码
	IDDocCopy      string `json:"id_doc_copy"`      // 证件正面照片，通过上传商户图片接口获取
	DocPeriodEnd   string `json:"doc_period_end"`   // 证件结束日期，格式如 "2022-06-06"、"长期"
	DocPeriodBegin string `json:"doc_period_begin"` // 证件有效期开始时间，格式如 "2022-06-06"
	IDDocAddress   string `json:"id_doc_address"`   // 证件居住地址
	IDDocCopyBack  string `json:"id_doc_copy_back"` // 证件反面照片，通过上传商户图片接口获取
}

// AccountInfo 结算银行账户
type AccountInfo struct {
	BankAccountType string `json:"bank_account_type"` // 账户类型，对公账户-"74"，对私账户-"75"
	AccountBank     string `json:"account_bank"`      // 开户银行，比如"工商银行"
	AccountName     string `json:"account_name"`      // 开户名称
	BankAddressCode string `json:"bank_address_code"` // 开户银行省市编码，例如 "110000"
	BankBranchID    string `json:"bank_branch_id"`    // 开户银行联行号，与开户银行全称二选一
	BankName        string `json:"bank_name"`         // 开户银行全称（含支行）
	AccountNumber   string `json:"account_number"`    // 银行帐号
}

// ContactInfo 超级管理员信息
type ContactInfo struct {
	ContactType                 string `json:"contact_type"`                  // 超级管理员类型，"65" 经营者/法人，"66" 经办人
	ContactName                 string `json:"contact_name"`                  // 超级管理员姓名
	ContactIDDocType            string `json:"contact_id_doc_type"`           // 超级管理员证件类型
	ContactIDCardNumber         string `json:"contact_id_card_number"`        // 超级管理员身份证件号码
	ContactIDDocCopy            string `json:"contact_id_doc_copy"`           // 超级管理员证件正面照片 id
	ContactIDDocCopyBack        string `json:"contact_id_doc_copy_back"`      // 超级管理员证件反面照片 id
	ContactIDDocPeriodBegin     string `json:"contact_id_doc_period_begin"`   // 超级管理员证件有效期开始时间
	ContactIDDocPeriodEnd       string `json:"contact_id_doc_period_end"`     // 超级管理员证件有效期结束时间
	BusinessAuthorizationLetter string `json:"business_authorization_letter"` // 业务办理授权函
	MobilePhone                 string `json:"mobile_phone"`                  // 超级管理员手机
	ContactEmail                string `json:"contact_email"`                 // 超级管理员邮箱
}

// BusinessLicense 营业执照
type BusinessLicense struct {
	BusinessLicenseCopy   string `json:"business_license_copy"`   // 证件扫描件图片 id，通过上传商户图片接口获取
	BusinessLicenseNumber string `json:"business_license_number"` // 证件注册号
	MerchantName          string `json:"merchant_name"`           // 商户名称
	LegalPerson           string `json:"legal_person"`            // 经营者/法定代表人姓名
	CompanyAddress        string `json:"company_address"`         // 注册地址
	BusinessTime          string `json:"business_time"`           // 营业期限
	CertType              string `json:"cert_type"`               // 登记证书类型
}

// Qualification 行业特殊资质资料
type Qualification struct {
	QualificationType string `json:"qualification_type"` // 行业特殊资质类型，如"速送"
	Qualifications    string `json:"qualifications"`     // 行业特殊资料，字符串数组构成的字符串
}

// ExtRegisterInfo 补充信息
type ExtRegisterInfo struct {
	DoorHeadFileID            string `json:"door_head_file_id"`            // 企业门头照 id
	StoreFileID               string `json:"store_file_id"`                // 商城截图 id
	OnlinePayFileID           string `json:"online_pay_file_id"`           // 确认订单付款界面截图 id
	MerchantScale             string `json:"merchant_scale"`               // 企业规模，枚举值
	AuthorizationLetterFileID string `json:"authorization_letter_file_id"` // 银行转账授权书图片 id
}

// RetailRegisterMchResponse 商户号进件响应
type RetailRegisterMchResponse struct {
	util.CommonError
	OrderNo string `json:"order_no"` // 进件申请单号，可用于查询进件状态
}

// RetailRegisterMch 商户号进件
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/bill/api_retailregistermch
func (s *B2B) RetailRegisterMch(in *RetailRegisterMchRequest) (out RetailRegisterMchResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(retailRegisterMchURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "RetailRegisterMch")
	return
}

// FinishProfitSharingOrderRequest 完成分账请求
type FinishProfitSharingOrderRequest struct {
	PaySig     string `json:"-"`            // pay_sig，请求签名，通过查询参数传递
	MchID      string `json:"mchid"`        // 子商户 id，在 B2b 小程序中进件完成返回的子商户 id
	OutTradeNo string `json:"out_trade_no"` // 订单 id，在 B2b 小程序中下单的订单 id
}

// FinishProfitSharingOrderResponse 完成分账响应
type FinishProfitSharingOrderResponse struct {
	util.CommonError
}

// FinishProfitSharingOrder 完成分账
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/bill/api_finishprofitsharingorder
func (s *B2B) FinishProfitSharingOrder(in *FinishProfitSharingOrderRequest) (out FinishProfitSharingOrderResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(finishProfitSharingOrderURL, accessToken, in.PaySig)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "FinishProfitSharingOrder")
	return
}

// SetAutoWithdrawRequest 微信支付自动提现请求
type SetAutoWithdrawRequest struct {
	PaySig    string `json:"-"`          // pay_sig，请求签名，通过查询参数传递
	MchID     string `json:"mchid"`      // 微信商户号
	Status    int64  `json:"status"`     // 自动提现状态
	RetainAmt int64  `json:"retain_amt"` // 留存额，单位：分
}

// SetAutoWithdrawResponse 微信支付自动提现响应
type SetAutoWithdrawResponse struct {
	util.CommonError
}

// SetAutoWithdraw 微信支付自动提现接口
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/bill/api_setautowithdraw
func (s *B2B) SetAutoWithdraw(in *SetAutoWithdrawRequest) (out SetAutoWithdrawResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(setAutoWithdrawURL, accessToken, in.PaySig)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "SetAutoWithdraw")
	return
}

// SetMchProfitRateRequest 报名微信支付技术服务费优惠活动请求
type SetMchProfitRateRequest struct {
	SubMchID   string `json:"sub_mchid"`   // 商户号
	ProfitRate int64  `json:"profit_rate"` // 费率，万分比，比如 40 指的是 0.40%
}

// SetMchProfitRateResponse 报名微信支付技术服务费优惠活动响应
type SetMchProfitRateResponse struct {
	util.CommonError
}

// SetMchProfitRate 报名微信支付技术服务费优惠活动
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/bill/api_setmchprofitrate
func (s *B2B) SetMchProfitRate(in *SetMchProfitRateRequest) (out SetMchProfitRateResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(setMchProfitRateURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "SetMchProfitRate")
	return
}

// UpdateWqfChargeFeeRequest 报名银行转账技术服务费优惠活动请求
type UpdateWqfChargeFeeRequest struct {
	SubMchID                    string `json:"sub_mchid"`                      // 商户号
	CertifiedChargeFeeNumerator int64  `json:"certified_charge_fee_numerator"` // 认证费率分子，分母统一为 10000，取值范围 [22,40]
}

// UpdateWqfChargeFeeResponse 报名银行转账技术服务费优惠活动响应
type UpdateWqfChargeFeeResponse struct {
	util.CommonError
	MerchantName                  string `json:"merchant_name"`                    // 商户名称
	CertifiedChargeFeeNumerator   int64  `json:"certified_charge_fee_numerator"`   // 认证费率分子，分母统一为 10000
	UncertifiedChargeFeeNumerator int64  `json:"uncertified_charge_fee_numerator"` // 未认证的费率分子，分母统一为 10000
}

// UpdateWqfChargeFee 报名银行转账技术服务费优惠活动
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/bill/api_updatewqfchargefee
func (s *B2B) UpdateWqfChargeFee(in *UpdateWqfChargeFeeRequest) (out UpdateWqfChargeFeeResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(updateWqfChargeFeeURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "UpdateWqfChargeFee")
	return
}

// DownLoadBillRequest 下载交易账单与资金账单请求
type DownLoadBillRequest struct {
	PaySig   string `json:"-"`         // pay_sig，请求签名，通过查询参数传递
	MchID    string `json:"mchid"`     // 微信商户号
	BillDate string `json:"bill_date"` // 账单日期，格式如 "20231102"
}

// DownLoadBillResponse 下载交易账单与资金账单响应
type DownLoadBillResponse struct {
	util.CommonError
	SuccessBillURL       string `json:"success_bill_url"`        // 微信支付渠道支付成功订单下载链接
	RefundBillURL        string `json:"refund_bill_url"`         // 微信支付渠道退款单下载链接
	AllBillURL           string `json:"all_bill_url"`            // 支付成功和退款单的下载链接
	FundBillURL          string `json:"fund_bill_url"`           // 微信支付渠道的资金账单下载链接
	EndedDayAvailAmt     int64  `json:"ended_day_avail_amt"`     // 日终账户可提现金额
	EndedDayFrozenAmt    int64  `json:"ended_day_frozen_amt"`    // 日终账户待结算金额
	EndedDayTotalAmt     int64  `json:"ended_day_total_amt"`     // 日终账户总金额，等于可提现金额+待结算金额
	ProfitSharingBillURL string `json:"profit_sharing_bill_url"` // 分账成功订单下载链接
	ProfitRefundBillURL  string `json:"profit_refund_bill_url"`  // 分账回退单下载链接
	BankpayFundBillURL   string `json:"bankpay_fund_bill_url"`   // 银行转账渠道资金帐单下载链接
}

// DownLoadBill 接口下载交易账单与资金账单
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/bill/api_downloadbill
func (s *B2B) DownLoadBill(in *DownLoadBillRequest) (out DownLoadBillResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(downLoadBillURL, accessToken, in.PaySig)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "DownLoadBill")
	return
}

// QueryProfitSharingRemainAmtRequest 查询分账剩余金额请求
type QueryProfitSharingRemainAmtRequest struct {
	PaySig     string `json:"-"`            // pay_sig，请求签名，通过查询参数传递
	MchID      string `json:"mchid"`        // 子商户 id，在 B2b 小程序中进件完成返回的子商户 id
	OutTradeNo string `json:"out_trade_no"` // 订单 id，在 B2b 小程序中下单的订单 id
}

// QueryProfitSharingRemainAmtResponse 查询分账剩余金额响应
type QueryProfitSharingRemainAmtResponse struct {
	util.CommonError
	RemainAmt int64 `json:"remain_amt"` // 冻结金额，单位：分
}

// QueryProfitSharingRemainAmt 查询分账剩余金额
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/bill/api_queryprofitsharingremainamt
func (s *B2B) QueryProfitSharingRemainAmt(in *QueryProfitSharingRemainAmtRequest) (out QueryProfitSharingRemainAmtResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(queryProfitSharingRemainAmtURL, accessToken, in.PaySig)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "QueryProfitSharingRemainAmt")
	return
}

// QueryRefundProfitSharingOrderRequest 查询分账回退结果请求
type QueryRefundProfitSharingOrderRequest struct {
	PaySig      string `json:"-"`             // pay_sig，请求签名，通过查询参数传递
	OutTradeNo  string `json:"out_trade_no"`  // 订单 id，在 B2b 小程序中下单后返回的订单 id
	OutRefundNo string `json:"out_refund_no"` // 退款单 id，在 B2b 小程序中下单后返回的退款单 id
	MchID       string `json:"mchid"`         // 商户号 id，发起这笔订单的商户号
	PayeeType   string `json:"payee_type"`    // 退款分账方类型，同添加分账方时填入的内容
	PayeeID     string `json:"payee_id"`      // 退款分账方 id，同添加分账方时填入的内容
}

// QueryRefundProfitSharingOrderResponse 查询分账回退结果响应
type QueryRefundProfitSharingOrderResponse struct {
	util.CommonError
	OrderStatus int64 `json:"order_status"` // 订单状态
}

// QueryRefundProfitSharingOrder 查询分账回退结果
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/bill/api_queryrefundprofitsharingorder
func (s *B2B) QueryRefundProfitSharingOrder(in *QueryRefundProfitSharingOrderRequest) (out QueryRefundProfitSharingOrderResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(queryRefundProfitSharingOrderURL, accessToken, in.PaySig)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "QueryRefundProfitSharingOrder")
	return
}

// QueryProfitSharingAccountRequest 查询分账方请求
type QueryProfitSharingAccountRequest struct {
	PaySig string `json:"-"`      // pay_sig，请求签名，通过查询参数传递
	Offset int64  `json:"offset"` // 分账方数据起始位置的偏移量，默认值为 0
	Limit  int64  `json:"limit"`  // 返回分账方数据的最大条数，默认值为 10
}

// ProfitSharingAccount 分账方信息
type ProfitSharingAccount struct {
	SharingAccountType string `json:"sharing_account_type"` // 分账方类型，"PAYEE_TYPE_EXTERNAL_USER" 外部用户，"PAYEE_TYPE_EXTERNAL_MERCHANT" 外部商户
	SharingAccount     string `json:"sharing_account"`      // 外部用户时为个人 openid，外部商户时为商户号
	AddTime            int64  `json:"add_time"`             // 添加时间
	UpdateTime         int64  `json:"update_time"`          // 更新时间
	Name               string `json:"name"`                 // 外部商户时返回对应的商户名称
}

// QueryProfitSharingAccountResponse 查询分账方响应
type QueryProfitSharingAccountResponse struct {
	util.CommonError
	AccountList []ProfitSharingAccount `json:"account_list"` // 分账方列表
}

// QueryProfitSharingAccount 查询分账方
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/bill/api_queryprofitsharingaccount
func (s *B2B) QueryProfitSharingAccount(in *QueryProfitSharingAccountRequest) (out QueryProfitSharingAccountResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(queryProfitSharingAccountURL, accessToken, in.PaySig)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "QueryProfitSharingAccount")
	return
}

// QueryProfitSharingOrderRequest 查询分账结果请求
type QueryProfitSharingOrderRequest struct {
	PaySig          string `json:"-"`                // pay_sig，请求签名，通过查询参数传递
	OutTradeNo      string `json:"out_trade_no"`     // 支付单 id，在 B2b 小程序中下单的订单 id
	ReceiverType    string `json:"receiver_type"`    // 分账接收方类型，同添加分账方时填入的内容
	ReceiverAccount string `json:"receiver_account"` // 分账接收方账号，同添加分账方时填入的内容
	MchID           string `json:"mchid"`            // 商户号
}

// QueryProfitSharingOrderResponse 查询分账结果响应
type QueryProfitSharingOrderResponse struct {
	util.CommonError
	OrderStatus int64 `json:"order_status"` // 分帐状态
}

// QueryProfitSharingOrder 查询分账结果
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/bill/api_queryprofitsharingorder
func (s *B2B) QueryProfitSharingOrder(in *QueryProfitSharingOrderRequest) (out QueryProfitSharingOrderResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(queryProfitSharingOrderURL, accessToken, in.PaySig)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "QueryProfitSharingOrder")
	return
}

// RetailGetMchOrderRequest 查询商户号开通状态请求
type RetailGetMchOrderRequest struct {
	OutRegistrationID string `json:"out_registration_id"` // 订单号，不填则表示拉出当前小程序的所有进件单
	PageIndex         int64  `json:"page_index"`          // 分页拉取偏移
	PageSize          int64  `json:"page_size"`           // 分页拉取总量限制
}

// RetailMchOrder 进件订单信息
type RetailMchOrder struct {
	Status               int64                 `json:"status"`                 // 0 初始化 1 资料校验中 2 待账户验证 3 审核中 4 已驳回 5 待签约 6 完成 7 已冻结 8 已作废 9 完成前平台额外准备
	InnerResp            *InnerResp            `json:"inner_resp"`             // 进件申请信息
	WqfRegisterStatement *WqfRegisterStatement `json:"wqf_register_statement"` // 银行转账开通状态
	WxPayRate            int64                 `json:"wx_pay_rate"`            // 微信支付技术服务费率，万分比
	WqfCertifiedRate     int64                 `json:"wqf_certified_rate"`     // 银行转账技术服务费率，万分比
	BindSceneStatus      int64                 `json:"bind_scene_status"`      // 与小程序关联状态
}

// InnerResp 进件申请信息
type InnerResp struct {
	SubMerchantRegistrationStatus *MchRegistrationStatus `json:"sub_merchant_registration_status"` // 二级商户申请状态
}

// MchRegistrationStatus 进件申请状态
type MchRegistrationStatus struct {
	ApplymentState     string             `json:"applyment_state"`      // 申请状态，CHECKING 资料校验中，ACCOUNT_NEED_VERIFY 待账户验证，AUDITING 审核中，REJECTED 已驳回，NEED_SIGN 待签约，FINISH 完成，FROZEN 已冻结，CANCELED 已作废
	ApplymentStateDesc string             `json:"applyment_state_desc"` // 申请状态描述
	SignState          string             `json:"sign_state"`           // 签约状态，UNSIGNED 未签约，SIGNED 已签约，NOT_SIGNABLE 不可签约
	SignURL            string             `json:"sign_url"`             // 签约链接
	SubMchID           string             `json:"sub_mchid"`            // 电商平台二级商户号
	AccountValidation  *AccountValidation `json:"account_validation"`   // 汇款账户验证信息
	AuditDetail        []AuditDetail      `json:"audit_detail"`         // 驳回原因详情
	LegalValidationURL string             `json:"legal_validation_url"` // 法人验证链接
}

// AccountValidation 汇款账户验证信息
type AccountValidation struct {
	AccountName              string `json:"account_name"`               // 账户名称
	AccountNo                string `json:"account_no"`                 // 账号
	PayAmount                int64  `json:"pay_amount"`                 // 支付金额
	DestinationAccountNumber string `json:"destination_account_number"` // 目的账号
	DestinationAccountName   string `json:"destination_account_name"`   // 目的账户名
	DestinationAccountBank   string `json:"destination_account_bank"`   // 目的银行
	City                     string `json:"city"`                       // 城市
	Remark                   string `json:"remark"`                     // 标志
	Deadline                 string `json:"deadline"`                   // 截止日期
}

// AuditDetail 驳回原因详情
type AuditDetail struct {
	ParamName    string `json:"param_name"`    // 参数名称
	RejectReason string `json:"reject_reason"` // 驳回理由
}

// WqfRegisterStatement 银行转账开通状态
type WqfRegisterStatement struct {
	WqfRegisterState     int64  `json:"wqf_register_state"`      // 0 未开通 1 开通中 2 开通成功 3 开通失败 4 申请驳回 5 申请开通中 6 申请开通失败
	WqfRegisterStateDesc string `json:"wqf_register_state_desc"` // 银行转账开通状态描述
	RequestNo            string `json:"request_no"`              // 银行转账开通单号
}

// RetailGetMchOrderResponse 查询商户号开通状态响应
type RetailGetMchOrderResponse struct {
	util.CommonError
	List  []RetailMchOrder `json:"list"`  // 订单列表
	Total int64            `json:"total"` // 返回订单总数
}

// RetailGetMchOrder 查询商户号开通状态
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/bill/api_retailgetmchorder
func (s *B2B) RetailGetMchOrder(in *RetailGetMchOrderRequest) (out RetailGetMchOrderResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(retailGetMchOrderURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "RetailGetMchOrder")
	return
}

// QueryWithdrawRequest 查询提现状态请求
type QueryWithdrawRequest struct {
	PaySig        string `json:"-"`               // pay_sig，请求签名，通过查询参数传递
	MchID         string `json:"mchid"`           // 微信商户号
	OutWithdrawNo string `json:"out_withdraw_no"` // 商户提现单号
}

// QueryWithdrawResponse 查询提现状态响应
type QueryWithdrawResponse struct {
	util.CommonError
	OutWithdrawNo  string `json:"out_withdraw_no"` // 商户提现单号
	WithdrawAmount int64  `json:"withdraw_amount"` // 提现金额
	Status         string `json:"status"`          // 提现状态，如 WITHDRAW_SUCC
	FailReason     string `json:"fail_reason"`     // 失败原因
}

// QueryWithdraw 查询提现状态
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/bill/api_querywithdraw
func (s *B2B) QueryWithdraw(in *QueryWithdrawRequest) (out QueryWithdrawResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(queryWithdrawURL, accessToken, in.PaySig)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "QueryWithdraw")
	return
}

// GetOrderRequest 查询订单请求
type GetOrderRequest struct {
	PaySig     string `json:"-"`            // pay_sig，请求签名，通过查询参数传递
	MchID      string `json:"mchid"`        // 微信商户号
	OutTradeNo string `json:"out_trade_no"` // 商户订单号，长度限制[6,32]，与 order_id 二选一
	OrderID    string `json:"order_id"`     // B2b 支付订单号，长度限制[1,32]，与 out_trade_no 二选一
}

// B2BAmount 订单金额
type B2BAmount struct {
	OrderAmount int64  `json:"order_amount"` // 订单总金额，单位为分
	PayerAmount int64  `json:"payer_amount"` // 用户支付金额，单位为分
	Currency    string `json:"currency"`     // 货币类型，仅支持人民币 "CNY"
}

// GetOrderResponse 查询订单响应
type GetOrderResponse struct {
	util.CommonError
	AppID                 string     `json:"appid"`                   // 小程序 appid
	MchID                 string     `json:"mchid"`                   // 微信商户号
	OutTradeNo            string     `json:"out_trade_no"`            // 商户订单号
	OrderID               string     `json:"order_id"`                // B2b 支付订单号
	PayStatus             string     `json:"pay_status"`              // 支付状态
	PayTime               string     `json:"pay_time"`                // 支付时间
	Attach                string     `json:"attach"`                  // 附加数据
	PayerOpenID           string     `json:"payer_openid"`            // 付款人 openid
	Amount                *B2BAmount `json:"amount"`                  // 订单金额
	WxpayTransactionID    string     `json:"wxpay_transaction_id"`    // 微信支付订单号
	Env                   int64      `json:"env"`                     // 环境
	SettleStatus          int64      `json:"settle_status"`           // 结算状态
	SettleFinishTime      string     `json:"settle_finish_time"`      // 结算完成时间
	PlatformProfitPercent int64      `json:"platform_profit_percent"` // 平台抽成比例
	PlatformProfitFee     int64      `json:"platform_profit_fee"`     // 平台抽成金额
	BankType              string     `json:"bank_type"`               // 银行类型
}

// GetOrder 查询订单
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/bill/api_getorder
func (s *B2B) GetOrder(in *GetOrderRequest) (out GetOrderResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getOrderURL, accessToken, in.PaySig)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetOrder")
	return
}

// GetMchBalanceRequest 查询账户余额请求
type GetMchBalanceRequest struct {
	PaySig string `json:"-"`     // pay_sig，请求签名，通过查询参数传递
	MchID  string `json:"mchid"` // 微信商户号
}

// MchBalance 账户余额
type MchBalance struct {
	BalanceType string `json:"balance_type"` // 余额类型，如 BALANCE_TYPE_AVAILABLE
	Amount      string `json:"amount"`       // 余额金额
	Currency    string `json:"currency"`     // 货币类型，如 CNY
}

// GetMchBalanceResponse 查询账户余额响应
type GetMchBalanceResponse struct {
	util.CommonError
	BalanceList []MchBalance `json:"balance_list"` // 账户余额列表
}

// GetMchBalance 查询账户余额
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/bill/api_getmchbalance
func (s *B2B) GetMchBalance(in *GetMchBalanceRequest) (out GetMchBalanceResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getMchBalanceURL, accessToken, in.PaySig)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetMchBalance")
	return
}

// GetRefundRequest 查询退款请求
type GetRefundRequest struct {
	PaySig      string `json:"-"`             // pay_sig，请求签名，通过查询参数传递
	MchID       string `json:"mchid"`         // 微信商户号
	OutRefundNo string `json:"out_refund_no"` // 商户退款单号
	RefundID    string `json:"refund_id"`     // 微信退款单号
}

// RefundAmount 退款金额
type RefundAmount struct {
	OrderAmount  int64  `json:"order_amount"`  // 订单总金额，单位为分
	RefundAmount int64  `json:"refund_amount"` // 退款金额，单位为分
	Currency     string `json:"currency"`      // 货币类型，仅支持人民币 "CNY"
}

// RefundChannelInfo 退款渠道信息
type RefundChannelInfo struct {
	Channel             string `json:"channel"`               // 退款渠道，如 ORIGINAL、BALANCE、OTHER_BALANCE
	UserReceivedAccount string `json:"user_received_account"` // 用户收到的金额
	FundsAccount        string `json:"funds_account"`         // 资金账户，如 AVAILABLE
}

// GetRefundResponse 查询退款响应
type GetRefundResponse struct {
	util.CommonError
	RefundID              string             `json:"refund_id"`                // 微信退款单号
	OutRefundNo           string             `json:"out_refund_no"`            // 商户退款单号
	OrderID               string             `json:"order_id"`                 // B2b 支付订单号
	OutTradeNo            string             `json:"out_trade_no"`             // 商户订单号
	CreateTime            string             `json:"create_time"`              // 创建时间
	RefundTime            string             `json:"refund_time"`              // 退款时间
	RefundStatus          string             `json:"refund_status"`            // 退款状态，如 REFUND_SUCC
	RefundDesc            string             `json:"refund_desc"`              // 退款状态描述
	Amount                *RefundAmount      `json:"amount"`                   // 退款金额
	WxpayRefundID         string             `json:"wxpay_refund_id"`          // 微信支付退款单号
	ReverseSettState      int64              `json:"reverse_sett_state"`       // 回退结算状态
	ReverseSettFinishTime string             `json:"reverse_sett_finish_time"` // 回退结算完成时间
	PlatformProfitPercent int64              `json:"platform_profit_percent"`  // 平台抽成比例
	ReverseSettAmt        int64              `json:"reverse_sett_amt"`         // 回退结算金额
	RefundChannelInfo     *RefundChannelInfo `json:"refund_channel_info"`      // 退款渠道信息
	Description           string             `json:"description"`              // 商品描述
}

// GetRefund 查询退款
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/bill/api_getrefund
func (s *B2B) GetRefund(in *GetRefundRequest) (out GetRefundResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getRefundURL, accessToken, in.PaySig)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetRefund")
	return
}

// GetWqfChargeFeeRequest 查询银行转账的技术服务费率请求
type GetWqfChargeFeeRequest struct {
	SubMchID string `json:"sub_mchid"` // 商户号
}

// GetWqfChargeFeeResponse 查询银行转账的技术服务费率响应
type GetWqfChargeFeeResponse struct {
	util.CommonError
	MerchantName                  string `json:"merchant_name"`                    // 商户名称
	CertifiedChargeFeeNumerator   int64  `json:"certified_charge_fee_numerator"`   // 认证费率分子，分母统一为 10000
	UncertifiedChargeFeeNumerator int64  `json:"uncertified_charge_fee_numerator"` // 未认证的费率分子，分母统一为 10000
}

// GetWqfChargeFee 查询银行转账的技术服务费率
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/bill/api_getwqfchargefee
func (s *B2B) GetWqfChargeFee(in *GetWqfChargeFeeRequest) (out GetWqfChargeFeeResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getWqfChargeFeeURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetWqfChargeFee")
	return
}

// AddProfitSharingAccountRequest 添加分账方请求
type AddProfitSharingAccountRequest struct {
	PaySig                    string `json:"-"`                            // pay_sig，请求签名，通过查询参数传递
	ProfitSharingRelationType string `json:"profit_sharing_relation_type"` // 分账接收方关系类型
	PayeeType                 string `json:"payee_type"`                   // 分账接收方类型
	PayeeID                   string `json:"payee_id"`                     // 分账接收方标识，外部用户填 openid，外部商户填商户号
	PayeeName                 string `json:"payee_name"`                   // 分账接收方名称，接收方为商户号时必填
}

// AddProfitSharingAccountResponse 添加分账方响应
type AddProfitSharingAccountResponse struct {
	util.CommonError
}

// AddProfitSharingAccount 添加分账方
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/bill/api_addprofitsharingaccount
func (s *B2B) AddProfitSharingAccount(in *AddProfitSharingAccountRequest) (out AddProfitSharingAccountResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(addProfitSharingAccountURL, accessToken, in.PaySig)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "AddProfitSharingAccount")
	return
}

// RegisterOnlyWqfRequest 申请开通银行转账请求
type RegisterOnlyWqfRequest struct {
	OutRegistrationID string `json:"out_registration_id"` // 订单号，原有进件订单号
}

// RegisterOnlyWqfResponse 申请开通银行转账响应
type RegisterOnlyWqfResponse struct {
	util.CommonError
}

// RegisterOnlyWqf 申请开通银行转账
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/bill/api_registeronlywqf
func (s *B2B) RegisterOnlyWqf(in *RegisterOnlyWqfRequest) (out RegisterOnlyWqfResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(registerOnlyWqfURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "RegisterOnlyWqf")
	return
}

// GetAppKeyRequest 获取密钥 AppKey 请求
type GetAppKeyRequest struct {
	MchID string `json:"mchid"` // 微信商户号
}

// GetAppKeyResponse 获取密钥 AppKey 响应
type GetAppKeyResponse struct {
	util.CommonError
	Appkey        string `json:"appkey"`         // 正式环境 AppKey
	SandboxAppkey string `json:"sandbox_appkey"` // 沙箱环境 AppKey
}

// GetAppKey 获取密钥 AppKey
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/bill/api_getappkey
func (s *B2B) GetAppKey(in *GetAppKeyRequest) (out GetAppKeyResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getAppKeyURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetAppKey")
	return
}

// MchInfo 商户信息
type MchInfo struct {
	SubMchID           string `json:"sub_mchid"`            // 商户号，有商户号不代表已开通成功，wxpay_status 为"完成"才能使用
	CompanyName        string `json:"company_name"`         // 企业名称
	BankName           string `json:"bank_name"`            // 开户银行
	BankAccount        string `json:"bank_account"`         // 银行账号，脱敏处理
	WxpayStatus        string `json:"wxpay_status"`         // 微信支付开通状态
	BankTransferStatus string `json:"bank_transfer_status"` // 银行转账开通状态
}

// GetMchInfoResponse 获取小程序下所有商户的信息响应
type GetMchInfoResponse struct {
	util.CommonError
	MchList []MchInfo `json:"mch_list"` // 商户列表
	Total   int64     `json:"total"`    // 总数
}

// GetMchInfo 获取小程序下所有商户的信息
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/bill/api_getmchinfo
func (s *B2B) GetMchInfo() (out GetMchInfoResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(getMchInfoURL, accessToken)
	response, err := util.PostJSON(uri, struct{}{})
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "GetMchInfo")
	return
}

// CreateProfitSharingOrderRequest 请求分账请求
type CreateProfitSharingOrderRequest struct {
	PaySig          string `json:"-"`                // pay_sig，请求签名，通过查询参数传递
	MchID           string `json:"mchid"`            // 子商户 id，在 B2b 小程序中进件完成返回的子商户 id
	OutTradeNo      string `json:"out_trade_no"`     // 支付单 id，在 B2b 小程序中下单的订单 id
	ProfitFee       int64  `json:"profit_fee"`       // 分账费用，单位：分，不超过支付单本身的金额
	ReceiverType    string `json:"receiver_type"`    // 分账接收方类型，同添加分账方时填入的内容
	ReceiverAccount string `json:"receiver_account"` // 分账接收方账号，同添加分账方时填入的内容
}

// CreateProfitSharingOrderResponse 请求分账响应
type CreateProfitSharingOrderResponse struct {
	util.CommonError
}

// CreateProfitSharingOrder 请求分账
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/bill/api_createprofitsharingorder
func (s *B2B) CreateProfitSharingOrder(in *CreateProfitSharingOrderRequest) (out CreateProfitSharingOrderResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(createProfitSharingOrderURL, accessToken, in.PaySig)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "CreateProfitSharingOrder")
	return
}

// RefundProfitSharingRequest 请求分账回退请求
type RefundProfitSharingRequest struct {
	PaySig      string `json:"-"`             // pay_sig，请求签名，通过查询参数传递
	OutTradeNo  string `json:"out_trade_no"`  // 订单 id，在 B2b 小程序中下单的订单 id
	OutRefundNo string `json:"out_refund_no"` // 退款单 id，在 B2b 小程序中下单的退款单 id
	PayeeType   string `json:"payee_type"`    // 退款分账方类型，同添加分账方时填入的内容
	PayeeID     string `json:"payee_id"`      // 退款分账方 id，同添加分账方时填入的内容
	MchID       string `json:"mchid"`         // 商户号 id，发起这笔订单的商户号
	RefundAmt   int64  `json:"refund_amt"`    // 退款金额，单位为分
}

// RefundProfitSharingResponse 请求分账回退响应
type RefundProfitSharingResponse struct {
	util.CommonError
}

// RefundProfitSharing 请求分账回退
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/bill/api_refundprofitsharing
func (s *B2B) RefundProfitSharing(in *RefundProfitSharingRequest) (out RefundProfitSharingResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(refundProfitSharingURL, accessToken, in.PaySig)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "RefundProfitSharing")
	return
}

// CreateWqfLinkRequest 跳转银行转账页面请求
type CreateWqfLinkRequest struct {
	RequestNo string `json:"request_no"` // 银行转账开通单号
}

// CreateWqfLinkResponse 跳转银行转账页面响应
type CreateWqfLinkResponse struct {
	util.CommonError
	URL        string `json:"url"`         // 银行转账页面链接
	ExpireTime string `json:"expire_time"` // 链接过期时间
}

// CreateWqfLink 跳转银行转账页面
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/bill/api_createwqflink
func (s *B2B) CreateWqfLink(in *CreateWqfLinkRequest) (out CreateWqfLinkResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(createWqfLinkURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "CreateWqfLink")
	return
}

// RefundOrderRequest 退款请求
type RefundOrderRequest struct {
	PaySig       string `json:"-"`             // pay_sig，请求签名，通过查询参数传递
	MchID        string `json:"mchid"`         // 微信商户号
	OutTradeNo   string `json:"out_trade_no"`  // 商户订单号，与 order_id 二选一
	OrderID      string `json:"order_id"`      // B2b 支付订单号，与 out_trade_no 二选一
	OutRefundNo  string `json:"out_refund_no"` // 商户退款单号
	RefundAmount int64  `json:"refund_amount"` // 退款金额，单位为分
	RefundFrom   int64  `json:"refund_from"`   // 退款资金来源
	RefundReason int64  `json:"refund_reason"` // 退款原因
	Description  string `json:"description"`   // 商品描述
}

// RefundOrderResponse 退款响应
type RefundOrderResponse struct {
	util.CommonError
	RefundID    string `json:"refund_id"`     // 微信退款单号
	OutRefundNo string `json:"out_refund_no"` // 商户退款单号
	OrderID     string `json:"order_id"`      // B2b 支付订单号
	OutTradeNo  string `json:"out_trade_no"`  // 商户订单号
}

// RefundOrder 退款
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/B2b/bill/api_refundorder
func (s *B2B) RefundOrder(in *RefundOrderRequest) (out RefundOrderResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(refundOrderURL, accessToken, in.PaySig)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "RefundOrder")
	return
}
