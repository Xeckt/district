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
	err := os.MkdirAll(Config.Bot.LogDir, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
	logFile, err = os.OpenFile(Config.Bot.LogDir+"/"+mainLog, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	l := slog.New(
		slogmulti.Fanout(
			slog.NewJSONHandler(logFile, nil),
			slog.NewTextHandler(os.Stdout, nil),
		),
	)
	return l
}
