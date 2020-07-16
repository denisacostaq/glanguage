package controllers

import (
	"encoding/json"
	"github.com/denisacostaq/glanguage/src/models"
	"github.com/denisacostaq/glanguage/src/services"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func TranslateWord(w http.ResponseWriter, r *http.Request) {
	word := models.Word{}
	if err := json.NewDecoder(r.Body).Decode(&word); err != nil {
		log.WithError(err).Errorln("unable to decode body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	eng := services.CreateDefaultEngine()
	if err := eng.TranslateWord(&word); err != nil {
		log.WithError(err).Errorln("unable to translate word")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(&models.Word{GopherWord: word.Translated()}); err != nil {
		log.WithError(err).Errorln("unable to encode response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func TranslateSentence(w http.ResponseWriter, r *http.Request) {
	sentence := models.Sentence{}
	if err := json.NewDecoder(r.Body).Decode(&sentence); err != nil {
		log.WithError(err).Errorln("unable to decode body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	eng := services.CreateDefaultEngine()
	if err := eng.TranslateSentence(&sentence); err != nil {
		log.WithError(err).Errorln("unable to translate sentence")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(&models.Sentence{GopherSentence: sentence.Translated()}); err != nil {
		log.WithError(err).Errorln("unable to encode response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func History(w http.ResponseWriter, r *http.Request) {
	eng := services.CreateDefaultEngine()
	history, err := eng.History()
	if err != nil {
		log.WithError(err).Errorln("unable to get history")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, err := services.SerializeFullHistoryAsJson(history)
	if err != nil {
		log.WithError(err).Errorln("unable to encode response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}