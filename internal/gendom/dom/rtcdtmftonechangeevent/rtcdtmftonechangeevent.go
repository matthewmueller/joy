package rtcdtmftonechangeevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type RTCDTMFToneChangeEvent struct {
	*event.Event
}

func (*RTCDTMFToneChangeEvent) GetTone() (tone string) {
	js.Rewrite("$<.tone")
	return tone
}
