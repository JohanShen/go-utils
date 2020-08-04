// 日志类
//  用于抽象日志处理层
package logger

type (
	// 日志类接口
	Logger interface {
		Debug(desc string, args ...*LogArg)
		Error(desc string, args ...*LogArg)
		Fatal(desc string, args ...*LogArg)
		Info(desc string, args ...*LogArg)
		Panic(desc string, args ...*LogArg)
		Warn(desc string, args ...*LogArg)
	}
)
