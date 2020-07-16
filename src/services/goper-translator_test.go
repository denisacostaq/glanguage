package services

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type dummyTranslator struct {
	suite.Suite
}

func (suite *dummyTranslator) SetupTest() {
}

func TestDummyTranslatorSuit(t *testing.T) {
	suite.Run(t, new(dummyTranslator))
}

func TranslatePathTests(suite *dummyTranslator, translator func(string) string, translations map[string]string) {
	// When
	res := make(map[string]string)
	for k := range translations {
		res[k] = translator(k)
	}

	// Then
	for k, v := range translations {
		suite.Equal(v, res[k])
	}
}

func (suite *dummyTranslator) TestTranslateWithVowelPrefix() {
	// Giving
	translations := map[string]string{
		"apple": "gapple",
		"orange": "gorange",
		"iceberg": "giceberg",
		"yellow": "yellow",
	}
	tr := NewDummyTranslator().(*DummyTranslator)

	// When
	TranslatePathTests(suite, tr.translateWithVowelPrefix, translations)
}

func (suite *dummyTranslator) TestTranslateWithXrPrefix() {
	// Giving
	translations := map[string]string{
		"xray": "gexray",
		"xreedfd": "gexreedfd",
		"xrxrdfd": "gexrxrdfd",
	}
	tr := NewDummyTranslator().(*DummyTranslator)

	// When
	TranslatePathTests(suite, tr.translateWithXrPrefix, translations)
}

func (suite *dummyTranslator) TestTranslateWithConsonantPrefix() {
	// Giving
	translations := map[string]string{
		"chair": "airchogo",
		"call": "allcogo",
		"phone": "onephogo",
		"thphone": "onethphogo",
	}
	tr := NewDummyTranslator().(*DummyTranslator)

	// When
	TranslatePathTests(suite, tr.translateWithConsonantPrefix, translations)
}

func (suite *dummyTranslator) TestTranslateWithConsonantPrefixFollowedBuQu() {
	// Giving
	translations := map[string]string{
		"square": "aresquogo",
		"trquare": "aretrquogo",
	}
	tr := NewDummyTranslator().(*DummyTranslator)

	// When
	TranslatePathTests(suite, tr.translateWithConsonantPrefixFollowedBuQu, translations)
}

func (suite *dummyTranslator) TestTranslateIgnoreApostrophe() {
	// Giving
	translations := map[string]string{
		"a'pple": "a'pple",
		"shouldn't": "shouldn't",
	}
	tr := NewDummyTranslator().(*DummyTranslator)

	// When
	res := make(map[string]string)
	for k := range translations {
		var err error
		res[k], err = tr.Translate2Gophers(k)
		suite.NoError(err)
	}

	// Then
	for k, v := range translations {
		suite.Equal(v, res[k])
	}

}