package server

import (
	"PokerGameAPI/domain/deck"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_CreateDeck(t *testing.T) {
	s := NewServer(deck.NewDeckMemoryRepository())

	req, err := http.NewRequest("GET", "api/create", nil)

	req.URL.Query().Add("shuffled", "true")

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.CreateDeck)
	handler.ServeHTTP(rr, req)

	AssertStatusCode(http.StatusOK, rr.Code, t)

	var deckResponse *DeckCreatedResponse = &DeckCreatedResponse{}

	jsonErr := json.Unmarshal(rr.Body.Bytes(), deckResponse)

	require.NoError(t, jsonErr)

	assert.Equal(t, false, deckResponse.Shuffled)
	assert.Equal(t, 52, deckResponse.Remaining)
}

func TestServer_OpenDeckMissingDeckID(t *testing.T) {
	s := NewServer(deck.NewDeckMemoryRepository())

	req, err := http.NewRequest("GET", "api/open", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.OpenDeck)
	handler.ServeHTTP(rr, req)

	AssertStatusCode(http.StatusBadRequest, rr.Code, t)

	AssertErrorIsNotEmpty(rr, t)
}

func TestServer_OpenDeck(t *testing.T) {
	server := NewServer(deck.NewDeckMemoryRepository())

	createdDeck := DoCreateDeckRequest(*server, t)

	deckId := createdDeck.Id

	d := DoOpenDeckRequest(*server, deckId, t)

	assert.Equal(t, deckId, d.GetId())
}

func TestServer_DrawCardMissingCount(t *testing.T) {
	server := NewServer(deck.NewDeckMemoryRepository())

	req, err := http.NewRequest("GET", "api/draw", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.DrawCards)
	handler.ServeHTTP(rr, req)

	AssertStatusCode(http.StatusBadRequest, rr.Code, t)

	AssertErrorIsNotEmpty(rr, t)
}

func TestServer_DrawCard(t *testing.T) {
	server := NewServer(deck.NewDeckMemoryRepository())

	createdDeck := DoCreateDeckRequest(*server, t)

	req, err := http.NewRequest("GET", "api/draw", nil)

	q := req.URL.Query()
	q.Add("deck_id", createdDeck.Id)
	q.Add("count", "10")
	req.URL.RawQuery = q.Encode()

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.DrawCards)
	handler.ServeHTTP(rr, req)

	AssertStatusCode(http.StatusOK, rr.Code, t)

	d := DoOpenDeckRequest(*server, createdDeck.Id, t)

	assert.Equal(t, 42, d.GetRemainingCards())
}

func AssertErrorIsNotEmpty(rr *httptest.ResponseRecorder, t *testing.T) {
	var errResponse *ErrorResponse = &ErrorResponse{}

	jsonErr := json.Unmarshal(rr.Body.Bytes(), errResponse)

	require.NoError(t, jsonErr)

	require.NotEmpty(t, errResponse.Error)
}

func AssertStatusCode(expectedCode int, actualCode int, t *testing.T) {
	if expectedCode != actualCode {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			actualCode, expectedCode)
	}
}

func DoCreateDeckRequest(server Server, t *testing.T) DeckCreatedResponse {

	req, err := http.NewRequest("GET", "api/create", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.CreateDeck)
	handler.ServeHTTP(rr, req)

	AssertStatusCode(http.StatusOK, rr.Code, t)

	var deckResponse *DeckCreatedResponse = &DeckCreatedResponse{}

	err = json.Unmarshal(rr.Body.Bytes(), deckResponse)

	require.NoError(t, err)

	return *deckResponse
}

func DoOpenDeckRequest(server Server, deckId string, t *testing.T) deck.IDeck {
	openRecorder := httptest.NewRecorder()
	openRequest, openReqErr := http.NewRequest("GET", "api/open", nil)

	q := openRequest.URL.Query()
	q.Add("deck_id", deckId)
	openRequest.URL.RawQuery = q.Encode()

	if openReqErr != nil {
		t.Fatal(openReqErr)
	}

	handler := http.HandlerFunc(server.OpenDeck)
	handler.ServeHTTP(openRecorder, openRequest)

	d := &deck.PokerDeck{}

	_ = json.Unmarshal(openRecorder.Body.Bytes(), d)

	return d
}
