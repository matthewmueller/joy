package rtcrtpparameters

type RTCRtpParameters struct {
	codecs                *[]*rtcrtpcodecparameters.RTCRtpCodecParameters
	degradationPreference *rtcdegradationpreference.RTCDegradationPreference
	encodings             *[]*rtcrtpencodingparameters.RTCRtpEncodingParameters
	headerExtensions      *[]*rtcrtpheaderextensionparameters.RTCRtpHeaderExtensionParameters
	muxId                 *string
	rtcp                  *rtcrtcpparameters.RTCRtcpParameters
}
