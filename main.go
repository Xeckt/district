package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	Token := ""

	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatalln("Failed to create Discord session:", err)
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
		fmt.Println("Failed to open Discord websocket:", err)
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
	fmt.Println("Member joined!")
}

func MemberLeft(s *discordgo.Session, m *discordgo.GuildMemberRemove) {
	fmt.Println("Member left!")
}

func MessageCreated(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println("Message created!")
}
