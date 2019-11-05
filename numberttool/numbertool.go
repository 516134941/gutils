package arraytool

import (
	"strings"
)

// IsNumeric  检测变量是否为数字或数字字符串。支持小数点、十六进制（hex）、科学计数法
func IsNumeric(val interface{}) bool {
	switch val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
	case float32, float64, complex64, complex128:
		return true
	case string:
		str := val.(string)
		if str == "" {
			return false
		}
		// Trim any whitespace
		str = strings.Trim(str, " \\t\\n\\r\\v\\f")
		if str[0] == '-' || str[0] == '+' {
			if len(str) == 1 {
				return false
			}
			str = str[1:]
		}
		// hex
		if len(str) > 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X') {
			for _, h := range str[2:] {
				if !((h >= '0' && h <= '9') || (h >= 'a' && h <= 'f') || (h >= 'A' && h <= 'F')) {
					return false
				}
			}
			return true
		}
		// 0-9,Point,Scientific
		p, s, l := 0, 0, len(str)
		for i, v := range str {
			if v == '.' { // Point
				if p > 0 || s > 0 || i+1 == l {
					return false
				}
				p = i
			} else if v == 'e' || v == 'E' { // Scientific
				if i == 0 || s > 0 || i+1 == l {
					return false
				}
				s = i
			} else if v < '0' || v > '9' {
				return false
			}
		}
		return true
	}

	return false
}

// IsInclouldArray 比较一个数组是否包含在另一个数组
func IsInclouldArray(arry3, arry4 []interface{}) bool {
	flag := false
	k := 0

	if len(arry3) == 0 {
		return true
	}
	if len(arry4) == 0 {
		return true
	}
	//  len(arry1) > len(arry2)
	arry1, arry2 := arry3, arry4

	if len(arry3) < len(arry4) {
		arry1, arry2 = arry4, arry3
	}

	for i := 0; i < len(arry1); i++ {
		for j := 0; j < len(arry2); j++ {
			if arry1[i] == arry2[j] {
				k++
				continue
			}
		}
	}

	if k == len(arry2){
		flag = true
	}
	return flag
}
