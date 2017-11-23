package audionode

import (
	"github.com/matthewmueller/golly/dom/audiobuffer"
	"github.com/matthewmueller/golly/dom/audioparam"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// AudioBufferSourceNode struct
// js:"AudioBufferSourceNode,omit"
type AudioBufferSourceNode struct {
	AudioNode
}

// Start fn
func (*AudioBufferSourceNode) Start(when *float32, offset *float32, duration *float32) {
	js.Rewrite("$<.start($1, $2, $3)", when, offset, duration)
}

// Stop fn
func (*AudioBufferSourceNode) Stop(when *float32) {
	js.Rewrite("$<.stop($1)", when)
}

// Buffer prop
func (*AudioBufferSourceNode) Buffer() (buffer *audiobuffer.AudioBuffer) {
	js.Rewrite("$<.buffer")
	return buffer
}

// Buffer prop
func (*AudioBufferSourceNode) SetBuffer(buffer *audiobuffer.AudioBuffer) {
	js.Rewrite("$<.buffer = $1", buffer)
}

// Detune prop
func (*AudioBufferSourceNode) Detune() (detune *audioparam.AudioParam) {
	js.Rewrite("$<.detune")
	return detune
}

// Loop prop
func (*AudioBufferSourceNode) Loop() (loop bool) {
	js.Rewrite("$<.loop")
	return loop
}

// Loop prop
func (*AudioBufferSourceNode) SetLoop(loop bool) {
	js.Rewrite("$<.loop = $1", loop)
}

// LoopEnd prop
func (*AudioBufferSourceNode) LoopEnd() (loopEnd float32) {
	js.Rewrite("$<.loopEnd")
	return loopEnd
}

// LoopEnd prop
func (*AudioBufferSourceNode) SetLoopEnd(loopEnd float32) {
	js.Rewrite("$<.loopEnd = $1", loopEnd)
}

// LoopStart prop
func (*AudioBufferSourceNode) LoopStart() (loopStart float32) {
	js.Rewrite("$<.loopStart")
	return loopStart
}

// LoopStart prop
func (*AudioBufferSourceNode) SetLoopStart(loopStart float32) {
	js.Rewrite("$<.loopStart = $1", loopStart)
}

// Onended prop
func (*AudioBufferSourceNode) Onended() (onended func(window.Event)) {
	js.Rewrite("$<.onended")
	return onended
}

// Onended prop
func (*AudioBufferSourceNode) SetOnended(onended func(window.Event)) {
	js.Rewrite("$<.onended = $1", onended)
}

// PlaybackRate prop
func (*AudioBufferSourceNode) PlaybackRate() (playbackRate *audioparam.AudioParam) {
	js.Rewrite("$<.playbackRate")
	return playbackRate
}
