package audionode

import "github.com/matthewmueller/golly/js"

// AnalyserNode struct
// js:"AnalyserNode,omit"
type AnalyserNode struct {
	AudioNode
}

// GetByteFrequencyData fn
func (*AnalyserNode) GetByteFrequencyData(array []uint8) {
	js.Rewrite("$<.getByteFrequencyData($1)", array)
}

// GetByteTimeDomainData fn
func (*AnalyserNode) GetByteTimeDomainData(array []uint8) {
	js.Rewrite("$<.getByteTimeDomainData($1)", array)
}

// GetFloatFrequencyData fn
func (*AnalyserNode) GetFloatFrequencyData(array []float32) {
	js.Rewrite("$<.getFloatFrequencyData($1)", array)
}

// GetFloatTimeDomainData fn
func (*AnalyserNode) GetFloatTimeDomainData(array []float32) {
	js.Rewrite("$<.getFloatTimeDomainData($1)", array)
}

// FftSize prop
func (*AnalyserNode) FftSize() (fftSize uint) {
	js.Rewrite("$<.fftSize")
	return fftSize
}

// FftSize prop
func (*AnalyserNode) SetFftSize(fftSize uint) {
	js.Rewrite("$<.fftSize = $1", fftSize)
}

// FrequencyBinCount prop
func (*AnalyserNode) FrequencyBinCount() (frequencyBinCount uint) {
	js.Rewrite("$<.frequencyBinCount")
	return frequencyBinCount
}

// MaxDecibels prop
func (*AnalyserNode) MaxDecibels() (maxDecibels float32) {
	js.Rewrite("$<.maxDecibels")
	return maxDecibels
}

// MaxDecibels prop
func (*AnalyserNode) SetMaxDecibels(maxDecibels float32) {
	js.Rewrite("$<.maxDecibels = $1", maxDecibels)
}

// MinDecibels prop
func (*AnalyserNode) MinDecibels() (minDecibels float32) {
	js.Rewrite("$<.minDecibels")
	return minDecibels
}

// MinDecibels prop
func (*AnalyserNode) SetMinDecibels(minDecibels float32) {
	js.Rewrite("$<.minDecibels = $1", minDecibels)
}

// SmoothingTimeConstant prop
func (*AnalyserNode) SmoothingTimeConstant() (smoothingTimeConstant float32) {
	js.Rewrite("$<.smoothingTimeConstant")
	return smoothingTimeConstant
}

// SmoothingTimeConstant prop
func (*AnalyserNode) SetSmoothingTimeConstant(smoothingTimeConstant float32) {
	js.Rewrite("$<.smoothingTimeConstant = $1", smoothingTimeConstant)
}
