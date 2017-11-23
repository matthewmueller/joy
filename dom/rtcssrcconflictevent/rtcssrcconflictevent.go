package rtcssrcconflictevent

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// RTCSsrcConflictEvent struct
// js:"RTCSsrcConflictEvent,omit"
type RTCSsrcConflictEvent struct {
	window.Event
}

// Ssrc prop
func (*RTCSsrcConflictEvent) Ssrc() (ssrc uint) {
	js.Rewrite("$<.ssrc")
	return ssrc
}
