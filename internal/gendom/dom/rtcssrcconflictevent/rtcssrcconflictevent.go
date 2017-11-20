package rtcssrcconflictevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type RTCSsrcConflictEvent struct {
	*event.Event
}

func (*RTCSsrcConflictEvent) GetSsrc() (ssrc uint) {
	js.Rewrite("$<.ssrc")
	return ssrc
}
