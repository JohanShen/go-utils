package logger

type (
	LogArg struct {
		Key        string      // 参数名
		IsTopLevel bool        // 是否在顶层输出
		Value      interface{} // 值
	}
	// 常见 KEY 设置
	LogArgKey string
)

// 内置 Key 预设
const (
	KeyUserid LogArgKey = "userid"
	KeyToken  LogArgKey = "token"
)

func (key LogArgKey) String() string {
	return string(key)
}

func ArgUserId(userid string) *LogArg {
	return &LogArg{
		Key:   KeyUserid.String(),
		Value: userid,
	}
}

func ArgAny(key string, val interface{}) *LogArg {
	return &LogArg{
		Key:   key,
		Value: val,
	}
}

func ArgToken(token string) *LogArg {
	return &LogArg{
		Key:   KeyToken.String(),
		Value: token,
	}
}
