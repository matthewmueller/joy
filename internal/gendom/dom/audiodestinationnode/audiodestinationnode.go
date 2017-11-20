package audiodestinationnode

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/audionode"
	"github.com/matthewmueller/golly/js"
)

type AudioDestinationNode struct {
	*audionode.AudioNode
}

func (*AudioDestinationNode) GetMaxChannelCount() (maxChannelCount uint) {
	js.Rewrite("$<.maxChannelCount")
	return maxChannelCount
}
