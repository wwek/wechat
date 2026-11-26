package express

import (
	"context"
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// 无忧退货 API URL
const (
	// 创建退货ID
	noWorryReturnAddURL = "https://api.weixin.qq.com/cgi-bin/express/delivery/no_worry_return/add?access_token=%s"
	// 查询退货ID状态
	noWorryReturnGetURL = "https://api.weixin.qq.com/cgi-bin/express/delivery/no_worry_return/get?access_token=%s"
	// 解绑退货ID
	noWorryReturnUnbindURL = "https://api.weixin.qq.com/cgi-bin/express/delivery/no_worry_return/unbind?access_token=%s"
)

// AddReturnID 创建退货ID
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/express/business/insurance_freight.html
func (express *Express) AddReturnID(ctx context.Context, req *AddReturnIDRequest) (res AddReturnIDResponse, err error) {
	accessToken, err := express.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(noWorryReturnAddURL, accessToken)
	response, err := util.PostJSONContext(ctx, uri, req)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &res, "AddReturnID")
	return
}

// GetReturnID 查询退货ID状态
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/express/business/insurance_freight.html
func (express *Express) GetReturnID(ctx context.Context, req *GetReturnIDRequest) (res GetReturnIDResponse, err error) {
	accessToken, err := express.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(noWorryReturnGetURL, accessToken)
	response, err := util.PostJSONContext(ctx, uri, req)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &res, "GetReturnID")
	return
}

// UnbindReturnID 解绑退货ID
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/express/business/insurance_freight.html
func (express *Express) UnbindReturnID(ctx context.Context, req *UnbindReturnIDRequest) (res UnbindReturnIDResponse, err error) {
	accessToken, err := express.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf(noWorryReturnUnbindURL, accessToken)
	response, err := util.PostJSONContext(ctx, uri, req)
	if err != nil {
		return
	}

	// 使用通用方法返回错误
	err = util.DecodeWithError(response, &res, "UnbindReturnID")
	return
}

// ReturnAddress 退货地址信息
type ReturnAddress struct {
	Name     string `json:"name"`      // 发/收件人姓名，不超过64字节
	Tel      string `json:"tel"`       // 发/收件人座机号码，若不填写则必须填写 mobile，不超过32字节
	Mobile   string `json:"mobile"`    // 发/收件人手机号码，若不填写则必须填写 tel，不超过32字节
	Company  string `json:"company"`   // 发/收件人公司名称，不超过64字节
	PostCode string `json:"post_code"` // 发/收件人邮编，不超过10字节
	Country  string `json:"country"`   // 发/收件人国家，不超过64字节
	Province string `json:"province"`  // 发/收件人省份，比如："广东省"，不超过64字节
	City     string `json:"city"`      // 发/收件人市/地区，比如："广州市"，不超过64字节
	Area     string `json:"area"`      // 发/收件人区/县，比如："海珠区"，不超过64字节
	Address  string `json:"address"`   // 发/收件人详细地址，比如："XX路XX号XX大厦XX"，不超过512字节
}

// ReturnGoods 退货商品信息
type ReturnGoods struct {
	Name string `json:"name"` // 退货商品的名称
	URL  string `json:"url"`  // 退货商品图片的url
}

// AddReturnIDRequest 创建退货ID请求参数
type AddReturnIDRequest struct {
	ShopOrderID string         `json:"shop_order_id"`       // 商家内部系统使用的退货编号
	BizAddr     *ReturnAddress `json:"biz_addr"`            // 商家退货地址
	UserAddr    *ReturnAddress `json:"user_addr,omitempty"` // 用户购物时的收货地址
	OpenID      string         `json:"openid"`              // 退货用户的openid
	OrderPath   string         `json:"order_path"`          // 退货订单在商家小程序的path。如投保时已传入订单商品信息，则以投保时传入的为准
	GoodsList   []*ReturnGoods `json:"goods_list"`          // 退货商品list，一个元素为对象的数组。如投保时已传入订单商品信息，则以投保时传入的为准
	OrderPrice  int64          `json:"order_price"`         // 退货订单的价格
	WxPayID     string         `json:"wx_pay_id"`           // 填写已投保的微信支付单号
}

// AddReturnIDResponse 创建退货ID返回参数
type AddReturnIDResponse struct {
	util.CommonError
	ReturnID string `json:"return_id"` // 退货ID
}

// GetReturnIDRequest 查询退货ID状态请求参数
type GetReturnIDRequest struct {
	ReturnID string `json:"return_id"` // 退货ID
}

// GetReturnIDResponse 查询退货ID状态返回参数
type GetReturnIDResponse struct {
	util.CommonError
	Status       int    `json:"status"`        // 退货方式，0.用户未填写退货信息 1.预约上门取件 2.填写自行寄回运单号
	WaybillID    string `json:"waybill_id"`    // 运单号
	OrderStatus  int    `json:"order_status"`  // 物流信息，0.已下单待揽件 1.已揽件 2.运输中 3.派件中 4.已签收 5.异常 6.代签收 7.揽收失败 8.签收失败（拒收，超区） 11.已取消 13.退件中 14.已退件 99.未知
	DeliveryID   string `json:"delivery_id"`   // 运力公司编码
	DeliveryName string `json:"delivery_name"` // 运力公司名称
}

// UnbindReturnIDRequest 解绑退货ID请求参数
type UnbindReturnIDRequest struct {
	ReturnID string `json:"return_id"` // 退货ID
}

// UnbindReturnIDResponse 解绑退货ID返回参数
type UnbindReturnIDResponse struct {
	util.CommonError
}
