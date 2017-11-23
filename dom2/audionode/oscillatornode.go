package audionode

import (
	"github.com/matthewmueller/golly/dom2/audioparam"
	"github.com/matthewmueller/golly/dom2/oscillatortype"
	"github.com/matthewmueller/golly/dom2/periodicwave"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// OscillatorNode struct
// js:"OscillatorNode,omit"
type OscillatorNode struct {
	AudioNode
}

// SetPeriodicWave
func (*OscillatorNode) SetPeriodicWave(periodicWave *periodicwave.PeriodicWave) {
	js.Rewrite("$<.SetPeriodicWave($1)", periodicWave)
}

// Start
func (*OscillatorNode) Start(when *float32) {
	js.Rewrite("$<.Start($1)", when)
}

// Stop
func (*OscillatorNode) Stop(when *float32) {
	js.Rewrite("$<.Stop($1)", when)
}

// Detune
func (*OscillatorNode) Detune() (detune *audioparam.AudioParam) {
	js.Rewrite("$<.Detune")
	return detune
}

// Frequency
func (*OscillatorNode) Frequency() (frequency *audioparam.AudioParam) {
	js.Rewrite("$<.Frequency")
	return frequency
}

// Onended
func (*OscillatorNode) Onended() (onended func(window.Event)) {
	js.Rewrite("$<.Onended")
	return onended
}

// Onended
func (*OscillatorNode) SetOnended(onended func(window.Event)) {
	js.Rewrite("$<.Onended = $1", onended)
}

// Type
func (*OscillatorNode) Type() (kind *oscillatortype.OscillatorType) {
	js.Rewrite("$<.Type")
	return kind
}

// Type
func (*OscillatorNode) SetType(kind *oscillatortype.OscillatorType) {
	js.Rewrite("$<.Type = $1", kind)
}
