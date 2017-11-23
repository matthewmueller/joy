package audionode

import (
	"github.com/matthewmueller/golly/dom/audioparam"
	"github.com/matthewmueller/golly/dom/biquadfiltertype"
	"github.com/matthewmueller/golly/js"
)

// BiquadFilterNode struct
// js:"BiquadFilterNode,omit"
type BiquadFilterNode struct {
	AudioNode
}

// GetFrequencyResponse fn
func (*BiquadFilterNode) GetFrequencyResponse(frequencyHz []float32, magResponse []float32, phaseResponse []float32) {
	js.Rewrite("$<.getFrequencyResponse($1, $2, $3)", frequencyHz, magResponse, phaseResponse)
}

// Detune prop
func (*BiquadFilterNode) Detune() (detune *audioparam.AudioParam) {
	js.Rewrite("$<.detune")
	return detune
}

// Frequency prop
func (*BiquadFilterNode) Frequency() (frequency *audioparam.AudioParam) {
	js.Rewrite("$<.frequency")
	return frequency
}

// Gain prop
func (*BiquadFilterNode) Gain() (gain *audioparam.AudioParam) {
	js.Rewrite("$<.gain")
	return gain
}

// Q prop
func (*BiquadFilterNode) Q() (Q *audioparam.AudioParam) {
	js.Rewrite("$<.Q")
	return Q
}

// Type prop
func (*BiquadFilterNode) Type() (kind *biquadfiltertype.BiquadFilterType) {
	js.Rewrite("$<.type")
	return kind
}

// Type prop
func (*BiquadFilterNode) SetType(kind *biquadfiltertype.BiquadFilterType) {
	js.Rewrite("$<.type = $1", kind)
}
