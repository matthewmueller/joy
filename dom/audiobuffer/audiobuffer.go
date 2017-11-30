package audiobuffer

import "github.com/matthewmueller/golly/js"

// AudioBuffer struct
// js:"AudioBuffer,omit"
type AudioBuffer struct {
}

// CopyFromChannel fn
// js:"copyFromChannel"
func (*AudioBuffer) CopyFromChannel(destination []float32, channelNumber int, startInChannel *uint) {
	js.Rewrite("$_.copyFromChannel($1, $2, $3)", destination, channelNumber, startInChannel)
}

// CopyToChannel fn
// js:"copyToChannel"
func (*AudioBuffer) CopyToChannel(source []float32, channelNumber int, startInChannel *uint) {
	js.Rewrite("$_.copyToChannel($1, $2, $3)", source, channelNumber, startInChannel)
}

// GetChannelData fn
// js:"getChannelData"
func (*AudioBuffer) GetChannelData(channel uint) (f []float32) {
	js.Rewrite("$_.getChannelData($1)", channel)
	return f
}

// Duration prop
// js:"duration"
func (*AudioBuffer) Duration() (duration float32) {
	js.Rewrite("$_.duration")
	return duration
}

// Length prop
// js:"length"
func (*AudioBuffer) Length() (length int) {
	js.Rewrite("$_.length")
	return length
}

// NumberOfChannels prop
// js:"numberOfChannels"
func (*AudioBuffer) NumberOfChannels() (numberOfChannels int) {
	js.Rewrite("$_.numberOfChannels")
	return numberOfChannels
}

// SampleRate prop
// js:"sampleRate"
func (*AudioBuffer) SampleRate() (sampleRate float32) {
	js.Rewrite("$_.sampleRate")
	return sampleRate
}
