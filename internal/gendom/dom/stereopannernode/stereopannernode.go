package stereopannernode

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/audionode"
	"github.com/matthewmueller/golly/internal/gendom/dom/audioparam"
	"github.com/matthewmueller/golly/js"
)

type StereoPannerNode struct {
	*audionode.AudioNode
}

func (*StereoPannerNode) GetPan() (pan *audioparam.AudioParam) {
	js.Rewrite("$<.pan")
	return pan
}
