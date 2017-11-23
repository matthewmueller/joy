package msmediakeyneededevent

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// MSMediaKeyNeededEvent struct
// js:"MSMediaKeyNeededEvent,omit"
type MSMediaKeyNeededEvent struct {
	window.Event
}

// InitData prop
func (*MSMediaKeyNeededEvent) InitData() (initData []uint8) {
	js.Rewrite("$<.initData")
	return initData
}
