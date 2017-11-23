package window

import "github.com/matthewmueller/golly/js"

// PageTransitionEvent struct
// js:"PageTransitionEvent,omit"
type PageTransitionEvent struct {
	Event
}

// Persisted prop
func (*PageTransitionEvent) Persisted() (persisted bool) {
	js.Rewrite("$<.persisted")
	return persisted
}
