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
		NewDummyTranslator().Translate2Gophers,
		)
}

func (de *DefaultTranslatorEngine) TranslateWord(word models.TranslationPair) error {
	tr, err := de.translate(word.English())
	if err != nil {
		log.WithFields(log.Fields{"word": word, "err": err}).Errorln("unable to translate word")
		return err
	}
	word.SetTranslated(tr)
	if err := de.ds.Save(word); err != nil {
		log.WithFields(log.Fields{"word": word, "err": err}).Errorln("error saving word")
		return err
	}
	return nil
}

func (de *DefaultTranslatorEngine) TranslateSentence(sentence models.TranslationPair) error {
	words := strings.Split(sentence.English(), " ")
	translatedWords := make([]string, len(words))
	for idx, word := range words {
		var err error
		translatedWords[idx], err = de.translate(word)
		if err != nil {
			log.WithFields(log.Fields{"word": word, "err": err}).Errorln("unable to translate word")
			log.WithFields(log.Fields{"sentence": sentence, "err": err}).Errorln("unable to translate sentence")
			return err
		}
	}
	sentence.SetTranslated(strings.Join(translatedWords, " "))
	if err := de.ds.Save(sentence); err != nil {
		log.WithFields(log.Fields{"sentence": sentence, "err": err}).Errorln("error saving sentence")
		return err
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
