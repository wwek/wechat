package tcb

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// 延时调用云函数
	addDelayedFunctionTaskURL = "https://api.weixin.qq.com/tcb/adddelayedfunctiontask"
)

// AddDelayedFunctionTaskRequest 延时调用云函数 请求参数
type AddDelayedFunctionTaskRequest struct {
	Env          string `json:"env,omitempty"`           // 环境ID
	FunctionName string `json:"function_name,omitempty"` // 函数名称
	Data         string `json:"data,omitempty"`          // 发送的数据包，格式必须为JSONString
	DelayTime    int64  `json:"delay_time,omitempty"`    // 延迟时间，单位：秒，合法范围：6秒-30天
}

// AddDelayedFunctionTaskResponse 延时调用云函数 返回结果
type AddDelayedFunctionTaskResponse struct {
	util.CommonError
}

// AddDelayedFunctionTask 延时调用云函数
//
// reference:https://developers.weixin.qq.com/miniprogram/dev/server/API/cloudbase/functions/api_adddelayedfunctiontask
func (s *Tcb) AddDelayedFunctionTask(in *AddDelayedFunctionTaskRequest) (out AddDelayedFunctionTaskResponse, err error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", addDelayedFunctionTaskURL, accessToken)
	response, err := util.PostJSON(uri, in)
	if err != nil {
		return
	}
	err = util.DecodeWithError(response, &out, "AddDelayedFunctionTask")
	return
}
