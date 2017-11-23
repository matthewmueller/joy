package audionode

import (
	"github.com/matthewmueller/golly/dom2/channelcountmode"
	"github.com/matthewmueller/golly/dom2/channelinterpretation"
	"github.com/matthewmueller/golly/dom2/window"
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
	SetChannelCount(channelCount uint)

	// ChannelCountMode prop
	// js:"channelCountMode"
	ChannelCountMode() (channelCountMode *channelcountmode.ChannelCountMode)

	// ChannelCountMode prop
	SetChannelCountMode(channelCountMode *channelcountmode.ChannelCountMode)

	// ChannelInterpretation prop
	// js:"channelInterpretation"
	ChannelInterpretation() (channelInterpretation *channelinterpretation.ChannelInterpretation)

	// ChannelInterpretation prop
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
