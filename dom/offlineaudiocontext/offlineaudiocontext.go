package offlineaudiocontext

import (
	"github.com/matthewmueller/joy/dom/audiobuffer"
	"github.com/matthewmueller/joy/dom/audiocontextstate"
	"github.com/matthewmueller/joy/dom/audiolistener"
	"github.com/matthewmueller/joy/dom/audionode"
	"github.com/matthewmueller/joy/dom/htmlmediaelement"
	"github.com/matthewmueller/joy/dom/offlineaudiocompletionevent"
	"github.com/matthewmueller/joy/dom/periodicwave"
	"github.com/matthewmueller/joy/dom/periodicwaveconstraints"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ audionode.AudioContext = (*OfflineAudioContext)(nil)
var _ window.EventTarget = (*OfflineAudioContext)(nil)

// New fn
func New(numberofchannels uint, length uint, samplerate float32) *OfflineAudioContext {
	macro.Rewrite("OfflineAudioContext")
	return &OfflineAudioContext{}
}

// OfflineAudioContext struct
// js:"OfflineAudioContext,omit"
type OfflineAudioContext struct {
}

// StartRendering fn
// js:"startRendering"
func (*OfflineAudioContext) StartRendering() (a *audiobuffer.AudioBuffer) {
	macro.Rewrite("await $_.startRendering()")
	return a
}

// Suspend fn
// js:"suspend"
func (*OfflineAudioContext) Suspend(suspendTime float32) {
	macro.Rewrite("await $_.suspend($1)", suspendTime)
}

// Close fn
// js:"close"
func (*OfflineAudioContext) Close() {
	macro.Rewrite("await $_.close()")
}

// CreateAnalyser fn
// js:"createAnalyser"
func (*OfflineAudioContext) CreateAnalyser() (a *audionode.AnalyserNode) {
	macro.Rewrite("$_.createAnalyser()")
	return a
}

// CreateBiquadFilter fn
// js:"createBiquadFilter"
func (*OfflineAudioContext) CreateBiquadFilter() (a *audionode.BiquadFilterNode) {
	macro.Rewrite("$_.createBiquadFilter()")
	return a
}

// CreateBuffer fn
// js:"createBuffer"
func (*OfflineAudioContext) CreateBuffer(numberOfChannels uint, length uint, sampleRate float32) (a *audiobuffer.AudioBuffer) {
	macro.Rewrite("$_.createBuffer($1, $2, $3)", numberOfChannels, length, sampleRate)
	return a
}

// CreateBufferSource fn
// js:"createBufferSource"
func (*OfflineAudioContext) CreateBufferSource() (a *audionode.AudioBufferSourceNode) {
	macro.Rewrite("$_.createBufferSource()")
	return a
}

// CreateChannelMerger fn
// js:"createChannelMerger"
func (*OfflineAudioContext) CreateChannelMerger(numberOfInputs *uint) (a *audionode.ChannelMergerNode) {
	macro.Rewrite("$_.createChannelMerger($1)", numberOfInputs)
	return a
}

// CreateChannelSplitter fn
// js:"createChannelSplitter"
func (*OfflineAudioContext) CreateChannelSplitter(numberOfOutputs *uint) (a *audionode.ChannelSplitterNode) {
	macro.Rewrite("$_.createChannelSplitter($1)", numberOfOutputs)
	return a
}

// CreateConvolver fn
// js:"createConvolver"
func (*OfflineAudioContext) CreateConvolver() (a *audionode.ConvolverNode) {
	macro.Rewrite("$_.createConvolver()")
	return a
}

// CreateDelay fn
// js:"createDelay"
func (*OfflineAudioContext) CreateDelay(maxDelayTime *float32) (a *audionode.DelayNode) {
	macro.Rewrite("$_.createDelay($1)", maxDelayTime)
	return a
}

// CreateDynamicsCompressor fn
// js:"createDynamicsCompressor"
func (*OfflineAudioContext) CreateDynamicsCompressor() (a *audionode.DynamicsCompressorNode) {
	macro.Rewrite("$_.createDynamicsCompressor()")
	return a
}

// CreateGain fn
// js:"createGain"
func (*OfflineAudioContext) CreateGain() (a *audionode.GainNode) {
	macro.Rewrite("$_.createGain()")
	return a
}

// CreateIIRFilter fn
// js:"createIIRFilter"
func (*OfflineAudioContext) CreateIIRFilter(feedforward []float32, feedback []float32) (a *audionode.IIRFilterNode) {
	macro.Rewrite("$_.createIIRFilter($1, $2)", feedforward, feedback)
	return a
}

// CreateMediaElementSource fn
// js:"createMediaElementSource"
func (*OfflineAudioContext) CreateMediaElementSource(mediaElement htmlmediaelement.HTMLMediaElement) (a *audionode.MediaElementAudioSourceNode) {
	macro.Rewrite("$_.createMediaElementSource($1)", mediaElement)
	return a
}

// CreateMediaStreamSource fn
// js:"createMediaStreamSource"
func (*OfflineAudioContext) CreateMediaStreamSource(mediaStream *window.MediaStream) (a *audionode.MediaStreamAudioSourceNode) {
	macro.Rewrite("$_.createMediaStreamSource($1)", mediaStream)
	return a
}

// CreateOscillator fn
// js:"createOscillator"
func (*OfflineAudioContext) CreateOscillator() (a *audionode.OscillatorNode) {
	macro.Rewrite("$_.createOscillator()")
	return a
}

// CreatePanner fn
// js:"createPanner"
func (*OfflineAudioContext) CreatePanner() (a *audionode.PannerNode) {
	macro.Rewrite("$_.createPanner()")
	return a
}

// CreatePeriodicWave fn
// js:"createPeriodicWave"
func (*OfflineAudioContext) CreatePeriodicWave(rl []float32, img []float32, constraints *periodicwaveconstraints.PeriodicWaveConstraints) (p *periodicwave.PeriodicWave) {
	macro.Rewrite("$_.createPeriodicWave($1, $2, $3)", rl, img, constraints)
	return p
}

// CreateScriptProcessor fn
// js:"createScriptProcessor"
func (*OfflineAudioContext) CreateScriptProcessor(bufferSize *uint, numberOfInputChannels *uint, numberOfOutputChannels *uint) (a *audionode.ScriptProcessorNode) {
	macro.Rewrite("$_.createScriptProcessor($1, $2, $3)", bufferSize, numberOfInputChannels, numberOfOutputChannels)
	return a
}

// CreateStereoPanner fn
// js:"createStereoPanner"
func (*OfflineAudioContext) CreateStereoPanner() (a *audionode.StereoPannerNode) {
	macro.Rewrite("$_.createStereoPanner()")
	return a
}

// CreateWaveShaper fn
// js:"createWaveShaper"
func (*OfflineAudioContext) CreateWaveShaper() (a *audionode.WaveShaperNode) {
	macro.Rewrite("$_.createWaveShaper()")
	return a
}

// DecodeAudioData fn
// js:"decodeAudioData"
func (*OfflineAudioContext) DecodeAudioData(audioData []byte, successCallback *func(decodedData *audiobuffer.AudioBuffer), errorCallback *func()) (a *audiobuffer.AudioBuffer) {
	macro.Rewrite("await $_.decodeAudioData($1, $2, $3)", audioData, successCallback, errorCallback)
	return a
}

// Resume fn
// js:"resume"
func (*OfflineAudioContext) Resume() {
	macro.Rewrite("await $_.resume()")
}

// AddEventListener fn
// js:"addEventListener"
func (*OfflineAudioContext) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*OfflineAudioContext) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*OfflineAudioContext) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Length prop
// js:"length"
func (*OfflineAudioContext) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

// Oncomplete prop
// js:"oncomplete"
func (*OfflineAudioContext) Oncomplete() (oncomplete func(*offlineaudiocompletionevent.OfflineAudioCompletionEvent)) {
	macro.Rewrite("$_.oncomplete")
	return oncomplete
}

// SetOncomplete prop
// js:"oncomplete"
func (*OfflineAudioContext) SetOncomplete(oncomplete func(*offlineaudiocompletionevent.OfflineAudioCompletionEvent)) {
	macro.Rewrite("$_.oncomplete = $1", oncomplete)
}

// CurrentTime prop
// js:"currentTime"
func (*OfflineAudioContext) CurrentTime() (currentTime float32) {
	macro.Rewrite("$_.currentTime")
	return currentTime
}

// Destination prop
// js:"destination"
func (*OfflineAudioContext) Destination() (destination *audionode.AudioDestinationNode) {
	macro.Rewrite("$_.destination")
	return destination
}

// Listener prop
// js:"listener"
func (*OfflineAudioContext) Listener() (listener *audiolistener.AudioListener) {
	macro.Rewrite("$_.listener")
	return listener
}

// Onstatechange prop
// js:"onstatechange"
func (*OfflineAudioContext) Onstatechange() (onstatechange func(window.Event)) {
	macro.Rewrite("$_.onstatechange")
	return onstatechange
}

// SetOnstatechange prop
// js:"onstatechange"
func (*OfflineAudioContext) SetOnstatechange(onstatechange func(window.Event)) {
	macro.Rewrite("$_.onstatechange = $1", onstatechange)
}

// SampleRate prop
// js:"sampleRate"
func (*OfflineAudioContext) SampleRate() (sampleRate float32) {
	macro.Rewrite("$_.sampleRate")
	return sampleRate
}

// State prop
// js:"state"
func (*OfflineAudioContext) State() (state *audiocontextstate.AudioContextState) {
	macro.Rewrite("$_.state")
	return state
}
