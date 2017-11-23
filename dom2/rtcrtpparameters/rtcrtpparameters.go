package rtcrtpparameters

import (
	"github.com/matthewmueller/golly/dom2/rtcdegradationpreference"
	"github.com/matthewmueller/golly/dom2/rtcrtcpparameters"
	"github.com/matthewmueller/golly/dom2/rtcrtpcodecparameters"
	"github.com/matthewmueller/golly/dom2/rtcrtpencodingparameters"
	"github.com/matthewmueller/golly/dom2/rtcrtpheaderextensionparameters"
)

type RTCRtpParameters struct {
	codecs                *[]*rtcrtpcodecparameters.RTCRtpCodecParameters
	degradationPreference *rtcdegradationpreference.RTCDegradationPreference
	encodings             *[]*rtcrtpencodingparameters.RTCRtpEncodingParameters
	headerExtensions      *[]*rtcrtpheaderextensionparameters.RTCRtpHeaderExtensionParameters
	muxId                 *string
	rtcp                  *rtcrtcpparameters.RTCRtcpParameters
}
