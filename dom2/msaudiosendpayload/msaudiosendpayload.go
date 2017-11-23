package msaudiosendpayload

type MSAudioSendPayload struct {
	*mspayloadbase.MSPayloadBase

	audioFECUsed    *bool
	samplingRate    *uint
	sendMutePercent *float32
	signal          *msaudiosendsignal.MSAudioSendSignal
}
