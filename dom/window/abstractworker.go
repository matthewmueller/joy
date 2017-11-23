package window

import "github.com/matthewmueller/golly/js"

// AbstractWorker struct
// js:"AbstractWorker,omit"
type AbstractWorker struct {
}

// Onerror prop
func (*AbstractWorker) Onerror() (onerror func(Event)) {
	js.Rewrite("$<.onerror")
	return onerror
}

// Onerror prop
func (*AbstractWorker) SetOnerror(onerror func(Event)) {
	js.Rewrite("$<.onerror = $1", onerror)
}
