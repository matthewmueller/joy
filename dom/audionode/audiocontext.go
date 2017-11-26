package audionode

import (
	"github.com/matthewmueller/golly/dom/audiobuffer"
	"github.com/matthewmueller/golly/dom/audiocontextstate"
	"github.com/matthewmueller/golly/dom/audiolistener"
	"github.com/matthewmueller/golly/dom/htmlmediaelement"
	"github.com/matthewmueller/golly/dom/periodicwave"
	"github.com/matthewmueller/golly/dom/periodicwaveconstraints"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// js:"AudioContext,omit"
type AudioContext interface {
	window.EventTarget

	// Close
	// js:"close,rewrite=cls"
	Close()

	// CreateAnalyser
	// js:"createAnalyser,rewrite=createanalyser"
	CreateAnalyser() (a *AnalyserNode)

	// CreateBiquadFilter
	// js:"createBiquadFilter,rewrite=createbiquadfilter"
	CreateBiquadFilter() (b *BiquadFilterNode)

	// CreateBuffer
	// js:"createBuffer,rewrite=createbuffer"
	CreateBuffer(numberOfChannels uint, length uint, sampleRate float32) (a *audiobuffer.AudioBuffer)

	// CreateBufferSource
	// js:"createBufferSource,rewrite=createbuffersource"
	CreateBufferSource() (a *AudioBufferSourceNode)

	// CreateChannelMerger
	// js:"createChannelMerger,rewrite=createchannelmerger"
	CreateChannelMerger(numberOfInputs *uint) (c *ChannelMergerNode)

	// CreateChannelSplitter
	// js:"createChannelSplitter,rewrite=createchannelsplitter"
	CreateChannelSplitter(numberOfOutputs *uint) (c *ChannelSplitterNode)

	// CreateConvolver
	// js:"createConvolver,rewrite=createconvolver"
	CreateConvolver() (c *ConvolverNode)

	// CreateDelay
	// js:"createDelay,rewrite=createdelay"
	CreateDelay(maxDelayTime *float32) (d *DelayNode)

	// CreateDynamicsCompressor
	// js:"createDynamicsCompressor,rewrite=createdynamicscompressor"
	CreateDynamicsCompressor() (d *DynamicsCompressorNode)

	// CreateGain
	// js:"createGain,rewrite=creategain"
	CreateGain() (g *GainNode)

	// CreateIIRFilter
	// js:"createIIRFilter,rewrite=createiirfilter"
	CreateIIRFilter(feedforward []float32, feedback []float32) (i *IIRFilterNode)

	// CreateMediaElementSource
	// js:"createMediaElementSource,rewrite=createmediaelementsource"
	CreateMediaElementSource(mediaElement htmlmediaelement.HTMLMediaElement) (m *MediaElementAudioSourceNode)

	// CreateMediaStreamSource
	// js:"createMediaStreamSource,rewrite=createmediastreamsource"
	CreateMediaStreamSource(mediaStream *window.MediaStream) (m *MediaStreamAudioSourceNode)

	// CreateOscillator
	// js:"createOscillator,rewrite=createoscillator"
	CreateOscillator() (o *OscillatorNode)

	// CreatePanner
	// js:"createPanner,rewrite=createpanner"
	CreatePanner() (p *PannerNode)

	// CreatePeriodicWave
	// js:"createPeriodicWave,rewrite=createperiodicwave"
	CreatePeriodicWave(rl []float32, img []float32, constraints *periodicwaveconstraints.PeriodicWaveConstraints) (p *periodicwave.PeriodicWave)

	// CreateScriptProcessor
	// js:"createScriptProcessor,rewrite=createscriptprocessor"
	CreateScriptProcessor(bufferSize *uint, numberOfInputChannels *uint, numberOfOutputChannels *uint) (s *ScriptProcessorNode)

	// CreateStereoPanner
	// js:"createStereoPanner,rewrite=createstereopanner"
	CreateStereoPanner() (s *StereoPannerNode)

	// CreateWaveShaper
	// js:"createWaveShaper,rewrite=createwaveshaper"
	CreateWaveShaper() (w *WaveShaperNode)

	// DecodeAudioData
	// js:"decodeAudioData,rewrite=decodeaudiodata"
	DecodeAudioData(audioData []byte, successCallback *func(decodedData *audiobuffer.AudioBuffer), errorCallback *func()) (a *audiobuffer.AudioBuffer)

	// Resume
	// js:"resume,rewrite=resume"
	Resume()

	// Suspend
	// js:"suspend,rewrite=suspend"
	Suspend()

	// CurrentTime prop
	// js:"currentTime,rewrite=currenttime"
	CurrentTime() (currentTime float32)

	// Destination prop
	// js:"destination,rewrite=destination"
	Destination() (destination *AudioDestinationNode)

	// Listener prop
	// js:"listener,rewrite=listener"
	Listener() (listener *audiolistener.AudioListener)

	// Onstatechange prop
	// js:"onstatechange,rewrite=onstatechange"
	Onstatechange() (onstatechange func(window.Event))

	// Onstatechange prop
	// js:"setonstatechange,rewrite=setonstatechange"
	SetOnstatechange(onstatechange func(window.Event))

	// SampleRate prop
	// js:"sampleRate,rewrite=samplerate"
	SampleRate() (sampleRate float32)

	// State prop
	// js:"state,rewrite=state"
	State() (state *audiocontextstate.AudioContextState)
}

// cls fn
func cls() {
	js.Rewrite("await $<.close()")
}

// createanalyser fn
func createanalyser() (a *AnalyserNode) {
	js.Rewrite("$<.createAnalyser()")
	return a
}

// createbiquadfilter fn
func createbiquadfilter() (b *BiquadFilterNode) {
	js.Rewrite("$<.createBiquadFilter()")
	return b
}

// createbuffer fn
func createbuffer(numberOfChannels uint, length uint, sampleRate float32) (a *audiobuffer.AudioBuffer) {
	js.Rewrite("$<.createBuffer($1, $2, $3)", numberOfChannels, length, sampleRate)
	return a
}

// createbuffersource fn
func createbuffersource() (a *AudioBufferSourceNode) {
	js.Rewrite("$<.createBufferSource()")
	return a
}

// createchannelmerger fn
func createchannelmerger(numberOfInputs *uint) (c *ChannelMergerNode) {
	js.Rewrite("$<.createChannelMerger($1)", numberOfInputs)
	return c
}

// createchannelsplitter fn
func createchannelsplitter(numberOfOutputs *uint) (c *ChannelSplitterNode) {
	js.Rewrite("$<.createChannelSplitter($1)", numberOfOutputs)
	return c
}

// createconvolver fn
func createconvolver() (c *ConvolverNode) {
	js.Rewrite("$<.createConvolver()")
	return c
}

// createdelay fn
func createdelay(maxDelayTime *float32) (d *DelayNode) {
	js.Rewrite("$<.createDelay($1)", maxDelayTime)
	return d
}

// createdynamicscompressor fn
func createdynamicscompressor() (d *DynamicsCompressorNode) {
	js.Rewrite("$<.createDynamicsCompressor()")
	return d
}

// creategain fn
func creategain() (g *GainNode) {
	js.Rewrite("$<.createGain()")
	return g
}

// createiirfilter fn
func createiirfilter(feedforward []float32, feedback []float32) (i *IIRFilterNode) {
	js.Rewrite("$<.createIIRFilter($1, $2)", feedforward, feedback)
	return i
}

// createmediaelementsource fn
func createmediaelementsource(mediaElement htmlmediaelement.HTMLMediaElement) (m *MediaElementAudioSourceNode) {
	js.Rewrite("$<.createMediaElementSource($1)", mediaElement)
	return m
}

// createmediastreamsource fn
func createmediastreamsource(mediaStream *window.MediaStream) (m *MediaStreamAudioSourceNode) {
	js.Rewrite("$<.createMediaStreamSource($1)", mediaStream)
	return m
}

// createoscillator fn
func createoscillator() (o *OscillatorNode) {
	js.Rewrite("$<.createOscillator()")
	return o
}

// createpanner fn
func createpanner() (p *PannerNode) {
	js.Rewrite("$<.createPanner()")
	return p
}

// createperiodicwave fn
func createperiodicwave(rl []float32, img []float32, constraints *periodicwaveconstraints.PeriodicWaveConstraints) (p *periodicwave.PeriodicWave) {
	js.Rewrite("$<.createPeriodicWave($1, $2, $3)", rl, img, constraints)
	return p
}

// createscriptprocessor fn
func createscriptprocessor(bufferSize *uint, numberOfInputChannels *uint, numberOfOutputChannels *uint) (s *ScriptProcessorNode) {
	js.Rewrite("$<.createScriptProcessor($1, $2, $3)", bufferSize, numberOfInputChannels, numberOfOutputChannels)
	return s
}

// createstereopanner fn
func createstereopanner() (s *StereoPannerNode) {
	js.Rewrite("$<.createStereoPanner()")
	return s
}

// createwaveshaper fn
func createwaveshaper() (w *WaveShaperNode) {
	js.Rewrite("$<.createWaveShaper()")
	return w
}

// decodeaudiodata fn
func decodeaudiodata(audioData []byte, successCallback *func(decodedData *audiobuffer.AudioBuffer), errorCallback *func()) (a *audiobuffer.AudioBuffer) {
	js.Rewrite("await $<.decodeAudioData($1, $2, $3)", audioData, successCallback, errorCallback)
	return a
}

// resume fn
func resume() {
	js.Rewrite("await $<.resume()")
}

// suspend fn
func suspend() {
	js.Rewrite("await $<.suspend()")
}

// currenttime prop
func currenttime() (currentTime float32) {
	js.Rewrite("$<.currentTime")
	return currentTime
}

// destination prop
func destination() (destination *AudioDestinationNode) {
	js.Rewrite("$<.destination")
	return destination
}

// listener prop
func listener() (listener *audiolistener.AudioListener) {
	js.Rewrite("$<.listener")
	return listener
}

// onstatechange prop
func onstatechange() (onstatechange func(window.Event)) {
	js.Rewrite("$<.onstatechange")
	return onstatechange
}

// setonstatechange prop
func setonstatechange(onstatechange func(window.Event)) {
	js.Rewrite("$<.onstatechange = onstatechange")
}

// samplerate prop
func samplerate() (sampleRate float32) {
	js.Rewrite("$<.sampleRate")
	return sampleRate
}

// state prop
func state() (state *audiocontextstate.AudioContextState) {
	js.Rewrite("$<.state")
	return state
}
