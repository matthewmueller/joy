package windowlocalstorage

import "github.com/matthewmueller/golly/dom/storage"

// WindowLocalStorage interface
// js:"WindowLocalStorage"
type WindowLocalStorage interface {

	// LocalStorage prop
	// js:"localStorage"
	LocalStorage() (localStorage *storage.Storage)
}
