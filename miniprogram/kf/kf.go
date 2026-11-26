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

package kf

import (
	"github.com/silenceper/wechat/v2/miniprogram/context"
)

// KF 小程序客服，包含客服子商户管理、客服消息与微信客服绑定相关接口
type KF struct {
	*context.Context
}

// NewKF 实例化
func NewKF(ctx *context.Context) *KF {
	return &KF{ctx}
}
