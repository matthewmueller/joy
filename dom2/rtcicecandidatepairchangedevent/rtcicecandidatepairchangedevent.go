package rtcicecandidatepairchangedevent

import (
	"github.com/matthewmueller/golly/dom2/rtcicecandidatepair"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"RTCIceCandidatePairChangedEvent,omit"
type RTCIceCandidatePairChangedEvent struct {
	window.Event
}

// Pair
func (*RTCIceCandidatePairChangedEvent) Pair() (pair *rtcicecandidatepair.RTCIceCandidatePair) {
	js.Rewrite("$<.Pair")
	return pair
}
