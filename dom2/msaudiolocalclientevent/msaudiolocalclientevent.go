package msaudiolocalclientevent

type MSAudioLocalClientEvent struct {
	*mslocalclienteventbase.MSLocalClientEventBase

	cpuInsufficientEventRatio             *float32
	deviceCaptureNotFunctioningEventRatio *float32
	deviceClippingEventRatio              *float32
	deviceEchoEventRatio                  *float32
	deviceGlitchesEventRatio              *float32
	deviceHalfDuplexAECEventRatio         *float32
	deviceHowlingEventCount               *uint
	deviceLowSNREventRatio                *float32
	deviceLowSpeechLevelEventRatio        *float32
	deviceMultipleEndpointsEventCount     *uint
	deviceNearEndToEchoRatioEventRatio    *float32
	deviceRenderMuteEventRatio            *float32
	deviceRenderNotFunctioningEventRatio  *float32
	deviceRenderZeroVolumeEventRatio      *float32
	networkDelayEventRatio                *float32
	networkSendQualityEventRatio          *float32
}
