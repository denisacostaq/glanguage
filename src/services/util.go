package services

import (
	"encoding/json"
	"errors"
	log "github.com/Sirupsen/logrus"
	"github.com/denisacostaq/glanguage/src/models"
	"strings"
)

var SerializationErr = errors.New("serialization error")

const vowels = "aeiou"

func isVowel(r rune) bool {
	return strings.ContainsRune(vowels, r)
}

func SerializeFullHistoryAsJson(history []models.TranslationPair) ([]byte, error) {
	respHistory := models.History{History: make([]map[string]string, len(history))}
	for idx := range history {
		m := make(map[string]string)
		m[history[idx].English()] = history[idx].Translated()
		respHistory.History[idx] = m
	}
	ser, err := json.Marshal(respHistory)
	if err != nil {
		log.WithError(err).Errorln("unable to serialize history")
		return nil, SerializationErr
	}
	return ser, nil
}
