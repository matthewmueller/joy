package window

import (
	"github.com/matthewmueller/golly/dom2/progresseventinit"
	"github.com/matthewmueller/golly/js"
)

// NewProgressEvent fn
func NewProgressEvent(typearg string, eventinitdict *progresseventinit.ProgressEventInit) *ProgressEvent {
	js.Rewrite("ProgressEvent")
	return &ProgressEvent{}
}

// ProgressEvent struct
// js:"ProgressEvent,omit"
type ProgressEvent struct {
	Event
}

// InitProgressEvent fn
func (*ProgressEvent) InitProgressEvent(typeArg string, canBubbleArg bool, cancelableArg bool, lengthComputableArg bool, loadedArg uint64, totalArg uint64) {
	js.Rewrite("$<.initProgressEvent($1, $2, $3, $4, $5, $6)", typeArg, canBubbleArg, cancelableArg, lengthComputableArg, loadedArg, totalArg)
}

// LengthComputable prop
func (*ProgressEvent) LengthComputable() (lengthComputable bool) {
	js.Rewrite("$<.lengthComputable")
	return lengthComputable
}

// Loaded prop
func (*ProgressEvent) Loaded() (loaded uint64) {
	js.Rewrite("$<.loaded")
	return loaded
}

// Total prop
func (*ProgressEvent) Total() (total uint64) {
	js.Rewrite("$<.total")
	return total
}
