package rtctransportstats

type RTCTransportStats struct {
	*RTCStats

	activeConnection        *bool
	bytesReceived           *uint64
	bytesSent               *uint64
	localCertificateId      *string
	remoteCertificateId     *string
	rtcpTransportStatsId    *string
	selectedCandidatePairId *string
}
