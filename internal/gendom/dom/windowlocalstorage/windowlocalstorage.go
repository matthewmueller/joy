package windowlocalstorage

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/storage"
	"github.com/matthewmueller/golly/js"
)

type WindowLocalStorage struct {
}

func (*WindowLocalStorage) GetLocalStorage() (localStorage *storage.Storage) {
	js.Rewrite("$<.localStorage")
	return localStorage
}
