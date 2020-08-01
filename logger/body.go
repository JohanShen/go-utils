package logger

type (
	Body struct {
		// 描述信息
		Desc string
		// 附加数据
		Data map[string]interface{}
	}
)

func MakeBody(desc string, data map[string]interface{}) *Body {
	return &Body{
		Desc: desc,
		Data: data,
	}
}
