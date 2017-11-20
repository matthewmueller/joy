package rtcdtlsparameters

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcdtlsfingerprint"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcdtlsrole"
)

type RTCDtlsParameters struct {
	fingerprints *[]*rtcdtlsfingerprint.RTCDtlsFingerprint
	role         *rtcdtlsrole.RTCDtlsRole
}
