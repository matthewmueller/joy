package windowsessionstorage

import "github.com/matthewmueller/joy/dom/storage"

// WindowSessionStorage interface
// js:"WindowSessionStorage"
type WindowSessionStorage interface {

	// SessionStorage prop
	// js:"sessionStorage"
	// jsrewrite:"$_.sessionStorage"
	SessionStorage() (sessionStorage *storage.Storage)
}
