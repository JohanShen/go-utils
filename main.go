package main

import (
	"fmt"
	"time"
	"utils/logger"
)

func main() {
	fmt.Println(1)

	////合建chan
	//c := make(chan os.Signal)
	////监听所有信号
	//signal.Notify(c)
	////阻塞直到有信号传入
	//fmt.Println("启动")
	//
	//go func() {
	//	s := <-c
	//	fmt.Println("退出信号", s)
	//}()
	//
	//time.Sleep(10*time.Second)

	currentTime := time.Now()
	var config = logger.Config{Name: "test", ConsoleLog: false, LogPath: "./%Y%M%D%H%I.log", WriteDelay: 200}
	var logger1 logger.Logger
	logger1 = logger.NewZapLogger(config)

	go func() {
		data := make(map[string]interface{})
		data["time"] = currentTime
		for i := 0; i < 10000; i++ {
			data["userid"] = i + i*2
			logger1.Debug(logger.MakeBody(fmt.Sprintf("男儿当自强 %d", i), data))
			time.Sleep(200 * time.Millisecond)
		}
	}()

	time.Sleep(300 * time.Second)
}
