package logger

type (
	// 日志类接口
	Logger interface {
		Debug(args ...*Body)
		Error(args ...*Body)
		Fatal(args ...*Body)
		Info(args ...*Body)
		Panic(args ...*Body)
		Warn(args ...*Body)
	}
	// 设置日志类接口
	SetLogger interface {
		SetLogger(logger *Logger)
	}
)
