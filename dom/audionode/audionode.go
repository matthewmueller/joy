package audionode

import (
	"github.com/matthewmueller/golly/dom/channelcountmode"
	"github.com/matthewmueller/golly/dom/channelinterpretation"
	"github.com/matthewmueller/golly/dom/window"
)

// js:"AudioNode,omit"
type AudioNode interface {
	window.EventTarget

	// Connect
	// js:"connect"
	Connect(destination AudioNode, output *uint, input *uint) (a AudioNode)

	// Disconnect
	// js:"disconnect"
	Disconnect()

	// ChannelCount prop
	// js:"channelCount"
	ChannelCount() (channelCount uint)

	// ChannelCount prop
	// js:"setchannelCount"
	SetChannelCount(channelCount uint)

	// ChannelCountMode prop
	// js:"channelCountMode"
	ChannelCountMode() (channelCountMode *channelcountmode.ChannelCountMode)

	// ChannelCountMode prop
	// js:"setchannelCountMode"
	SetChannelCountMode(channelCountMode *channelcountmode.ChannelCountMode)

	// ChannelInterpretation prop
	// js:"channelInterpretation"
	ChannelInterpretation() (channelInterpretation *channelinterpretation.ChannelInterpretation)

	// ChannelInterpretation prop
	// js:"setchannelInterpretation"
	SetChannelInterpretation(channelInterpretation *channelinterpretation.ChannelInterpretation)

	// Context prop
	// js:"context"
	Context() (context AudioContext)

	// NumberOfInputs prop
	// js:"numberOfInputs"
	NumberOfInputs() (numberOfInputs uint)

	// NumberOfOutputs prop
	// js:"numberOfOutputs"
	NumberOfOutputs() (numberOfOutputs uint)
}
