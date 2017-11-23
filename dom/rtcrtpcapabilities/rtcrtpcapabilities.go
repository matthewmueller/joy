package rtcrtpcapabilities

import (
	"github.com/matthewmueller/golly/dom/rtcrtpcodeccapability"
	"github.com/matthewmueller/golly/dom/rtcrtpheaderextension"
)

type RTCRtpCapabilities struct {
	codecs           *[]*rtcrtpcodeccapability.RTCRtpCodecCapability
	fecMechanisms    *[]string
	headerExtensions *[]*rtcrtpheaderextension.RTCRtpHeaderExtension
}
