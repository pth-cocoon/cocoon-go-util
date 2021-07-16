package stringutil

import (
	"math/big"
	"strconv"
)

// ReverseString 反转字符串
func ReverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

// ToString 转换到字符串，兼容类型：int,float,error,bool
func ToString(o interface{}) string {
	switch o.(type) {
	case int:
		return ToString(int64(o.(int)))
	case int8:
		return ToString(int64(o.(int8)))
	case int16:
		return ToString(int64(o.(int16)))
	case int32:
		return ToString(int64(o.(uint8)))
	case int64:
		return big.NewInt(o.(int64)).String()
	case float32:
		return strconv.FormatFloat(float64(o.(float32)), 'f', -1, 32)
	case float64:
		return big.NewFloat(o.(float64)).String()
	case error:
		return o.(error).Error()
	case bool:
		{
			if o == true {
				return "true"
			} else {
				return "false"
			}

		}
	}
	return ""
}
