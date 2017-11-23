package msaudiosendpayload

import (
	"github.com/matthewmueller/golly/dom2/msaudiosendsignal"
	"github.com/matthewmueller/golly/dom2/mspayloadbase"
)

type MSAudioSendPayload struct {
	*mspayloadbase.MSPayloadBase

	audioFECUsed    *bool
	samplingRate    *uint
	sendMutePercent *float32
	signal          *msaudiosendsignal.MSAudioSendSignal
}
