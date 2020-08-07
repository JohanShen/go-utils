package redis

import (
	"encoding/json"
	"fmt"
	"github.com/JohanShen/go-utils/logger"
	"github.com/JohanShen/go-utils/utils"
	"github.com/vmihailenco/msgpack/v5"
	"math/rand"
	"testing"
)

func TestConfig_Build(t *testing.T) {

	logger1 := logger.DefaultLogger()
	config := DefaultRedisConfig().SetLogger(logger1)
	config.Addrs = []string{"10.0.3.107:6379"}
	r := config.Build()

	str, _ := json.Marshal(config)

	//t.Log(config)
	if val, err := msgpack.Marshal(config); err == nil {
		vval1, _ := utils.GZipBytes(val)
		vval2, _ := utils.ZipBytes(val)
		t.Log(len(str), len(val), len(vval1), len(vval2))

		if set := r.Set("oj8k:good", val, 0); set.Err() != nil {
			t.Logf("写入失败 %+v", set)
		}
		r.HSetNX("oj8k:good1", "MMM", val)

		val1, _ := r.GetRaw("oj8k:good")
		var varl2, varl3 *Config
		err := msgpack.Unmarshal(val1, &varl2)
		if err != nil {
			t.Error(err)
		}
		t.Log(varl2)

		val3, _ := r.HGetRaw("oj8k:good1", "MMM")
		err1 := msgpack.Unmarshal(val3, &varl3)
		if err1 != nil {
			t.Error(err1)
		}
		t.Log(varl3)

		vall := r.HGetAll("oj8k:good1")
		t.Log(vall)
	}

}

func TestConfig_SetLogger(t *testing.T) {
	logger1 := logger.DefaultLogger()
	config := DefaultRedisConfig().SetLogger(logger1)
	config.Addrs = []string{"10.0.3.107:6379"}
	r := config.Build()

	testbigdata := make(map[string]string)

	for i := 0; i < 1000; i++ {
		testbigdata[randSeq(5)] = "This is a redis client GUI tool written based on Java SWT and Jedis. It's my objective to build the most convenient redis client GUI tool in the world. In the first place, it will facilitate in editing redis data, such as: add, update, delete, search, cut, copy, paste etc." // randSeq(355)
	}

	if set := r.Set("oj8k:bigdata1", testbigdata, 0); set.Err() != nil {
		t.Logf("oj8k:bigdata1 写入失败 %+v", set.Err())
	}
	if set := r.Set("oj8k:bigdata1", testbigdata, 0); set.Err() != nil {
		t.Logf("oj8k:bigdata1 写入失败 %+v", set)
	}

	str, _ := json.Marshal(testbigdata)

	if set := r.Set("oj8k:bigdata2", str, 0); set.Err() != nil {
		t.Logf("oj8k:bigdata2 写入失败 %+v", set)
	}

	if val, err := utils.StructToBytes(testbigdata); err == nil {
		vval1, _ := utils.GZipBytes(val)
		vval2, _ := utils.ZipBytes(str)
		t.Log("json", len(str))
		t.Log("msgpack", len(val))
		t.Log("gzip", len(vval1), float64(len(vval1))/float64(len(str))*100)
		t.Log("zip", len(vval2), float64(len(vval2))/float64(len(str))*100)

		if set := r.Set("oj8k:bigdata3", vval1, 0); set.Err() != nil {
			t.Logf("oj8k:bigdata3 写入失败 %+v", set)
		}
		if set := r.Set("oj8k:bigdata4", vval2, 0); set.Err() != nil {
			t.Logf("oj8k:bigdata4 写入失败 %+v", set)
		}

		mu1 := r.MemoryUsage("oj8k:bigdata1", 1)
		mu2 := r.MemoryUsage("oj8k:bigdata2", 2)
		mu3 := r.MemoryUsage("oj8k:bigdata3", 3)
		mu4 := r.MemoryUsage("oj8k:bigdata4", 4)
		mu5 := r.MemoryUsage("")

		t.Log("内存使用情况：", mu1, mu2, mu3, mu4, mu5)

		var obj2 map[string]string
		if err2 := utils.BytesToStruct(val, &obj2); err2 == nil {
			t.Log(len(obj2))
		}
	}

}

func TestDefaultRedisConfig(t *testing.T) {
	logger1 := logger.DefaultLogger()
	config := DefaultRedisConfig().SetLogger(logger1)
	config.Addrs = []string{"10.0.3.107:6379"}
	r := config.Build()

	t.Log(r)
}

func TestConfig_Build2(t *testing.T) {

	logger1 := logger.DefaultLogger()
	config := DefaultRedisConfig().SetLogger(logger1)
	config.Addrs = []string{"10.0.3.107:6379"}
	r := config.Build()
	pipe := r.Pipeline()

	for i := 0; i < 150; i++ {
		//pipe.Set()
		key1 := fmt.Sprintf("oj8k:pipe:key%d", i)
		key2 := fmt.Sprintf("oj8k:set:%d", i)
		if set := pipe.Set(key1, i, 0); set.Err() != nil {
			t.Logf("%v 写入失败 %+v", key1, set)
		}
		pipe.Get(key2)
		//if i==55{
		//	if set := pipe.Set(key2, config, 0); set.Err() != nil {
		//		t.Logf("%v 写入失败 %+v", key2, set)
		//	}
		//}
		//if set := r.Client.Set(key2, i, 0); set.Err() != nil {
		//	t.Logf("%v 写入失败 %+v", key2, set)
		//}
	}

	//a := pipe.FlushAll()
	//t.Log(a)

	b, err := pipe.Exec()

	if err == nil {
		for i, v := range b {
			t.Log(i, v)
		}
	}
	t.Log(err)

}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
