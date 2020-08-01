package logger

/* 定义接口 */
type Logger interface {
	Debug(args ...*Body)
	Error(args ...*Body)
	Fatal(args ...*Body)
	Info(args ...*Body)
	Panic(args ...*Body)
	Warn(args ...*Body)
}
