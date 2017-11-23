package rtcdtlsparameters

type RTCDtlsParameters struct {
	fingerprints *[]*rtcdtlsfingerprint.RTCDtlsFingerprint
	role         *rtcdtlsrole.RTCDtlsRole
}
