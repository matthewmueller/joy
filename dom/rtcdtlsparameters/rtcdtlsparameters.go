package rtcdtlsparameters

import (
	"github.com/matthewmueller/golly/dom/rtcdtlsfingerprint"
	"github.com/matthewmueller/golly/dom/rtcdtlsrole"
)

type RTCDtlsParameters struct {
	fingerprints *[]*rtcdtlsfingerprint.RTCDtlsFingerprint
	role         *rtcdtlsrole.RTCDtlsRole
}
