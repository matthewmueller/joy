package rtcdtlsparameters

import (
	"github.com/matthewmueller/golly/dom2/rtcdtlsfingerprint"
	"github.com/matthewmueller/golly/dom2/rtcdtlsrole"
)

type RTCDtlsParameters struct {
	fingerprints *[]*rtcdtlsfingerprint.RTCDtlsFingerprint
	role         *rtcdtlsrole.RTCDtlsRole
}
