package waveshapernode

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/audionode"
	"github.com/matthewmueller/golly/internal/gendom/dom/oversampletype"
	"github.com/matthewmueller/golly/js"
)

type WaveShaperNode struct {
	*audionode.AudioNode
}

func (*WaveShaperNode) GetCurve() (curve []float32) {
	js.Rewrite("$<.curve")
	return curve
}

func (*WaveShaperNode) SetCurve(curve []float32) {
	js.Rewrite("$<.curve = $1", curve)
}

func (*WaveShaperNode) GetOversample() (oversample *oversampletype.OverSampleType) {
	js.Rewrite("$<.oversample")
	return oversample
}

func (*WaveShaperNode) SetOversample(oversample *oversampletype.OverSampleType) {
	js.Rewrite("$<.oversample = $1", oversample)
}
