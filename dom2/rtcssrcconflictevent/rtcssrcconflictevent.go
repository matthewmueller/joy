package rtcssrcconflictevent

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// RTCSsrcConflictEvent struct
// js:"RTCSsrcConflictEvent,omit"
type RTCSsrcConflictEvent struct {
	window.Event
}

// Ssrc
func (*RTCSsrcConflictEvent) Ssrc() (ssrc uint) {
	js.Rewrite("$<.Ssrc")
	return ssrc
}
