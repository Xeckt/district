package main

import (
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

func MemberJoined(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	Dislog.Info("Member joined guild", slog.String("member", m.Member.User.String()))
}

func MemberLeft(s *discordgo.Session, m *discordgo.GuildMemberRemove) {
	Dislog.Info("Member left guild", slog.String("member", m.Member.User.String()))
}

func AddMemberHandler(dh DistrictHandler) {
	dh.AddHandlers(
		MemberJoined,
		MemberLeft,
	)
}
