package clipboardevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/datatransfer"
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type ClipboardEvent struct {
	*event.Event
}

func (*ClipboardEvent) GetClipboardData() (clipboardData *datatransfer.DataTransfer) {
	js.Rewrite("$<.clipboardData")
	return clipboardData
}
