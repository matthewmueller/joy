package rtcicecandidatedictionary

import (
	"github.com/matthewmueller/joy/dom/rtcicecandidatetype"
	"github.com/matthewmueller/joy/dom/rtciceprotocol"
	"github.com/matthewmueller/joy/dom/rtcicetcpcandidatetype"
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
