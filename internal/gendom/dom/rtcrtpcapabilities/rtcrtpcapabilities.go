package rtcrtpcapabilities

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcrtpcodeccapability"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcrtpheaderextension"
)

type RTCRtpCapabilities struct {
	codecs           *[]*rtcrtpcodeccapability.RTCRtpCodecCapability
	fecMechanisms    *[]string
	headerExtensions *[]*rtcrtpheaderextension.RTCRtpHeaderExtension
}
