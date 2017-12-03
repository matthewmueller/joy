package rtcicecandidatepair

import "github.com/matthewmueller/joy/dom/rtcicecandidatedictionary"

type RTCIceCandidatePair struct {
	local  *rtcicecandidatedictionary.RTCIceCandidateDictionary
	remote *rtcicecandidatedictionary.RTCIceCandidateDictionary
}
