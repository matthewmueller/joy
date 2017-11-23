package audionode

import (
	"github.com/matthewmueller/golly/dom2/audioparam"
	"github.com/matthewmueller/golly/js"
)

// StereoPannerNode struct
// js:"StereoPannerNode,omit"
type StereoPannerNode struct {
	AudioNode
}

// Pan
func (*StereoPannerNode) Pan() (pan *audioparam.AudioParam) {
	js.Rewrite("$<.Pan")
	return pan
}
