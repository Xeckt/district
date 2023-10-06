package main

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	if len(Config.Bot.Token) == 0 {
		Dislog.Fatal("Token is empty!")
		return
	}
	
	dg, err := discordgo.New("Bot " + Config.Bot.Token)
	if err != nil {
		Dislog.Fatal(err.Error())
		return
	}

	dg.AddHandler(MemberJoined)
	dg.AddHandler(MemberLeft)
	dg.AddHandler(MessageCreated)

	dg.Identify.Intents |= discordgo.IntentsGuilds
	dg.Identify.Intents |= discordgo.IntentsGuildMembers
	dg.Identify.Intents |= discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		Dislog.Fatal(err.Error())
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func MemberJoined(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	Dislog.Info("Member joined guild", zap.String("member", m.Member.User.String()))
}

func MemberLeft(s *discordgo.Session, m *discordgo.GuildMemberRemove) {
	Dislog.Info("Member left guild", zap.String("member", m.Member.User.String()))
}

func MessageCreated(s *discordgo.Session, m *discordgo.MessageCreate) {
	Dislog.Info("Message created", zap.String("message", m.Content))
}
