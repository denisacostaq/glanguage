package models

type Word struct {
	EnglishWord string `json:"english-word,omitempty"`
	GopherWord string `json:"gopher-word,omitempty"`
}

func (w *Word) English() string {
	return w.EnglishWord
}

func (w *Word) SetEnglish(word string) {
	w.EnglishWord = word
}

func (w *Word) Translated() string {
	return w.GopherWord
}

func (w *Word) SetTranslated(word string) {
	w.GopherWord = word
}

type Sentence struct {
	EnglishSentence string `json:"english-sentence,omitempty"`
	GopherSentence string `json:"gopher-sentence,omitempty"`
}

func (s *Sentence) English() string {
	return s.EnglishSentence
}

func (s *Sentence) SetEnglish(word string) {
	s.EnglishSentence = word
}

func (s *Sentence) Translated() string {
	return s.GopherSentence
}

func (s *Sentence) SetTranslated(word string) {
	s.GopherSentence = word
}

type History struct {
	History []map[string]string `json:"history"`
}