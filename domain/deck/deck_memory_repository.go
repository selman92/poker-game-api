package deck

import "errors"

type DeckMemoryRepository struct {
	decks map[string]IDeck
}

func NewDeckMemoryRepository() *DeckMemoryRepository {
	return &DeckMemoryRepository{decks: make(map[string]IDeck)}
}

func (repository *DeckMemoryRepository) Add(deck IDeck) {
	repository.decks[deck.GetId()] = deck
}

func (repository *DeckMemoryRepository) Update(deck IDeck) {
	repository.decks[deck.GetId()] = deck
}

func (repository *DeckMemoryRepository) Get(deckId string) (IDeck, error) {
	deck, exists := repository.decks[deckId]

	if !exists {
		return deck, errors.New("The deck with the specified ID doesn't exist.")
	}

	return deck, nil
}
