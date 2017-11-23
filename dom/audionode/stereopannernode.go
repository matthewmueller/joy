package audionode

import (
	"github.com/matthewmueller/golly/dom/audioparam"
	"github.com/matthewmueller/golly/js"
)

// StereoPannerNode struct
// js:"StereoPannerNode,omit"
type StereoPannerNode struct {
	AudioNode
}

// Pan prop
func (*StereoPannerNode) Pan() (pan *audioparam.AudioParam) {
	js.Rewrite("$<.pan")
	return pan
}
