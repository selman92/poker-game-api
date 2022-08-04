package main

import (
	"PokerGameAPI/domain/deck"
	"PokerGameAPI/server"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	LoadEnvironmentVariables()

	s := server.NewServer(deck.NewDeckMemoryRepository())

	s.Start()
}

func LoadEnvironmentVariables() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
