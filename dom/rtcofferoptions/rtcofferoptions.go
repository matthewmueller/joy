package rtcofferoptions

type RTCOfferOptions struct {
	iceRestart             *bool
	offerToReceiveAudio    *int
	offerToReceiveVideo    *int
	voiceActivityDetection *bool
}
