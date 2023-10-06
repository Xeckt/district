package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Dislog *zap.Logger

const mainLog = "district.log"

func init() {
	Dislog = create()
	defer Dislog.Sync()
}

func create() *zap.Logger {
	consoleConfig := zapcore.EncoderConfig{
		LevelKey:       "L",
		NameKey:        "N",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "M",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	consoleEncoder := zapcore.NewConsoleEncoder(consoleConfig)
	consoleLogLevel := zapcore.InfoLevel

	logFile, _ := os.OpenFile(config.LogDir+"/"+mainLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(logFile)

	var core zapcore.Core

	if !config.EnableLog {
		if config.EnableDebug {
			return zap.New(zapcore.NewCore(consoleEncoder, os.Stdout, consoleLogLevel),
				zap.AddCaller(),
				zap.AddStacktrace(zapcore.ErrorLevel),
			)
		}
		return zap.New(zapcore.NewCore(consoleEncoder, os.Stdout, consoleLogLevel))
	}

	fileConfig := zap.NewProductionEncoderConfig()
	fileConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(fileConfig)
	fileLogLevel := zapcore.InfoLevel

	core = zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, fileLogLevel),
		zapcore.NewCore(consoleEncoder, os.Stdout, consoleLogLevel),
	)

	if config.EnableDebug {
		return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.DebugLevel))
	}
	
	return zap.New(core)
}
