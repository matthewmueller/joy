package rtcicecandidatepair

import "github.com/matthewmueller/golly/dom2/rtcicecandidatedictionary"

type RTCIceCandidatePair struct {
	local  *rtcicecandidatedictionary.RTCIceCandidateDictionary
	remote *rtcicecandidatedictionary.RTCIceCandidateDictionary
}
