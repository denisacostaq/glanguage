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
	tr := NewGopherTranslator().(*GopherTranslator)

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
	tr := NewGopherTranslator().(*GopherTranslator)

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
		"my": "ymogo",
		"jjjj": "jjjjogo",
		"t": "togo",
		"year": "earyogo",
		"tyear": "yeartogo",
	}
	tr := NewGopherTranslator().(*GopherTranslator)

	// When
	TranslatePathTests(suite, tr.translateWithConsonantPrefix, translations)
}

func (suite *dummyTranslator) TestTranslateWithConsonantPrefixFollowedBuQu() {
	// Giving
	translations := map[string]string{
		"square": "aresquogo",
		"trquare": "aretrquogo",
		"qu": "qu",
		"rqu": "rquogo",
		"drquee": "eedrquogo",
		"equ": "equ",
		"qqu": "qquogo",
	}
	tr := NewGopherTranslator().(*GopherTranslator)

	// When
	TranslatePathTests(suite, tr.translateWithConsonantPrefixFollowedBuQu, translations)
}

func (suite *dummyTranslator) TestTranslateIgnoreApostrophe() {
	// Giving
	translations := map[string]error {
		"a'pple": InvalidWordErr,
		"shouldn't": InvalidWordErr,
	}
	tr := NewGopherTranslator().(*GopherTranslator)

	// When
	res := make(map[string]error)
	for k := range translations {
		var err error
		_, res[k] = tr.Translate(k)
		suite.NoError(err)
	}

	// Then
	for k, v := range translations {
		suite.Equal(v, res[k])
	}

}