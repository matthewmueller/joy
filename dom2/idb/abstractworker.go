package idb

import "github.com/matthewmueller/golly/js"

// AbstractWorker struct
// js:"AbstractWorker,omit"
type AbstractWorker struct {
}

// Onerror
func (*AbstractWorker) Onerror() (onerror func(Event)) {
	js.Rewrite("$<.Onerror")
	return onerror
}

// Onerror
func (*AbstractWorker) SetOnerror(onerror func(Event)) {
	js.Rewrite("$<.Onerror = $1", onerror)
}
