package audionode

import (
	"github.com/matthewmueller/golly/dom/audioparam"
	"github.com/matthewmueller/golly/dom/channelcountmode"
	"github.com/matthewmueller/golly/dom/channelinterpretation"
	"github.com/matthewmueller/golly/dom/oscillatortype"
	"github.com/matthewmueller/golly/dom/periodicwave"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
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
	js.Rewrite("$_.setPeriodicWave($1)", periodicWave)
}

// Start fn
// js:"start"
func (*OscillatorNode) Start(when *float32) {
	js.Rewrite("$_.start($1)", when)
}

// Stop fn
// js:"stop"
func (*OscillatorNode) Stop(when *float32) {
	js.Rewrite("$_.stop($1)", when)
}

// Connect fn
// js:"connect"
func (*OscillatorNode) Connect(destination AudioNode, output *uint, input *uint) (a AudioNode) {
	js.Rewrite("$_.connect($1, $2, $3)", destination, output, input)
	return a
}

// Disconnect fn
// js:"disconnect"
func (*OscillatorNode) Disconnect() {
	js.Rewrite("$_.disconnect()")
}

// AddEventListener fn
// js:"addEventListener"
func (*OscillatorNode) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*OscillatorNode) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*OscillatorNode) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Detune prop
// js:"detune"
func (*OscillatorNode) Detune() (detune *audioparam.AudioParam) {
	js.Rewrite("$_.detune")
	return detune
}

// Frequency prop
// js:"frequency"
func (*OscillatorNode) Frequency() (frequency *audioparam.AudioParam) {
	js.Rewrite("$_.frequency")
	return frequency
}

// Onended prop
// js:"onended"
func (*OscillatorNode) Onended() (onended func(window.Event)) {
	js.Rewrite("$_.onended")
	return onended
}

// SetOnended prop
// js:"onended"
func (*OscillatorNode) SetOnended(onended func(window.Event)) {
	js.Rewrite("$_.onended = $1", onended)
}

// Type prop
// js:"type"
func (*OscillatorNode) Type() (kind *oscillatortype.OscillatorType) {
	js.Rewrite("$_.type")
	return kind
}

// SetType prop
// js:"type"
func (*OscillatorNode) SetType(kind *oscillatortype.OscillatorType) {
	js.Rewrite("$_.type = $1", kind)
}

// ChannelCount prop
// js:"channelCount"
func (*OscillatorNode) ChannelCount() (channelCount uint) {
	js.Rewrite("$_.channelCount")
	return channelCount
}

// SetChannelCount prop
// js:"channelCount"
func (*OscillatorNode) SetChannelCount(channelCount uint) {
	js.Rewrite("$_.channelCount = $1", channelCount)
}

// ChannelCountMode prop
// js:"channelCountMode"
func (*OscillatorNode) ChannelCountMode() (channelCountMode *channelcountmode.ChannelCountMode) {
	js.Rewrite("$_.channelCountMode")
	return channelCountMode
}

// SetChannelCountMode prop
// js:"channelCountMode"
func (*OscillatorNode) SetChannelCountMode(channelCountMode *channelcountmode.ChannelCountMode) {
	js.Rewrite("$_.channelCountMode = $1", channelCountMode)
}

// ChannelInterpretation prop
// js:"channelInterpretation"
func (*OscillatorNode) ChannelInterpretation() (channelInterpretation *channelinterpretation.ChannelInterpretation) {
	js.Rewrite("$_.channelInterpretation")
	return channelInterpretation
}

// SetChannelInterpretation prop
// js:"channelInterpretation"
func (*OscillatorNode) SetChannelInterpretation(channelInterpretation *channelinterpretation.ChannelInterpretation) {
	js.Rewrite("$_.channelInterpretation = $1", channelInterpretation)
}

// Context prop
// js:"context"
func (*OscillatorNode) Context() (context AudioContext) {
	js.Rewrite("$_.context")
	return context
}

// NumberOfInputs prop
// js:"numberOfInputs"
func (*OscillatorNode) NumberOfInputs() (numberOfInputs uint) {
	js.Rewrite("$_.numberOfInputs")
	return numberOfInputs
}

// NumberOfOutputs prop
// js:"numberOfOutputs"
func (*OscillatorNode) NumberOfOutputs() (numberOfOutputs uint) {
	js.Rewrite("$_.numberOfOutputs")
	return numberOfOutputs
}
