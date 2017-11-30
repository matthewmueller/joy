package rtcrtpparameters

import (
	"github.com/matthewmueller/golly/dom/rtcdegradationpreference"
	"github.com/matthewmueller/golly/dom/rtcrtcpparameters"
	"github.com/matthewmueller/golly/dom/rtcrtpcodecparameters"
	"github.com/matthewmueller/golly/dom/rtcrtpencodingparameters"
	"github.com/matthewmueller/golly/dom/rtcrtpheaderextensionparameters"
)

type RTCRtpParameters struct {
	codecs                *[]*rtcrtpcodecparameters.RTCRtpCodecParameters
	degradationPreference *rtcdegradationpreference.RTCDegradationPreference
	encodings             *[]*rtcrtpencodingparameters.RTCRtpEncodingParameters
	headerExtensions      *[]*rtcrtpheaderextensionparameters.RTCRtpHeaderExtensionParameters
	muxId                 *string
	rtcp                  *rtcrtcpparameters.RTCRtcpParameters
}
