package mediasource

import (
	"github.com/matthewmueller/golly/dom2/avtrack"
	"github.com/matthewmueller/golly/dom2/sourcebufferlist"
	"github.com/matthewmueller/golly/dom2/window"
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

// AddSourceBuffer
func (*MediaSource) AddSourceBuffer(kind string) (a *avtrack.SourceBuffer) {
	js.Rewrite("$<.AddSourceBuffer($1)", kind)
	return a
}

// EndOfStream
func (*MediaSource) EndOfStream(err *string) {
	js.Rewrite("$<.EndOfStream($1)", err)
}

// IsTypeSupported
func (*MediaSource) IsTypeSupported(kind string) (b bool) {
	js.Rewrite("$<.IsTypeSupported($1)", kind)
	return b
}

// RemoveSourceBuffer
func (*MediaSource) RemoveSourceBuffer(sourceBuffer *avtrack.SourceBuffer) {
	js.Rewrite("$<.RemoveSourceBuffer($1)", sourceBuffer)
}

// ActiveSourceBuffers
func (*MediaSource) ActiveSourceBuffers() (activeSourceBuffers *sourcebufferlist.SourceBufferList) {
	js.Rewrite("$<.ActiveSourceBuffers")
	return activeSourceBuffers
}

// Duration
func (*MediaSource) Duration() (duration float32) {
	js.Rewrite("$<.Duration")
	return duration
}

// Duration
func (*MediaSource) SetDuration(duration float32) {
	js.Rewrite("$<.Duration = $1", duration)
}

// ReadyState
func (*MediaSource) ReadyState() (readyState string) {
	js.Rewrite("$<.ReadyState")
	return readyState
}

// SourceBuffers
func (*MediaSource) SourceBuffers() (sourceBuffers *sourcebufferlist.SourceBufferList) {
	js.Rewrite("$<.SourceBuffers")
	return sourceBuffers
}
