package main

import (
	"PokerGameAPI/domain/deck"
	"PokerGameAPI/server"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {
	server := server.NewServer(deck.NewDeckMemoryRepository())

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/api/create", server.CreateDeck)
	router.Get("/api/open", server.OpenDeck)
	router.Get("/api/draw", server.DrawCards)

	log.Fatal(http.ListenAndServe(":6592", router))
}
