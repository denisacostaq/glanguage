package services

import (
	log "github.com/sirupsen/logrus"
	"strings"
)

type DummyTranslator struct {
}

func NewDummyTranslator() Translator {
	return &DummyTranslator{}
}

// translatePath1 1. If a word starts with a vowel letter, add prefix “g” to
// the word (ex. apple => gapple)
func (tr *DummyTranslator) translatePath1(str string) string {
	var sb strings.Builder
	if isVowel(rune(str[0])) {
		sb.WriteRune('g')
		sb.WriteString(str)
		return sb.String()
	}
	return str
}

// translatePath2 2. If a word starts with the consonant letters “xr”, add the prefix
// “ge” to the begging of the word. Such words as “xray” actually sound in the beginning
// with vowel sound as you pronounce them so a true gopher would say “gexray”.
func (tr *DummyTranslator) translatePath2(str string) string {
	var sb strings.Builder
	if strings.HasPrefix(str, "xr") {
		sb.WriteString("ge")
		sb.WriteString(str)
		return sb.String()
	}
	return str
}

// translatePath3 3. If a word starts with a consonant sound, move it to the end
// of the word and then add “ogo” suffix to the word. Consonant sounds can be made
// up of multiple consonants, a.k.a. a consonant cluster (e.g. "chair" -> "airchogo”).
func (tr *DummyTranslator) translatePath3(str string) string {
	var sb strings.Builder
	if !isVowel(rune(str[0])) {
		var consonantPrefixLen int
		for consonantPrefixLen = 0; !isVowel(rune(str[consonantPrefixLen])); consonantPrefixLen++ {}
		sb.WriteString(str[consonantPrefixLen:])
		sb.WriteString(str[:consonantPrefixLen])
		sb.WriteString("ogo")
		return sb.String()
	}
	return str
}

// translatePath4 4. If a word starts with a consonant sound followed by "qu", move it
// to the end of the word, and then add "ogo" suffix to the word
// (e.g. "square" -> "aresquogo").
func (tr *DummyTranslator) translatePath4(str string) string {
	var sb strings.Builder
	if !isVowel(rune(str[0])) {
		var consonantPrefixLen int
		for consonantPrefixLen = 0; !isVowel(rune(str[consonantPrefixLen])) && str[consonantPrefixLen] != 'q'; consonantPrefixLen++ {}
		log.Errorln(str, isVowel(rune(str[0])), consonantPrefixLen, str[consonantPrefixLen:], strings.HasPrefix(str[consonantPrefixLen:], "qu"))
		if strings.HasPrefix(str[consonantPrefixLen:], "qu") {
			sb.WriteString(str[consonantPrefixLen+2:])
			sb.WriteString(str[:consonantPrefixLen+2])
			sb.WriteString("ogo")
			return sb.String()
		}
	}
	return str
}

func (tr *DummyTranslator) Translate2Gophers(word string) (string, error) {
	const apostropheUnicode = rune(0x27)
	log.WithField("word", word).Info("translating word...")
	if len(word) == 0 || strings.ContainsRune(word, apostropheUnicode) {
		return word, nil
	}
	translated := tr.translatePath1(word)
	if word == translated {
		translated = tr.translatePath2(word)
	}
	if word == translated {
		translated = tr.translatePath3(word)
	}
	if word == translated {
		translated = tr.translatePath4(word)
	}
	log.WithFields(log.Fields{"word": word, "translated": translated}).Debug("translating word...")
	return translated, nil
}
