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

	cards, _, err := deck.DrawCards(1)

	require.NoError(t, err)

	assert.NotEqual(t, "AS", cards[0].Code)
}

func TestPokerdeck_CreateCustomDeck(t *testing.T) {

	c1, _ := card.NewCardFromCode("AS")
	c2, _ := card.NewCardFromCode("2D")
	c3, _ := card.NewCardFromCode("KH")

	deck := NewPokerDeck(false, []card.Card{
		*c1, *c2, *c3,
	})

	assert.Equal(t, 3, deck.Remaining)

	cards, _, err := deck.DrawCards(3)

	require.NoError(t, err)

	assert.Equal(t, "AS", cards[0].Code)
	assert.Equal(t, "SPADES", cards[0].Suite)
	assert.Equal(t, "A", cards[0].Value)

	assert.Equal(t, "2D", cards[1].Code)
	assert.Equal(t, "DIAMONDS", cards[1].Suite)
	assert.Equal(t, "2", cards[1].Value)

	assert.Equal(t, "KH", cards[2].Code)
	assert.Equal(t, "HEARTS", cards[2].Suite)
	assert.Equal(t, "KING", cards[2].Value)
}

func TestPokerDeck_DrawCard(t *testing.T) {
	deck := NewPokerDeck(false, []card.Card{})

	assert.Equal(t, 52, deck.Remaining)

	_, updatedDeck, _ := deck.DrawCards(4)

	assert.Equal(t, 48, updatedDeck.GetRemainingCards())
}

func TestPokerDeck_DrawCardError(t *testing.T) {
	c, _ := card.NewCardFromCode("AS")

	deck := NewPokerDeck(false, []card.Card{*c})

	assert.Equal(t, 1, deck.Remaining)

	_, _, err := deck.DrawCards(2)

	require.Error(t, err)
}
