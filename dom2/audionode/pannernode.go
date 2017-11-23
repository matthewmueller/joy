package audionode

import (
	"github.com/matthewmueller/golly/dom2/distancemodeltype"
	"github.com/matthewmueller/golly/dom2/panningmodeltype"
	"github.com/matthewmueller/golly/js"
)

// js:"PannerNode,omit"
type PannerNode struct {
	AudioNode
}

// SetOrientation
func (*PannerNode) SetOrientation(x float32, y float32, z float32) {
	js.Rewrite("$<.SetOrientation($1, $2, $3)", x, y, z)
}

// SetPosition
func (*PannerNode) SetPosition(x float32, y float32, z float32) {
	js.Rewrite("$<.SetPosition($1, $2, $3)", x, y, z)
}

// SetVelocity
func (*PannerNode) SetVelocity(x float32, y float32, z float32) {
	js.Rewrite("$<.SetVelocity($1, $2, $3)", x, y, z)
}

// ConeInnerAngle
func (*PannerNode) ConeInnerAngle() (coneInnerAngle float32) {
	js.Rewrite("$<.ConeInnerAngle")
	return coneInnerAngle
}

// ConeInnerAngle
func (*PannerNode) SetConeInnerAngle(coneInnerAngle float32) {
	js.Rewrite("$<.ConeInnerAngle = $1", coneInnerAngle)
}

// ConeOuterAngle
func (*PannerNode) ConeOuterAngle() (coneOuterAngle float32) {
	js.Rewrite("$<.ConeOuterAngle")
	return coneOuterAngle
}

// ConeOuterAngle
func (*PannerNode) SetConeOuterAngle(coneOuterAngle float32) {
	js.Rewrite("$<.ConeOuterAngle = $1", coneOuterAngle)
}

// ConeOuterGain
func (*PannerNode) ConeOuterGain() (coneOuterGain float32) {
	js.Rewrite("$<.ConeOuterGain")
	return coneOuterGain
}

// ConeOuterGain
func (*PannerNode) SetConeOuterGain(coneOuterGain float32) {
	js.Rewrite("$<.ConeOuterGain = $1", coneOuterGain)
}

// DistanceModel
func (*PannerNode) DistanceModel() (distanceModel *distancemodeltype.DistanceModelType) {
	js.Rewrite("$<.DistanceModel")
	return distanceModel
}

// DistanceModel
func (*PannerNode) SetDistanceModel(distanceModel *distancemodeltype.DistanceModelType) {
	js.Rewrite("$<.DistanceModel = $1", distanceModel)
}

// MaxDistance
func (*PannerNode) MaxDistance() (maxDistance float32) {
	js.Rewrite("$<.MaxDistance")
	return maxDistance
}

// MaxDistance
func (*PannerNode) SetMaxDistance(maxDistance float32) {
	js.Rewrite("$<.MaxDistance = $1", maxDistance)
}

// PanningModel
func (*PannerNode) PanningModel() (panningModel *panningmodeltype.PanningModelType) {
	js.Rewrite("$<.PanningModel")
	return panningModel
}

// PanningModel
func (*PannerNode) SetPanningModel(panningModel *panningmodeltype.PanningModelType) {
	js.Rewrite("$<.PanningModel = $1", panningModel)
}

// RefDistance
func (*PannerNode) RefDistance() (refDistance float32) {
	js.Rewrite("$<.RefDistance")
	return refDistance
}

// RefDistance
func (*PannerNode) SetRefDistance(refDistance float32) {
	js.Rewrite("$<.RefDistance = $1", refDistance)
}

// RolloffFactor
func (*PannerNode) RolloffFactor() (rolloffFactor float32) {
	js.Rewrite("$<.RolloffFactor")
	return rolloffFactor
}

// RolloffFactor
func (*PannerNode) SetRolloffFactor(rolloffFactor float32) {
	js.Rewrite("$<.RolloffFactor = $1", rolloffFactor)
}
