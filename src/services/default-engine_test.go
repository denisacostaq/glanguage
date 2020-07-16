package services

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type defaultEngine struct {
	suite.Suite
}

func (suite *defaultEngine) SetupTest() {
}

func TestDefaultEngineSuit(t *testing.T) {
	suite.Run(t, new(defaultEngine))
}

func (suite *defaultEngine) TestRemoveLastSymbol() {
	// Giving
	translations := map[string]string{
		"": "",
		"a": "",
		"a.": "a",
		"fdfdf?": "fdfdf",
	}

	// When
	res := make(map[string]string)
	for k := range translations {
		res[k] = removeLastSymbol(k)
	}

	// Then
	for k, v := range translations {
		suite.Equal(v, res[k])
	}
}

func (suite *defaultEngine) TestValidateSentence() {
	// Giving
	translations := map[string]error{
		"": EmptyValueErr,
		"a": SentenceTooShortErr,
		"aa": SentenceTooShortErr,
		"a a": MalformedSentenceErr,
		" Yes": MalformedSentenceErr,
		"All Yes.": nil,
		"Yes, all?": nil,
		"Yes all!": nil,
	}

	// When
	res := make(map[string]error)
	for k := range translations {
		res[k] = validateSentence(k)
	}

	// Then
	for k, v := range translations {
		suite.Equal(v, res[k])
	}
}