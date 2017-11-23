package windowsessionstorage

import (
	"github.com/matthewmueller/golly/dom/storage"
	"github.com/matthewmueller/golly/js"
)

// WindowSessionStorage struct
// js:"WindowSessionStorage,omit"
type WindowSessionStorage struct {
}

// SessionStorage prop
func (*WindowSessionStorage) SessionStorage() (sessionStorage *storage.Storage) {
	js.Rewrite("$<.sessionStorage")
	return sessionStorage
}
