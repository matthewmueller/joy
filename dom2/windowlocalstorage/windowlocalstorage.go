package windowlocalstorage

import "github.com/matthewmueller/golly/dom2/storage"

// js:"WindowLocalStorage,omit"
type WindowLocalStorage interface {

	// LocalStorage
	LocalStorage() (localStorage *storage.Storage)
}
