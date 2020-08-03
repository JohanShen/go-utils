package logger

import (
	"fmt"
	"testing"
	"time"
	"utils/utils"
)

func TestDefaultLogger(t *testing.T) {
	var logger1 Logger
	logger1 = DefaultLogger()
	//go func(logger1 Logger) {

	for i := 0; i < 2000; i++ {
		logger1.Debug(fmt.Sprintf("男儿当自强 %d", i), ArgUserId(utils.AnyToStr(i+i*2)), ArgAny("time", time.Now()), ArgAny("i", i))
	}
	//}(logger1)

}
