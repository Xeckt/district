package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type District struct {
	Bot BotConfig `yaml:"district"`
}

type BotConfig struct {
	Version       string `yaml:"version"`
	Token         string `yaml:"token"`
	EnableLogFile bool   `yaml:"enableLogFile"`
	EnableDebug   bool   `yaml:"enableDebug"`
	LogDir        string `yaml:"logDir"`
}

var Config District

const configFile = "config.yml"

func init() {
	yml, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(yml, &Config)
	if err != nil {
		log.Fatal(err)
	}
	Config.Bot.Token = os.ExpandEnv(Config.Bot.Token)
}
