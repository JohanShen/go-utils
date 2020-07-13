package queue

import (
	"../timer"
	"../utils"
	"fmt"
	"sync"
)

type delayQueue struct {

	queue []interface{}
	callback func([]interface{}) error
	timer *timer.Timer
	locker *sync.Mutex
}


func  NewDelayQueue(interval int64, callback func([]interface{}) error) *delayQueue{
	dq := &delayQueue{	}
	dq.queue = make([]interface{}, 50, 100)
	dq.timer = timer.NewTimer(interval)
	dq.timer.Elapsed.Bind(dq.handleCallBack)
	dq.callback = callback
	return dq
}

// 推入延迟处理的数据
func (obj *delayQueue) Push(items ...interface{})  {
	defer func() {
		obj.locker.Unlock()
		if !obj.timer.IsRunning{
			obj.timer.Start()
		}
	}()

	obj.locker.Lock()
	for _, v := range items {
		obj.queue = append(obj.queue, v)
	}
}


// 延迟处理回调
func (obj *delayQueue) handleCallBack(arg *timer.EventArg) error{

	if obj.callback!=nil{

		obj.timer.Stop()

		len1 := len(obj.queue)
		idx := utils.IntMax(len1, 100)

		obj.locker.Lock()
		//单次回调传入的数组
		result := obj.queue[:idx]
		//从原数组中移除已经回调的数据
		obj.queue = obj.queue[idx:]
		obj.locker.Unlock()


		if err := obj.callback(result); err!=nil{
			//执行出错了
			fmt.Println("回调出错：", err)
		}

		if len(obj.queue)>0{
			obj.timer.Start()
		}
	}

	return nil
}