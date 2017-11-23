package audionode

import (
	"github.com/matthewmueller/golly/dom2/audioparam"
	"github.com/matthewmueller/golly/js"
)

// js:"BiquadFilterNode,omit"
type BiquadFilterNode struct {
	AudioNode
}

// GetFrequencyResponse
func (*BiquadFilterNode) GetFrequencyResponse(frequencyHz []float32, magResponse []float32, phaseResponse []float32) {
	js.Rewrite("$<.GetFrequencyResponse($1, $2, $3)", frequencyHz, magResponse, phaseResponse)
}

// Detune
func (*BiquadFilterNode) Detune() (detune *audioparam.AudioParam) {
	js.Rewrite("$<.Detune")
	return detune
}

// Frequency
func (*BiquadFilterNode) Frequency() (frequency *audioparam.AudioParam) {
	js.Rewrite("$<.Frequency")
	return frequency
}

// Gain
func (*BiquadFilterNode) Gain() (gain *audioparam.AudioParam) {
	js.Rewrite("$<.Gain")
	return gain
}

// Q
func (*BiquadFilterNode) Q() (Q *audioparam.AudioParam) {
	js.Rewrite("$<.Q")
	return Q
}

// Type
func (*BiquadFilterNode) Type() (kind *biquadfiltertype.BiquadFilterType) {
	js.Rewrite("$<.Type")
	return kind
}

// Type
func (*BiquadFilterNode) SetType(kind *biquadfiltertype.BiquadFilterType) {
	js.Rewrite("$<.Type = $1", kind)
}
