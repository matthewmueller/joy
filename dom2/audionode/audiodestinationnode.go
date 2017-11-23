package audionode

import "github.com/matthewmueller/golly/js"

// js:"AudioDestinationNode,omit"
type AudioDestinationNode struct {
	AudioNode
}

// MaxChannelCount
func (*AudioDestinationNode) MaxChannelCount() (maxChannelCount uint) {
	js.Rewrite("$<.MaxChannelCount")
	return maxChannelCount
}
