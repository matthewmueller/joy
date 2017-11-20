package mediasource

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/sourcebuffer"
	"github.com/matthewmueller/golly/internal/gendom/dom/sourcebufferlist"
	"github.com/matthewmueller/golly/js"
)

type MediaSource struct {
	*eventtarget.EventTarget
}

func (*MediaSource) AddSourceBuffer(kind string) (s *sourcebuffer.SourceBuffer) {
	js.Rewrite("$<.addSourceBuffer($1)", kind)
	return s
}

func (*MediaSource) EndOfStream(err string) {
	js.Rewrite("$<.endOfStream($1)", err)
}

func (*MediaSource) IsTypeSupported(kind string) (b bool) {
	js.Rewrite("$<.isTypeSupported($1)", kind)
	return b
}

func (*MediaSource) RemoveSourceBuffer(sourceBuffer *sourcebuffer.SourceBuffer) {
	js.Rewrite("$<.removeSourceBuffer($1)", sourceBuffer)
}

func (*MediaSource) GetActiveSourceBuffers() (activeSourceBuffers *sourcebufferlist.SourceBufferList) {
	js.Rewrite("$<.activeSourceBuffers")
	return activeSourceBuffers
}

func (*MediaSource) GetDuration() (duration float32) {
	js.Rewrite("$<.duration")
	return duration
}

func (*MediaSource) SetDuration(duration float32) {
	js.Rewrite("$<.duration = $1", duration)
}

func (*MediaSource) GetReadyState() (readyState string) {
	js.Rewrite("$<.readyState")
	return readyState
}

func (*MediaSource) GetSourceBuffers() (sourceBuffers *sourcebufferlist.SourceBufferList) {
	js.Rewrite("$<.sourceBuffers")
	return sourceBuffers
}
