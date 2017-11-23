package audionode

import (
	"github.com/matthewmueller/golly/dom/audioparam"
	"github.com/matthewmueller/golly/js"
)

// DelayNode struct
// js:"DelayNode,omit"
type DelayNode struct {
	AudioNode
}

// DelayTime prop
func (*DelayNode) DelayTime() (delayTime *audioparam.AudioParam) {
	js.Rewrite("$<.delayTime")
	return delayTime
}
