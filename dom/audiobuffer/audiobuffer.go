package audiobuffer

import "github.com/matthewmueller/golly/js"

// AudioBuffer struct
// js:"AudioBuffer,omit"
type AudioBuffer struct {
}

// CopyFromChannel fn
func (*AudioBuffer) CopyFromChannel(destination []float32, channelNumber int, startInChannel *uint) {
	js.Rewrite("$<.copyFromChannel($1, $2, $3)", destination, channelNumber, startInChannel)
}

// CopyToChannel fn
func (*AudioBuffer) CopyToChannel(source []float32, channelNumber int, startInChannel *uint) {
	js.Rewrite("$<.copyToChannel($1, $2, $3)", source, channelNumber, startInChannel)
}

// GetChannelData fn
func (*AudioBuffer) GetChannelData(channel uint) (f []float32) {
	js.Rewrite("$<.getChannelData($1)", channel)
	return f
}

// Duration prop
func (*AudioBuffer) Duration() (duration float32) {
	js.Rewrite("$<.duration")
	return duration
}

// Length prop
func (*AudioBuffer) Length() (length int) {
	js.Rewrite("$<.length")
	return length
}

// NumberOfChannels prop
func (*AudioBuffer) NumberOfChannels() (numberOfChannels int) {
	js.Rewrite("$<.numberOfChannels")
	return numberOfChannels
}

// SampleRate prop
func (*AudioBuffer) SampleRate() (sampleRate float32) {
	js.Rewrite("$<.sampleRate")
	return sampleRate
}
