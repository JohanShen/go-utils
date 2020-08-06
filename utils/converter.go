package utils

import "github.com/vmihailenco/msgpack/v5"
import "unsafe"

// 字符串转字节组
func StringToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// 字节组转字符串
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func StructToBytes(obj interface{}) ([]byte, error) {
	val, err := msgpack.Marshal(obj)
	if err != nil {
		return []byte{}, err
	}
	return val, nil
}

func BytesToStruct(bytes []byte, obj interface{}) error {
	return msgpack.Unmarshal(bytes, &obj)
}

// struct 转 bytes 代码备用研究
// 用 unsafe 方法实现
//func MyStructToBytes(s interface{}) []byte {
//	var sizeOfMyStruct = int(unsafe.Sizeof(interface{}))
//	var x reflect.SliceHeader
//	x.Len = sizeOfMyStruct
//	x.Cap = sizeOfMyStruct
//	x.Data = uintptr(unsafe.Pointer(s))
//	return *(*[]byte)(unsafe.Pointer(&x))
//}
//
//func BytesToMyStruct(b []byte) *MyStruct {
//	return (*MyStruct)(unsafe.Pointer(
//		(*reflect.SliceHeader)(unsafe.Pointer(&b)).Data,
//	))
//}
