package audiobuffer

import "github.com/matthewmueller/joy/macro"

// AudioBuffer struct
// js:"AudioBuffer,omit"
type AudioBuffer struct {
}

// CopyFromChannel fn
// js:"copyFromChannel"
func (*AudioBuffer) CopyFromChannel(destination []float32, channelNumber int, startInChannel *uint) {
	macro.Rewrite("$_.copyFromChannel($1, $2, $3)", destination, channelNumber, startInChannel)
}

// CopyToChannel fn
// js:"copyToChannel"
func (*AudioBuffer) CopyToChannel(source []float32, channelNumber int, startInChannel *uint) {
	macro.Rewrite("$_.copyToChannel($1, $2, $3)", source, channelNumber, startInChannel)
}

// GetChannelData fn
// js:"getChannelData"
func (*AudioBuffer) GetChannelData(channel uint) (f []float32) {
	macro.Rewrite("$_.getChannelData($1)", channel)
	return f
}

// Duration prop
// js:"duration"
func (*AudioBuffer) Duration() (duration float32) {
	macro.Rewrite("$_.duration")
	return duration
}

// Length prop
// js:"length"
func (*AudioBuffer) Length() (length int) {
	macro.Rewrite("$_.length")
	return length
}

// NumberOfChannels prop
// js:"numberOfChannels"
func (*AudioBuffer) NumberOfChannels() (numberOfChannels int) {
	macro.Rewrite("$_.numberOfChannels")
	return numberOfChannels
}

// SampleRate prop
// js:"sampleRate"
func (*AudioBuffer) SampleRate() (sampleRate float32) {
	macro.Rewrite("$_.sampleRate")
	return sampleRate
}
