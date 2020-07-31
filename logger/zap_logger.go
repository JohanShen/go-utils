package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type zapLogger struct {
	logger      *zap.Logger
	sugarLogger *zap.SugaredLogger
}

var createdObject map[string]*zapLogger

func NewZapLogger(name string, path string) *zapLogger {

	if len(name) == 0 {
		name = "default"
	}
	if len(path) == 0 {
		path = ""
	}
	if obj, ok := createdObject[name]; ok {
		return obj
	}

	writer := getLogWriter()
	encoder := getEncoder()
	//core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	// 实现多个输出
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, writer, zapcore.DebugLevel),
		//同时将日志输出到控制台，NewJSONEncoder 是结构化输出
		zapcore.NewCore(zapcore.NewConsoleEncoder(getEncoderConfig()), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), zapcore.DebugLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))

	return &zapLogger{
		logger:      logger,
		sugarLogger: logger.Sugar(),
	}
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

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	return zapcore.AddSync(file)
}

func (z *zapLogger) Debug(args ...interface{}) {
	defer z.logger.Sync()
	z.logger.Debug("ddd", zap.String("token", "md5"), zap.Any("msg", args))
}
func (z *zapLogger) Error(args ...interface{}) {
	defer z.sugarLogger.Sync()
	z.sugarLogger.Error(args)
}
func (z *zapLogger) Fatal(args ...interface{}) {
	defer z.sugarLogger.Sync()
	z.sugarLogger.Fatal(args)
}
func (z *zapLogger) Info(args ...interface{}) {
	defer z.sugarLogger.Sync()
	z.sugarLogger.Info(args)
}
func (z *zapLogger) Panic(args ...interface{}) {
	defer z.sugarLogger.Sync()
	z.sugarLogger.Panic(args)
}
func (z *zapLogger) Warn(args ...interface{}) {
	defer z.sugarLogger.Sync()
	z.sugarLogger.Warn(args)
}
