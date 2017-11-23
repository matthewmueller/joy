package mediasource

import (
	"github.com/matthewmueller/golly/dom/avtrack"
	"github.com/matthewmueller/golly/dom/sourcebufferlist"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New() *MediaSource {
	js.Rewrite("MediaSource")
	return &MediaSource{}
}

// MediaSource struct
// js:"MediaSource,omit"
type MediaSource struct {
	window.EventTarget
}

// AddSourceBuffer fn
func (*MediaSource) AddSourceBuffer(kind string) (a *avtrack.SourceBuffer) {
	js.Rewrite("$<.addSourceBuffer($1)", kind)
	return a
}

// EndOfStream fn
func (*MediaSource) EndOfStream(err *string) {
	js.Rewrite("$<.endOfStream($1)", err)
}

// IsTypeSupported fn
func (*MediaSource) IsTypeSupported(kind string) (b bool) {
	js.Rewrite("$<.isTypeSupported($1)", kind)
	return b
}

// RemoveSourceBuffer fn
func (*MediaSource) RemoveSourceBuffer(sourceBuffer *avtrack.SourceBuffer) {
	js.Rewrite("$<.removeSourceBuffer($1)", sourceBuffer)
}

// ActiveSourceBuffers prop
func (*MediaSource) ActiveSourceBuffers() (activeSourceBuffers *sourcebufferlist.SourceBufferList) {
	js.Rewrite("$<.activeSourceBuffers")
	return activeSourceBuffers
}

// Duration prop
func (*MediaSource) Duration() (duration float32) {
	js.Rewrite("$<.duration")
	return duration
}

// Duration prop
func (*MediaSource) SetDuration(duration float32) {
	js.Rewrite("$<.duration = $1", duration)
}

// ReadyState prop
func (*MediaSource) ReadyState() (readyState string) {
	js.Rewrite("$<.readyState")
	return readyState
}

// SourceBuffers prop
func (*MediaSource) SourceBuffers() (sourceBuffers *sourcebufferlist.SourceBufferList) {
	js.Rewrite("$<.sourceBuffers")
	return sourceBuffers
}
