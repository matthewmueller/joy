package msaudiosendpayload

import "github.com/matthewmueller/golly/internal/gendom/dom/msaudiosendsignal"

type MSAudioSendPayload struct {
	*MSPayloadBase

	audioFECUsed    *bool
	samplingRate    *uint
	sendMutePercent *float32
	signal          *msaudiosendsignal.MSAudioSendSignal
}
