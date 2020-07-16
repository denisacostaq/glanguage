package services

import (
	"github.com/denisacostaq/glanguage/src/models"
	"errors"
)

var SentenceTooShortErr = errors.New("sentence should have more than one word")
var EmptyValueErr = errors.New("you should send a value")
var MalformedSentenceErr = errors.New("sentences should end with dot, question or exclamation mark")
var InvalidSentenceErr = errors.New("invalid sentence")
var InternalErr = errors.New("an internal error occurred")
var WordTooShortErr = errors.New("word should have at least a character")
var InvalidWordErr = errors.New("invalid word")

// Translator specify translation operations
type Translator interface {
	// Translate specify the translation from English to another language specification
	Translate(word string) (string, error)
}

// TranslatorEngine handle the translations operations to an specific
// language. In addition should handle/track the history translations
type TranslatorEngine interface {
	// TranslateWord translate a word from one language to another
	TranslateWord(word models.TranslationPair) error

	// TranslateSentence translate a sentence from one language to another
	TranslateSentence(sentence models.TranslationPair) error

	// History return all the previous translations requests
	History() ([]models.TranslationPair, error)
}