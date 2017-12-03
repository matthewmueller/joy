package rtcrtpcodecparameters

import "github.com/matthewmueller/joy/dom/rtcrtcpfeedback"

type RTCRtpCodecParameters struct {
	clockRate    *uint
	maxptime     *uint
	name         *string
	numChannels  *uint
	parameters   *interface{}
	payloadType  *interface{}
	ptime        *uint
	rtcpFeedback *[]*rtcrtcpfeedback.RTCRtcpFeedback
}
