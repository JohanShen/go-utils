package utils

import (
	"strings"
	"testing"
)

var arys = []string{"a", "aa", "abca", "A", "aA", "aBca"}

type testdata struct {
	str   string
	mcase bool
	value int
}

func TestFixedLeft1(t *testing.T) {

	type ttt struct {
		str1 string
		str2 string
	}
	b := byte('*')
	l := 5
	data := make([]ttt, 0, 10)
	strs := []string{"", "a", "中文a", "テスト", "provë", "îmtîhan", "परीक्षा", "פּרובירן", "փորձարկում", "పరీక్ష", "ทดสอบ"}
	for _, str := range strs {
		//ol := len(str)
		//bl := len(StringToBytes(str))
		rl := len([]rune(str))
		//t.Log(ol,bl, rl)
		l1 := l - rl

		str1 := str
		if l1 > 0 {
			str1 = str + strings.Repeat(string(b), l1)
		} else {
			str1 = string([]rune(str1)[0:l]) //str1[0:l]
		}
		data = append(data, ttt{str, str1})
	}

	for i, d := range data {
		val := FixedLeft(d.str1, b, l)
		t.Log(i, d.str2, val, d.str2 == val)
	}

}

func TestFixedLeft(t *testing.T) {

	type ttt struct {
		str1 string
		str2 string
	}
	b := byte('*')
	l := 5
	data := make([]ttt, 0, 10)
	strs := []string{"", "a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa"}
	for _, str := range strs {
		l1 := l - len(str)
		str1 := str
		if l1 > 0 {
			str1 = str + strings.Repeat(string(b), l1)
		} else {
			str1 = str1[0:l]
		}
		data = append(data, ttt{str, str1})
	}

	for i, d := range data {
		val := FixedRight(d.str1, b, l)
		t.Log(i, d.str2, val, d.str2 == val)
	}

}

func TestFixedRight(t *testing.T) {

	type ttt struct {
		str1 string
		str2 string
	}
	b := byte('*')
	l := 5
	data := make([]ttt, 0, 10)
	strs := []string{"", "a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa"}
	for _, str := range strs {
		l1 := l - len(str)
		str1 := str
		if l1 > 0 {
			str1 = strings.Repeat(string(b), l1) + str
		} else {
			l2 := len(str) - l
			str1 = str1[l2:]
		}
		data = append(data, ttt{str, str1})
	}

	for i, d := range data {
		val := FixedLeft(d.str1, b, l)
		t.Log(i, d.str2, val, d.str2 == val)
	}
}

func TestIndexOf(t *testing.T) {

	data := make([]*testdata, 0)
	data = append(data, &testdata{"a", false, 0})
	data = append(data, &testdata{"aa", false, 1})
	data = append(data, &testdata{"AA", false, 1})
	data = append(data, &testdata{"aA", false, 1})
	data = append(data, &testdata{"A", true, 3})
	data = append(data, &testdata{"AA", true, -1})
	data = append(data, &testdata{"aA", true, 4})

	t.Log(arys)

	for _, v := range data {
		if a := IndexOf(arys, v.str, v.mcase); a != v.value {
			t.Error("期待值不匹配", v.str, v.value, a)
		} else {
			t.Log(v.str, v.value, a)
		}
	}
	t.Log("全部测试通过")
}

func TestIndexOfMatchCase(t *testing.T) {

	data := make([]*testdata, 0)
	data = append(data, &testdata{"a", true, 0})
	data = append(data, &testdata{"aa", true, 1})
	data = append(data, &testdata{"AA", true, -1})
	data = append(data, &testdata{"aA", true, 4})
	data = append(data, &testdata{"A", true, 3})
	data = append(data, &testdata{"abca", true, 2})
	data = append(data, &testdata{"aA", true, 4})

	t.Log(arys)

	for _, v := range data {
		if a := IndexOfMatchCase(arys, v.str); a != v.value {
			t.Error("期待值不匹配", v.str, v.value, a)
		} else {
			t.Log(v.str, v.value, a)
		}
	}
	t.Log("全部测试通过")
}

func TestIndexOfWithoutCase(t *testing.T) {

	data := make([]*testdata, 0)
	data = append(data, &testdata{"a", true, 0})
	data = append(data, &testdata{"aa", true, 1})
	data = append(data, &testdata{"AA", true, 1})
	data = append(data, &testdata{"aA", true, 1})
	data = append(data, &testdata{"A", true, 0})
	data = append(data, &testdata{"abca", true, 2})
	data = append(data, &testdata{"aA", true, 1})

	t.Log(arys)

	for _, v := range data {
		if a := IndexOfWithoutCase(arys, v.str); a != v.value {
			t.Error("期待值不匹配", v.str, v.value, a)
		} else {
			t.Log(v.str, v.value, a)
		}
	}
	t.Log("全部测试通过")
}
