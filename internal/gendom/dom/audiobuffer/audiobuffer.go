package audiobuffer

import "github.com/matthewmueller/golly/js"

type AudioBuffer struct {
}

func (*AudioBuffer) CopyFromChannel(destination []float32, channelNumber int, startInChannel uint) {
	js.Rewrite("$<.copyFromChannel($1, $2, $3)", destination, channelNumber, startInChannel)
}

func (*AudioBuffer) CopyToChannel(source []float32, channelNumber int, startInChannel uint) {
	js.Rewrite("$<.copyToChannel($1, $2, $3)", source, channelNumber, startInChannel)
}

func (*AudioBuffer) GetChannelData(channel uint) (f []float32) {
	js.Rewrite("$<.getChannelData($1)", channel)
	return f
}

func (*AudioBuffer) GetDuration() (duration float32) {
	js.Rewrite("$<.duration")
	return duration
}

func (*AudioBuffer) GetLength() (length int) {
	js.Rewrite("$<.length")
	return length
}

func (*AudioBuffer) GetNumberOfChannels() (numberOfChannels int) {
	js.Rewrite("$<.numberOfChannels")
	return numberOfChannels
}

func (*AudioBuffer) GetSampleRate() (sampleRate float32) {
	js.Rewrite("$<.sampleRate")
	return sampleRate
}
