package deck

import (
	"PokerGameAPI/domain/card"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPokerdeck_CreateDeck(t *testing.T) {
	deck := NewPokerDeck(false, []card.Card{})

	assert.Equal(t, 52, deck.Remaining)
}

func TestPokerdeck_CreateShuffledDeck(t *testing.T) {
	deck := NewPokerDeck(true, []card.Card{})

	assert.Equal(t, 52, deck.Remaining)
	assert.Equal(t, true, deck.Shuffled)

	cards, err := deck.DrawCards(1)

	require.NoError(t, err)

	assert.NotEqual(t, "AS", cards[0].Code)
}

func TestPokerdeck_CreateCustomDeck(t *testing.T) {
	deck := NewPokerDeck(false, []card.Card{
		*card.NewCardFromCode("AS"),
		*card.NewCardFromCode("1D"),
		*card.NewCardFromCode("KH"),
	})

	assert.Equal(t, 3, deck.Remaining)

	cards, err := deck.DrawCards(3)

	require.NoError(t, err)

	assert.Equal(t, "AS", cards[0].Code)
	assert.Equal(t, "SPADES", cards[0].Suite)
	assert.Equal(t, "A", cards[0].Value)

	assert.Equal(t, "1D", cards[1].Code)
	assert.Equal(t, "DIAMONDS", cards[1].Suite)
	assert.Equal(t, "1", cards[1].Value)

	assert.Equal(t, "KH", cards[2].Code)
	assert.Equal(t, "HEARTS", cards[2].Suite)
	assert.Equal(t, "KING", cards[2].Value)
}

func TestPokerDeck_DrawCard(t *testing.T) {
	deck := NewPokerDeck(false, []card.Card{})

	assert.Equal(t, 52, deck.Remaining)

	deck.DrawCards(4)

	assert.Equal(t, 48, deck.Remaining)
}

func TestPokerDeck_DrawCardError(t *testing.T) {
	deck := NewPokerDeck(false, []card.Card{
		card.NewCardFromCode("AS"),
	})

	assert.Equal(t, 1, deck.Remaining)

	_, err := deck.DrawCards(2)

	require.Error(t, err)
}
