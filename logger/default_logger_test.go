package logger

import (
	"fmt"
	"testing"
	"time"
)

func TestDefaultLogger(t *testing.T) {
	var logger1 Logger
	logger1 = DefaultLogger()
	//go func(logger1 Logger) {
	bodies := make([]*Body, 0, 1000)
	for i := 0; i < 2000; i++ {
		data := make(map[string]interface{})
		data["time"] = time.Now()
		data["userid"] = i + i*2
		data["i"] = i
		//time.Sleep(15 * time.Millisecond)
		body := MakeBody(fmt.Sprintf("男儿当自强 %d", i), data)
		bodies = append(bodies, body)
	}
	logger1.Debug(bodies...)
	//}(logger1)

}
