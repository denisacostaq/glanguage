package services

import (
	"github.com/denisacostaq/glanguage/src/models"
	"github.com/denisacostaq/glanguage/src/repository"
	log "github.com/sirupsen/logrus"
	"strings"
)

var dataSource = repository.NewInMemoryDataSource()

type DefaultTranslatorEngine struct {
	ds repository.DataSource
	translate func(word string) (string, error)
}

func NewDefaultTranslatorEngine(ds repository.DataSource, tr func(word string) (string, error)) (TranslatorEngine) {
	return &DefaultTranslatorEngine{ds: ds, translate: tr}
}

// CreateDefaultEngine create an engine that receive English words and get back
// the Gopher's language equivalent
func CreateDefaultEngine() TranslatorEngine {
	return NewDefaultTranslatorEngine(
		dataSource,
		NewGopherTranslator().Translate,
		)
}

func (de *DefaultTranslatorEngine) TranslateWord(word models.TranslationPair) error {
	tr, err := de.translate(word.English())
	if err != nil {
		log.WithFields(log.Fields{"word": word, "err": err}).Errorln("unable to translate word")
		return InvalidWordErr
	}
	word.SetTranslated(tr)
	if err := de.ds.Save(word); err != nil {
		log.WithFields(log.Fields{"word": word, "err": err}).Errorln("error saving word")
		return InternalErr
	}
	return nil
}

func removeLastSymbol(sentence string) string {
	if len(sentence) == 0 || len(sentence) == 1 {
		return ""
	}
	return sentence[:len(sentence)-1]
}

func validateSentence(sentence string) error {
	if len(sentence) == 0 {
		return EmptyValueErr
	}
	if !strings.Contains(sentence, " ") {
		return SentenceTooShortErr
	}
	lastSymbol := rune(sentence[len(sentence)-1])
	dotMark := rune(0x2E)
	questionMark := rune(0x3F)
	exclamationMark := rune(0x21)
	if lastSymbol != dotMark && lastSymbol != questionMark && lastSymbol != exclamationMark {
		return MalformedSentenceErr
	}
	return nil
}

func (de *DefaultTranslatorEngine) TranslateSentence(sentence models.TranslationPair) error {
	if err := validateSentence(sentence.English()); err != nil {
		log.WithError(err).Errorln(InvalidSentenceErr)
		return InvalidSentenceErr
	}
	lastSymbol := rune(sentence.English()[len(sentence.English())-1])
	words := strings.Split(removeLastSymbol(sentence.English()), " ")
	translatedWords := make([]string, len(words))
	ignored := 0
	for idx, word := range words {
		if strings.ContainsRune(word, apostropheUnicode) {
			ignored++
			continue
		}
		var err error
		translatedWords[idx-ignored], err = de.translate(word)
		if err != nil {
			log.WithFields(log.Fields{"word": word, "err": err}).Errorln("unable to translate word")
			log.WithFields(log.Fields{"sentence": sentence, "err": err}).Errorln("unable to translate sentence")
			return err
		}
	}
	sentence.SetTranslated(strings.Join(translatedWords[:len(translatedWords) - ignored], " ") + string(lastSymbol))
	if err := de.ds.Save(sentence); err != nil {
		log.WithFields(log.Fields{"sentence": sentence, "err": err}).Errorln("error saving sentence")
		return InternalErr
	}
	return nil
}

func (de *DefaultTranslatorEngine) History() ([]models.TranslationPair, error) {
	words, err := de.ds.List()
	if err != nil {
		log.WithError(err).Errorln("unable to get history from the data source")
		return nil, err
	}
	return words, nil
}
