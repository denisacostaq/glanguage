package repository

import (
	"container/list"
	"errors"
	"github.com/denisacostaq/glanguage/src/models"
	"sync"
)

var UnexpectedTypeErr = errors.New("unable to get object as expected type")

type InMemoryDataSource struct {
	sync.RWMutex
	history *list.List
}

func NewInMemoryDataSource() DataSource {
	return &InMemoryDataSource{history: list.New()}
}

func (ds *InMemoryDataSource) Save(word models.TranslationPair) error {
	ds.Lock()
	defer ds.Unlock()
	ds.history.PushBack(word)
	return nil
}

func (ds *InMemoryDataSource) List() ([]models.TranslationPair, error) {
	ds.RLock()
	defer ds.RUnlock()
	words := make([]models.TranslationPair, ds.history.Len())
	for e, idx := ds.history.Front(), 0; e != nil; e = e.Next() {
		val, ok := e.Value.(models.TranslationPair)
		if !ok {
			return words, UnexpectedTypeErr
		}
		words[idx] = val
		idx++
	}
	return words, nil
}
