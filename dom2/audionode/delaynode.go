package audionode

import (
	"github.com/matthewmueller/golly/dom2/audioparam"
	"github.com/matthewmueller/golly/js"
)

// DelayNode struct
// js:"DelayNode,omit"
type DelayNode struct {
	AudioNode
}

// DelayTime
func (*DelayNode) DelayTime() (delayTime *audioparam.AudioParam) {
	js.Rewrite("$<.DelayTime")
	return delayTime
}
