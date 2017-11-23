package rtcrtpparameters

import (
	"github.com/matthewmueller/golly/dom2/rtcdegradationpreference"
	"github.com/matthewmueller/golly/dom2/rtcrtcpparameters"
	"github.com/matthewmueller/golly/dom2/rtcrtpencodingparameters"
)

type RTCRtpParameters struct {
	codecs                *[]*rtcrtpcodecparameters.RTCRtpCodecParameters
	degradationPreference *rtcdegradationpreference.RTCDegradationPreference
	encodings             *[]*rtcrtpencodingparameters.RTCRtpEncodingParameters
	headerExtensions      *[]*rtcrtpheaderextensionparameters.RTCRtpHeaderExtensionParameters
	muxId                 *string
	rtcp                  *rtcrtcpparameters.RTCRtcpParameters
}
