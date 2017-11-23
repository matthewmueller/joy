package audionode

import "github.com/matthewmueller/golly/js"

// AudioDestinationNode struct
// js:"AudioDestinationNode,omit"
type AudioDestinationNode struct {
	AudioNode
}

// MaxChannelCount prop
func (*AudioDestinationNode) MaxChannelCount() (maxChannelCount uint) {
	js.Rewrite("$<.maxChannelCount")
	return maxChannelCount
}
