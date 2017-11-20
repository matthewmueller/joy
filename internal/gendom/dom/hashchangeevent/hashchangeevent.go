package hashchangeevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type HashChangeEvent struct {
	*event.Event
}

func (*HashChangeEvent) GetNewURL() (newURL string) {
	js.Rewrite("$<.newURL")
	return newURL
}

func (*HashChangeEvent) GetOldURL() (oldURL string) {
	js.Rewrite("$<.oldURL")
	return oldURL
}
