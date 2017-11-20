package audionode

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/audiocontext"
	"github.com/matthewmueller/golly/internal/gendom/dom/channelcountmode"
	"github.com/matthewmueller/golly/internal/gendom/dom/channelinterpretation"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/js"
)

type AudioNode struct {
	*eventtarget.EventTarget
}

func (*AudioNode) Connect(destination *AudioNode, output uint, input uint) (a *AudioNode) {
	js.Rewrite("$<.connect($1, $2, $3)", destination, output, input)
	return a
}

func (*AudioNode) Disconnect() {
	js.Rewrite("$<.disconnect()")
}

func (*AudioNode) GetChannelCount() (channelCount uint) {
	js.Rewrite("$<.channelCount")
	return channelCount
}

func (*AudioNode) SetChannelCount(channelCount uint) {
	js.Rewrite("$<.channelCount = $1", channelCount)
}

func (*AudioNode) GetChannelCountMode() (channelCountMode *channelcountmode.ChannelCountMode) {
	js.Rewrite("$<.channelCountMode")
	return channelCountMode
}

func (*AudioNode) SetChannelCountMode(channelCountMode *channelcountmode.ChannelCountMode) {
	js.Rewrite("$<.channelCountMode = $1", channelCountMode)
}

func (*AudioNode) GetChannelInterpretation() (channelInterpretation *channelinterpretation.ChannelInterpretation) {
	js.Rewrite("$<.channelInterpretation")
	return channelInterpretation
}

func (*AudioNode) SetChannelInterpretation(channelInterpretation *channelinterpretation.ChannelInterpretation) {
	js.Rewrite("$<.channelInterpretation = $1", channelInterpretation)
}

func (*AudioNode) GetContext() (context *audiocontext.AudioContext) {
	js.Rewrite("$<.context")
	return context
}

func (*AudioNode) GetNumberOfInputs() (numberOfInputs uint) {
	js.Rewrite("$<.numberOfInputs")
	return numberOfInputs
}

func (*AudioNode) GetNumberOfOutputs() (numberOfOutputs uint) {
	js.Rewrite("$<.numberOfOutputs")
	return numberOfOutputs
}
