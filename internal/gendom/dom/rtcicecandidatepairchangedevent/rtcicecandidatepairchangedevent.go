package rtcicecandidatepairchangedevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcicecandidatepair"
	"github.com/matthewmueller/golly/js"
)

type RTCIceCandidatePairChangedEvent struct {
	*event.Event
}

func (*RTCIceCandidatePairChangedEvent) GetPair() (pair *rtcicecandidatepair.RTCIceCandidatePair) {
	js.Rewrite("$<.pair")
	return pair
}
