package main

import (
	"district/internal"
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
	internal.Dislog.Info("Shutting down district session")
	err := s.Close()
	if err != nil {
		internal.Dislog.Error("Error shutting down district", err)
	}
	internal.Dislog.Info("Closing log file")
	err = internal.LogFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done. Goodbye!")
}

func setupRequirements(s *discordgo.Session) {
	dh := internal.HandlerManager{Session: s}
	im := internal.IntentManager{Session: s}

	internal.AddMemberHandler(dh)
	internal.AddMessageHandler(dh)
	internal.AddIntents(im)
}

func main() {
	var wg sync.WaitGroup

	if len(internal.Config.Bot.Token) == 0 {
		internal.Dislog.Error("Token is empty!")
		return
	}

	dg, err := discordgo.New("Bot " + internal.Config.Bot.Token)
	if err != nil {
		internal.Dislog.Error("Could not create session with district", err)
		return
	}

	err = dg.Open()
	if err != nil {
		internal.Dislog.Error("Error opening district session", err)
		return
	}

	setupRequirements(dg)

	internal.Dislog.Info("district is now running. CTRL+C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	wg.Add(1)
	go cleanup(dg, &wg)
	wg.Wait()
}
