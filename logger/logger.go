package logger

/* 定义接口 */
type Logger interface {
	Debug(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Info(args ...interface{})
	Panic(args ...interface{})
	Warn(args ...interface{})
}
