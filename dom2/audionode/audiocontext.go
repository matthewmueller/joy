package audionode

import (
	"github.com/matthewmueller/golly/dom2/audiobuffer"
	"github.com/matthewmueller/golly/dom2/audiocontextstate"
	"github.com/matthewmueller/golly/dom2/htmlmediaelement"
	"github.com/matthewmueller/golly/dom2/periodicwave"
	"github.com/matthewmueller/golly/dom2/periodicwaveconstraints"
	"github.com/matthewmueller/golly/dom2/window"
)

// js:"AudioContext,omit"
type AudioContext interface {
	window.EventTarget

	// Close
	Close()

	// CreateAnalyser
	CreateAnalyser() (a *AnalyserNode)

	// CreateBiquadFilter
	CreateBiquadFilter() (b *BiquadFilterNode)

	// CreateBuffer
	CreateBuffer(numberOfChannels uint, length uint, sampleRate float32) (a *audiobuffer.AudioBuffer)

	// CreateBufferSource
	CreateBufferSource() (a *AudioBufferSourceNode)

	// CreateChannelMerger
	CreateChannelMerger(numberOfInputs *uint) (c *ChannelMergerNode)

	// CreateChannelSplitter
	CreateChannelSplitter(numberOfOutputs *uint) (c *ChannelSplitterNode)

	// CreateConvolver
	CreateConvolver() (c *ConvolverNode)

	// CreateDelay
	CreateDelay(maxDelayTime *float32) (d *DelayNode)

	// CreateDynamicsCompressor
	CreateDynamicsCompressor() (d *DynamicsCompressorNode)

	// CreateGain
	CreateGain() (g *GainNode)

	// CreateIIRFilter
	CreateIIRFilter(feedforward []float32, feedback []float32) (i *IIRFilterNode)

	// CreateMediaElementSource
	CreateMediaElementSource(mediaElement htmlmediaelement.HTMLMediaElement) (m *MediaElementAudioSourceNode)

	// CreateMediaStreamSource
	CreateMediaStreamSource(mediaStream *window.MediaStream) (m *MediaStreamAudioSourceNode)

	// CreateOscillator
	CreateOscillator() (o *OscillatorNode)

	// CreatePanner
	CreatePanner() (p *PannerNode)

	// CreatePeriodicWave
	CreatePeriodicWave(rl []float32, img []float32, constraints *periodicwaveconstraints.PeriodicWaveConstraints) (p *periodicwave.PeriodicWave)

	// CreateScriptProcessor
	CreateScriptProcessor(bufferSize *uint, numberOfInputChannels *uint, numberOfOutputChannels *uint) (s *ScriptProcessorNode)

	// CreateStereoPanner
	CreateStereoPanner() (s *StereoPannerNode)

	// CreateWaveShaper
	CreateWaveShaper() (w *WaveShaperNode)

	// DecodeAudioData
	DecodeAudioData(audioData []byte, successCallback *func(decodedData *audiobuffer.AudioBuffer), errorCallback *func()) (a *audiobuffer.AudioBuffer)

	// Resume
	Resume()

	// Suspend
	Suspend()

	// CurrentTime
	CurrentTime() (currentTime float32)

	// Destination
	Destination() (destination *AudioDestinationNode)

	// Listener
	Listener() (listener *audiolistener.AudioListener)

	// Onstatechange
	Onstatechange() (onstatechange func(window.Event))

	// Onstatechange
	SetOnstatechange(onstatechange func(window.Event))

	// SampleRate
	SampleRate() (sampleRate float32)

	// State
	State() (state *audiocontextstate.AudioContextState)
}
