package main

import (
	"log"
	"log/slog"
	"os"

	slogmulti "github.com/samber/slog-multi"
)

var Dislog *slog.Logger
var LogFile *os.File

func init() {
	Dislog = create()
}

const (
	mainLog = "district.log"
)

func create() *slog.Logger {
	var hOpts slog.HandlerOptions
	var l *slog.Logger

	err := os.MkdirAll(Config.Bot.LogDir, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
	LogFile, err = os.OpenFile(Config.Bot.LogDir+"/"+mainLog, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
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
	if !Config.Bot.EnableLogFile {
		l = slog.New(slog.NewTextHandler(os.Stdout, &hOpts))
		return l
	}
	l = slog.New(
		slogmulti.Fanout(
			slog.NewJSONHandler(LogFile, &hOpts),
			slog.NewTextHandler(os.Stdout, &hOpts),
		),
	)
	return l
}
