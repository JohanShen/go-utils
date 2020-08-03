package utils

import "strings"

// 获取字符串在字符串数组中的位置
// -1 表示字符串不在数组中
func IndexOf(ary []string, str string, isMatchCase bool) (index int) {
	for i, v := range ary {
		if v == str || (!isMatchCase && strings.ToUpper(v) == strings.ToUpper(str)) {
			return i
		}
	}
	return -1
}

// 获取字符串在字符串数组中的位置
//  区分大小写
//	-1 表示字符串不在数组中
func IndexOfMatchCase(ary []string, str string) (index int) {
	return IndexOf(ary, str, true)
}

// 获取字符串在字符串数组中的位置
//  不区分大小写
//	-1 表示字符串不在数组中
func IndexOfWithoutCase(ary []string, str string) (index int) {
	return IndexOf(ary, str, false)
}
