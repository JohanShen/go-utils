package logger

import "fmt"

type (
	// 默认 logger 的实现，输出到控制台
	defaultLogger struct{}
)

// 创建 logger 默认实现对象
func DefaultLogger() *defaultLogger {
	return &defaultLogger{}
}

func printOnConsole(level Level, desc string, args ...*LogArg) {
	data := make(map[string]interface{})
	for _, v := range args {
		data[v.Key] = v.Value
	}
	if len(data) == 0 {
		fmt.Println(level.String(), desc)
	} else {
		fmt.Println(level.String(), desc, data)
	}
}

func (myself *defaultLogger) Debug(desc string, args ...*LogArg) {
	printOnConsole(DebugLevel, desc, args...)
}
func (myself *defaultLogger) Error(desc string, args ...*LogArg) {
	printOnConsole(ErrorLevel, desc, args...)
}
func (myself *defaultLogger) Fatal(desc string, args ...*LogArg) {
	printOnConsole(FatalLevel, desc, args...)
}
func (myself *defaultLogger) Info(desc string, args ...*LogArg) {
	printOnConsole(InfoLevel, desc, args...)
}
func (myself *defaultLogger) Panic(desc string, args ...*LogArg) {
	printOnConsole(PanicLevel, desc, args...)
}
func (myself *defaultLogger) Warn(desc string, args ...*LogArg) {
	printOnConsole(WarnLevel, desc, args...)
}
