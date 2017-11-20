package sourcebuffer

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/appendmode"
	"github.com/matthewmueller/golly/internal/gendom/dom/audiotracklist"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/msstream"
	"github.com/matthewmueller/golly/internal/gendom/dom/timeranges"
	"github.com/matthewmueller/golly/internal/gendom/dom/videotracklist"
	"github.com/matthewmueller/golly/js"
)

type SourceBuffer struct {
	*eventtarget.EventTarget
}

func (*SourceBuffer) Abort() {
	js.Rewrite("$<.abort()")
}

func (*SourceBuffer) AppendBuffer(data interface{}) {
	js.Rewrite("$<.appendBuffer($1)", data)
}

func (*SourceBuffer) AppendStream(stream *msstream.MSStream, maxSize uint64) {
	js.Rewrite("$<.appendStream($1, $2)", stream, maxSize)
}

func (*SourceBuffer) Remove(start float32, end float32) {
	js.Rewrite("$<.remove($1, $2)", start, end)
}

func (*SourceBuffer) GetAppendWindowEnd() (appendWindowEnd float32) {
	js.Rewrite("$<.appendWindowEnd")
	return appendWindowEnd
}

func (*SourceBuffer) SetAppendWindowEnd(appendWindowEnd float32) {
	js.Rewrite("$<.appendWindowEnd = $1", appendWindowEnd)
}

func (*SourceBuffer) GetAppendWindowStart() (appendWindowStart float32) {
	js.Rewrite("$<.appendWindowStart")
	return appendWindowStart
}

func (*SourceBuffer) SetAppendWindowStart(appendWindowStart float32) {
	js.Rewrite("$<.appendWindowStart = $1", appendWindowStart)
}

func (*SourceBuffer) GetAudioTracks() (audioTracks *audiotracklist.AudioTrackList) {
	js.Rewrite("$<.audioTracks")
	return audioTracks
}

func (*SourceBuffer) GetBuffered() (buffered *timeranges.TimeRanges) {
	js.Rewrite("$<.buffered")
	return buffered
}

func (*SourceBuffer) GetMode() (mode *appendmode.AppendMode) {
	js.Rewrite("$<.mode")
	return mode
}

func (*SourceBuffer) SetMode(mode *appendmode.AppendMode) {
	js.Rewrite("$<.mode = $1", mode)
}

func (*SourceBuffer) GetTimestampOffset() (timestampOffset float32) {
	js.Rewrite("$<.timestampOffset")
	return timestampOffset
}

func (*SourceBuffer) SetTimestampOffset(timestampOffset float32) {
	js.Rewrite("$<.timestampOffset = $1", timestampOffset)
}

func (*SourceBuffer) GetUpdating() (updating bool) {
	js.Rewrite("$<.updating")
	return updating
}

func (*SourceBuffer) GetVideoTracks() (videoTracks *videotracklist.VideoTrackList) {
	js.Rewrite("$<.videoTracks")
	return videoTracks
}
