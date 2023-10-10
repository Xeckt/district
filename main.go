package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func cleanup(s *discordgo.Session) {
	Dislog.Info("Cleaning up")
	go func() {
		Dislog.Info("Shutting down district session")
		err := s.Close()
		if err != nil {
			Dislog.Error("Error shutting down district", err)
		}
		Dislog.Info("Closing log file")
		err = logFile.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Done. Goodbye!")
	}()
}

func main() {
	if len(Config.Bot.Token) == 0 {
		Dislog.Error("Token is empty!")
		return
	}

	dg, err := discordgo.New("Bot " + Config.Bot.Token)
	if err != nil {
		Dislog.Error("Could not create session with district", err)
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
		Dislog.Error("Error opening district session", err)
		return
	}

	Dislog.Info("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	defer cleanup(dg)
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
