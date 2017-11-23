package audionode

import "github.com/matthewmueller/golly/js"

// AudioDestinationNode struct
// js:"AudioDestinationNode,omit"
type AudioDestinationNode struct {
	AudioNode
}

// MaxChannelCount
func (*AudioDestinationNode) MaxChannelCount() (maxChannelCount uint) {
	js.Rewrite("$<.MaxChannelCount")
	return maxChannelCount
}
