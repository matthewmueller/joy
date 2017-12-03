package audionode

import (
	"github.com/matthewmueller/joy/dom/audiobuffer"
	"github.com/matthewmueller/joy/dom/audiocontextstate"
	"github.com/matthewmueller/joy/dom/audiolistener"
	"github.com/matthewmueller/joy/dom/htmlmediaelement"
	"github.com/matthewmueller/joy/dom/periodicwave"
	"github.com/matthewmueller/joy/dom/periodicwaveconstraints"
	"github.com/matthewmueller/joy/dom/window"
)

// AudioContext interface
// js:"AudioContext"
type AudioContext interface {
	window.EventTarget

	// Close
	// js:"close"
	// jsrewrite:"await $_.close()"
	Close()

	// CreateAnalyser
	// js:"createAnalyser"
	// jsrewrite:"$_.createAnalyser()"
	CreateAnalyser() (a *AnalyserNode)

	// CreateBiquadFilter
	// js:"createBiquadFilter"
	// jsrewrite:"$_.createBiquadFilter()"
	CreateBiquadFilter() (b *BiquadFilterNode)

	// CreateBuffer
	// js:"createBuffer"
	// jsrewrite:"$_.createBuffer($1, $2, $3)"
	CreateBuffer(numberOfChannels uint, length uint, sampleRate float32) (a *audiobuffer.AudioBuffer)

	// CreateBufferSource
	// js:"createBufferSource"
	// jsrewrite:"$_.createBufferSource()"
	CreateBufferSource() (a *AudioBufferSourceNode)

	// CreateChannelMerger
	// js:"createChannelMerger"
	// jsrewrite:"$_.createChannelMerger($1)"
	CreateChannelMerger(numberOfInputs *uint) (c *ChannelMergerNode)

	// CreateChannelSplitter
	// js:"createChannelSplitter"
	// jsrewrite:"$_.createChannelSplitter($1)"
	CreateChannelSplitter(numberOfOutputs *uint) (c *ChannelSplitterNode)

	// CreateConvolver
	// js:"createConvolver"
	// jsrewrite:"$_.createConvolver()"
	CreateConvolver() (c *ConvolverNode)

	// CreateDelay
	// js:"createDelay"
	// jsrewrite:"$_.createDelay($1)"
	CreateDelay(maxDelayTime *float32) (d *DelayNode)

	// CreateDynamicsCompressor
	// js:"createDynamicsCompressor"
	// jsrewrite:"$_.createDynamicsCompressor()"
	CreateDynamicsCompressor() (d *DynamicsCompressorNode)

	// CreateGain
	// js:"createGain"
	// jsrewrite:"$_.createGain()"
	CreateGain() (g *GainNode)

	// CreateIIRFilter
	// js:"createIIRFilter"
	// jsrewrite:"$_.createIIRFilter($1, $2)"
	CreateIIRFilter(feedforward []float32, feedback []float32) (i *IIRFilterNode)

	// CreateMediaElementSource
	// js:"createMediaElementSource"
	// jsrewrite:"$_.createMediaElementSource($1)"
	CreateMediaElementSource(mediaElement htmlmediaelement.HTMLMediaElement) (m *MediaElementAudioSourceNode)

	// CreateMediaStreamSource
	// js:"createMediaStreamSource"
	// jsrewrite:"$_.createMediaStreamSource($1)"
	CreateMediaStreamSource(mediaStream *window.MediaStream) (m *MediaStreamAudioSourceNode)

	// CreateOscillator
	// js:"createOscillator"
	// jsrewrite:"$_.createOscillator()"
	CreateOscillator() (o *OscillatorNode)

	// CreatePanner
	// js:"createPanner"
	// jsrewrite:"$_.createPanner()"
	CreatePanner() (p *PannerNode)

	// CreatePeriodicWave
	// js:"createPeriodicWave"
	// jsrewrite:"$_.createPeriodicWave($1, $2, $3)"
	CreatePeriodicWave(rl []float32, img []float32, constraints *periodicwaveconstraints.PeriodicWaveConstraints) (p *periodicwave.PeriodicWave)

	// CreateScriptProcessor
	// js:"createScriptProcessor"
	// jsrewrite:"$_.createScriptProcessor($1, $2, $3)"
	CreateScriptProcessor(bufferSize *uint, numberOfInputChannels *uint, numberOfOutputChannels *uint) (s *ScriptProcessorNode)

	// CreateStereoPanner
	// js:"createStereoPanner"
	// jsrewrite:"$_.createStereoPanner()"
	CreateStereoPanner() (s *StereoPannerNode)

	// CreateWaveShaper
	// js:"createWaveShaper"
	// jsrewrite:"$_.createWaveShaper()"
	CreateWaveShaper() (w *WaveShaperNode)

	// DecodeAudioData
	// js:"decodeAudioData"
	// jsrewrite:"await $_.decodeAudioData($1, $2, $3)"
	DecodeAudioData(audioData []byte, successCallback *func(decodedData *audiobuffer.AudioBuffer), errorCallback *func()) (a *audiobuffer.AudioBuffer)

	// Resume
	// js:"resume"
	// jsrewrite:"await $_.resume()"
	Resume()

	// Suspend
	// js:"suspend"
	// jsrewrite:"await $_.suspend($1)"
	Suspend(suspendTime float32)

	// CurrentTime prop
	// js:"currentTime"
	// jsrewrite:"$_.currentTime"
	CurrentTime() (currentTime float32)

	// Destination prop
	// js:"destination"
	// jsrewrite:"$_.destination"
	Destination() (destination *AudioDestinationNode)

	// Listener prop
	// js:"listener"
	// jsrewrite:"$_.listener"
	Listener() (listener *audiolistener.AudioListener)

	// Onstatechange prop
	// js:"onstatechange"
	// jsrewrite:"$_.onstatechange"
	Onstatechange() (onstatechange func(window.Event))

	// SetOnstatechange prop
	// js:"onstatechange"
	// jsrewrite:"$_.onstatechange = $1"
	SetOnstatechange(onstatechange func(window.Event))

	// SampleRate prop
	// js:"sampleRate"
	// jsrewrite:"$_.sampleRate"
	SampleRate() (sampleRate float32)

	// State prop
	// js:"state"
	// jsrewrite:"$_.state"
	State() (state *audiocontextstate.AudioContextState)
}
