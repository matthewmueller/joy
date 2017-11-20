package pannernode

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/audionode"
	"github.com/matthewmueller/golly/internal/gendom/dom/distancemodeltype"
	"github.com/matthewmueller/golly/internal/gendom/dom/panningmodeltype"
	"github.com/matthewmueller/golly/js"
)

type PannerNode struct {
	*audionode.AudioNode
}

func (*PannerNode) SetOrientation(x float32, y float32, z float32) {
	js.Rewrite("$<.setOrientation($1, $2, $3)", x, y, z)
}

func (*PannerNode) SetPosition(x float32, y float32, z float32) {
	js.Rewrite("$<.setPosition($1, $2, $3)", x, y, z)
}

func (*PannerNode) SetVelocity(x float32, y float32, z float32) {
	js.Rewrite("$<.setVelocity($1, $2, $3)", x, y, z)
}

func (*PannerNode) GetConeInnerAngle() (coneInnerAngle float32) {
	js.Rewrite("$<.coneInnerAngle")
	return coneInnerAngle
}

func (*PannerNode) SetConeInnerAngle(coneInnerAngle float32) {
	js.Rewrite("$<.coneInnerAngle = $1", coneInnerAngle)
}

func (*PannerNode) GetConeOuterAngle() (coneOuterAngle float32) {
	js.Rewrite("$<.coneOuterAngle")
	return coneOuterAngle
}

func (*PannerNode) SetConeOuterAngle(coneOuterAngle float32) {
	js.Rewrite("$<.coneOuterAngle = $1", coneOuterAngle)
}

func (*PannerNode) GetConeOuterGain() (coneOuterGain float32) {
	js.Rewrite("$<.coneOuterGain")
	return coneOuterGain
}

func (*PannerNode) SetConeOuterGain(coneOuterGain float32) {
	js.Rewrite("$<.coneOuterGain = $1", coneOuterGain)
}

func (*PannerNode) GetDistanceModel() (distanceModel *distancemodeltype.DistanceModelType) {
	js.Rewrite("$<.distanceModel")
	return distanceModel
}

func (*PannerNode) SetDistanceModel(distanceModel *distancemodeltype.DistanceModelType) {
	js.Rewrite("$<.distanceModel = $1", distanceModel)
}

func (*PannerNode) GetMaxDistance() (maxDistance float32) {
	js.Rewrite("$<.maxDistance")
	return maxDistance
}

func (*PannerNode) SetMaxDistance(maxDistance float32) {
	js.Rewrite("$<.maxDistance = $1", maxDistance)
}

func (*PannerNode) GetPanningModel() (panningModel *panningmodeltype.PanningModelType) {
	js.Rewrite("$<.panningModel")
	return panningModel
}

func (*PannerNode) SetPanningModel(panningModel *panningmodeltype.PanningModelType) {
	js.Rewrite("$<.panningModel = $1", panningModel)
}

func (*PannerNode) GetRefDistance() (refDistance float32) {
	js.Rewrite("$<.refDistance")
	return refDistance
}

func (*PannerNode) SetRefDistance(refDistance float32) {
	js.Rewrite("$<.refDistance = $1", refDistance)
}

func (*PannerNode) GetRolloffFactor() (rolloffFactor float32) {
	js.Rewrite("$<.rolloffFactor")
	return rolloffFactor
}

func (*PannerNode) SetRolloffFactor(rolloffFactor float32) {
	js.Rewrite("$<.rolloffFactor = $1", rolloffFactor)
}
