package utils

import (
	"fmt"
	"go/types"
	"reflect"
	"strconv"
	"testing"
)

type (
	mint int
	nint int64
	uInt uint
	nstr string
	aint *int
	astr *string
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
	var f8 []*int
	var f9 *int
	var f10 aint = &a
	var f11 string = string(n)
	var f12 *string = &f11
	var f13 astr = &f11
	var f14 *astr

	//fmt.Printf("obj = %#v , %#v", *f10, reflect.ValueOf(f10).Kind())

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
	t.Log(AnyToStr(f8))
	t.Log(AnyToStr(f9))
	t.Log(AnyToStr(f10))
	t.Log(AnyToStr(f11))
	t.Log(AnyToStr(f12))
	t.Log(AnyToStr(f13))
	t.Log(AnyToStr(f14))

	//t.Errorf("f7 = %#v", AnyToStr(f7))
	//t.Fatal(AnyToStr(f7))
	//t.Skip(f7)
}

func TestIsStrEmptyOrNull(t *testing.T) {

	a1 := ""
	a2 := "\n"
	a3 := " "
	a4 := "\r"
	a5 := "\t"
	a6 := "\v"
	a7 := "\f"

	t.Logf("a1 = %#v : %#v", a1, IsStrEmptyOrNull(a1))
	t.Logf("a2 = %#v : %#v", a2, IsStrEmptyOrNull(a2))
	t.Logf("a3 = %#v : %#v", a3, IsStrEmptyOrNull(a3))
	t.Logf("a4 = %#v : %#v", a4, IsStrEmptyOrNull(a4))
	t.Logf("a5 = %#v : %#v", a5, IsStrEmptyOrNull(a5))
	t.Logf("a6 = %#v : %#v", a6, IsStrEmptyOrNull(a6))
	t.Logf("a7 = %#v : %#v", a7, IsStrEmptyOrNull(a7))

}

func TestIsObjectNil(t *testing.T) {
	var a1 int
	var a2 *int
	var a3 mint
	var a4 aint
	var a5 string
	var a6 *string = &a5
	var a7 astr
	var a8 interface{} = 9
	var a9 types.Object

	t.Logf("a1 = %#v : %#v", a1, IsObjectNil(a1))
	t.Logf("a2 = %#v : %#v", a2, IsObjectNil(a2))
	t.Logf("a3 = %#v : %#v", a3, IsObjectNil(a3))
	t.Logf("a4 = %#v : %#v", a4, IsObjectNil(a4))
	t.Logf("a5 = %#v : %#v", a5, IsObjectNil(a5))
	t.Logf("a6 = %#v : %#v", a6, IsObjectNil(a6))
	t.Logf("a7 = %#v : %#v", a7, IsObjectNil(a7))
	t.Logf("a8 = %#v : %#v", a8, IsObjectNil(a8))
	t.Logf("a9 = %#v : %#v", a9, IsObjectNil(a9))
}

func BenchmarkAnyToStr(b *testing.B) {
	var u1 uInt = 9845
	for i := 0; i < 1000000; i++ {
		AnyToStr(u1)
	}
}

func ExampleAnyToStr() {
	var u1 uInt = 9845
	var f chan int
	fmt.Println(AnyToStr(u1))
	fmt.Println(AnyToStr(f))
	// Output:
	// 9845
	// chan int
}
