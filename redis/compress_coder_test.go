package redis

import (
	"github.com/JohanShen/go-utils/v1/logger"
	"testing"
)

func TestCompressCoder_Encoder(t *testing.T) {

	logger1 := logger.DefaultLogger()
	config := DefaultRedisConfig().SetLogger(logger1)
	config.Addrs = []string{"10.0.3.107:6379"}
	r := config.Build().UseValueCoder(&CompressCoder{})

	t.Log(r.Set("test:json:keyzip", config, 0))

}

func TestCompressCoder_DeCoder(t *testing.T) {

	logger1 := logger.DefaultLogger()
	config := DefaultRedisConfig().SetLogger(logger1)
	config.Addrs = []string{"10.0.3.107:6379"}
	r1 := config.Build()
	r := r1.UseValueCoder(&CompressCoder{})

	r1.Set("test:json:keytext", "1233333", 0)

	var config1 Config
	t.Log(r.GetByCoder("test:json:keyzip", &config1))
	t.Log(config1)

}
