package main

import (
	"PokerGameAPI/domain/deck"
	"PokerGameAPI/server"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	LoadEnvironmentVariables()

	s := server.NewServer(deck.NewDeckMemoryRepository())

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/api/create", s.CreateDeck)
	router.Get("/api/open", s.OpenDeck)
	router.Get("/api/draw", s.DrawCards)

	fmt.Println("Server started listening on port: " + os.Getenv("PORT"))

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}

func LoadEnvironmentVariables() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
