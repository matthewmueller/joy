package rtcrtpcapabilities

import (
	"github.com/matthewmueller/golly/dom2/rtcrtpcodeccapability"
	"github.com/matthewmueller/golly/dom2/rtcrtpheaderextension"
)

type RTCRtpCapabilities struct {
	codecs           *[]*rtcrtpcodeccapability.RTCRtpCodecCapability
	fecMechanisms    *[]string
	headerExtensions *[]*rtcrtpheaderextension.RTCRtpHeaderExtension
}
