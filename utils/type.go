package utils

import "reflect"

func GetType(val interface{}) reflect.Type {

	return reflect.TypeOf(val)
}
func GetTypeName(val interface{}) string {

	return reflect.TypeOf(val).Kind().String()
}

func GetTypeFullName(val interface{}) (string, string, string) {
	t := reflect.TypeOf(val)
	return t.Kind().String(), t.String(), t.PkgPath()
}
