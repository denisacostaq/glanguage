package models

// TranslationPair hold a pair of words/sentences. The english version
// an the equivalent in the gophers's language
type TranslationPair interface {
	// English get back the English version for the expression
	English() string

	// SetEnglish set the English version for the expression
	SetEnglish(string)

	// Translated get back the Translated version for the expression
	Translated() string

	// SetTranslated set the Translated version for the expression
	SetTranslated(string)
}
