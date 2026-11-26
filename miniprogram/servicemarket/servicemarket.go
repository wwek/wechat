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

package servicemarket

import (
	"fmt"

	"github.com/silenceper/wechat/v2/miniprogram/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	// invokeServiceURL 调用服务市场接口
	invokeServiceURL = "https://api.weixin.qq.com/wxa/servicemarket?access_token=%s"
	// serviceMarketRetrieveURL 异步获取处理数据
	serviceMarketRetrieveURL = "https://api.weixin.qq.com/wxa/servicemarketretrieve?access_token=%s"
)

// ServiceMarket 微信服务市场
type ServiceMarket struct {
	*context.Context
}

// NewServiceMarket 实例化
func NewServiceMarket(ctx *context.Context) *ServiceMarket {
	return &ServiceMarket{ctx}
}

// InvokeServiceRequest 调用服务市场接口请求
type InvokeServiceRequest struct {
	Service     string `json:"service"`       // 服务 ID，在每个接口详情页面中均可以找到服务id
	API         string `json:"api"`           // 接口名；在每个接口详情页面中均可以找到api name
	Data        string `json:"data"`          // 服务提供方接口定义的 JSON 格式的数据，请看每个服务的接口文档详情
	ClientMsgID string `json:"client_msg_id"` // 随机字符串 ID，调用方请求的唯一标识
	Async       bool   `json:"async"`         // 是否是异步API，当是异步调用时，必须填true
}

// InvokeServiceResponse 调用服务市场接口响应
type InvokeServiceResponse struct {
	util.CommonError
	Data      string `json:"data"`       // 回包信息
	RequestID string `json:"request_id"` // 异步调用才会返回，唯一id
}

// InvokeService 调用服务市场接口
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/wx-service-market/api_invokeservice
func (s *ServiceMarket) InvokeService(in *InvokeServiceRequest) (out InvokeServiceResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(invokeServiceURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "InvokeService")
	return
}

// ServiceMarketRetrieveRequest 异步获取处理数据请求
type ServiceMarketRetrieveRequest struct {
	RequestID string `json:"request_id"` // 调用接口 invokeService 接口返回的request_id
}

// ServiceMarketRetrieveResponse 异步获取处理数据响应
type ServiceMarketRetrieveResponse struct {
	util.CommonError
	Data      string `json:"data"`       // 回包信息
	RequestID string `json:"request_id"` // 唯一id
}

// ServiceMarketRetrieve 异步获取处理数据
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/wx-service-market/api_servicemarketretrieve
func (s *ServiceMarket) ServiceMarketRetrieve(in *ServiceMarketRetrieveRequest) (out ServiceMarketRetrieveResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(serviceMarketRetrieveURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "ServiceMarketRetrieve")
	return
}
