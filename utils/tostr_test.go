package utils

import (
	"reflect"
	"strconv"
	"testing"
)

type (
	mint int
	nint int64
	uInt uint
	nstr string
)

func TestAnyToStr2(t *testing.T) {

	var g mint = 9

	var h interface{} = 5

	var j nint

	var k = strconv.FormatInt(int64(h.(int)), 10)

	j = nint(g)

	ttt := reflect.TypeOf(5)
	r := reflect.ValueOf(j)
	r1 := r.Convert(ttt)

	t.Log(k, j, reflect.TypeOf(j), r1)
}

func TestAnyToStr(t *testing.T) {
	var a int = 1
	var b int8 = 4
	var c int16 = 5
	var d int32 = 6
	var e int64 = 7
	var f chan int
	var g mint = 9
	var h nint = 98
	var n nstr = "fdf"
	var u1 uInt = 9845
	var f1 float32 = 999.99
	var f2 func(...uInt) (<-chan int, chan<- bool, error)
	var f3 map[nstr]interface{}
	var f4 struct{}
	var f5 reflect.Type
	var f6 []int
	var f7 [3]int

	t.Log(AnyToStr(a))
	t.Log(AnyToStr(b))
	t.Log(AnyToStr(c))
	t.Log(AnyToStr(d))
	t.Log(AnyToStr(e))
	t.Log(AnyToStr(f))
	t.Log(AnyToStr(g))
	t.Log(AnyToStr(h))
	t.Log(AnyToStr(n))
	t.Log(AnyToStr(u1))
	t.Log(AnyToStr(f1))
	t.Log(AnyToStr(f2))
	t.Log(AnyToStr(f3))
	t.Log(AnyToStr(f4))
	t.Log(AnyToStr(f5))
	t.Log(AnyToStr(f6))
	t.Log(AnyToStr(f7))
}
