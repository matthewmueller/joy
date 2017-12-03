package audionode

import (
	"github.com/matthewmueller/joy/dom/audiobuffer"
	"github.com/matthewmueller/joy/dom/channelcountmode"
	"github.com/matthewmueller/joy/dom/channelinterpretation"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/js"
)

var _ AudioNode = (*ConvolverNode)(nil)
var _ window.EventTarget = (*ConvolverNode)(nil)

// ConvolverNode struct
// js:"ConvolverNode,omit"
type ConvolverNode struct {
}

// Connect fn
// js:"connect"
func (*ConvolverNode) Connect(destination AudioNode, output *uint, input *uint) (a AudioNode) {
	js.Rewrite("$_.connect($1, $2, $3)", destination, output, input)
	return a
}

// Disconnect fn
// js:"disconnect"
func (*ConvolverNode) Disconnect() {
	js.Rewrite("$_.disconnect()")
}

// AddEventListener fn
// js:"addEventListener"
func (*ConvolverNode) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*ConvolverNode) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*ConvolverNode) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Buffer prop
// js:"buffer"
func (*ConvolverNode) Buffer() (buffer *audiobuffer.AudioBuffer) {
	js.Rewrite("$_.buffer")
	return buffer
}

// SetBuffer prop
// js:"buffer"
func (*ConvolverNode) SetBuffer(buffer *audiobuffer.AudioBuffer) {
	js.Rewrite("$_.buffer = $1", buffer)
}

// Normalize prop
// js:"normalize"
func (*ConvolverNode) Normalize() (normalize bool) {
	js.Rewrite("$_.normalize")
	return normalize
}

// SetNormalize prop
// js:"normalize"
func (*ConvolverNode) SetNormalize(normalize bool) {
	js.Rewrite("$_.normalize = $1", normalize)
}

// ChannelCount prop
// js:"channelCount"
func (*ConvolverNode) ChannelCount() (channelCount uint) {
	js.Rewrite("$_.channelCount")
	return channelCount
}

// SetChannelCount prop
// js:"channelCount"
func (*ConvolverNode) SetChannelCount(channelCount uint) {
	js.Rewrite("$_.channelCount = $1", channelCount)
}

// ChannelCountMode prop
// js:"channelCountMode"
func (*ConvolverNode) ChannelCountMode() (channelCountMode *channelcountmode.ChannelCountMode) {
	js.Rewrite("$_.channelCountMode")
	return channelCountMode
}

// SetChannelCountMode prop
// js:"channelCountMode"
func (*ConvolverNode) SetChannelCountMode(channelCountMode *channelcountmode.ChannelCountMode) {
	js.Rewrite("$_.channelCountMode = $1", channelCountMode)
}

// ChannelInterpretation prop
// js:"channelInterpretation"
func (*ConvolverNode) ChannelInterpretation() (channelInterpretation *channelinterpretation.ChannelInterpretation) {
	js.Rewrite("$_.channelInterpretation")
	return channelInterpretation
}

// SetChannelInterpretation prop
// js:"channelInterpretation"
func (*ConvolverNode) SetChannelInterpretation(channelInterpretation *channelinterpretation.ChannelInterpretation) {
	js.Rewrite("$_.channelInterpretation = $1", channelInterpretation)
}

// Context prop
// js:"context"
func (*ConvolverNode) Context() (context AudioContext) {
	js.Rewrite("$_.context")
	return context
}

// NumberOfInputs prop
// js:"numberOfInputs"
func (*ConvolverNode) NumberOfInputs() (numberOfInputs uint) {
	js.Rewrite("$_.numberOfInputs")
	return numberOfInputs
}

// NumberOfOutputs prop
// js:"numberOfOutputs"
func (*ConvolverNode) NumberOfOutputs() (numberOfOutputs uint) {
	js.Rewrite("$_.numberOfOutputs")
	return numberOfOutputs
}
