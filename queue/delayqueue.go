package queue

import (
	"fmt"
	"sync"
	"utils/timer"
	"utils/utils"
)

type DelayQueue struct {
	queue    []interface{}
	callback func([]interface{}) error
	timer    *timer.Timer
	locker   *sync.Mutex
}

func NewDelayQueue(interval int64, callback func([]interface{}) error) *DelayQueue {
	dq := &DelayQueue{}
	dq.locker = &sync.Mutex{}
	dq.queue = make([]interface{}, 0, 100)
	dq.timer = timer.NewTimer(interval)
	dq.timer.Elapsed.Bind(dq.handleCallBack)
	dq.callback = callback
	return dq
}

// 推入延迟处理的数据
func (obj *DelayQueue) Push(items ...interface{}) {
	defer func() {
		obj.locker.Unlock()
		if !obj.timer.IsRunning {
			obj.timer.Start()
		}
	}()

	obj.locker.Lock()
	for _, v := range items {
		obj.queue = append(obj.queue, v)
	}
}

// 延迟处理回调
func (obj *DelayQueue) handleCallBack(arg *timer.EventArg) error {

	//fmt.Println(arg)

	if obj.callback != nil {

		obj.timer.Stop()

		len1 := len(obj.queue)
		idx := utils.IfElseInt(len1 > 100, 100, len1)

		//fmt.Println("回调：", len1, idx)

		obj.locker.Lock()
		//单次回调传入的数组
		result := obj.queue[:idx]
		//从原数组中移除已经回调的数据
		obj.queue = obj.queue[idx:]
		obj.locker.Unlock()

		//fmt.Println("回调：", result)

		go obj.execute(result)

		if len(obj.queue) > 0 {
			obj.timer.Start()
		}
	}

	return nil
}

// 执行回调函数
func (obj *DelayQueue) execute(result []interface{}) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出了错1：", "", err)
		}
	}()

	if err := obj.callback(result); err != nil {
		fmt.Println("出了错2：", err)
	}
}
