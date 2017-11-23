package window

import (
	"github.com/matthewmueller/golly/dom2/datatransfer"
	"github.com/matthewmueller/golly/js"
)

// js:"ClipboardEvent,omit"
type ClipboardEvent struct {
	Event
}

// ClipboardData
func (*ClipboardEvent) ClipboardData() (clipboardData *datatransfer.DataTransfer) {
	js.Rewrite("$<.ClipboardData")
	return clipboardData
}
