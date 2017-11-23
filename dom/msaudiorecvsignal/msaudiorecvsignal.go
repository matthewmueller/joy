package msaudiorecvsignal

type MSAudioRecvSignal struct {
	initialSignalLevelRMS     *float32
	recvNoiseLevelCh1         *int
	recvSignalLevelCh1        *int
	renderLoopbackSignalLevel *float32
	renderNoiseLevel          *float32
	renderSignalLevel         *float32
}
