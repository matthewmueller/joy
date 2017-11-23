package rtcicecandidatepairchangedevent

import (
	"github.com/matthewmueller/golly/dom/rtcicecandidatepair"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// RTCIceCandidatePairChangedEvent struct
// js:"RTCIceCandidatePairChangedEvent,omit"
type RTCIceCandidatePairChangedEvent struct {
	window.Event
}

// Pair prop
func (*RTCIceCandidatePairChangedEvent) Pair() (pair *rtcicecandidatepair.RTCIceCandidatePair) {
	js.Rewrite("$<.pair")
	return pair
}
