package windowlocalstorage

import (
	"github.com/matthewmueller/golly/dom2/storage"
	"github.com/matthewmueller/golly/js"
)

// WindowLocalStorage struct
// js:"WindowLocalStorage,omit"
type WindowLocalStorage struct {
}

// LocalStorage
func (*WindowLocalStorage) LocalStorage() (localStorage *storage.Storage) {
	js.Rewrite("$<.LocalStorage")
	return localStorage
}
