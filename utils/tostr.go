package utils

import (
	"fmt"
	"reflect"
	"strings"
)

// 判断对象是否为空
func IsObjectNil(obj interface{}) bool {
	// 对象是否为空
	// 空引用
	if obj == nil {
		return true
	}
	// TODO 需要检查一下空指针的是否需要判断
	// 字符串类型判断
	val := fmt.Sprintf("%v", obj)
	if len(val) == 0 || val == "<nil>" {
		return true
	}
	return false
}

// 判断字符串是否为空或空字符串
func IsStrEmptyOrNull(str string) bool {
	if len(str) == 0 || strings.Trim(str, " ") == "" {
		return true
	}
	return false
}

// 任意类型转成字符串
func AnyToStr(obj interface{}) string {
	if obj == nil {
		return ""
	}
	val := fmt.Sprintf("%v", obj)
	if len(val) == 0 || val == "<nil>" {
		if val = reflect.TypeOf(obj).String(); len(val) == 0 {
			val = reflect.TypeOf(obj).Kind().String()
		}
	}
	return val
}

// 转换成字符串
// 代码保留 研究用
//func AnyToStr1(a interface{}) string {
//	b := ""
//
//	switch a.(type) {
//	case int:
//		b = strconv.FormatInt(int64(a.(int)), 10)
//	case int8:
//		b = strconv.FormatInt(int64(a.(int8)), 10)
//	case int16:
//		b = strconv.FormatInt(int64(a.(int16)), 10)
//	case int32:
//		b = strconv.FormatInt(int64(a.(int32)), 10)
//	case int64:
//		b = strconv.FormatInt(a.(int64), 10)
//	case float32:
//		b = fmt.Sprintf("%f", a)
//		//b = strconv.FormatFloat(float64(a.(float32)), 'f', -1, 64)
//	case float64:
//		b = fmt.Sprintf("%f", a)
//		b = strconv.FormatFloat(a.(float64), 'f', -1, 64)
//	case string:
//		b = a.(string)
//	default:
//		val := reflect.ValueOf(a)
//		switch strings.ToLower(reflect.TypeOf(a).Kind().String()) {
//		case "int","int8","int16","int32","int64":
//
//			vtype := reflect.TypeOf(int64(0))
//			r1 := val.Convert(vtype).Int()
//			b = strconv.FormatInt(r1, 10)
//
//		case "uint","uint8","uint16","uint32","uint64":
//
//			vtype := reflect.TypeOf(uint64(0))
//			r1 := val.Convert(vtype).Uint()
//			b = strconv.FormatUint(r1, 10)
//		case "string":
//
//			vtype := reflect.TypeOf(string(0))
//			r1 := val.Convert(vtype).String()
//			b = r1
//
//		default:
//			b = GetTypeName(a)
//		}
//	}
//	return b
//}
