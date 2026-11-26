package express

import (
	"context"
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// 物流轨迹推送 API URL
const (
	// 推送已绑定物流轨迹信息
	deliveryPathNotifyURL = "https://api.weixin.qq.com/cgi-bin/express/delivery/pathnotify?access_token=%s"
	// 用户手机状态查询
	deliveryUserQueryURL = "https://api.weixin.qq.com/cgi-bin/express/delivery/userquery?access_token=%s"
)

// DeliveryPathNotify 推送已绑定物流轨迹信息
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/express/business/express_open_msg.html
func (express *Express) DeliveryPathNotify(ctx context.Context, req *DeliveryPathNotifyRequest) (res DeliveryPathNotifyResponse, err error) {
	accessToken, err := express.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(deliveryPathNotifyURL, accessToken)
	response, err := util.PostJSONContext(ctx, uri, req)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &res, "DeliveryPathNotify")
	return
}

// DeliveryUserQuery 用户手机状态查询
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/express/business/express_open_msg.html
func (express *Express) DeliveryUserQuery(ctx context.Context, req *DeliveryUserQueryRequest) (res DeliveryUserQueryResponse, err error) {
	accessToken, err := express.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(deliveryUserQueryURL, accessToken)
	response, err := util.PostJSONContext(ctx, uri, req)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &res, "DeliveryUserQuery")
	return
}

// DeliveryContact 物流寄件人与收件人信息
type DeliveryContact struct {
	Name     string `json:"name"`             // 姓名（收件人必传，寄件人非必传）
	Phone    string `json:"phone,omitempty"`  // 电话（收件人必传，寄件人非必传）
	Province string `json:"province"`         // 省
	City     string `json:"city"`             // 市
	Area     string `json:"area"`             // 区/县
	Street   string `json:"street,omitempty"` // 街道
	Address  string `json:"address,omitempty"`
	ID       string `json:"id,omitempty"` // 地址id
}

// DeliveryPath 物流轨迹信息
type DeliveryPath struct {
	ActionTime           int64  `json:"action_time"`                      // 轨迹变化Unix时间戳
	ActionType           int    `json:"action_type"`                      // 轨迹变化类型，与普通单保持一致，参见下方其他说明action_type定义
	ActionMsg            string `json:"action_msg"`                       // 轨迹变化具体信息说明，展示在快递轨迹详情页中
	PickupCourierName    string `json:"pickup_courier_name,omitempty"`    // 取件员姓名,当分配取件员成功时返回
	PickupCourierPhone   string `json:"pickup_courier_phone,omitempty"`   // 取件员电话,当分配取件员成功时返回
	DeliveryCourierName  string `json:"delivery_courier_name,omitempty"`  // 派件员姓名,当分配派件员成功时返回
	DeliveryCourierPhone string `json:"delivery_courier_phone,omitempty"` // 派件员电话,当分配派件员成功时返回
}

// DeliveryPathNotifyRequest 推送已绑定物流轨迹信息请求参数
type DeliveryPathNotifyRequest struct {
	Sender     *DeliveryContact `json:"sender"`      // 寄件人
	Receiver   *DeliveryContact `json:"receiver"`    // 收件人
	WaybillID  string           `json:"waybill_id"`  // 运单号
	Path       *DeliveryPath    `json:"path"`        // 当前需要推送消息的轨迹
	CreateTime int64            `json:"create_time"` // 运单创建时间，unix时间戳（秒）
}

// DeliveryPathNotifyResponse 推送已绑定物流轨迹信息返回参数
type DeliveryPathNotifyResponse struct {
	util.CommonError
	Exist int `json:"exist"` // 用户是否绑定 0-未绑定 1-已绑定
}

// DeliveryUserQueryRequest 用户手机状态查询请求参数
type DeliveryUserQueryRequest struct {
	Phone string `json:"phone"` // 手机号
}

// DeliveryUserQueryResponse 用户手机状态查询返回参数
type DeliveryUserQueryResponse struct {
	util.CommonError
	Exist int `json:"exist"` // 用户是否绑定 0-未绑定 1-已绑定
}
