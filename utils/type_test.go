package utils

import (
	"errors"
	"testing"
)

type diy struct {
	A int
}

type Iface interface{}

type Int int

var aInt int
var aInt32 int32
var bString string
var cMap map[string]interface{}
var dObj interface{}
var eDiy diy
var fChan chan int
var gArray []chan int
var hArray [1]chan int
var iInt Int
var jFun func()
var kFun func(string) string
var lErr error
var mInterface Iface
var nFloat float32
var oBool bool

func init() {
	dObj = cMap
	lErr = errors.New("")
	mInterface = new(Iface)
}

func TestGetType(t *testing.T) {
	t.Log(GetType(aInt))
	t.Log(GetType(aInt32))
	t.Log(GetType(bString))
	t.Log(GetType(cMap))
	t.Log(GetType(dObj))
	t.Log(GetType(eDiy))
	t.Log(GetType(fChan))
	t.Log(GetType(gArray))
	t.Log(GetType(hArray))
	t.Log(GetType(iInt))
	t.Log(GetType(jFun))
	t.Log(GetType(kFun))
	t.Log(GetType(lErr))
	t.Log(GetType(mInterface))
	t.Log(GetType(nFloat))
	t.Log(GetType(oBool))
}

func TestGetTypeFullName(t *testing.T) {
	t.Log(GetTypeFullName(aInt))
	t.Log(GetTypeFullName(aInt32))
	t.Log(GetTypeFullName(bString))
	t.Log(GetTypeFullName(cMap))
	t.Log(GetTypeFullName(dObj))
	t.Log(GetTypeFullName(eDiy))
	t.Log(GetTypeFullName(fChan))
	t.Log(GetTypeFullName(gArray))
	t.Log(GetTypeFullName(hArray))
	t.Log(GetTypeFullName(iInt))
	t.Log(GetTypeFullName(jFun))
	t.Log(GetTypeFullName(kFun))
	t.Log(GetTypeFullName(lErr))
	t.Log(GetTypeFullName(mInterface))
	t.Log(GetTypeFullName(nFloat))
	t.Log(GetTypeFullName(oBool))
}

func TestGetTypeName(t *testing.T) {
	t.Log(GetTypeName(aInt))
	t.Log(GetTypeName(aInt32))
	t.Log(GetTypeName(bString))
	t.Log(GetTypeName(cMap))
	t.Log(GetTypeName(dObj))
	t.Log(GetTypeName(eDiy))
	t.Log(GetTypeName(fChan))
	t.Log(GetTypeName(gArray))
	t.Log(GetTypeName(hArray))
	t.Log(GetTypeName(iInt))
	t.Log(GetTypeName(jFun))
	t.Log(GetTypeName(kFun))
	t.Log(GetTypeName(lErr))
	t.Log(GetTypeName(mInterface))
	t.Log(GetTypeName(nFloat))
	t.Log(GetTypeName(oBool))
}
