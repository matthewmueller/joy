package audionode

import (
	"github.com/matthewmueller/golly/dom2/channelcountmode"
	"github.com/matthewmueller/golly/dom2/window"
)

// js:"AudioNode,omit"
type AudioNode interface {
	window.EventTarget

	// Connect
	Connect(destination AudioNode, output *uint, input *uint) (a AudioNode)

	// Disconnect
	Disconnect()

	// ChannelCount
	ChannelCount() (channelCount uint)

	// ChannelCount
	SetChannelCount(channelCount uint)

	// ChannelCountMode
	ChannelCountMode() (channelCountMode *channelcountmode.ChannelCountMode)

	// ChannelCountMode
	SetChannelCountMode(channelCountMode *channelcountmode.ChannelCountMode)

	// ChannelInterpretation
	ChannelInterpretation() (channelInterpretation *channelinterpretation.ChannelInterpretation)

	// ChannelInterpretation
	SetChannelInterpretation(channelInterpretation *channelinterpretation.ChannelInterpretation)

	// Context
	Context() (context AudioContext)

	// NumberOfInputs
	NumberOfInputs() (numberOfInputs uint)

	// NumberOfOutputs
	NumberOfOutputs() (numberOfOutputs uint)
}
