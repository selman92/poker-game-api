package deck

type Repository interface {
	Add(deck IDeck)
	Update(deck IDeck)
	Get(deckId string) (IDeck, error)
}
