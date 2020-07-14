package queue

import (
	"fmt"
	"testing"
	"time"
)

func TestSlic(t *testing.T) {
	ary1 := make([]int, 0, 100)

	t.Logf("%d", len(ary1))
	t.Logf("%p", ary1)
	ary2 := append(ary1, 0)
	for i := 0; i < 150; i++ {
		ary2 = append(ary2, i)
		//ary2[i] = i
		t.Logf("%p", ary2)
		t.Log(len(ary2))
	}

	ary3 := ary2[:10]
	t.Logf("ary3 : %p", ary3)

	ary4 := ary2[10:]
	t.Logf("ary4 : %p", ary4)

}

func callback(datas []interface{}) error {
	fmt.Println("回调  - ", datas)
	for i, v := range datas { //range 关键字
		//fmt.Println(i,v)
		fmt.Print(" - ", i, v)
	}
	return nil
}

func TestDelayQueuePush(t *testing.T) {

	q := NewDelayQueue(1550, callback)

	go func(q *DelayQueue) {
		for i := 0; i < 15000; i++ {
			q.Push(i)
			if i%150 == 0 {
				time.Sleep(time.Microsecond * 350)
			}
		}
	}(q)

	time.Sleep(time.Second * 10)
}
