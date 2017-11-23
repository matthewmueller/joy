package audionode

import (
	"github.com/matthewmueller/golly/dom2/audioparam"
	"github.com/matthewmueller/golly/js"
)

// DynamicsCompressorNode struct
// js:"DynamicsCompressorNode,omit"
type DynamicsCompressorNode struct {
	AudioNode
}

// Attack
func (*DynamicsCompressorNode) Attack() (attack *audioparam.AudioParam) {
	js.Rewrite("$<.Attack")
	return attack
}

// Knee
func (*DynamicsCompressorNode) Knee() (knee *audioparam.AudioParam) {
	js.Rewrite("$<.Knee")
	return knee
}

// Ratio
func (*DynamicsCompressorNode) Ratio() (ratio *audioparam.AudioParam) {
	js.Rewrite("$<.Ratio")
	return ratio
}

// Reduction
func (*DynamicsCompressorNode) Reduction() (reduction float32) {
	js.Rewrite("$<.Reduction")
	return reduction
}

// Release
func (*DynamicsCompressorNode) Release() (release *audioparam.AudioParam) {
	js.Rewrite("$<.Release")
	return release
}

// Threshold
func (*DynamicsCompressorNode) Threshold() (threshold *audioparam.AudioParam) {
	js.Rewrite("$<.Threshold")
	return threshold
}
