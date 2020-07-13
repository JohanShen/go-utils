package timer

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {}

// 事件回调时回传的参数
type EventArg struct {
	Sender *Timer
	Msg    string
}

// 事件
type Event struct {
	events []func(*EventArg) error
}

// 绑定事件
// 支持链式绑定
// timer.Bind(func1).Bind(func2)
func (e *Event) Bind(fun func(*EventArg) error) *Event {
	e.events = append(e.events, fun)
	return e
}

type Timer struct {
	Elapsed   Event  //达到间隔时的回调
	Name      string //名称
	IsRunning bool   //运行状态

	mutex     *sync.Mutex //互斥锁
	interval  int64
	state     chan int
	startTime time.Time
	stopTime  time.Time
	wait      *sync.WaitGroup
	ticker    *time.Ticker
}

// 创建新的实例化对象
func NewTimer(interval int64) *Timer {
	if interval <= 1 {
		panic(errors.New("non-positive interval for NewTimer"))
	}
	name := fmt.Sprintf("%x", rand.Uint64())
	obj := &Timer{interval: interval, Name: name, IsRunning: false}
	obj.mutex = &sync.Mutex{}
	obj.wait = &sync.WaitGroup{}
	return obj
}

func initNewObj(obj *Timer) {
	obj.stopTime = time.Unix(0, 0)
	obj.ticker = time.NewTicker(time.Duration(obj.interval) * time.Millisecond)
	obj.state = make(chan int, 1)
}

// 执行单个回调函数
func execute(obj *Timer, fun func( *EventArg) error, now *time.Time) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出了错1：", "", err)
		}
	}()

	arg := &EventArg{Sender: obj, Msg: fmt.Sprintf("回调执行时间：%s ", now)}
	if err := fun(arg); err != nil {
		fmt.Println("出了错2：", err)
	}
}

func lisenTicker(obj *Timer) {
	// fmt.Println(fmt.Sprintf("%s 准备监听", obj.Name))

	defer func(obj *Timer) {
		obj.wait.Done()
		//fmt.Print(fmt.Sprintf("========[][][][] %s next ", obj.Name))
	}(obj)

	obj.wait.Add(1)

Loop:

	for {
		select {
		case now := <-obj.ticker.C:
			//执行回调函数
			for _, f := range obj.Elapsed.events {
				// 将最终的执行函数单独包装成方法
				// 有利于其中一个或多个回调函数出错时，保证程序继续运行
				execute(obj, f, &now)
			}
		case state, ok := <-obj.state:
			if !ok {
				break Loop
			}
			if state == 1 {
				//obj.mutex.Lock()
				fmt.Println(fmt.Sprintf("%s will stop. %t %d", obj.Name, ok, state))
				obj.IsRunning = false
				close(obj.state)
				obj.wait.Done()

				//obj.mutex.Unlock()
				break Loop
			}

		}

		//fmt.Print(fmt.Sprintf("%s next ", obj.Name))
	}

	//fmt.Println(fmt.Sprintf("%s 退出监听", obj.Name))
}

// 设置定时器的名称
func (obj *Timer) SetName(name string) *Timer {
	obj.Name = name
	return obj
}

//设置执行间隔
func (obj *Timer) SetInterval(interval int64) {

	obj.mutex.Lock()
	obj.interval = interval
	if obj.ticker != nil {
		//fmt.Println()
		//fmt.Print("set interval")
		obj.ticker.Stop()
		//obj.state <- 2
		obj.ticker = time.NewTicker(time.Duration(obj.interval) * time.Millisecond)
		//fmt.Print(" ok")
	}
	obj.mutex.Unlock()
	if obj.IsRunning {
		go lisenTicker(obj)
	}
	//fmt.Printf("%p", obj.ticker)
}

// 开始计时器
// 需要通过 NewTimer 方法获取实例化
func (obj *Timer) Start() {

	// if obj.wait == nil {
	// 	panic(errors.New("尚未初始化，请通过 New() 方法初始化对象。"))
	// 	return
	// }
	if obj.IsRunning {
		return
	}
	fmt.Println(fmt.Sprintf("%s start %d", obj.Name, obj.interval))

	obj.mutex.Lock()
	obj.startTime = time.Now()
	obj.IsRunning = true

	initNewObj(obj)

	obj.mutex.Unlock()
	//obj.wait.Add(1)
	go lisenTicker(obj)
	//obj.wait.Wait()
}

// 停止计时器
func (obj *Timer) Stop() {

	if obj.state == nil || obj.ticker == nil {
		return
	}

	defer func() {
		obj.mutex.Unlock()
		if err := recover(); err != nil {
			//
		} else {
			//
		}
	}()

	obj.mutex.Lock()
	if !obj.IsRunning {
		return
	}

	obj.stopTime = time.Now()
	obj.wait.Add(1)

	fmt.Println(fmt.Sprintf("%s go to stop timer .", obj.Name))
	obj.ticker.Stop()
	obj.state <- 1
	obj.wait.Wait() //这个的作用就是确保协程退出后执行下面的代码
	fmt.Println(fmt.Sprintf("%s stopped spent %s", obj.Name, obj.stopTime.Sub(obj.startTime)))

}
