package audionode

import (
	"github.com/matthewmueller/golly/dom/channelcountmode"
	"github.com/matthewmueller/golly/dom/channelinterpretation"
	"github.com/matthewmueller/golly/dom/window"
)

// AudioNode interface
// js:"AudioNode"
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

	// SetChannelCount prop
	// js:"channelCount"
	SetChannelCount(channelCount uint)

	// ChannelCountMode prop
	// js:"channelCountMode"
	ChannelCountMode() (channelCountMode *channelcountmode.ChannelCountMode)

	// SetChannelCountMode prop
	// js:"channelCountMode"
	SetChannelCountMode(channelCountMode *channelcountmode.ChannelCountMode)

	// ChannelInterpretation prop
	// js:"channelInterpretation"
	ChannelInterpretation() (channelInterpretation *channelinterpretation.ChannelInterpretation)

	// SetChannelInterpretation prop
	// js:"channelInterpretation"
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
