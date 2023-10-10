package main

import "github.com/bwmarrin/discordgo"

func SpecifyIntents(im HandlerManager) {
	im.AddIntents(
		discordgo.IntentsGuilds,
		discordgo.IntentsGuildMembers,
		discordgo.IntentsGuildMessages,
	)
}
