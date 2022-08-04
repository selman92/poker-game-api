package deck

import (
	"PokerGameAPI/domain/card"
	"strings"
)

func NewDeck(shuffled bool, cardCodes string) IDeck {
	if cardCodes != "" {
		cards := CreateCardsFromCodes(cardCodes)

		return NewPokerDeck(shuffled, cards)
	}

	return NewPokerDeck(shuffled, []card.Card{})
}

func CreateCardsFromCodes(codes string) []card.Card {
	parsedCodes := ParseCardCodes(codes)

	cards := make([]card.Card, 0)

	for _, c := range parsedCodes {
		if c == "" {
			continue
		}

		newCard, err := card.NewCardFromCode(c)

		// Ignore invalid values.
		if err != nil {
			continue
		}

		cards = append(cards, *newCard)
	}

	return cards
}

func ParseCardCodes(codes string) []string {
	return strings.Split(codes, ",")
}
