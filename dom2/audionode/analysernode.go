package audionode

import "github.com/matthewmueller/golly/js"

// js:"AnalyserNode,omit"
type AnalyserNode struct {
	AudioNode
}

// GetByteFrequencyData
func (*AnalyserNode) GetByteFrequencyData(array []uint8) {
	js.Rewrite("$<.GetByteFrequencyData($1)", array)
}

// GetByteTimeDomainData
func (*AnalyserNode) GetByteTimeDomainData(array []uint8) {
	js.Rewrite("$<.GetByteTimeDomainData($1)", array)
}

// GetFloatFrequencyData
func (*AnalyserNode) GetFloatFrequencyData(array []float32) {
	js.Rewrite("$<.GetFloatFrequencyData($1)", array)
}

// GetFloatTimeDomainData
func (*AnalyserNode) GetFloatTimeDomainData(array []float32) {
	js.Rewrite("$<.GetFloatTimeDomainData($1)", array)
}

// FftSize
func (*AnalyserNode) FftSize() (fftSize uint) {
	js.Rewrite("$<.FftSize")
	return fftSize
}

// FftSize
func (*AnalyserNode) SetFftSize(fftSize uint) {
	js.Rewrite("$<.FftSize = $1", fftSize)
}

// FrequencyBinCount
func (*AnalyserNode) FrequencyBinCount() (frequencyBinCount uint) {
	js.Rewrite("$<.FrequencyBinCount")
	return frequencyBinCount
}

// MaxDecibels
func (*AnalyserNode) MaxDecibels() (maxDecibels float32) {
	js.Rewrite("$<.MaxDecibels")
	return maxDecibels
}

// MaxDecibels
func (*AnalyserNode) SetMaxDecibels(maxDecibels float32) {
	js.Rewrite("$<.MaxDecibels = $1", maxDecibels)
}

// MinDecibels
func (*AnalyserNode) MinDecibels() (minDecibels float32) {
	js.Rewrite("$<.MinDecibels")
	return minDecibels
}

// MinDecibels
func (*AnalyserNode) SetMinDecibels(minDecibels float32) {
	js.Rewrite("$<.MinDecibels = $1", minDecibels)
}

// SmoothingTimeConstant
func (*AnalyserNode) SmoothingTimeConstant() (smoothingTimeConstant float32) {
	js.Rewrite("$<.SmoothingTimeConstant")
	return smoothingTimeConstant
}

// SmoothingTimeConstant
func (*AnalyserNode) SetSmoothingTimeConstant(smoothingTimeConstant float32) {
	js.Rewrite("$<.SmoothingTimeConstant = $1", smoothingTimeConstant)
}
