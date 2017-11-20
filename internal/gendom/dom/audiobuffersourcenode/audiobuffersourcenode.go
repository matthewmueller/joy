package audiobuffersourcenode

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/audiobuffer"
	"github.com/matthewmueller/golly/internal/gendom/dom/audionode"
	"github.com/matthewmueller/golly/internal/gendom/dom/audioparam"
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type AudioBufferSourceNode struct {
	*audionode.AudioNode
}

func (*AudioBufferSourceNode) Start(when float32, offset float32, duration float32) {
	js.Rewrite("$<.start($1, $2, $3)", when, offset, duration)
}

func (*AudioBufferSourceNode) Stop(when float32) {
	js.Rewrite("$<.stop($1)", when)
}

func (*AudioBufferSourceNode) GetBuffer() (buffer *audiobuffer.AudioBuffer) {
	js.Rewrite("$<.buffer")
	return buffer
}

func (*AudioBufferSourceNode) SetBuffer(buffer *audiobuffer.AudioBuffer) {
	js.Rewrite("$<.buffer = $1", buffer)
}

func (*AudioBufferSourceNode) GetDetune() (detune *audioparam.AudioParam) {
	js.Rewrite("$<.detune")
	return detune
}

func (*AudioBufferSourceNode) GetLoop() (loop bool) {
	js.Rewrite("$<.loop")
	return loop
}

func (*AudioBufferSourceNode) SetLoop(loop bool) {
	js.Rewrite("$<.loop = $1", loop)
}

func (*AudioBufferSourceNode) GetLoopEnd() (loopEnd float32) {
	js.Rewrite("$<.loopEnd")
	return loopEnd
}

func (*AudioBufferSourceNode) SetLoopEnd(loopEnd float32) {
	js.Rewrite("$<.loopEnd = $1", loopEnd)
}

func (*AudioBufferSourceNode) GetLoopStart() (loopStart float32) {
	js.Rewrite("$<.loopStart")
	return loopStart
}

func (*AudioBufferSourceNode) SetLoopStart(loopStart float32) {
	js.Rewrite("$<.loopStart = $1", loopStart)
}

func (*AudioBufferSourceNode) GetOnended() (e *event.Event) {
	js.Rewrite("$<.onended")
	return e
}

func (*AudioBufferSourceNode) SetOnended(e *event.Event) {
	js.Rewrite("$<.onended = $1", e)
}

func (*AudioBufferSourceNode) GetPlaybackRate() (playbackRate *audioparam.AudioParam) {
	js.Rewrite("$<.playbackRate")
	return playbackRate
}
