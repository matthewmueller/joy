package audiocontext

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/analysernode"
	"github.com/matthewmueller/golly/internal/gendom/dom/audiobuffer"
	"github.com/matthewmueller/golly/internal/gendom/dom/audiobuffersourcenode"
	"github.com/matthewmueller/golly/internal/gendom/dom/audiocontextstate"
	"github.com/matthewmueller/golly/internal/gendom/dom/audiodestinationnode"
	"github.com/matthewmueller/golly/internal/gendom/dom/audiolistener"
	"github.com/matthewmueller/golly/internal/gendom/dom/biquadfilternode"
	"github.com/matthewmueller/golly/internal/gendom/dom/channelmergernode"
	"github.com/matthewmueller/golly/internal/gendom/dom/channelsplitternode"
	"github.com/matthewmueller/golly/internal/gendom/dom/convolvernode"
	"github.com/matthewmueller/golly/internal/gendom/dom/delaynode"
	"github.com/matthewmueller/golly/internal/gendom/dom/dynamicscompressornode"
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/gainnode"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlmediaelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/iirfilternode"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediaelementaudiosourcenode"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediastream"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediastreamaudiosourcenode"
	"github.com/matthewmueller/golly/internal/gendom/dom/oscillatornode"
	"github.com/matthewmueller/golly/internal/gendom/dom/pannernode"
	"github.com/matthewmueller/golly/internal/gendom/dom/periodicwave"
	"github.com/matthewmueller/golly/internal/gendom/dom/periodicwaveconstraints"
	"github.com/matthewmueller/golly/internal/gendom/dom/scriptprocessornode"
	"github.com/matthewmueller/golly/internal/gendom/dom/stereopannernode"
	"github.com/matthewmueller/golly/internal/gendom/dom/waveshapernode"
	"github.com/matthewmueller/golly/js"
)

type AudioContext struct {
	*eventtarget.EventTarget
}

func (*AudioContext) Close() {
	js.Rewrite("await $<.close()")
}

func (*AudioContext) CreateAnalyser() (a *analysernode.AnalyserNode) {
	js.Rewrite("$<.createAnalyser()")
	return a
}

func (*AudioContext) CreateBiquadFilter() (b *biquadfilternode.BiquadFilterNode) {
	js.Rewrite("$<.createBiquadFilter()")
	return b
}

func (*AudioContext) CreateBuffer(numberOfChannels uint, length uint, sampleRate float32) (a *audiobuffer.AudioBuffer) {
	js.Rewrite("$<.createBuffer($1, $2, $3)", numberOfChannels, length, sampleRate)
	return a
}

func (*AudioContext) CreateBufferSource() (a *audiobuffersourcenode.AudioBufferSourceNode) {
	js.Rewrite("$<.createBufferSource()")
	return a
}

func (*AudioContext) CreateChannelMerger(numberOfInputs uint) (c *channelmergernode.ChannelMergerNode) {
	js.Rewrite("$<.createChannelMerger($1)", numberOfInputs)
	return c
}

func (*AudioContext) CreateChannelSplitter(numberOfOutputs uint) (c *channelsplitternode.ChannelSplitterNode) {
	js.Rewrite("$<.createChannelSplitter($1)", numberOfOutputs)
	return c
}

func (*AudioContext) CreateConvolver() (c *convolvernode.ConvolverNode) {
	js.Rewrite("$<.createConvolver()")
	return c
}

func (*AudioContext) CreateDelay(maxDelayTime float32) (d *delaynode.DelayNode) {
	js.Rewrite("$<.createDelay($1)", maxDelayTime)
	return d
}

func (*AudioContext) CreateDynamicsCompressor() (d *dynamicscompressornode.DynamicsCompressorNode) {
	js.Rewrite("$<.createDynamicsCompressor()")
	return d
}

func (*AudioContext) CreateGain() (g *gainnode.GainNode) {
	js.Rewrite("$<.createGain()")
	return g
}

func (*AudioContext) CreateIIRFilter(feedforward []float32, feedback []float32) (i *iirfilternode.IIRFilterNode) {
	js.Rewrite("$<.createIIRFilter($1, $2)", feedforward, feedback)
	return i
}

func (*AudioContext) CreateMediaElementSource(mediaElement *htmlmediaelement.HTMLMediaElement) (m *mediaelementaudiosourcenode.MediaElementAudioSourceNode) {
	js.Rewrite("$<.createMediaElementSource($1)", mediaElement)
	return m
}

func (*AudioContext) CreateMediaStreamSource(mediaStream *mediastream.MediaStream) (m *mediastreamaudiosourcenode.MediaStreamAudioSourceNode) {
	js.Rewrite("$<.createMediaStreamSource($1)", mediaStream)
	return m
}

func (*AudioContext) CreateOscillator() (o *oscillatornode.OscillatorNode) {
	js.Rewrite("$<.createOscillator()")
	return o
}

func (*AudioContext) CreatePanner() (p *pannernode.PannerNode) {
	js.Rewrite("$<.createPanner()")
	return p
}

func (*AudioContext) CreatePeriodicWave(rl []float32, img []float32, constraints *periodicwaveconstraints.PeriodicWaveConstraints) (p *periodicwave.PeriodicWave) {
	js.Rewrite("$<.createPeriodicWave($1, $2, $3)", rl, img, constraints)
	return p
}

func (*AudioContext) CreateScriptProcessor(bufferSize uint, numberOfInputChannels uint, numberOfOutputChannels uint) (s *scriptprocessornode.ScriptProcessorNode) {
	js.Rewrite("$<.createScriptProcessor($1, $2, $3)", bufferSize, numberOfInputChannels, numberOfOutputChannels)
	return s
}

func (*AudioContext) CreateStereoPanner() (s *stereopannernode.StereoPannerNode) {
	js.Rewrite("$<.createStereoPanner()")
	return s
}

func (*AudioContext) CreateWaveShaper() (w *waveshapernode.WaveShaperNode) {
	js.Rewrite("$<.createWaveShaper()")
	return w
}

func (*AudioContext) DecodeAudioData(audioData []byte, successCallback func(decodedData *audiobuffer.AudioBuffer), errorCallback func()) (a *audiobuffer.AudioBuffer) {
	js.Rewrite("await $<.decodeAudioData($1, $2, $3)", audioData, successCallback, errorCallback)
	return a
}

func (*AudioContext) Resume() {
	js.Rewrite("await $<.resume()")
}

func (*AudioContext) Suspend() {
	js.Rewrite("await $<.suspend()")
}

func (*AudioContext) GetCurrentTime() (currentTime float32) {
	js.Rewrite("$<.currentTime")
	return currentTime
}

func (*AudioContext) GetDestination() (destination *audiodestinationnode.AudioDestinationNode) {
	js.Rewrite("$<.destination")
	return destination
}

func (*AudioContext) GetListener() (listener *audiolistener.AudioListener) {
	js.Rewrite("$<.listener")
	return listener
}

func (*AudioContext) GetOnstatechange() (e *event.Event) {
	js.Rewrite("$<.onstatechange")
	return e
}

func (*AudioContext) SetOnstatechange(e *event.Event) {
	js.Rewrite("$<.onstatechange = $1", e)
}

func (*AudioContext) GetSampleRate() (sampleRate float32) {
	js.Rewrite("$<.sampleRate")
	return sampleRate
}

func (*AudioContext) GetState() (state *audiocontextstate.AudioContextState) {
	js.Rewrite("$<.state")
	return state
}
