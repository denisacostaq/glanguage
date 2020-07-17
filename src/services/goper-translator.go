package services

import (
	log "github.com/sirupsen/logrus"
	"strings"
)

// GopherTranslator	specify the translation from English to the
// Gophers's language specification
type GopherTranslator struct {
}

func NewGopherTranslator() Translator {
	return &GopherTranslator{}
}

// translateWithVowelPrefix 1. If a word starts with a vowel letter, add prefix “g” to
// the word (ex. apple => gapple)
func (tr *GopherTranslator) translateWithVowelPrefix(word string) string {
	var sb strings.Builder
	if isVowel(rune(word[0])) {
		sb.WriteRune('g')
		sb.WriteString(word)
		return sb.String()
	}
	return word
}

// translateWithXrPrefix 2. If a word starts with the consonant letters “xr”, add the prefix
// “ge” to the begging of the word. Such words as “xray” actually sound in the beginning
// with vowel sound as you pronounce them so a true gopher would say “gexray”.
func (tr *GopherTranslator) translateWithXrPrefix(word string) string {
	var sb strings.Builder
	if strings.HasPrefix(word, "xr") {
		sb.WriteString("ge")
		sb.WriteString(word)
		return sb.String()
	}
	return word
}

// translateWithConsonantPrefix 3. If a word starts with a consonant sound, move it to the end
// of the word and then add “ogo” suffix to the word. Consonant sounds can be made
// up of multiple consonants, a.k.a. a consonant cluster (e.g. "chair" -> "airchogo”).
func (tr *GopherTranslator) translateWithConsonantPrefix(word string) string {
	var sb strings.Builder
	if !isVowel(rune(word[0])) {
		const suffix = "ogo"
		var consonantPrefixLen int
		for consonantPrefixLen = 1; consonantPrefixLen < len(word) && !isLikeVowel(rune(word[consonantPrefixLen])); consonantPrefixLen++ {}
		if consonantPrefixLen == len(word) {
			return word + suffix
		}
		sb.WriteString(word[consonantPrefixLen:])
		sb.WriteString(word[:consonantPrefixLen])
		sb.WriteString(suffix)
		return sb.String()
	}
	return word
}

// translateWithConsonantPrefixFollowedBuQu 4. If a word starts with a consonant sound followed by "qu", move it
// to the end of the word, and then add "ogo" suffix to the word
// (e.g. "square" -> "aresquogo").
func (tr *GopherTranslator) translateWithConsonantPrefixFollowedBuQu(word string) string {
	var sb strings.Builder
	const qu = "qu"
	if !isVowel(rune(word[0])) && strings.Contains(word, qu) {
		var consonantPrefixLen int
		var quIndex = 0
		if strings.HasPrefix(word, qu) {
			if !strings.Contains(word[len(qu):], qu) {
				return word
			}
			quIndex = strings.Index(word[len(qu):], qu) + len(qu)
		} else {
			quIndex = strings.Index(word, qu)
		}
		for consonantPrefixLen = 1; !isLikeVowel(rune(word[consonantPrefixLen])) && (word[consonantPrefixLen] != 'q' || consonantPrefixLen < quIndex); consonantPrefixLen++ {}
		log.Warningln(consonantPrefixLen, word, strings.HasPrefix(word[consonantPrefixLen:], "qu"), word[consonantPrefixLen+len(qu):])
		log.Warningln(word[:consonantPrefixLen+len(qu)])
		if strings.HasPrefix(word[consonantPrefixLen:], "qu") {
			sb.WriteString(word[consonantPrefixLen+len(qu):])
			sb.WriteString(word[:consonantPrefixLen+len(qu)])
			sb.WriteString("ogo")
			return sb.String()
		}
	}
	return word
}

const apostropheUnicode = rune(0x27)
func (tr *GopherTranslator) Translate(word string) (string, error) {
	if len(word) == 0 {
		log.Errorln("empty word received")
		return "", WordTooShortErr
	}
	if strings.ContainsRune(word, apostropheUnicode) {
		log.WithField("word", word).Errorln("gophers don’t understand shortened versions of words or apostrophes")
		return "", InvalidWordErr
	}
	log.WithField("word", word).Info("translating word...")
	translated := tr.translateWithVowelPrefix(word)
	if word == translated {
		translated = tr.translateWithXrPrefix(word)
	}
	if word == translated {
		translated = tr.translateWithConsonantPrefixFollowedBuQu(word)
	}
	if word == translated {
		translated = tr.translateWithConsonantPrefix(word)
	}
	log.WithFields(log.Fields{"word": word, "translated": translated}).Debug("translating word...")
	return translated, nil
}
