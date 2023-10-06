package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type District struct {
	Version string `yaml:"version"`
	Token   string `yaml:"token"`
}

var config District

const configFile = "config.yml"

func init() {
	yml, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(yml, &config)
	if err != nil {
		log.Fatal(err)
	}
	config.Token = os.ExpandEnv(config.Token)
}
