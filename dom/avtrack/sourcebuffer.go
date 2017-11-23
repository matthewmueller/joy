package avtrack

import (
	"github.com/matthewmueller/golly/dom/appendmode"
	"github.com/matthewmueller/golly/dom/msstream"
	"github.com/matthewmueller/golly/dom/timeranges"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// SourceBuffer struct
// js:"SourceBuffer,omit"
type SourceBuffer struct {
	window.EventTarget
}

// Abort fn
func (*SourceBuffer) Abort() {
	js.Rewrite("$<.abort()")
}

// AppendBuffer fn
func (*SourceBuffer) AppendBuffer(data interface{}) {
	js.Rewrite("$<.appendBuffer($1)", data)
}

// AppendStream fn
func (*SourceBuffer) AppendStream(stream *msstream.MSStream, maxSize *uint64) {
	js.Rewrite("$<.appendStream($1, $2)", stream, maxSize)
}

// Remove fn
func (*SourceBuffer) Remove(start float32, end float32) {
	js.Rewrite("$<.remove($1, $2)", start, end)
}

// AppendWindowEnd prop
func (*SourceBuffer) AppendWindowEnd() (appendWindowEnd float32) {
	js.Rewrite("$<.appendWindowEnd")
	return appendWindowEnd
}

// AppendWindowEnd prop
func (*SourceBuffer) SetAppendWindowEnd(appendWindowEnd float32) {
	js.Rewrite("$<.appendWindowEnd = $1", appendWindowEnd)
}

// AppendWindowStart prop
func (*SourceBuffer) AppendWindowStart() (appendWindowStart float32) {
	js.Rewrite("$<.appendWindowStart")
	return appendWindowStart
}

// AppendWindowStart prop
func (*SourceBuffer) SetAppendWindowStart(appendWindowStart float32) {
	js.Rewrite("$<.appendWindowStart = $1", appendWindowStart)
}

// AudioTracks prop
func (*SourceBuffer) AudioTracks() (audioTracks *AudioTrackList) {
	js.Rewrite("$<.audioTracks")
	return audioTracks
}

// Buffered prop
func (*SourceBuffer) Buffered() (buffered *timeranges.TimeRanges) {
	js.Rewrite("$<.buffered")
	return buffered
}

// Mode prop
func (*SourceBuffer) Mode() (mode *appendmode.AppendMode) {
	js.Rewrite("$<.mode")
	return mode
}

// Mode prop
func (*SourceBuffer) SetMode(mode *appendmode.AppendMode) {
	js.Rewrite("$<.mode = $1", mode)
}

// TimestampOffset prop
func (*SourceBuffer) TimestampOffset() (timestampOffset float32) {
	js.Rewrite("$<.timestampOffset")
	return timestampOffset
}

// TimestampOffset prop
func (*SourceBuffer) SetTimestampOffset(timestampOffset float32) {
	js.Rewrite("$<.timestampOffset = $1", timestampOffset)
}

// Updating prop
func (*SourceBuffer) Updating() (updating bool) {
	js.Rewrite("$<.updating")
	return updating
}

// VideoTracks prop
func (*SourceBuffer) VideoTracks() (videoTracks *VideoTrackList) {
	js.Rewrite("$<.videoTracks")
	return videoTracks
}
