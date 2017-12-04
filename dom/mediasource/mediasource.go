package mediasource

import (
	"github.com/matthewmueller/joy/dom/avtrack"
	"github.com/matthewmueller/joy/dom/sourcebufferlist"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.EventTarget = (*MediaSource)(nil)

// New fn
func New() *MediaSource {
	macro.Rewrite("new MediaSource()")
	return &MediaSource{}
}

// MediaSource struct
// js:"MediaSource,omit"
type MediaSource struct {
}

// AddSourceBuffer fn
// js:"addSourceBuffer"
func (*MediaSource) AddSourceBuffer(kind string) (a *avtrack.SourceBuffer) {
	macro.Rewrite("$_.addSourceBuffer($1)", kind)
	return a
}

// EndOfStream fn
// js:"endOfStream"
func (*MediaSource) EndOfStream(err *string) {
	macro.Rewrite("$_.endOfStream($1)", err)
}

// IsTypeSupported fn
// js:"isTypeSupported"
func (*MediaSource) IsTypeSupported(kind string) (b bool) {
	macro.Rewrite("$_.isTypeSupported($1)", kind)
	return b
}

// RemoveSourceBuffer fn
// js:"removeSourceBuffer"
func (*MediaSource) RemoveSourceBuffer(sourceBuffer *avtrack.SourceBuffer) {
	macro.Rewrite("$_.removeSourceBuffer($1)", sourceBuffer)
}

// AddEventListener fn
// js:"addEventListener"
func (*MediaSource) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*MediaSource) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*MediaSource) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// ActiveSourceBuffers prop
// js:"activeSourceBuffers"
func (*MediaSource) ActiveSourceBuffers() (activeSourceBuffers *sourcebufferlist.SourceBufferList) {
	macro.Rewrite("$_.activeSourceBuffers")
	return activeSourceBuffers
}

// Duration prop
// js:"duration"
func (*MediaSource) Duration() (duration float32) {
	macro.Rewrite("$_.duration")
	return duration
}

// SetDuration prop
// js:"duration"
func (*MediaSource) SetDuration(duration float32) {
	macro.Rewrite("$_.duration = $1", duration)
}

// ReadyState prop
// js:"readyState"
func (*MediaSource) ReadyState() (readyState string) {
	macro.Rewrite("$_.readyState")
	return readyState
}

// SourceBuffers prop
// js:"sourceBuffers"
func (*MediaSource) SourceBuffers() (sourceBuffers *sourcebufferlist.SourceBufferList) {
	macro.Rewrite("$_.sourceBuffers")
	return sourceBuffers
}
