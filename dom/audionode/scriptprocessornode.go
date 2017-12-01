package audionode

import (
	"github.com/matthewmueller/golly/dom/audioprocessingevent"
	"github.com/matthewmueller/golly/dom/channelcountmode"
	"github.com/matthewmueller/golly/dom/channelinterpretation"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

var _ AudioNode = (*ScriptProcessorNode)(nil)
var _ window.EventTarget = (*ScriptProcessorNode)(nil)

// ScriptProcessorNode struct
// js:"ScriptProcessorNode,omit"
type ScriptProcessorNode struct {
}

// Connect fn
// js:"connect"
func (*ScriptProcessorNode) Connect(destination AudioNode, output *uint, input *uint) (a AudioNode) {
	js.Rewrite("$_.connect($1, $2, $3)", destination, output, input)
	return a
}

// Disconnect fn
// js:"disconnect"
func (*ScriptProcessorNode) Disconnect() {
	js.Rewrite("$_.disconnect()")
}

// AddEventListener fn
// js:"addEventListener"
func (*ScriptProcessorNode) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*ScriptProcessorNode) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*ScriptProcessorNode) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// BufferSize prop
// js:"bufferSize"
func (*ScriptProcessorNode) BufferSize() (bufferSize int) {
	js.Rewrite("$_.bufferSize")
	return bufferSize
}

// Onaudioprocess prop
// js:"onaudioprocess"
func (*ScriptProcessorNode) Onaudioprocess() (onaudioprocess func(*audioprocessingevent.AudioProcessingEvent)) {
	js.Rewrite("$_.onaudioprocess")
	return onaudioprocess
}

// SetOnaudioprocess prop
// js:"onaudioprocess"
func (*ScriptProcessorNode) SetOnaudioprocess(onaudioprocess func(*audioprocessingevent.AudioProcessingEvent)) {
	js.Rewrite("$_.onaudioprocess = $1", onaudioprocess)
}

// ChannelCount prop
// js:"channelCount"
func (*ScriptProcessorNode) ChannelCount() (channelCount uint) {
	js.Rewrite("$_.channelCount")
	return channelCount
}

// SetChannelCount prop
// js:"channelCount"
func (*ScriptProcessorNode) SetChannelCount(channelCount uint) {
	js.Rewrite("$_.channelCount = $1", channelCount)
}

// ChannelCountMode prop
// js:"channelCountMode"
func (*ScriptProcessorNode) ChannelCountMode() (channelCountMode *channelcountmode.ChannelCountMode) {
	js.Rewrite("$_.channelCountMode")
	return channelCountMode
}

// SetChannelCountMode prop
// js:"channelCountMode"
func (*ScriptProcessorNode) SetChannelCountMode(channelCountMode *channelcountmode.ChannelCountMode) {
	js.Rewrite("$_.channelCountMode = $1", channelCountMode)
}

// ChannelInterpretation prop
// js:"channelInterpretation"
func (*ScriptProcessorNode) ChannelInterpretation() (channelInterpretation *channelinterpretation.ChannelInterpretation) {
	js.Rewrite("$_.channelInterpretation")
	return channelInterpretation
}

// SetChannelInterpretation prop
// js:"channelInterpretation"
func (*ScriptProcessorNode) SetChannelInterpretation(channelInterpretation *channelinterpretation.ChannelInterpretation) {
	js.Rewrite("$_.channelInterpretation = $1", channelInterpretation)
}

// Context prop
// js:"context"
func (*ScriptProcessorNode) Context() (context AudioContext) {
	js.Rewrite("$_.context")
	return context
}

// NumberOfInputs prop
// js:"numberOfInputs"
func (*ScriptProcessorNode) NumberOfInputs() (numberOfInputs uint) {
	js.Rewrite("$_.numberOfInputs")
	return numberOfInputs
}

// NumberOfOutputs prop
// js:"numberOfOutputs"
func (*ScriptProcessorNode) NumberOfOutputs() (numberOfOutputs uint) {
	js.Rewrite("$_.numberOfOutputs")
	return numberOfOutputs
}
