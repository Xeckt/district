package main

import (
	"log/slog"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func MessageCreated(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	Dislog.Info("Message created", slog.String("message", m.Content))
	ListUsage(s, m)
}

func ListUsage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "!user") {
		Dislog.Info("!user command found", slog.String("user_id", m.Author.ID))
		queryUser(s, m)
	} else if strings.HasPrefix(m.Content, "!server") {
		Dislog.Info("!server command found", slog.String("server_id", m.GuildID))
		queryServer(s, m)
	} else {
		// don't update on command
		Dislog.Info("Updating message count", slog.String("server_id", m.GuildID), slog.String("user_id", m.Author.ID))
		updateMessageCount(db, m.GuildID, m.Author.ID)
	}
}

func AddMessageHandler(dh DistrictHandler) {
	dh.AddHandlers(
		MessageCreated,
	)
}
