package window

import "github.com/matthewmueller/golly/js"

// PageTransitionEvent struct
// js:"PageTransitionEvent,omit"
type PageTransitionEvent struct {
	Event
}

// Persisted
func (*PageTransitionEvent) Persisted() (persisted bool) {
	js.Rewrite("$<.Persisted")
	return persisted
}
