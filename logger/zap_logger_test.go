package logger

import (
	"fmt"
	"github.com/JohanShen/go-utils/v1/utils"
	"os"
	"os/signal"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestABC(t *testing.T) {

	var a interface{}
	a = 43
	b := numberic2str(a)

	var a1 interface{}
	a1 = "66"
	b1 := numberic2str(a1)
	//b = a.(string)

	t.Log(a, reflect.TypeOf(a))
	t.Log(b, reflect.TypeOf(b))

	t.Log(a1, reflect.TypeOf(a1))
	t.Log(b1, reflect.TypeOf(b1))

}

func numberic2str(a interface{}) string {
	b := ""
	switch a.(type) {
	case int:
		b = strconv.Itoa(a.(int))
	case int16:
		b = strconv.FormatInt(int64(a.(int16)), 10)
	case int8:
		b = strconv.FormatInt(int64(a.(int8)), 10)
	case int32:
		b = strconv.FormatInt(int64(a.(int32)), 10)
	case int64:
		b = strconv.FormatInt(a.(int64), 10)
	case string:
		b = a.(string)
	}
	return b
}

func TestDebug(t *testing.T) {

	//合建chan
	c := make(chan os.Signal)
	//监听所有信号
	signal.Notify(c)
	//阻塞直到有信号传入
	fmt.Println("启动")

	go func() {
		s := <-c
		fmt.Println("退出信号", s)
	}()

	currentTime := time.Now()

	fmt.Println("Current Time in String: ", currentTime.String())

	fmt.Println("MM-DD-YYYY : ", currentTime.Format("01月02日2006年 Monday"))
	fmt.Println("MM-DD-YYYY : ", currentTime.Format("06年1月2日3:4:5 Mon"))

	fmt.Println("MM-DD-YYYY : ", utils.XTime(currentTime).Format("%Y年%m月%D日%h:%I:%s 333333"))

	//path := "./d/f/d/txt.log"

	a, _ := os.Getwd()
	t.Log(a)

	//os.Chdir(path)
	//os.MkdirAll(path, 0666)

	fileInfo, err := os.Stat("test.txt")
	t.Log(fileInfo, err)
	if err != nil {
		if os.IsNotExist(err) {
			t.Log("文件不存在")
		}
	}

	fileInfo, err = os.Stat("test.log")
	t.Log(fileInfo, err)

	var config = &Config{Name: "test", ConsoleLog: false, LogPath: "./%Y/%M%D/%H%I.log", WriteDelay: 50, ShowInTopLevel: []string{"token", "userid"}}
	var logger1 Logger
	logger1 = NewZapLogger(config)

	go func(logger1 Logger) {
		for i := 0; i < 20000000; i++ {
			data := make([]*LogArg, 0, 5)
			data = append(data, ArgToken(utils.AnyToStr(time.Now().Unix())))
			data = append(data, ArgUserId(utils.AnyToStr(i+i*2)))
			data = append(data, ArgAny("time", time.Now()))
			data = append(data, ArgAny("i", i))

			logger1.Debug(fmt.Sprintf("男儿当自强 %d", i), data...)
			//time.Sleep(15 * time.Millisecond)
		}
	}(logger1)

	time.Sleep(190 * time.Second)
	//logger1.Debug(MakeInfoBody("男儿当自强", "123", "", currentTime))
	//logger1.Info(MakeDebugBody("123", "ddd", "", "c"))

}
