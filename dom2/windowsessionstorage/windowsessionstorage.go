package windowsessionstorage

import (
	"github.com/matthewmueller/golly/dom2/storage"
	"github.com/matthewmueller/golly/js"
)

// WindowSessionStorage struct
// js:"WindowSessionStorage,omit"
type WindowSessionStorage struct {
}

// SessionStorage
func (*WindowSessionStorage) SessionStorage() (sessionStorage *storage.Storage) {
	js.Rewrite("$<.SessionStorage")
	return sessionStorage
}
