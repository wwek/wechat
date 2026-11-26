package util

import (
	"strconv"
	"strings"
)

// 微信部分接口的同一字段，在官方文档的参数表与返回示例中类型并不一致
// （例如 business_id 参数表写 string、返回示例写数字）。直接按其中一种建模，
// 一旦与线上实际返回不符，json.Unmarshal 会直接报错并丢弃整个响应。
// 下面几个类型用于兼容这类字段：解析失败不会导致整个响应解码失败。

// FlexString 兼容微信返回字符串或数字的字段
type FlexString string

// UnmarshalJSON 解析字符串或数字
func (v *FlexString) UnmarshalJSON(data []byte) error {
	s := strings.Trim(strings.TrimSpace(string(data)), `"`)
	if s == "null" {
		s = ""
	}
	*v = FlexString(s)
	return nil
}

// String 转为字符串
func (v FlexString) String() string {
	return string(v)
}

// FlexInt64 兼容微信返回数字或字符串的字段
type FlexInt64 int64

// UnmarshalJSON 解析数字或字符串
func (v *FlexInt64) UnmarshalJSON(data []byte) error {
	s := strings.Trim(strings.TrimSpace(string(data)), `"`)
	if s == "" || s == "null" {
		*v = 0
		return nil
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	*v = FlexInt64(int64(f))
	return nil
}

// Int64 转为 int64
func (v FlexInt64) Int64() int64 {
	return int64(v)
}

// FlexBool 兼容微信返回布尔、数字或字符串的字段
type FlexBool bool

// UnmarshalJSON 解析布尔、数字或字符串
func (v *FlexBool) UnmarshalJSON(data []byte) error {
	switch strings.ToLower(strings.Trim(strings.TrimSpace(string(data)), `"`)) {
	case "true", "1":
		*v = true
	case "false", "0", "", "null":
		*v = false
	default:
		f, err := strconv.ParseFloat(strings.Trim(strings.TrimSpace(string(data)), `"`), 64)
		if err != nil {
			return err
		}
		*v = f != 0
	}
	return nil
}

// Bool 转为 bool
func (v FlexBool) Bool() bool {
	return bool(v)
}
