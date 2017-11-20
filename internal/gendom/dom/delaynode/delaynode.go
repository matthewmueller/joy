package delaynode

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/audionode"
	"github.com/matthewmueller/golly/internal/gendom/dom/audioparam"
	"github.com/matthewmueller/golly/js"
)

type DelayNode struct {
	*audionode.AudioNode
}

func (*DelayNode) GetDelayTime() (delayTime *audioparam.AudioParam) {
	js.Rewrite("$<.delayTime")
	return delayTime
}
