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

package hardware

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/miniprogram/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	// sendHardwareDeviceMessageURL 发送设备消息
	sendHardwareDeviceMessageURL = "https://api.weixin.qq.com/cgi-bin/message/device/subscribe/send?access_token=%s"
	// getSnTicketURL 获取设备票据
	getSnTicketURL = "https://api.weixin.qq.com/wxa/getsnticket?access_token=%s"
	// createIotGroupIdURL 创建设备组
	createIotGroupIdURL = "https://api.weixin.qq.com/wxa/business/group/createid?access_token=%s"
	// getIotGroupInfoURL 查询设备组信息
	getIotGroupInfoURL = "https://api.weixin.qq.com/wxa/business/group/getinfo?access_token=%s"
	// addIotGroupDeviceURL 设备组添加设备
	addIotGroupDeviceURL = "https://api.weixin.qq.com/wxa/business/group/adddevice?access_token=%s"
	// removeIotGroupDeviceURL 设备组删除设备
	removeIotGroupDeviceURL = "https://api.weixin.qq.com/wxa/business/group/removedevice?access_token=%s"
	// getLicensePkgListURL 查询 license 资源包列表
	getLicensePkgListURL = "https://api.weixin.qq.com/wxa/business/license/getpkglist?access_token=%s"
	// activeLicenseDeviceURL 激活设备 license
	activeLicenseDeviceURL = "https://api.weixin.qq.com/wxa/business/license/activedevice?access_token=%s"
	// getLicenseDeviceInfoURL 查询设备激活详情
	getLicenseDeviceInfoURL = "https://api.weixin.qq.com/wxa/business/license/getdeviceinfo?access_token=%s"
)

// Hardware 硬件设备
type Hardware struct {
	*context.Context
}

// NewHardware 实例化硬件设备
func NewHardware(ctx *context.Context) *Hardware {
	return &Hardware{ctx}
}

// DeviceListItem 设备列表项，包含设备型号与设备序列号
type DeviceListItem struct {
	ModelID string `json:"model_id"` // 设备型号 id。通过微信公众平台注册设备获得。
	SN      string `json:"sn"`       // 设备唯一序列号。由厂商分配。
	ErrCode int64  `json:"errcode"`  // 该设备的错误码，0 表示成功
}

// DeviceExpireItem 设备激活详情列表项，包含设备型号、设备序列号与过期时间
type DeviceExpireItem struct {
	ModelID    string `json:"model_id"`    // 设备型号唯一标识
	SN         string `json:"sn"`          // 设备的唯一标识
	ExpireTime int64  `json:"expire_time"` // 设备的过期时间
	ErrCode    int64  `json:"errcode"`     // 该设备的错误码，0 表示成功
}

// ActiveLicenseDeviceItem 激活 license 的设备列表项
type ActiveLicenseDeviceItem struct {
	ModelID      string `json:"model_id"`      // 设备型号 id。通过微信公众平台注册设备获得。
	SN           string `json:"sn"`            // 设备唯一序列号。由厂商分配。
	ActiveNumber int64  `json:"active_number"` // 激活码序号，任意 uint32 整数（需与之前使用过的不重复）。主要用于防止重复请求导致重复激活。
}

// LicensePkgItem license 资源包列表项
type LicensePkgItem struct {
	PkgID     string `json:"pkg_id"`     // 资源包 ID
	PkgType   int64  `json:"pkg_type"`   // 资源包类型
	StartTime int64  `json:"start_time"` // 资源包下单时间
	EndTime   int64  `json:"end_time"`   // 资源包过期时间
	PkgStatus int64  `json:"pkg_status"` // 资源包状态，1为已生效，2为未生效，3为已过期
	Used      int64  `json:"used"`       // 已使用额度
	All       int64  `json:"all"`        // 资源包总量
}

// SendHardwareDeviceMessageRequest 发送设备消息请求
type SendHardwareDeviceMessageRequest struct {
	TemplateID       string          `json:"template_id"`                 // 所需下发的订阅模板id
	SN               string          `json:"sn"`                          // 设备唯一序列号。由厂商分配，长度不能超过128字节。字符只接受数字，大小写字母，下划线（_）和连字符（-）。
	Page             string          `json:"page"`                        // 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
	ToOpenIDList     []string        `json:"to_openid_list"`              // 接收者（用户）的 openid 列表
	MiniprogramState string          `json:"miniprogram_state,omitempty"` // 跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
	ModelID          string          `json:"modelId"`                     // 设备型号 id ，通过注册设备获得。
	Data             json.RawMessage `json:"data"`                        // 模板内容，格式形如 { "key1": { "value": "xxx" }, "key2": { "value": "xxx" } } ，value 为枚举值。
	Lang             string          `json:"lang"`                        // 进入小程序查看”的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN
}

// SendHardwareDeviceMessageResponse 发送设备消息响应
type SendHardwareDeviceMessageResponse struct {
	util.CommonError
}

// SendHardwareDeviceMessage 发送设备消息
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/hardware-device/api_sendhardwaredevicemessage
func (s *Hardware) SendHardwareDeviceMessage(in *SendHardwareDeviceMessageRequest) (out SendHardwareDeviceMessageResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	response, err := util.PostJSON(fmt.Sprintf(sendHardwareDeviceMessageURL, accessToken), in)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &out, "SendHardwareDeviceMessage")
	return
}

// GetSnTicketRequest 获取设备票据请求
type GetSnTicketRequest struct {
	SN      string `json:"sn"`       // sn 设备唯一序列号。由厂商分配，长度不能超过128字节。字符只接受数字，大小写字母，下划线（_）和连字符（-）。
	ModelID string `json:"model_id"` // 设备型号 id ，通过注册设备获得。
}

// GetSnTicketResponse 获取设备票据响应
type GetSnTicketResponse struct {
	util.CommonError
	SNTicket string `json:"sn_ticket"` // 设备票据，5分钟内有效。
}

// GetSnTicket 获取设备票据
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/hardware-device/api_getsnticket
func (s *Hardware) GetSnTicket(in *GetSnTicketRequest) (out GetSnTicketResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	response, err := util.PostJSON(fmt.Sprintf(getSnTicketURL, accessToken), in)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &out, "GetSnTicket")
	return
}

// CreateIotGroupIdRequest 创建设备组请求
type CreateIotGroupIdRequest struct {
	ModelID   string `json:"model_id"`   // 设备型号的唯一标识。
	GroupName string `json:"group_name"` // 设备组的名称（创建时时决定，无法修改）
}

// CreateIotGroupIdResponse 创建设备组响应
type CreateIotGroupIdResponse struct {
	util.CommonError
	GroupID string `json:"group_id"` // 设备组的唯一标识
}

// CreateIotGroupId 创建设备组
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/hardware-device/api_createiotgroupid
func (s *Hardware) CreateIotGroupId(in *CreateIotGroupIdRequest) (out CreateIotGroupIdResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	response, err := util.PostJSON(fmt.Sprintf(createIotGroupIdURL, accessToken), in)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &out, "CreateIotGroupId")
	return
}

// GetIotGroupInfoRequest 查询设备组信息请求
type GetIotGroupInfoRequest struct {
	GroupID string `json:"group_id"` // 设备组的唯一标识
}

// GetIotGroupInfoResponse 查询设备组信息响应
type GetIotGroupInfoResponse struct {
	util.CommonError
	GroupName  string           `json:"group_name"`  // 设备名称
	DeviceList []DeviceListItem `json:"device_list"` // 设备列表
	ModelID    string           `json:"model_id"`    // 设备型号的唯一标识
	ModelType  string           `json:"model_type"`  // 设备类型
}

// GetIotGroupInfo 查询设备组信息
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/hardware-device/api_getiotgroupinfo
func (s *Hardware) GetIotGroupInfo(in *GetIotGroupInfoRequest) (out GetIotGroupInfoResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	response, err := util.PostJSON(fmt.Sprintf(getIotGroupInfoURL, accessToken), in)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &out, "GetIotGroupInfo")
	return
}

// AddIotGroupDeviceRequest 设备组添加设备请求
type AddIotGroupDeviceRequest struct {
	GroupID    string           `json:"group_id"`            // 设备组的唯一标识
	DeviceList []DeviceListItem `json:"device_list"`         // 设备列表
	ForceAdd   bool             `json:"force_add,omitempty"` // 是否强制更新设备列表，等于 true 时将已存在其它设备组中的设备移除并添加到当前设备组，慎用。
}

// AddIotGroupDeviceResponse 设备组添加设备响应
type AddIotGroupDeviceResponse struct {
	util.CommonError
	DeviceList []DeviceListItem `json:"device_list"` // 设备列表
}

// AddIotGroupDevice 设备组添加设备
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/hardware-device/api_addiotgroupdevice
func (s *Hardware) AddIotGroupDevice(in *AddIotGroupDeviceRequest) (out AddIotGroupDeviceResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	response, err := util.PostJSON(fmt.Sprintf(addIotGroupDeviceURL, accessToken), in)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &out, "AddIotGroupDevice")
	return
}

// RemoveIotGroupDeviceRequest 设备组删除设备请求
type RemoveIotGroupDeviceRequest struct {
	GroupID    string           `json:"group_id"`    // 设备组的唯一标识
	DeviceList []DeviceListItem `json:"device_list"` // 设备列表
}

// RemoveIotGroupDeviceResponse 设备组删除设备响应
type RemoveIotGroupDeviceResponse struct {
	util.CommonError
	DeviceList []DeviceListItem `json:"device_list"` // 设备列表
}

// RemoveIotGroupDevice 设备组删除设备
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/hardware-device/api_removeiotgroupdevice
func (s *Hardware) RemoveIotGroupDevice(in *RemoveIotGroupDeviceRequest) (out RemoveIotGroupDeviceResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	response, err := util.PostJSON(fmt.Sprintf(removeIotGroupDeviceURL, accessToken), in)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &out, "RemoveIotGroupDevice")
	return
}

// GetLicensePkgListRequest 查询 license 资源包列表请求
type GetLicensePkgListRequest struct {
	PkgType int64 `json:"pkg_type"` // 资源包类型，0：测试体验包，1：A 类设备，2：B 类设备，3：C 类设备，4：D 类设备，5：E类设备
}

// GetLicensePkgListResponse 查询 license 资源包列表响应
type GetLicensePkgListResponse struct {
	util.CommonError
	PkgList         []LicensePkgItem `json:"pkg_list"`          // 资源包列表
	MaxActiveNumber int64            `json:"max_active_number"` // 最大激活码序号，已废弃。
}

// GetLicensePkgList 查询 license 资源包列表
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/hardware-device/api_getlicensepkglist
func (s *Hardware) GetLicensePkgList(in *GetLicensePkgListRequest) (out GetLicensePkgListResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	response, err := util.PostJSON(fmt.Sprintf(getLicensePkgListURL, accessToken), in)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &out, "GetLicensePkgList")
	return
}

// ActiveLicenseDeviceRequest 激活设备 license 请求
type ActiveLicenseDeviceRequest struct {
	DeviceList []ActiveLicenseDeviceItem `json:"device_list"` // 待激活的设备列表
	PkgType    int64                     `json:"pkg_type"`    // 资源包类型，0：测试体验包（默认），1：A 类设备，2：B 类设备，3：C 类设备，4：D 类设备，5：E类设备
}

// ActiveLicenseDeviceResponse 激活设备 license 响应
type ActiveLicenseDeviceResponse struct {
	util.CommonError
	DeviceList []DeviceExpireItem `json:"device_list"` // 设备列表
}

// ActiveLicenseDevice 激活设备 license
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/hardware-device/api_activelicensedevice
func (s *Hardware) ActiveLicenseDevice(in *ActiveLicenseDeviceRequest) (out ActiveLicenseDeviceResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	response, err := util.PostJSON(fmt.Sprintf(activeLicenseDeviceURL, accessToken), in)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &out, "ActiveLicenseDevice")
	return
}

// GetLicenseDeviceInfoRequest 查询设备激活详情请求
type GetLicenseDeviceInfoRequest struct {
	DeviceList []DeviceListItem `json:"device_list"` // 设备列表
}

// GetLicenseDeviceInfoResponse 查询设备激活详情响应
type GetLicenseDeviceInfoResponse struct {
	util.CommonError
	DeviceList []DeviceExpireItem `json:"device_list"` // 设备列表
}

// GetLicenseDeviceInfo 查询设备激活详情
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/hardware-device/api_getlicensedeviceinfo
func (s *Hardware) GetLicenseDeviceInfo(in *GetLicenseDeviceInfoRequest) (out GetLicenseDeviceInfoResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	response, err := util.PostJSON(fmt.Sprintf(getLicenseDeviceInfoURL, accessToken), in)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, &out, "GetLicenseDeviceInfo")
	return
}
