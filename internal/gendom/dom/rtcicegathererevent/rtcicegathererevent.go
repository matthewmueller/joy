package rtcicegathererevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type RTCIceGathererEvent struct {
	*event.Event
}

func (*RTCIceGathererEvent) GetCandidate() (candidate interface{}) {
	js.Rewrite("$<.candidate")
	return candidate
}
