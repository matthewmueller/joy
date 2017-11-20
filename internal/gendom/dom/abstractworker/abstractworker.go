package abstractworker

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type AbstractWorker struct {
}

func (*AbstractWorker) GetOnerror() (e *event.Event) {
	js.Rewrite("$<.onerror")
	return e
}

func (*AbstractWorker) SetOnerror(e *event.Event) {
	js.Rewrite("$<.onerror = $1", e)
}
