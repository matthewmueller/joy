package analysernode

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/audionode"
	"github.com/matthewmueller/golly/js"
)

type AnalyserNode struct {
	*audionode.AudioNode
}

func (*AnalyserNode) GetByteFrequencyData(array []uint8) {
	js.Rewrite("$<.getByteFrequencyData($1)", array)
}

func (*AnalyserNode) GetByteTimeDomainData(array []uint8) {
	js.Rewrite("$<.getByteTimeDomainData($1)", array)
}

func (*AnalyserNode) GetFloatFrequencyData(array []float32) {
	js.Rewrite("$<.getFloatFrequencyData($1)", array)
}

func (*AnalyserNode) GetFloatTimeDomainData(array []float32) {
	js.Rewrite("$<.getFloatTimeDomainData($1)", array)
}

func (*AnalyserNode) GetFftSize() (fftSize uint) {
	js.Rewrite("$<.fftSize")
	return fftSize
}

func (*AnalyserNode) SetFftSize(fftSize uint) {
	js.Rewrite("$<.fftSize = $1", fftSize)
}

func (*AnalyserNode) GetFrequencyBinCount() (frequencyBinCount uint) {
	js.Rewrite("$<.frequencyBinCount")
	return frequencyBinCount
}

func (*AnalyserNode) GetMaxDecibels() (maxDecibels float32) {
	js.Rewrite("$<.maxDecibels")
	return maxDecibels
}

func (*AnalyserNode) SetMaxDecibels(maxDecibels float32) {
	js.Rewrite("$<.maxDecibels = $1", maxDecibels)
}

func (*AnalyserNode) GetMinDecibels() (minDecibels float32) {
	js.Rewrite("$<.minDecibels")
	return minDecibels
}

func (*AnalyserNode) SetMinDecibels(minDecibels float32) {
	js.Rewrite("$<.minDecibels = $1", minDecibels)
}

func (*AnalyserNode) GetSmoothingTimeConstant() (smoothingTimeConstant float32) {
	js.Rewrite("$<.smoothingTimeConstant")
	return smoothingTimeConstant
}

func (*AnalyserNode) SetSmoothingTimeConstant(smoothingTimeConstant float32) {
	js.Rewrite("$<.smoothingTimeConstant = $1", smoothingTimeConstant)
}
