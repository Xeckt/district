package internal

import (
	"github.com/bwmarrin/discordgo"
	"log/slog"
)

func MessageCreated(s *discordgo.Session, m *discordgo.MessageCreate) {
	Dislog.Info("Message created", slog.String("message", m.Content))
}

func AddMessageHandler(dh DistrictHandler) {
	dh.AddHandlers(
		MessageCreated,
	)
}
