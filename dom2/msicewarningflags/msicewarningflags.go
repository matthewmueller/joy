package msicewarningflags

type MSIceWarningFlags struct {
	allocationMessageIntegrityFailed *bool
	alternateServerReceived          *bool
	connCheckMessageIntegrityFailed  *bool
	connCheckOtherError              *bool
	fipsAllocationFailure            *bool
	multipleRelayServersAttempted    *bool
	noRelayServersConfigured         *bool
	portRangeExhausted               *bool
	pseudoTLSFailure                 *bool
	tcpNatConnectivityFailed         *bool
	tcpRelayConnectivityFailed       *bool
	turnAuthUnknownUsernameError     *bool
	turnTcpAllocateFailed            *bool
	turnTcpSendFailed                *bool
	turnTcpTimedOut                  *bool
	turnTurnTcpConnectivityFailed    *bool
	turnUdpAllocateFailed            *bool
	turnUdpSendFailed                *bool
	udpLocalConnectivityFailed       *bool
	udpNatConnectivityFailed         *bool
	udpRelayConnectivityFailed       *bool
	useCandidateChecksFailed         *bool
}
