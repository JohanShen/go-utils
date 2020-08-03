package logger

type (
	ZapBody struct {
		// 日志等级
		LogLevel Level
		// 描述信息
		Desc string
		// 顶级数据
		TopLevelData map[string]interface{}
		// 附加数据
		Data map[string]interface{}
	}
)
