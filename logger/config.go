package logger

type (
	Config struct {
		// 当前配置名称
		Name string
		// 日志存放的绝对路径
		//	路径中支持的变量
		//		年份	%y	强制2位	%Y 4位
		//		月份	%m	自然位	%M 补齐2位
		//		日期	%d	自然位	%D 补齐2位
		//		小时	%h	自然位	%H 补齐2位
		//		分钟	%mi	自然位	%Mi 补齐2位
		//		秒数	%s	自然位	%S 补齐2位
		LogPath string
		// 日志是否输出到控制台
		ConsoleLog bool
		// 是否延迟写入
		//	0 表示实时写入，不会做延迟处理
		//	n(>0) 表示延迟 n 毫秒写入，可以降低 IO 写入次数
		WriteDelay int
	}
)
