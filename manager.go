package main

import (
	"github.com/bwmarrin/discordgo"
	"log/slog"
	"reflect"
	"runtime"
)

type DistrictHandler interface {
	AddHandlers(handler ...any)
	AddIntents(intents ...discordgo.Intent)
}

type HandlerManager struct {
	*discordgo.Session
}

func (hm HandlerManager) AddHandlers(handlers ...any) {
	for _, h := range handlers {
		hm.AddHandler(h)
		Dislog.With(
			slog.Group("handler",
				slog.String("name", runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()),
			),
		).Info("Handler added")
	}
}

func (hm HandlerManager) AddIntents(intents ...discordgo.Intent) {
	for _, i := range intents {
		hm.Identify.Intents |= i
		Dislog.With(
			slog.Group("intent",
				slog.String("name", IntentString(i)),
			),
		).Info("intent specified")
	}
}
