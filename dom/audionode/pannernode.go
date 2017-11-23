package audionode

import (
	"github.com/matthewmueller/golly/dom2/distancemodeltype"
	"github.com/matthewmueller/golly/dom2/panningmodeltype"
	"github.com/matthewmueller/golly/js"
)

// PannerNode struct
// js:"PannerNode,omit"
type PannerNode struct {
	AudioNode
}

// SetOrientation fn
func (*PannerNode) SetOrientation(x float32, y float32, z float32) {
	js.Rewrite("$<.setOrientation($1, $2, $3)", x, y, z)
}

// SetPosition fn
func (*PannerNode) SetPosition(x float32, y float32, z float32) {
	js.Rewrite("$<.setPosition($1, $2, $3)", x, y, z)
}

// SetVelocity fn
func (*PannerNode) SetVelocity(x float32, y float32, z float32) {
	js.Rewrite("$<.setVelocity($1, $2, $3)", x, y, z)
}

// ConeInnerAngle prop
func (*PannerNode) ConeInnerAngle() (coneInnerAngle float32) {
	js.Rewrite("$<.coneInnerAngle")
	return coneInnerAngle
}

// ConeInnerAngle prop
func (*PannerNode) SetConeInnerAngle(coneInnerAngle float32) {
	js.Rewrite("$<.coneInnerAngle = $1", coneInnerAngle)
}

// ConeOuterAngle prop
func (*PannerNode) ConeOuterAngle() (coneOuterAngle float32) {
	js.Rewrite("$<.coneOuterAngle")
	return coneOuterAngle
}

// ConeOuterAngle prop
func (*PannerNode) SetConeOuterAngle(coneOuterAngle float32) {
	js.Rewrite("$<.coneOuterAngle = $1", coneOuterAngle)
}

// ConeOuterGain prop
func (*PannerNode) ConeOuterGain() (coneOuterGain float32) {
	js.Rewrite("$<.coneOuterGain")
	return coneOuterGain
}

// ConeOuterGain prop
func (*PannerNode) SetConeOuterGain(coneOuterGain float32) {
	js.Rewrite("$<.coneOuterGain = $1", coneOuterGain)
}

// DistanceModel prop
func (*PannerNode) DistanceModel() (distanceModel *distancemodeltype.DistanceModelType) {
	js.Rewrite("$<.distanceModel")
	return distanceModel
}

// DistanceModel prop
func (*PannerNode) SetDistanceModel(distanceModel *distancemodeltype.DistanceModelType) {
	js.Rewrite("$<.distanceModel = $1", distanceModel)
}

// MaxDistance prop
func (*PannerNode) MaxDistance() (maxDistance float32) {
	js.Rewrite("$<.maxDistance")
	return maxDistance
}

// MaxDistance prop
func (*PannerNode) SetMaxDistance(maxDistance float32) {
	js.Rewrite("$<.maxDistance = $1", maxDistance)
}

// PanningModel prop
func (*PannerNode) PanningModel() (panningModel *panningmodeltype.PanningModelType) {
	js.Rewrite("$<.panningModel")
	return panningModel
}

// PanningModel prop
func (*PannerNode) SetPanningModel(panningModel *panningmodeltype.PanningModelType) {
	js.Rewrite("$<.panningModel = $1", panningModel)
}

// RefDistance prop
func (*PannerNode) RefDistance() (refDistance float32) {
	js.Rewrite("$<.refDistance")
	return refDistance
}

// RefDistance prop
func (*PannerNode) SetRefDistance(refDistance float32) {
	js.Rewrite("$<.refDistance = $1", refDistance)
}

// RolloffFactor prop
func (*PannerNode) RolloffFactor() (rolloffFactor float32) {
	js.Rewrite("$<.rolloffFactor")
	return rolloffFactor
}

// RolloffFactor prop
func (*PannerNode) SetRolloffFactor(rolloffFactor float32) {
	js.Rewrite("$<.rolloffFactor = $1", rolloffFactor)
}
