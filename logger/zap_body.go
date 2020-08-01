package logger

type (
	ZapBody struct {
		// 日志等级
		LogLevel Level
		// 用户 ID
		UserId string
		//  用户 token
		Token string
		Body
	}

	ZapInfoBody ZapBody
)

func (body *Body) ToZapBody(level Level) *ZapBody {
	zBody := &ZapBody{
		LogLevel: level,
	}
	if val, ok := zBody.Data["userid"]; ok {
		zBody.UserId = val.(string)
	}
	if val, ok := zBody.Data["token"]; ok {
		zBody.Token = val.(string)
	}
	zBody.Data = body.Data
	zBody.Desc = body.Desc
	return zBody
}
