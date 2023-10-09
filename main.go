package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	if len(Config.Bot.Token) == 0 {
		Dislog.Info("test?!?!?!?!?")
		Dislog.Error("Token is empty!")
		return
	}

	dg, err := discordgo.New("Bot " + Config.Bot.Token)
	if err != nil {
		Dislog.Error(err.Error())
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
		Dislog.Error(err.Error())
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	Dislog.Info("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	Dislog.Info("Shutting down District")
	err = dg.Close()
	if err != nil {
		Dislog.Error(err.Error())
	}
	Dislog.Info("Closing log file.")
	err = logFile.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func MemberJoined(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	Dislog.Info("Member joined guild", slog.String("member", m.Member.User.String()))
}

func MemberLeft(s *discordgo.Session, m *discordgo.GuildMemberRemove) {
	Dislog.Info("Member left guild", slog.String("member", m.Member.User.String()))
}

func MessageCreated(s *discordgo.Session, m *discordgo.MessageCreate) {
	Dislog.Info("Message created", slog.String("message", m.Content))
}
