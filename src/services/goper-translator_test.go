package services

import (
	"github.com/stretchrcom/testify/suite"
	"testing"
)

/*	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/mock"*/
	//"github.com/stretchr/testify/suite"

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

func (suite *dummyTranslator) TestTranslatePath1() {
	// Giving
	translations := map[string]string{
		"apple": "gapple",
		"orange": "gorange",
		"iceberg": "giceberg",
		"yellow": "yellow",
	}
	tr := NewDummyTranslator().(*DummyTranslator)

	// When
	TranslatePathTests(suite, tr.translatePath1, translations)
}

func (suite *dummyTranslator) TestTranslatePath2() {
	// Giving
	translations := map[string]string{
		"xray": "gexray",
		"xreedfd": "gexreedfd",
		"xrxrdfd": "gexrxrdfd",
	}
	tr := NewDummyTranslator().(*DummyTranslator)

	// When
	TranslatePathTests(suite, tr.translatePath2, translations)
}

func (suite *dummyTranslator) TestTranslatePath3() {
	// Giving
	translations := map[string]string{
		"chair": "airchogo",
		"call": "allcogo",
		"phone": "onephogo",
		"thphone": "onethphogo",
	}
	tr := NewDummyTranslator().(*DummyTranslator)

	// When
	TranslatePathTests(suite, tr.translatePath3, translations)
}

func (suite *dummyTranslator) TestTranslatePath4() {
	// Giving
	translations := map[string]string{
		"square": "aresquogo",
		"trquare": "aretrquogo",
	}
	tr := NewDummyTranslator().(*DummyTranslator)

	// When
	TranslatePathTests(suite, tr.translatePath4, translations)
}