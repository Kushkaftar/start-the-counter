package main

import (
	"github.com/Kushkaftar/start-the-counter/iternal/app/apiserver"

	"flag"
	"github.com/BurntSushi/toml"
	"log"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "путь до конфигурациооного файла")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
