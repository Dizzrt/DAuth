package dlog

import (
	"time"

	"github.com/Dizzrt/DAuth/backend/common/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const timeFormat = "2006-01-02 15:04:05 Z07"

var (
	logger  *zap.Logger
	slogger *zap.SugaredLogger
)

// -------------------- logger init start --------------------
func init() {
	hook := &lumberjack.Logger{
		Filename:   config.LogGetPath(),
		MaxSize:    config.LogGetMaxSize(),
		MaxBackups: config.LogGetMaxBackups(),
		MaxAge:     config.LogGetMaxAge(),
		LocalTime:  config.LogGetIsLocalTime(),
		Compress:   config.LogGetIsComporess(),
	}

	writeSyncer := zapcore.BufferedWriteSyncer{
		WS:   zapcore.AddSync(hook),
		Size: 4096,
	}

	cores := []zapcore.Core{
		zapcore.NewCore(encoder(), zapcore.AddSync(&writeSyncer), zapcore.DebugLevel),
	}

	tee := zapcore.NewTee(cores...)
	l := zap.New(tee).WithOptions(zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))

	logger = l
	slogger = l.Sugar()
}

func encoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(
		zapcore.EncoderConfig{
			MessageKey:     "msg",
			LevelKey:       "level",
			TimeKey:        "time",
			NameKey:        "logger",
			CallerKey:      "caller",
			StacktraceKey:  "stacktrace",
			FunctionKey:    zapcore.OmitKey,
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeDuration: zapcore.MillisDurationEncoder,
			EncodeLevel:    levelEncoder,
			EncodeCaller:   callerEncoder,
			EncodeTime:     timeEncoder,
		})
}

func levelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + t.Format(timeFormat) + "]")
}

func callerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + caller.TrimmedPath() + "]")
}

// -------------------- logger init end --------------------

func Sync() error {
	return logger.Sync()
}

func L() *zap.Logger {
	return logger.WithOptions(zap.AddCallerSkip(-1))
}

func SL() *zap.SugaredLogger {
	return slogger.WithOptions(zap.AddCallerSkip(-1))
}

// -------------------- slogger functions mapping --------------------
func Debug(args ...interface{}) {
	slogger.Debug(args...)
}

func Info(args ...interface{}) {
	slogger.Info(args...)
}

func Warn(args ...interface{}) {
	slogger.Warn(args...)
}

func Error(args ...interface{}) {
	slogger.Error(args...)
}

func DPanic(args ...interface{}) {
	slogger.DPanic(args...)
}

func Panic(args ...interface{}) {
	slogger.Panic(args...)
}

func Fatal(args ...interface{}) {
	slogger.Fatal(args...)
}

func Debugf(template string, args ...interface{}) {
	slogger.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	slogger.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	slogger.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	slogger.Errorf(template, args...)
}

func DPanicf(template string, args ...interface{}) {
	slogger.DPanicf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	slogger.Panicf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	slogger.Fatalf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	slogger.Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	slogger.Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	slogger.Warnw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	slogger.Errorw(msg, keysAndValues...)
}

func DPanicw(msg string, keysAndValues ...interface{}) {
	slogger.DPanicw(msg, keysAndValues...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	slogger.Panicw(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	slogger.Fatalw(msg, keysAndValues...)
}

func Debugln(args ...interface{}) {
	slogger.Debugln(args...)
}

func Infoln(args ...interface{}) {
	slogger.Infoln(args...)
}

func Warnln(args ...interface{}) {
	slogger.Warnln(args...)
}

func Errorln(args ...interface{}) {
	slogger.Errorln(args...)
}

func DPanicln(args ...interface{}) {
	slogger.DPanicln(args...)
}

func Panicln(args ...interface{}) {
	slogger.Panicln(args...)
}

func Fatalln(args ...interface{}) {
	slogger.Fatalln(args...)
}

// -------------------- logger functions mapping --------------------
func FDebug(msg string, fields ...zapcore.Field) {
	logger.Debug(msg, fields...)
}

func FInfo(msg string, fields ...zapcore.Field) {
	logger.Info(msg, fields...)
}

func FWarn(msg string, fields ...zapcore.Field) {
	logger.Warn(msg, fields...)
}

func FError(msg string, fields ...zapcore.Field) {
	logger.Error(msg, fields...)
}

func FDPanic(msg string, fields ...zapcore.Field) {
	logger.DPanic(msg, fields...)
}

func FPanic(msg string, fields ...zapcore.Field) {
	logger.Panic(msg, fields...)
}

func FFatal(msg string, fields ...zapcore.Field) {
	logger.Fatal(msg, fields...)
}
