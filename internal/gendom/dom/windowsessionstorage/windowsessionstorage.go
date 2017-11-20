package windowsessionstorage

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/storage"
	"github.com/matthewmueller/golly/js"
)

type WindowSessionStorage struct {
}

func (*WindowSessionStorage) GetSessionStorage() (sessionStorage *storage.Storage) {
	js.Rewrite("$<.sessionStorage")
	return sessionStorage
}
