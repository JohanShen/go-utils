package logger

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestDebug(t *testing.T) {

	var str string

	t.Log(str == "")

	currentTime := time.Now()

	fmt.Println("Current Time in String: ", currentTime.String())

	fmt.Println("MM-DD-YYYY : ", currentTime.Format("01月02日2006年"))
	fmt.Println("MM-DD-YYYY : ", currentTime.Format("2006年1月2日"))

	path := "./d/f/d/txt.log"

	a, _ := os.Getwd()
	t.Log(a)

	os.Chdir(path)
	os.MkdirAll(path, 0666)

	fileInfo, err := os.Stat("test.txt")
	t.Log(fileInfo, err)
	if err != nil {
		if os.IsNotExist(err) {
			t.Log("文件不存在")
		}
	}

	fileInfo, err = os.Stat("test.log")
	t.Log(fileInfo, err)

	var logger Logger
	logger = NewZapLogger("default", "")

	logger.Debug(currentTime)
	logger.Debug("男儿当自强")
	logger.Info("c")

}
