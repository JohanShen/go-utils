package utils

import (
	"hash/crc32"
	"strings"
)

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

// 向右靠齐固定字符长度
// 不支持中文
//func FixedRight(str string, s byte, l int) string {
//
//	rs := string(s)
//	smask := strings.Repeat(rs, l)
//	if len(str)==0 {
//		return smask
//	}
//	n := StringToBytes(smask)
//	copy(n, str)
//	return BytesToString(n)
//
//}

// 向左靠齐固定字符长度
// 一个中文算三位长度
//func FixedLeft(str string, s byte, l int) string {
//
//	rs := string(s)
//	if len(str)==0 {
//		return strings.Repeat(rs, l)
//	}
//	r := l / len(str)
//	if len(str)%l != 0 {
//		r++
//	}
//	n1 := StringToBytes(strings.Repeat(str, r))
//	if len(n1) > l {
//		n1 = n1[len(n1)-l:]
//	}
//	left1 := l - len(str)
//	if left1 > 0 {
//		n2 := StringToBytes(strings.Repeat(rs, left1))
//		copy(n1, n2)
//	}
//	return BytesToString(n1)
//
//}

// 向右靠齐固定字符长度
func FixedRight(str string, s byte, l int) string {

	rs := string(s)
	smask := strings.Repeat(rs, l)
	if len(str) == 0 {
		return smask
	}
	n := StringToBytes(smask)
	rstr := []rune(str)
	//copy(n, str)
	sb := strings.Builder{}
	for i := 0; i < len(n); i++ {
		if i >= len(rstr) {
			sb.WriteByte(n[i])
		} else {
			sb.WriteRune(rstr[i])
		}
	}
	return sb.String()

}

// 向左靠齐固定字符长度
func FixedLeft(str string, s byte, l int) string {

	rs := string(s)
	if len(str) == 0 {
		return strings.Repeat(rs, l)
	}
	rstr := []rune(str)
	r := l / len(rstr)
	if len(rstr)%l != 0 {
		r++
	}
	n1 := StringToBytes(strings.Repeat(str, r))
	if len(n1) > l {
		n1 = n1[len(n1)-l:]
	}
	left1 := l - len(rstr)
	sb := strings.Builder{}
	if left1 > 0 {
		n2 := StringToBytes(strings.Repeat(rs, left1))
		//copy(n1, n2)

		for i := 0; i < len(n1); i++ {
			//fmt.Println(i, left1, l, len(rstr))
			if i >= left1 {
				sb.WriteRune(rstr[i-len(n2)])
			} else {
				sb.WriteByte(n2[i])
			}
		}
	} else {
		sb.Write(n1)
	}
	return sb.String()

}

// 获得季字符串的哈希
func HashCode(s string) uint32 {
	v := crc32.ChecksumIEEE(StringToBytes(s))
	return v
}
