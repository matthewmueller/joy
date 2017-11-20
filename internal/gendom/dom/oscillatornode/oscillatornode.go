package oscillatornode

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/audionode"
	"github.com/matthewmueller/golly/internal/gendom/dom/audioparam"
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/oscillatortype"
	"github.com/matthewmueller/golly/internal/gendom/dom/periodicwave"
	"github.com/matthewmueller/golly/js"
)

type OscillatorNode struct {
	*audionode.AudioNode
}

func (*OscillatorNode) SetPeriodicWave(periodicWave *periodicwave.PeriodicWave) {
	js.Rewrite("$<.setPeriodicWave($1)", periodicWave)
}

func (*OscillatorNode) Start(when float32) {
	js.Rewrite("$<.start($1)", when)
}

func (*OscillatorNode) Stop(when float32) {
	js.Rewrite("$<.stop($1)", when)
}

func (*OscillatorNode) GetDetune() (detune *audioparam.AudioParam) {
	js.Rewrite("$<.detune")
	return detune
}

func (*OscillatorNode) GetFrequency() (frequency *audioparam.AudioParam) {
	js.Rewrite("$<.frequency")
	return frequency
}

func (*OscillatorNode) GetOnended() (e *event.Event) {
	js.Rewrite("$<.onended")
	return e
}

func (*OscillatorNode) SetOnended(e *event.Event) {
	js.Rewrite("$<.onended = $1", e)
}

func (*OscillatorNode) GetType() (kind *oscillatortype.OscillatorType) {
	js.Rewrite("$<.type")
	return kind
}

func (*OscillatorNode) SetType(kind *oscillatortype.OscillatorType) {
	js.Rewrite("$<.type = $1", kind)
}
