package audionode

import (
	"github.com/matthewmueller/joy/dom/audioparam"
	"github.com/matthewmueller/joy/dom/channelcountmode"
	"github.com/matthewmueller/joy/dom/channelinterpretation"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ AudioNode = (*DynamicsCompressorNode)(nil)
var _ window.EventTarget = (*DynamicsCompressorNode)(nil)

// DynamicsCompressorNode struct
// js:"DynamicsCompressorNode,omit"
type DynamicsCompressorNode struct {
}

// Connect fn
// js:"connect"
func (*DynamicsCompressorNode) Connect(destination AudioNode, output *uint, input *uint) (a AudioNode) {
	macro.Rewrite("$_.connect($1, $2, $3)", destination, output, input)
	return a
}

// Disconnect fn
// js:"disconnect"
func (*DynamicsCompressorNode) Disconnect() {
	macro.Rewrite("$_.disconnect()")
}

// AddEventListener fn
// js:"addEventListener"
func (*DynamicsCompressorNode) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*DynamicsCompressorNode) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*DynamicsCompressorNode) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Attack prop
// js:"attack"
func (*DynamicsCompressorNode) Attack() (attack *audioparam.AudioParam) {
	macro.Rewrite("$_.attack")
	return attack
}

// Knee prop
// js:"knee"
func (*DynamicsCompressorNode) Knee() (knee *audioparam.AudioParam) {
	macro.Rewrite("$_.knee")
	return knee
}

// Ratio prop
// js:"ratio"
func (*DynamicsCompressorNode) Ratio() (ratio *audioparam.AudioParam) {
	macro.Rewrite("$_.ratio")
	return ratio
}

// Reduction prop
// js:"reduction"
func (*DynamicsCompressorNode) Reduction() (reduction float32) {
	macro.Rewrite("$_.reduction")
	return reduction
}

// Release prop
// js:"release"
func (*DynamicsCompressorNode) Release() (release *audioparam.AudioParam) {
	macro.Rewrite("$_.release")
	return release
}

// Threshold prop
// js:"threshold"
func (*DynamicsCompressorNode) Threshold() (threshold *audioparam.AudioParam) {
	macro.Rewrite("$_.threshold")
	return threshold
}

// ChannelCount prop
// js:"channelCount"
func (*DynamicsCompressorNode) ChannelCount() (channelCount uint) {
	macro.Rewrite("$_.channelCount")
	return channelCount
}

// SetChannelCount prop
// js:"channelCount"
func (*DynamicsCompressorNode) SetChannelCount(channelCount uint) {
	macro.Rewrite("$_.channelCount = $1", channelCount)
}

// ChannelCountMode prop
// js:"channelCountMode"
func (*DynamicsCompressorNode) ChannelCountMode() (channelCountMode *channelcountmode.ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode")
	return channelCountMode
}

// SetChannelCountMode prop
// js:"channelCountMode"
func (*DynamicsCompressorNode) SetChannelCountMode(channelCountMode *channelcountmode.ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode = $1", channelCountMode)
}

// ChannelInterpretation prop
// js:"channelInterpretation"
func (*DynamicsCompressorNode) ChannelInterpretation() (channelInterpretation *channelinterpretation.ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation")
	return channelInterpretation
}

// SetChannelInterpretation prop
// js:"channelInterpretation"
func (*DynamicsCompressorNode) SetChannelInterpretation(channelInterpretation *channelinterpretation.ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation = $1", channelInterpretation)
}

// Context prop
// js:"context"
func (*DynamicsCompressorNode) Context() (context AudioContext) {
	macro.Rewrite("$_.context")
	return context
}

// NumberOfInputs prop
// js:"numberOfInputs"
func (*DynamicsCompressorNode) NumberOfInputs() (numberOfInputs uint) {
	macro.Rewrite("$_.numberOfInputs")
	return numberOfInputs
}

// NumberOfOutputs prop
// js:"numberOfOutputs"
func (*DynamicsCompressorNode) NumberOfOutputs() (numberOfOutputs uint) {
	macro.Rewrite("$_.numberOfOutputs")
	return numberOfOutputs
}
