package main

import (
	"GolangWebsocket/internal/wsserver"
	"github.com/BurntSushi/toml"
	"log"
)

var (
	configPath string = "config/wsserver.toml"
)

func main() {
	config := wsserver.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	server := wsserver.NewServer(config)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
