package audionode

import (
	"github.com/matthewmueller/joy/dom/audiobuffer"
	"github.com/matthewmueller/joy/dom/audioparam"
	"github.com/matthewmueller/joy/dom/channelcountmode"
	"github.com/matthewmueller/joy/dom/channelinterpretation"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ AudioNode = (*AudioBufferSourceNode)(nil)
var _ window.EventTarget = (*AudioBufferSourceNode)(nil)

// AudioBufferSourceNode struct
// js:"AudioBufferSourceNode,omit"
type AudioBufferSourceNode struct {
}

// Start fn
// js:"start"
func (*AudioBufferSourceNode) Start(when *float32, offset *float32, duration *float32) {
	macro.Rewrite("$_.start($1, $2, $3)", when, offset, duration)
}

// Stop fn
// js:"stop"
func (*AudioBufferSourceNode) Stop(when *float32) {
	macro.Rewrite("$_.stop($1)", when)
}

// Connect fn
// js:"connect"
func (*AudioBufferSourceNode) Connect(destination AudioNode, output *uint, input *uint) (a AudioNode) {
	macro.Rewrite("$_.connect($1, $2, $3)", destination, output, input)
	return a
}

// Disconnect fn
// js:"disconnect"
func (*AudioBufferSourceNode) Disconnect() {
	macro.Rewrite("$_.disconnect()")
}

// AddEventListener fn
// js:"addEventListener"
func (*AudioBufferSourceNode) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*AudioBufferSourceNode) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*AudioBufferSourceNode) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Buffer prop
// js:"buffer"
func (*AudioBufferSourceNode) Buffer() (buffer *audiobuffer.AudioBuffer) {
	macro.Rewrite("$_.buffer")
	return buffer
}

// SetBuffer prop
// js:"buffer"
func (*AudioBufferSourceNode) SetBuffer(buffer *audiobuffer.AudioBuffer) {
	macro.Rewrite("$_.buffer = $1", buffer)
}

// Detune prop
// js:"detune"
func (*AudioBufferSourceNode) Detune() (detune *audioparam.AudioParam) {
	macro.Rewrite("$_.detune")
	return detune
}

// Loop prop
// js:"loop"
func (*AudioBufferSourceNode) Loop() (loop bool) {
	macro.Rewrite("$_.loop")
	return loop
}

// SetLoop prop
// js:"loop"
func (*AudioBufferSourceNode) SetLoop(loop bool) {
	macro.Rewrite("$_.loop = $1", loop)
}

// LoopEnd prop
// js:"loopEnd"
func (*AudioBufferSourceNode) LoopEnd() (loopEnd float32) {
	macro.Rewrite("$_.loopEnd")
	return loopEnd
}

// SetLoopEnd prop
// js:"loopEnd"
func (*AudioBufferSourceNode) SetLoopEnd(loopEnd float32) {
	macro.Rewrite("$_.loopEnd = $1", loopEnd)
}

// LoopStart prop
// js:"loopStart"
func (*AudioBufferSourceNode) LoopStart() (loopStart float32) {
	macro.Rewrite("$_.loopStart")
	return loopStart
}

// SetLoopStart prop
// js:"loopStart"
func (*AudioBufferSourceNode) SetLoopStart(loopStart float32) {
	macro.Rewrite("$_.loopStart = $1", loopStart)
}

// Onended prop
// js:"onended"
func (*AudioBufferSourceNode) Onended() (onended func(window.Event)) {
	macro.Rewrite("$_.onended")
	return onended
}

// SetOnended prop
// js:"onended"
func (*AudioBufferSourceNode) SetOnended(onended func(window.Event)) {
	macro.Rewrite("$_.onended = $1", onended)
}

// PlaybackRate prop
// js:"playbackRate"
func (*AudioBufferSourceNode) PlaybackRate() (playbackRate *audioparam.AudioParam) {
	macro.Rewrite("$_.playbackRate")
	return playbackRate
}

// ChannelCount prop
// js:"channelCount"
func (*AudioBufferSourceNode) ChannelCount() (channelCount uint) {
	macro.Rewrite("$_.channelCount")
	return channelCount
}

// SetChannelCount prop
// js:"channelCount"
func (*AudioBufferSourceNode) SetChannelCount(channelCount uint) {
	macro.Rewrite("$_.channelCount = $1", channelCount)
}

// ChannelCountMode prop
// js:"channelCountMode"
func (*AudioBufferSourceNode) ChannelCountMode() (channelCountMode *channelcountmode.ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode")
	return channelCountMode
}

// SetChannelCountMode prop
// js:"channelCountMode"
func (*AudioBufferSourceNode) SetChannelCountMode(channelCountMode *channelcountmode.ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode = $1", channelCountMode)
}

// ChannelInterpretation prop
// js:"channelInterpretation"
func (*AudioBufferSourceNode) ChannelInterpretation() (channelInterpretation *channelinterpretation.ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation")
	return channelInterpretation
}

// SetChannelInterpretation prop
// js:"channelInterpretation"
func (*AudioBufferSourceNode) SetChannelInterpretation(channelInterpretation *channelinterpretation.ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation = $1", channelInterpretation)
}

// Context prop
// js:"context"
func (*AudioBufferSourceNode) Context() (context AudioContext) {
	macro.Rewrite("$_.context")
	return context
}

// NumberOfInputs prop
// js:"numberOfInputs"
func (*AudioBufferSourceNode) NumberOfInputs() (numberOfInputs uint) {
	macro.Rewrite("$_.numberOfInputs")
	return numberOfInputs
}

// NumberOfOutputs prop
// js:"numberOfOutputs"
func (*AudioBufferSourceNode) NumberOfOutputs() (numberOfOutputs uint) {
	macro.Rewrite("$_.numberOfOutputs")
	return numberOfOutputs
}
