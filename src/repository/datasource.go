package repository

import (
	"errors"
	"github.com/denisacostaq/glanguage/src/models"
)

var DataSourceErr = errors.New("data source is not working")

// DataSource specify the operations for storing data
type DataSource interface {
	// Save a translation
	Save(word models.TranslationPair) error

	// List get back all the saved translations
	List() ([]models.TranslationPair, error)
}