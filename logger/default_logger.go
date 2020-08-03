package logger

import "fmt"

type (
	// 默认 logger 的实现
	defaultLogger struct{}
)

// 默认的日志实现 啥也不干
func DefaultLogger() *defaultLogger {
	return &defaultLogger{}
}

func printOnConsole(level Level, args ...*Body) {
	for _, v := range args {
		fmt.Println(level.String(), v.Desc, v.Data)
	}
}

func (myself *defaultLogger) Debug(args ...*Body) {
	printOnConsole(DebugLevel, args...)
}
func (myself *defaultLogger) Error(args ...*Body) {
	printOnConsole(ErrorLevel, args...)
}
func (myself *defaultLogger) Fatal(args ...*Body) {
	printOnConsole(FatalLevel, args...)
}
func (myself *defaultLogger) Info(args ...*Body) {
	printOnConsole(InfoLevel, args...)
}
func (myself *defaultLogger) Panic(args ...*Body) {
	printOnConsole(PanicLevel, args...)
}
func (myself *defaultLogger) Warn(args ...*Body) {
	printOnConsole(WarnLevel, args...)
}
