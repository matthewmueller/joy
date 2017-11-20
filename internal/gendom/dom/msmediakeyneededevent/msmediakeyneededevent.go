package msmediakeyneededevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type MSMediaKeyNeededEvent struct {
	*event.Event
}

func (*MSMediaKeyNeededEvent) GetInitData() (initData []uint8) {
	js.Rewrite("$<.initData")
	return initData
}
