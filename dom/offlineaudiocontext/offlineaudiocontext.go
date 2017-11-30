package offlineaudiocontext

import (
	"github.com/matthewmueller/golly/dom/audiobuffer"
	"github.com/matthewmueller/golly/dom/audiocontextstate"
	"github.com/matthewmueller/golly/dom/audiolistener"
	"github.com/matthewmueller/golly/dom/audionode"
	"github.com/matthewmueller/golly/dom/htmlmediaelement"
	"github.com/matthewmueller/golly/dom/offlineaudiocompletionevent"
	"github.com/matthewmueller/golly/dom/periodicwave"
	"github.com/matthewmueller/golly/dom/periodicwaveconstraints"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

var _ audionode.AudioContext = (*OfflineAudioContext)(nil)
var _ window.EventTarget = (*OfflineAudioContext)(nil)

// OfflineAudioContext struct
// js:"OfflineAudioContext,omit"
type OfflineAudioContext struct {
}

// StartRendering fn
// js:"startRendering"
func (*OfflineAudioContext) StartRendering() (a *audiobuffer.AudioBuffer) {
	js.Rewrite("await $_.startRendering()")
	return a
}

// Suspend fn
// js:"suspend"
func (*OfflineAudioContext) Suspend(suspendTime float32) {
	js.Rewrite("await $_.suspend($1)", suspendTime)
}

// Close fn
// js:"close"
func (*OfflineAudioContext) Close() {
	js.Rewrite("await $_.close()")
}

// CreateAnalyser fn
// js:"createAnalyser"
func (*OfflineAudioContext) CreateAnalyser() (a *audionode.AnalyserNode) {
	js.Rewrite("$_.createAnalyser()")
	return a
}

// CreateBiquadFilter fn
// js:"createBiquadFilter"
func (*OfflineAudioContext) CreateBiquadFilter() (a *audionode.BiquadFilterNode) {
	js.Rewrite("$_.createBiquadFilter()")
	return a
}

// CreateBuffer fn
// js:"createBuffer"
func (*OfflineAudioContext) CreateBuffer(numberOfChannels uint, length uint, sampleRate float32) (a *audiobuffer.AudioBuffer) {
	js.Rewrite("$_.createBuffer($1, $2, $3)", numberOfChannels, length, sampleRate)
	return a
}

// CreateBufferSource fn
// js:"createBufferSource"
func (*OfflineAudioContext) CreateBufferSource() (a *audionode.AudioBufferSourceNode) {
	js.Rewrite("$_.createBufferSource()")
	return a
}

// CreateChannelMerger fn
// js:"createChannelMerger"
func (*OfflineAudioContext) CreateChannelMerger(numberOfInputs *uint) (a *audionode.ChannelMergerNode) {
	js.Rewrite("$_.createChannelMerger($1)", numberOfInputs)
	return a
}

// CreateChannelSplitter fn
// js:"createChannelSplitter"
func (*OfflineAudioContext) CreateChannelSplitter(numberOfOutputs *uint) (a *audionode.ChannelSplitterNode) {
	js.Rewrite("$_.createChannelSplitter($1)", numberOfOutputs)
	return a
}

// CreateConvolver fn
// js:"createConvolver"
func (*OfflineAudioContext) CreateConvolver() (a *audionode.ConvolverNode) {
	js.Rewrite("$_.createConvolver()")
	return a
}

// CreateDelay fn
// js:"createDelay"
func (*OfflineAudioContext) CreateDelay(maxDelayTime *float32) (a *audionode.DelayNode) {
	js.Rewrite("$_.createDelay($1)", maxDelayTime)
	return a
}

// CreateDynamicsCompressor fn
// js:"createDynamicsCompressor"
func (*OfflineAudioContext) CreateDynamicsCompressor() (a *audionode.DynamicsCompressorNode) {
	js.Rewrite("$_.createDynamicsCompressor()")
	return a
}

// CreateGain fn
// js:"createGain"
func (*OfflineAudioContext) CreateGain() (a *audionode.GainNode) {
	js.Rewrite("$_.createGain()")
	return a
}

// CreateIIRFilter fn
// js:"createIIRFilter"
func (*OfflineAudioContext) CreateIIRFilter(feedforward []float32, feedback []float32) (a *audionode.IIRFilterNode) {
	js.Rewrite("$_.createIIRFilter($1, $2)", feedforward, feedback)
	return a
}

// CreateMediaElementSource fn
// js:"createMediaElementSource"
func (*OfflineAudioContext) CreateMediaElementSource(mediaElement htmlmediaelement.HTMLMediaElement) (a *audionode.MediaElementAudioSourceNode) {
	js.Rewrite("$_.createMediaElementSource($1)", mediaElement)
	return a
}

// CreateMediaStreamSource fn
// js:"createMediaStreamSource"
func (*OfflineAudioContext) CreateMediaStreamSource(mediaStream *window.MediaStream) (a *audionode.MediaStreamAudioSourceNode) {
	js.Rewrite("$_.createMediaStreamSource($1)", mediaStream)
	return a
}

// CreateOscillator fn
// js:"createOscillator"
func (*OfflineAudioContext) CreateOscillator() (a *audionode.OscillatorNode) {
	js.Rewrite("$_.createOscillator()")
	return a
}

// CreatePanner fn
// js:"createPanner"
func (*OfflineAudioContext) CreatePanner() (a *audionode.PannerNode) {
	js.Rewrite("$_.createPanner()")
	return a
}

// CreatePeriodicWave fn
// js:"createPeriodicWave"
func (*OfflineAudioContext) CreatePeriodicWave(rl []float32, img []float32, constraints *periodicwaveconstraints.PeriodicWaveConstraints) (p *periodicwave.PeriodicWave) {
	js.Rewrite("$_.createPeriodicWave($1, $2, $3)", rl, img, constraints)
	return p
}

// CreateScriptProcessor fn
// js:"createScriptProcessor"
func (*OfflineAudioContext) CreateScriptProcessor(bufferSize *uint, numberOfInputChannels *uint, numberOfOutputChannels *uint) (a *audionode.ScriptProcessorNode) {
	js.Rewrite("$_.createScriptProcessor($1, $2, $3)", bufferSize, numberOfInputChannels, numberOfOutputChannels)
	return a
}

// CreateStereoPanner fn
// js:"createStereoPanner"
func (*OfflineAudioContext) CreateStereoPanner() (a *audionode.StereoPannerNode) {
	js.Rewrite("$_.createStereoPanner()")
	return a
}

// CreateWaveShaper fn
// js:"createWaveShaper"
func (*OfflineAudioContext) CreateWaveShaper() (a *audionode.WaveShaperNode) {
	js.Rewrite("$_.createWaveShaper()")
	return a
}

// DecodeAudioData fn
// js:"decodeAudioData"
func (*OfflineAudioContext) DecodeAudioData(audioData []byte, successCallback *func(decodedData *audiobuffer.AudioBuffer), errorCallback *func()) (a *audiobuffer.AudioBuffer) {
	js.Rewrite("await $_.decodeAudioData($1, $2, $3)", audioData, successCallback, errorCallback)
	return a
}

// Resume fn
// js:"resume"
func (*OfflineAudioContext) Resume() {
	js.Rewrite("await $_.resume()")
}

// AddEventListener fn
// js:"addEventListener"
func (*OfflineAudioContext) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*OfflineAudioContext) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*OfflineAudioContext) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Length prop
// js:"length"
func (*OfflineAudioContext) Length() (length uint) {
	js.Rewrite("$_.length")
	return length
}

// Oncomplete prop
// js:"oncomplete"
func (*OfflineAudioContext) Oncomplete() (oncomplete func(*offlineaudiocompletionevent.OfflineAudioCompletionEvent)) {
	js.Rewrite("$_.oncomplete")
	return oncomplete
}

// SetOncomplete prop
// js:"oncomplete"
func (*OfflineAudioContext) SetOncomplete(oncomplete func(*offlineaudiocompletionevent.OfflineAudioCompletionEvent)) {
	js.Rewrite("$_.oncomplete = $1", oncomplete)
}

// CurrentTime prop
// js:"currentTime"
func (*OfflineAudioContext) CurrentTime() (currentTime float32) {
	js.Rewrite("$_.currentTime")
	return currentTime
}

// Destination prop
// js:"destination"
func (*OfflineAudioContext) Destination() (destination *audionode.AudioDestinationNode) {
	js.Rewrite("$_.destination")
	return destination
}

// Listener prop
// js:"listener"
func (*OfflineAudioContext) Listener() (listener *audiolistener.AudioListener) {
	js.Rewrite("$_.listener")
	return listener
}

// Onstatechange prop
// js:"onstatechange"
func (*OfflineAudioContext) Onstatechange() (onstatechange func(window.Event)) {
	js.Rewrite("$_.onstatechange")
	return onstatechange
}

// SetOnstatechange prop
// js:"onstatechange"
func (*OfflineAudioContext) SetOnstatechange(onstatechange func(window.Event)) {
	js.Rewrite("$_.onstatechange = $1", onstatechange)
}

// SampleRate prop
// js:"sampleRate"
func (*OfflineAudioContext) SampleRate() (sampleRate float32) {
	js.Rewrite("$_.sampleRate")
	return sampleRate
}

// State prop
// js:"state"
func (*OfflineAudioContext) State() (state *audiocontextstate.AudioContextState) {
	js.Rewrite("$_.state")
	return state
}
