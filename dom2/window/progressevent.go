package window

import "github.com/matthewmueller/golly/js"

// js:"ProgressEvent,omit"
type ProgressEvent struct {
	Event
}

// InitProgressEvent
func (*ProgressEvent) InitProgressEvent(typeArg string, canBubbleArg bool, cancelableArg bool, lengthComputableArg bool, loadedArg uint64, totalArg uint64) {
	js.Rewrite("$<.InitProgressEvent($1, $2, $3, $4, $5, $6)", typeArg, canBubbleArg, cancelableArg, lengthComputableArg, loadedArg, totalArg)
}

// LengthComputable
func (*ProgressEvent) LengthComputable() (lengthComputable bool) {
	js.Rewrite("$<.LengthComputable")
	return lengthComputable
}

// Loaded
func (*ProgressEvent) Loaded() (loaded uint64) {
	js.Rewrite("$<.Loaded")
	return loaded
}

// Total
func (*ProgressEvent) Total() (total uint64) {
	js.Rewrite("$<.Total")
	return total
}
