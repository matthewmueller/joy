package dynamicscompressornode

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/audionode"
	"github.com/matthewmueller/golly/internal/gendom/dom/audioparam"
	"github.com/matthewmueller/golly/js"
)

type DynamicsCompressorNode struct {
	*audionode.AudioNode
}

func (*DynamicsCompressorNode) GetAttack() (attack *audioparam.AudioParam) {
	js.Rewrite("$<.attack")
	return attack
}

func (*DynamicsCompressorNode) GetKnee() (knee *audioparam.AudioParam) {
	js.Rewrite("$<.knee")
	return knee
}

func (*DynamicsCompressorNode) GetRatio() (ratio *audioparam.AudioParam) {
	js.Rewrite("$<.ratio")
	return ratio
}

func (*DynamicsCompressorNode) GetReduction() (reduction float32) {
	js.Rewrite("$<.reduction")
	return reduction
}

func (*DynamicsCompressorNode) GetRelease() (release *audioparam.AudioParam) {
	js.Rewrite("$<.release")
	return release
}

func (*DynamicsCompressorNode) GetThreshold() (threshold *audioparam.AudioParam) {
	js.Rewrite("$<.threshold")
	return threshold
}
