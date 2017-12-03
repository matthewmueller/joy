package audionode

import (
	"github.com/matthewmueller/joy/dom/audioprocessingevent"
	"github.com/matthewmueller/joy/dom/channelcountmode"
	"github.com/matthewmueller/joy/dom/channelinterpretation"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
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
	macro.Rewrite("$_.connect($1, $2, $3)", destination, output, input)
	return a
}

// Disconnect fn
// js:"disconnect"
func (*ScriptProcessorNode) Disconnect() {
	macro.Rewrite("$_.disconnect()")
}

// AddEventListener fn
// js:"addEventListener"
func (*ScriptProcessorNode) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*ScriptProcessorNode) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*ScriptProcessorNode) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// BufferSize prop
// js:"bufferSize"
func (*ScriptProcessorNode) BufferSize() (bufferSize int) {
	macro.Rewrite("$_.bufferSize")
	return bufferSize
}

// Onaudioprocess prop
// js:"onaudioprocess"
func (*ScriptProcessorNode) Onaudioprocess() (onaudioprocess func(*audioprocessingevent.AudioProcessingEvent)) {
	macro.Rewrite("$_.onaudioprocess")
	return onaudioprocess
}

// SetOnaudioprocess prop
// js:"onaudioprocess"
func (*ScriptProcessorNode) SetOnaudioprocess(onaudioprocess func(*audioprocessingevent.AudioProcessingEvent)) {
	macro.Rewrite("$_.onaudioprocess = $1", onaudioprocess)
}

// ChannelCount prop
// js:"channelCount"
func (*ScriptProcessorNode) ChannelCount() (channelCount uint) {
	macro.Rewrite("$_.channelCount")
	return channelCount
}

// SetChannelCount prop
// js:"channelCount"
func (*ScriptProcessorNode) SetChannelCount(channelCount uint) {
	macro.Rewrite("$_.channelCount = $1", channelCount)
}

// ChannelCountMode prop
// js:"channelCountMode"
func (*ScriptProcessorNode) ChannelCountMode() (channelCountMode *channelcountmode.ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode")
	return channelCountMode
}

// SetChannelCountMode prop
// js:"channelCountMode"
func (*ScriptProcessorNode) SetChannelCountMode(channelCountMode *channelcountmode.ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode = $1", channelCountMode)
}

// ChannelInterpretation prop
// js:"channelInterpretation"
func (*ScriptProcessorNode) ChannelInterpretation() (channelInterpretation *channelinterpretation.ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation")
	return channelInterpretation
}

// SetChannelInterpretation prop
// js:"channelInterpretation"
func (*ScriptProcessorNode) SetChannelInterpretation(channelInterpretation *channelinterpretation.ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation = $1", channelInterpretation)
}

// Context prop
// js:"context"
func (*ScriptProcessorNode) Context() (context AudioContext) {
	macro.Rewrite("$_.context")
	return context
}

// NumberOfInputs prop
// js:"numberOfInputs"
func (*ScriptProcessorNode) NumberOfInputs() (numberOfInputs uint) {
	macro.Rewrite("$_.numberOfInputs")
	return numberOfInputs
}

// NumberOfOutputs prop
// js:"numberOfOutputs"
func (*ScriptProcessorNode) NumberOfOutputs() (numberOfOutputs uint) {
	macro.Rewrite("$_.numberOfOutputs")
	return numberOfOutputs
}
