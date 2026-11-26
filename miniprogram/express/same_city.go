package express

import (
	"context"
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// 同城配送扣费主体与城市 API URL
const (
	// 查询扣费主体
	intracityGetPayModeURL = "https://api.weixin.qq.com/cgi-bin/express/intracity/getpaymode?access_token=%s"
	// 查询支持同城配送的城市
	intracityGetCityURL = "https://api.weixin.qq.com/cgi-bin/express/intracity/getcity?access_token=%s"
	// 设置扣费主体
	intracitySetPayModeURL = "https://api.weixin.qq.com/cgi-bin/express/intracity/setpaymode?access_token=%s"
)

// IntracityGetPayMode 查询扣费主体
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/express/business/intracity_service.html
func (express *Express) IntracityGetPayMode(ctx context.Context, req *IntracityGetPayModeRequest) (res IntracityGetPayModeResponse, err error) {
	accessToken, err := express.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(intracityGetPayModeURL, accessToken)
	response, err := util.PostJSONContext(ctx, uri, req)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &res, "IntracityGetPayMode")
	return
}

// IntracityGetCity 查询支持同城配送的城市
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/express/business/intracity_service.html
func (express *Express) IntracityGetCity(ctx context.Context, req *IntracityGetCityRequest) (res IntracityGetCityResponse, err error) {
	accessToken, err := express.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(intracityGetCityURL, accessToken)
	response, err := util.PostJSONContext(ctx, uri, req)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &res, "IntracityGetCity")
	return
}

// IntracitySetPayMode 设置扣费主体
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/express/business/intracity_service.html
func (express *Express) IntracitySetPayMode(ctx context.Context, req *IntracitySetPayModeRequest) (res IntracitySetPayModeResponse, err error) {
	accessToken, err := express.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(intracitySetPayModeURL, accessToken)
	response, err := util.PostJSONContext(ctx, uri, req)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &res, "IntracitySetPayMode")
	return
}

// IntracityGetPayModeRequest 查询扣费主体请求参数
type IntracityGetPayModeRequest struct {
	AppID string `json:"appid"` // 小程序appid。必须要和access_token匹配
}

// IntracityGetPayModeResponse 查询扣费主体返回参数
type IntracityGetPayModeResponse struct {
	util.CommonError
	PayMode           PayMode `json:"pay_mode"`                      // 扣费主体。门店：PAY_MODE_STORE；小程序:PAY_MODE_APP；服务商：PAY_MODE_COMPONENT
	PayAppID          string  `json:"pay_appid,omitempty"`           // 扣费appid。扣费主体为小程序时返回
	PayComponentAppID string  `json:"pay_component_appid,omitempty"` // 扣费component_appid。扣费主体为服务商时返回
}

// IntracityGetCityRequest 查询支持同城配送的城市请求参数
type IntracityGetCityRequest struct {
	ServiceTransID string `json:"service_trans_id,omitempty"` // 返回指定运力的支持城市信息。SFTC:顺丰同城，DADA:达达，不传则返回所有结果
}

// IntracityGetCityResponse 查询支持同城配送的城市返回参数
type IntracityGetCityResponse struct {
	util.CommonError
	SupportList []*IntracitySupportCity `json:"support_list"` // 运力支持的城市信息
}

// IntracitySupportCity 运力支持的城市信息
type IntracitySupportCity struct {
	ServiceTransID string               `json:"service_trans_id"` // 运力ID
	CityList       []*IntracityCityInfo `json:"city_list"`        // 城市列表
}

// IntracityCityInfo 同城配送城市信息
type IntracityCityInfo struct {
	CityCode int    `json:"city_code"` // 城市行政区域编码
	CityName string `json:"city_name"` // 城市名称
}

// IntracitySetPayModeRequest 设置扣费主体请求参数
type IntracitySetPayModeRequest struct {
	AppID   string  `json:"appid"`    // 小程序appid。必须要和access_token匹配
	PayMode PayMode `json:"pay_mode"` // 扣费主体。门店：PAY_MODE_STORE；小程序:PAY_MODE_APP；服务商：PAY_MODE_COMPONENT
}

// IntracitySetPayModeResponse 设置扣费主体返回参数
type IntracitySetPayModeResponse struct {
	util.CommonError
}
