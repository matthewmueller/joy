package audiobuffer

import "github.com/matthewmueller/golly/js"

// AudioBuffer struct
// js:"AudioBuffer,omit"
type AudioBuffer struct {
}

// CopyFromChannel
func (*AudioBuffer) CopyFromChannel(destination []float32, channelNumber int, startInChannel *uint) {
	js.Rewrite("$<.CopyFromChannel($1, $2, $3)", destination, channelNumber, startInChannel)
}

// CopyToChannel
func (*AudioBuffer) CopyToChannel(source []float32, channelNumber int, startInChannel *uint) {
	js.Rewrite("$<.CopyToChannel($1, $2, $3)", source, channelNumber, startInChannel)
}

// GetChannelData
func (*AudioBuffer) GetChannelData(channel uint) (f []float32) {
	js.Rewrite("$<.GetChannelData($1)", channel)
	return f
}

// Duration
func (*AudioBuffer) Duration() (duration float32) {
	js.Rewrite("$<.Duration")
	return duration
}

// Length
func (*AudioBuffer) Length() (length int) {
	js.Rewrite("$<.Length")
	return length
}

// NumberOfChannels
func (*AudioBuffer) NumberOfChannels() (numberOfChannels int) {
	js.Rewrite("$<.NumberOfChannels")
	return numberOfChannels
}

// SampleRate
func (*AudioBuffer) SampleRate() (sampleRate float32) {
	js.Rewrite("$<.SampleRate")
	return sampleRate
}
