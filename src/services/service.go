package services

import (
	"github.com/denisacostaq/glanguage/src/models"
)

// Translator specify translation operations
type Translator interface {
	// Translate2Gophers specify the translation from English to the
	// Gophers's language specification
	Translate2Gophers(word string) (string, error)
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