package main

import (
	"fmt"
	"time"
	"utils/queue"
)

func main() {
	fmt.Println(1)

	q := queue.NewDelayQueue(1550, callback)
	go func(q *queue.DelayQueue) {
		for i := 0; i < 150; i++ {
			q.Push(i)
		}
	}(q)

	time.Sleep(time.Second * 10)

}

func callback(datas []interface{}) error {
	for i, v := range datas { //range 关键字
		//fmt.Println(i,v)
		fmt.Println("回调  - ", i, v)
	}
	return nil
}
