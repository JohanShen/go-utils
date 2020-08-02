package logger

import (
	"utils/utils"
)

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

func ToZapBody(body *Body, level Level) *ZapBody {
	zBody := &ZapBody{
		LogLevel: level,
	}
	if val, ok := body.Data["userid"]; ok {
		zBody.UserId = utils.AnyToStr(val)
	}
	if val, ok := body.Data["token"]; ok {
		zBody.Token = utils.AnyToStr(val)
	}
	zBody.Data = body.Data
	zBody.Desc = body.Desc
	return zBody
}
