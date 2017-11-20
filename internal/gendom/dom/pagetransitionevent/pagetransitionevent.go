package pagetransitionevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type PageTransitionEvent struct {
	*event.Event
}

func (*PageTransitionEvent) GetPersisted() (persisted bool) {
	js.Rewrite("$<.persisted")
	return persisted
}
