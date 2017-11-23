package window

import (
	"github.com/matthewmueller/golly/dom/datatransfer"
	"github.com/matthewmueller/golly/js"
)

// ClipboardEvent struct
// js:"ClipboardEvent,omit"
type ClipboardEvent struct {
	Event
}

// ClipboardData prop
func (*ClipboardEvent) ClipboardData() (clipboardData *datatransfer.DataTransfer) {
	js.Rewrite("$<.clipboardData")
	return clipboardData
}
