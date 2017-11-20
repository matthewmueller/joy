package gainnode

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/audionode"
	"github.com/matthewmueller/golly/internal/gendom/dom/audioparam"
	"github.com/matthewmueller/golly/js"
)

type GainNode struct {
	*audionode.AudioNode
}

func (*GainNode) GetGain() (gain *audioparam.AudioParam) {
	js.Rewrite("$<.gain")
	return gain
}
