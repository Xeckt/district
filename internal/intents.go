package internal

import (
	"github.com/bwmarrin/discordgo"
)

type IntentHandler interface {
	SpecifyIntents(...any)
}

type IntentManager struct {
	*discordgo.Session
}

func (im IntentManager) SpecifyIntents(intents ...discordgo.Intent) {
	for _, i := range intents {
		im.Identify.Intents |= i
	}
}

func AddIntents(im IntentManager) {
	im.SpecifyIntents(
		discordgo.IntentsGuilds,
		discordgo.IntentsGuildMembers,
		discordgo.IntentsGuildMessages,
	)
}
