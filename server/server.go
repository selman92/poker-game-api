package server

import (
	"PokerGameAPI/domain/deck"
	"encoding/json"
	"net/http"
	"strconv"
)

type Server struct {
	repository deck.Repository
}

func NewServer(repository deck.Repository) *Server {
	return &Server{repository: repository}
}

type DeckCreatedResponse struct {
	Id        string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (s Server) CreateDeck(w http.ResponseWriter, r *http.Request) {
	shuffled, err := strconv.ParseBool(r.URL.Query().Get("shuffled"))

	if err != nil {
		shuffled = false
	}

	cardCodes := r.URL.Query().Get("cards")

	newDeck := deck.NewDeck(shuffled, cardCodes)

	s.repository.Add(newDeck)

	createdResponse := DeckCreatedResponse{Id: newDeck.GetId(), Shuffled: newDeck.IsShuffled(), Remaining: newDeck.GetRemainingCards()}

	responseBytes, _ := json.Marshal(createdResponse)

	w.Header().Add("Content-Type", "application/json")

	if _, writeErr := w.Write(responseBytes); writeErr != nil {
		http.Error(w, CreateErrorResponse("Could not write to the response."), http.StatusInternalServerError)
		return
	}
}

func (s Server) OpenDeck(w http.ResponseWriter, r *http.Request) {
	deckId := r.URL.Query().Get("deck_id")

	if deckId == "" {
		http.Error(w, CreateErrorResponse("Deck ID must be provided."), http.StatusBadRequest)
		return
	}

	d, err := s.repository.Get(deckId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deckBytes, _ := json.Marshal(d)

	w.Header().Add("Content-Type", "application/json")

	if _, writeErr := w.Write(deckBytes); writeErr != nil {
		http.Error(w, CreateErrorResponse("Could not write to the response."), http.StatusInternalServerError)
		return
	}
}

func (s Server) DrawCards(w http.ResponseWriter, r *http.Request) {
	deckId := r.URL.Query().Get("deck_id")

	if deckId == "" {
		http.Error(w, CreateErrorResponse("Deck ID must be provided."), http.StatusBadRequest)
		return
	}

	count := r.URL.Query().Get("count")

	if count == "" {
		http.Error(w, CreateErrorResponse("Count must be provided."), http.StatusBadRequest)
		return
	}

	countInt, parseErr := strconv.Atoi(count)
	if parseErr != nil {
		http.Error(w, CreateErrorResponse("Count is invalid."), http.StatusBadRequest)
		return
	}

	deckToDraw, err := s.repository.Get(deckId)

	if err != nil {
		http.Error(w, CreateErrorResponse(err.Error()), http.StatusBadRequest)
		return
	}

	cards, updatedDeck, drawErr := deckToDraw.DrawCards(countInt)

	if drawErr != nil {
		http.Error(w, CreateErrorResponse(drawErr.Error()), http.StatusBadRequest)
		return
	}

	s.repository.Update(updatedDeck)

	responseBytes, _ := json.Marshal(cards)

	w.Header().Add("Content-Type", "application/json")
	w.Write(responseBytes)
}

func CreateErrorResponse(message string) string {
	errResponse := ErrorResponse{Error: message}

	bytes, _ := json.Marshal(errResponse)

	return string(bytes)
}
