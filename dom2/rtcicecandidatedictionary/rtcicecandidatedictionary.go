package rtcicecandidatedictionary

import (
	"github.com/matthewmueller/golly/dom2/rtcicecandidatetype"
	"github.com/matthewmueller/golly/dom2/rtciceprotocol"
	"github.com/matthewmueller/golly/dom2/rtcicetcpcandidatetype"
)

type RTCIceCandidateDictionary struct {
	foundation       *string
	ip               *string
	msMTurnSessionId *string
	port             *uint8
	priority         *uint
	protocol         *rtciceprotocol.RTCIceProtocol
	relatedAddress   *string
	relatedPort      *uint8
	tcpType          *rtcicetcpcandidatetype.RTCIceTCPCandidateType
	kind             *rtcicecandidatetype.RTCIceCandidateType
}
