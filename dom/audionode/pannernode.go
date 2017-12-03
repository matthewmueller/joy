package audionode

import (
	"github.com/matthewmueller/joy/dom/channelcountmode"
	"github.com/matthewmueller/joy/dom/channelinterpretation"
	"github.com/matthewmueller/joy/dom/distancemodeltype"
	"github.com/matthewmueller/joy/dom/panningmodeltype"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ AudioNode = (*PannerNode)(nil)
var _ window.EventTarget = (*PannerNode)(nil)

// PannerNode struct
// js:"PannerNode,omit"
type PannerNode struct {
}

// SetOrientation fn
// js:"setOrientation"
func (*PannerNode) SetOrientation(x float32, y float32, z float32) {
	macro.Rewrite("$_.setOrientation($1, $2, $3)", x, y, z)
}

// SetPosition fn
// js:"setPosition"
func (*PannerNode) SetPosition(x float32, y float32, z float32) {
	macro.Rewrite("$_.setPosition($1, $2, $3)", x, y, z)
}

// SetVelocity fn
// js:"setVelocity"
func (*PannerNode) SetVelocity(x float32, y float32, z float32) {
	macro.Rewrite("$_.setVelocity($1, $2, $3)", x, y, z)
}

// Connect fn
// js:"connect"
func (*PannerNode) Connect(destination AudioNode, output *uint, input *uint) (a AudioNode) {
	macro.Rewrite("$_.connect($1, $2, $3)", destination, output, input)
	return a
}

// Disconnect fn
// js:"disconnect"
func (*PannerNode) Disconnect() {
	macro.Rewrite("$_.disconnect()")
}

// AddEventListener fn
// js:"addEventListener"
func (*PannerNode) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*PannerNode) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*PannerNode) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// ConeInnerAngle prop
// js:"coneInnerAngle"
func (*PannerNode) ConeInnerAngle() (coneInnerAngle float32) {
	macro.Rewrite("$_.coneInnerAngle")
	return coneInnerAngle
}

// SetConeInnerAngle prop
// js:"coneInnerAngle"
func (*PannerNode) SetConeInnerAngle(coneInnerAngle float32) {
	macro.Rewrite("$_.coneInnerAngle = $1", coneInnerAngle)
}

// ConeOuterAngle prop
// js:"coneOuterAngle"
func (*PannerNode) ConeOuterAngle() (coneOuterAngle float32) {
	macro.Rewrite("$_.coneOuterAngle")
	return coneOuterAngle
}

// SetConeOuterAngle prop
// js:"coneOuterAngle"
func (*PannerNode) SetConeOuterAngle(coneOuterAngle float32) {
	macro.Rewrite("$_.coneOuterAngle = $1", coneOuterAngle)
}

// ConeOuterGain prop
// js:"coneOuterGain"
func (*PannerNode) ConeOuterGain() (coneOuterGain float32) {
	macro.Rewrite("$_.coneOuterGain")
	return coneOuterGain
}

// SetConeOuterGain prop
// js:"coneOuterGain"
func (*PannerNode) SetConeOuterGain(coneOuterGain float32) {
	macro.Rewrite("$_.coneOuterGain = $1", coneOuterGain)
}

// DistanceModel prop
// js:"distanceModel"
func (*PannerNode) DistanceModel() (distanceModel *distancemodeltype.DistanceModelType) {
	macro.Rewrite("$_.distanceModel")
	return distanceModel
}

// SetDistanceModel prop
// js:"distanceModel"
func (*PannerNode) SetDistanceModel(distanceModel *distancemodeltype.DistanceModelType) {
	macro.Rewrite("$_.distanceModel = $1", distanceModel)
}

// MaxDistance prop
// js:"maxDistance"
func (*PannerNode) MaxDistance() (maxDistance float32) {
	macro.Rewrite("$_.maxDistance")
	return maxDistance
}

// SetMaxDistance prop
// js:"maxDistance"
func (*PannerNode) SetMaxDistance(maxDistance float32) {
	macro.Rewrite("$_.maxDistance = $1", maxDistance)
}

// PanningModel prop
// js:"panningModel"
func (*PannerNode) PanningModel() (panningModel *panningmodeltype.PanningModelType) {
	macro.Rewrite("$_.panningModel")
	return panningModel
}

// SetPanningModel prop
// js:"panningModel"
func (*PannerNode) SetPanningModel(panningModel *panningmodeltype.PanningModelType) {
	macro.Rewrite("$_.panningModel = $1", panningModel)
}

// RefDistance prop
// js:"refDistance"
func (*PannerNode) RefDistance() (refDistance float32) {
	macro.Rewrite("$_.refDistance")
	return refDistance
}

// SetRefDistance prop
// js:"refDistance"
func (*PannerNode) SetRefDistance(refDistance float32) {
	macro.Rewrite("$_.refDistance = $1", refDistance)
}

// RolloffFactor prop
// js:"rolloffFactor"
func (*PannerNode) RolloffFactor() (rolloffFactor float32) {
	macro.Rewrite("$_.rolloffFactor")
	return rolloffFactor
}

// SetRolloffFactor prop
// js:"rolloffFactor"
func (*PannerNode) SetRolloffFactor(rolloffFactor float32) {
	macro.Rewrite("$_.rolloffFactor = $1", rolloffFactor)
}

// ChannelCount prop
// js:"channelCount"
func (*PannerNode) ChannelCount() (channelCount uint) {
	macro.Rewrite("$_.channelCount")
	return channelCount
}

// SetChannelCount prop
// js:"channelCount"
func (*PannerNode) SetChannelCount(channelCount uint) {
	macro.Rewrite("$_.channelCount = $1", channelCount)
}

// ChannelCountMode prop
// js:"channelCountMode"
func (*PannerNode) ChannelCountMode() (channelCountMode *channelcountmode.ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode")
	return channelCountMode
}

// SetChannelCountMode prop
// js:"channelCountMode"
func (*PannerNode) SetChannelCountMode(channelCountMode *channelcountmode.ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode = $1", channelCountMode)
}

// ChannelInterpretation prop
// js:"channelInterpretation"
func (*PannerNode) ChannelInterpretation() (channelInterpretation *channelinterpretation.ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation")
	return channelInterpretation
}

// SetChannelInterpretation prop
// js:"channelInterpretation"
func (*PannerNode) SetChannelInterpretation(channelInterpretation *channelinterpretation.ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation = $1", channelInterpretation)
}

// Context prop
// js:"context"
func (*PannerNode) Context() (context AudioContext) {
	macro.Rewrite("$_.context")
	return context
}

// NumberOfInputs prop
// js:"numberOfInputs"
func (*PannerNode) NumberOfInputs() (numberOfInputs uint) {
	macro.Rewrite("$_.numberOfInputs")
	return numberOfInputs
}

// NumberOfOutputs prop
// js:"numberOfOutputs"
func (*PannerNode) NumberOfOutputs() (numberOfOutputs uint) {
	macro.Rewrite("$_.numberOfOutputs")
	return numberOfOutputs
}
