package audionode

import (
	"github.com/matthewmueller/golly/dom2/oversampletype"
	"github.com/matthewmueller/golly/js"
)

// WaveShaperNode struct
// js:"WaveShaperNode,omit"
type WaveShaperNode struct {
	AudioNode
}

// Curve
func (*WaveShaperNode) Curve() (curve []float32) {
	js.Rewrite("$<.Curve")
	return curve
}

// Curve
func (*WaveShaperNode) SetCurve(curve []float32) {
	js.Rewrite("$<.Curve = $1", curve)
}

// Oversample
func (*WaveShaperNode) Oversample() (oversample *oversampletype.OverSampleType) {
	js.Rewrite("$<.Oversample")
	return oversample
}

// Oversample
func (*WaveShaperNode) SetOversample(oversample *oversampletype.OverSampleType) {
	js.Rewrite("$<.Oversample = $1", oversample)
}
