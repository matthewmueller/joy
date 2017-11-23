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

// SetPeriodicWave fn
func (*OscillatorNode) SetPeriodicWave(periodicWave *periodicwave.PeriodicWave) {
	js.Rewrite("$<.setPeriodicWave($1)", periodicWave)
}

// Start fn
func (*OscillatorNode) Start(when *float32) {
	js.Rewrite("$<.start($1)", when)
}

// Stop fn
func (*OscillatorNode) Stop(when *float32) {
	js.Rewrite("$<.stop($1)", when)
}

// Detune prop
func (*OscillatorNode) Detune() (detune *audioparam.AudioParam) {
	js.Rewrite("$<.detune")
	return detune
}

// Frequency prop
func (*OscillatorNode) Frequency() (frequency *audioparam.AudioParam) {
	js.Rewrite("$<.frequency")
	return frequency
}

// Onended prop
func (*OscillatorNode) Onended() (onended func(window.Event)) {
	js.Rewrite("$<.onended")
	return onended
}

// Onended prop
func (*OscillatorNode) SetOnended(onended func(window.Event)) {
	js.Rewrite("$<.onended = $1", onended)
}

// Type prop
func (*OscillatorNode) Type() (kind *oscillatortype.OscillatorType) {
	js.Rewrite("$<.type")
	return kind
}

// Type prop
func (*OscillatorNode) SetType(kind *oscillatortype.OscillatorType) {
	js.Rewrite("$<.type = $1", kind)
}
