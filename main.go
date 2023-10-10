package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func cleanup(s *discordgo.Session, wg *sync.WaitGroup) {
	defer wg.Done()
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
}

func main() {
	var wg sync.WaitGroup

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

	wg.Add(1)
	go cleanup(dg, &wg)
	wg.Wait()
}
