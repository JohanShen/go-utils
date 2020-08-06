package redis

import (
	"testing"
	"utils/logger"
)

func TestJsonCoder_DeCoder(t *testing.T) {

}

func TestJsonCoder_Encoder(t *testing.T) {

	logger1 := logger.DefaultLogger()
	config := DefaultRedisConfig().SetLogger(logger1)
	config.Addrs = []string{"10.0.3.107:6379"}
	r := config.Build().UseValueCoder(&JsonCoder{})

	t.Log(r.Set("test:json:keyjson", config, 0))

}
