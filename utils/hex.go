package utils

import (
	"errors"
	"strings"
)

const (
	hexStr = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// 十进制转任意进制
// val 十进制值，n 目标进制 2-62
func Oct2Any(val uint64, n int8) (string, error) {

	if n < 2 || n > 62 {
		return "", errors.New("n 的范围为 2-62 ")
	}

	remainder := uint64(0)
	id := make([]string, 32)
	for i := 1; val != 0; i++ {
		remainder = val % uint64(n)
		//fmt.Printf("v = %v, %v \r\n", remainder, val)
		str := string(hexStr[remainder])
		//fmt.Printf("v = %v \r\n", str)
		id[32-i] = str
		val = val / uint64(n)
	}
	return strings.Join(id, ""), nil
}
