package audionode

import (
	"github.com/matthewmueller/joy/dom/audioparam"
	"github.com/matthewmueller/joy/dom/channelcountmode"
	"github.com/matthewmueller/joy/dom/channelinterpretation"
	"github.com/matthewmueller/joy/dom/oscillatortype"
	"github.com/matthewmueller/joy/dom/periodicwave"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ AudioNode = (*OscillatorNode)(nil)
var _ window.EventTarget = (*OscillatorNode)(nil)

// OscillatorNode struct
// js:"OscillatorNode,omit"
type OscillatorNode struct {
}

// SetPeriodicWave fn
// js:"setPeriodicWave"
func (*OscillatorNode) SetPeriodicWave(periodicWave *periodicwave.PeriodicWave) {
	macro.Rewrite("$_.setPeriodicWave($1)", periodicWave)
}

// Start fn
// js:"start"
func (*OscillatorNode) Start(when *float32) {
	macro.Rewrite("$_.start($1)", when)
}

// Stop fn
// js:"stop"
func (*OscillatorNode) Stop(when *float32) {
	macro.Rewrite("$_.stop($1)", when)
}

// Connect fn
// js:"connect"
func (*OscillatorNode) Connect(destination AudioNode, output *uint, input *uint) (a AudioNode) {
	macro.Rewrite("$_.connect($1, $2, $3)", destination, output, input)
	return a
}

// Disconnect fn
// js:"disconnect"
func (*OscillatorNode) Disconnect() {
	macro.Rewrite("$_.disconnect()")
}

// AddEventListener fn
// js:"addEventListener"
func (*OscillatorNode) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*OscillatorNode) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*OscillatorNode) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Detune prop
// js:"detune"
func (*OscillatorNode) Detune() (detune *audioparam.AudioParam) {
	macro.Rewrite("$_.detune")
	return detune
}

// Frequency prop
// js:"frequency"
func (*OscillatorNode) Frequency() (frequency *audioparam.AudioParam) {
	macro.Rewrite("$_.frequency")
	return frequency
}

// Onended prop
// js:"onended"
func (*OscillatorNode) Onended() (onended func(window.Event)) {
	macro.Rewrite("$_.onended")
	return onended
}

// SetOnended prop
// js:"onended"
func (*OscillatorNode) SetOnended(onended func(window.Event)) {
	macro.Rewrite("$_.onended = $1", onended)
}

// Type prop
// js:"type"
func (*OscillatorNode) Type() (kind *oscillatortype.OscillatorType) {
	macro.Rewrite("$_.type")
	return kind
}

// SetType prop
// js:"type"
func (*OscillatorNode) SetType(kind *oscillatortype.OscillatorType) {
	macro.Rewrite("$_.type = $1", kind)
}

// ChannelCount prop
// js:"channelCount"
func (*OscillatorNode) ChannelCount() (channelCount uint) {
	macro.Rewrite("$_.channelCount")
	return channelCount
}

// SetChannelCount prop
// js:"channelCount"
func (*OscillatorNode) SetChannelCount(channelCount uint) {
	macro.Rewrite("$_.channelCount = $1", channelCount)
}

// ChannelCountMode prop
// js:"channelCountMode"
func (*OscillatorNode) ChannelCountMode() (channelCountMode *channelcountmode.ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode")
	return channelCountMode
}

// SetChannelCountMode prop
// js:"channelCountMode"
func (*OscillatorNode) SetChannelCountMode(channelCountMode *channelcountmode.ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode = $1", channelCountMode)
}

// ChannelInterpretation prop
// js:"channelInterpretation"
func (*OscillatorNode) ChannelInterpretation() (channelInterpretation *channelinterpretation.ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation")
	return channelInterpretation
}

// SetChannelInterpretation prop
// js:"channelInterpretation"
func (*OscillatorNode) SetChannelInterpretation(channelInterpretation *channelinterpretation.ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation = $1", channelInterpretation)
}

// Context prop
// js:"context"
func (*OscillatorNode) Context() (context AudioContext) {
	macro.Rewrite("$_.context")
	return context
}

// NumberOfInputs prop
// js:"numberOfInputs"
func (*OscillatorNode) NumberOfInputs() (numberOfInputs uint) {
	macro.Rewrite("$_.numberOfInputs")
	return numberOfInputs
}

// NumberOfOutputs prop
// js:"numberOfOutputs"
func (*OscillatorNode) NumberOfOutputs() (numberOfOutputs uint) {
	macro.Rewrite("$_.numberOfOutputs")
	return numberOfOutputs
}
