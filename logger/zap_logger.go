package logger

import (
	"fmt"
	"github.com/JohanShen/go-utils/v1/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// zap 日志的具体实现
//  可能存在的BUG：
//   1 日志不是顺序输出的
//   2 日志没有按生成时间输出到对应文件
//
type zapLogger struct {
	config      *Config
	logger      *zap.Logger
	sugarLogger *zap.SugaredLogger
	delayQueue  []*ZapBody   // 存放延迟处理的列队
	delayNum    int          // 延迟写入的数量
	ticker      *time.Ticker // 用于处理延迟数据
	stopTicker  chan os.Signal
	locker      *sync.Mutex
}

var createdObject map[string]*zapLogger

func NewZapLogger(config *Config) *zapLogger {

	if len(config.Name) == 0 {
		config.Name = "default"
	}

	if obj, ok := createdObject[config.Name]; ok {
		return obj
	}

	obj := &zapLogger{
		config: config,
	}

	if config.WriteDelay > 0 {
		//delay = queue.NewDelayQueue(config.WriteMode, obj.callback)
		config.WriteDelay = utils.IfElseInt(config.WriteDelay > 10, config.WriteDelay, 10)
		config.WriteDelay = utils.IfElseInt(config.WriteDelay > 10000, 10000, config.WriteDelay)

		obj.locker = &sync.Mutex{}
		obj.delayQueue = make([]*ZapBody, 0, 10000)
		obj.delayNum = 10000
		obj.stopTicker = make(chan os.Signal, 1)
		obj.ticker = time.NewTicker(time.Duration(config.WriteDelay) * time.Millisecond)
		signal.Notify(obj.stopTicker, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

		go obj.listenTicker()
	}

	cores := make([]zapcore.Core, 0, 2)
	// 实现多个输出
	if len(config.LogPath) > 0 {
		encoder := getEncoder()
		writer := getLogWriter(config.LogPath)
		cores = append(cores, zapcore.NewCore(encoder, writer, zapcore.DebugLevel))
	}
	if config.ConsoleLog {
		cores = append(cores, zapcore.NewCore(zapcore.NewConsoleEncoder(getEncoderConfig()), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), zapcore.DebugLevel))
	}
	core := zapcore.NewTee(cores...)

	//zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))
	obj.logger = zap.New(core, zap.AddStacktrace(zap.WarnLevel))
	obj.sugarLogger = obj.logger.Sugar()

	return obj
}

func getEncoderConfig() zapcore.EncoderConfig {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return encoderConfig
}
func getEncoder() zapcore.Encoder {
	encoderConfig := getEncoderConfig()
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(path string) zapcore.WriteSyncer {
	//now := time.Now()
	//logPath := utils.XTime(now).Format(path)
	//file, _ := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	file := NewTimeWriter(path)
	return zapcore.AddSync(file)
}

// 监听定时器
func (z *zapLogger) listenTicker() {

	defer z.ticker.Stop()

Loop:
	for {
		select {
		case <-z.ticker.C:

			len1 := len(z.delayQueue)
			idx := utils.IfElseInt(len1 > z.delayNum, z.delayNum, len1)

			if idx > 0 {
				z.locker.Lock()
				//单次回调传入的数组
				result := z.delayQueue[:idx]
				//从原数组中移除已经回调的数据
				z.delayQueue = z.delayQueue[idx:]
				//result1 := make([]ZapBody,1000)
				//copy(result1, result)
				z.locker.Unlock()

				fmt.Println("总长度：", len1, "取出数：", idx)
				// 将最终的执行函数单独包装成方法
				// 有利于其中一个或多个回调函数出错时，保证程序继续运行
				go z.writeMsg(result...)
			}

		case <-z.stopTicker:
			z.locker.Lock()
			len1 := len(z.delayQueue)
			if len1 > 0 {
				// 将剩余未写入的一次性写入
				//fmt.Println("剩余：", len1)
				go z.writeMsg(z.delayQueue...)
			}
			z.locker.Unlock()

			break Loop

		}
	}
	//println("退出延迟等候列队")
}

func (z *zapLogger) writeMsg(bodies ...*ZapBody) {

	defer func() {
		if err := recover(); err != nil {
			// 这个位置出错说明IO操作出现问题
			// 无法写入到错误 为了程序稳定 也不宜 panic
			println("写入日志时出错 writeMsg() panic ", err)
		}
	}()

	defer func() {
		if err := z.logger.Sync(); err != nil {
			println("写入日志时出错 writeMsg() Sync() ", err)
		}
	}()

	for _, body := range bodies {
		fields := make([]zap.Field, 0, len(body.TopLevelData)+2)
		fields = append(fields, zap.String("name", z.config.Name))
		for key, val := range body.TopLevelData {
			fields = append(fields, zap.Any(key, val))
		}
		if body.Data != nil {
			fields = append(fields, zap.Any("data", body.Data))
		}

		switch body.LogLevel {
		case DebugLevel:
			z.logger.Debug(body.Desc, fields...)
		case ErrorLevel:
			z.logger.Error(body.Desc, fields...)
		case WarnLevel:
			z.logger.Warn(body.Desc, fields...)
		case PanicLevel:
			z.logger.Panic(body.Desc, fields...)
		case FatalLevel:
			z.logger.Fatal(body.Desc, fields...)
		default:
			z.logger.Info(body.Desc, fields...)
		}
	}
}

func (z *zapLogger) write(level Level, desc string, args ...*LogArg) {

	body := &ZapBody{
		LogLevel:     level,
		Desc:         desc,
		TopLevelData: make(map[string]interface{}),
	}
	for _, item := range args {
		if len(z.config.ShowInTopLevel) > 0 && utils.IndexOfWithoutCase(z.config.ShowInTopLevel, item.Key) > -1 {
			body.TopLevelData[item.Key] = item.Value
		} else {
			if body.Data == nil {
				body.Data = make(map[string]interface{})
			}
			body.Data[item.Key] = item.Value
		}

	}

	if z.config.WriteDelay == 0 {
		//需要即时处理的数据
		go z.writeMsg(body)
	} else {
		//延迟处理的数据
		//此处可能会存在性能问题
		z.locker.Lock()
		z.delayQueue = append(z.delayQueue, body)
		z.locker.Unlock()
	}

}

func (z *zapLogger) Debug(desc string, args ...*LogArg) {
	z.write(DebugLevel, desc, args...)
}
func (z *zapLogger) Error(desc string, args ...*LogArg) {
	z.write(ErrorLevel, desc, args...)
}
func (z *zapLogger) Fatal(desc string, args ...*LogArg) {
	z.write(FatalLevel, desc, args...)
}
func (z *zapLogger) Info(desc string, args ...*LogArg) {
	z.write(InfoLevel, desc, args...)
}
func (z *zapLogger) Panic(desc string, args ...*LogArg) {
	z.write(PanicLevel, desc, args...)
}
func (z *zapLogger) Warn(desc string, args ...*LogArg) {
	z.write(WarnLevel, desc, args...)
}
