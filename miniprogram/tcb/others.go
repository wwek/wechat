package tcb

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// 云开发上报接口
	cloudbaseReportURL = "https://api.weixin.qq.com/tcb/cloudbasereport"
	// 创建发短信任务
	createSendSmsTaskURL = "https://api.weixin.qq.com/tcb/createsendsmstask"
	// 描述扩展上传文件信息
	describeExtensionUploadInfoURL = "https://api.weixin.qq.com/tcb/describeextensionuploadinfo"
	// 查询短信记录
	describeSmsRecordsURL = "https://api.weixin.qq.com/tcb/describesmsrecords"
	// 获取cloudID对应的数据
	getOpenDataURL = "https://api.weixin.qq.com/wxa/getopendata"
	// 获取云开发数据
	getCloudBaseStatisticsURL = "https://api.weixin.qq.com/tcb/getstatistics"
	// 获取实时语音签名
	getCloudBaseVoIPSignURL = "https://api.weixin.qq.com/wxa/getvoipsign"
	// 获取腾讯云API调用凭证
	getCloudTokenURL = "https://api.weixin.qq.com/tcb/getqcloudtoken"
)

// CloudbaseReportAPIRequest 云开发上报接口 请求参数
type CloudbaseReportAPIRequest struct {
	ReportAction string `json:"report_action,omitempty"` // 上报动作，目前支持（sendSmsTask：发送短信；openH5：H5 打开）
	EnvID        string `json:"env_id,omitempty"`        // 环境 ID
	ActivityID   string `json:"activity_id,omitempty"`   // 活动 ID
	TaskID       string `json:"task_id,omitempty"`       // 任务 ID【report_action 取 sendSmsTask 时必填】
	PhoneCount   string `json:"phone_count,omitempty"`   // 下发手机号数量【report_action 取 sendSmsTask 时必填】
	ChannelID    string `json:"channel_id,omitempty"`    // 渠道 ID（云开发 CMS 使用 _cms_sms_）【report_action 取 openH5 时必填】
	SessionID    string `json:"session_id,omitempty"`    // 会话 ID【report_action 取 openH5 时必填】
}

// CloudbaseReportAPIResponse 云开发上报接口 返回结果
type CloudbaseReportAPIResponse struct {
	util.CommonError
}

// CreateSendSmsTaskRequest 创建发短信任务 请求参数
type CreateSendSmsTaskRequest struct {
	Env        string `json:"env,omitempty"`         // 环境 ID
	FileURL    string `json:"file_url,omitempty"`    // 短信 CSV 文件地址CodeUri
	TemplateID string `json:"template_id,omitempty"` // 短信模版 ID 默认值：844110（销类短信模版 ID)
}

// CreateSendSmsTaskResponse 创建发短信任务 返回结果
type CreateSendSmsTaskResponse struct {
	util.CommonError
	QueryID string `json:"query_id"` // 查询 ID
}

// ExtensionFile 待上传的文件信息
type ExtensionFile struct {
	FileType string `json:"FileType"` // 文件类型。枚举值 FUNCTION：函数代码 STATIC：静态托管代码 SMS：短信文件
	FileName string `json:"FileName"` // 文件名，长度不超过24
}

// ExtensionFileData 待上传文件的信息
type ExtensionFileData struct {
	CodeURI   string `json:"CodeUri"`   // 模板里使用的地址
	UploadURL string `json:"UploadUrl"` // 上传文件的临时地址，含签名
	CustomKey string `json:"CustomKey"` // 自定义密钥。如果为空，则表示不需要加密
	MaxSize   int64  `json:"MaxSize"`   // 文件大小限制，单位M，客户端上传前需要主动检查文件大小，超过限制的文件会被删除。
}

// DescribeExtensionUploadInfoRequest 描述扩展上传文件信息 请求参数
type DescribeExtensionUploadInfoRequest struct {
	ExtensionFiles []ExtensionFile `json:"ExtensionFiles,omitempty"` // 待上传的文件列表
}

// DescribeExtensionUploadInfoResponse 描述扩展上传文件信息 返回结果
type DescribeExtensionUploadInfoResponse struct {
	util.CommonError
	FilesData []ExtensionFileData `json:"FilesData"` // 待上传文件的信息数组
	RequestID string              `json:"RequestId"` // 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
}

// SmsRecord 发送记录
type SmsRecord struct {
	Mobile       string `json:"Mobile"`       // 手机号码
	Content      string `json:"Content"`      // 短信内容
	ContentSize  int64  `json:"ContentSize"`  // 短信内容长度
	Fee          int64  `json:"Fee"`          // 计费条数
	CreateTime   string `json:"CreateTime"`   // 发送时间
	ReceivedTime string `json:"ReceivedTime"` // 用户接收时间
	Status       string `json:"Status"`       // sent(成功), error(失败)
	Remarks      string `json:"Remarks"`      // 备注
}

// DescribeSmsRecordsRequest 查询短信记录 请求参数
type DescribeSmsRecordsRequest struct {
	EnvID      string `json:"EnvId,omitempty"`      // 环境 ID
	StartDate  string `json:"StartDate,omitempty"`  // 开始日期, 如:2021-01-01
	EndDate    string `json:"EndDate,omitempty"`    // 结束日期, 如2021-01-07
	Mobile     string `json:"Mobile,omitempty"`     // 电话号码
	QueryID    string `json:"QueryId,omitempty"`    // 查询ID
	PageNumber int64  `json:"PageNumber,omitempty"` // 页码(1开始)
	PageSize   int64  `json:"PageSize,omitempty"`   // 每页条目数
}

// DescribeSmsRecordsResponse 查询短信记录 返回结果
type DescribeSmsRecordsResponse struct {
	util.CommonError
	SmsRecords []SmsRecord `json:"SmsRecords"` // 发送记录列表
	TotalCount int64       `json:"TotalCount"` // 记录总数
	RequestID  string      `json:"RequestId"`  // 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
}

// GetOpenDataRequest 获取cloudID对应的数据 请求参数
type GetOpenDataRequest struct {
	OpenID      string   `json:"openid,omitempty"`       // 用户openid（敏感信息需要传入），作为 query 参数发送
	CloudIDList []string `json:"cloudid_list,omitempty"` // CloudID 列表
}

// OpenData 开放数据
type OpenData struct {
	CloudID string                 `json:"cloud_id"` // cloud id
	JSON    map[string]interface{} `json:"json"`     // 数据详情
}

// GetOpenDataResponse 获取cloudID对应的数据 返回结果
type GetOpenDataResponse struct {
	util.CommonError
	DataList []OpenData `json:"data_list"` // 开放数据列表
}

// StatisticsColumn 数据列定义
type StatisticsColumn struct {
	ColID       string `json:"col_id"`        // 列 id
	ColName     string `json:"col_name"`      // 列名
	ColDataType string `json:"col_data_type"` // 数据类型（0:string；1:number；2:double）
}

// StatisticsRow 数据行
type StatisticsRow struct {
	DataValue []interface{} `json:"data_value"` // 数据值
}

// GetCloudBaseStatisticsRequest 获取云开发数据 请求参数
type GetCloudBaseStatisticsRequest struct {
	Action     string `json:"action,omitempty"`      // 获取动作，目前支持（smsMarketingOverviewData：短信营销概览数据；smsMarketingConversionData：短信营销转化数据；smsMarketingRealTimeData：短信营销实时数据）
	BeginDate  int64  `json:"begin_date,omitempty"`  // 开始时间戳
	EndDate    int64  `json:"end_date,omitempty"`    // 结束时间戳
	PageLimit  int64  `json:"page_limit,omitempty"`  // 分页 limit【action 取 smsMarketingOverviewData、smsMarketingConversionData 时必填】
	PageOffset int64  `json:"page_offset,omitempty"` // 分页 offset【action 取 smsMarketingOverviewData、smsMarketingConversionData 时必填】
}

// GetCloudBaseStatisticsResponse 获取云开发数据 返回结果
type GetCloudBaseStatisticsResponse struct {
	util.CommonError
	DataColumn []StatisticsColumn `json:"data_column"` // 数据列定义
	DataValue  []StatisticsRow    `json:"data_value"`  // 数据行
	TotalNum   int64              `json:"total_num"`   // 总行数
}

// GetCloudBaseVoIPSignRequest 获取实时语音签名 请求参数
type GetCloudBaseVoIPSignRequest struct {
	GroupID   string `json:"group_id,omitempty"`  // 游戏房间的标识
	Timestamp int64  `json:"timestamp,omitempty"` // 生成这个随机字符串的 UNIX 时间戳（精确到秒）
	Nonce     string `json:"nonce,omitempty"`     // 随机字符串，长度应小于 128
}

// GetCloudBaseVoIPSignResponse 获取实时语音签名 返回结果
type GetCloudBaseVoIPSignResponse struct {
	util.CommonError
	Signature string `json:"signature"` // 签名
}

// GetCloudTokenRequest 获取腾讯云API调用凭证 请求参数
type GetCloudTokenRequest struct {
	Lifespan int64 `json:"lifespan,omitempty"` // 有效期（单位为秒，最大7200）
}

// GetCloudTokenResponse 获取腾讯云API调用凭证 返回结果
type GetCloudTokenResponse struct {
	util.CommonError
	SecretID    string `json:"secretid"`     // secretid
	SecretKey   string `json:"secretkey"`    // secretkey
	Token       string `json:"token"`        // token
	ExpiredTime int64  `json:"expired_time"` // 过期时间戳
}

// CloudbaseReportAPI 云开发上报接口
//
// reference:https://developers.weixin.qq.com/miniprogram/dev/server/API/cloudbase/others/api_cloudbasereportapi
func (s *Tcb) CloudbaseReportAPI(in *CloudbaseReportAPIRequest) (out CloudbaseReportAPIResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", cloudbaseReportURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	err = util.DecodeWithError(response, &out, "CloudbaseReportAPI")
	return
}

// CreateSendSmsTask 创建发短信任务
//
// reference:https://developers.weixin.qq.com/miniprogram/dev/server/API/cloudbase/others/api_createsendsmstask
func (s *Tcb) CreateSendSmsTask(in *CreateSendSmsTaskRequest) (out CreateSendSmsTaskResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", createSendSmsTaskURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	err = util.DecodeWithError(response, &out, "CreateSendSmsTask")
	return
}

// DescribeExtensionUploadInfo 描述扩展上传文件信息
//
// reference:https://developers.weixin.qq.com/miniprogram/dev/server/API/cloudbase/others/api_describeextensionuploadinfo
func (s *Tcb) DescribeExtensionUploadInfo(in *DescribeExtensionUploadInfoRequest) (out DescribeExtensionUploadInfoResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", describeExtensionUploadInfoURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	err = util.DecodeWithError(response, &out, "DescribeExtensionUploadInfo")
	return
}

// DescribeSmsRecords 查询短信记录
//
// reference:https://developers.weixin.qq.com/miniprogram/dev/server/API/cloudbase/others/api_describesmsrecords
func (s *Tcb) DescribeSmsRecords(in *DescribeSmsRecordsRequest) (out DescribeSmsRecordsResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", describeSmsRecordsURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	err = util.DecodeWithError(response, &out, "DescribeSmsRecords")
	return
}

// GetOpenData 获取cloudID对应的数据
//
// reference:https://developers.weixin.qq.com/miniprogram/dev/server/API/cloudbase/others/api_getopendata
func (s *Tcb) GetOpenData(in *GetOpenDataRequest) (out GetOpenDataResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", getOpenDataURL, accessToken)
	// openid 属于查询参数，附加在 url 上，不放入请求体
	if in.OpenID != "" {
		uri = fmt.Sprintf("%s&%s", uri, util.Query(map[string]interface{}{"openid": in.OpenID}))
	}
	body := map[string]interface{}{"cloudid_list": in.CloudIDList}
	response, err := util.PostJSON(uri, body)
	if err != nil {
		return
	}
	err = util.DecodeWithError(response, &out, "GetOpenData")
	return
}

// GetCloudBaseStatistics 获取云开发数据
//
// reference:https://developers.weixin.qq.com/miniprogram/dev/server/API/cloudbase/others/api_getcloudbasestatistics
func (s *Tcb) GetCloudBaseStatistics(in *GetCloudBaseStatisticsRequest) (out GetCloudBaseStatisticsResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", getCloudBaseStatisticsURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	err = util.DecodeWithError(response, &out, "GetCloudBaseStatistics")
	return
}

// GetCloudBaseVoIPSign 获取实时语音签名
//
// reference:https://developers.weixin.qq.com/miniprogram/dev/server/API/cloudbase/others/api_getcloudbasevoipsign
func (s *Tcb) GetCloudBaseVoIPSign(in *GetCloudBaseVoIPSignRequest) (out GetCloudBaseVoIPSignResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", getCloudBaseVoIPSignURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	err = util.DecodeWithError(response, &out, "GetCloudBaseVoIPSign")
	return
}

// GetCloudToken 获取腾讯云API调用凭证
//
// reference:https://developers.weixin.qq.com/miniprogram/dev/server/API/cloudbase/others/api_getcloudtoken
func (s *Tcb) GetCloudToken(in *GetCloudTokenRequest) (out GetCloudTokenResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", getCloudTokenURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	err = util.DecodeWithError(response, &out, "GetCloudToken")
	return
}
