package timer

import (
	"fmt"
	"runtime"
	"testing"
	"time"
	"utils/logger"
)

func TestNewTimer(t *testing.T) {

	var config = &logger.Config{Name: "test", ConsoleLog: false, LogPath: "./%Y/%M%D/%H%I.log", WriteDelay: 50, ShowInTopLevel: []string{"token", "userid"}}
	var logger1 logger.Logger

	logger1 = logger.NewZapLogger(config)
	//logger1 = logger.DefaultLogger()

	timer := NewTimer(500).UseLogger(logger1)
	timer.Elapsed.Bind(handleCallBack)
	timer.Start()
	print("协程数量:", runtime.NumGoroutine())
	go func(timer *Timer) {
		for i := 0; i < 100; i++ {
			if err := timer.SetInterval(int64(i*i) + 1); err == nil {

			}
			time.Sleep(time.Millisecond * 250)
		}
	}(timer)

	time.Sleep(time.Second * 30)
	timer.Stop()
	print("协程数量:", runtime.NumGoroutine())
}

// 延迟处理回调
func handleCallBack(arg *EventArg) error {
	//panic(arg)
	fmt.Println(arg)
	return nil
}
