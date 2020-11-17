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
	// 指针类型 判断指向的内存是否为空
	vof := reflect.ValueOf(obj)
	if vof.Kind() == reflect.Ptr {
		//fmt.Printf("obj = %v, %v, %#v", vof.IsNil(), vof.IsZero(),reflect.ValueOf(obj).Elem())
		//|| vof.Elem().Len()==0
		if vof.IsNil() || vof.IsZero() {
			return true
		}
	} else {
		// 字符串类型判断
		val := fmt.Sprintf("%v", obj)
		if val == "<nil>" {
			return true
		}
	}
	return false
}

var spaceChar = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}

// 判断字符串是否为空或空字符串
func IsStrEmptyOrNull(str string) bool {
	if len(str) == 0 {
		return true
	}
	for _, v := range spaceChar {
		str = strings.Trim(str, string(v))
		//fmt.Printf("V = %#v , %#v \n",string(v), v)
	}
	return len(str) == 0
}

// 任意类型转成字符串
func AnyToStr(obj interface{}) string {
	if obj == nil {
		return ""
	}
	val := ""
	vof := reflect.ValueOf(obj)
	if vof.Kind() == reflect.Ptr {
		if vof.IsNil() || vof.IsZero() {
			val = fmt.Sprint("")
			//fmt.Printf("obj = %v, %v, %#v", vof.IsNil(), vof.IsZero(),reflect.ValueOf(obj).Elem())
		} else {
			val1 := vof.Elem()
			val = fmt.Sprintf("%v", val1)
		}
	} else {
		val = fmt.Sprintf("%v", obj)

		if len(val) == 0 || val == "<nil>" {
			if val = reflect.TypeOf(obj).String(); len(val) == 0 {
				val = reflect.TypeOf(obj).Kind().String()
			}
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
//		str := reflect.ValueOf(a)
//		switch strings.ToLower(reflect.TypeOf(a).Kind().String()) {
//		case "int","int8","int16","int32","int64":
//
//			vtype := reflect.TypeOf(int64(0))
//			r1 := str.Convert(vtype).Int()
//			b = strconv.FormatInt(r1, 10)
//
//		case "uint","uint8","uint16","uint32","uint64":
//
//			vtype := reflect.TypeOf(uint64(0))
//			r1 := str.Convert(vtype).Uint()
//			b = strconv.FormatUint(r1, 10)
//		case "string":
//
//			vtype := reflect.TypeOf(string(0))
//			r1 := str.Convert(vtype).String()
//			b = r1
//
//		default:
//			b = GetTypeName(a)
//		}
//	}
//	return b
//}
