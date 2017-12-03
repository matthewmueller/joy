package audionode

import (
	"github.com/matthewmueller/joy/dom/audioparam"
	"github.com/matthewmueller/joy/dom/biquadfiltertype"
	"github.com/matthewmueller/joy/dom/channelcountmode"
	"github.com/matthewmueller/joy/dom/channelinterpretation"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ AudioNode = (*BiquadFilterNode)(nil)
var _ window.EventTarget = (*BiquadFilterNode)(nil)

// BiquadFilterNode struct
// js:"BiquadFilterNode,omit"
type BiquadFilterNode struct {
}

// GetFrequencyResponse fn
// js:"getFrequencyResponse"
func (*BiquadFilterNode) GetFrequencyResponse(frequencyHz []float32, magResponse []float32, phaseResponse []float32) {
	macro.Rewrite("$_.getFrequencyResponse($1, $2, $3)", frequencyHz, magResponse, phaseResponse)
}

// Connect fn
// js:"connect"
func (*BiquadFilterNode) Connect(destination AudioNode, output *uint, input *uint) (a AudioNode) {
	macro.Rewrite("$_.connect($1, $2, $3)", destination, output, input)
	return a
}

// Disconnect fn
// js:"disconnect"
func (*BiquadFilterNode) Disconnect() {
	macro.Rewrite("$_.disconnect()")
}

// AddEventListener fn
// js:"addEventListener"
func (*BiquadFilterNode) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*BiquadFilterNode) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*BiquadFilterNode) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Detune prop
// js:"detune"
func (*BiquadFilterNode) Detune() (detune *audioparam.AudioParam) {
	macro.Rewrite("$_.detune")
	return detune
}

// Frequency prop
// js:"frequency"
func (*BiquadFilterNode) Frequency() (frequency *audioparam.AudioParam) {
	macro.Rewrite("$_.frequency")
	return frequency
}

// Gain prop
// js:"gain"
func (*BiquadFilterNode) Gain() (gain *audioparam.AudioParam) {
	macro.Rewrite("$_.gain")
	return gain
}

// Q prop
// js:"Q"
func (*BiquadFilterNode) Q() (Q *audioparam.AudioParam) {
	macro.Rewrite("$_.Q")
	return Q
}

// Type prop
// js:"type"
func (*BiquadFilterNode) Type() (kind *biquadfiltertype.BiquadFilterType) {
	macro.Rewrite("$_.type")
	return kind
}

// SetType prop
// js:"type"
func (*BiquadFilterNode) SetType(kind *biquadfiltertype.BiquadFilterType) {
	macro.Rewrite("$_.type = $1", kind)
}

// ChannelCount prop
// js:"channelCount"
func (*BiquadFilterNode) ChannelCount() (channelCount uint) {
	macro.Rewrite("$_.channelCount")
	return channelCount
}

// SetChannelCount prop
// js:"channelCount"
func (*BiquadFilterNode) SetChannelCount(channelCount uint) {
	macro.Rewrite("$_.channelCount = $1", channelCount)
}

// ChannelCountMode prop
// js:"channelCountMode"
func (*BiquadFilterNode) ChannelCountMode() (channelCountMode *channelcountmode.ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode")
	return channelCountMode
}

// SetChannelCountMode prop
// js:"channelCountMode"
func (*BiquadFilterNode) SetChannelCountMode(channelCountMode *channelcountmode.ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode = $1", channelCountMode)
}

// ChannelInterpretation prop
// js:"channelInterpretation"
func (*BiquadFilterNode) ChannelInterpretation() (channelInterpretation *channelinterpretation.ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation")
	return channelInterpretation
}

// SetChannelInterpretation prop
// js:"channelInterpretation"
func (*BiquadFilterNode) SetChannelInterpretation(channelInterpretation *channelinterpretation.ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation = $1", channelInterpretation)
}

// Context prop
// js:"context"
func (*BiquadFilterNode) Context() (context AudioContext) {
	macro.Rewrite("$_.context")
	return context
}

// NumberOfInputs prop
// js:"numberOfInputs"
func (*BiquadFilterNode) NumberOfInputs() (numberOfInputs uint) {
	macro.Rewrite("$_.numberOfInputs")
	return numberOfInputs
}

// NumberOfOutputs prop
// js:"numberOfOutputs"
func (*BiquadFilterNode) NumberOfOutputs() (numberOfOutputs uint) {
	macro.Rewrite("$_.numberOfOutputs")
	return numberOfOutputs
}
