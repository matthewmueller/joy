package audionode

import (
	"github.com/matthewmueller/joy/dom/audioparam"
	"github.com/matthewmueller/joy/dom/channelcountmode"
	"github.com/matthewmueller/joy/dom/channelinterpretation"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/js"
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
	js.Rewrite("$_.connect($1, $2, $3)", destination, output, input)
	return a
}

// Disconnect fn
// js:"disconnect"
func (*DynamicsCompressorNode) Disconnect() {
	js.Rewrite("$_.disconnect()")
}

// AddEventListener fn
// js:"addEventListener"
func (*DynamicsCompressorNode) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*DynamicsCompressorNode) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*DynamicsCompressorNode) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Attack prop
// js:"attack"
func (*DynamicsCompressorNode) Attack() (attack *audioparam.AudioParam) {
	js.Rewrite("$_.attack")
	return attack
}

// Knee prop
// js:"knee"
func (*DynamicsCompressorNode) Knee() (knee *audioparam.AudioParam) {
	js.Rewrite("$_.knee")
	return knee
}

// Ratio prop
// js:"ratio"
func (*DynamicsCompressorNode) Ratio() (ratio *audioparam.AudioParam) {
	js.Rewrite("$_.ratio")
	return ratio
}

// Reduction prop
// js:"reduction"
func (*DynamicsCompressorNode) Reduction() (reduction float32) {
	js.Rewrite("$_.reduction")
	return reduction
}

// Release prop
// js:"release"
func (*DynamicsCompressorNode) Release() (release *audioparam.AudioParam) {
	js.Rewrite("$_.release")
	return release
}

// Threshold prop
// js:"threshold"
func (*DynamicsCompressorNode) Threshold() (threshold *audioparam.AudioParam) {
	js.Rewrite("$_.threshold")
	return threshold
}

// ChannelCount prop
// js:"channelCount"
func (*DynamicsCompressorNode) ChannelCount() (channelCount uint) {
	js.Rewrite("$_.channelCount")
	return channelCount
}

// SetChannelCount prop
// js:"channelCount"
func (*DynamicsCompressorNode) SetChannelCount(channelCount uint) {
	js.Rewrite("$_.channelCount = $1", channelCount)
}

// ChannelCountMode prop
// js:"channelCountMode"
func (*DynamicsCompressorNode) ChannelCountMode() (channelCountMode *channelcountmode.ChannelCountMode) {
	js.Rewrite("$_.channelCountMode")
	return channelCountMode
}

// SetChannelCountMode prop
// js:"channelCountMode"
func (*DynamicsCompressorNode) SetChannelCountMode(channelCountMode *channelcountmode.ChannelCountMode) {
	js.Rewrite("$_.channelCountMode = $1", channelCountMode)
}

// ChannelInterpretation prop
// js:"channelInterpretation"
func (*DynamicsCompressorNode) ChannelInterpretation() (channelInterpretation *channelinterpretation.ChannelInterpretation) {
	js.Rewrite("$_.channelInterpretation")
	return channelInterpretation
}

// SetChannelInterpretation prop
// js:"channelInterpretation"
func (*DynamicsCompressorNode) SetChannelInterpretation(channelInterpretation *channelinterpretation.ChannelInterpretation) {
	js.Rewrite("$_.channelInterpretation = $1", channelInterpretation)
}

// Context prop
// js:"context"
func (*DynamicsCompressorNode) Context() (context AudioContext) {
	js.Rewrite("$_.context")
	return context
}

// NumberOfInputs prop
// js:"numberOfInputs"
func (*DynamicsCompressorNode) NumberOfInputs() (numberOfInputs uint) {
	js.Rewrite("$_.numberOfInputs")
	return numberOfInputs
}

// NumberOfOutputs prop
// js:"numberOfOutputs"
func (*DynamicsCompressorNode) NumberOfOutputs() (numberOfOutputs uint) {
	js.Rewrite("$_.numberOfOutputs")
	return numberOfOutputs
}
