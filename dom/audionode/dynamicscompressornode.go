package audionode

import (
	"github.com/matthewmueller/golly/dom/audioparam"
	"github.com/matthewmueller/golly/js"
)

// DynamicsCompressorNode struct
// js:"DynamicsCompressorNode,omit"
type DynamicsCompressorNode struct {
	AudioNode
}

// Attack prop
func (*DynamicsCompressorNode) Attack() (attack *audioparam.AudioParam) {
	js.Rewrite("$<.attack")
	return attack
}

// Knee prop
func (*DynamicsCompressorNode) Knee() (knee *audioparam.AudioParam) {
	js.Rewrite("$<.knee")
	return knee
}

// Ratio prop
func (*DynamicsCompressorNode) Ratio() (ratio *audioparam.AudioParam) {
	js.Rewrite("$<.ratio")
	return ratio
}

// Reduction prop
func (*DynamicsCompressorNode) Reduction() (reduction float32) {
	js.Rewrite("$<.reduction")
	return reduction
}

// Release prop
func (*DynamicsCompressorNode) Release() (release *audioparam.AudioParam) {
	js.Rewrite("$<.release")
	return release
}

// Threshold prop
func (*DynamicsCompressorNode) Threshold() (threshold *audioparam.AudioParam) {
	js.Rewrite("$<.threshold")
	return threshold
}
