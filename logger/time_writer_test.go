package logger

import (
	"fmt"
	"testing"
	"time"
	"utils/utils"
)

func TestNewTimeWriter(t *testing.T) {
	tw := NewTimeWriter("./%Y%M/%D/%H%I.log")

	for i := 0; i < 1000000; i++ {
		str := fmt.Sprintf("abc%v 中文 %#v\r\n", i, time.Now().Unix())
		b := utils.StringToBytes(str)
		t.Logf("内容 %#v , %#v", str, b)
		n, err := tw.Write(b)
		t.Logf("当前写入文件 %#v ,历史数量 %#v 写入长度 %#v 状态 %#v", tw.CurrentPath, len(tw.files), n, err)
		// 为了模式出实际效果 写入速度控制慢一点
		//time.Sleep(time.Millisecond * 20)
	}
}
