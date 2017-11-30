package rtcicecandidatepair

import "github.com/matthewmueller/golly/dom/rtcicecandidatedictionary"

type RTCIceCandidatePair struct {
	local  *rtcicecandidatedictionary.RTCIceCandidateDictionary
	remote *rtcicecandidatedictionary.RTCIceCandidateDictionary
}
