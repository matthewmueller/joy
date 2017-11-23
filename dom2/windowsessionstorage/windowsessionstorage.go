package windowsessionstorage

import "github.com/matthewmueller/golly/dom2/storage"

// js:"WindowSessionStorage,omit"
type WindowSessionStorage interface {

	// SessionStorage
	SessionStorage() (sessionStorage *storage.Storage)
}
