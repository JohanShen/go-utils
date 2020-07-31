package timer

import (
	"fmt"
	"testing"
	"time"
)

func TestNewTimer(t *testing.T) {

	//logger := logger.NewZapLogger("default", "")
	timer := NewTimer(500)
	timer.Elapsed.Bind(handleCallBack)
	timer.Start()

	time.Sleep(time.Second * 10)
}

// 延迟处理回调
func handleCallBack(arg *EventArg) error {
	fmt.Print(arg)
	return nil
}
