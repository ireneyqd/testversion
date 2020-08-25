package testversion

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Init() {
	settingLogger()
	settingJlogger()
	//defer loggerr.Sync()
}

var logger *zap.Logger

func settingLogger() {
	jsonEncoder := getJSONEncoder()
	consoleEncoder := getConsoleEncoder()
	writeFileSyncer := getLogFileWriter()
	writeConsoleSyncer := getLogConsoleWriter()
	core := zapcore.NewTee(
		zapcore.NewCore(jsonEncoder, writeFileSyncer, zapcore.DebugLevel),
		zapcore.NewCore(consoleEncoder, writeConsoleSyncer, zapcore.DebugLevel),
	)

	logger = zap.New(
		core, zap.AddCaller(),
		zap.AddCallerSkip(2),
		zap.Fields(zap.String("type", "jlog")))
}

func getJSONEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogFileWriter() zapcore.WriteSyncer {
	fileWriter := &lumberjack.Logger{
		Filename: "./test.log",
		// When log file size achieves to MaxSize, logrus will compression that.
		MaxSize: 100, // megabytes
		// When backups count larger than MaxBackups, logrus will remove the oldest backups.
		MaxBackups: 100,
		MaxAge:     60,   // days
		Compress:   true, // disabled by default
	}
	return zapcore.AddSync(fileWriter)
}

func getConsoleEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogConsoleWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}

type jlogger struct {
	sugar   *zap.SugaredLogger
}

var jloggerInstance jlogger

func settingJlogger() {
	jloggerInstance = jlogger{
		sugar: logger.Sugar(),
	}
}
