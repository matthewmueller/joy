package audionode

import (
	"github.com/matthewmueller/golly/dom/oversampletype"
	"github.com/matthewmueller/golly/js"
)

// WaveShaperNode struct
// js:"WaveShaperNode,omit"
type WaveShaperNode struct {
	AudioNode
}

// Curve prop
func (*WaveShaperNode) Curve() (curve []float32) {
	js.Rewrite("$<.curve")
	return curve
}

// Curve prop
func (*WaveShaperNode) SetCurve(curve []float32) {
	js.Rewrite("$<.curve = $1", curve)
}

// Oversample prop
func (*WaveShaperNode) Oversample() (oversample *oversampletype.OverSampleType) {
	js.Rewrite("$<.oversample")
	return oversample
}

// Oversample prop
func (*WaveShaperNode) SetOversample(oversample *oversampletype.OverSampleType) {
	js.Rewrite("$<.oversample = $1", oversample)
}
