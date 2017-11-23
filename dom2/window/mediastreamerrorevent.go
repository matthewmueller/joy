package window

import (
	"github.com/matthewmueller/golly/dom2/mediastreamerror"
	"github.com/matthewmueller/golly/js"
)

// js:"MediaStreamErrorEvent,omit"
type MediaStreamErrorEvent struct {
	Event
}

// Error
func (*MediaStreamErrorEvent) Error() (err *mediastreamerror.MediaStreamError) {
	js.Rewrite("$<.Error")
	return err
}
