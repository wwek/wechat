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

package delivery

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// updateOrderURL 更新配送单状态
	updateOrderURL = "https://api.weixin.qq.com/cgi-bin/express/local/delivery/update_order?access_token=%s"
)

// LocalAgent 骑手信息
type LocalAgent struct {
	Name             string  `json:"name"`               // 骑手姓名
	Phone            string  `json:"phone"`              // 骑手电话
	IsPhoneEncrypted int64   `json:"is_phone_encrypted"` // 电话是否加密。默认是0不加密，1表示加密
	Lng              float64 `json:"lng"`                // 骑手经度
	Lat              float64 `json:"lat"`                // 骑手纬度
}

// UpdateOrderRequest 更新配送单状态请求参数
type UpdateOrderRequest struct {
	WxToken              string      `json:"wx_token"`               // 下单事件中推送的token字段
	OrderStatus          int64       `json:"order_status"`           // 订单状态
	WaybillID            string      `json:"waybill_id"`             // 配送单id
	ActionMsg            string      `json:"action_msg"`             // 附加信息
	ActionTime           int64       `json:"action_time"`            // 状态变更时间点，Unix秒级时间戳
	Agent                *LocalAgent `json:"agent"`                  // 骑手信息, 骑手接单时需返回
	ShopID               string      `json:"shopid"`                 // 商家id，由配送公司分配，可以是dev_id或者appkey
	ShopOrderID          string      `json:"shop_order_id"`          // 唯一标识订单的 ID，由商户生成
	ShopNo               string      `json:"shop_no"`                // 商家门店编号，在配送公司侧登记
	WxaPath              string      `json:"wxa_path"`               // 配送公司小程序跳转路径，用于用户收到消息会间接跳转到这个页面
	ExpectedDeliveryTime int64       `json:"expected_delivery_time"` // 预计送达时间戳，骑手接单时需返回
}

// UpdateOrderResponse 更新配送单状态返回参数
type UpdateOrderResponse struct {
	util.CommonError
	ResultCode int64  `json:"resultcode"` // 错误码
	ResultMsg  string `json:"resultmsg"`  // 错误描述
}

// UpdateOrder 更新配送单状态
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/immediate-delivery/deliver-by-provider/api_updateorder
func (s *Delivery) UpdateOrder(in *UpdateOrderRequest) (out UpdateOrderResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(updateOrderURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &out, "UpdateOrder")
	return
}
