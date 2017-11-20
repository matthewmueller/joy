package iirfilternode

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/audionode"
	"github.com/matthewmueller/golly/js"
)

type IIRFilterNode struct {
	*audionode.AudioNode
}

func (*IIRFilterNode) GetFrequencyResponse(frequencyHz []float32, magResponse []float32, phaseResponse []float32) {
	js.Rewrite("$<.getFrequencyResponse($1, $2, $3)", frequencyHz, magResponse, phaseResponse)
}
