package queue

import (
	"fmt"
	"testing"
	"time"
)

func TestSlic(t *testing.T) {

	var delay DelayQueue

	t.Log("对象", delay)
	t.Logf("%p", delay)

	ary1 := make([]int, 0, 100)

	t.Logf("%d", len(ary1))
	t.Logf("%p", ary1)
	ary2 := append(ary1, 0)
	for i := 0; i < 10; i++ {
		ary2 = append(ary2, i)
		//ary2[i] = i
		t.Logf("%p", ary2)
		t.Log(len(ary2))
	}

	ary3 := ary2[:10]
	t.Logf("ary3 : %p", ary3)

	ary4 := ary2[10:]
	t.Logf("ary4 : %p", ary4)

	var a string = "123"

	b := a
	fmt.Println("a addr = ", &a)
	fmt.Println("b addr = ", &b)

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

	c := make(chan int, 10)

	go func(q *DelayQueue) {
		for i := 0; i < 15000; i++ {
			q.Push(i)
			c <- i
			if i%150 == 0 {
				time.Sleep(time.Microsecond * 350)
			}
		}
	}(q)

	go func() {
		//for i := range c {
		//	fmt.Println(i)
		//}
		fmt.Println("length :", len(c))
		for {
			if len(c) == 0 {
				time.Sleep(time.Microsecond * 100)
			}
			fmt.Println(<-c)
		}
	}()

	time.Sleep(time.Second * 10)
}
