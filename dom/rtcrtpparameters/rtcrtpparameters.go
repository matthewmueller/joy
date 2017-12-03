package rtcrtpparameters

import (
	"github.com/matthewmueller/joy/dom/rtcdegradationpreference"
	"github.com/matthewmueller/joy/dom/rtcrtcpparameters"
	"github.com/matthewmueller/joy/dom/rtcrtpcodecparameters"
	"github.com/matthewmueller/joy/dom/rtcrtpencodingparameters"
	"github.com/matthewmueller/joy/dom/rtcrtpheaderextensionparameters"
)

type RTCRtpParameters struct {
	codecs                *[]*rtcrtpcodecparameters.RTCRtpCodecParameters
	degradationPreference *rtcdegradationpreference.RTCDegradationPreference
	encodings             *[]*rtcrtpencodingparameters.RTCRtpEncodingParameters
	headerExtensions      *[]*rtcrtpheaderextensionparameters.RTCRtpHeaderExtensionParameters
	muxId                 *string
	rtcp                  *rtcrtcpparameters.RTCRtcpParameters
}
