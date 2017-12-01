package rtcicecandidatedictionary

import (
	"github.com/matthewmueller/golly/dom/rtcicecandidatetype"
	"github.com/matthewmueller/golly/dom/rtciceprotocol"
	"github.com/matthewmueller/golly/dom/rtcicetcpcandidatetype"
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
