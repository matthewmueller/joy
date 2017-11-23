package avtrack

import (
	"github.com/matthewmueller/golly/dom2/appendmode"
	"github.com/matthewmueller/golly/dom2/msstream"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"SourceBuffer,omit"
type SourceBuffer struct {
	window.EventTarget
}

// Abort
func (*SourceBuffer) Abort() {
	js.Rewrite("$<.Abort()")
}

// AppendBuffer
func (*SourceBuffer) AppendBuffer(data interface{}) {
	js.Rewrite("$<.AppendBuffer($1)", data)
}

// AppendStream
func (*SourceBuffer) AppendStream(stream *msstream.MSStream, maxSize *uint64) {
	js.Rewrite("$<.AppendStream($1, $2)", stream, maxSize)
}

// Remove
func (*SourceBuffer) Remove(start float32, end float32) {
	js.Rewrite("$<.Remove($1, $2)", start, end)
}

// AppendWindowEnd
func (*SourceBuffer) AppendWindowEnd() (appendWindowEnd float32) {
	js.Rewrite("$<.AppendWindowEnd")
	return appendWindowEnd
}

// AppendWindowEnd
func (*SourceBuffer) SetAppendWindowEnd(appendWindowEnd float32) {
	js.Rewrite("$<.AppendWindowEnd = $1", appendWindowEnd)
}

// AppendWindowStart
func (*SourceBuffer) AppendWindowStart() (appendWindowStart float32) {
	js.Rewrite("$<.AppendWindowStart")
	return appendWindowStart
}

// AppendWindowStart
func (*SourceBuffer) SetAppendWindowStart(appendWindowStart float32) {
	js.Rewrite("$<.AppendWindowStart = $1", appendWindowStart)
}

// AudioTracks
func (*SourceBuffer) AudioTracks() (audioTracks *AudioTrackList) {
	js.Rewrite("$<.AudioTracks")
	return audioTracks
}

// Buffered
func (*SourceBuffer) Buffered() (buffered *timeranges.TimeRanges) {
	js.Rewrite("$<.Buffered")
	return buffered
}

// Mode
func (*SourceBuffer) Mode() (mode *appendmode.AppendMode) {
	js.Rewrite("$<.Mode")
	return mode
}

// Mode
func (*SourceBuffer) SetMode(mode *appendmode.AppendMode) {
	js.Rewrite("$<.Mode = $1", mode)
}

// TimestampOffset
func (*SourceBuffer) TimestampOffset() (timestampOffset float32) {
	js.Rewrite("$<.TimestampOffset")
	return timestampOffset
}

// TimestampOffset
func (*SourceBuffer) SetTimestampOffset(timestampOffset float32) {
	js.Rewrite("$<.TimestampOffset = $1", timestampOffset)
}

// Updating
func (*SourceBuffer) Updating() (updating bool) {
	js.Rewrite("$<.Updating")
	return updating
}

// VideoTracks
func (*SourceBuffer) VideoTracks() (videoTracks *VideoTrackList) {
	js.Rewrite("$<.VideoTracks")
	return videoTracks
}
