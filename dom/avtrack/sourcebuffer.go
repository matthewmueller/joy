package avtrack

import (
	"github.com/matthewmueller/joy/dom/appendmode"
	"github.com/matthewmueller/joy/dom/msstream"
	"github.com/matthewmueller/joy/dom/timeranges"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/js"
)

var _ window.EventTarget = (*SourceBuffer)(nil)

// SourceBuffer struct
// js:"SourceBuffer,omit"
type SourceBuffer struct {
}

// Abort fn
// js:"abort"
func (*SourceBuffer) Abort() {
	js.Rewrite("$_.abort()")
}

// AppendBuffer fn
// js:"appendBuffer"
func (*SourceBuffer) AppendBuffer(data interface{}) {
	js.Rewrite("$_.appendBuffer($1)", data)
}

// AppendStream fn
// js:"appendStream"
func (*SourceBuffer) AppendStream(stream *msstream.MSStream, maxSize *uint64) {
	js.Rewrite("$_.appendStream($1, $2)", stream, maxSize)
}

// Remove fn
// js:"remove"
func (*SourceBuffer) Remove(start float32, end float32) {
	js.Rewrite("$_.remove($1, $2)", start, end)
}

// AddEventListener fn
// js:"addEventListener"
func (*SourceBuffer) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*SourceBuffer) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*SourceBuffer) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// AppendWindowEnd prop
// js:"appendWindowEnd"
func (*SourceBuffer) AppendWindowEnd() (appendWindowEnd float32) {
	js.Rewrite("$_.appendWindowEnd")
	return appendWindowEnd
}

// SetAppendWindowEnd prop
// js:"appendWindowEnd"
func (*SourceBuffer) SetAppendWindowEnd(appendWindowEnd float32) {
	js.Rewrite("$_.appendWindowEnd = $1", appendWindowEnd)
}

// AppendWindowStart prop
// js:"appendWindowStart"
func (*SourceBuffer) AppendWindowStart() (appendWindowStart float32) {
	js.Rewrite("$_.appendWindowStart")
	return appendWindowStart
}

// SetAppendWindowStart prop
// js:"appendWindowStart"
func (*SourceBuffer) SetAppendWindowStart(appendWindowStart float32) {
	js.Rewrite("$_.appendWindowStart = $1", appendWindowStart)
}

// AudioTracks prop
// js:"audioTracks"
func (*SourceBuffer) AudioTracks() (audioTracks *AudioTrackList) {
	js.Rewrite("$_.audioTracks")
	return audioTracks
}

// Buffered prop
// js:"buffered"
func (*SourceBuffer) Buffered() (buffered *timeranges.TimeRanges) {
	js.Rewrite("$_.buffered")
	return buffered
}

// Mode prop
// js:"mode"
func (*SourceBuffer) Mode() (mode *appendmode.AppendMode) {
	js.Rewrite("$_.mode")
	return mode
}

// SetMode prop
// js:"mode"
func (*SourceBuffer) SetMode(mode *appendmode.AppendMode) {
	js.Rewrite("$_.mode = $1", mode)
}

// TimestampOffset prop
// js:"timestampOffset"
func (*SourceBuffer) TimestampOffset() (timestampOffset float32) {
	js.Rewrite("$_.timestampOffset")
	return timestampOffset
}

// SetTimestampOffset prop
// js:"timestampOffset"
func (*SourceBuffer) SetTimestampOffset(timestampOffset float32) {
	js.Rewrite("$_.timestampOffset = $1", timestampOffset)
}

// Updating prop
// js:"updating"
func (*SourceBuffer) Updating() (updating bool) {
	js.Rewrite("$_.updating")
	return updating
}

// VideoTracks prop
// js:"videoTracks"
func (*SourceBuffer) VideoTracks() (videoTracks *VideoTrackList) {
	js.Rewrite("$_.videoTracks")
	return videoTracks
}
