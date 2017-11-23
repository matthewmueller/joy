package audionode

import "github.com/matthewmueller/golly/js"

// IIRFilterNode struct
// js:"IIRFilterNode,omit"
type IIRFilterNode struct {
	AudioNode
}

// GetFrequencyResponse fn
func (*IIRFilterNode) GetFrequencyResponse(frequencyHz []float32, magResponse []float32, phaseResponse []float32) {
	js.Rewrite("$<.getFrequencyResponse($1, $2, $3)", frequencyHz, magResponse, phaseResponse)
}
