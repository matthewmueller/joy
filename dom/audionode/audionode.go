package audionode

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// js:"AudioNode,omit"
type AudioNode interface {
	window.EventTarget

	// Connect
	// js:"connect,rewrite=connect"
	Connect(destination AudioNode, output *uint, input *uint) (a AudioNode)

	// Disconnect
	// js:"disconnect,rewrite=disconnect"
	Disconnect()

	// ChannelCount prop
	// js:"channelCount,rewrite=channelcount"
	ChannelCount() (channelCount uint)

	// ChannelCount prop
	// js:"setchannelCount,rewrite=setchannelcount"
	SetChannelCount(channelCount uint)

	// ChannelCountMode prop
	// js:"channelCountMode,rewrite=channelcountmode"
	ChannelCountMode() (channelCountMode *channelcountmode.ChannelCountMode)

	// ChannelCountMode prop
	// js:"setchannelCountMode,rewrite=setchannelcountmode"
	SetChannelCountMode(channelCountMode *channelcountmode.ChannelCountMode)

	// ChannelInterpretation prop
	// js:"channelInterpretation,rewrite=channelinterpretation"
	ChannelInterpretation() (channelInterpretation *channelinterpretation.ChannelInterpretation)

	// ChannelInterpretation prop
	// js:"setchannelInterpretation,rewrite=setchannelinterpretation"
	SetChannelInterpretation(channelInterpretation *channelinterpretation.ChannelInterpretation)

	// Context prop
	// js:"context,rewrite=context"
	Context() (context AudioContext)

	// NumberOfInputs prop
	// js:"numberOfInputs,rewrite=numberofinputs"
	NumberOfInputs() (numberOfInputs uint)

	// NumberOfOutputs prop
	// js:"numberOfOutputs,rewrite=numberofoutputs"
	NumberOfOutputs() (numberOfOutputs uint)
}

// connect fn
func connect(destination AudioNode, output *uint, input *uint) (a AudioNode) {
	js.Rewrite("$<.connect($1, $2, $3)", destination, output, input)
	return a
}

// disconnect fn
func disconnect() {
	js.Rewrite("$<.disconnect()")
}

// channelcount prop
func channelcount() (channelCount uint) {
	js.Rewrite("$<.channelCount")
	return channelCount
}

// setchannelcount prop
func setchannelcount(channelCount uint) {
	js.Rewrite("$<.channelCount = channelCount")
}

// channelcountmode prop
func channelcountmode() (channelCountMode *channelcountmode.ChannelCountMode) {
	js.Rewrite("$<.channelCountMode")
	return channelCountMode
}

// setchannelcountmode prop
func setchannelcountmode(channelCountMode *channelcountmode.ChannelCountMode) {
	js.Rewrite("$<.channelCountMode = channelCountMode")
}

// channelinterpretation prop
func channelinterpretation() (channelInterpretation *channelinterpretation.ChannelInterpretation) {
	js.Rewrite("$<.channelInterpretation")
	return channelInterpretation
}

// setchannelinterpretation prop
func setchannelinterpretation(channelInterpretation *channelinterpretation.ChannelInterpretation) {
	js.Rewrite("$<.channelInterpretation = channelInterpretation")
}

// context prop
func context() (context AudioContext) {
	js.Rewrite("$<.context")
	return context
}

// numberofinputs prop
func numberofinputs() (numberOfInputs uint) {
	js.Rewrite("$<.numberOfInputs")
	return numberOfInputs
}

// numberofoutputs prop
func numberofoutputs() (numberOfOutputs uint) {
	js.Rewrite("$<.numberOfOutputs")
	return numberOfOutputs
}
