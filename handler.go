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
		Dislog.Info("Handler added", slog.String("Handler", runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()))
	}
}

func (hm HandlerManager) AddIntents(intents ...discordgo.Intent) {
	for _, i := range intents {
		hm.Identify.Intents |= i
	}
}

func SpecifyIntents(im HandlerManager) {
	im.AddIntents(
		discordgo.IntentsGuilds,
		discordgo.IntentsGuildMembers,
		discordgo.IntentsGuildMessages,
	)
}
