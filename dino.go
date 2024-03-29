package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/echeadle/dino/dynowebportal"
)

type configuration struct {
	ServerAddress string `json:"webserver"`
}

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}

	config := new(configuration)
	json.NewDecoder(file).Decode(config)
	log.Println("Starting web server on address ", config.ServerAddress)
	dynowebportal.RunWebPortal(config.ServerAddress)
}
