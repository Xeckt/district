package main

import (
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

func MessageCreated(s *discordgo.Session, m *discordgo.MessageCreate) {
	Dislog.Info("Message created", slog.String("message", m.Content))
}

func AddMessageHandler(dh DistrictHandler) {
	dh.AddHandlers(
		MessageCreated,
	)
}
