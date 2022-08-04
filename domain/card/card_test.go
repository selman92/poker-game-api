package card

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCard_CreateFromCode(t *testing.T) {
	card, err := NewCardFromCode("AC")

	assert.NoError(t, err)

	assert.Equal(t, "A", card.Value)
	assert.Equal(t, "CLUBS", card.Suite)
	assert.Equal(t, "AC", card.Code)
}
