package timer

import (
	"errors"
	"fmt"
	"github.com/JohanShen/go-utils/v1/logger"
	"math/rand"
	"reflect"
	"runtime"
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

	logger    logger.Logger // 日志记录器
	mutex     *sync.Mutex   //互斥锁
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

//设置日志
func (obj *Timer) UseLogger(logger logger.Logger) *Timer {
	obj.logger = logger
	return obj
}

// 执行单个回调函数
func execute(obj *Timer, fun func(*EventArg) error, now *time.Time) {

	funcName := runtime.FuncForPC(reflect.ValueOf(fun).Pointer()).Name()

	defer func() {
		if err := recover(); err != nil {
			logger.Panic(obj.logger, fmt.Sprintf("%s 回调错误 panic ", obj.Name), err)
		}
	}()

	logger.Debug(obj.logger, fmt.Sprintf("%s 回调 %s", obj.Name, funcName))
	arg := &EventArg{Sender: obj, Msg: fmt.Sprintf("回调执行时间：%s ", now)}
	if err := fun(arg); err != nil {
		logger.Error(obj.logger, fmt.Sprintf("%s 回调错误 error", obj.Name), logger.ArgAny("err", err))
	}
}

func (obj *Timer) lisenTicker() {

	logger.Debug(obj.logger, fmt.Sprintf("%s 准备监听", obj.Name))

	defer func(obj *Timer) {
		obj.wait.Done()
		logger.Debug(obj.logger, fmt.Sprintf("%s 退出监听函数 ", obj.Name))
	}(obj)

	obj.wait.Add(1)

Loop:

	for {
		select {
		case now := <-obj.ticker.C:
			//执行回调函数
			logger.Debug(obj.logger, fmt.Sprintf("%s 回调函数-START[%d]", obj.Name, runtime.NumGoroutine()))
			for _, f := range obj.Elapsed.events {
				// 将最终的执行函数单独包装成方法
				// 有利于其中一个或多个回调函数出错时，保证程序继续运行
				go execute(obj, f, &now)
			}
			logger.Debug(obj.logger, fmt.Sprintf("%s 回调函数-END", obj.Name))
		case state, ok := <-obj.state:
			logger.Debug(obj.logger, fmt.Sprintf("%s 状态更改 state(value = %v, isclose = %v)", obj.Name, state, ok))
			if !ok {
				logger.Debug(obj.logger, fmt.Sprintf("%s 通道已被关闭", obj.Name))
				break Loop
			}
			if state == 1 {
				//obj.mutex.Lock()
				logger.Debug(obj.logger, fmt.Sprintf("%s 收到关闭通知", obj.Name))
				obj.wait.Done()

				//obj.mutex.Unlock()
				break Loop
			}

		}

		logger.Debug(obj.logger, fmt.Sprintf("%s 准备接受下一轮信号", obj.Name))
	}

	logger.Debug(obj.logger, fmt.Sprintf("%s 完成监听", obj.Name))
}

// 设置定时器的名称
func (obj *Timer) SetName(name string) *Timer {

	logger.Debug(obj.logger, fmt.Sprintf("%s 更改名称 %[1]s => %s", obj.Name, name))
	obj.Name = name
	return obj
}

//设置执行间隔
func (obj *Timer) SetInterval(interval int64) error {

	if interval <= 1 {
		return errors.New("non-positive interval for NewTimer")
	}

	obj.mutex.Lock()
	obj.interval = interval
	logger.Debug(obj.logger, fmt.Sprintf("%s 更改间隔 %[1]s -> %d", obj.Name, interval))
	if obj.ticker != nil {
		//fmt.Println()
		//fmt.Print("set interval")
		obj.ticker.Stop()
		obj.wait.Add(1)
		obj.state <- 1
		obj.wait.Wait()
		obj.ticker = time.NewTicker(time.Duration(obj.interval) * time.Millisecond)
		//obj.state = make(chan int, 1)
		//fmt.Print(" ok")
	}
	logger.Debug(obj.logger, fmt.Sprintf("%s 更改间隔成功 %[1]s -> %d", obj.Name, interval))
	obj.mutex.Unlock()
	if obj.IsRunning {
		go obj.lisenTicker()
	}
	//fmt.Printf("%p", obj.ticker)
	return nil
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

	logger.Debug(obj.logger, fmt.Sprintf("%s 执行间隔 %d", obj.Name, obj.interval))

	obj.mutex.Lock()
	obj.startTime = time.Now()
	obj.IsRunning = true

	initNewObj(obj)

	obj.mutex.Unlock()
	//obj.wait.Add(1)
	go obj.lisenTicker()
	//obj.wait.Wait()
}

// 停止计时器
func (obj *Timer) Stop() {

	if obj.state == nil || obj.ticker == nil || !obj.IsRunning {
		return
	}

	defer func() {
		obj.mutex.Unlock()
		if err := recover(); err != nil {
			//
			logger.Panic(obj.logger, fmt.Sprintf("%s 停止时出错", obj.Name), err)
		} else {
			//
			logger.Debug(obj.logger, fmt.Sprintf("%s 停止成功", obj.Name))
		}
	}()

	obj.mutex.Lock()
	obj.stopTime = time.Now()
	obj.wait.Add(1)

	logger.Debug(obj.logger, fmt.Sprintf("%s 停止-START", obj.Name))
	obj.ticker.Stop()
	obj.state <- 1
	obj.wait.Wait() //这个的作用就是确保协程退出后执行下面的代码
	obj.IsRunning = false
	close(obj.state)
	logger.Debug(obj.logger, fmt.Sprintf("%s 停止-STOPPED", obj.Name))

}
