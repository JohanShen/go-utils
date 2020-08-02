package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"utils/utils"
)

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

func NewZapLogger(config Config) *zapLogger {

	if len(config.Name) == 0 {
		config.Name = "default"
	}

	if obj, ok := createdObject[config.Name]; ok {
		return obj
	}

	obj := &zapLogger{
		config: &config,
	}

	if config.WriteDelay > 0 {
		//delay = queue.NewDelayQueue(config.WriteMode, obj.callback)
		config.WriteDelay = utils.IfElseInt(config.WriteDelay > 100, config.WriteDelay, 100)
		config.WriteDelay = utils.IfElseInt(config.WriteDelay > 10000, 10000, config.WriteDelay)

		obj.locker = &sync.Mutex{}
		obj.delayQueue = make([]*ZapBody, 0, 100)
		obj.delayNum = 10
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

	obj.logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))
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
	now := time.Now()
	logPath := utils.XTime(now).Format(path)
	file, _ := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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
				z.locker.Unlock()

				//fmt.Println("总长度：", len1, "取出数：", idx)
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
			println(err)
		}
	}()

	defer z.logger.Sync()

	for _, body := range bodies {
		fields := make([]zap.Field, 0, 5)
		fields = append(fields, zap.String("name", z.config.Name))
		fields = append(fields, zap.String("userid", body.UserId))
		fields = append(fields, zap.Any("data", body.Data))
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

func (z *zapLogger) Write(level Level, args ...*Body) {

	msgs := make([]*ZapBody, 0, len(args))
	for _, item := range args {
		nitem := item.ToZapBody(level)
		msgs = append(msgs, nitem)
	}

	if z.config.WriteDelay == 0 {
		//需要即时处理的数据
		for _, item := range msgs {
			z.writeMsg(item)
		}
	} else {
		//延迟处理的数据
		z.locker.Lock()
		for _, item := range msgs {
			z.delayQueue = append(z.delayQueue, item)
		}
		z.locker.Unlock()
	}
}

func (z *zapLogger) Debug(args ...*Body) {
	z.Write(DebugLevel, args...)
}
func (z *zapLogger) Error(args ...*Body) {
	z.Write(ErrorLevel, args...)
}
func (z *zapLogger) Fatal(args ...*Body) {
	z.Write(FatalLevel, args...)
}
func (z *zapLogger) Info(args ...*Body) {
	z.Write(InfoLevel, args...)
}
func (z *zapLogger) Panic(args ...*Body) {
	z.Write(PanicLevel, args...)
}
func (z *zapLogger) Warn(args ...*Body) {
	z.Write(WarnLevel, args...)
}