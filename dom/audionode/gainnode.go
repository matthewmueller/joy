package audionode

import (
	"github.com/matthewmueller/golly/dom/audioparam"
	"github.com/matthewmueller/golly/js"
)

// GainNode struct
// js:"GainNode,omit"
type GainNode struct {
	AudioNode
}

// Gain prop
func (*GainNode) Gain() (gain *audioparam.AudioParam) {
	js.Rewrite("$<.gain")
	return gain
}
