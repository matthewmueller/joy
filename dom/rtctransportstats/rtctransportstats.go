package rtctransportstats

import "github.com/matthewmueller/joy/dom/rtcstats"

type RTCTransportStats struct {
	*rtcstats.RTCStats

	activeConnection        *bool
	bytesReceived           *uint64
	bytesSent               *uint64
	localCertificateId      *string
	remoteCertificateId     *string
	rtcpTransportStatsId    *string
	selectedCandidatePairId *string
}
