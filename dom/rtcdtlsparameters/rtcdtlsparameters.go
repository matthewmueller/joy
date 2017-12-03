package rtcdtlsparameters

import (
	"github.com/matthewmueller/joy/dom/rtcdtlsfingerprint"
	"github.com/matthewmueller/joy/dom/rtcdtlsrole"
)

type RTCDtlsParameters struct {
	fingerprints *[]*rtcdtlsfingerprint.RTCDtlsFingerprint
	role         *rtcdtlsrole.RTCDtlsRole
}
