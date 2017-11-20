package rtcrtpparameters

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcdegradationpreference"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcrtcpparameters"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcrtpcodecparameters"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcrtpencodingparameters"
	"github.com/matthewmueller/golly/internal/gendom/dom/rtcrtpheaderextensionparameters"
)

type RTCRtpParameters struct {
	codecs                *[]*rtcrtpcodecparameters.RTCRtpCodecParameters
	degradationPreference *rtcdegradationpreference.RTCDegradationPreference
	encodings             *[]*rtcrtpencodingparameters.RTCRtpEncodingParameters
	headerExtensions      *[]*rtcrtpheaderextensionparameters.RTCRtpHeaderExtensionParameters
	muxId                 *string
	rtcp                  *rtcrtcpparameters.RTCRtcpParameters
}
