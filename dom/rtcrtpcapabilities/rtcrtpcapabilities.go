package rtcrtpcapabilities

import (
	"github.com/matthewmueller/joy/dom/rtcrtpcodeccapability"
	"github.com/matthewmueller/joy/dom/rtcrtpheaderextension"
)

type RTCRtpCapabilities struct {
	codecs           *[]*rtcrtpcodeccapability.RTCRtpCodecCapability
	fecMechanisms    *[]string
	headerExtensions *[]*rtcrtpheaderextension.RTCRtpHeaderExtension
}
