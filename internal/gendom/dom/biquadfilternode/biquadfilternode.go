package biquadfilternode

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/audionode"
	"github.com/matthewmueller/golly/internal/gendom/dom/audioparam"
	"github.com/matthewmueller/golly/internal/gendom/dom/biquadfiltertype"
	"github.com/matthewmueller/golly/js"
)

type BiquadFilterNode struct {
	*audionode.AudioNode
}

func (*BiquadFilterNode) GetFrequencyResponse(frequencyHz []float32, magResponse []float32, phaseResponse []float32) {
	js.Rewrite("$<.getFrequencyResponse($1, $2, $3)", frequencyHz, magResponse, phaseResponse)
}

func (*BiquadFilterNode) GetDetune() (detune *audioparam.AudioParam) {
	js.Rewrite("$<.detune")
	return detune
}

func (*BiquadFilterNode) GetFrequency() (frequency *audioparam.AudioParam) {
	js.Rewrite("$<.frequency")
	return frequency
}

func (*BiquadFilterNode) GetGain() (gain *audioparam.AudioParam) {
	js.Rewrite("$<.gain")
	return gain
}

func (*BiquadFilterNode) GetQ() (Q *audioparam.AudioParam) {
	js.Rewrite("$<.Q")
	return Q
}

func (*BiquadFilterNode) GetType() (kind *biquadfiltertype.BiquadFilterType) {
	js.Rewrite("$<.type")
	return kind
}

func (*BiquadFilterNode) SetType(kind *biquadfiltertype.BiquadFilterType) {
	js.Rewrite("$<.type = $1", kind)
}
