package card

import (
	"errors"
	"strings"
)

var (
	Values = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "JACK", "QUEEN", "KING"}
	Suites = []string{"SPADES", "DIAMONDS", "CLUBS", "HEARTS"}
)

type Card struct {
	Suite string `json:"suite"`
	Value string `json:"value"`
	Code  string `json:"code"`
}

func NewCard(suite string, value string) *Card {
	return &Card{Value: value, Suite: suite, Code: GetCode(suite, value)}
}

func NewCardFromCode(code string) (*Card, error) {
	suite, suiteError := GetSuite(string(code[len(code)-1]))

	if suiteError != nil {
		return nil, suiteError
	}

	value, valueError := GetValue(code[0 : len(code)-1])

	if valueError != nil {
		return nil, valueError
	}

	return &Card{Value: value, Suite: suite, Code: code}, nil
}

func IsValidCardValue(value string) bool {
	for _, v := range Values {
		if v == value {
			return true
		}
	}

	return false
}

func GetValue(code string) (string, error) {
	switch code {
	case "J":
		return "JACK", nil
	case "Q":
		return "QUEEN", nil
	case "K":
		return "KING", nil
	default:
		if !IsValidCardValue(code) {
			return "", errors.New("The card value is invalid: " + code)
		}
		return code, nil
	}
}

func GetSuite(code string) (string, error) {
	switch strings.ToUpper(code) {
	case "S":
		return "SPADES", nil
	case "D":
		return "DIAMONDS", nil
	case "C":
		return "CLUBS", nil
	case "H":
		return "HEARTS", nil
	default:
		return "", errors.New("The suite is invalid: " + code)
	}
}

func GetCode(suite string, value string) string {

	if IsSpecialValue(value) {
		return string(value[0]) + string(suite[0])
	}
	
	return value + string(suite[0])
}

func IsSpecialValue(value string) bool {
	if value == "JACK" || value == "QUEEN" || value == "KING" {
		return true
	}

	return false
}
