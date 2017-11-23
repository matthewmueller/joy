package audionode

import (
	"github.com/matthewmueller/golly/dom2/audioparam"
	"github.com/matthewmueller/golly/js"
)

// js:"GainNode,omit"
type GainNode struct {
	AudioNode
}

// Gain
func (*GainNode) Gain() (gain *audioparam.AudioParam) {
	js.Rewrite("$<.Gain")
	return gain
}
