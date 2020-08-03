package utils

import "testing"

var arys = []string{"a", "aa", "abca", "A", "aA", "aBca"}

type testdata struct {
	str   string
	mcase bool
	value int
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
