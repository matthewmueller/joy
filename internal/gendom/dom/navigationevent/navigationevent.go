package navigationevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type NavigationEvent struct {
	*event.Event
}

func (*NavigationEvent) GetURI() (uri string) {
	js.Rewrite("$<.uri")
	return uri
}
