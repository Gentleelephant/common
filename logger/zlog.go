package logger

import (
	"github.com/Gentleelephant/common/consts"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

const (
	logTmFmt = "2006-01-02 15:04:05"
)

func getDefaultEncoderConfig(v *viper.Viper) *zapcore.EncoderConfig {
	encoderConfig := &zapcore.EncoderConfig{
		MessageKey:       v.GetString(consts.LoggerMessageKey),
		LevelKey:         v.GetString(consts.LoggerLevelKey),
		TimeKey:          v.GetString(consts.LoggerTimeKey),
		NameKey:          v.GetString(consts.LoggerNameKey),
		CallerKey:        v.GetString(consts.LoggerCallerKey),
		FunctionKey:      v.GetString(consts.LoggerFunctionKey),
		StacktraceKey:    v.GetString(consts.LoggerStacktraceKey),
		LineEnding:       v.GetString(consts.LoggerLineEnding),
		EncodeLevel:      cEncodeLevel,
		EncodeTime:       cEncodeTime,
		EncodeDuration:   getDefaultEncodeDuration(),
		EncodeCaller:     cEncodeCaller,
		EncodeName:       getDefaultEncodeName(),
		ConsoleSeparator: viper.GetString(consts.LoggerSeparator),
	}
	return encoderConfig
}

func GetDefaultZapConfig(v *viper.Viper) *zap.Config {
	encoding := v.GetString(consts.LoggerEncoding)
	if encoding == "" {
		encoding = "console"
	}
	development := v.GetBool(consts.LoggerDevelopmentMode)
	disableCaller := v.GetBool(consts.LoggerDisableCaller)
	disableStacktrace := v.GetBool(consts.LoggerDisableStacktrace)

	config := zap.Config{
		Level:             zap.AtomicLevel{},
		Development:       development,
		DisableCaller:     disableCaller,
		DisableStacktrace: disableStacktrace,
		Sampling:          getDefaultSampling(),
		Encoding:          encoding,
		EncoderConfig:     *getDefaultEncoderConfig(v),
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stdout"},
		InitialFields:     nil,
	}
	return &config
}

func GetDefaultLogger(v *viper.Viper) *zap.Logger {
	encoder := GetDefaultEncoder(v)
	writeSyncer := GetWriteSyncer(v)
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, writeSyncer, getDefaultLevel(v)),
		zapcore.NewCore(zapcore.NewConsoleEncoder(*getDefaultEncoderConfig(v)), zapcore.AddSync(os.Stdout), getDefaultLevel(v)),
	)
	logger := zap.New(core, zap.AddCaller())
	return logger
}

// GetDefaultEncoder 自定义日志格式显示
func GetDefaultEncoder(v *viper.Viper) zapcore.Encoder {
	encoding := v.GetString(consts.LoggerEncoding)
	if encoding == "" || encoding == "json" {
		return zapcore.NewJSONEncoder(*getDefaultEncoderConfig(v))
	}
	return zapcore.NewConsoleEncoder(*getDefaultEncoderConfig(v))
}

// GetWriteSyncer 自定义的WriteSyncer
func GetWriteSyncer(v *viper.Viper) zapcore.WriteSyncer {
	maxSize := v.GetInt(consts.LoggerMaxSize)
	if maxSize == 0 {
		maxSize = 100
	}
	maxBackups := v.GetInt(consts.LoggerMaxBackups)
	if maxBackups == 0 {
		maxBackups = 5
	}
	maxAge := v.GetInt(consts.LoggerMaxAge)
	if maxAge == 0 {
		maxAge = 1
	}
	lumberJackLogger := &lumberjack.Logger{
		Filename:   v.GetString(consts.LoggerOutputPath),
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// cEncodeLevel 自定义日志级别显示
func cEncodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(level.CapitalString())
}

// cEncodeTime 自定义时间格式显示
func cEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(logTmFmt))
}

// cEncodeCaller 自定义行号显示
func cEncodeCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(caller.TrimmedPath())
}

func getDefaultEncodeDuration() zapcore.DurationEncoder {
	return zapcore.SecondsDurationEncoder
}

func getDefaultEncodeName() zapcore.NameEncoder {
	return zapcore.FullNameEncoder
}

func getDefaultSampling() *zap.SamplingConfig {
	return nil
}

func getDefaultLevel(v *viper.Viper) zapcore.Level {
	level := v.GetString(consts.LoggerLevel)
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
