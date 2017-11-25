package audionode

import (
	"github.com/matthewmueller/golly/dom/audiobuffer"
	"github.com/matthewmueller/golly/dom/audiocontextstate"
	"github.com/matthewmueller/golly/dom/audiolistener"
	"github.com/matthewmueller/golly/dom/htmlmediaelement"
	"github.com/matthewmueller/golly/dom/periodicwave"
	"github.com/matthewmueller/golly/dom/periodicwaveconstraints"
	"github.com/matthewmueller/golly/dom/window"
)

// js:"AudioContext,omit"
type AudioContext interface {
	window.EventTarget

	// Close
	// js:"close"
	Close()

	// CreateAnalyser
	// js:"createAnalyser"
	CreateAnalyser() (a *AnalyserNode)

	// CreateBiquadFilter
	// js:"createBiquadFilter"
	CreateBiquadFilter() (b *BiquadFilterNode)

	// CreateBuffer
	// js:"createBuffer"
	CreateBuffer(numberOfChannels uint, length uint, sampleRate float32) (a *audiobuffer.AudioBuffer)

	// CreateBufferSource
	// js:"createBufferSource"
	CreateBufferSource() (a *AudioBufferSourceNode)

	// CreateChannelMerger
	// js:"createChannelMerger"
	CreateChannelMerger(numberOfInputs *uint) (c *ChannelMergerNode)

	// CreateChannelSplitter
	// js:"createChannelSplitter"
	CreateChannelSplitter(numberOfOutputs *uint) (c *ChannelSplitterNode)

	// CreateConvolver
	// js:"createConvolver"
	CreateConvolver() (c *ConvolverNode)

	// CreateDelay
	// js:"createDelay"
	CreateDelay(maxDelayTime *float32) (d *DelayNode)

	// CreateDynamicsCompressor
	// js:"createDynamicsCompressor"
	CreateDynamicsCompressor() (d *DynamicsCompressorNode)

	// CreateGain
	// js:"createGain"
	CreateGain() (g *GainNode)

	// CreateIIRFilter
	// js:"createIIRFilter"
	CreateIIRFilter(feedforward []float32, feedback []float32) (i *IIRFilterNode)

	// CreateMediaElementSource
	// js:"createMediaElementSource"
	CreateMediaElementSource(mediaElement htmlmediaelement.HTMLMediaElement) (m *MediaElementAudioSourceNode)

	// CreateMediaStreamSource
	// js:"createMediaStreamSource"
	CreateMediaStreamSource(mediaStream *window.MediaStream) (m *MediaStreamAudioSourceNode)

	// CreateOscillator
	// js:"createOscillator"
	CreateOscillator() (o *OscillatorNode)

	// CreatePanner
	// js:"createPanner"
	CreatePanner() (p *PannerNode)

	// CreatePeriodicWave
	// js:"createPeriodicWave"
	CreatePeriodicWave(rl []float32, img []float32, constraints *periodicwaveconstraints.PeriodicWaveConstraints) (p *periodicwave.PeriodicWave)

	// CreateScriptProcessor
	// js:"createScriptProcessor"
	CreateScriptProcessor(bufferSize *uint, numberOfInputChannels *uint, numberOfOutputChannels *uint) (s *ScriptProcessorNode)

	// CreateStereoPanner
	// js:"createStereoPanner"
	CreateStereoPanner() (s *StereoPannerNode)

	// CreateWaveShaper
	// js:"createWaveShaper"
	CreateWaveShaper() (w *WaveShaperNode)

	// DecodeAudioData
	// js:"decodeAudioData"
	DecodeAudioData(audioData []byte, successCallback *func(decodedData *audiobuffer.AudioBuffer), errorCallback *func()) (a *audiobuffer.AudioBuffer)

	// Resume
	// js:"resume"
	Resume()

	// Suspend
	// js:"suspend"
	Suspend()

	// CurrentTime prop
	// js:"currentTime"
	CurrentTime() (currentTime float32)

	// Destination prop
	// js:"destination"
	Destination() (destination *AudioDestinationNode)

	// Listener prop
	// js:"listener"
	Listener() (listener *audiolistener.AudioListener)

	// Onstatechange prop
	// js:"onstatechange"
	Onstatechange() (onstatechange func(window.Event))

	// Onstatechange prop
	// js:"setonstatechange"
	SetOnstatechange(onstatechange func(window.Event))

	// SampleRate prop
	// js:"sampleRate"
	SampleRate() (sampleRate float32)

	// State prop
	// js:"state"
	State() (state *audiocontextstate.AudioContextState)
}
