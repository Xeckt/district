package main

import (
	slogmulti "github.com/samber/slog-multi"
	"log"
	"log/slog"
	"os"
)

var Dislog *slog.Logger
var logFile *os.File

const mainLog = "district.log"

func init() {
	Dislog = create()
}

func create() *slog.Logger {
	var hOpts slog.HandlerOptions
	var l *slog.Logger

	err := os.MkdirAll(Config.Bot.LogDir, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
	logFile, err = os.OpenFile(Config.Bot.LogDir+"/"+mainLog, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	if Config.Bot.EnableDebug {
		hOpts = slog.HandlerOptions{
			AddSource:   true,
			Level:       slog.LevelDebug,
			ReplaceAttr: nil,
		}
	}
	if Config.Bot.EnableLogFile {
		l = slog.New(slog.NewTextHandler(os.Stdout, &hOpts))
		return l
	}
	l = slog.New(
		slogmulti.Fanout(
			slog.NewJSONHandler(logFile, &hOpts),
			slog.NewTextHandler(os.Stdout, &hOpts),
		),
	)
	return l
}
