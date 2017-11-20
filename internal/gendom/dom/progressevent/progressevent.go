package progressevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type ProgressEvent struct {
	*event.Event
}

func (*ProgressEvent) InitProgressEvent(typeArg string, canBubbleArg bool, cancelableArg bool, lengthComputableArg bool, loadedArg uint64, totalArg uint64) {
	js.Rewrite("$<.initProgressEvent($1, $2, $3, $4, $5, $6)", typeArg, canBubbleArg, cancelableArg, lengthComputableArg, loadedArg, totalArg)
}

func (*ProgressEvent) GetLengthComputable() (lengthComputable bool) {
	js.Rewrite("$<.lengthComputable")
	return lengthComputable
}

func (*ProgressEvent) GetLoaded() (loaded uint64) {
	js.Rewrite("$<.loaded")
	return loaded
}

func (*ProgressEvent) GetTotal() (total uint64) {
	js.Rewrite("$<.total")
	return total
}
