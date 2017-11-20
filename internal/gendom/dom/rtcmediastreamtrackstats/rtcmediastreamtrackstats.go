package rtcmediastreamtrackstats

type RTCMediaStreamTrackStats struct {
	*RTCStats

	audioLevel                *float32
	echoReturnLoss            *float32
	echoReturnLossEnhancement *float32
	frameHeight               *uint
	framesCorrupted           *uint
	framesDecoded             *uint
	framesDropped             *uint
	framesPerSecond           *float32
	framesReceived            *uint
	framesSent                *uint
	frameWidth                *uint
	remoteSource              *bool
	ssrcIds                   *[]string
	trackIdentifier           *string
}
