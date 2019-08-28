package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"

	"app/config"
	"app/server"
)

func main() {
	conf := config.New()

	err := envconfig.Process("ART", conf)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = server.Initialise(conf)
	if err != nil {
		log.Fatal(err.Error())
	}
}