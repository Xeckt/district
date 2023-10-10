package main

import (
	"github.com/bwmarrin/discordgo"
	"log/slog"
	"reflect"
	"runtime"
)

type DistrictHandler interface {
	AddHandlers(handler ...any)
}

type HandlerManager struct {
	*discordgo.Session
}

func (a HandlerManager) AddHandlers(handlers ...any) {
	for _, h := range handlers {
		a.AddHandler(h)
		Dislog.Info("Handler added", slog.String("Handler", runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()))
	}
}
