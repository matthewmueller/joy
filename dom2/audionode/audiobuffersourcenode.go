package audionode

import (
	"github.com/matthewmueller/golly/dom2/audiobuffer"
	"github.com/matthewmueller/golly/dom2/audioparam"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"AudioBufferSourceNode,omit"
type AudioBufferSourceNode struct {
	AudioNode
}

// Start
func (*AudioBufferSourceNode) Start(when *float32, offset *float32, duration *float32) {
	js.Rewrite("$<.Start($1, $2, $3)", when, offset, duration)
}

// Stop
func (*AudioBufferSourceNode) Stop(when *float32) {
	js.Rewrite("$<.Stop($1)", when)
}

// Buffer
func (*AudioBufferSourceNode) Buffer() (buffer *audiobuffer.AudioBuffer) {
	js.Rewrite("$<.Buffer")
	return buffer
}

// Buffer
func (*AudioBufferSourceNode) SetBuffer(buffer *audiobuffer.AudioBuffer) {
	js.Rewrite("$<.Buffer = $1", buffer)
}

// Detune
func (*AudioBufferSourceNode) Detune() (detune *audioparam.AudioParam) {
	js.Rewrite("$<.Detune")
	return detune
}

// Loop
func (*AudioBufferSourceNode) Loop() (loop bool) {
	js.Rewrite("$<.Loop")
	return loop
}

// Loop
func (*AudioBufferSourceNode) SetLoop(loop bool) {
	js.Rewrite("$<.Loop = $1", loop)
}

// LoopEnd
func (*AudioBufferSourceNode) LoopEnd() (loopEnd float32) {
	js.Rewrite("$<.LoopEnd")
	return loopEnd
}

// LoopEnd
func (*AudioBufferSourceNode) SetLoopEnd(loopEnd float32) {
	js.Rewrite("$<.LoopEnd = $1", loopEnd)
}

// LoopStart
func (*AudioBufferSourceNode) LoopStart() (loopStart float32) {
	js.Rewrite("$<.LoopStart")
	return loopStart
}

// LoopStart
func (*AudioBufferSourceNode) SetLoopStart(loopStart float32) {
	js.Rewrite("$<.LoopStart = $1", loopStart)
}

// Onended
func (*AudioBufferSourceNode) Onended() (onended func(window.Event)) {
	js.Rewrite("$<.Onended")
	return onended
}

// Onended
func (*AudioBufferSourceNode) SetOnended(onended func(window.Event)) {
	js.Rewrite("$<.Onended = $1", onended)
}

// PlaybackRate
func (*AudioBufferSourceNode) PlaybackRate() (playbackRate *audioparam.AudioParam) {
	js.Rewrite("$<.PlaybackRate")
	return playbackRate
}
