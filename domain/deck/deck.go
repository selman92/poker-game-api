package deck

import "PokerGameAPI/domain/card"

type IDeck interface {
	DrawCards(count int) ([]card.Card, IDeck, error)
	GetId() string
	IsShuffled() bool
	GetRemainingCards() int
}
