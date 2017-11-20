package rtcicecandidatepair

import "github.com/matthewmueller/golly/internal/gendom/dom/rtcicecandidatedictionary"

type RTCIceCandidatePair struct {
	local  *rtcicecandidatedictionary.RTCIceCandidateDictionary
	remote *rtcicecandidatedictionary.RTCIceCandidateDictionary
}
