package deck

import (
	"PokerGameAPI/domain/card"
	"errors"
	"github.com/google/uuid"
	"math/rand"
)

const DefaultCardCount = 52

type PokerDeck struct {
	Id        string      `json:"deck_id"`
	Shuffled  bool        `json:"shuffled"`
	Remaining int         `json:"remaining"`
	Cards     []card.Card `json:"cards"`
}

func NewPokerDeck(shuffled bool, cards []card.Card) PokerDeck {
	if len(cards) == 0 {
		cards = BuildCards()
	}

	if shuffled {
		rand.Shuffle(len(cards), func(i, j int) {
			cards[i], cards[j] = cards[j], cards[i]
		})
	}

	return PokerDeck{Id: uuid.New().String(), Shuffled: shuffled, Cards: cards, Remaining: len(cards)}
}

func (deck PokerDeck) GetId() string {
	return deck.Id
}

func (deck PokerDeck) IsShuffled() bool {
	return deck.Shuffled
}

func (deck PokerDeck) GetRemainingCards() int {
	return deck.Remaining
}

func (deck PokerDeck) DrawCards(count int) ([]card.Card, IDeck, error) {
	if count > deck.Remaining {
		return nil, nil, errors.New("There are not enough cards in the deck.")
	}

	cards := deck.Cards[0:count]

	deck.Cards = deck.Cards[count:]
	deck.Remaining -= count

	return cards, deck, nil
}

func BuildCards() []card.Card {
	cards := make([]card.Card, DefaultCardCount)
	idx := 0

	for _, suite := range card.Suites {
		for _, value := range card.Values {
			cards[idx] = *card.NewCard(suite, value)
			idx++
		}
	}

	return cards
}
