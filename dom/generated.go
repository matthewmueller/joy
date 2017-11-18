package dom

import (
	"time"

	"github.com/matthewmueller/golly/js"
)

type AppendMode string
type AudioContextState string
type BiquadFilterType string
type CanvasFillRule string
type ChannelCountMode string
type ChannelInterpretation string
type DistanceModelType string
type ExpandGranularity string
type GamepadInputEmulationType string
type IDBCursorDirection string
type IDBRequestReadyState string
type IDBTransactionMode string
type ListeningState string
type MediaDeviceKind string
type MediaKeyMessageType string
type MediaKeySessionType string
type MediaKeysRequirement string
type MediaKeyStatus string
type MediaStreamTrackState string
type MSCredentialType string
type MSIceAddrType string
type MSIceType string
type MSStatsType string
type MSTransportType string
type MSWebViewPermissionState string
type MSWebViewPermissionType string
type NavigationReason string
type NavigationType string
type NotificationDirection string
type NotificationPermission string
type OscillatorType string
type OverSampleType string
type PanningModelType string
type PaymentComplete string
type PaymentShippingType string
type PushEncryptionKeyName string
type PushPermissionState string
type ReferrerPolicy string
type RequestCache string
type RequestCredentials string
type RequestDestination string
type RequestMode string
type RequestRedirect string
type RequestType string
type ResponseType string
type RTCBundlePolicy string
type RTCDegradationPreference string
type RTCDtlsRole string
type RTCDtlsTransportState string
type RTCIceCandidateType string
type RTCIceComponent string
type RTCIceConnectionState string
type RTCIceGathererState string
type RTCIceGatheringState string
type RTCIceGatherPolicy string
type RTCIceProtocol string
type RTCIceRole string
type RTCIceTcpCandidateType string
type RTCIceTransportPolicy string
type RTCIceTransportState string
type RTCSdpType string
type RTCSignalingState string
type RTCStatsIceCandidatePairState string
type RTCStatsIceCandidateType string
type RTCStatsType string
type ScopedCredentialType string
type ServiceWorkerState string
type Transport string
type VideoFacingModeEnum string
type VisibilityState string
type XMLHttpRequestResponseType string

type Account struct {
	displayName   string
	id            string
	imageURL      *string
	name          *string
	rpDisplayName string
}

type Algorithm struct {
	name *string
}

type AnimationEventInit struct {
	*EventInit

	animationName *string
	elapsedTime   *float32
}

type AssertionOptions struct {
	allowList      []*ScopedCredentialDescriptor
	extensions     *WebAuthnExtensions
	rpId           *string
	timeoutSeconds *uint
}

type CacheQueryOptions struct {
	cacheName    *string
	ignoreMethod *bool
	ignoreSearch *bool
	ignoreVary   *bool
}

type ClientData struct {
	challenge    string
	extensions   *WebAuthnExtensions
	hashAlg      interface{}
	origin       string
	rpId         string
	tokenBinding *string
}

type CloseEventInit struct {
	*EventInit

	code     *uint8
	reason   *string
	wasClean *bool
}

type CompositionEventInit struct {
	*UIEventInit

	data *string
}

type ConfirmSiteSpecificExceptionsInformation struct {
	*ExceptionInformation

	arrayOfDomainStrings []string
}

type ConstrainBooleanParameters struct {
	exact *bool
	ideal *bool
}

type ConstrainDOMStringParameters struct {
	exact interface{}
	ideal interface{}
}

type ConstrainDoubleRange struct {
	*DoubleRange

	exact *float32
	ideal *float32
}

type ConstrainLongRange struct {
	*LongRange

	exact *int
	ideal *int
}

type ConstrainVideoFacingModeParameters struct {
	exact interface{}
	ideal interface{}
}

type CustomEventInit struct {
	*EventInit

	detail interface{}
}

type DeviceAccelerationDict struct {
	x *float32
	y *float32
	z *float32
}

type DeviceLightEventInit struct {
	*EventInit

	value *float32
}

type DeviceMotionEventInit struct {
	*EventInit

	acceleration                 *DeviceAccelerationDict
	accelerationIncludingGravity *DeviceAccelerationDict
	interval                     *float32
	rotationRate                 *DeviceRotationRateDict
}

type DeviceOrientationEventInit struct {
	*EventInit

	absolute *bool
	alpha    *float32
	beta     *float32
	gamma    *float32
}

type DeviceRotationRateDict struct {
	alpha *float32
	beta  *float32
	gamma *float32
}

type DOMRectInit struct {
	height *float32
	width  *float32
	x      *float32
	y      *float32
}

type DoubleRange struct {
	max *float32
	min *float32
}

type ErrorEventInit struct {
	*EventInit

	colno    *uint
	err      interface{}
	filename *string
	lineno   *uint
	message  *string
}

type EventInit struct {
	bubbles    *bool
	cancelable *bool
}

type EventModifierInit struct {
	*UIEventInit

	altKey             *bool
	ctrlKey            *bool
	metaKey            *bool
	modifierAltGraph   *bool
	modifierCapsLock   *bool
	modifierFn         *bool
	modifierFnLock     *bool
	modifierHyper      *bool
	modifierNumLock    *bool
	modifierOS         *bool
	modifierScrollLock *bool
	modifierSuper      *bool
	modifierSymbol     *bool
	modifierSymbolLock *bool
	shiftKey           *bool
}

type ExceptionInformation struct {
	domain *string
}

type FocusEventInit struct {
	*UIEventInit

	relatedTarget *EventTarget
}

type FocusNavigationEventInit struct {
	*EventInit

	navigationReason *string
	originHeight     *float32
	originLeft       *float32
	originTop        *float32
	originWidth      *float32
}

type FocusNavigationOrigin struct {
	originHeight *float32
	originLeft   *float32
	originTop    *float32
	originWidth  *float32
}

type GamepadEventInit struct {
	*EventInit

	gamepad *Gamepad
}

type GetNotificationOptions struct {
	tag *string
}

type HashChangeEventInit struct {
	*EventInit

	newURL *string
	oldURL *string
}

type IDBIndexParameters struct {
	unique *bool
}

type IDBObjectStoreParameters struct {
	keyPath *string
}

type IntersectionObserverEntryInit struct {
	boundingClientRect *DOMRectInit
	intersectionRect   *DOMRectInit
	rootBounds         *DOMRectInit
	target             *Element
	time               int
}

type IntersectionObserverInit struct {
	root       *Element
	rootMargin *string
	threshold  interface{}
}

type KeyAlgorithm struct {
	name *string
}

type KeyboardEventInit struct {
	*EventModifierInit

	key      *string
	location *uint
	repeat   *bool
}

type LongRange struct {
	max *int
	min *int
}

type MediaEncryptedEventInit struct {
	*EventInit

	initData     []byte
	initDataType *string
}

type MediaKeyMessageEventInit struct {
	*EventInit

	message     []byte
	messageType *MediaKeyMessageType
}

type MediaKeySystemConfiguration struct {
	audioCapabilities     []*MediaKeySystemMediaCapability
	distinctiveIdentifier *MediaKeysRequirement
	initDataTypes         []string
	persistentState       *MediaKeysRequirement
	videoCapabilities     []*MediaKeySystemMediaCapability
}

type MediaKeySystemMediaCapability struct {
	contentType *string
	robustness  *string
}

type MediaStreamConstraints struct {
	audio interface{}
	video interface{}
}

type MediaStreamErrorEventInit struct {
	*EventInit

	err *MediaStreamError
}

type MediaStreamEventInit struct {
	*EventInit

	stream *MediaStream
}

type MediaStreamTrackEventInit struct {
	*EventInit

	track *MediaStreamTrack
}

type MediaTrackCapabilities struct {
	aspectRatio      interface{}
	deviceId         *string
	echoCancellation []bool
	facingMode       *string
	frameRate        interface{}
	groupId          *string
	height           interface{}
	sampleRate       interface{}
	sampleSize       interface{}
	volume           interface{}
	width            interface{}
}

type MediaTrackConstraints struct {
	*MediaTrackConstraintSet

	advanced []*MediaTrackConstraintSet
}

type MediaTrackConstraintSet struct {
	aspectRatio     interface{}
	deviceId        interface{}
	echoCancelation interface{}
	facingMode      interface{}
	frameRate       interface{}
	groupId         interface{}
	height          interface{}
	sampleRate      interface{}
	sampleSize      interface{}
	volume          interface{}
	width           interface{}
}

type MediaTrackSettings struct {
	aspectRatio      *float32
	deviceId         *string
	echoCancellation *bool
	facingMode       *string
	frameRate        *float32
	groupId          *string
	height           *int
	sampleRate       *int
	sampleSize       *int
	volume           *float32
	width            *int
}

type MediaTrackSupportedConstraints struct {
	aspectRatio      *bool
	deviceId         *bool
	echoCancellation *bool
	facingMode       *bool
	frameRate        *bool
	groupId          *bool
	height           *bool
	sampleRate       *bool
	sampleSize       *bool
	volume           *bool
	width            *bool
}

type MessageEventInit struct {
	*EventInit

	data   interface{}
	origin *string
	ports  []*MessagePort
	source *Window
}

type MouseEventInit struct {
	*EventModifierInit

	button        *int8
	buttons       *uint8
	clientX       *int
	clientY       *int
	relatedTarget *EventTarget
	screenX       *int
	screenY       *int
}

type MSAccountInfo struct {
	accountImageUri *string
	accountName     *string
	rpDisplayName   string
	userDisplayName string
	userId          *string
}

type MSAudioLocalClientEvent struct {
	*MSLocalClientEventBase

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

type MSAudioRecvPayload struct {
	*MSPayloadBase

	burstLossLength1          *float32
	burstLossLength2          *float32
	burstLossLength3          *float32
	burstLossLength4          *float32
	burstLossLength5          *float32
	burstLossLength6          *float32
	burstLossLength7          *float32
	burstLossLength8OrHigher  *float32
	fecRecvDistance1          *float32
	fecRecvDistance2          *float32
	fecRecvDistance3          *float32
	packetReorderDepthAvg     *int
	packetReorderDepthMax     *int
	packetReorderRatio        *float32
	ratioCompressedSamplesAvg *float32
	ratioConcealedSamplesAvg  *float32
	ratioStretchedSamplesAvg  *float32
	samplingRate              *uint
	signal                    *MSAudioRecvSignal
}

type MSAudioRecvSignal struct {
	initialSignalLevelRMS     *float32
	recvNoiseLevelCh1         *int
	recvSignalLevelCh1        *int
	renderLoopbackSignalLevel *float32
	renderNoiseLevel          *float32
	renderSignalLevel         *float32
}

type MSAudioSendPayload struct {
	*MSPayloadBase

	audioFECUsed    *bool
	samplingRate    *uint
	sendMutePercent *float32
	signal          *MSAudioSendSignal
}

type MSAudioSendSignal struct {
	noiseLevel         *int
	sendNoiseLevelCh1  *int
	sendSignalLevelCh1 *int
}

type MSConnectivity struct {
	iceType         *MSIceType
	iceWarningFlags *MSIceWarningFlags
	relayAddress    *MSRelayAddress
}

type MSCredentialFilter struct {
	accept []*MSCredentialSpec
}

type MSCredentialParameters struct {
	kind *MSCredentialType
}

type MSCredentialSpec struct {
	id   *string
	kind *MSCredentialType
}

type MSDelay struct {
	roundTrip    *uint
	roundTripMax *uint
}

type MSDescription struct {
	*RTCStats

	connectivity         *MSConnectivity
	deviceDevName        *string
	localAddr            *MSIPAddressInfo
	networkconnectivity  *MSNetworkConnectivityInfo
	reflexiveLocalIPAddr *MSIPAddressInfo
	remoteAddr           *MSIPAddressInfo
	transport            *RTCIceProtocol
}

type MSFIDOCredentialParameters struct {
	*MSCredentialParameters

	algorithm      interface{}
	authenticators []string
}

type MSIceWarningFlags struct {
	allocationMessageIntegrityFailed *bool
	alternateServerReceived          *bool
	connCheckMessageIntegrityFailed  *bool
	connCheckOtherError              *bool
	fipsAllocationFailure            *bool
	multipleRelayServersAttempted    *bool
	noRelayServersConfigured         *bool
	portRangeExhausted               *bool
	pseudoTLSFailure                 *bool
	tcpNatConnectivityFailed         *bool
	tcpRelayConnectivityFailed       *bool
	turnAuthUnknownUsernameError     *bool
	turnTcpAllocateFailed            *bool
	turnTcpSendFailed                *bool
	turnTcpTimedOut                  *bool
	turnTurnTcpConnectivityFailed    *bool
	turnUdpAllocateFailed            *bool
	turnUdpSendFailed                *bool
	udpLocalConnectivityFailed       *bool
	udpNatConnectivityFailed         *bool
	udpRelayConnectivityFailed       *bool
	useCandidateChecksFailed         *bool
}

type MSIPAddressInfo struct {
	ipAddr                  *string
	manufacturerMacAddrMask *string
	port                    *uint8
}

type MSJitter struct {
	interArrival    *uint
	interArrivalMax *uint
	interArrivalSD  *float32
}

type MSLocalClientEventBase struct {
	*RTCStats

	networkBandwidthLowEventRatio   *float32
	networkReceiveQualityEventRatio *float32
}

type MSNetwork struct {
	*RTCStats

	delay       *MSDelay
	jitter      *MSJitter
	packetLoss  *MSPacketLoss
	utilization *MSUtilization
}

type MSNetworkConnectivityInfo struct {
	linkspeed                *uint
	networkConnectionDetails *string
	vpn                      *bool
}

type MSNetworkInterfaceType struct {
	interfaceTypeEthernet *bool
	interfaceTypePPP      *bool
	interfaceTypeTunnel   *bool
	interfaceTypeWireless *bool
	interfaceTypeWWAN     *bool
}

type MSOutboundNetwork struct {
	*MSNetwork

	appliedBandwidthLimit *uint
}

type MSPacketLoss struct {
	lossRate    *float32
	lossRateMax *float32
}

type MSPayloadBase struct {
	*RTCStats

	payloadDescription *string
}

type MSPortRange struct {
	max *uint8
	min *uint8
}

type MSRelayAddress struct {
	port         *uint8
	relayAddress *string
}

type MSSignatureParameters struct {
	userPrompt *string
}

type MSTransportDiagnosticsStats struct {
	*RTCStats

	allocationTimeInMs     *uint
	baseAddress            *string
	baseInterface          *MSNetworkInterfaceType
	iceRole                *RTCIceRole
	iceWarningFlags        *MSIceWarningFlags
	interfaces             *MSNetworkInterfaceType
	localAddress           *string
	localAddrType          *MSIceAddrType
	localInterface         *MSNetworkInterfaceType
	localMR                *string
	localMRTCPPort         *uint8
	localSite              *string
	msRtcEngineVersion     *string
	networkName            *string
	numConsentReqReceived  *uint
	numConsentReqSent      *uint
	numConsentRespReceived *uint
	numConsentRespSent     *uint
	portRangeMax           *uint8
	portRangeMin           *uint8
	protocol               *RTCIceProtocol
	remoteAddress          *string
	remoteAddrType         *MSIceAddrType
	remoteMR               *string
	remoteMRTCPPort        *uint8
	remoteSite             *string
	rtpRtcpMux             *bool
	stunVer                *uint
}

type MSUtilization struct {
	bandwidthEstimation       *uint64
	bandwidthEstimationAvg    *uint64
	bandwidthEstimationMax    *uint64
	bandwidthEstimationMin    *uint64
	bandwidthEstimationStdDev *uint64
	packets                   *uint64
}

type MSVideoPayload struct {
	*MSPayloadBase

	durationSeconds     *float32
	resolution          *string
	videoBitRateAvg     *uint
	videoBitRateMax     *uint
	videoFrameRateAvg   *float32
	videoPacketLossRate *float32
}

type MSVideoRecvPayload struct {
	*MSVideoPayload

	lowBitRateCallPercent                                *float32
	lowFrameRateCallPercent                              *float32
	recvBitRateAverage                                   *uint
	recvBitRateMaximum                                   *uint
	recvCodecType                                        *string
	recvFpsHarmonicAverage                               *float32
	recvFrameRateAverage                                 *float32
	recvNumResSwitches                                   *float32
	recvReorderBufferMaxSuccessfullyOrderedExtent        *uint
	recvReorderBufferMaxSuccessfullyOrderedLateTime      *uint
	recvReorderBufferPacketsDroppedDueToBufferExhaustion *uint
	recvReorderBufferPacketsDroppedDueToTimeout          *uint
	recvReorderBufferReorderedPackets                    *uint
	recvResolutionHeight                                 *uint
	recvResolutionWidth                                  *uint
	recvVideoStreamsMax                                  *uint
	recvVideoStreamsMin                                  *uint
	recvVideoStreamsMode                                 *int
	reorderBufferTotalPackets                            *uint
	videoFrameLossRate                                   *float32
	videoPostFECPLR                                      *float32
	videoResolutions                                     *MSVideoResolutionDistribution
}

type MSVideoResolutionDistribution struct {
	cifQuality   *uint
	h1080Quality *uint
	h1440Quality *uint
	h2160Quality *uint
	h720Quality  *uint
	vgaQuality   *uint
}

type MSVideoSendPayload struct {
	*MSVideoPayload

	sendBitRateAverage   *uint64
	sendBitRateMaximum   *uint64
	sendFrameRateAverage *float32
	sendResolutionHeight *uint
	sendResolutionWidth  *uint
	sendVideoStreamsMax  *uint
}

type MsZoomToOptions struct {
	animate     *string
	contentX    *int
	contentY    *int
	scaleFactor *float32
	viewportX   *string
	viewportY   *string
}

type MutationObserverInit struct {
	attributeFilter       []string
	attributeOldValue     *bool
	attributes            *bool
	characterData         *bool
	characterDataOldValue *bool
	childList             *bool
	subtree               *bool
}

type NotificationOptions struct {
	body *string
	dir  *NotificationDirection
	icon *string
	lang *string
	tag  *string
}

type ObjectURLOptions struct {
	oneTimeOnly *bool
}

type PaymentCurrencyAmount struct {
	currency       string
	currencySystem *string
	value          string
}

type PaymentDetails struct {
	displayItems    []*PaymentItem
	err             *string
	modifiers       []*PaymentDetailsModifier
	shippingOptions []*PaymentShippingOption
	total           *PaymentItem
}

type PaymentDetailsModifier struct {
	additionalDisplayItems []*PaymentItem
	data                   interface{}
	supportedMethods       []string
	total                  *PaymentItem
}

type PaymentItem struct {
	amount  *PaymentCurrencyAmount
	label   string
	pending *bool
}

type PaymentMethodData struct {
	data             interface{}
	supportedMethods []string
}

type PaymentOptions struct {
	requestPayerEmail *bool
	requestPayerName  *bool
	requestPayerPhone *bool
	requestShipping   *bool
	shippingType      *string
}

type PaymentRequestUpdateEventInit struct {
	*EventInit
}

type PaymentShippingOption struct {
	amount   *PaymentCurrencyAmount
	id       string
	label    string
	selected *bool
}

type PeriodicWaveConstraints struct {
	disableNormalization *bool
}

type PointerEventInit struct {
	*MouseEventInit

	height      *int
	isPrimary   *bool
	pointerId   *int
	pointerType *string
	pressure    *float32
	tiltX       *int
	tiltY       *int
	width       *int
}

type PopStateEventInit struct {
	*EventInit

	state interface{}
}

type PositionOptions struct {
	enableHighAccuracy *bool
	maximumAge         *int
	timeout            *int
}

type ProgressEventInit struct {
	*EventInit

	lengthComputable *bool
	loaded           *uint64
	total            *uint64
}

type PushSubscriptionOptionsInit struct {
	applicationServerKey []byte
	userVisibleOnly      *bool
}

type RegistrationOptions struct {
	scope *string
}

type RequestInit struct {
	body           interface{}
	cache          *RequestCache
	credentials    *RequestCredentials
	headers        interface{}
	integrity      *string
	keepalive      *bool
	method         *string
	mode           *RequestMode
	redirect       *RequestRedirect
	referrer       *string
	referrerPolicy *ReferrerPolicy
	window         interface{}
}

type ResponseInit struct {
	headers    interface{}
	status     *uint8
	statusText *string
}

type RTCConfiguration struct {
	bundlePolicy       *RTCBundlePolicy
	iceServers         []*RTCIceServer
	iceTransportPolicy *RTCIceTransportPolicy
	peerIdentity       *string
}

type RTCDtlsFingerprint struct {
	algorithm *string
	value     *string
}

type RTCDtlsParameters struct {
	fingerprints []*RTCDtlsFingerprint
	role         *RTCDtlsRole
}

type RTCDTMFToneChangeEventInit struct {
	*EventInit

	tone *string
}

type RTCIceCandidateAttributes struct {
	*RTCStats

	addressSourceUrl *string
	candidateType    *RTCStatsIceCandidateType
	ipAddress        *string
	portNumber       *int
	priority         *int
	transport        *string
}

type RTCIceCandidateComplete struct {
}

type RTCIceCandidateDictionary struct {
	foundation       *string
	ip               *string
	msMTurnSessionId *string
	port             *uint8
	priority         *uint
	protocol         *RTCIceProtocol
	relatedAddress   *string
	relatedPort      *uint8
	tcpType          *RTCIceTcpCandidateType
	kind             *RTCIceCandidateType
}

type RTCIceCandidateInit struct {
	candidate     *string
	sdpMid        *string
	sdpMLineIndex *uint8
}

type RTCIceCandidatePair struct {
	local  *RTCIceCandidateDictionary
	remote *RTCIceCandidateDictionary
}

type RTCIceCandidatePairStats struct {
	*RTCStats

	availableIncomingBitrate *float32
	availableOutgoingBitrate *float32
	bytesReceived            *uint64
	bytesSent                *uint64
	localCandidateId         *string
	nominated                *bool
	priority                 *uint64
	readable                 *bool
	remoteCandidateId        *string
	roundTripTime            *float32
	state                    *RTCStatsIceCandidatePairState
	transportId              *string
	writable                 *bool
}

type RTCIceGatherOptions struct {
	gatherPolicy *RTCIceGatherPolicy
	iceservers   []*RTCIceServer
	portRange    *MSPortRange
}

type RTCIceParameters struct {
	iceLite          *bool
	password         *string
	usernameFragment *string
}

type RTCIceServer struct {
	credential *string
	urls       interface{}
	username   *string
}

type RTCInboundRTPStreamStats struct {
	*RTCRTPStreamStats

	bytesReceived   *uint64
	fractionLost    *float32
	jitter          *float32
	packetsLost     *uint
	packetsReceived *uint
}

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
	ssrcIds                   []string
	trackIdentifier           *string
}

type RTCOfferOptions struct {
	iceRestart             *bool
	offerToReceiveAudio    *int
	offerToReceiveVideo    *int
	voiceActivityDetection *bool
}

type RTCOutboundRTPStreamStats struct {
	*RTCRTPStreamStats

	bytesSent     *uint64
	packetsSent   *uint
	roundTripTime *float32
	targetBitrate *float32
}

type RTCPeerConnectionIceEventInit struct {
	*EventInit

	candidate *RTCIceCandidate
}

type RTCRtcpFeedback struct {
	parameter *string
	kind      *string
}

type RTCRtcpParameters struct {
	cname       *string
	mux         *bool
	reducedSize *bool
	ssrc        *uint
}

type RTCRtpCapabilities struct {
	codecs           []*RTCRtpCodecCapability
	fecMechanisms    []string
	headerExtensions []*RTCRtpHeaderExtension
}

type RTCRtpCodecCapability struct {
	clockRate             *uint
	kind                  *string
	maxptime              *uint
	maxSpatialLayers      *uint8
	maxTemporalLayers     *uint8
	name                  *string
	numChannels           *uint
	options               interface{}
	parameters            interface{}
	preferredPayloadType  *byte
	ptime                 *uint
	rtcpFeedback          []*RTCRtcpFeedback
	svcMultiStreamSupport *bool
}

type RTCRtpCodecParameters struct {
	clockRate    *uint
	maxptime     *uint
	name         *string
	numChannels  *uint
	parameters   interface{}
	payloadType  interface{}
	ptime        *uint
	rtcpFeedback []*RTCRtcpFeedback
}

type RTCRtpContributingSource struct {
	audioLevel *byte
	csrc       *uint
	timestamp  *int
}

type RTCRtpEncodingParameters struct {
	active                *bool
	codecPayloadType      *byte
	dependencyEncodingIds []string
	encodingId            *string
	fec                   *RTCRtpFecParameters
	framerateScale        *float32
	maxBitrate            *float32
	maxFramerate          *uint
	minQuality            *float32
	priority              *float32
	resolutionScale       *float32
	rtx                   *RTCRtpRtxParameters
	ssrc                  *uint
	ssrcRange             *RTCSsrcRange
}

type RTCRtpFecParameters struct {
	mechanism *string
	ssrc      *uint
}

type RTCRtpHeaderExtension struct {
	kind             *string
	preferredEncrypt *bool
	preferredId      *uint8
	uri              *string
}

type RTCRtpHeaderExtensionParameters struct {
	encrypt *bool
	id      *uint8
	uri     *string
}

type RTCRtpParameters struct {
	codecs                []*RTCRtpCodecParameters
	degradationPreference *RTCDegradationPreference
	encodings             []*RTCRtpEncodingParameters
	headerExtensions      []*RTCRtpHeaderExtensionParameters
	muxId                 *string
	rtcp                  *RTCRtcpParameters
}

type RTCRtpRtxParameters struct {
	ssrc *uint
}

type RTCRTPStreamStats struct {
	*RTCStats

	associateStatsId *string
	codecId          *string
	firCount         *uint
	isRemote         *bool
	mediaTrackId     *string
	nackCount        *uint
	pliCount         *uint
	sliCount         *uint
	ssrc             *string
	transportId      *string
}

type RTCRtpUnhandled struct {
	muxId       *string
	payloadType *byte
	ssrc        *uint
}

type RTCSessionDescriptionInit struct {
	sdp  *string
	kind *RTCSdpType
}

type RTCSrtpKeyParam struct {
	keyMethod *string
	keySalt   *string
	lifetime  *string
	mkiLength *uint8
	mkiValue  *uint8
}

type RTCSrtpSdesParameters struct {
	cryptoSuite   *string
	keyParams     []*RTCSrtpKeyParam
	sessionParams []string
	tag           *uint8
}

type RTCSsrcRange struct {
	max *uint
	min *uint
}

type RTCStats struct {
	id        *string
	msType    *MSStatsType
	timestamp *int
	kind      *RTCStatsType
}

type RTCStatsReport struct {
}

type RTCTransportStats struct {
	*RTCStats

	activeConnection        *bool
	bytesReceived           *uint64
	bytesSent               *uint64
	localCertificateId      *string
	remoteCertificateId     *string
	rtcpTransportStatsId    *string
	selectedCandidatePairId *string
}

type ScopedCredentialDescriptor struct {
	id         []byte
	transports []*Transport
	kind       *ScopedCredentialType
}

type ScopedCredentialOptions struct {
	excludeList    []*ScopedCredentialDescriptor
	extensions     *WebAuthnExtensions
	rpId           *string
	timeoutSeconds *uint
}

type ScopedCredentialParameters struct {
	algorithm interface{}
	kind      *ScopedCredentialType
}

type ServiceWorkerMessageEventInit struct {
	*EventInit

	data        interface{}
	lastEventId *string
	origin      *string
	ports       []*MessagePort
	source      interface{}
}

type SpeechSynthesisEventInit struct {
	*EventInit

	charIndex   *uint
	elapsedTime *float32
	name        *string
	utterance   *SpeechSynthesisUtterance
}

type StoreExceptionsInformation struct {
	*ExceptionInformation

	detailURI         *string
	explanationString *string
	siteName          *string
}

type StoreSiteSpecificExceptionsInformation struct {
	*StoreExceptionsInformation

	arrayOfDomainStrings []string
}

type TrackEventInit struct {
	*EventInit

	track interface{}
}

type TransitionEventInit struct {
	*EventInit

	elapsedTime  *float32
	propertyName *string
}

type UIEventInit struct {
	*EventInit

	detail *int
	view   *Window
}

type WebAuthnExtensions struct {
}

type WebGLContextAttributes struct {
	alpha                 *bool
	antialias             *bool
	depth                 *bool
	premultipliedAlpha    *bool
	preserveDrawingBuffer *bool
	stencil               *bool
}

type WebGLContextEventInit struct {
	*EventInit

	statusMessage *string
}

type WheelEventInit struct {
	*MouseEventInit

	deltaMode *uint
	deltaX    *float32
	deltaY    *float32
	deltaZ    *float32
}

type AnalyserNode struct {
	*AudioNode
}

func (*AnalyserNode) GetByteFrequencyData(array []uint8) {
	js.Rewrite("$<.getByteFrequencyData($1)", array)
}

func (*AnalyserNode) GetByteTimeDomainData(array []uint8) {
	js.Rewrite("$<.getByteTimeDomainData($1)", array)
}

func (*AnalyserNode) GetFloatFrequencyData(array []float32) {
	js.Rewrite("$<.getFloatFrequencyData($1)", array)
}

func (*AnalyserNode) GetFloatTimeDomainData(array []float32) {
	js.Rewrite("$<.getFloatTimeDomainData($1)", array)
}

func (*AnalyserNode) GetFftSize() (fftSize uint) {
	js.Rewrite("$<.fftSize")
	return fftSize
}

func (*AnalyserNode) SetFftSize(fftSize uint) {
	js.Rewrite("$<.fftSize = $1", fftSize)
}

func (*AnalyserNode) GetFrequencyBinCount() (frequencyBinCount uint) {
	js.Rewrite("$<.frequencyBinCount")
	return frequencyBinCount
}

func (*AnalyserNode) GetMaxDecibels() (maxDecibels float32) {
	js.Rewrite("$<.maxDecibels")
	return maxDecibels
}

func (*AnalyserNode) SetMaxDecibels(maxDecibels float32) {
	js.Rewrite("$<.maxDecibels = $1", maxDecibels)
}

func (*AnalyserNode) GetMinDecibels() (minDecibels float32) {
	js.Rewrite("$<.minDecibels")
	return minDecibels
}

func (*AnalyserNode) SetMinDecibels(minDecibels float32) {
	js.Rewrite("$<.minDecibels = $1", minDecibels)
}

func (*AnalyserNode) GetSmoothingTimeConstant() (smoothingTimeConstant float32) {
	js.Rewrite("$<.smoothingTimeConstant")
	return smoothingTimeConstant
}

func (*AnalyserNode) SetSmoothingTimeConstant(smoothingTimeConstant float32) {
	js.Rewrite("$<.smoothingTimeConstant = $1", smoothingTimeConstant)
}

type ANGLE_instanced_arrays struct {
}

func (*ANGLE_instanced_arrays) DrawArraysInstancedANGLE(mode uint, first int, count int, primcount int) {
	js.Rewrite("$<.drawArraysInstancedANGLE($1, $2, $3, $4)", mode, first, count, primcount)
}

func (*ANGLE_instanced_arrays) DrawElementsInstancedANGLE(mode uint, count int, kind uint, offset int64, primcount int) {
	js.Rewrite("$<.drawElementsInstancedANGLE($1, $2, $3, $4, $5)", mode, count, kind, offset, primcount)
}

func (*ANGLE_instanced_arrays) VertexAttribDivisorANGLE(index uint, divisor uint) {
	js.Rewrite("$<.vertexAttribDivisorANGLE($1, $2)", index, divisor)
}

type AnimationEvent struct {
	*Event
}

func (*AnimationEvent) InitAnimationEvent(typeArg string, canBubbleArg bool, cancelableArg bool, animationNameArg string, elapsedTimeArg float32) {
	js.Rewrite("$<.initAnimationEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, animationNameArg, elapsedTimeArg)
}

func (*AnimationEvent) GetAnimationName() (animationName string) {
	js.Rewrite("$<.animationName")
	return animationName
}

func (*AnimationEvent) GetElapsedTime() (elapsedTime float32) {
	js.Rewrite("$<.elapsedTime")
	return elapsedTime
}

type ApplicationCache struct {
	*EventTarget
}

func (*ApplicationCache) Abort() {
	js.Rewrite("$<.abort()")
}

func (*ApplicationCache) SwapCache() {
	js.Rewrite("$<.swapCache()")
}

func (*ApplicationCache) Update() {
	js.Rewrite("$<.update()")
}

func (*ApplicationCache) GetOncached() (cached *Event) {
	js.Rewrite("$<.oncached")
	return cached
}

func (*ApplicationCache) SetOncached(cached *Event) {
	js.Rewrite("$<.oncached = $1", cached)
}

func (*ApplicationCache) GetOnchecking() (checking *Event) {
	js.Rewrite("$<.onchecking")
	return checking
}

func (*ApplicationCache) SetOnchecking(checking *Event) {
	js.Rewrite("$<.onchecking = $1", checking)
}

func (*ApplicationCache) GetOndownloading() (downloading *Event) {
	js.Rewrite("$<.ondownloading")
	return downloading
}

func (*ApplicationCache) SetOndownloading(downloading *Event) {
	js.Rewrite("$<.ondownloading = $1", downloading)
}

func (*ApplicationCache) GetOnerror() (err *Event) {
	js.Rewrite("$<.onerror")
	return err
}

func (*ApplicationCache) SetOnerror(err *Event) {
	js.Rewrite("$<.onerror = $1", err)
}

func (*ApplicationCache) GetOnnoupdate() (noupdate *Event) {
	js.Rewrite("$<.onnoupdate")
	return noupdate
}

func (*ApplicationCache) SetOnnoupdate(noupdate *Event) {
	js.Rewrite("$<.onnoupdate = $1", noupdate)
}

func (*ApplicationCache) GetOnobsolete() (obsolete *Event) {
	js.Rewrite("$<.onobsolete")
	return obsolete
}

func (*ApplicationCache) SetOnobsolete(obsolete *Event) {
	js.Rewrite("$<.onobsolete = $1", obsolete)
}

func (*ApplicationCache) GetOnprogress() (progress *ProgressEvent) {
	js.Rewrite("$<.onprogress")
	return progress
}

func (*ApplicationCache) SetOnprogress(progress *ProgressEvent) {
	js.Rewrite("$<.onprogress = $1", progress)
}

func (*ApplicationCache) GetOnupdateready() (updateready *Event) {
	js.Rewrite("$<.onupdateready")
	return updateready
}

func (*ApplicationCache) SetOnupdateready(updateready *Event) {
	js.Rewrite("$<.onupdateready = $1", updateready)
}

func (*ApplicationCache) GetStatus() (status uint8) {
	js.Rewrite("$<.status")
	return status
}

type Attr struct {
	*Node
}

func (*Attr) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*Attr) GetOwnerElement() (ownerElement *Element) {
	js.Rewrite("$<.ownerElement")
	return ownerElement
}

func (*Attr) GetPrefix() (prefix string) {
	js.Rewrite("$<.prefix")
	return prefix
}

func (*Attr) GetSpecified() (specified bool) {
	js.Rewrite("$<.specified")
	return specified
}

func (*Attr) GetValue() (value string) {
	js.Rewrite("$<.value")
	return value
}

func (*Attr) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}

type AudioBuffer struct {
}

func (*AudioBuffer) CopyFromChannel(destination []float32, channelNumber int, startInChannel uint) {
	js.Rewrite("$<.copyFromChannel($1, $2, $3)", destination, channelNumber, startInChannel)
}

func (*AudioBuffer) CopyToChannel(source []float32, channelNumber int, startInChannel uint) {
	js.Rewrite("$<.copyToChannel($1, $2, $3)", source, channelNumber, startInChannel)
}

func (*AudioBuffer) GetChannelData(channel uint) (f []float32) {
	js.Rewrite("$<.getChannelData($1)", channel)
	return f
}

func (*AudioBuffer) GetDuration() (duration float32) {
	js.Rewrite("$<.duration")
	return duration
}

func (*AudioBuffer) GetLength() (length int) {
	js.Rewrite("$<.length")
	return length
}

func (*AudioBuffer) GetNumberOfChannels() (numberOfChannels int) {
	js.Rewrite("$<.numberOfChannels")
	return numberOfChannels
}

func (*AudioBuffer) GetSampleRate() (sampleRate float32) {
	js.Rewrite("$<.sampleRate")
	return sampleRate
}

type AudioBufferSourceNode struct {
	*AudioNode
}

func (*AudioBufferSourceNode) Start(when float32, offset float32, duration float32) {
	js.Rewrite("$<.start($1, $2, $3)", when, offset, duration)
}

func (*AudioBufferSourceNode) Stop(when float32) {
	js.Rewrite("$<.stop($1)", when)
}

func (*AudioBufferSourceNode) GetBuffer() (buffer *AudioBuffer) {
	js.Rewrite("$<.buffer")
	return buffer
}

func (*AudioBufferSourceNode) SetBuffer(buffer *AudioBuffer) {
	js.Rewrite("$<.buffer = $1", buffer)
}

func (*AudioBufferSourceNode) GetDetune() (detune *AudioParam) {
	js.Rewrite("$<.detune")
	return detune
}

func (*AudioBufferSourceNode) GetLoop() (loop bool) {
	js.Rewrite("$<.loop")
	return loop
}

func (*AudioBufferSourceNode) SetLoop(loop bool) {
	js.Rewrite("$<.loop = $1", loop)
}

func (*AudioBufferSourceNode) GetLoopEnd() (loopEnd float32) {
	js.Rewrite("$<.loopEnd")
	return loopEnd
}

func (*AudioBufferSourceNode) SetLoopEnd(loopEnd float32) {
	js.Rewrite("$<.loopEnd = $1", loopEnd)
}

func (*AudioBufferSourceNode) GetLoopStart() (loopStart float32) {
	js.Rewrite("$<.loopStart")
	return loopStart
}

func (*AudioBufferSourceNode) SetLoopStart(loopStart float32) {
	js.Rewrite("$<.loopStart = $1", loopStart)
}

func (*AudioBufferSourceNode) GetOnended() (e *Event) {
	js.Rewrite("$<.onended")
	return e
}

func (*AudioBufferSourceNode) SetOnended(e *Event) {
	js.Rewrite("$<.onended = $1", e)
}

func (*AudioBufferSourceNode) GetPlaybackRate() (playbackRate *AudioParam) {
	js.Rewrite("$<.playbackRate")
	return playbackRate
}

type AudioContext struct {
	*EventTarget
}

func (*AudioContext) Close() {
	js.Rewrite("await $<.close()")
}

func (*AudioContext) CreateAnalyser() (a *AnalyserNode) {
	js.Rewrite("$<.createAnalyser()")
	return a
}

func (*AudioContext) CreateBiquadFilter() (b *BiquadFilterNode) {
	js.Rewrite("$<.createBiquadFilter()")
	return b
}

func (*AudioContext) CreateBuffer(numberOfChannels uint, length uint, sampleRate float32) (a *AudioBuffer) {
	js.Rewrite("$<.createBuffer($1, $2, $3)", numberOfChannels, length, sampleRate)
	return a
}

func (*AudioContext) CreateBufferSource() (a *AudioBufferSourceNode) {
	js.Rewrite("$<.createBufferSource()")
	return a
}

func (*AudioContext) CreateChannelMerger(numberOfInputs uint) (c *ChannelMergerNode) {
	js.Rewrite("$<.createChannelMerger($1)", numberOfInputs)
	return c
}

func (*AudioContext) CreateChannelSplitter(numberOfOutputs uint) (c *ChannelSplitterNode) {
	js.Rewrite("$<.createChannelSplitter($1)", numberOfOutputs)
	return c
}

func (*AudioContext) CreateConvolver() (c *ConvolverNode) {
	js.Rewrite("$<.createConvolver()")
	return c
}

func (*AudioContext) CreateDelay(maxDelayTime float32) (d *DelayNode) {
	js.Rewrite("$<.createDelay($1)", maxDelayTime)
	return d
}

func (*AudioContext) CreateDynamicsCompressor() (d *DynamicsCompressorNode) {
	js.Rewrite("$<.createDynamicsCompressor()")
	return d
}

func (*AudioContext) CreateGain() (g *GainNode) {
	js.Rewrite("$<.createGain()")
	return g
}

func (*AudioContext) CreateIIRFilter(feedforward []float32, feedback []float32) (i *IIRFilterNode) {
	js.Rewrite("$<.createIIRFilter($1, $2)", feedforward, feedback)
	return i
}

func (*AudioContext) CreateMediaElementSource(mediaElement *HTMLMediaElement) (m *MediaElementAudioSourceNode) {
	js.Rewrite("$<.createMediaElementSource($1)", mediaElement)
	return m
}

func (*AudioContext) CreateMediaStreamSource(mediaStream *MediaStream) (m *MediaStreamAudioSourceNode) {
	js.Rewrite("$<.createMediaStreamSource($1)", mediaStream)
	return m
}

func (*AudioContext) CreateOscillator() (o *OscillatorNode) {
	js.Rewrite("$<.createOscillator()")
	return o
}

func (*AudioContext) CreatePanner() (p *PannerNode) {
	js.Rewrite("$<.createPanner()")
	return p
}

func (*AudioContext) CreatePeriodicWave(rl []float32, img []float32, constraints *PeriodicWaveConstraints) (p *PeriodicWave) {
	js.Rewrite("$<.createPeriodicWave($1, $2, $3)", rl, img, constraints)
	return p
}

func (*AudioContext) CreateScriptProcessor(bufferSize uint, numberOfInputChannels uint, numberOfOutputChannels uint) (s *ScriptProcessorNode) {
	js.Rewrite("$<.createScriptProcessor($1, $2, $3)", bufferSize, numberOfInputChannels, numberOfOutputChannels)
	return s
}

func (*AudioContext) CreateStereoPanner() (s *StereoPannerNode) {
	js.Rewrite("$<.createStereoPanner()")
	return s
}

func (*AudioContext) CreateWaveShaper() (w *WaveShaperNode) {
	js.Rewrite("$<.createWaveShaper()")
	return w
}

func (*AudioContext) DecodeAudioData(audioData []byte, successCallback func(decodedData *AudioBuffer), errorCallback func()) (a *AudioBuffer) {
	js.Rewrite("await $<.decodeAudioData($1, $2, $3)", audioData, successCallback, errorCallback)
	return a
}

func (*AudioContext) Resume() {
	js.Rewrite("await $<.resume()")
}

func (*AudioContext) Suspend() {
	js.Rewrite("await $<.suspend()")
}

func (*AudioContext) GetCurrentTime() (currentTime float32) {
	js.Rewrite("$<.currentTime")
	return currentTime
}

func (*AudioContext) GetDestination() (destination *AudioDestinationNode) {
	js.Rewrite("$<.destination")
	return destination
}

func (*AudioContext) GetListener() (listener *AudioListener) {
	js.Rewrite("$<.listener")
	return listener
}

func (*AudioContext) GetOnstatechange() (e *Event) {
	js.Rewrite("$<.onstatechange")
	return e
}

func (*AudioContext) SetOnstatechange(e *Event) {
	js.Rewrite("$<.onstatechange = $1", e)
}

func (*AudioContext) GetSampleRate() (sampleRate float32) {
	js.Rewrite("$<.sampleRate")
	return sampleRate
}

func (*AudioContext) GetState() (state *AudioContextState) {
	js.Rewrite("$<.state")
	return state
}

type AudioDestinationNode struct {
	*AudioNode
}

func (*AudioDestinationNode) GetMaxChannelCount() (maxChannelCount uint) {
	js.Rewrite("$<.maxChannelCount")
	return maxChannelCount
}

type AudioListener struct {
}

func (*AudioListener) SetOrientation(x float32, y float32, z float32, xUp float32, yUp float32, zUp float32) {
	js.Rewrite("$<.setOrientation($1, $2, $3, $4, $5, $6)", x, y, z, xUp, yUp, zUp)
}

func (*AudioListener) SetPosition(x float32, y float32, z float32) {
	js.Rewrite("$<.setPosition($1, $2, $3)", x, y, z)
}

func (*AudioListener) SetVelocity(x float32, y float32, z float32) {
	js.Rewrite("$<.setVelocity($1, $2, $3)", x, y, z)
}

func (*AudioListener) GetDopplerFactor() (dopplerFactor float32) {
	js.Rewrite("$<.dopplerFactor")
	return dopplerFactor
}

func (*AudioListener) SetDopplerFactor(dopplerFactor float32) {
	js.Rewrite("$<.dopplerFactor = $1", dopplerFactor)
}

func (*AudioListener) GetSpeedOfSound() (speedOfSound float32) {
	js.Rewrite("$<.speedOfSound")
	return speedOfSound
}

func (*AudioListener) SetSpeedOfSound(speedOfSound float32) {
	js.Rewrite("$<.speedOfSound = $1", speedOfSound)
}

type AudioNode struct {
	*EventTarget
}

func (*AudioNode) Connect(destination *AudioNode, output uint, input uint) (a *AudioNode) {
	js.Rewrite("$<.connect($1, $2, $3)", destination, output, input)
	return a
}

func (*AudioNode) Disconnect() {
	js.Rewrite("$<.disconnect()")
}

func (*AudioNode) GetChannelCount() (channelCount uint) {
	js.Rewrite("$<.channelCount")
	return channelCount
}

func (*AudioNode) SetChannelCount(channelCount uint) {
	js.Rewrite("$<.channelCount = $1", channelCount)
}

func (*AudioNode) GetChannelCountMode() (channelCountMode *ChannelCountMode) {
	js.Rewrite("$<.channelCountMode")
	return channelCountMode
}

func (*AudioNode) SetChannelCountMode(channelCountMode *ChannelCountMode) {
	js.Rewrite("$<.channelCountMode = $1", channelCountMode)
}

func (*AudioNode) GetChannelInterpretation() (channelInterpretation *ChannelInterpretation) {
	js.Rewrite("$<.channelInterpretation")
	return channelInterpretation
}

func (*AudioNode) SetChannelInterpretation(channelInterpretation *ChannelInterpretation) {
	js.Rewrite("$<.channelInterpretation = $1", channelInterpretation)
}

func (*AudioNode) GetContext() (context *AudioContext) {
	js.Rewrite("$<.context")
	return context
}

func (*AudioNode) GetNumberOfInputs() (numberOfInputs uint) {
	js.Rewrite("$<.numberOfInputs")
	return numberOfInputs
}

func (*AudioNode) GetNumberOfOutputs() (numberOfOutputs uint) {
	js.Rewrite("$<.numberOfOutputs")
	return numberOfOutputs
}

type AudioParam struct {
}

func (*AudioParam) CancelScheduledValues(startTime float32) (a *AudioParam) {
	js.Rewrite("$<.cancelScheduledValues($1)", startTime)
	return a
}

func (*AudioParam) ExponentialRampToValueAtTime(value float32, endTime float32) (a *AudioParam) {
	js.Rewrite("$<.exponentialRampToValueAtTime($1, $2)", value, endTime)
	return a
}

func (*AudioParam) LinearRampToValueAtTime(value float32, endTime float32) (a *AudioParam) {
	js.Rewrite("$<.linearRampToValueAtTime($1, $2)", value, endTime)
	return a
}

func (*AudioParam) SetTargetAtTime(target float32, startTime float32, timeConstant float32) (a *AudioParam) {
	js.Rewrite("$<.setTargetAtTime($1, $2, $3)", target, startTime, timeConstant)
	return a
}

func (*AudioParam) SetValueAtTime(value float32, startTime float32) (a *AudioParam) {
	js.Rewrite("$<.setValueAtTime($1, $2)", value, startTime)
	return a
}

func (*AudioParam) SetValueCurveAtTime(values []float32, startTime float32, duration float32) (a *AudioParam) {
	js.Rewrite("$<.setValueCurveAtTime($1, $2, $3)", values, startTime, duration)
	return a
}

func (*AudioParam) GetDefaultValue() (defaultValue float32) {
	js.Rewrite("$<.defaultValue")
	return defaultValue
}

func (*AudioParam) GetValue() (value float32) {
	js.Rewrite("$<.value")
	return value
}

func (*AudioParam) SetValue(value float32) {
	js.Rewrite("$<.value = $1", value)
}

type AudioProcessingEvent struct {
	*Event
}

func (*AudioProcessingEvent) GetInputBuffer() (inputBuffer *AudioBuffer) {
	js.Rewrite("$<.inputBuffer")
	return inputBuffer
}

func (*AudioProcessingEvent) GetOutputBuffer() (outputBuffer *AudioBuffer) {
	js.Rewrite("$<.outputBuffer")
	return outputBuffer
}

func (*AudioProcessingEvent) GetPlaybackTime() (playbackTime float32) {
	js.Rewrite("$<.playbackTime")
	return playbackTime
}

type AudioTrack struct {
}

func (*AudioTrack) GetEnabled() (enabled bool) {
	js.Rewrite("$<.enabled")
	return enabled
}

func (*AudioTrack) SetEnabled(enabled bool) {
	js.Rewrite("$<.enabled = $1", enabled)
}

func (*AudioTrack) GetID() (id string) {
	js.Rewrite("$<.id")
	return id
}

func (*AudioTrack) GetKind() (kind string) {
	js.Rewrite("$<.kind")
	return kind
}

func (*AudioTrack) SetKind(kind string) {
	js.Rewrite("$<.kind = $1", kind)
}

func (*AudioTrack) GetLabel() (label string) {
	js.Rewrite("$<.label")
	return label
}

func (*AudioTrack) GetLanguage() (language string) {
	js.Rewrite("$<.language")
	return language
}

func (*AudioTrack) SetLanguage(language string) {
	js.Rewrite("$<.language = $1", language)
}

func (*AudioTrack) GetSourceBuffer() (sourceBuffer *SourceBuffer) {
	js.Rewrite("$<.sourceBuffer")
	return sourceBuffer
}

type AudioTrackList struct {
	*EventTarget
}

func (*AudioTrackList) GetTrackByID(id string) (a *AudioTrack) {
	js.Rewrite("$<.getTrackById($1)", id)
	return a
}

func (*AudioTrackList) Item(index uint) (a *AudioTrack) {
	js.Rewrite("$<.item($1)", index)
	return a
}

func (*AudioTrackList) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

func (*AudioTrackList) GetOnaddtrack() (addtrack *TrackEvent) {
	js.Rewrite("$<.onaddtrack")
	return addtrack
}

func (*AudioTrackList) SetOnaddtrack(addtrack *TrackEvent) {
	js.Rewrite("$<.onaddtrack = $1", addtrack)
}

func (*AudioTrackList) GetOnchange() (change *Event) {
	js.Rewrite("$<.onchange")
	return change
}

func (*AudioTrackList) SetOnchange(change *Event) {
	js.Rewrite("$<.onchange = $1", change)
}

func (*AudioTrackList) GetOnremovetrack() (removetrack *TrackEvent) {
	js.Rewrite("$<.onremovetrack")
	return removetrack
}

func (*AudioTrackList) SetOnremovetrack(removetrack *TrackEvent) {
	js.Rewrite("$<.onremovetrack = $1", removetrack)
}

type BarProp struct {
}

func (*BarProp) GetVisible() (visible bool) {
	js.Rewrite("$<.visible")
	return visible
}

type BeforeUnloadEvent struct {
	*Event
}

func (*BeforeUnloadEvent) GetReturnValue() (returnValue string) {
	js.Rewrite("$<.returnValue")
	return returnValue
}

func (*BeforeUnloadEvent) SetReturnValue(returnValue string) {
	js.Rewrite("$<.returnValue = $1", returnValue)
}

type BiquadFilterNode struct {
	*AudioNode
}

func (*BiquadFilterNode) GetFrequencyResponse(frequencyHz []float32, magResponse []float32, phaseResponse []float32) {
	js.Rewrite("$<.getFrequencyResponse($1, $2, $3)", frequencyHz, magResponse, phaseResponse)
}

func (*BiquadFilterNode) GetDetune() (detune *AudioParam) {
	js.Rewrite("$<.detune")
	return detune
}

func (*BiquadFilterNode) GetFrequency() (frequency *AudioParam) {
	js.Rewrite("$<.frequency")
	return frequency
}

func (*BiquadFilterNode) GetGain() (gain *AudioParam) {
	js.Rewrite("$<.gain")
	return gain
}

func (*BiquadFilterNode) GetQ() (Q *AudioParam) {
	js.Rewrite("$<.Q")
	return Q
}

func (*BiquadFilterNode) GetType() (kind *BiquadFilterType) {
	js.Rewrite("$<.type")
	return kind
}

func (*BiquadFilterNode) SetType(kind *BiquadFilterType) {
	js.Rewrite("$<.type = $1", kind)
}

type Blob struct {
}

func (*Blob) MsClose() {
	js.Rewrite("$<.msClose()")
}

func (*Blob) MsDetachStream() (i interface{}) {
	js.Rewrite("$<.msDetachStream()")
	return i
}

func (*Blob) Slice(start int64, end int64, contentType string) (b *Blob) {
	js.Rewrite("$<.slice($1, $2, $3)", start, end, contentType)
	return b
}

func (*Blob) GetSize() (size uint64) {
	js.Rewrite("$<.size")
	return size
}

func (*Blob) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

type Cache struct {
}

func (*Cache) Add(request *Request) {
	js.Rewrite("await $<.add($1)", request)
}

func (*Cache) AddAll(requests []*Request) {
	js.Rewrite("await $<.addAll($1)", requests)
}

func (*Cache) Delete(request *Request, options *CacheQueryOptions) (b bool) {
	js.Rewrite("await $<.delete($1, $2)", request, options)
	return b
}

func (*Cache) Keys(request *Request, options *CacheQueryOptions) (r []*Request) {
	js.Rewrite("await $<.keys($1, $2)", request, options)
	return r
}

func (*Cache) Match(request *Request, options *CacheQueryOptions) (r *Response) {
	js.Rewrite("await $<.match($1, $2)", request, options)
	return r
}

func (*Cache) MatchAll(request *Request, options *CacheQueryOptions) (r []*Response) {
	js.Rewrite("await $<.matchAll($1, $2)", request, options)
	return r
}

func (*Cache) Put(request *Request, response *Response) {
	js.Rewrite("await $<.put($1, $2)", request, response)
}

type CacheStorage struct {
}

func (*CacheStorage) Delete(cacheName string) (b bool) {
	js.Rewrite("await $<.delete($1)", cacheName)
	return b
}

func (*CacheStorage) Has(cacheName string) (b bool) {
	js.Rewrite("await $<.has($1)", cacheName)
	return b
}

func (*CacheStorage) Keys() (s []string) {
	js.Rewrite("await $<.keys()")
	return s
}

func (*CacheStorage) Match(request *Request, options *CacheQueryOptions) (i interface{}) {
	js.Rewrite("await $<.match($1, $2)", request, options)
	return i
}

func (*CacheStorage) Open(cacheName string) (c *Cache) {
	js.Rewrite("await $<.open($1)", cacheName)
	return c
}

type CanvasGradient struct {
}

func (*CanvasGradient) AddColorStop(offset float32, color string) {
	js.Rewrite("$<.addColorStop($1, $2)", offset, color)
}

type CanvasPattern struct {
}

type CanvasRenderingContext2D struct {
	*CanvasPathMethods
}

func (*CanvasRenderingContext2D) BeginPath() {
	js.Rewrite("$<.beginPath()")
}

func (*CanvasRenderingContext2D) ClearRect(x float32, y float32, w float32, h float32) {
	js.Rewrite("$<.clearRect($1, $2, $3, $4)", x, y, w, h)
}

func (*CanvasRenderingContext2D) Clip(fillRule *CanvasFillRule) {
	js.Rewrite("$<.clip($1)", fillRule)
}

func (*CanvasRenderingContext2D) CreateImageData(imageDataOrSw interface{}, sh float32) (i *ImageData) {
	js.Rewrite("$<.createImageData($1, $2)", imageDataOrSw, sh)
	return i
}

func (*CanvasRenderingContext2D) CreateLinearGradient(x0 float32, y0 float32, x1 float32, y1 float32) (c *CanvasGradient) {
	js.Rewrite("$<.createLinearGradient($1, $2, $3, $4)", x0, y0, x1, y1)
	return c
}

func (*CanvasRenderingContext2D) CreatePattern(image interface{}, repetition string) (c *CanvasPattern) {
	js.Rewrite("$<.createPattern($1, $2)", image, repetition)
	return c
}

func (*CanvasRenderingContext2D) CreateRadialGradient(x0 float32, y0 float32, r0 float32, x1 float32, y1 float32, r1 float32) (c *CanvasGradient) {
	js.Rewrite("$<.createRadialGradient($1, $2, $3, $4, $5, $6)", x0, y0, r0, x1, y1, r1)
	return c
}

func (*CanvasRenderingContext2D) DrawFocusIfNeeded(element *Element) {
	js.Rewrite("$<.drawFocusIfNeeded($1)", element)
}

func (*CanvasRenderingContext2D) DrawImage(image interface{}, offsetX float32, offsetY float32, width float32, height float32, canvasOffsetX float32, canvasOffsetY float32, canvasImageWidth float32, canvasImageHeight float32) {
	js.Rewrite("$<.drawImage($1, $2, $3, $4, $5, $6, $7, $8, $9)", image, offsetX, offsetY, width, height, canvasOffsetX, canvasOffsetY, canvasImageWidth, canvasImageHeight)
}

func (*CanvasRenderingContext2D) Fill(fillRule *CanvasFillRule) {
	js.Rewrite("$<.fill($1)", fillRule)
}

func (*CanvasRenderingContext2D) FillRect(x float32, y float32, w float32, h float32) {
	js.Rewrite("$<.fillRect($1, $2, $3, $4)", x, y, w, h)
}

func (*CanvasRenderingContext2D) FillText(text string, x float32, y float32, maxWidth float32) {
	js.Rewrite("$<.fillText($1, $2, $3, $4)", text, x, y, maxWidth)
}

func (*CanvasRenderingContext2D) GetImageData(sx float32, sy float32, sw float32, sh float32) (i *ImageData) {
	js.Rewrite("$<.getImageData($1, $2, $3, $4)", sx, sy, sw, sh)
	return i
}

func (*CanvasRenderingContext2D) GetLineDash() (f []float32) {
	js.Rewrite("$<.getLineDash()")
	return f
}

func (*CanvasRenderingContext2D) IsPointInPath(x float32, y float32, fillRule *CanvasFillRule) (b bool) {
	js.Rewrite("$<.isPointInPath($1, $2, $3)", x, y, fillRule)
	return b
}

func (*CanvasRenderingContext2D) MeasureText(text string) (t *TextMetrics) {
	js.Rewrite("$<.measureText($1)", text)
	return t
}

func (*CanvasRenderingContext2D) PutImageData(imagedata *ImageData, dx float32, dy float32, dirtyX float32, dirtyY float32, dirtyWidth float32, dirtyHeight float32) {
	js.Rewrite("$<.putImageData($1, $2, $3, $4, $5, $6, $7)", imagedata, dx, dy, dirtyX, dirtyY, dirtyWidth, dirtyHeight)
}

func (*CanvasRenderingContext2D) Restore() {
	js.Rewrite("$<.restore()")
}

func (*CanvasRenderingContext2D) Rotate(angle float32) {
	js.Rewrite("$<.rotate($1)", angle)
}

func (*CanvasRenderingContext2D) Save() {
	js.Rewrite("$<.save()")
}

func (*CanvasRenderingContext2D) Scale(x float32, y float32) {
	js.Rewrite("$<.scale($1, $2)", x, y)
}

func (*CanvasRenderingContext2D) SetLineDash(segments []float32) {
	js.Rewrite("$<.setLineDash($1)", segments)
}

func (*CanvasRenderingContext2D) SetTransform(m11 float32, m12 float32, m21 float32, m22 float32, dx float32, dy float32) {
	js.Rewrite("$<.setTransform($1, $2, $3, $4, $5, $6)", m11, m12, m21, m22, dx, dy)
}

func (*CanvasRenderingContext2D) Stroke(path *Path2D) {
	js.Rewrite("$<.stroke($1)", path)
}

func (*CanvasRenderingContext2D) StrokeRect(x float32, y float32, w float32, h float32) {
	js.Rewrite("$<.strokeRect($1, $2, $3, $4)", x, y, w, h)
}

func (*CanvasRenderingContext2D) StrokeText(text string, x float32, y float32, maxWidth float32) {
	js.Rewrite("$<.strokeText($1, $2, $3, $4)", text, x, y, maxWidth)
}

func (*CanvasRenderingContext2D) Transform(m11 float32, m12 float32, m21 float32, m22 float32, dx float32, dy float32) {
	js.Rewrite("$<.transform($1, $2, $3, $4, $5, $6)", m11, m12, m21, m22, dx, dy)
}

func (*CanvasRenderingContext2D) Translate(x float32, y float32) {
	js.Rewrite("$<.translate($1, $2)", x, y)
}

func (*CanvasRenderingContext2D) GetCanvas() (canvas *HTMLCanvasElement) {
	js.Rewrite("$<.canvas")
	return canvas
}

func (*CanvasRenderingContext2D) GetFillStyle() (fillStyle interface{}) {
	js.Rewrite("$<.fillStyle")
	return fillStyle
}

func (*CanvasRenderingContext2D) SetFillStyle(fillStyle interface{}) {
	js.Rewrite("$<.fillStyle = $1", fillStyle)
}

func (*CanvasRenderingContext2D) GetFont() (font string) {
	js.Rewrite("$<.font")
	return font
}

func (*CanvasRenderingContext2D) SetFont(font string) {
	js.Rewrite("$<.font = $1", font)
}

func (*CanvasRenderingContext2D) GetGlobalAlpha() (globalAlpha float32) {
	js.Rewrite("$<.globalAlpha")
	return globalAlpha
}

func (*CanvasRenderingContext2D) SetGlobalAlpha(globalAlpha float32) {
	js.Rewrite("$<.globalAlpha = $1", globalAlpha)
}

func (*CanvasRenderingContext2D) GetGlobalCompositeOperation() (globalCompositeOperation string) {
	js.Rewrite("$<.globalCompositeOperation")
	return globalCompositeOperation
}

func (*CanvasRenderingContext2D) SetGlobalCompositeOperation(globalCompositeOperation string) {
	js.Rewrite("$<.globalCompositeOperation = $1", globalCompositeOperation)
}

func (*CanvasRenderingContext2D) GetImageSmoothingEnabled() (imageSmoothingEnabled bool) {
	js.Rewrite("$<.imageSmoothingEnabled")
	return imageSmoothingEnabled
}

func (*CanvasRenderingContext2D) SetImageSmoothingEnabled(imageSmoothingEnabled bool) {
	js.Rewrite("$<.imageSmoothingEnabled = $1", imageSmoothingEnabled)
}

func (*CanvasRenderingContext2D) GetLineCap() (lineCap string) {
	js.Rewrite("$<.lineCap")
	return lineCap
}

func (*CanvasRenderingContext2D) SetLineCap(lineCap string) {
	js.Rewrite("$<.lineCap = $1", lineCap)
}

func (*CanvasRenderingContext2D) GetLineDashOffset() (lineDashOffset float32) {
	js.Rewrite("$<.lineDashOffset")
	return lineDashOffset
}

func (*CanvasRenderingContext2D) SetLineDashOffset(lineDashOffset float32) {
	js.Rewrite("$<.lineDashOffset = $1", lineDashOffset)
}

func (*CanvasRenderingContext2D) GetLineJoin() (lineJoin string) {
	js.Rewrite("$<.lineJoin")
	return lineJoin
}

func (*CanvasRenderingContext2D) SetLineJoin(lineJoin string) {
	js.Rewrite("$<.lineJoin = $1", lineJoin)
}

func (*CanvasRenderingContext2D) GetLineWidth() (lineWidth float32) {
	js.Rewrite("$<.lineWidth")
	return lineWidth
}

func (*CanvasRenderingContext2D) SetLineWidth(lineWidth float32) {
	js.Rewrite("$<.lineWidth = $1", lineWidth)
}

func (*CanvasRenderingContext2D) GetMiterLimit() (miterLimit float32) {
	js.Rewrite("$<.miterLimit")
	return miterLimit
}

func (*CanvasRenderingContext2D) SetMiterLimit(miterLimit float32) {
	js.Rewrite("$<.miterLimit = $1", miterLimit)
}

func (*CanvasRenderingContext2D) GetMsFillRule() (msFillRule *CanvasFillRule) {
	js.Rewrite("$<.msFillRule")
	return msFillRule
}

func (*CanvasRenderingContext2D) SetMsFillRule(msFillRule *CanvasFillRule) {
	js.Rewrite("$<.msFillRule = $1", msFillRule)
}

func (*CanvasRenderingContext2D) GetShadowBlur() (shadowBlur float32) {
	js.Rewrite("$<.shadowBlur")
	return shadowBlur
}

func (*CanvasRenderingContext2D) SetShadowBlur(shadowBlur float32) {
	js.Rewrite("$<.shadowBlur = $1", shadowBlur)
}

func (*CanvasRenderingContext2D) GetShadowColor() (shadowColor string) {
	js.Rewrite("$<.shadowColor")
	return shadowColor
}

func (*CanvasRenderingContext2D) SetShadowColor(shadowColor string) {
	js.Rewrite("$<.shadowColor = $1", shadowColor)
}

func (*CanvasRenderingContext2D) GetShadowOffsetX() (shadowOffsetX float32) {
	js.Rewrite("$<.shadowOffsetX")
	return shadowOffsetX
}

func (*CanvasRenderingContext2D) SetShadowOffsetX(shadowOffsetX float32) {
	js.Rewrite("$<.shadowOffsetX = $1", shadowOffsetX)
}

func (*CanvasRenderingContext2D) GetShadowOffsetY() (shadowOffsetY float32) {
	js.Rewrite("$<.shadowOffsetY")
	return shadowOffsetY
}

func (*CanvasRenderingContext2D) SetShadowOffsetY(shadowOffsetY float32) {
	js.Rewrite("$<.shadowOffsetY = $1", shadowOffsetY)
}

func (*CanvasRenderingContext2D) GetStrokeStyle() (strokeStyle interface{}) {
	js.Rewrite("$<.strokeStyle")
	return strokeStyle
}

func (*CanvasRenderingContext2D) SetStrokeStyle(strokeStyle interface{}) {
	js.Rewrite("$<.strokeStyle = $1", strokeStyle)
}

func (*CanvasRenderingContext2D) GetTextAlign() (textAlign string) {
	js.Rewrite("$<.textAlign")
	return textAlign
}

func (*CanvasRenderingContext2D) SetTextAlign(textAlign string) {
	js.Rewrite("$<.textAlign = $1", textAlign)
}

func (*CanvasRenderingContext2D) GetTextBaseline() (textBaseline string) {
	js.Rewrite("$<.textBaseline")
	return textBaseline
}

func (*CanvasRenderingContext2D) SetTextBaseline(textBaseline string) {
	js.Rewrite("$<.textBaseline = $1", textBaseline)
}

type CDATASection struct {
	*Text
}

type ChannelMergerNode struct {
	*AudioNode
}

type ChannelSplitterNode struct {
	*AudioNode
}

type CharacterData struct {
	*Node
	*ChildNode
}

func (*CharacterData) AppendData(arg string) {
	js.Rewrite("$<.appendData($1)", arg)
}

func (*CharacterData) DeleteData(offset uint, count uint) {
	js.Rewrite("$<.deleteData($1, $2)", offset, count)
}

func (*CharacterData) InsertData(offset uint, arg string) {
	js.Rewrite("$<.insertData($1, $2)", offset, arg)
}

func (*CharacterData) ReplaceData(offset uint, count uint, arg string) {
	js.Rewrite("$<.replaceData($1, $2, $3)", offset, count, arg)
}

func (*CharacterData) SubstringData(offset uint, count uint) (s string) {
	js.Rewrite("$<.substringData($1, $2)", offset, count)
	return s
}

func (*CharacterData) GetData() (data string) {
	js.Rewrite("$<.data")
	return data
}

func (*CharacterData) SetData(data string) {
	js.Rewrite("$<.data = $1", data)
}

func (*CharacterData) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

type ClientRect struct {
}

func (*ClientRect) GetBottom() (bottom int) {
	js.Rewrite("$<.bottom")
	return bottom
}

func (*ClientRect) SetBottom(bottom int) {
	js.Rewrite("$<.bottom = $1", bottom)
}

func (*ClientRect) GetHeight() (height float32) {
	js.Rewrite("$<.height")
	return height
}

func (*ClientRect) GetLeft() (left int) {
	js.Rewrite("$<.left")
	return left
}

func (*ClientRect) SetLeft(left int) {
	js.Rewrite("$<.left = $1", left)
}

func (*ClientRect) GetRight() (right int) {
	js.Rewrite("$<.right")
	return right
}

func (*ClientRect) SetRight(right int) {
	js.Rewrite("$<.right = $1", right)
}

func (*ClientRect) GetTop() (top int) {
	js.Rewrite("$<.top")
	return top
}

func (*ClientRect) SetTop(top int) {
	js.Rewrite("$<.top = $1", top)
}

func (*ClientRect) GetWidth() (width float32) {
	js.Rewrite("$<.width")
	return width
}

type ClientRectList struct {
}

func (*ClientRectList) Item(index uint) (c *ClientRect) {
	js.Rewrite("$<.item($1)", index)
	return c
}

func (*ClientRectList) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

type ClipboardEvent struct {
	*Event
}

func (*ClipboardEvent) GetClipboardData() (clipboardData *DataTransfer) {
	js.Rewrite("$<.clipboardData")
	return clipboardData
}

type CloseEvent struct {
	*Event
}

func (*CloseEvent) InitCloseEvent(typeArg string, canBubbleArg bool, cancelableArg bool, wasCleanArg bool, codeArg uint8, reasonArg string) {
	js.Rewrite("$<.initCloseEvent($1, $2, $3, $4, $5, $6)", typeArg, canBubbleArg, cancelableArg, wasCleanArg, codeArg, reasonArg)
}

func (*CloseEvent) GetCode() (code uint8) {
	js.Rewrite("$<.code")
	return code
}

func (*CloseEvent) GetReason() (reason string) {
	js.Rewrite("$<.reason")
	return reason
}

func (*CloseEvent) GetWasClean() (wasClean bool) {
	js.Rewrite("$<.wasClean")
	return wasClean
}

type Comment struct {
	*CharacterData
}

func (*Comment) GetText() (text string) {
	js.Rewrite("$<.text")
	return text
}

func (*Comment) SetText(text string) {
	js.Rewrite("$<.text = $1", text)
}

type CompositionEvent struct {
	*UIEvent
}

func (*CompositionEvent) InitCompositionEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, dataArg string, locale string) {
	js.Rewrite("$<.initCompositionEvent($1, $2, $3, $4, $5, $6)", typeArg, canBubbleArg, cancelableArg, viewArg, dataArg, locale)
}

func (*CompositionEvent) GetData() (data string) {
	js.Rewrite("$<.data")
	return data
}

func (*CompositionEvent) GetLocale() (locale string) {
	js.Rewrite("$<.locale")
	return locale
}

type Console struct {
}

func (*Console) Assert(test bool, message string, optionalParams interface{}) {
	js.Rewrite("$<.assert($1, $2, $3)", test, message, optionalParams)
}

func (*Console) Clear() {
	js.Rewrite("$<.clear()")
}

func (*Console) Count(countTitle string) {
	js.Rewrite("$<.count($1)", countTitle)
}

func (*Console) Debug(message string, optionalParams interface{}) {
	js.Rewrite("$<.debug($1, $2)", message, optionalParams)
}

func (*Console) Dir(value interface{}, optionalParams interface{}) {
	js.Rewrite("$<.dir($1, $2)", value, optionalParams)
}

func (*Console) Dirxml(value interface{}) {
	js.Rewrite("$<.dirxml($1)", value)
}

func (*Console) Error(message string, optionalParams interface{}) {
	js.Rewrite("$<.error($1, $2)", message, optionalParams)
}

func (*Console) Exception(message string, optionalParams interface{}) {
	js.Rewrite("$<.exception($1, $2)", message, optionalParams)
}

func (*Console) Group(groupTitle string) {
	js.Rewrite("$<.group($1)", groupTitle)
}

func (*Console) GroupCollapsed(groupTitle string) {
	js.Rewrite("$<.groupCollapsed($1)", groupTitle)
}

func (*Console) GroupEnd() {
	js.Rewrite("$<.groupEnd()")
}

func (*Console) Info(message string, optionalParams interface{}) {
	js.Rewrite("$<.info($1, $2)", message, optionalParams)
}

func (*Console) Log(message string, optionalParams interface{}) {
	js.Rewrite("$<.log($1, $2)", message, optionalParams)
}

func (*Console) MsIsIndependentlyComposed(element *Element) (b bool) {
	js.Rewrite("$<.msIsIndependentlyComposed($1)", element)
	return b
}

func (*Console) Profile(reportName string) {
	js.Rewrite("$<.profile($1)", reportName)
}

func (*Console) ProfileEnd() {
	js.Rewrite("$<.profileEnd()")
}

func (*Console) Select(element *Element) {
	js.Rewrite("$<.select($1)", element)
}

func (*Console) Table(data interface{}) {
	js.Rewrite("$<.table($1)", data)
}

func (*Console) Time(timerName string) {
	js.Rewrite("$<.time($1)", timerName)
}

func (*Console) TimeEnd(timerName string) {
	js.Rewrite("$<.timeEnd($1)", timerName)
}

func (*Console) Trace() {
	js.Rewrite("$<.trace()")
}

func (*Console) Warn(message string, optionalParams interface{}) {
	js.Rewrite("$<.warn($1, $2)", message, optionalParams)
}

type ConvolverNode struct {
	*AudioNode
}

func (*ConvolverNode) GetBuffer() (buffer *AudioBuffer) {
	js.Rewrite("$<.buffer")
	return buffer
}

func (*ConvolverNode) SetBuffer(buffer *AudioBuffer) {
	js.Rewrite("$<.buffer = $1", buffer)
}

func (*ConvolverNode) GetNormalize() (normalize bool) {
	js.Rewrite("$<.normalize")
	return normalize
}

func (*ConvolverNode) SetNormalize(normalize bool) {
	js.Rewrite("$<.normalize = $1", normalize)
}

type Coordinates struct {
}

func (*Coordinates) GetAccuracy() (accuracy float32) {
	js.Rewrite("$<.accuracy")
	return accuracy
}

func (*Coordinates) GetAltitude() (altitude float32) {
	js.Rewrite("$<.altitude")
	return altitude
}

func (*Coordinates) GetAltitudeAccuracy() (altitudeAccuracy float32) {
	js.Rewrite("$<.altitudeAccuracy")
	return altitudeAccuracy
}

func (*Coordinates) GetHeading() (heading float32) {
	js.Rewrite("$<.heading")
	return heading
}

func (*Coordinates) GetLatitude() (latitude float32) {
	js.Rewrite("$<.latitude")
	return latitude
}

func (*Coordinates) GetLongitude() (longitude float32) {
	js.Rewrite("$<.longitude")
	return longitude
}

func (*Coordinates) GetSpeed() (speed float32) {
	js.Rewrite("$<.speed")
	return speed
}

type Crypto struct {
	*RandomSource
}

func (*Crypto) GetSubtle() (subtle *SubtleCrypto) {
	js.Rewrite("$<.subtle")
	return subtle
}

type CryptoKey struct {
}

func (*CryptoKey) GetAlgorithm() (algorithm *KeyAlgorithm) {
	js.Rewrite("$<.algorithm")
	return algorithm
}

func (*CryptoKey) GetExtractable() (extractable bool) {
	js.Rewrite("$<.extractable")
	return extractable
}

func (*CryptoKey) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*CryptoKey) GetUsages() (usages []string) {
	js.Rewrite("$<.usages")
	return usages
}

type CryptoKeyPair struct {
}

func (*CryptoKeyPair) GetPrivateKey() (privateKey *CryptoKey) {
	js.Rewrite("$<.privateKey")
	return privateKey
}

func (*CryptoKeyPair) SetPrivateKey(privateKey *CryptoKey) {
	js.Rewrite("$<.privateKey = $1", privateKey)
}

func (*CryptoKeyPair) GetPublicKey() (publicKey *CryptoKey) {
	js.Rewrite("$<.publicKey")
	return publicKey
}

func (*CryptoKeyPair) SetPublicKey(publicKey *CryptoKey) {
	js.Rewrite("$<.publicKey = $1", publicKey)
}

type CSS struct {
}

func (*CSS) Supports(property string, value string) (b bool) {
	js.Rewrite("$<.supports($1, $2)", property, value)
	return b
}

type CSSConditionRule struct {
	*CSSGroupingRule
}

func (*CSSConditionRule) GetConditionText() (conditionText string) {
	js.Rewrite("$<.conditionText")
	return conditionText
}

func (*CSSConditionRule) SetConditionText(conditionText string) {
	js.Rewrite("$<.conditionText = $1", conditionText)
}

type CSSFontFaceRule struct {
	*CSSRule
}

func (*CSSFontFaceRule) GetStyle() (style *CSSStyleDeclaration) {
	js.Rewrite("$<.style")
	return style
}

type CSSGroupingRule struct {
	*CSSRule
}

func (*CSSGroupingRule) DeleteRule(index uint) {
	js.Rewrite("$<.deleteRule($1)", index)
}

func (*CSSGroupingRule) InsertRule(rule string, index uint) (u uint) {
	js.Rewrite("$<.insertRule($1, $2)", rule, index)
	return u
}

func (*CSSGroupingRule) GetCSSRules() (cssRules *CSSRuleList) {
	js.Rewrite("$<.cssRules")
	return cssRules
}

type CSSImportRule struct {
	*CSSRule
}

func (*CSSImportRule) GetHref() (href string) {
	js.Rewrite("$<.href")
	return href
}

func (*CSSImportRule) GetMedia() (media *MediaList) {
	js.Rewrite("$<.media")
	return media
}

func (*CSSImportRule) GetStyleSheet() (styleSheet *CSSStyleSheet) {
	js.Rewrite("$<.styleSheet")
	return styleSheet
}

type CSSKeyframeRule struct {
	*CSSRule
}

func (*CSSKeyframeRule) GetKeyText() (keyText string) {
	js.Rewrite("$<.keyText")
	return keyText
}

func (*CSSKeyframeRule) SetKeyText(keyText string) {
	js.Rewrite("$<.keyText = $1", keyText)
}

func (*CSSKeyframeRule) GetStyle() (style *CSSStyleDeclaration) {
	js.Rewrite("$<.style")
	return style
}

type CSSKeyframesRule struct {
	*CSSRule
}

func (*CSSKeyframesRule) AppendRule(rule string) {
	js.Rewrite("$<.appendRule($1)", rule)
}

func (*CSSKeyframesRule) DeleteRule(rule string) {
	js.Rewrite("$<.deleteRule($1)", rule)
}

func (*CSSKeyframesRule) FindRule(rule string) (c *CSSKeyframeRule) {
	js.Rewrite("$<.findRule($1)", rule)
	return c
}

func (*CSSKeyframesRule) GetCSSRules() (cssRules *CSSRuleList) {
	js.Rewrite("$<.cssRules")
	return cssRules
}

func (*CSSKeyframesRule) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*CSSKeyframesRule) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

type CSSMediaRule struct {
	*CSSConditionRule
}

func (*CSSMediaRule) GetMedia() (media *MediaList) {
	js.Rewrite("$<.media")
	return media
}

type CSSNamespaceRule struct {
	*CSSRule
}

func (*CSSNamespaceRule) GetNamespaceURI() (namespaceURI string) {
	js.Rewrite("$<.namespaceURI")
	return namespaceURI
}

func (*CSSNamespaceRule) GetPrefix() (prefix string) {
	js.Rewrite("$<.prefix")
	return prefix
}

type CSSPageRule struct {
	*CSSRule
}

func (*CSSPageRule) GetPseudoClass() (pseudoClass string) {
	js.Rewrite("$<.pseudoClass")
	return pseudoClass
}

func (*CSSPageRule) GetSelector() (selector string) {
	js.Rewrite("$<.selector")
	return selector
}

func (*CSSPageRule) GetSelectorText() (selectorText string) {
	js.Rewrite("$<.selectorText")
	return selectorText
}

func (*CSSPageRule) SetSelectorText(selectorText string) {
	js.Rewrite("$<.selectorText = $1", selectorText)
}

func (*CSSPageRule) GetStyle() (style *CSSStyleDeclaration) {
	js.Rewrite("$<.style")
	return style
}

type CSSRule struct {
}

func (*CSSRule) GetCSSText() (cssText string) {
	js.Rewrite("$<.cssText")
	return cssText
}

func (*CSSRule) SetCSSText(cssText string) {
	js.Rewrite("$<.cssText = $1", cssText)
}

func (*CSSRule) GetParentRule() (parentRule *CSSRule) {
	js.Rewrite("$<.parentRule")
	return parentRule
}

func (*CSSRule) GetParentStyleSheet() (parentStyleSheet *CSSStyleSheet) {
	js.Rewrite("$<.parentStyleSheet")
	return parentStyleSheet
}

func (*CSSRule) GetType() (kind uint8) {
	js.Rewrite("$<.type")
	return kind
}

type CSSRuleList struct {
}

func (*CSSRuleList) Item(index uint) (c *CSSRule) {
	js.Rewrite("$<.item($1)", index)
	return c
}

func (*CSSRuleList) GetLength() (length int) {
	js.Rewrite("$<.length")
	return length
}

type CSSStyleDeclaration struct {
}

func (*CSSStyleDeclaration) GetPropertyPriority(propertyName string) (s string) {
	js.Rewrite("$<.getPropertyPriority($1)", propertyName)
	return s
}

func (*CSSStyleDeclaration) GetPropertyValue(propertyName string) (s string) {
	js.Rewrite("$<.getPropertyValue($1)", propertyName)
	return s
}

func (*CSSStyleDeclaration) Item(index uint) (s string) {
	js.Rewrite("$<.item($1)", index)
	return s
}

func (*CSSStyleDeclaration) RemoveProperty(propertyName string) (s string) {
	js.Rewrite("$<.removeProperty($1)", propertyName)
	return s
}

func (*CSSStyleDeclaration) SetProperty(propertyName string, value string, priority string) {
	js.Rewrite("$<.setProperty($1, $2, $3)", propertyName, value, priority)
}

func (*CSSStyleDeclaration) GetAlignContent() (alignContent string) {
	js.Rewrite("$<.alignContent")
	return alignContent
}

func (*CSSStyleDeclaration) SetAlignContent(alignContent string) {
	js.Rewrite("$<.alignContent = $1", alignContent)
}

func (*CSSStyleDeclaration) GetAlignItems() (alignItems string) {
	js.Rewrite("$<.alignItems")
	return alignItems
}

func (*CSSStyleDeclaration) SetAlignItems(alignItems string) {
	js.Rewrite("$<.alignItems = $1", alignItems)
}

func (*CSSStyleDeclaration) GetAlignmentBaseline() (alignmentBaseline string) {
	js.Rewrite("$<.alignmentBaseline")
	return alignmentBaseline
}

func (*CSSStyleDeclaration) SetAlignmentBaseline(alignmentBaseline string) {
	js.Rewrite("$<.alignmentBaseline = $1", alignmentBaseline)
}

func (*CSSStyleDeclaration) GetAlignSelf() (alignSelf string) {
	js.Rewrite("$<.alignSelf")
	return alignSelf
}

func (*CSSStyleDeclaration) SetAlignSelf(alignSelf string) {
	js.Rewrite("$<.alignSelf = $1", alignSelf)
}

func (*CSSStyleDeclaration) GetAnimation() (animation string) {
	js.Rewrite("$<.animation")
	return animation
}

func (*CSSStyleDeclaration) SetAnimation(animation string) {
	js.Rewrite("$<.animation = $1", animation)
}

func (*CSSStyleDeclaration) GetAnimationDelay() (animationDelay string) {
	js.Rewrite("$<.animationDelay")
	return animationDelay
}

func (*CSSStyleDeclaration) SetAnimationDelay(animationDelay string) {
	js.Rewrite("$<.animationDelay = $1", animationDelay)
}

func (*CSSStyleDeclaration) GetAnimationDirection() (animationDirection string) {
	js.Rewrite("$<.animationDirection")
	return animationDirection
}

func (*CSSStyleDeclaration) SetAnimationDirection(animationDirection string) {
	js.Rewrite("$<.animationDirection = $1", animationDirection)
}

func (*CSSStyleDeclaration) GetAnimationDuration() (animationDuration string) {
	js.Rewrite("$<.animationDuration")
	return animationDuration
}

func (*CSSStyleDeclaration) SetAnimationDuration(animationDuration string) {
	js.Rewrite("$<.animationDuration = $1", animationDuration)
}

func (*CSSStyleDeclaration) GetAnimationFillMode() (animationFillMode string) {
	js.Rewrite("$<.animationFillMode")
	return animationFillMode
}

func (*CSSStyleDeclaration) SetAnimationFillMode(animationFillMode string) {
	js.Rewrite("$<.animationFillMode = $1", animationFillMode)
}

func (*CSSStyleDeclaration) GetAnimationIterationCount() (animationIterationCount string) {
	js.Rewrite("$<.animationIterationCount")
	return animationIterationCount
}

func (*CSSStyleDeclaration) SetAnimationIterationCount(animationIterationCount string) {
	js.Rewrite("$<.animationIterationCount = $1", animationIterationCount)
}

func (*CSSStyleDeclaration) GetAnimationName() (animationName string) {
	js.Rewrite("$<.animationName")
	return animationName
}

func (*CSSStyleDeclaration) SetAnimationName(animationName string) {
	js.Rewrite("$<.animationName = $1", animationName)
}

func (*CSSStyleDeclaration) GetAnimationPlayState() (animationPlayState string) {
	js.Rewrite("$<.animationPlayState")
	return animationPlayState
}

func (*CSSStyleDeclaration) SetAnimationPlayState(animationPlayState string) {
	js.Rewrite("$<.animationPlayState = $1", animationPlayState)
}

func (*CSSStyleDeclaration) GetAnimationTimingFunction() (animationTimingFunction string) {
	js.Rewrite("$<.animationTimingFunction")
	return animationTimingFunction
}

func (*CSSStyleDeclaration) SetAnimationTimingFunction(animationTimingFunction string) {
	js.Rewrite("$<.animationTimingFunction = $1", animationTimingFunction)
}

func (*CSSStyleDeclaration) GetBackfaceVisibility() (backfaceVisibility string) {
	js.Rewrite("$<.backfaceVisibility")
	return backfaceVisibility
}

func (*CSSStyleDeclaration) SetBackfaceVisibility(backfaceVisibility string) {
	js.Rewrite("$<.backfaceVisibility = $1", backfaceVisibility)
}

func (*CSSStyleDeclaration) GetBackground() (background string) {
	js.Rewrite("$<.background")
	return background
}

func (*CSSStyleDeclaration) SetBackground(background string) {
	js.Rewrite("$<.background = $1", background)
}

func (*CSSStyleDeclaration) GetBackgroundAttachment() (backgroundAttachment string) {
	js.Rewrite("$<.backgroundAttachment")
	return backgroundAttachment
}

func (*CSSStyleDeclaration) SetBackgroundAttachment(backgroundAttachment string) {
	js.Rewrite("$<.backgroundAttachment = $1", backgroundAttachment)
}

func (*CSSStyleDeclaration) GetBackgroundClip() (backgroundClip string) {
	js.Rewrite("$<.backgroundClip")
	return backgroundClip
}

func (*CSSStyleDeclaration) SetBackgroundClip(backgroundClip string) {
	js.Rewrite("$<.backgroundClip = $1", backgroundClip)
}

func (*CSSStyleDeclaration) GetBackgroundColor() (backgroundColor string) {
	js.Rewrite("$<.backgroundColor")
	return backgroundColor
}

func (*CSSStyleDeclaration) SetBackgroundColor(backgroundColor string) {
	js.Rewrite("$<.backgroundColor = $1", backgroundColor)
}

func (*CSSStyleDeclaration) GetBackgroundImage() (backgroundImage string) {
	js.Rewrite("$<.backgroundImage")
	return backgroundImage
}

func (*CSSStyleDeclaration) SetBackgroundImage(backgroundImage string) {
	js.Rewrite("$<.backgroundImage = $1", backgroundImage)
}

func (*CSSStyleDeclaration) GetBackgroundOrigin() (backgroundOrigin string) {
	js.Rewrite("$<.backgroundOrigin")
	return backgroundOrigin
}

func (*CSSStyleDeclaration) SetBackgroundOrigin(backgroundOrigin string) {
	js.Rewrite("$<.backgroundOrigin = $1", backgroundOrigin)
}

func (*CSSStyleDeclaration) GetBackgroundPosition() (backgroundPosition string) {
	js.Rewrite("$<.backgroundPosition")
	return backgroundPosition
}

func (*CSSStyleDeclaration) SetBackgroundPosition(backgroundPosition string) {
	js.Rewrite("$<.backgroundPosition = $1", backgroundPosition)
}

func (*CSSStyleDeclaration) GetBackgroundPositionX() (backgroundPositionX string) {
	js.Rewrite("$<.backgroundPositionX")
	return backgroundPositionX
}

func (*CSSStyleDeclaration) SetBackgroundPositionX(backgroundPositionX string) {
	js.Rewrite("$<.backgroundPositionX = $1", backgroundPositionX)
}

func (*CSSStyleDeclaration) GetBackgroundPositionY() (backgroundPositionY string) {
	js.Rewrite("$<.backgroundPositionY")
	return backgroundPositionY
}

func (*CSSStyleDeclaration) SetBackgroundPositionY(backgroundPositionY string) {
	js.Rewrite("$<.backgroundPositionY = $1", backgroundPositionY)
}

func (*CSSStyleDeclaration) GetBackgroundRepeat() (backgroundRepeat string) {
	js.Rewrite("$<.backgroundRepeat")
	return backgroundRepeat
}

func (*CSSStyleDeclaration) SetBackgroundRepeat(backgroundRepeat string) {
	js.Rewrite("$<.backgroundRepeat = $1", backgroundRepeat)
}

func (*CSSStyleDeclaration) GetBackgroundSize() (backgroundSize string) {
	js.Rewrite("$<.backgroundSize")
	return backgroundSize
}

func (*CSSStyleDeclaration) SetBackgroundSize(backgroundSize string) {
	js.Rewrite("$<.backgroundSize = $1", backgroundSize)
}

func (*CSSStyleDeclaration) GetBaselineShift() (baselineShift string) {
	js.Rewrite("$<.baselineShift")
	return baselineShift
}

func (*CSSStyleDeclaration) SetBaselineShift(baselineShift string) {
	js.Rewrite("$<.baselineShift = $1", baselineShift)
}

func (*CSSStyleDeclaration) GetBorder() (border string) {
	js.Rewrite("$<.border")
	return border
}

func (*CSSStyleDeclaration) SetBorder(border string) {
	js.Rewrite("$<.border = $1", border)
}

func (*CSSStyleDeclaration) GetBorderBottom() (borderBottom string) {
	js.Rewrite("$<.borderBottom")
	return borderBottom
}

func (*CSSStyleDeclaration) SetBorderBottom(borderBottom string) {
	js.Rewrite("$<.borderBottom = $1", borderBottom)
}

func (*CSSStyleDeclaration) GetBorderBottomColor() (borderBottomColor string) {
	js.Rewrite("$<.borderBottomColor")
	return borderBottomColor
}

func (*CSSStyleDeclaration) SetBorderBottomColor(borderBottomColor string) {
	js.Rewrite("$<.borderBottomColor = $1", borderBottomColor)
}

func (*CSSStyleDeclaration) GetBorderBottomLeftRadius() (borderBottomLeftRadius string) {
	js.Rewrite("$<.borderBottomLeftRadius")
	return borderBottomLeftRadius
}

func (*CSSStyleDeclaration) SetBorderBottomLeftRadius(borderBottomLeftRadius string) {
	js.Rewrite("$<.borderBottomLeftRadius = $1", borderBottomLeftRadius)
}

func (*CSSStyleDeclaration) GetBorderBottomRightRadius() (borderBottomRightRadius string) {
	js.Rewrite("$<.borderBottomRightRadius")
	return borderBottomRightRadius
}

func (*CSSStyleDeclaration) SetBorderBottomRightRadius(borderBottomRightRadius string) {
	js.Rewrite("$<.borderBottomRightRadius = $1", borderBottomRightRadius)
}

func (*CSSStyleDeclaration) GetBorderBottomStyle() (borderBottomStyle string) {
	js.Rewrite("$<.borderBottomStyle")
	return borderBottomStyle
}

func (*CSSStyleDeclaration) SetBorderBottomStyle(borderBottomStyle string) {
	js.Rewrite("$<.borderBottomStyle = $1", borderBottomStyle)
}

func (*CSSStyleDeclaration) GetBorderBottomWidth() (borderBottomWidth string) {
	js.Rewrite("$<.borderBottomWidth")
	return borderBottomWidth
}

func (*CSSStyleDeclaration) SetBorderBottomWidth(borderBottomWidth string) {
	js.Rewrite("$<.borderBottomWidth = $1", borderBottomWidth)
}

func (*CSSStyleDeclaration) GetBorderCollapse() (borderCollapse string) {
	js.Rewrite("$<.borderCollapse")
	return borderCollapse
}

func (*CSSStyleDeclaration) SetBorderCollapse(borderCollapse string) {
	js.Rewrite("$<.borderCollapse = $1", borderCollapse)
}

func (*CSSStyleDeclaration) GetBorderColor() (borderColor string) {
	js.Rewrite("$<.borderColor")
	return borderColor
}

func (*CSSStyleDeclaration) SetBorderColor(borderColor string) {
	js.Rewrite("$<.borderColor = $1", borderColor)
}

func (*CSSStyleDeclaration) GetBorderImage() (borderImage string) {
	js.Rewrite("$<.borderImage")
	return borderImage
}

func (*CSSStyleDeclaration) SetBorderImage(borderImage string) {
	js.Rewrite("$<.borderImage = $1", borderImage)
}

func (*CSSStyleDeclaration) GetBorderImageOutset() (borderImageOutset string) {
	js.Rewrite("$<.borderImageOutset")
	return borderImageOutset
}

func (*CSSStyleDeclaration) SetBorderImageOutset(borderImageOutset string) {
	js.Rewrite("$<.borderImageOutset = $1", borderImageOutset)
}

func (*CSSStyleDeclaration) GetBorderImageRepeat() (borderImageRepeat string) {
	js.Rewrite("$<.borderImageRepeat")
	return borderImageRepeat
}

func (*CSSStyleDeclaration) SetBorderImageRepeat(borderImageRepeat string) {
	js.Rewrite("$<.borderImageRepeat = $1", borderImageRepeat)
}

func (*CSSStyleDeclaration) GetBorderImageSlice() (borderImageSlice string) {
	js.Rewrite("$<.borderImageSlice")
	return borderImageSlice
}

func (*CSSStyleDeclaration) SetBorderImageSlice(borderImageSlice string) {
	js.Rewrite("$<.borderImageSlice = $1", borderImageSlice)
}

func (*CSSStyleDeclaration) GetBorderImageSource() (borderImageSource string) {
	js.Rewrite("$<.borderImageSource")
	return borderImageSource
}

func (*CSSStyleDeclaration) SetBorderImageSource(borderImageSource string) {
	js.Rewrite("$<.borderImageSource = $1", borderImageSource)
}

func (*CSSStyleDeclaration) GetBorderImageWidth() (borderImageWidth string) {
	js.Rewrite("$<.borderImageWidth")
	return borderImageWidth
}

func (*CSSStyleDeclaration) SetBorderImageWidth(borderImageWidth string) {
	js.Rewrite("$<.borderImageWidth = $1", borderImageWidth)
}

func (*CSSStyleDeclaration) GetBorderLeft() (borderLeft string) {
	js.Rewrite("$<.borderLeft")
	return borderLeft
}

func (*CSSStyleDeclaration) SetBorderLeft(borderLeft string) {
	js.Rewrite("$<.borderLeft = $1", borderLeft)
}

func (*CSSStyleDeclaration) GetBorderLeftColor() (borderLeftColor string) {
	js.Rewrite("$<.borderLeftColor")
	return borderLeftColor
}

func (*CSSStyleDeclaration) SetBorderLeftColor(borderLeftColor string) {
	js.Rewrite("$<.borderLeftColor = $1", borderLeftColor)
}

func (*CSSStyleDeclaration) GetBorderLeftStyle() (borderLeftStyle string) {
	js.Rewrite("$<.borderLeftStyle")
	return borderLeftStyle
}

func (*CSSStyleDeclaration) SetBorderLeftStyle(borderLeftStyle string) {
	js.Rewrite("$<.borderLeftStyle = $1", borderLeftStyle)
}

func (*CSSStyleDeclaration) GetBorderLeftWidth() (borderLeftWidth string) {
	js.Rewrite("$<.borderLeftWidth")
	return borderLeftWidth
}

func (*CSSStyleDeclaration) SetBorderLeftWidth(borderLeftWidth string) {
	js.Rewrite("$<.borderLeftWidth = $1", borderLeftWidth)
}

func (*CSSStyleDeclaration) GetBorderRadius() (borderRadius string) {
	js.Rewrite("$<.borderRadius")
	return borderRadius
}

func (*CSSStyleDeclaration) SetBorderRadius(borderRadius string) {
	js.Rewrite("$<.borderRadius = $1", borderRadius)
}

func (*CSSStyleDeclaration) GetBorderRight() (borderRight string) {
	js.Rewrite("$<.borderRight")
	return borderRight
}

func (*CSSStyleDeclaration) SetBorderRight(borderRight string) {
	js.Rewrite("$<.borderRight = $1", borderRight)
}

func (*CSSStyleDeclaration) GetBorderRightColor() (borderRightColor string) {
	js.Rewrite("$<.borderRightColor")
	return borderRightColor
}

func (*CSSStyleDeclaration) SetBorderRightColor(borderRightColor string) {
	js.Rewrite("$<.borderRightColor = $1", borderRightColor)
}

func (*CSSStyleDeclaration) GetBorderRightStyle() (borderRightStyle string) {
	js.Rewrite("$<.borderRightStyle")
	return borderRightStyle
}

func (*CSSStyleDeclaration) SetBorderRightStyle(borderRightStyle string) {
	js.Rewrite("$<.borderRightStyle = $1", borderRightStyle)
}

func (*CSSStyleDeclaration) GetBorderRightWidth() (borderRightWidth string) {
	js.Rewrite("$<.borderRightWidth")
	return borderRightWidth
}

func (*CSSStyleDeclaration) SetBorderRightWidth(borderRightWidth string) {
	js.Rewrite("$<.borderRightWidth = $1", borderRightWidth)
}

func (*CSSStyleDeclaration) GetBorderSpacing() (borderSpacing string) {
	js.Rewrite("$<.borderSpacing")
	return borderSpacing
}

func (*CSSStyleDeclaration) SetBorderSpacing(borderSpacing string) {
	js.Rewrite("$<.borderSpacing = $1", borderSpacing)
}

func (*CSSStyleDeclaration) GetBorderStyle() (borderStyle string) {
	js.Rewrite("$<.borderStyle")
	return borderStyle
}

func (*CSSStyleDeclaration) SetBorderStyle(borderStyle string) {
	js.Rewrite("$<.borderStyle = $1", borderStyle)
}

func (*CSSStyleDeclaration) GetBorderTop() (borderTop string) {
	js.Rewrite("$<.borderTop")
	return borderTop
}

func (*CSSStyleDeclaration) SetBorderTop(borderTop string) {
	js.Rewrite("$<.borderTop = $1", borderTop)
}

func (*CSSStyleDeclaration) GetBorderTopColor() (borderTopColor string) {
	js.Rewrite("$<.borderTopColor")
	return borderTopColor
}

func (*CSSStyleDeclaration) SetBorderTopColor(borderTopColor string) {
	js.Rewrite("$<.borderTopColor = $1", borderTopColor)
}

func (*CSSStyleDeclaration) GetBorderTopLeftRadius() (borderTopLeftRadius string) {
	js.Rewrite("$<.borderTopLeftRadius")
	return borderTopLeftRadius
}

func (*CSSStyleDeclaration) SetBorderTopLeftRadius(borderTopLeftRadius string) {
	js.Rewrite("$<.borderTopLeftRadius = $1", borderTopLeftRadius)
}

func (*CSSStyleDeclaration) GetBorderTopRightRadius() (borderTopRightRadius string) {
	js.Rewrite("$<.borderTopRightRadius")
	return borderTopRightRadius
}

func (*CSSStyleDeclaration) SetBorderTopRightRadius(borderTopRightRadius string) {
	js.Rewrite("$<.borderTopRightRadius = $1", borderTopRightRadius)
}

func (*CSSStyleDeclaration) GetBorderTopStyle() (borderTopStyle string) {
	js.Rewrite("$<.borderTopStyle")
	return borderTopStyle
}

func (*CSSStyleDeclaration) SetBorderTopStyle(borderTopStyle string) {
	js.Rewrite("$<.borderTopStyle = $1", borderTopStyle)
}

func (*CSSStyleDeclaration) GetBorderTopWidth() (borderTopWidth string) {
	js.Rewrite("$<.borderTopWidth")
	return borderTopWidth
}

func (*CSSStyleDeclaration) SetBorderTopWidth(borderTopWidth string) {
	js.Rewrite("$<.borderTopWidth = $1", borderTopWidth)
}

func (*CSSStyleDeclaration) GetBorderWidth() (borderWidth string) {
	js.Rewrite("$<.borderWidth")
	return borderWidth
}

func (*CSSStyleDeclaration) SetBorderWidth(borderWidth string) {
	js.Rewrite("$<.borderWidth = $1", borderWidth)
}

func (*CSSStyleDeclaration) GetBottom() (bottom string) {
	js.Rewrite("$<.bottom")
	return bottom
}

func (*CSSStyleDeclaration) SetBottom(bottom string) {
	js.Rewrite("$<.bottom = $1", bottom)
}

func (*CSSStyleDeclaration) GetBoxShadow() (boxShadow string) {
	js.Rewrite("$<.boxShadow")
	return boxShadow
}

func (*CSSStyleDeclaration) SetBoxShadow(boxShadow string) {
	js.Rewrite("$<.boxShadow = $1", boxShadow)
}

func (*CSSStyleDeclaration) GetBoxSizing() (boxSizing string) {
	js.Rewrite("$<.boxSizing")
	return boxSizing
}

func (*CSSStyleDeclaration) SetBoxSizing(boxSizing string) {
	js.Rewrite("$<.boxSizing = $1", boxSizing)
}

func (*CSSStyleDeclaration) GetBreakAfter() (breakAfter string) {
	js.Rewrite("$<.breakAfter")
	return breakAfter
}

func (*CSSStyleDeclaration) SetBreakAfter(breakAfter string) {
	js.Rewrite("$<.breakAfter = $1", breakAfter)
}

func (*CSSStyleDeclaration) GetBreakBefore() (breakBefore string) {
	js.Rewrite("$<.breakBefore")
	return breakBefore
}

func (*CSSStyleDeclaration) SetBreakBefore(breakBefore string) {
	js.Rewrite("$<.breakBefore = $1", breakBefore)
}

func (*CSSStyleDeclaration) GetBreakInside() (breakInside string) {
	js.Rewrite("$<.breakInside")
	return breakInside
}

func (*CSSStyleDeclaration) SetBreakInside(breakInside string) {
	js.Rewrite("$<.breakInside = $1", breakInside)
}

func (*CSSStyleDeclaration) GetCaptionSide() (captionSide string) {
	js.Rewrite("$<.captionSide")
	return captionSide
}

func (*CSSStyleDeclaration) SetCaptionSide(captionSide string) {
	js.Rewrite("$<.captionSide = $1", captionSide)
}

func (*CSSStyleDeclaration) GetClear() (clear string) {
	js.Rewrite("$<.clear")
	return clear
}

func (*CSSStyleDeclaration) SetClear(clear string) {
	js.Rewrite("$<.clear = $1", clear)
}

func (*CSSStyleDeclaration) GetClip() (clip string) {
	js.Rewrite("$<.clip")
	return clip
}

func (*CSSStyleDeclaration) SetClip(clip string) {
	js.Rewrite("$<.clip = $1", clip)
}

func (*CSSStyleDeclaration) GetClipPath() (clipPath string) {
	js.Rewrite("$<.clipPath")
	return clipPath
}

func (*CSSStyleDeclaration) SetClipPath(clipPath string) {
	js.Rewrite("$<.clipPath = $1", clipPath)
}

func (*CSSStyleDeclaration) GetClipRule() (clipRule string) {
	js.Rewrite("$<.clipRule")
	return clipRule
}

func (*CSSStyleDeclaration) SetClipRule(clipRule string) {
	js.Rewrite("$<.clipRule = $1", clipRule)
}

func (*CSSStyleDeclaration) GetColor() (color string) {
	js.Rewrite("$<.color")
	return color
}

func (*CSSStyleDeclaration) SetColor(color string) {
	js.Rewrite("$<.color = $1", color)
}

func (*CSSStyleDeclaration) GetColorInterpolationFilters() (colorInterpolationFilters string) {
	js.Rewrite("$<.colorInterpolationFilters")
	return colorInterpolationFilters
}

func (*CSSStyleDeclaration) SetColorInterpolationFilters(colorInterpolationFilters string) {
	js.Rewrite("$<.colorInterpolationFilters = $1", colorInterpolationFilters)
}

func (*CSSStyleDeclaration) GetColumnCount() (columnCount interface{}) {
	js.Rewrite("$<.columnCount")
	return columnCount
}

func (*CSSStyleDeclaration) SetColumnCount(columnCount interface{}) {
	js.Rewrite("$<.columnCount = $1", columnCount)
}

func (*CSSStyleDeclaration) GetColumnFill() (columnFill string) {
	js.Rewrite("$<.columnFill")
	return columnFill
}

func (*CSSStyleDeclaration) SetColumnFill(columnFill string) {
	js.Rewrite("$<.columnFill = $1", columnFill)
}

func (*CSSStyleDeclaration) GetColumnGap() (columnGap interface{}) {
	js.Rewrite("$<.columnGap")
	return columnGap
}

func (*CSSStyleDeclaration) SetColumnGap(columnGap interface{}) {
	js.Rewrite("$<.columnGap = $1", columnGap)
}

func (*CSSStyleDeclaration) GetColumnRule() (columnRule string) {
	js.Rewrite("$<.columnRule")
	return columnRule
}

func (*CSSStyleDeclaration) SetColumnRule(columnRule string) {
	js.Rewrite("$<.columnRule = $1", columnRule)
}

func (*CSSStyleDeclaration) GetColumnRuleColor() (columnRuleColor interface{}) {
	js.Rewrite("$<.columnRuleColor")
	return columnRuleColor
}

func (*CSSStyleDeclaration) SetColumnRuleColor(columnRuleColor interface{}) {
	js.Rewrite("$<.columnRuleColor = $1", columnRuleColor)
}

func (*CSSStyleDeclaration) GetColumnRuleStyle() (columnRuleStyle string) {
	js.Rewrite("$<.columnRuleStyle")
	return columnRuleStyle
}

func (*CSSStyleDeclaration) SetColumnRuleStyle(columnRuleStyle string) {
	js.Rewrite("$<.columnRuleStyle = $1", columnRuleStyle)
}

func (*CSSStyleDeclaration) GetColumnRuleWidth() (columnRuleWidth interface{}) {
	js.Rewrite("$<.columnRuleWidth")
	return columnRuleWidth
}

func (*CSSStyleDeclaration) SetColumnRuleWidth(columnRuleWidth interface{}) {
	js.Rewrite("$<.columnRuleWidth = $1", columnRuleWidth)
}

func (*CSSStyleDeclaration) GetColumns() (columns string) {
	js.Rewrite("$<.columns")
	return columns
}

func (*CSSStyleDeclaration) SetColumns(columns string) {
	js.Rewrite("$<.columns = $1", columns)
}

func (*CSSStyleDeclaration) GetColumnSpan() (columnSpan string) {
	js.Rewrite("$<.columnSpan")
	return columnSpan
}

func (*CSSStyleDeclaration) SetColumnSpan(columnSpan string) {
	js.Rewrite("$<.columnSpan = $1", columnSpan)
}

func (*CSSStyleDeclaration) GetColumnWidth() (columnWidth interface{}) {
	js.Rewrite("$<.columnWidth")
	return columnWidth
}

func (*CSSStyleDeclaration) SetColumnWidth(columnWidth interface{}) {
	js.Rewrite("$<.columnWidth = $1", columnWidth)
}

func (*CSSStyleDeclaration) GetContent() (content string) {
	js.Rewrite("$<.content")
	return content
}

func (*CSSStyleDeclaration) SetContent(content string) {
	js.Rewrite("$<.content = $1", content)
}

func (*CSSStyleDeclaration) GetCounterIncrement() (counterIncrement string) {
	js.Rewrite("$<.counterIncrement")
	return counterIncrement
}

func (*CSSStyleDeclaration) SetCounterIncrement(counterIncrement string) {
	js.Rewrite("$<.counterIncrement = $1", counterIncrement)
}

func (*CSSStyleDeclaration) GetCounterReset() (counterReset string) {
	js.Rewrite("$<.counterReset")
	return counterReset
}

func (*CSSStyleDeclaration) SetCounterReset(counterReset string) {
	js.Rewrite("$<.counterReset = $1", counterReset)
}

func (*CSSStyleDeclaration) GetCSSFloat() (cssFloat string) {
	js.Rewrite("$<.cssFloat")
	return cssFloat
}

func (*CSSStyleDeclaration) SetCSSFloat(cssFloat string) {
	js.Rewrite("$<.cssFloat = $1", cssFloat)
}

func (*CSSStyleDeclaration) GetCSSText() (cssText string) {
	js.Rewrite("$<.cssText")
	return cssText
}

func (*CSSStyleDeclaration) SetCSSText(cssText string) {
	js.Rewrite("$<.cssText = $1", cssText)
}

func (*CSSStyleDeclaration) GetCursor() (cursor string) {
	js.Rewrite("$<.cursor")
	return cursor
}

func (*CSSStyleDeclaration) SetCursor(cursor string) {
	js.Rewrite("$<.cursor = $1", cursor)
}

func (*CSSStyleDeclaration) GetDirection() (direction string) {
	js.Rewrite("$<.direction")
	return direction
}

func (*CSSStyleDeclaration) SetDirection(direction string) {
	js.Rewrite("$<.direction = $1", direction)
}

func (*CSSStyleDeclaration) GetDisplay() (display string) {
	js.Rewrite("$<.display")
	return display
}

func (*CSSStyleDeclaration) SetDisplay(display string) {
	js.Rewrite("$<.display = $1", display)
}

func (*CSSStyleDeclaration) GetDominantBaseline() (dominantBaseline string) {
	js.Rewrite("$<.dominantBaseline")
	return dominantBaseline
}

func (*CSSStyleDeclaration) SetDominantBaseline(dominantBaseline string) {
	js.Rewrite("$<.dominantBaseline = $1", dominantBaseline)
}

func (*CSSStyleDeclaration) GetEmptyCells() (emptyCells string) {
	js.Rewrite("$<.emptyCells")
	return emptyCells
}

func (*CSSStyleDeclaration) SetEmptyCells(emptyCells string) {
	js.Rewrite("$<.emptyCells = $1", emptyCells)
}

func (*CSSStyleDeclaration) GetEnableBackground() (enableBackground string) {
	js.Rewrite("$<.enableBackground")
	return enableBackground
}

func (*CSSStyleDeclaration) SetEnableBackground(enableBackground string) {
	js.Rewrite("$<.enableBackground = $1", enableBackground)
}

func (*CSSStyleDeclaration) GetFill() (fill string) {
	js.Rewrite("$<.fill")
	return fill
}

func (*CSSStyleDeclaration) SetFill(fill string) {
	js.Rewrite("$<.fill = $1", fill)
}

func (*CSSStyleDeclaration) GetFillOpacity() (fillOpacity string) {
	js.Rewrite("$<.fillOpacity")
	return fillOpacity
}

func (*CSSStyleDeclaration) SetFillOpacity(fillOpacity string) {
	js.Rewrite("$<.fillOpacity = $1", fillOpacity)
}

func (*CSSStyleDeclaration) GetFillRule() (fillRule string) {
	js.Rewrite("$<.fillRule")
	return fillRule
}

func (*CSSStyleDeclaration) SetFillRule(fillRule string) {
	js.Rewrite("$<.fillRule = $1", fillRule)
}

func (*CSSStyleDeclaration) GetFilter() (filter string) {
	js.Rewrite("$<.filter")
	return filter
}

func (*CSSStyleDeclaration) SetFilter(filter string) {
	js.Rewrite("$<.filter = $1", filter)
}

func (*CSSStyleDeclaration) GetFlex() (flex string) {
	js.Rewrite("$<.flex")
	return flex
}

func (*CSSStyleDeclaration) SetFlex(flex string) {
	js.Rewrite("$<.flex = $1", flex)
}

func (*CSSStyleDeclaration) GetFlexBasis() (flexBasis string) {
	js.Rewrite("$<.flexBasis")
	return flexBasis
}

func (*CSSStyleDeclaration) SetFlexBasis(flexBasis string) {
	js.Rewrite("$<.flexBasis = $1", flexBasis)
}

func (*CSSStyleDeclaration) GetFlexDirection() (flexDirection string) {
	js.Rewrite("$<.flexDirection")
	return flexDirection
}

func (*CSSStyleDeclaration) SetFlexDirection(flexDirection string) {
	js.Rewrite("$<.flexDirection = $1", flexDirection)
}

func (*CSSStyleDeclaration) GetFlexFlow() (flexFlow string) {
	js.Rewrite("$<.flexFlow")
	return flexFlow
}

func (*CSSStyleDeclaration) SetFlexFlow(flexFlow string) {
	js.Rewrite("$<.flexFlow = $1", flexFlow)
}

func (*CSSStyleDeclaration) GetFlexGrow() (flexGrow string) {
	js.Rewrite("$<.flexGrow")
	return flexGrow
}

func (*CSSStyleDeclaration) SetFlexGrow(flexGrow string) {
	js.Rewrite("$<.flexGrow = $1", flexGrow)
}

func (*CSSStyleDeclaration) GetFlexShrink() (flexShrink string) {
	js.Rewrite("$<.flexShrink")
	return flexShrink
}

func (*CSSStyleDeclaration) SetFlexShrink(flexShrink string) {
	js.Rewrite("$<.flexShrink = $1", flexShrink)
}

func (*CSSStyleDeclaration) GetFlexWrap() (flexWrap string) {
	js.Rewrite("$<.flexWrap")
	return flexWrap
}

func (*CSSStyleDeclaration) SetFlexWrap(flexWrap string) {
	js.Rewrite("$<.flexWrap = $1", flexWrap)
}

func (*CSSStyleDeclaration) GetFloodColor() (floodColor string) {
	js.Rewrite("$<.floodColor")
	return floodColor
}

func (*CSSStyleDeclaration) SetFloodColor(floodColor string) {
	js.Rewrite("$<.floodColor = $1", floodColor)
}

func (*CSSStyleDeclaration) GetFloodOpacity() (floodOpacity string) {
	js.Rewrite("$<.floodOpacity")
	return floodOpacity
}

func (*CSSStyleDeclaration) SetFloodOpacity(floodOpacity string) {
	js.Rewrite("$<.floodOpacity = $1", floodOpacity)
}

func (*CSSStyleDeclaration) GetFont() (font string) {
	js.Rewrite("$<.font")
	return font
}

func (*CSSStyleDeclaration) SetFont(font string) {
	js.Rewrite("$<.font = $1", font)
}

func (*CSSStyleDeclaration) GetFontFamily() (fontFamily string) {
	js.Rewrite("$<.fontFamily")
	return fontFamily
}

func (*CSSStyleDeclaration) SetFontFamily(fontFamily string) {
	js.Rewrite("$<.fontFamily = $1", fontFamily)
}

func (*CSSStyleDeclaration) GetFontFeatureSettings() (fontFeatureSettings string) {
	js.Rewrite("$<.fontFeatureSettings")
	return fontFeatureSettings
}

func (*CSSStyleDeclaration) SetFontFeatureSettings(fontFeatureSettings string) {
	js.Rewrite("$<.fontFeatureSettings = $1", fontFeatureSettings)
}

func (*CSSStyleDeclaration) GetFontSize() (fontSize string) {
	js.Rewrite("$<.fontSize")
	return fontSize
}

func (*CSSStyleDeclaration) SetFontSize(fontSize string) {
	js.Rewrite("$<.fontSize = $1", fontSize)
}

func (*CSSStyleDeclaration) GetFontSizeAdjust() (fontSizeAdjust string) {
	js.Rewrite("$<.fontSizeAdjust")
	return fontSizeAdjust
}

func (*CSSStyleDeclaration) SetFontSizeAdjust(fontSizeAdjust string) {
	js.Rewrite("$<.fontSizeAdjust = $1", fontSizeAdjust)
}

func (*CSSStyleDeclaration) GetFontStretch() (fontStretch string) {
	js.Rewrite("$<.fontStretch")
	return fontStretch
}

func (*CSSStyleDeclaration) SetFontStretch(fontStretch string) {
	js.Rewrite("$<.fontStretch = $1", fontStretch)
}

func (*CSSStyleDeclaration) GetFontStyle() (fontStyle string) {
	js.Rewrite("$<.fontStyle")
	return fontStyle
}

func (*CSSStyleDeclaration) SetFontStyle(fontStyle string) {
	js.Rewrite("$<.fontStyle = $1", fontStyle)
}

func (*CSSStyleDeclaration) GetFontVariant() (fontVariant string) {
	js.Rewrite("$<.fontVariant")
	return fontVariant
}

func (*CSSStyleDeclaration) SetFontVariant(fontVariant string) {
	js.Rewrite("$<.fontVariant = $1", fontVariant)
}

func (*CSSStyleDeclaration) GetFontWeight() (fontWeight string) {
	js.Rewrite("$<.fontWeight")
	return fontWeight
}

func (*CSSStyleDeclaration) SetFontWeight(fontWeight string) {
	js.Rewrite("$<.fontWeight = $1", fontWeight)
}

func (*CSSStyleDeclaration) GetGlyphOrientationHorizontal() (glyphOrientationHorizontal string) {
	js.Rewrite("$<.glyphOrientationHorizontal")
	return glyphOrientationHorizontal
}

func (*CSSStyleDeclaration) SetGlyphOrientationHorizontal(glyphOrientationHorizontal string) {
	js.Rewrite("$<.glyphOrientationHorizontal = $1", glyphOrientationHorizontal)
}

func (*CSSStyleDeclaration) GetGlyphOrientationVertical() (glyphOrientationVertical string) {
	js.Rewrite("$<.glyphOrientationVertical")
	return glyphOrientationVertical
}

func (*CSSStyleDeclaration) SetGlyphOrientationVertical(glyphOrientationVertical string) {
	js.Rewrite("$<.glyphOrientationVertical = $1", glyphOrientationVertical)
}

func (*CSSStyleDeclaration) GetHeight() (height string) {
	js.Rewrite("$<.height")
	return height
}

func (*CSSStyleDeclaration) SetHeight(height string) {
	js.Rewrite("$<.height = $1", height)
}

func (*CSSStyleDeclaration) GetImeMode() (imeMode string) {
	js.Rewrite("$<.imeMode")
	return imeMode
}

func (*CSSStyleDeclaration) SetImeMode(imeMode string) {
	js.Rewrite("$<.imeMode = $1", imeMode)
}

func (*CSSStyleDeclaration) GetJustifyContent() (justifyContent string) {
	js.Rewrite("$<.justifyContent")
	return justifyContent
}

func (*CSSStyleDeclaration) SetJustifyContent(justifyContent string) {
	js.Rewrite("$<.justifyContent = $1", justifyContent)
}

func (*CSSStyleDeclaration) GetKerning() (kerning string) {
	js.Rewrite("$<.kerning")
	return kerning
}

func (*CSSStyleDeclaration) SetKerning(kerning string) {
	js.Rewrite("$<.kerning = $1", kerning)
}

func (*CSSStyleDeclaration) GetLayoutGrid() (layoutGrid string) {
	js.Rewrite("$<.layoutGrid")
	return layoutGrid
}

func (*CSSStyleDeclaration) SetLayoutGrid(layoutGrid string) {
	js.Rewrite("$<.layoutGrid = $1", layoutGrid)
}

func (*CSSStyleDeclaration) GetLayoutGridChar() (layoutGridChar string) {
	js.Rewrite("$<.layoutGridChar")
	return layoutGridChar
}

func (*CSSStyleDeclaration) SetLayoutGridChar(layoutGridChar string) {
	js.Rewrite("$<.layoutGridChar = $1", layoutGridChar)
}

func (*CSSStyleDeclaration) GetLayoutGridLine() (layoutGridLine string) {
	js.Rewrite("$<.layoutGridLine")
	return layoutGridLine
}

func (*CSSStyleDeclaration) SetLayoutGridLine(layoutGridLine string) {
	js.Rewrite("$<.layoutGridLine = $1", layoutGridLine)
}

func (*CSSStyleDeclaration) GetLayoutGridMode() (layoutGridMode string) {
	js.Rewrite("$<.layoutGridMode")
	return layoutGridMode
}

func (*CSSStyleDeclaration) SetLayoutGridMode(layoutGridMode string) {
	js.Rewrite("$<.layoutGridMode = $1", layoutGridMode)
}

func (*CSSStyleDeclaration) GetLayoutGridType() (layoutGridType string) {
	js.Rewrite("$<.layoutGridType")
	return layoutGridType
}

func (*CSSStyleDeclaration) SetLayoutGridType(layoutGridType string) {
	js.Rewrite("$<.layoutGridType = $1", layoutGridType)
}

func (*CSSStyleDeclaration) GetLeft() (left string) {
	js.Rewrite("$<.left")
	return left
}

func (*CSSStyleDeclaration) SetLeft(left string) {
	js.Rewrite("$<.left = $1", left)
}

func (*CSSStyleDeclaration) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

func (*CSSStyleDeclaration) GetLetterSpacing() (letterSpacing string) {
	js.Rewrite("$<.letterSpacing")
	return letterSpacing
}

func (*CSSStyleDeclaration) SetLetterSpacing(letterSpacing string) {
	js.Rewrite("$<.letterSpacing = $1", letterSpacing)
}

func (*CSSStyleDeclaration) GetLightingColor() (lightingColor string) {
	js.Rewrite("$<.lightingColor")
	return lightingColor
}

func (*CSSStyleDeclaration) SetLightingColor(lightingColor string) {
	js.Rewrite("$<.lightingColor = $1", lightingColor)
}

func (*CSSStyleDeclaration) GetLineBreak() (lineBreak string) {
	js.Rewrite("$<.lineBreak")
	return lineBreak
}

func (*CSSStyleDeclaration) SetLineBreak(lineBreak string) {
	js.Rewrite("$<.lineBreak = $1", lineBreak)
}

func (*CSSStyleDeclaration) GetLineHeight() (lineHeight string) {
	js.Rewrite("$<.lineHeight")
	return lineHeight
}

func (*CSSStyleDeclaration) SetLineHeight(lineHeight string) {
	js.Rewrite("$<.lineHeight = $1", lineHeight)
}

func (*CSSStyleDeclaration) GetListStyle() (listStyle string) {
	js.Rewrite("$<.listStyle")
	return listStyle
}

func (*CSSStyleDeclaration) SetListStyle(listStyle string) {
	js.Rewrite("$<.listStyle = $1", listStyle)
}

func (*CSSStyleDeclaration) GetListStyleImage() (listStyleImage string) {
	js.Rewrite("$<.listStyleImage")
	return listStyleImage
}

func (*CSSStyleDeclaration) SetListStyleImage(listStyleImage string) {
	js.Rewrite("$<.listStyleImage = $1", listStyleImage)
}

func (*CSSStyleDeclaration) GetListStylePosition() (listStylePosition string) {
	js.Rewrite("$<.listStylePosition")
	return listStylePosition
}

func (*CSSStyleDeclaration) SetListStylePosition(listStylePosition string) {
	js.Rewrite("$<.listStylePosition = $1", listStylePosition)
}

func (*CSSStyleDeclaration) GetListStyleType() (listStyleType string) {
	js.Rewrite("$<.listStyleType")
	return listStyleType
}

func (*CSSStyleDeclaration) SetListStyleType(listStyleType string) {
	js.Rewrite("$<.listStyleType = $1", listStyleType)
}

func (*CSSStyleDeclaration) GetMargin() (margin string) {
	js.Rewrite("$<.margin")
	return margin
}

func (*CSSStyleDeclaration) SetMargin(margin string) {
	js.Rewrite("$<.margin = $1", margin)
}

func (*CSSStyleDeclaration) GetMarginBottom() (marginBottom string) {
	js.Rewrite("$<.marginBottom")
	return marginBottom
}

func (*CSSStyleDeclaration) SetMarginBottom(marginBottom string) {
	js.Rewrite("$<.marginBottom = $1", marginBottom)
}

func (*CSSStyleDeclaration) GetMarginLeft() (marginLeft string) {
	js.Rewrite("$<.marginLeft")
	return marginLeft
}

func (*CSSStyleDeclaration) SetMarginLeft(marginLeft string) {
	js.Rewrite("$<.marginLeft = $1", marginLeft)
}

func (*CSSStyleDeclaration) GetMarginRight() (marginRight string) {
	js.Rewrite("$<.marginRight")
	return marginRight
}

func (*CSSStyleDeclaration) SetMarginRight(marginRight string) {
	js.Rewrite("$<.marginRight = $1", marginRight)
}

func (*CSSStyleDeclaration) GetMarginTop() (marginTop string) {
	js.Rewrite("$<.marginTop")
	return marginTop
}

func (*CSSStyleDeclaration) SetMarginTop(marginTop string) {
	js.Rewrite("$<.marginTop = $1", marginTop)
}

func (*CSSStyleDeclaration) GetMarker() (marker string) {
	js.Rewrite("$<.marker")
	return marker
}

func (*CSSStyleDeclaration) SetMarker(marker string) {
	js.Rewrite("$<.marker = $1", marker)
}

func (*CSSStyleDeclaration) GetMarkerEnd() (markerEnd string) {
	js.Rewrite("$<.markerEnd")
	return markerEnd
}

func (*CSSStyleDeclaration) SetMarkerEnd(markerEnd string) {
	js.Rewrite("$<.markerEnd = $1", markerEnd)
}

func (*CSSStyleDeclaration) GetMarkerMid() (markerMid string) {
	js.Rewrite("$<.markerMid")
	return markerMid
}

func (*CSSStyleDeclaration) SetMarkerMid(markerMid string) {
	js.Rewrite("$<.markerMid = $1", markerMid)
}

func (*CSSStyleDeclaration) GetMarkerStart() (markerStart string) {
	js.Rewrite("$<.markerStart")
	return markerStart
}

func (*CSSStyleDeclaration) SetMarkerStart(markerStart string) {
	js.Rewrite("$<.markerStart = $1", markerStart)
}

func (*CSSStyleDeclaration) GetMask() (mask string) {
	js.Rewrite("$<.mask")
	return mask
}

func (*CSSStyleDeclaration) SetMask(mask string) {
	js.Rewrite("$<.mask = $1", mask)
}

func (*CSSStyleDeclaration) GetMaxHeight() (maxHeight string) {
	js.Rewrite("$<.maxHeight")
	return maxHeight
}

func (*CSSStyleDeclaration) SetMaxHeight(maxHeight string) {
	js.Rewrite("$<.maxHeight = $1", maxHeight)
}

func (*CSSStyleDeclaration) GetMaxWidth() (maxWidth string) {
	js.Rewrite("$<.maxWidth")
	return maxWidth
}

func (*CSSStyleDeclaration) SetMaxWidth(maxWidth string) {
	js.Rewrite("$<.maxWidth = $1", maxWidth)
}

func (*CSSStyleDeclaration) GetMinHeight() (minHeight string) {
	js.Rewrite("$<.minHeight")
	return minHeight
}

func (*CSSStyleDeclaration) SetMinHeight(minHeight string) {
	js.Rewrite("$<.minHeight = $1", minHeight)
}

func (*CSSStyleDeclaration) GetMinWidth() (minWidth string) {
	js.Rewrite("$<.minWidth")
	return minWidth
}

func (*CSSStyleDeclaration) SetMinWidth(minWidth string) {
	js.Rewrite("$<.minWidth = $1", minWidth)
}

func (*CSSStyleDeclaration) GetMsContentZoomChaining() (msContentZoomChaining string) {
	js.Rewrite("$<.msContentZoomChaining")
	return msContentZoomChaining
}

func (*CSSStyleDeclaration) SetMsContentZoomChaining(msContentZoomChaining string) {
	js.Rewrite("$<.msContentZoomChaining = $1", msContentZoomChaining)
}

func (*CSSStyleDeclaration) GetMsContentZooming() (msContentZooming string) {
	js.Rewrite("$<.msContentZooming")
	return msContentZooming
}

func (*CSSStyleDeclaration) SetMsContentZooming(msContentZooming string) {
	js.Rewrite("$<.msContentZooming = $1", msContentZooming)
}

func (*CSSStyleDeclaration) GetMsContentZoomLimit() (msContentZoomLimit string) {
	js.Rewrite("$<.msContentZoomLimit")
	return msContentZoomLimit
}

func (*CSSStyleDeclaration) SetMsContentZoomLimit(msContentZoomLimit string) {
	js.Rewrite("$<.msContentZoomLimit = $1", msContentZoomLimit)
}

func (*CSSStyleDeclaration) GetMsContentZoomLimitMax() (msContentZoomLimitMax interface{}) {
	js.Rewrite("$<.msContentZoomLimitMax")
	return msContentZoomLimitMax
}

func (*CSSStyleDeclaration) SetMsContentZoomLimitMax(msContentZoomLimitMax interface{}) {
	js.Rewrite("$<.msContentZoomLimitMax = $1", msContentZoomLimitMax)
}

func (*CSSStyleDeclaration) GetMsContentZoomLimitMin() (msContentZoomLimitMin interface{}) {
	js.Rewrite("$<.msContentZoomLimitMin")
	return msContentZoomLimitMin
}

func (*CSSStyleDeclaration) SetMsContentZoomLimitMin(msContentZoomLimitMin interface{}) {
	js.Rewrite("$<.msContentZoomLimitMin = $1", msContentZoomLimitMin)
}

func (*CSSStyleDeclaration) GetMsContentZoomSnap() (msContentZoomSnap string) {
	js.Rewrite("$<.msContentZoomSnap")
	return msContentZoomSnap
}

func (*CSSStyleDeclaration) SetMsContentZoomSnap(msContentZoomSnap string) {
	js.Rewrite("$<.msContentZoomSnap = $1", msContentZoomSnap)
}

func (*CSSStyleDeclaration) GetMsContentZoomSnapPoints() (msContentZoomSnapPoints string) {
	js.Rewrite("$<.msContentZoomSnapPoints")
	return msContentZoomSnapPoints
}

func (*CSSStyleDeclaration) SetMsContentZoomSnapPoints(msContentZoomSnapPoints string) {
	js.Rewrite("$<.msContentZoomSnapPoints = $1", msContentZoomSnapPoints)
}

func (*CSSStyleDeclaration) GetMsContentZoomSnapType() (msContentZoomSnapType string) {
	js.Rewrite("$<.msContentZoomSnapType")
	return msContentZoomSnapType
}

func (*CSSStyleDeclaration) SetMsContentZoomSnapType(msContentZoomSnapType string) {
	js.Rewrite("$<.msContentZoomSnapType = $1", msContentZoomSnapType)
}

func (*CSSStyleDeclaration) GetMsFlowFrom() (msFlowFrom string) {
	js.Rewrite("$<.msFlowFrom")
	return msFlowFrom
}

func (*CSSStyleDeclaration) SetMsFlowFrom(msFlowFrom string) {
	js.Rewrite("$<.msFlowFrom = $1", msFlowFrom)
}

func (*CSSStyleDeclaration) GetMsFlowInto() (msFlowInto string) {
	js.Rewrite("$<.msFlowInto")
	return msFlowInto
}

func (*CSSStyleDeclaration) SetMsFlowInto(msFlowInto string) {
	js.Rewrite("$<.msFlowInto = $1", msFlowInto)
}

func (*CSSStyleDeclaration) GetMsFontFeatureSettings() (msFontFeatureSettings string) {
	js.Rewrite("$<.msFontFeatureSettings")
	return msFontFeatureSettings
}

func (*CSSStyleDeclaration) SetMsFontFeatureSettings(msFontFeatureSettings string) {
	js.Rewrite("$<.msFontFeatureSettings = $1", msFontFeatureSettings)
}

func (*CSSStyleDeclaration) GetMsGridColumn() (msGridColumn interface{}) {
	js.Rewrite("$<.msGridColumn")
	return msGridColumn
}

func (*CSSStyleDeclaration) SetMsGridColumn(msGridColumn interface{}) {
	js.Rewrite("$<.msGridColumn = $1", msGridColumn)
}

func (*CSSStyleDeclaration) GetMsGridColumnAlign() (msGridColumnAlign string) {
	js.Rewrite("$<.msGridColumnAlign")
	return msGridColumnAlign
}

func (*CSSStyleDeclaration) SetMsGridColumnAlign(msGridColumnAlign string) {
	js.Rewrite("$<.msGridColumnAlign = $1", msGridColumnAlign)
}

func (*CSSStyleDeclaration) GetMsGridColumns() (msGridColumns string) {
	js.Rewrite("$<.msGridColumns")
	return msGridColumns
}

func (*CSSStyleDeclaration) SetMsGridColumns(msGridColumns string) {
	js.Rewrite("$<.msGridColumns = $1", msGridColumns)
}

func (*CSSStyleDeclaration) GetMsGridColumnSpan() (msGridColumnSpan interface{}) {
	js.Rewrite("$<.msGridColumnSpan")
	return msGridColumnSpan
}

func (*CSSStyleDeclaration) SetMsGridColumnSpan(msGridColumnSpan interface{}) {
	js.Rewrite("$<.msGridColumnSpan = $1", msGridColumnSpan)
}

func (*CSSStyleDeclaration) GetMsGridRow() (msGridRow interface{}) {
	js.Rewrite("$<.msGridRow")
	return msGridRow
}

func (*CSSStyleDeclaration) SetMsGridRow(msGridRow interface{}) {
	js.Rewrite("$<.msGridRow = $1", msGridRow)
}

func (*CSSStyleDeclaration) GetMsGridRowAlign() (msGridRowAlign string) {
	js.Rewrite("$<.msGridRowAlign")
	return msGridRowAlign
}

func (*CSSStyleDeclaration) SetMsGridRowAlign(msGridRowAlign string) {
	js.Rewrite("$<.msGridRowAlign = $1", msGridRowAlign)
}

func (*CSSStyleDeclaration) GetMsGridRows() (msGridRows string) {
	js.Rewrite("$<.msGridRows")
	return msGridRows
}

func (*CSSStyleDeclaration) SetMsGridRows(msGridRows string) {
	js.Rewrite("$<.msGridRows = $1", msGridRows)
}

func (*CSSStyleDeclaration) GetMsGridRowSpan() (msGridRowSpan interface{}) {
	js.Rewrite("$<.msGridRowSpan")
	return msGridRowSpan
}

func (*CSSStyleDeclaration) SetMsGridRowSpan(msGridRowSpan interface{}) {
	js.Rewrite("$<.msGridRowSpan = $1", msGridRowSpan)
}

func (*CSSStyleDeclaration) GetMsHighContrastAdjust() (msHighContrastAdjust string) {
	js.Rewrite("$<.msHighContrastAdjust")
	return msHighContrastAdjust
}

func (*CSSStyleDeclaration) SetMsHighContrastAdjust(msHighContrastAdjust string) {
	js.Rewrite("$<.msHighContrastAdjust = $1", msHighContrastAdjust)
}

func (*CSSStyleDeclaration) GetMsHyphenateLimitChars() (msHyphenateLimitChars string) {
	js.Rewrite("$<.msHyphenateLimitChars")
	return msHyphenateLimitChars
}

func (*CSSStyleDeclaration) SetMsHyphenateLimitChars(msHyphenateLimitChars string) {
	js.Rewrite("$<.msHyphenateLimitChars = $1", msHyphenateLimitChars)
}

func (*CSSStyleDeclaration) GetMsHyphenateLimitLines() (msHyphenateLimitLines interface{}) {
	js.Rewrite("$<.msHyphenateLimitLines")
	return msHyphenateLimitLines
}

func (*CSSStyleDeclaration) SetMsHyphenateLimitLines(msHyphenateLimitLines interface{}) {
	js.Rewrite("$<.msHyphenateLimitLines = $1", msHyphenateLimitLines)
}

func (*CSSStyleDeclaration) GetMsHyphenateLimitZone() (msHyphenateLimitZone interface{}) {
	js.Rewrite("$<.msHyphenateLimitZone")
	return msHyphenateLimitZone
}

func (*CSSStyleDeclaration) SetMsHyphenateLimitZone(msHyphenateLimitZone interface{}) {
	js.Rewrite("$<.msHyphenateLimitZone = $1", msHyphenateLimitZone)
}

func (*CSSStyleDeclaration) GetMsHyphens() (msHyphens string) {
	js.Rewrite("$<.msHyphens")
	return msHyphens
}

func (*CSSStyleDeclaration) SetMsHyphens(msHyphens string) {
	js.Rewrite("$<.msHyphens = $1", msHyphens)
}

func (*CSSStyleDeclaration) GetMsImeAlign() (msImeAlign string) {
	js.Rewrite("$<.msImeAlign")
	return msImeAlign
}

func (*CSSStyleDeclaration) SetMsImeAlign(msImeAlign string) {
	js.Rewrite("$<.msImeAlign = $1", msImeAlign)
}

func (*CSSStyleDeclaration) GetMsOverflowStyle() (msOverflowStyle string) {
	js.Rewrite("$<.msOverflowStyle")
	return msOverflowStyle
}

func (*CSSStyleDeclaration) SetMsOverflowStyle(msOverflowStyle string) {
	js.Rewrite("$<.msOverflowStyle = $1", msOverflowStyle)
}

func (*CSSStyleDeclaration) GetMsScrollChaining() (msScrollChaining string) {
	js.Rewrite("$<.msScrollChaining")
	return msScrollChaining
}

func (*CSSStyleDeclaration) SetMsScrollChaining(msScrollChaining string) {
	js.Rewrite("$<.msScrollChaining = $1", msScrollChaining)
}

func (*CSSStyleDeclaration) GetMsScrollLimit() (msScrollLimit string) {
	js.Rewrite("$<.msScrollLimit")
	return msScrollLimit
}

func (*CSSStyleDeclaration) SetMsScrollLimit(msScrollLimit string) {
	js.Rewrite("$<.msScrollLimit = $1", msScrollLimit)
}

func (*CSSStyleDeclaration) GetMsScrollLimitXMax() (msScrollLimitXMax interface{}) {
	js.Rewrite("$<.msScrollLimitXMax")
	return msScrollLimitXMax
}

func (*CSSStyleDeclaration) SetMsScrollLimitXMax(msScrollLimitXMax interface{}) {
	js.Rewrite("$<.msScrollLimitXMax = $1", msScrollLimitXMax)
}

func (*CSSStyleDeclaration) GetMsScrollLimitXMin() (msScrollLimitXMin interface{}) {
	js.Rewrite("$<.msScrollLimitXMin")
	return msScrollLimitXMin
}

func (*CSSStyleDeclaration) SetMsScrollLimitXMin(msScrollLimitXMin interface{}) {
	js.Rewrite("$<.msScrollLimitXMin = $1", msScrollLimitXMin)
}

func (*CSSStyleDeclaration) GetMsScrollLimitYMax() (msScrollLimitYMax interface{}) {
	js.Rewrite("$<.msScrollLimitYMax")
	return msScrollLimitYMax
}

func (*CSSStyleDeclaration) SetMsScrollLimitYMax(msScrollLimitYMax interface{}) {
	js.Rewrite("$<.msScrollLimitYMax = $1", msScrollLimitYMax)
}

func (*CSSStyleDeclaration) GetMsScrollLimitYMin() (msScrollLimitYMin interface{}) {
	js.Rewrite("$<.msScrollLimitYMin")
	return msScrollLimitYMin
}

func (*CSSStyleDeclaration) SetMsScrollLimitYMin(msScrollLimitYMin interface{}) {
	js.Rewrite("$<.msScrollLimitYMin = $1", msScrollLimitYMin)
}

func (*CSSStyleDeclaration) GetMsScrollRails() (msScrollRails string) {
	js.Rewrite("$<.msScrollRails")
	return msScrollRails
}

func (*CSSStyleDeclaration) SetMsScrollRails(msScrollRails string) {
	js.Rewrite("$<.msScrollRails = $1", msScrollRails)
}

func (*CSSStyleDeclaration) GetMsScrollSnapPointsX() (msScrollSnapPointsX string) {
	js.Rewrite("$<.msScrollSnapPointsX")
	return msScrollSnapPointsX
}

func (*CSSStyleDeclaration) SetMsScrollSnapPointsX(msScrollSnapPointsX string) {
	js.Rewrite("$<.msScrollSnapPointsX = $1", msScrollSnapPointsX)
}

func (*CSSStyleDeclaration) GetMsScrollSnapPointsY() (msScrollSnapPointsY string) {
	js.Rewrite("$<.msScrollSnapPointsY")
	return msScrollSnapPointsY
}

func (*CSSStyleDeclaration) SetMsScrollSnapPointsY(msScrollSnapPointsY string) {
	js.Rewrite("$<.msScrollSnapPointsY = $1", msScrollSnapPointsY)
}

func (*CSSStyleDeclaration) GetMsScrollSnapType() (msScrollSnapType string) {
	js.Rewrite("$<.msScrollSnapType")
	return msScrollSnapType
}

func (*CSSStyleDeclaration) SetMsScrollSnapType(msScrollSnapType string) {
	js.Rewrite("$<.msScrollSnapType = $1", msScrollSnapType)
}

func (*CSSStyleDeclaration) GetMsScrollSnapX() (msScrollSnapX string) {
	js.Rewrite("$<.msScrollSnapX")
	return msScrollSnapX
}

func (*CSSStyleDeclaration) SetMsScrollSnapX(msScrollSnapX string) {
	js.Rewrite("$<.msScrollSnapX = $1", msScrollSnapX)
}

func (*CSSStyleDeclaration) GetMsScrollSnapY() (msScrollSnapY string) {
	js.Rewrite("$<.msScrollSnapY")
	return msScrollSnapY
}

func (*CSSStyleDeclaration) SetMsScrollSnapY(msScrollSnapY string) {
	js.Rewrite("$<.msScrollSnapY = $1", msScrollSnapY)
}

func (*CSSStyleDeclaration) GetMsScrollTranslation() (msScrollTranslation string) {
	js.Rewrite("$<.msScrollTranslation")
	return msScrollTranslation
}

func (*CSSStyleDeclaration) SetMsScrollTranslation(msScrollTranslation string) {
	js.Rewrite("$<.msScrollTranslation = $1", msScrollTranslation)
}

func (*CSSStyleDeclaration) GetMsTextCombineHorizontal() (msTextCombineHorizontal string) {
	js.Rewrite("$<.msTextCombineHorizontal")
	return msTextCombineHorizontal
}

func (*CSSStyleDeclaration) SetMsTextCombineHorizontal(msTextCombineHorizontal string) {
	js.Rewrite("$<.msTextCombineHorizontal = $1", msTextCombineHorizontal)
}

func (*CSSStyleDeclaration) GetMsTextSizeAdjust() (msTextSizeAdjust interface{}) {
	js.Rewrite("$<.msTextSizeAdjust")
	return msTextSizeAdjust
}

func (*CSSStyleDeclaration) SetMsTextSizeAdjust(msTextSizeAdjust interface{}) {
	js.Rewrite("$<.msTextSizeAdjust = $1", msTextSizeAdjust)
}

func (*CSSStyleDeclaration) GetMsTouchAction() (msTouchAction string) {
	js.Rewrite("$<.msTouchAction")
	return msTouchAction
}

func (*CSSStyleDeclaration) SetMsTouchAction(msTouchAction string) {
	js.Rewrite("$<.msTouchAction = $1", msTouchAction)
}

func (*CSSStyleDeclaration) GetMsTouchSelect() (msTouchSelect string) {
	js.Rewrite("$<.msTouchSelect")
	return msTouchSelect
}

func (*CSSStyleDeclaration) SetMsTouchSelect(msTouchSelect string) {
	js.Rewrite("$<.msTouchSelect = $1", msTouchSelect)
}

func (*CSSStyleDeclaration) GetMsUserSelect() (msUserSelect string) {
	js.Rewrite("$<.msUserSelect")
	return msUserSelect
}

func (*CSSStyleDeclaration) SetMsUserSelect(msUserSelect string) {
	js.Rewrite("$<.msUserSelect = $1", msUserSelect)
}

func (*CSSStyleDeclaration) GetMsWrapFlow() (msWrapFlow string) {
	js.Rewrite("$<.msWrapFlow")
	return msWrapFlow
}

func (*CSSStyleDeclaration) SetMsWrapFlow(msWrapFlow string) {
	js.Rewrite("$<.msWrapFlow = $1", msWrapFlow)
}

func (*CSSStyleDeclaration) GetMsWrapMargin() (msWrapMargin interface{}) {
	js.Rewrite("$<.msWrapMargin")
	return msWrapMargin
}

func (*CSSStyleDeclaration) SetMsWrapMargin(msWrapMargin interface{}) {
	js.Rewrite("$<.msWrapMargin = $1", msWrapMargin)
}

func (*CSSStyleDeclaration) GetMsWrapThrough() (msWrapThrough string) {
	js.Rewrite("$<.msWrapThrough")
	return msWrapThrough
}

func (*CSSStyleDeclaration) SetMsWrapThrough(msWrapThrough string) {
	js.Rewrite("$<.msWrapThrough = $1", msWrapThrough)
}

func (*CSSStyleDeclaration) GetOpacity() (opacity string) {
	js.Rewrite("$<.opacity")
	return opacity
}

func (*CSSStyleDeclaration) SetOpacity(opacity string) {
	js.Rewrite("$<.opacity = $1", opacity)
}

func (*CSSStyleDeclaration) GetOrder() (order string) {
	js.Rewrite("$<.order")
	return order
}

func (*CSSStyleDeclaration) SetOrder(order string) {
	js.Rewrite("$<.order = $1", order)
}

func (*CSSStyleDeclaration) GetOrphans() (orphans string) {
	js.Rewrite("$<.orphans")
	return orphans
}

func (*CSSStyleDeclaration) SetOrphans(orphans string) {
	js.Rewrite("$<.orphans = $1", orphans)
}

func (*CSSStyleDeclaration) GetOutline() (outline string) {
	js.Rewrite("$<.outline")
	return outline
}

func (*CSSStyleDeclaration) SetOutline(outline string) {
	js.Rewrite("$<.outline = $1", outline)
}

func (*CSSStyleDeclaration) GetOutlineColor() (outlineColor string) {
	js.Rewrite("$<.outlineColor")
	return outlineColor
}

func (*CSSStyleDeclaration) SetOutlineColor(outlineColor string) {
	js.Rewrite("$<.outlineColor = $1", outlineColor)
}

func (*CSSStyleDeclaration) GetOutlineOffset() (outlineOffset string) {
	js.Rewrite("$<.outlineOffset")
	return outlineOffset
}

func (*CSSStyleDeclaration) SetOutlineOffset(outlineOffset string) {
	js.Rewrite("$<.outlineOffset = $1", outlineOffset)
}

func (*CSSStyleDeclaration) GetOutlineStyle() (outlineStyle string) {
	js.Rewrite("$<.outlineStyle")
	return outlineStyle
}

func (*CSSStyleDeclaration) SetOutlineStyle(outlineStyle string) {
	js.Rewrite("$<.outlineStyle = $1", outlineStyle)
}

func (*CSSStyleDeclaration) GetOutlineWidth() (outlineWidth string) {
	js.Rewrite("$<.outlineWidth")
	return outlineWidth
}

func (*CSSStyleDeclaration) SetOutlineWidth(outlineWidth string) {
	js.Rewrite("$<.outlineWidth = $1", outlineWidth)
}

func (*CSSStyleDeclaration) GetOverflow() (overflow string) {
	js.Rewrite("$<.overflow")
	return overflow
}

func (*CSSStyleDeclaration) SetOverflow(overflow string) {
	js.Rewrite("$<.overflow = $1", overflow)
}

func (*CSSStyleDeclaration) GetOverflowX() (overflowX string) {
	js.Rewrite("$<.overflowX")
	return overflowX
}

func (*CSSStyleDeclaration) SetOverflowX(overflowX string) {
	js.Rewrite("$<.overflowX = $1", overflowX)
}

func (*CSSStyleDeclaration) GetOverflowY() (overflowY string) {
	js.Rewrite("$<.overflowY")
	return overflowY
}

func (*CSSStyleDeclaration) SetOverflowY(overflowY string) {
	js.Rewrite("$<.overflowY = $1", overflowY)
}

func (*CSSStyleDeclaration) GetPadding() (padding string) {
	js.Rewrite("$<.padding")
	return padding
}

func (*CSSStyleDeclaration) SetPadding(padding string) {
	js.Rewrite("$<.padding = $1", padding)
}

func (*CSSStyleDeclaration) GetPaddingBottom() (paddingBottom string) {
	js.Rewrite("$<.paddingBottom")
	return paddingBottom
}

func (*CSSStyleDeclaration) SetPaddingBottom(paddingBottom string) {
	js.Rewrite("$<.paddingBottom = $1", paddingBottom)
}

func (*CSSStyleDeclaration) GetPaddingLeft() (paddingLeft string) {
	js.Rewrite("$<.paddingLeft")
	return paddingLeft
}

func (*CSSStyleDeclaration) SetPaddingLeft(paddingLeft string) {
	js.Rewrite("$<.paddingLeft = $1", paddingLeft)
}

func (*CSSStyleDeclaration) GetPaddingRight() (paddingRight string) {
	js.Rewrite("$<.paddingRight")
	return paddingRight
}

func (*CSSStyleDeclaration) SetPaddingRight(paddingRight string) {
	js.Rewrite("$<.paddingRight = $1", paddingRight)
}

func (*CSSStyleDeclaration) GetPaddingTop() (paddingTop string) {
	js.Rewrite("$<.paddingTop")
	return paddingTop
}

func (*CSSStyleDeclaration) SetPaddingTop(paddingTop string) {
	js.Rewrite("$<.paddingTop = $1", paddingTop)
}

func (*CSSStyleDeclaration) GetPageBreakAfter() (pageBreakAfter string) {
	js.Rewrite("$<.pageBreakAfter")
	return pageBreakAfter
}

func (*CSSStyleDeclaration) SetPageBreakAfter(pageBreakAfter string) {
	js.Rewrite("$<.pageBreakAfter = $1", pageBreakAfter)
}

func (*CSSStyleDeclaration) GetPageBreakBefore() (pageBreakBefore string) {
	js.Rewrite("$<.pageBreakBefore")
	return pageBreakBefore
}

func (*CSSStyleDeclaration) SetPageBreakBefore(pageBreakBefore string) {
	js.Rewrite("$<.pageBreakBefore = $1", pageBreakBefore)
}

func (*CSSStyleDeclaration) GetPageBreakInside() (pageBreakInside string) {
	js.Rewrite("$<.pageBreakInside")
	return pageBreakInside
}

func (*CSSStyleDeclaration) SetPageBreakInside(pageBreakInside string) {
	js.Rewrite("$<.pageBreakInside = $1", pageBreakInside)
}

func (*CSSStyleDeclaration) GetParentRule() (parentRule *CSSRule) {
	js.Rewrite("$<.parentRule")
	return parentRule
}

func (*CSSStyleDeclaration) GetPerspective() (perspective string) {
	js.Rewrite("$<.perspective")
	return perspective
}

func (*CSSStyleDeclaration) SetPerspective(perspective string) {
	js.Rewrite("$<.perspective = $1", perspective)
}

func (*CSSStyleDeclaration) GetPerspectiveOrigin() (perspectiveOrigin string) {
	js.Rewrite("$<.perspectiveOrigin")
	return perspectiveOrigin
}

func (*CSSStyleDeclaration) SetPerspectiveOrigin(perspectiveOrigin string) {
	js.Rewrite("$<.perspectiveOrigin = $1", perspectiveOrigin)
}

func (*CSSStyleDeclaration) GetPointerEvents() (pointerEvents string) {
	js.Rewrite("$<.pointerEvents")
	return pointerEvents
}

func (*CSSStyleDeclaration) SetPointerEvents(pointerEvents string) {
	js.Rewrite("$<.pointerEvents = $1", pointerEvents)
}

func (*CSSStyleDeclaration) GetPosition() (position string) {
	js.Rewrite("$<.position")
	return position
}

func (*CSSStyleDeclaration) SetPosition(position string) {
	js.Rewrite("$<.position = $1", position)
}

func (*CSSStyleDeclaration) GetQuotes() (quotes string) {
	js.Rewrite("$<.quotes")
	return quotes
}

func (*CSSStyleDeclaration) SetQuotes(quotes string) {
	js.Rewrite("$<.quotes = $1", quotes)
}

func (*CSSStyleDeclaration) GetRight() (right string) {
	js.Rewrite("$<.right")
	return right
}

func (*CSSStyleDeclaration) SetRight(right string) {
	js.Rewrite("$<.right = $1", right)
}

func (*CSSStyleDeclaration) GetRotate() (rotate string) {
	js.Rewrite("$<.rotate")
	return rotate
}

func (*CSSStyleDeclaration) SetRotate(rotate string) {
	js.Rewrite("$<.rotate = $1", rotate)
}

func (*CSSStyleDeclaration) GetRubyAlign() (rubyAlign string) {
	js.Rewrite("$<.rubyAlign")
	return rubyAlign
}

func (*CSSStyleDeclaration) SetRubyAlign(rubyAlign string) {
	js.Rewrite("$<.rubyAlign = $1", rubyAlign)
}

func (*CSSStyleDeclaration) GetRubyOverhang() (rubyOverhang string) {
	js.Rewrite("$<.rubyOverhang")
	return rubyOverhang
}

func (*CSSStyleDeclaration) SetRubyOverhang(rubyOverhang string) {
	js.Rewrite("$<.rubyOverhang = $1", rubyOverhang)
}

func (*CSSStyleDeclaration) GetRubyPosition() (rubyPosition string) {
	js.Rewrite("$<.rubyPosition")
	return rubyPosition
}

func (*CSSStyleDeclaration) SetRubyPosition(rubyPosition string) {
	js.Rewrite("$<.rubyPosition = $1", rubyPosition)
}

func (*CSSStyleDeclaration) GetScale() (scale string) {
	js.Rewrite("$<.scale")
	return scale
}

func (*CSSStyleDeclaration) SetScale(scale string) {
	js.Rewrite("$<.scale = $1", scale)
}

func (*CSSStyleDeclaration) GetStopColor() (stopColor string) {
	js.Rewrite("$<.stopColor")
	return stopColor
}

func (*CSSStyleDeclaration) SetStopColor(stopColor string) {
	js.Rewrite("$<.stopColor = $1", stopColor)
}

func (*CSSStyleDeclaration) GetStopOpacity() (stopOpacity string) {
	js.Rewrite("$<.stopOpacity")
	return stopOpacity
}

func (*CSSStyleDeclaration) SetStopOpacity(stopOpacity string) {
	js.Rewrite("$<.stopOpacity = $1", stopOpacity)
}

func (*CSSStyleDeclaration) GetStroke() (stroke string) {
	js.Rewrite("$<.stroke")
	return stroke
}

func (*CSSStyleDeclaration) SetStroke(stroke string) {
	js.Rewrite("$<.stroke = $1", stroke)
}

func (*CSSStyleDeclaration) GetStrokeDasharray() (strokeDasharray string) {
	js.Rewrite("$<.strokeDasharray")
	return strokeDasharray
}

func (*CSSStyleDeclaration) SetStrokeDasharray(strokeDasharray string) {
	js.Rewrite("$<.strokeDasharray = $1", strokeDasharray)
}

func (*CSSStyleDeclaration) GetStrokeDashoffset() (strokeDashoffset string) {
	js.Rewrite("$<.strokeDashoffset")
	return strokeDashoffset
}

func (*CSSStyleDeclaration) SetStrokeDashoffset(strokeDashoffset string) {
	js.Rewrite("$<.strokeDashoffset = $1", strokeDashoffset)
}

func (*CSSStyleDeclaration) GetStrokeLinecap() (strokeLinecap string) {
	js.Rewrite("$<.strokeLinecap")
	return strokeLinecap
}

func (*CSSStyleDeclaration) SetStrokeLinecap(strokeLinecap string) {
	js.Rewrite("$<.strokeLinecap = $1", strokeLinecap)
}

func (*CSSStyleDeclaration) GetStrokeLinejoin() (strokeLinejoin string) {
	js.Rewrite("$<.strokeLinejoin")
	return strokeLinejoin
}

func (*CSSStyleDeclaration) SetStrokeLinejoin(strokeLinejoin string) {
	js.Rewrite("$<.strokeLinejoin = $1", strokeLinejoin)
}

func (*CSSStyleDeclaration) GetStrokeMiterlimit() (strokeMiterlimit string) {
	js.Rewrite("$<.strokeMiterlimit")
	return strokeMiterlimit
}

func (*CSSStyleDeclaration) SetStrokeMiterlimit(strokeMiterlimit string) {
	js.Rewrite("$<.strokeMiterlimit = $1", strokeMiterlimit)
}

func (*CSSStyleDeclaration) GetStrokeOpacity() (strokeOpacity string) {
	js.Rewrite("$<.strokeOpacity")
	return strokeOpacity
}

func (*CSSStyleDeclaration) SetStrokeOpacity(strokeOpacity string) {
	js.Rewrite("$<.strokeOpacity = $1", strokeOpacity)
}

func (*CSSStyleDeclaration) GetStrokeWidth() (strokeWidth string) {
	js.Rewrite("$<.strokeWidth")
	return strokeWidth
}

func (*CSSStyleDeclaration) SetStrokeWidth(strokeWidth string) {
	js.Rewrite("$<.strokeWidth = $1", strokeWidth)
}

func (*CSSStyleDeclaration) GetTableLayout() (tableLayout string) {
	js.Rewrite("$<.tableLayout")
	return tableLayout
}

func (*CSSStyleDeclaration) SetTableLayout(tableLayout string) {
	js.Rewrite("$<.tableLayout = $1", tableLayout)
}

func (*CSSStyleDeclaration) GetTextAlign() (textAlign string) {
	js.Rewrite("$<.textAlign")
	return textAlign
}

func (*CSSStyleDeclaration) SetTextAlign(textAlign string) {
	js.Rewrite("$<.textAlign = $1", textAlign)
}

func (*CSSStyleDeclaration) GetTextAlignLast() (textAlignLast string) {
	js.Rewrite("$<.textAlignLast")
	return textAlignLast
}

func (*CSSStyleDeclaration) SetTextAlignLast(textAlignLast string) {
	js.Rewrite("$<.textAlignLast = $1", textAlignLast)
}

func (*CSSStyleDeclaration) GetTextAnchor() (textAnchor string) {
	js.Rewrite("$<.textAnchor")
	return textAnchor
}

func (*CSSStyleDeclaration) SetTextAnchor(textAnchor string) {
	js.Rewrite("$<.textAnchor = $1", textAnchor)
}

func (*CSSStyleDeclaration) GetTextDecoration() (textDecoration string) {
	js.Rewrite("$<.textDecoration")
	return textDecoration
}

func (*CSSStyleDeclaration) SetTextDecoration(textDecoration string) {
	js.Rewrite("$<.textDecoration = $1", textDecoration)
}

func (*CSSStyleDeclaration) GetTextIndent() (textIndent string) {
	js.Rewrite("$<.textIndent")
	return textIndent
}

func (*CSSStyleDeclaration) SetTextIndent(textIndent string) {
	js.Rewrite("$<.textIndent = $1", textIndent)
}

func (*CSSStyleDeclaration) GetTextJustify() (textJustify string) {
	js.Rewrite("$<.textJustify")
	return textJustify
}

func (*CSSStyleDeclaration) SetTextJustify(textJustify string) {
	js.Rewrite("$<.textJustify = $1", textJustify)
}

func (*CSSStyleDeclaration) GetTextKashida() (textKashida string) {
	js.Rewrite("$<.textKashida")
	return textKashida
}

func (*CSSStyleDeclaration) SetTextKashida(textKashida string) {
	js.Rewrite("$<.textKashida = $1", textKashida)
}

func (*CSSStyleDeclaration) GetTextKashidaSpace() (textKashidaSpace string) {
	js.Rewrite("$<.textKashidaSpace")
	return textKashidaSpace
}

func (*CSSStyleDeclaration) SetTextKashidaSpace(textKashidaSpace string) {
	js.Rewrite("$<.textKashidaSpace = $1", textKashidaSpace)
}

func (*CSSStyleDeclaration) GetTextOverflow() (textOverflow string) {
	js.Rewrite("$<.textOverflow")
	return textOverflow
}

func (*CSSStyleDeclaration) SetTextOverflow(textOverflow string) {
	js.Rewrite("$<.textOverflow = $1", textOverflow)
}

func (*CSSStyleDeclaration) GetTextShadow() (textShadow string) {
	js.Rewrite("$<.textShadow")
	return textShadow
}

func (*CSSStyleDeclaration) SetTextShadow(textShadow string) {
	js.Rewrite("$<.textShadow = $1", textShadow)
}

func (*CSSStyleDeclaration) GetTextTransform() (textTransform string) {
	js.Rewrite("$<.textTransform")
	return textTransform
}

func (*CSSStyleDeclaration) SetTextTransform(textTransform string) {
	js.Rewrite("$<.textTransform = $1", textTransform)
}

func (*CSSStyleDeclaration) GetTextUnderlinePosition() (textUnderlinePosition string) {
	js.Rewrite("$<.textUnderlinePosition")
	return textUnderlinePosition
}

func (*CSSStyleDeclaration) SetTextUnderlinePosition(textUnderlinePosition string) {
	js.Rewrite("$<.textUnderlinePosition = $1", textUnderlinePosition)
}

func (*CSSStyleDeclaration) GetTop() (top string) {
	js.Rewrite("$<.top")
	return top
}

func (*CSSStyleDeclaration) SetTop(top string) {
	js.Rewrite("$<.top = $1", top)
}

func (*CSSStyleDeclaration) GetTouchAction() (touchAction string) {
	js.Rewrite("$<.touchAction")
	return touchAction
}

func (*CSSStyleDeclaration) SetTouchAction(touchAction string) {
	js.Rewrite("$<.touchAction = $1", touchAction)
}

func (*CSSStyleDeclaration) GetTransform() (transform string) {
	js.Rewrite("$<.transform")
	return transform
}

func (*CSSStyleDeclaration) SetTransform(transform string) {
	js.Rewrite("$<.transform = $1", transform)
}

func (*CSSStyleDeclaration) GetTransformOrigin() (transformOrigin string) {
	js.Rewrite("$<.transformOrigin")
	return transformOrigin
}

func (*CSSStyleDeclaration) SetTransformOrigin(transformOrigin string) {
	js.Rewrite("$<.transformOrigin = $1", transformOrigin)
}

func (*CSSStyleDeclaration) GetTransformStyle() (transformStyle string) {
	js.Rewrite("$<.transformStyle")
	return transformStyle
}

func (*CSSStyleDeclaration) SetTransformStyle(transformStyle string) {
	js.Rewrite("$<.transformStyle = $1", transformStyle)
}

func (*CSSStyleDeclaration) GetTransition() (transition string) {
	js.Rewrite("$<.transition")
	return transition
}

func (*CSSStyleDeclaration) SetTransition(transition string) {
	js.Rewrite("$<.transition = $1", transition)
}

func (*CSSStyleDeclaration) GetTransitionDelay() (transitionDelay string) {
	js.Rewrite("$<.transitionDelay")
	return transitionDelay
}

func (*CSSStyleDeclaration) SetTransitionDelay(transitionDelay string) {
	js.Rewrite("$<.transitionDelay = $1", transitionDelay)
}

func (*CSSStyleDeclaration) GetTransitionDuration() (transitionDuration string) {
	js.Rewrite("$<.transitionDuration")
	return transitionDuration
}

func (*CSSStyleDeclaration) SetTransitionDuration(transitionDuration string) {
	js.Rewrite("$<.transitionDuration = $1", transitionDuration)
}

func (*CSSStyleDeclaration) GetTransitionProperty() (transitionProperty string) {
	js.Rewrite("$<.transitionProperty")
	return transitionProperty
}

func (*CSSStyleDeclaration) SetTransitionProperty(transitionProperty string) {
	js.Rewrite("$<.transitionProperty = $1", transitionProperty)
}

func (*CSSStyleDeclaration) GetTransitionTimingFunction() (transitionTimingFunction string) {
	js.Rewrite("$<.transitionTimingFunction")
	return transitionTimingFunction
}

func (*CSSStyleDeclaration) SetTransitionTimingFunction(transitionTimingFunction string) {
	js.Rewrite("$<.transitionTimingFunction = $1", transitionTimingFunction)
}

func (*CSSStyleDeclaration) GetTranslate() (translate string) {
	js.Rewrite("$<.translate")
	return translate
}

func (*CSSStyleDeclaration) SetTranslate(translate string) {
	js.Rewrite("$<.translate = $1", translate)
}

func (*CSSStyleDeclaration) GetUnicodeBidi() (unicodeBidi string) {
	js.Rewrite("$<.unicodeBidi")
	return unicodeBidi
}

func (*CSSStyleDeclaration) SetUnicodeBidi(unicodeBidi string) {
	js.Rewrite("$<.unicodeBidi = $1", unicodeBidi)
}

func (*CSSStyleDeclaration) GetVerticalAlign() (verticalAlign string) {
	js.Rewrite("$<.verticalAlign")
	return verticalAlign
}

func (*CSSStyleDeclaration) SetVerticalAlign(verticalAlign string) {
	js.Rewrite("$<.verticalAlign = $1", verticalAlign)
}

func (*CSSStyleDeclaration) GetVisibility() (visibility string) {
	js.Rewrite("$<.visibility")
	return visibility
}

func (*CSSStyleDeclaration) SetVisibility(visibility string) {
	js.Rewrite("$<.visibility = $1", visibility)
}

func (*CSSStyleDeclaration) GetWebkitAlignContent() (webkitAlignContent string) {
	js.Rewrite("$<.webkitAlignContent")
	return webkitAlignContent
}

func (*CSSStyleDeclaration) SetWebkitAlignContent(webkitAlignContent string) {
	js.Rewrite("$<.webkitAlignContent = $1", webkitAlignContent)
}

func (*CSSStyleDeclaration) GetWebkitAlignItems() (webkitAlignItems string) {
	js.Rewrite("$<.webkitAlignItems")
	return webkitAlignItems
}

func (*CSSStyleDeclaration) SetWebkitAlignItems(webkitAlignItems string) {
	js.Rewrite("$<.webkitAlignItems = $1", webkitAlignItems)
}

func (*CSSStyleDeclaration) GetWebkitAlignSelf() (webkitAlignSelf string) {
	js.Rewrite("$<.webkitAlignSelf")
	return webkitAlignSelf
}

func (*CSSStyleDeclaration) SetWebkitAlignSelf(webkitAlignSelf string) {
	js.Rewrite("$<.webkitAlignSelf = $1", webkitAlignSelf)
}

func (*CSSStyleDeclaration) GetWebkitAnimation() (webkitAnimation string) {
	js.Rewrite("$<.webkitAnimation")
	return webkitAnimation
}

func (*CSSStyleDeclaration) SetWebkitAnimation(webkitAnimation string) {
	js.Rewrite("$<.webkitAnimation = $1", webkitAnimation)
}

func (*CSSStyleDeclaration) GetWebkitAnimationDelay() (webkitAnimationDelay string) {
	js.Rewrite("$<.webkitAnimationDelay")
	return webkitAnimationDelay
}

func (*CSSStyleDeclaration) SetWebkitAnimationDelay(webkitAnimationDelay string) {
	js.Rewrite("$<.webkitAnimationDelay = $1", webkitAnimationDelay)
}

func (*CSSStyleDeclaration) GetWebkitAnimationDirection() (webkitAnimationDirection string) {
	js.Rewrite("$<.webkitAnimationDirection")
	return webkitAnimationDirection
}

func (*CSSStyleDeclaration) SetWebkitAnimationDirection(webkitAnimationDirection string) {
	js.Rewrite("$<.webkitAnimationDirection = $1", webkitAnimationDirection)
}

func (*CSSStyleDeclaration) GetWebkitAnimationDuration() (webkitAnimationDuration string) {
	js.Rewrite("$<.webkitAnimationDuration")
	return webkitAnimationDuration
}

func (*CSSStyleDeclaration) SetWebkitAnimationDuration(webkitAnimationDuration string) {
	js.Rewrite("$<.webkitAnimationDuration = $1", webkitAnimationDuration)
}

func (*CSSStyleDeclaration) GetWebkitAnimationFillMode() (webkitAnimationFillMode string) {
	js.Rewrite("$<.webkitAnimationFillMode")
	return webkitAnimationFillMode
}

func (*CSSStyleDeclaration) SetWebkitAnimationFillMode(webkitAnimationFillMode string) {
	js.Rewrite("$<.webkitAnimationFillMode = $1", webkitAnimationFillMode)
}

func (*CSSStyleDeclaration) GetWebkitAnimationIterationCount() (webkitAnimationIterationCount string) {
	js.Rewrite("$<.webkitAnimationIterationCount")
	return webkitAnimationIterationCount
}

func (*CSSStyleDeclaration) SetWebkitAnimationIterationCount(webkitAnimationIterationCount string) {
	js.Rewrite("$<.webkitAnimationIterationCount = $1", webkitAnimationIterationCount)
}

func (*CSSStyleDeclaration) GetWebkitAnimationName() (webkitAnimationName string) {
	js.Rewrite("$<.webkitAnimationName")
	return webkitAnimationName
}

func (*CSSStyleDeclaration) SetWebkitAnimationName(webkitAnimationName string) {
	js.Rewrite("$<.webkitAnimationName = $1", webkitAnimationName)
}

func (*CSSStyleDeclaration) GetWebkitAnimationPlayState() (webkitAnimationPlayState string) {
	js.Rewrite("$<.webkitAnimationPlayState")
	return webkitAnimationPlayState
}

func (*CSSStyleDeclaration) SetWebkitAnimationPlayState(webkitAnimationPlayState string) {
	js.Rewrite("$<.webkitAnimationPlayState = $1", webkitAnimationPlayState)
}

func (*CSSStyleDeclaration) GetWebkitAnimationTimingFunction() (webkitAnimationTimingFunction string) {
	js.Rewrite("$<.webkitAnimationTimingFunction")
	return webkitAnimationTimingFunction
}

func (*CSSStyleDeclaration) SetWebkitAnimationTimingFunction(webkitAnimationTimingFunction string) {
	js.Rewrite("$<.webkitAnimationTimingFunction = $1", webkitAnimationTimingFunction)
}

func (*CSSStyleDeclaration) GetWebkitAppearance() (webkitAppearance string) {
	js.Rewrite("$<.webkitAppearance")
	return webkitAppearance
}

func (*CSSStyleDeclaration) SetWebkitAppearance(webkitAppearance string) {
	js.Rewrite("$<.webkitAppearance = $1", webkitAppearance)
}

func (*CSSStyleDeclaration) GetWebkitBackfaceVisibility() (webkitBackfaceVisibility string) {
	js.Rewrite("$<.webkitBackfaceVisibility")
	return webkitBackfaceVisibility
}

func (*CSSStyleDeclaration) SetWebkitBackfaceVisibility(webkitBackfaceVisibility string) {
	js.Rewrite("$<.webkitBackfaceVisibility = $1", webkitBackfaceVisibility)
}

func (*CSSStyleDeclaration) GetWebkitBackgroundClip() (webkitBackgroundClip string) {
	js.Rewrite("$<.webkitBackgroundClip")
	return webkitBackgroundClip
}

func (*CSSStyleDeclaration) SetWebkitBackgroundClip(webkitBackgroundClip string) {
	js.Rewrite("$<.webkitBackgroundClip = $1", webkitBackgroundClip)
}

func (*CSSStyleDeclaration) GetWebkitBackgroundOrigin() (webkitBackgroundOrigin string) {
	js.Rewrite("$<.webkitBackgroundOrigin")
	return webkitBackgroundOrigin
}

func (*CSSStyleDeclaration) SetWebkitBackgroundOrigin(webkitBackgroundOrigin string) {
	js.Rewrite("$<.webkitBackgroundOrigin = $1", webkitBackgroundOrigin)
}

func (*CSSStyleDeclaration) GetWebkitBackgroundSize() (webkitBackgroundSize string) {
	js.Rewrite("$<.webkitBackgroundSize")
	return webkitBackgroundSize
}

func (*CSSStyleDeclaration) SetWebkitBackgroundSize(webkitBackgroundSize string) {
	js.Rewrite("$<.webkitBackgroundSize = $1", webkitBackgroundSize)
}

func (*CSSStyleDeclaration) GetWebkitBorderBottomLeftRadius() (webkitBorderBottomLeftRadius string) {
	js.Rewrite("$<.webkitBorderBottomLeftRadius")
	return webkitBorderBottomLeftRadius
}

func (*CSSStyleDeclaration) SetWebkitBorderBottomLeftRadius(webkitBorderBottomLeftRadius string) {
	js.Rewrite("$<.webkitBorderBottomLeftRadius = $1", webkitBorderBottomLeftRadius)
}

func (*CSSStyleDeclaration) GetWebkitBorderBottomRightRadius() (webkitBorderBottomRightRadius string) {
	js.Rewrite("$<.webkitBorderBottomRightRadius")
	return webkitBorderBottomRightRadius
}

func (*CSSStyleDeclaration) SetWebkitBorderBottomRightRadius(webkitBorderBottomRightRadius string) {
	js.Rewrite("$<.webkitBorderBottomRightRadius = $1", webkitBorderBottomRightRadius)
}

func (*CSSStyleDeclaration) GetWebkitBorderImage() (webkitBorderImage string) {
	js.Rewrite("$<.webkitBorderImage")
	return webkitBorderImage
}

func (*CSSStyleDeclaration) SetWebkitBorderImage(webkitBorderImage string) {
	js.Rewrite("$<.webkitBorderImage = $1", webkitBorderImage)
}

func (*CSSStyleDeclaration) GetWebkitBorderRadius() (webkitBorderRadius string) {
	js.Rewrite("$<.webkitBorderRadius")
	return webkitBorderRadius
}

func (*CSSStyleDeclaration) SetWebkitBorderRadius(webkitBorderRadius string) {
	js.Rewrite("$<.webkitBorderRadius = $1", webkitBorderRadius)
}

func (*CSSStyleDeclaration) GetWebkitBorderTopLeftRadius() (webkitBorderTopLeftRadius string) {
	js.Rewrite("$<.webkitBorderTopLeftRadius")
	return webkitBorderTopLeftRadius
}

func (*CSSStyleDeclaration) SetWebkitBorderTopLeftRadius(webkitBorderTopLeftRadius string) {
	js.Rewrite("$<.webkitBorderTopLeftRadius = $1", webkitBorderTopLeftRadius)
}

func (*CSSStyleDeclaration) GetWebkitBorderTopRightRadius() (webkitBorderTopRightRadius string) {
	js.Rewrite("$<.webkitBorderTopRightRadius")
	return webkitBorderTopRightRadius
}

func (*CSSStyleDeclaration) SetWebkitBorderTopRightRadius(webkitBorderTopRightRadius string) {
	js.Rewrite("$<.webkitBorderTopRightRadius = $1", webkitBorderTopRightRadius)
}

func (*CSSStyleDeclaration) GetWebkitBoxAlign() (webkitBoxAlign string) {
	js.Rewrite("$<.webkitBoxAlign")
	return webkitBoxAlign
}

func (*CSSStyleDeclaration) SetWebkitBoxAlign(webkitBoxAlign string) {
	js.Rewrite("$<.webkitBoxAlign = $1", webkitBoxAlign)
}

func (*CSSStyleDeclaration) GetWebkitBoxDirection() (webkitBoxDirection string) {
	js.Rewrite("$<.webkitBoxDirection")
	return webkitBoxDirection
}

func (*CSSStyleDeclaration) SetWebkitBoxDirection(webkitBoxDirection string) {
	js.Rewrite("$<.webkitBoxDirection = $1", webkitBoxDirection)
}

func (*CSSStyleDeclaration) GetWebkitBoxFlex() (webkitBoxFlex string) {
	js.Rewrite("$<.webkitBoxFlex")
	return webkitBoxFlex
}

func (*CSSStyleDeclaration) SetWebkitBoxFlex(webkitBoxFlex string) {
	js.Rewrite("$<.webkitBoxFlex = $1", webkitBoxFlex)
}

func (*CSSStyleDeclaration) GetWebkitBoxOrdinalGroup() (webkitBoxOrdinalGroup string) {
	js.Rewrite("$<.webkitBoxOrdinalGroup")
	return webkitBoxOrdinalGroup
}

func (*CSSStyleDeclaration) SetWebkitBoxOrdinalGroup(webkitBoxOrdinalGroup string) {
	js.Rewrite("$<.webkitBoxOrdinalGroup = $1", webkitBoxOrdinalGroup)
}

func (*CSSStyleDeclaration) GetWebkitBoxOrient() (webkitBoxOrient string) {
	js.Rewrite("$<.webkitBoxOrient")
	return webkitBoxOrient
}

func (*CSSStyleDeclaration) SetWebkitBoxOrient(webkitBoxOrient string) {
	js.Rewrite("$<.webkitBoxOrient = $1", webkitBoxOrient)
}

func (*CSSStyleDeclaration) GetWebkitBoxPack() (webkitBoxPack string) {
	js.Rewrite("$<.webkitBoxPack")
	return webkitBoxPack
}

func (*CSSStyleDeclaration) SetWebkitBoxPack(webkitBoxPack string) {
	js.Rewrite("$<.webkitBoxPack = $1", webkitBoxPack)
}

func (*CSSStyleDeclaration) GetWebkitBoxSizing() (webkitBoxSizing string) {
	js.Rewrite("$<.webkitBoxSizing")
	return webkitBoxSizing
}

func (*CSSStyleDeclaration) SetWebkitBoxSizing(webkitBoxSizing string) {
	js.Rewrite("$<.webkitBoxSizing = $1", webkitBoxSizing)
}

func (*CSSStyleDeclaration) GetWebkitColumnBreakAfter() (webkitColumnBreakAfter string) {
	js.Rewrite("$<.webkitColumnBreakAfter")
	return webkitColumnBreakAfter
}

func (*CSSStyleDeclaration) SetWebkitColumnBreakAfter(webkitColumnBreakAfter string) {
	js.Rewrite("$<.webkitColumnBreakAfter = $1", webkitColumnBreakAfter)
}

func (*CSSStyleDeclaration) GetWebkitColumnBreakBefore() (webkitColumnBreakBefore string) {
	js.Rewrite("$<.webkitColumnBreakBefore")
	return webkitColumnBreakBefore
}

func (*CSSStyleDeclaration) SetWebkitColumnBreakBefore(webkitColumnBreakBefore string) {
	js.Rewrite("$<.webkitColumnBreakBefore = $1", webkitColumnBreakBefore)
}

func (*CSSStyleDeclaration) GetWebkitColumnBreakInside() (webkitColumnBreakInside string) {
	js.Rewrite("$<.webkitColumnBreakInside")
	return webkitColumnBreakInside
}

func (*CSSStyleDeclaration) SetWebkitColumnBreakInside(webkitColumnBreakInside string) {
	js.Rewrite("$<.webkitColumnBreakInside = $1", webkitColumnBreakInside)
}

func (*CSSStyleDeclaration) GetWebkitColumnCount() (webkitColumnCount interface{}) {
	js.Rewrite("$<.webkitColumnCount")
	return webkitColumnCount
}

func (*CSSStyleDeclaration) SetWebkitColumnCount(webkitColumnCount interface{}) {
	js.Rewrite("$<.webkitColumnCount = $1", webkitColumnCount)
}

func (*CSSStyleDeclaration) GetWebkitColumnGap() (webkitColumnGap interface{}) {
	js.Rewrite("$<.webkitColumnGap")
	return webkitColumnGap
}

func (*CSSStyleDeclaration) SetWebkitColumnGap(webkitColumnGap interface{}) {
	js.Rewrite("$<.webkitColumnGap = $1", webkitColumnGap)
}

func (*CSSStyleDeclaration) GetWebkitColumnRule() (webkitColumnRule string) {
	js.Rewrite("$<.webkitColumnRule")
	return webkitColumnRule
}

func (*CSSStyleDeclaration) SetWebkitColumnRule(webkitColumnRule string) {
	js.Rewrite("$<.webkitColumnRule = $1", webkitColumnRule)
}

func (*CSSStyleDeclaration) GetWebkitColumnRuleColor() (webkitColumnRuleColor interface{}) {
	js.Rewrite("$<.webkitColumnRuleColor")
	return webkitColumnRuleColor
}

func (*CSSStyleDeclaration) SetWebkitColumnRuleColor(webkitColumnRuleColor interface{}) {
	js.Rewrite("$<.webkitColumnRuleColor = $1", webkitColumnRuleColor)
}

func (*CSSStyleDeclaration) GetWebkitColumnRuleStyle() (webkitColumnRuleStyle string) {
	js.Rewrite("$<.webkitColumnRuleStyle")
	return webkitColumnRuleStyle
}

func (*CSSStyleDeclaration) SetWebkitColumnRuleStyle(webkitColumnRuleStyle string) {
	js.Rewrite("$<.webkitColumnRuleStyle = $1", webkitColumnRuleStyle)
}

func (*CSSStyleDeclaration) GetWebkitColumnRuleWidth() (webkitColumnRuleWidth interface{}) {
	js.Rewrite("$<.webkitColumnRuleWidth")
	return webkitColumnRuleWidth
}

func (*CSSStyleDeclaration) SetWebkitColumnRuleWidth(webkitColumnRuleWidth interface{}) {
	js.Rewrite("$<.webkitColumnRuleWidth = $1", webkitColumnRuleWidth)
}

func (*CSSStyleDeclaration) GetWebkitColumns() (webkitColumns string) {
	js.Rewrite("$<.webkitColumns")
	return webkitColumns
}

func (*CSSStyleDeclaration) SetWebkitColumns(webkitColumns string) {
	js.Rewrite("$<.webkitColumns = $1", webkitColumns)
}

func (*CSSStyleDeclaration) GetWebkitColumnSpan() (webkitColumnSpan string) {
	js.Rewrite("$<.webkitColumnSpan")
	return webkitColumnSpan
}

func (*CSSStyleDeclaration) SetWebkitColumnSpan(webkitColumnSpan string) {
	js.Rewrite("$<.webkitColumnSpan = $1", webkitColumnSpan)
}

func (*CSSStyleDeclaration) GetWebkitColumnWidth() (webkitColumnWidth interface{}) {
	js.Rewrite("$<.webkitColumnWidth")
	return webkitColumnWidth
}

func (*CSSStyleDeclaration) SetWebkitColumnWidth(webkitColumnWidth interface{}) {
	js.Rewrite("$<.webkitColumnWidth = $1", webkitColumnWidth)
}

func (*CSSStyleDeclaration) GetWebkitFilter() (webkitFilter string) {
	js.Rewrite("$<.webkitFilter")
	return webkitFilter
}

func (*CSSStyleDeclaration) SetWebkitFilter(webkitFilter string) {
	js.Rewrite("$<.webkitFilter = $1", webkitFilter)
}

func (*CSSStyleDeclaration) GetWebkitFlex() (webkitFlex string) {
	js.Rewrite("$<.webkitFlex")
	return webkitFlex
}

func (*CSSStyleDeclaration) SetWebkitFlex(webkitFlex string) {
	js.Rewrite("$<.webkitFlex = $1", webkitFlex)
}

func (*CSSStyleDeclaration) GetWebkitFlexBasis() (webkitFlexBasis string) {
	js.Rewrite("$<.webkitFlexBasis")
	return webkitFlexBasis
}

func (*CSSStyleDeclaration) SetWebkitFlexBasis(webkitFlexBasis string) {
	js.Rewrite("$<.webkitFlexBasis = $1", webkitFlexBasis)
}

func (*CSSStyleDeclaration) GetWebkitFlexDirection() (webkitFlexDirection string) {
	js.Rewrite("$<.webkitFlexDirection")
	return webkitFlexDirection
}

func (*CSSStyleDeclaration) SetWebkitFlexDirection(webkitFlexDirection string) {
	js.Rewrite("$<.webkitFlexDirection = $1", webkitFlexDirection)
}

func (*CSSStyleDeclaration) GetWebkitFlexFlow() (webkitFlexFlow string) {
	js.Rewrite("$<.webkitFlexFlow")
	return webkitFlexFlow
}

func (*CSSStyleDeclaration) SetWebkitFlexFlow(webkitFlexFlow string) {
	js.Rewrite("$<.webkitFlexFlow = $1", webkitFlexFlow)
}

func (*CSSStyleDeclaration) GetWebkitFlexGrow() (webkitFlexGrow string) {
	js.Rewrite("$<.webkitFlexGrow")
	return webkitFlexGrow
}

func (*CSSStyleDeclaration) SetWebkitFlexGrow(webkitFlexGrow string) {
	js.Rewrite("$<.webkitFlexGrow = $1", webkitFlexGrow)
}

func (*CSSStyleDeclaration) GetWebkitFlexShrink() (webkitFlexShrink string) {
	js.Rewrite("$<.webkitFlexShrink")
	return webkitFlexShrink
}

func (*CSSStyleDeclaration) SetWebkitFlexShrink(webkitFlexShrink string) {
	js.Rewrite("$<.webkitFlexShrink = $1", webkitFlexShrink)
}

func (*CSSStyleDeclaration) GetWebkitFlexWrap() (webkitFlexWrap string) {
	js.Rewrite("$<.webkitFlexWrap")
	return webkitFlexWrap
}

func (*CSSStyleDeclaration) SetWebkitFlexWrap(webkitFlexWrap string) {
	js.Rewrite("$<.webkitFlexWrap = $1", webkitFlexWrap)
}

func (*CSSStyleDeclaration) GetWebkitJustifyContent() (webkitJustifyContent string) {
	js.Rewrite("$<.webkitJustifyContent")
	return webkitJustifyContent
}

func (*CSSStyleDeclaration) SetWebkitJustifyContent(webkitJustifyContent string) {
	js.Rewrite("$<.webkitJustifyContent = $1", webkitJustifyContent)
}

func (*CSSStyleDeclaration) GetWebkitOrder() (webkitOrder string) {
	js.Rewrite("$<.webkitOrder")
	return webkitOrder
}

func (*CSSStyleDeclaration) SetWebkitOrder(webkitOrder string) {
	js.Rewrite("$<.webkitOrder = $1", webkitOrder)
}

func (*CSSStyleDeclaration) GetWebkitPerspective() (webkitPerspective string) {
	js.Rewrite("$<.webkitPerspective")
	return webkitPerspective
}

func (*CSSStyleDeclaration) SetWebkitPerspective(webkitPerspective string) {
	js.Rewrite("$<.webkitPerspective = $1", webkitPerspective)
}

func (*CSSStyleDeclaration) GetWebkitPerspectiveOrigin() (webkitPerspectiveOrigin string) {
	js.Rewrite("$<.webkitPerspectiveOrigin")
	return webkitPerspectiveOrigin
}

func (*CSSStyleDeclaration) SetWebkitPerspectiveOrigin(webkitPerspectiveOrigin string) {
	js.Rewrite("$<.webkitPerspectiveOrigin = $1", webkitPerspectiveOrigin)
}

func (*CSSStyleDeclaration) GetWebkitTapHighlightColor() (webkitTapHighlightColor string) {
	js.Rewrite("$<.webkitTapHighlightColor")
	return webkitTapHighlightColor
}

func (*CSSStyleDeclaration) SetWebkitTapHighlightColor(webkitTapHighlightColor string) {
	js.Rewrite("$<.webkitTapHighlightColor = $1", webkitTapHighlightColor)
}

func (*CSSStyleDeclaration) GetWebkitTextFillColor() (webkitTextFillColor string) {
	js.Rewrite("$<.webkitTextFillColor")
	return webkitTextFillColor
}

func (*CSSStyleDeclaration) SetWebkitTextFillColor(webkitTextFillColor string) {
	js.Rewrite("$<.webkitTextFillColor = $1", webkitTextFillColor)
}

func (*CSSStyleDeclaration) GetWebkitTextSizeAdjust() (webkitTextSizeAdjust interface{}) {
	js.Rewrite("$<.webkitTextSizeAdjust")
	return webkitTextSizeAdjust
}

func (*CSSStyleDeclaration) SetWebkitTextSizeAdjust(webkitTextSizeAdjust interface{}) {
	js.Rewrite("$<.webkitTextSizeAdjust = $1", webkitTextSizeAdjust)
}

func (*CSSStyleDeclaration) GetWebkitTextStroke() (webkitTextStroke string) {
	js.Rewrite("$<.webkitTextStroke")
	return webkitTextStroke
}

func (*CSSStyleDeclaration) SetWebkitTextStroke(webkitTextStroke string) {
	js.Rewrite("$<.webkitTextStroke = $1", webkitTextStroke)
}

func (*CSSStyleDeclaration) GetWebkitTextStrokeColor() (webkitTextStrokeColor string) {
	js.Rewrite("$<.webkitTextStrokeColor")
	return webkitTextStrokeColor
}

func (*CSSStyleDeclaration) SetWebkitTextStrokeColor(webkitTextStrokeColor string) {
	js.Rewrite("$<.webkitTextStrokeColor = $1", webkitTextStrokeColor)
}

func (*CSSStyleDeclaration) GetWebkitTextStrokeWidth() (webkitTextStrokeWidth string) {
	js.Rewrite("$<.webkitTextStrokeWidth")
	return webkitTextStrokeWidth
}

func (*CSSStyleDeclaration) SetWebkitTextStrokeWidth(webkitTextStrokeWidth string) {
	js.Rewrite("$<.webkitTextStrokeWidth = $1", webkitTextStrokeWidth)
}

func (*CSSStyleDeclaration) GetWebkitTransform() (webkitTransform string) {
	js.Rewrite("$<.webkitTransform")
	return webkitTransform
}

func (*CSSStyleDeclaration) SetWebkitTransform(webkitTransform string) {
	js.Rewrite("$<.webkitTransform = $1", webkitTransform)
}

func (*CSSStyleDeclaration) GetWebkitTransformOrigin() (webkitTransformOrigin string) {
	js.Rewrite("$<.webkitTransformOrigin")
	return webkitTransformOrigin
}

func (*CSSStyleDeclaration) SetWebkitTransformOrigin(webkitTransformOrigin string) {
	js.Rewrite("$<.webkitTransformOrigin = $1", webkitTransformOrigin)
}

func (*CSSStyleDeclaration) GetWebkitTransformStyle() (webkitTransformStyle string) {
	js.Rewrite("$<.webkitTransformStyle")
	return webkitTransformStyle
}

func (*CSSStyleDeclaration) SetWebkitTransformStyle(webkitTransformStyle string) {
	js.Rewrite("$<.webkitTransformStyle = $1", webkitTransformStyle)
}

func (*CSSStyleDeclaration) GetWebkitTransition() (webkitTransition string) {
	js.Rewrite("$<.webkitTransition")
	return webkitTransition
}

func (*CSSStyleDeclaration) SetWebkitTransition(webkitTransition string) {
	js.Rewrite("$<.webkitTransition = $1", webkitTransition)
}

func (*CSSStyleDeclaration) GetWebkitTransitionDelay() (webkitTransitionDelay string) {
	js.Rewrite("$<.webkitTransitionDelay")
	return webkitTransitionDelay
}

func (*CSSStyleDeclaration) SetWebkitTransitionDelay(webkitTransitionDelay string) {
	js.Rewrite("$<.webkitTransitionDelay = $1", webkitTransitionDelay)
}

func (*CSSStyleDeclaration) GetWebkitTransitionDuration() (webkitTransitionDuration string) {
	js.Rewrite("$<.webkitTransitionDuration")
	return webkitTransitionDuration
}

func (*CSSStyleDeclaration) SetWebkitTransitionDuration(webkitTransitionDuration string) {
	js.Rewrite("$<.webkitTransitionDuration = $1", webkitTransitionDuration)
}

func (*CSSStyleDeclaration) GetWebkitTransitionProperty() (webkitTransitionProperty string) {
	js.Rewrite("$<.webkitTransitionProperty")
	return webkitTransitionProperty
}

func (*CSSStyleDeclaration) SetWebkitTransitionProperty(webkitTransitionProperty string) {
	js.Rewrite("$<.webkitTransitionProperty = $1", webkitTransitionProperty)
}

func (*CSSStyleDeclaration) GetWebkitTransitionTimingFunction() (webkitTransitionTimingFunction string) {
	js.Rewrite("$<.webkitTransitionTimingFunction")
	return webkitTransitionTimingFunction
}

func (*CSSStyleDeclaration) SetWebkitTransitionTimingFunction(webkitTransitionTimingFunction string) {
	js.Rewrite("$<.webkitTransitionTimingFunction = $1", webkitTransitionTimingFunction)
}

func (*CSSStyleDeclaration) GetWebkitUserModify() (webkitUserModify string) {
	js.Rewrite("$<.webkitUserModify")
	return webkitUserModify
}

func (*CSSStyleDeclaration) SetWebkitUserModify(webkitUserModify string) {
	js.Rewrite("$<.webkitUserModify = $1", webkitUserModify)
}

func (*CSSStyleDeclaration) GetWebkitUserSelect() (webkitUserSelect string) {
	js.Rewrite("$<.webkitUserSelect")
	return webkitUserSelect
}

func (*CSSStyleDeclaration) SetWebkitUserSelect(webkitUserSelect string) {
	js.Rewrite("$<.webkitUserSelect = $1", webkitUserSelect)
}

func (*CSSStyleDeclaration) GetWebkitWritingMode() (webkitWritingMode string) {
	js.Rewrite("$<.webkitWritingMode")
	return webkitWritingMode
}

func (*CSSStyleDeclaration) SetWebkitWritingMode(webkitWritingMode string) {
	js.Rewrite("$<.webkitWritingMode = $1", webkitWritingMode)
}

func (*CSSStyleDeclaration) GetWhiteSpace() (whiteSpace string) {
	js.Rewrite("$<.whiteSpace")
	return whiteSpace
}

func (*CSSStyleDeclaration) SetWhiteSpace(whiteSpace string) {
	js.Rewrite("$<.whiteSpace = $1", whiteSpace)
}

func (*CSSStyleDeclaration) GetWidows() (widows string) {
	js.Rewrite("$<.widows")
	return widows
}

func (*CSSStyleDeclaration) SetWidows(widows string) {
	js.Rewrite("$<.widows = $1", widows)
}

func (*CSSStyleDeclaration) GetWidth() (width string) {
	js.Rewrite("$<.width")
	return width
}

func (*CSSStyleDeclaration) SetWidth(width string) {
	js.Rewrite("$<.width = $1", width)
}

func (*CSSStyleDeclaration) GetWordBreak() (wordBreak string) {
	js.Rewrite("$<.wordBreak")
	return wordBreak
}

func (*CSSStyleDeclaration) SetWordBreak(wordBreak string) {
	js.Rewrite("$<.wordBreak = $1", wordBreak)
}

func (*CSSStyleDeclaration) GetWordSpacing() (wordSpacing string) {
	js.Rewrite("$<.wordSpacing")
	return wordSpacing
}

func (*CSSStyleDeclaration) SetWordSpacing(wordSpacing string) {
	js.Rewrite("$<.wordSpacing = $1", wordSpacing)
}

func (*CSSStyleDeclaration) GetWordWrap() (wordWrap string) {
	js.Rewrite("$<.wordWrap")
	return wordWrap
}

func (*CSSStyleDeclaration) SetWordWrap(wordWrap string) {
	js.Rewrite("$<.wordWrap = $1", wordWrap)
}

func (*CSSStyleDeclaration) GetWritingMode() (writingMode string) {
	js.Rewrite("$<.writingMode")
	return writingMode
}

func (*CSSStyleDeclaration) SetWritingMode(writingMode string) {
	js.Rewrite("$<.writingMode = $1", writingMode)
}

func (*CSSStyleDeclaration) GetZIndex() (zIndex string) {
	js.Rewrite("$<.zIndex")
	return zIndex
}

func (*CSSStyleDeclaration) SetZIndex(zIndex string) {
	js.Rewrite("$<.zIndex = $1", zIndex)
}

func (*CSSStyleDeclaration) GetZoom() (zoom string) {
	js.Rewrite("$<.zoom")
	return zoom
}

func (*CSSStyleDeclaration) SetZoom(zoom string) {
	js.Rewrite("$<.zoom = $1", zoom)
}

type CSSStyleRule struct {
	*CSSRule
}

func (*CSSStyleRule) GetReadOnly() (readOnly bool) {
	js.Rewrite("$<.readOnly")
	return readOnly
}

func (*CSSStyleRule) GetSelectorText() (selectorText string) {
	js.Rewrite("$<.selectorText")
	return selectorText
}

func (*CSSStyleRule) SetSelectorText(selectorText string) {
	js.Rewrite("$<.selectorText = $1", selectorText)
}

func (*CSSStyleRule) GetStyle() (style *CSSStyleDeclaration) {
	js.Rewrite("$<.style")
	return style
}

type CSSStyleSheet struct {
	*StyleSheet
}

func (*CSSStyleSheet) AddImport(bstrURL string, lIndex int) (i int) {
	js.Rewrite("$<.addImport($1, $2)", bstrURL, lIndex)
	return i
}

func (*CSSStyleSheet) AddPageRule(bstrSelector string, bstrStyle string, lIndex int) (i int) {
	js.Rewrite("$<.addPageRule($1, $2, $3)", bstrSelector, bstrStyle, lIndex)
	return i
}

func (*CSSStyleSheet) AddRule(bstrSelector string, bstrStyle string, lIndex int) (i int) {
	js.Rewrite("$<.addRule($1, $2, $3)", bstrSelector, bstrStyle, lIndex)
	return i
}

func (*CSSStyleSheet) DeleteRule(index uint) {
	js.Rewrite("$<.deleteRule($1)", index)
}

func (*CSSStyleSheet) InsertRule(rule string, index uint) (u uint) {
	js.Rewrite("$<.insertRule($1, $2)", rule, index)
	return u
}

func (*CSSStyleSheet) RemoveImport(lIndex int) {
	js.Rewrite("$<.removeImport($1)", lIndex)
}

func (*CSSStyleSheet) RemoveRule(lIndex int) {
	js.Rewrite("$<.removeRule($1)", lIndex)
}

func (*CSSStyleSheet) GetCSSRules() (cssRules *CSSRuleList) {
	js.Rewrite("$<.cssRules")
	return cssRules
}

func (*CSSStyleSheet) GetCSSText() (cssText string) {
	js.Rewrite("$<.cssText")
	return cssText
}

func (*CSSStyleSheet) SetCSSText(cssText string) {
	js.Rewrite("$<.cssText = $1", cssText)
}

func (*CSSStyleSheet) GetID() (id string) {
	js.Rewrite("$<.id")
	return id
}

func (*CSSStyleSheet) GetImports() (imports *StyleSheetList) {
	js.Rewrite("$<.imports")
	return imports
}

func (*CSSStyleSheet) GetIsAlternate() (isAlternate bool) {
	js.Rewrite("$<.isAlternate")
	return isAlternate
}

func (*CSSStyleSheet) GetIsPrefAlternate() (isPrefAlternate bool) {
	js.Rewrite("$<.isPrefAlternate")
	return isPrefAlternate
}

func (*CSSStyleSheet) GetOwnerRule() (ownerRule *CSSRule) {
	js.Rewrite("$<.ownerRule")
	return ownerRule
}

func (*CSSStyleSheet) GetOwningElement() (owningElement *Element) {
	js.Rewrite("$<.owningElement")
	return owningElement
}

func (*CSSStyleSheet) GetPages() (pages *StyleSheetPageList) {
	js.Rewrite("$<.pages")
	return pages
}

func (*CSSStyleSheet) GetReadOnly() (readOnly bool) {
	js.Rewrite("$<.readOnly")
	return readOnly
}

func (*CSSStyleSheet) GetRules() (rules *CSSRuleList) {
	js.Rewrite("$<.rules")
	return rules
}

type CSSSupportsRule struct {
	*CSSConditionRule
}

type CustomEvent struct {
	*Event
}

func (*CustomEvent) InitCustomEvent(typeArg string, canBubbleArg bool, cancelableArg bool, detailArg interface{}) {
	js.Rewrite("$<.initCustomEvent($1, $2, $3, $4)", typeArg, canBubbleArg, cancelableArg, detailArg)
}

func (*CustomEvent) GetDetail() (detail interface{}) {
	js.Rewrite("$<.detail")
	return detail
}

type DataCue struct {
	*TextTrackCue
}

func (*DataCue) GetData() (data []byte) {
	js.Rewrite("$<.data")
	return data
}

func (*DataCue) SetData(data []byte) {
	js.Rewrite("$<.data = $1", data)
}

type DataTransfer struct {
}

func (*DataTransfer) ClearData(format string) (b bool) {
	js.Rewrite("$<.clearData($1)", format)
	return b
}

func (*DataTransfer) GetData(format string) (s string) {
	js.Rewrite("$<.getData($1)", format)
	return s
}

func (*DataTransfer) SetData(format string, data string) (b bool) {
	js.Rewrite("$<.setData($1, $2)", format, data)
	return b
}

func (*DataTransfer) GetDropEffect() (dropEffect string) {
	js.Rewrite("$<.dropEffect")
	return dropEffect
}

func (*DataTransfer) SetDropEffect(dropEffect string) {
	js.Rewrite("$<.dropEffect = $1", dropEffect)
}

func (*DataTransfer) GetEffectAllowed() (effectAllowed string) {
	js.Rewrite("$<.effectAllowed")
	return effectAllowed
}

func (*DataTransfer) SetEffectAllowed(effectAllowed string) {
	js.Rewrite("$<.effectAllowed = $1", effectAllowed)
}

func (*DataTransfer) GetFiles() (files *FileList) {
	js.Rewrite("$<.files")
	return files
}

func (*DataTransfer) GetItems() (items *DataTransferItemList) {
	js.Rewrite("$<.items")
	return items
}

func (*DataTransfer) GetTypes() (types *DOMStringList) {
	js.Rewrite("$<.types")
	return types
}

type DataTransferItem struct {
}

func (*DataTransferItem) GetAsFile() (f *File) {
	js.Rewrite("$<.getAsFile()")
	return f
}

func (*DataTransferItem) GetAsString(_callback func(data string)) {
	js.Rewrite("$<.getAsString($1)", _callback)
}

func (*DataTransferItem) WebkitGetAsEntry() (i interface{}) {
	js.Rewrite("$<.webkitGetAsEntry()")
	return i
}

func (*DataTransferItem) GetKind() (kind string) {
	js.Rewrite("$<.kind")
	return kind
}

func (*DataTransferItem) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

type DataTransferItemList struct {
}

func (*DataTransferItemList) Add(data *File) (d *DataTransferItem) {
	js.Rewrite("$<.add($1)", data)
	return d
}

func (*DataTransferItemList) Clear() {
	js.Rewrite("$<.clear()")
}

func (*DataTransferItemList) Item(index uint) (f *File) {
	js.Rewrite("$<.item($1)", index)
	return f
}

func (*DataTransferItemList) Remove(index uint) {
	js.Rewrite("$<.remove($1)", index)
}

func (*DataTransferItemList) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

type DeferredPermissionRequest struct {
}

func (*DeferredPermissionRequest) Allow() {
	js.Rewrite("$<.allow()")
}

func (*DeferredPermissionRequest) Deny() {
	js.Rewrite("$<.deny()")
}

func (*DeferredPermissionRequest) GetID() (id uint) {
	js.Rewrite("$<.id")
	return id
}

func (*DeferredPermissionRequest) GetType() (kind *MSWebViewPermissionType) {
	js.Rewrite("$<.type")
	return kind
}

func (*DeferredPermissionRequest) GetURI() (uri string) {
	js.Rewrite("$<.uri")
	return uri
}

type DelayNode struct {
	*AudioNode
}

func (*DelayNode) GetDelayTime() (delayTime *AudioParam) {
	js.Rewrite("$<.delayTime")
	return delayTime
}

type DeviceAcceleration struct {
}

func (*DeviceAcceleration) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*DeviceAcceleration) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*DeviceAcceleration) GetZ() (z float32) {
	js.Rewrite("$<.z")
	return z
}

type DeviceLightEvent struct {
	*Event
}

func (*DeviceLightEvent) GetValue() (value float32) {
	js.Rewrite("$<.value")
	return value
}

type DeviceMotionEvent struct {
	*Event
}

func (*DeviceMotionEvent) InitDeviceMotionEvent(kind string, bubbles bool, cancelable bool, acceleration *DeviceAccelerationDict, accelerationIncludingGravity *DeviceAccelerationDict, rotationRate *DeviceRotationRateDict, interval float32) {
	js.Rewrite("$<.initDeviceMotionEvent($1, $2, $3, $4, $5, $6, $7)", kind, bubbles, cancelable, acceleration, accelerationIncludingGravity, rotationRate, interval)
}

func (*DeviceMotionEvent) GetAcceleration() (acceleration *DeviceAcceleration) {
	js.Rewrite("$<.acceleration")
	return acceleration
}

func (*DeviceMotionEvent) GetAccelerationIncludingGravity() (accelerationIncludingGravity *DeviceAcceleration) {
	js.Rewrite("$<.accelerationIncludingGravity")
	return accelerationIncludingGravity
}

func (*DeviceMotionEvent) GetInterval() (interval float32) {
	js.Rewrite("$<.interval")
	return interval
}

func (*DeviceMotionEvent) GetRotationRate() (rotationRate *DeviceRotationRate) {
	js.Rewrite("$<.rotationRate")
	return rotationRate
}

type DeviceOrientationEvent struct {
	*Event
}

func (*DeviceOrientationEvent) InitDeviceOrientationEvent(kind string, bubbles bool, cancelable bool, alpha float32, beta float32, gamma float32, absolute bool) {
	js.Rewrite("$<.initDeviceOrientationEvent($1, $2, $3, $4, $5, $6, $7)", kind, bubbles, cancelable, alpha, beta, gamma, absolute)
}

func (*DeviceOrientationEvent) GetAbsolute() (absolute bool) {
	js.Rewrite("$<.absolute")
	return absolute
}

func (*DeviceOrientationEvent) GetAlpha() (alpha float32) {
	js.Rewrite("$<.alpha")
	return alpha
}

func (*DeviceOrientationEvent) GetBeta() (beta float32) {
	js.Rewrite("$<.beta")
	return beta
}

func (*DeviceOrientationEvent) GetGamma() (gamma float32) {
	js.Rewrite("$<.gamma")
	return gamma
}

type DeviceRotationRate struct {
}

func (*DeviceRotationRate) GetAlpha() (alpha float32) {
	js.Rewrite("$<.alpha")
	return alpha
}

func (*DeviceRotationRate) GetBeta() (beta float32) {
	js.Rewrite("$<.beta")
	return beta
}

func (*DeviceRotationRate) GetGamma() (gamma float32) {
	js.Rewrite("$<.gamma")
	return gamma
}

type Document struct {
	*Node
	*GlobalEventHandlers
	*NodeSelector
	*DocumentEvent
}

func (*Document) AdoptNode(source *Node) (n *Node) {
	js.Rewrite("$<.adoptNode($1)", source)
	return n
}

func (*Document) CaptureEvents() {
	js.Rewrite("$<.captureEvents()")
}

func (*Document) CaretRangeFromPoint(x float32, y float32) (r *Range) {
	js.Rewrite("$<.caretRangeFromPoint($1, $2)", x, y)
	return r
}

func (*Document) Clear() {
	js.Rewrite("$<.clear()")
}

func (*Document) Close() {
	js.Rewrite("$<.close()")
}

func (*Document) CreateAttribute(name string) (a *Attr) {
	js.Rewrite("$<.createAttribute($1)", name)
	return a
}

func (*Document) CreateAttributeNS(namespaceURI string, qualifiedName string) (a *Attr) {
	js.Rewrite("$<.createAttributeNS($1, $2)", namespaceURI, qualifiedName)
	return a
}

func (*Document) CreateCDATASection(data string) (c *CDATASection) {
	js.Rewrite("$<.createCDATASection($1)", data)
	return c
}

func (*Document) CreateComment(data string) (c *Comment) {
	js.Rewrite("$<.createComment($1)", data)
	return c
}

func (*Document) CreateDocumentFragment() (d *DocumentFragment) {
	js.Rewrite("$<.createDocumentFragment()")
	return d
}

func (*Document) CreateElement(tagName string) (e *Element) {
	js.Rewrite("$<.createElement($1)", tagName)
	return e
}

func (*Document) CreateElementNS(namespaceURI string, qualifiedName string) (e *Element) {
	js.Rewrite("$<.createElementNS($1, $2)", namespaceURI, qualifiedName)
	return e
}

func (*Document) CreateExpression(expression string, resolver *XPathNSResolver) (x *XPathExpression) {
	js.Rewrite("$<.createExpression($1, $2)", expression, resolver)
	return x
}

func (*Document) CreateNodeIterator(root *Node, whatToShow uint, filter *NodeFilter, entityReferenceExpansion bool) (n *NodeIterator) {
	js.Rewrite("$<.createNodeIterator($1, $2, $3, $4)", root, whatToShow, filter, entityReferenceExpansion)
	return n
}

func (*Document) CreateNSResolver(nodeResolver *Node) (x *XPathNSResolver) {
	js.Rewrite("$<.createNSResolver($1)", nodeResolver)
	return x
}

func (*Document) CreateProcessingInstruction(target string, data string) (p *ProcessingInstruction) {
	js.Rewrite("$<.createProcessingInstruction($1, $2)", target, data)
	return p
}

func (*Document) CreateRange() (r *Range) {
	js.Rewrite("$<.createRange()")
	return r
}

func (*Document) CreateTextNode(data string) (t *Text) {
	js.Rewrite("$<.createTextNode($1)", data)
	return t
}

func (*Document) CreateTouch(view *Window, target *EventTarget, identifier int, pageX int, pageY int, screenX int, screenY int) (t *Touch) {
	js.Rewrite("$<.createTouch($1, $2, $3, $4, $5, $6, $7)", view, target, identifier, pageX, pageY, screenX, screenY)
	return t
}

func (*Document) CreateTouchList(touches *Touch) (t *TouchList) {
	js.Rewrite("$<.createTouchList($1)", touches)
	return t
}

func (*Document) CreateTreeWalker(root *Node, whatToShow uint, filter *NodeFilter, entityReferenceExpansion bool) (t *TreeWalker) {
	js.Rewrite("$<.createTreeWalker($1, $2, $3, $4)", root, whatToShow, filter, entityReferenceExpansion)
	return t
}

func (*Document) ElementFromPoint(x int, y int) (e *Element) {
	js.Rewrite("$<.elementFromPoint($1, $2)", x, y)
	return e
}

func (*Document) Evaluate(expression string, contextNode *Node, resolver *XPathNSResolver, kind uint8, result *XPathResult) (x *XPathResult) {
	js.Rewrite("$<.evaluate($1, $2, $3, $4, $5)", expression, contextNode, resolver, kind, result)
	return x
}

func (*Document) ExecCommand(commandId string, showUI bool, value interface{}) (b bool) {
	js.Rewrite("$<.execCommand($1, $2, $3)", commandId, showUI, value)
	return b
}

func (*Document) ExecCommandShowHelp(commandId string) (b bool) {
	js.Rewrite("$<.execCommandShowHelp($1)", commandId)
	return b
}

func (*Document) ExitFullscreen() {
	js.Rewrite("$<.exitFullscreen()")
}

func (*Document) ExitPointerLock() {
	js.Rewrite("$<.exitPointerLock()")
}

func (*Document) Focus() {
	js.Rewrite("$<.focus()")
}

func (*Document) GetElementByID(elementId string) (e *Element) {
	js.Rewrite("$<.getElementById($1)", elementId)
	return e
}

func (*Document) GetElementsByClassName(classNames string) (n *NodeList) {
	js.Rewrite("$<.getElementsByClassName($1)", classNames)
	return n
}

func (*Document) GetElementsByName(elementName string) (n *NodeList) {
	js.Rewrite("$<.getElementsByName($1)", elementName)
	return n
}

func (*Document) GetElementsByTagName(tagname string) (n *NodeList) {
	js.Rewrite("$<.getElementsByTagName($1)", tagname)
	return n
}

func (*Document) GetElementsByTagNameNS(namespaceURI string, localName string) (n *NodeList) {
	js.Rewrite("$<.getElementsByTagNameNS($1, $2)", namespaceURI, localName)
	return n
}

func (*Document) GetSelection() (s *Selection) {
	js.Rewrite("$<.getSelection()")
	return s
}

func (*Document) HasFocus() (b bool) {
	js.Rewrite("$<.hasFocus()")
	return b
}

func (*Document) ImportNode(importedNode *Node, deep bool) (n *Node) {
	js.Rewrite("$<.importNode($1, $2)", importedNode, deep)
	return n
}

func (*Document) MsElementsFromPoint(x float32, y float32) (n *NodeList) {
	js.Rewrite("$<.msElementsFromPoint($1, $2)", x, y)
	return n
}

func (*Document) MsElementsFromRect(left float32, top float32, width float32, height float32) (n *NodeList) {
	js.Rewrite("$<.msElementsFromRect($1, $2, $3, $4)", left, top, width, height)
	return n
}

func (*Document) Open(url string, name string, features string, replace bool) (i interface{}) {
	js.Rewrite("$<.open($1, $2, $3, $4)", url, name, features, replace)
	return i
}

func (*Document) QueryCommandEnabled(commandId string) (b bool) {
	js.Rewrite("$<.queryCommandEnabled($1)", commandId)
	return b
}

func (*Document) QueryCommandIndeterm(commandId string) (b bool) {
	js.Rewrite("$<.queryCommandIndeterm($1)", commandId)
	return b
}

func (*Document) QueryCommandState(commandId string) (b bool) {
	js.Rewrite("$<.queryCommandState($1)", commandId)
	return b
}

func (*Document) QueryCommandSupported(commandId string) (b bool) {
	js.Rewrite("$<.queryCommandSupported($1)", commandId)
	return b
}

func (*Document) QueryCommandText(commandId string) (s string) {
	js.Rewrite("$<.queryCommandText($1)", commandId)
	return s
}

func (*Document) QueryCommandValue(commandId string) (s string) {
	js.Rewrite("$<.queryCommandValue($1)", commandId)
	return s
}

func (*Document) ReleaseEvents() {
	js.Rewrite("$<.releaseEvents()")
}

func (*Document) UpdateSettings() {
	js.Rewrite("$<.updateSettings()")
}

func (*Document) WebkitCancelFullScreen() {
	js.Rewrite("$<.webkitCancelFullScreen()")
}

func (*Document) WebkitExitFullscreen() {
	js.Rewrite("$<.webkitExitFullscreen()")
}

func (*Document) Write(content string) {
	js.Rewrite("$<.write($1)", content)
}

func (*Document) Writeln(content string) {
	js.Rewrite("$<.writeln($1)", content)
}

func (*Document) GetActiveElement() (activeElement *Element) {
	js.Rewrite("$<.activeElement")
	return activeElement
}

func (*Document) GetAlinkColor() (alinkColor string) {
	js.Rewrite("$<.alinkColor")
	return alinkColor
}

func (*Document) SetAlinkColor(alinkColor string) {
	js.Rewrite("$<.alinkColor = $1", alinkColor)
}

func (*Document) GetAll() (all *HTMLAllCollection) {
	js.Rewrite("$<.all")
	return all
}

func (*Document) GetAnchors() (anchors *HTMLCollection) {
	js.Rewrite("$<.anchors")
	return anchors
}

func (*Document) GetApplets() (applets *HTMLCollection) {
	js.Rewrite("$<.applets")
	return applets
}

func (*Document) GetBgColor() (bgColor string) {
	js.Rewrite("$<.bgColor")
	return bgColor
}

func (*Document) SetBgColor(bgColor string) {
	js.Rewrite("$<.bgColor = $1", bgColor)
}

func (*Document) GetBody() (body *HTMLElement) {
	js.Rewrite("$<.body")
	return body
}

func (*Document) SetBody(body *HTMLElement) {
	js.Rewrite("$<.body = $1", body)
}

func (*Document) GetCharacterSet() (characterSet string) {
	js.Rewrite("$<.characterSet")
	return characterSet
}

func (*Document) GetCharset() (charset string) {
	js.Rewrite("$<.charset")
	return charset
}

func (*Document) SetCharset(charset string) {
	js.Rewrite("$<.charset = $1", charset)
}

func (*Document) GetCompatMode() (compatMode string) {
	js.Rewrite("$<.compatMode")
	return compatMode
}

func (*Document) GetCookie() (cookie string) {
	js.Rewrite("$<.cookie")
	return cookie
}

func (*Document) SetCookie(cookie string) {
	js.Rewrite("$<.cookie = $1", cookie)
}

func (*Document) GetCurrentScript() (currentScript interface{}) {
	js.Rewrite("$<.currentScript")
	return currentScript
}

func (*Document) GetDefaultView() (defaultView *Window) {
	js.Rewrite("$<.defaultView")
	return defaultView
}

func (*Document) GetDesignMode() (designMode string) {
	js.Rewrite("$<.designMode")
	return designMode
}

func (*Document) SetDesignMode(designMode string) {
	js.Rewrite("$<.designMode = $1", designMode)
}

func (*Document) GetDir() (dir string) {
	js.Rewrite("$<.dir")
	return dir
}

func (*Document) SetDir(dir string) {
	js.Rewrite("$<.dir = $1", dir)
}

func (*Document) GetDoctype() (doctype *DocumentType) {
	js.Rewrite("$<.doctype")
	return doctype
}

func (*Document) GetDocumentElement() (documentElement *Element) {
	js.Rewrite("$<.documentElement")
	return documentElement
}

func (*Document) GetDomain() (domain string) {
	js.Rewrite("$<.domain")
	return domain
}

func (*Document) SetDomain(domain string) {
	js.Rewrite("$<.domain = $1", domain)
}

func (*Document) GetEmbeds() (embeds *HTMLCollection) {
	js.Rewrite("$<.embeds")
	return embeds
}

func (*Document) GetFgColor() (fgColor string) {
	js.Rewrite("$<.fgColor")
	return fgColor
}

func (*Document) SetFgColor(fgColor string) {
	js.Rewrite("$<.fgColor = $1", fgColor)
}

func (*Document) GetForms() (forms *HTMLCollection) {
	js.Rewrite("$<.forms")
	return forms
}

func (*Document) GetFullscreenElement() (fullscreenElement *Element) {
	js.Rewrite("$<.fullscreenElement")
	return fullscreenElement
}

func (*Document) GetFullscreenEnabled() (fullscreenEnabled bool) {
	js.Rewrite("$<.fullscreenEnabled")
	return fullscreenEnabled
}

func (*Document) GetHead() (head *HTMLHeadElement) {
	js.Rewrite("$<.head")
	return head
}

func (*Document) GetHidden() (hidden bool) {
	js.Rewrite("$<.hidden")
	return hidden
}

func (*Document) GetImages() (images *HTMLCollection) {
	js.Rewrite("$<.images")
	return images
}

func (*Document) GetImplementation() (implementation *DOMImplementation) {
	js.Rewrite("$<.implementation")
	return implementation
}

func (*Document) GetInputEncoding() (inputEncoding string) {
	js.Rewrite("$<.inputEncoding")
	return inputEncoding
}

func (*Document) GetLastModified() (lastModified string) {
	js.Rewrite("$<.lastModified")
	return lastModified
}

func (*Document) GetLinkColor() (linkColor string) {
	js.Rewrite("$<.linkColor")
	return linkColor
}

func (*Document) SetLinkColor(linkColor string) {
	js.Rewrite("$<.linkColor = $1", linkColor)
}

func (*Document) GetLinks() (links *HTMLCollection) {
	js.Rewrite("$<.links")
	return links
}

func (*Document) GetLocation() (location *Location) {
	js.Rewrite("$<.location")
	return location
}

func (*Document) GetMsCapsLockWarningOff() (msCapsLockWarningOff bool) {
	js.Rewrite("$<.msCapsLockWarningOff")
	return msCapsLockWarningOff
}

func (*Document) SetMsCapsLockWarningOff(msCapsLockWarningOff bool) {
	js.Rewrite("$<.msCapsLockWarningOff = $1", msCapsLockWarningOff)
}

func (*Document) GetMsCSSOMElementFloatMetrics() (msCSSOMElementFloatMetrics bool) {
	js.Rewrite("$<.msCSSOMElementFloatMetrics")
	return msCSSOMElementFloatMetrics
}

func (*Document) SetMsCSSOMElementFloatMetrics(msCSSOMElementFloatMetrics bool) {
	js.Rewrite("$<.msCSSOMElementFloatMetrics = $1", msCSSOMElementFloatMetrics)
}

func (*Document) GetOnabort() (e *Event) {
	js.Rewrite("$<.onabort")
	return e
}

func (*Document) SetOnabort(e *Event) {
	js.Rewrite("$<.onabort = $1", e)
}

func (*Document) GetOnactivate() (e *Event) {
	js.Rewrite("$<.onactivate")
	return e
}

func (*Document) SetOnactivate(e *Event) {
	js.Rewrite("$<.onactivate = $1", e)
}

func (*Document) GetOnbeforeactivate() (e *Event) {
	js.Rewrite("$<.onbeforeactivate")
	return e
}

func (*Document) SetOnbeforeactivate(e *Event) {
	js.Rewrite("$<.onbeforeactivate = $1", e)
}

func (*Document) GetOnbeforedeactivate() (e *Event) {
	js.Rewrite("$<.onbeforedeactivate")
	return e
}

func (*Document) SetOnbeforedeactivate(e *Event) {
	js.Rewrite("$<.onbeforedeactivate = $1", e)
}

func (*Document) GetOnblur() (e *Event) {
	js.Rewrite("$<.onblur")
	return e
}

func (*Document) SetOnblur(e *Event) {
	js.Rewrite("$<.onblur = $1", e)
}

func (*Document) GetOncanplay() (e *Event) {
	js.Rewrite("$<.oncanplay")
	return e
}

func (*Document) SetOncanplay(e *Event) {
	js.Rewrite("$<.oncanplay = $1", e)
}

func (*Document) GetOncanplaythrough() (e *Event) {
	js.Rewrite("$<.oncanplaythrough")
	return e
}

func (*Document) SetOncanplaythrough(e *Event) {
	js.Rewrite("$<.oncanplaythrough = $1", e)
}

func (*Document) GetOnchange() (e *Event) {
	js.Rewrite("$<.onchange")
	return e
}

func (*Document) SetOnchange(e *Event) {
	js.Rewrite("$<.onchange = $1", e)
}

func (*Document) GetOnclick() (e *Event) {
	js.Rewrite("$<.onclick")
	return e
}

func (*Document) SetOnclick(e *Event) {
	js.Rewrite("$<.onclick = $1", e)
}

func (*Document) GetOncontextmenu() (e *Event) {
	js.Rewrite("$<.oncontextmenu")
	return e
}

func (*Document) SetOncontextmenu(e *Event) {
	js.Rewrite("$<.oncontextmenu = $1", e)
}

func (*Document) GetOndblclick() (e *Event) {
	js.Rewrite("$<.ondblclick")
	return e
}

func (*Document) SetOndblclick(e *Event) {
	js.Rewrite("$<.ondblclick = $1", e)
}

func (*Document) GetOndeactivate() (e *Event) {
	js.Rewrite("$<.ondeactivate")
	return e
}

func (*Document) SetOndeactivate(e *Event) {
	js.Rewrite("$<.ondeactivate = $1", e)
}

func (*Document) GetOndrag() (e *Event) {
	js.Rewrite("$<.ondrag")
	return e
}

func (*Document) SetOndrag(e *Event) {
	js.Rewrite("$<.ondrag = $1", e)
}

func (*Document) GetOndragend() (e *Event) {
	js.Rewrite("$<.ondragend")
	return e
}

func (*Document) SetOndragend(e *Event) {
	js.Rewrite("$<.ondragend = $1", e)
}

func (*Document) GetOndragenter() (e *Event) {
	js.Rewrite("$<.ondragenter")
	return e
}

func (*Document) SetOndragenter(e *Event) {
	js.Rewrite("$<.ondragenter = $1", e)
}

func (*Document) GetOndragleave() (e *Event) {
	js.Rewrite("$<.ondragleave")
	return e
}

func (*Document) SetOndragleave(e *Event) {
	js.Rewrite("$<.ondragleave = $1", e)
}

func (*Document) GetOndragover() (e *Event) {
	js.Rewrite("$<.ondragover")
	return e
}

func (*Document) SetOndragover(e *Event) {
	js.Rewrite("$<.ondragover = $1", e)
}

func (*Document) GetOndragstart() (e *Event) {
	js.Rewrite("$<.ondragstart")
	return e
}

func (*Document) SetOndragstart(e *Event) {
	js.Rewrite("$<.ondragstart = $1", e)
}

func (*Document) GetOndrop() (e *Event) {
	js.Rewrite("$<.ondrop")
	return e
}

func (*Document) SetOndrop(e *Event) {
	js.Rewrite("$<.ondrop = $1", e)
}

func (*Document) GetOndurationchange() (e *Event) {
	js.Rewrite("$<.ondurationchange")
	return e
}

func (*Document) SetOndurationchange(e *Event) {
	js.Rewrite("$<.ondurationchange = $1", e)
}

func (*Document) GetOnemptied() (e *Event) {
	js.Rewrite("$<.onemptied")
	return e
}

func (*Document) SetOnemptied(e *Event) {
	js.Rewrite("$<.onemptied = $1", e)
}

func (*Document) GetOnended() (e *Event) {
	js.Rewrite("$<.onended")
	return e
}

func (*Document) SetOnended(e *Event) {
	js.Rewrite("$<.onended = $1", e)
}

func (*Document) GetOnerror() (e *Event) {
	js.Rewrite("$<.onerror")
	return e
}

func (*Document) SetOnerror(e *Event) {
	js.Rewrite("$<.onerror = $1", e)
}

func (*Document) GetOnfocus() (e *Event) {
	js.Rewrite("$<.onfocus")
	return e
}

func (*Document) SetOnfocus(e *Event) {
	js.Rewrite("$<.onfocus = $1", e)
}

func (*Document) GetOnfullscreenchange() (fullscreenchange *Event) {
	js.Rewrite("$<.onfullscreenchange")
	return fullscreenchange
}

func (*Document) SetOnfullscreenchange(fullscreenchange *Event) {
	js.Rewrite("$<.onfullscreenchange = $1", fullscreenchange)
}

func (*Document) GetOnfullscreenerror() (fullscreenerror *Event) {
	js.Rewrite("$<.onfullscreenerror")
	return fullscreenerror
}

func (*Document) SetOnfullscreenerror(fullscreenerror *Event) {
	js.Rewrite("$<.onfullscreenerror = $1", fullscreenerror)
}

func (*Document) GetOninput() (e *Event) {
	js.Rewrite("$<.oninput")
	return e
}

func (*Document) SetOninput(e *Event) {
	js.Rewrite("$<.oninput = $1", e)
}

func (*Document) GetOninvalid() (e *Event) {
	js.Rewrite("$<.oninvalid")
	return e
}

func (*Document) SetOninvalid(e *Event) {
	js.Rewrite("$<.oninvalid = $1", e)
}

func (*Document) GetOnkeydown() (e *Event) {
	js.Rewrite("$<.onkeydown")
	return e
}

func (*Document) SetOnkeydown(e *Event) {
	js.Rewrite("$<.onkeydown = $1", e)
}

func (*Document) GetOnkeypress() (e *Event) {
	js.Rewrite("$<.onkeypress")
	return e
}

func (*Document) SetOnkeypress(e *Event) {
	js.Rewrite("$<.onkeypress = $1", e)
}

func (*Document) GetOnkeyup() (e *Event) {
	js.Rewrite("$<.onkeyup")
	return e
}

func (*Document) SetOnkeyup(e *Event) {
	js.Rewrite("$<.onkeyup = $1", e)
}

func (*Document) GetOnload() (e *Event) {
	js.Rewrite("$<.onload")
	return e
}

func (*Document) SetOnload(e *Event) {
	js.Rewrite("$<.onload = $1", e)
}

func (*Document) GetOnloadeddata() (e *Event) {
	js.Rewrite("$<.onloadeddata")
	return e
}

func (*Document) SetOnloadeddata(e *Event) {
	js.Rewrite("$<.onloadeddata = $1", e)
}

func (*Document) GetOnloadedmetadata() (e *Event) {
	js.Rewrite("$<.onloadedmetadata")
	return e
}

func (*Document) SetOnloadedmetadata(e *Event) {
	js.Rewrite("$<.onloadedmetadata = $1", e)
}

func (*Document) GetOnloadstart() (e *Event) {
	js.Rewrite("$<.onloadstart")
	return e
}

func (*Document) SetOnloadstart(e *Event) {
	js.Rewrite("$<.onloadstart = $1", e)
}

func (*Document) GetOnmousedown() (e *Event) {
	js.Rewrite("$<.onmousedown")
	return e
}

func (*Document) SetOnmousedown(e *Event) {
	js.Rewrite("$<.onmousedown = $1", e)
}

func (*Document) GetOnmousemove() (e *Event) {
	js.Rewrite("$<.onmousemove")
	return e
}

func (*Document) SetOnmousemove(e *Event) {
	js.Rewrite("$<.onmousemove = $1", e)
}

func (*Document) GetOnmouseout() (e *Event) {
	js.Rewrite("$<.onmouseout")
	return e
}

func (*Document) SetOnmouseout(e *Event) {
	js.Rewrite("$<.onmouseout = $1", e)
}

func (*Document) GetOnmouseover() (e *Event) {
	js.Rewrite("$<.onmouseover")
	return e
}

func (*Document) SetOnmouseover(e *Event) {
	js.Rewrite("$<.onmouseover = $1", e)
}

func (*Document) GetOnmouseup() (e *Event) {
	js.Rewrite("$<.onmouseup")
	return e
}

func (*Document) SetOnmouseup(e *Event) {
	js.Rewrite("$<.onmouseup = $1", e)
}

func (*Document) GetOnmousewheel() (e *Event) {
	js.Rewrite("$<.onmousewheel")
	return e
}

func (*Document) SetOnmousewheel(e *Event) {
	js.Rewrite("$<.onmousewheel = $1", e)
}

func (*Document) GetOnmscontentzoom() (MSContentZoom *UIEvent) {
	js.Rewrite("$<.onmscontentzoom")
	return MSContentZoom
}

func (*Document) SetOnmscontentzoom(MSContentZoom *UIEvent) {
	js.Rewrite("$<.onmscontentzoom = $1", MSContentZoom)
}

func (*Document) GetOnmsgesturechange() (e *Event) {
	js.Rewrite("$<.onmsgesturechange")
	return e
}

func (*Document) SetOnmsgesturechange(e *Event) {
	js.Rewrite("$<.onmsgesturechange = $1", e)
}

func (*Document) GetOnmsgesturedoubletap() (e *Event) {
	js.Rewrite("$<.onmsgesturedoubletap")
	return e
}

func (*Document) SetOnmsgesturedoubletap(e *Event) {
	js.Rewrite("$<.onmsgesturedoubletap = $1", e)
}

func (*Document) GetOnmsgestureend() (e *Event) {
	js.Rewrite("$<.onmsgestureend")
	return e
}

func (*Document) SetOnmsgestureend(e *Event) {
	js.Rewrite("$<.onmsgestureend = $1", e)
}

func (*Document) GetOnmsgesturehold() (e *Event) {
	js.Rewrite("$<.onmsgesturehold")
	return e
}

func (*Document) SetOnmsgesturehold(e *Event) {
	js.Rewrite("$<.onmsgesturehold = $1", e)
}

func (*Document) GetOnmsgesturestart() (e *Event) {
	js.Rewrite("$<.onmsgesturestart")
	return e
}

func (*Document) SetOnmsgesturestart(e *Event) {
	js.Rewrite("$<.onmsgesturestart = $1", e)
}

func (*Document) GetOnmsgesturetap() (e *Event) {
	js.Rewrite("$<.onmsgesturetap")
	return e
}

func (*Document) SetOnmsgesturetap(e *Event) {
	js.Rewrite("$<.onmsgesturetap = $1", e)
}

func (*Document) GetOnmsinertiastart() (e *Event) {
	js.Rewrite("$<.onmsinertiastart")
	return e
}

func (*Document) SetOnmsinertiastart(e *Event) {
	js.Rewrite("$<.onmsinertiastart = $1", e)
}

func (*Document) GetOnmsmanipulationstatechanged() (MSManipulationStateChanged *MSManipulationEvent) {
	js.Rewrite("$<.onmsmanipulationstatechanged")
	return MSManipulationStateChanged
}

func (*Document) SetOnmsmanipulationstatechanged(MSManipulationStateChanged *MSManipulationEvent) {
	js.Rewrite("$<.onmsmanipulationstatechanged = $1", MSManipulationStateChanged)
}

func (*Document) GetOnmspointercancel() (e *Event) {
	js.Rewrite("$<.onmspointercancel")
	return e
}

func (*Document) SetOnmspointercancel(e *Event) {
	js.Rewrite("$<.onmspointercancel = $1", e)
}

func (*Document) GetOnmspointerdown() (e *Event) {
	js.Rewrite("$<.onmspointerdown")
	return e
}

func (*Document) SetOnmspointerdown(e *Event) {
	js.Rewrite("$<.onmspointerdown = $1", e)
}

func (*Document) GetOnmspointerenter() (e *Event) {
	js.Rewrite("$<.onmspointerenter")
	return e
}

func (*Document) SetOnmspointerenter(e *Event) {
	js.Rewrite("$<.onmspointerenter = $1", e)
}

func (*Document) GetOnmspointerleave() (e *Event) {
	js.Rewrite("$<.onmspointerleave")
	return e
}

func (*Document) SetOnmspointerleave(e *Event) {
	js.Rewrite("$<.onmspointerleave = $1", e)
}

func (*Document) GetOnmspointermove() (e *Event) {
	js.Rewrite("$<.onmspointermove")
	return e
}

func (*Document) SetOnmspointermove(e *Event) {
	js.Rewrite("$<.onmspointermove = $1", e)
}

func (*Document) GetOnmspointerout() (e *Event) {
	js.Rewrite("$<.onmspointerout")
	return e
}

func (*Document) SetOnmspointerout(e *Event) {
	js.Rewrite("$<.onmspointerout = $1", e)
}

func (*Document) GetOnmspointerover() (e *Event) {
	js.Rewrite("$<.onmspointerover")
	return e
}

func (*Document) SetOnmspointerover(e *Event) {
	js.Rewrite("$<.onmspointerover = $1", e)
}

func (*Document) GetOnmspointerup() (e *Event) {
	js.Rewrite("$<.onmspointerup")
	return e
}

func (*Document) SetOnmspointerup(e *Event) {
	js.Rewrite("$<.onmspointerup = $1", e)
}

func (*Document) GetOnmssitemodejumplistitemremoved() (mssitemodejumplistitemremoved *MSSiteModeEvent) {
	js.Rewrite("$<.onmssitemodejumplistitemremoved")
	return mssitemodejumplistitemremoved
}

func (*Document) SetOnmssitemodejumplistitemremoved(mssitemodejumplistitemremoved *MSSiteModeEvent) {
	js.Rewrite("$<.onmssitemodejumplistitemremoved = $1", mssitemodejumplistitemremoved)
}

func (*Document) GetOnmsthumbnailclick() (msthumbnailclick *MSSiteModeEvent) {
	js.Rewrite("$<.onmsthumbnailclick")
	return msthumbnailclick
}

func (*Document) SetOnmsthumbnailclick(msthumbnailclick *MSSiteModeEvent) {
	js.Rewrite("$<.onmsthumbnailclick = $1", msthumbnailclick)
}

func (*Document) GetOnpause() (e *Event) {
	js.Rewrite("$<.onpause")
	return e
}

func (*Document) SetOnpause(e *Event) {
	js.Rewrite("$<.onpause = $1", e)
}

func (*Document) GetOnplay() (e *Event) {
	js.Rewrite("$<.onplay")
	return e
}

func (*Document) SetOnplay(e *Event) {
	js.Rewrite("$<.onplay = $1", e)
}

func (*Document) GetOnplaying() (e *Event) {
	js.Rewrite("$<.onplaying")
	return e
}

func (*Document) SetOnplaying(e *Event) {
	js.Rewrite("$<.onplaying = $1", e)
}

func (*Document) GetOnpointerlockchange() (e *Event) {
	js.Rewrite("$<.onpointerlockchange")
	return e
}

func (*Document) SetOnpointerlockchange(e *Event) {
	js.Rewrite("$<.onpointerlockchange = $1", e)
}

func (*Document) GetOnpointerlockerror() (e *Event) {
	js.Rewrite("$<.onpointerlockerror")
	return e
}

func (*Document) SetOnpointerlockerror(e *Event) {
	js.Rewrite("$<.onpointerlockerror = $1", e)
}

func (*Document) GetOnprogress() (e *Event) {
	js.Rewrite("$<.onprogress")
	return e
}

func (*Document) SetOnprogress(e *Event) {
	js.Rewrite("$<.onprogress = $1", e)
}

func (*Document) GetOnratechange() (e *Event) {
	js.Rewrite("$<.onratechange")
	return e
}

func (*Document) SetOnratechange(e *Event) {
	js.Rewrite("$<.onratechange = $1", e)
}

func (*Document) GetOnreadystatechange() (readystatechange *Event) {
	js.Rewrite("$<.onreadystatechange")
	return readystatechange
}

func (*Document) SetOnreadystatechange(readystatechange *Event) {
	js.Rewrite("$<.onreadystatechange = $1", readystatechange)
}

func (*Document) GetOnreset() (e *Event) {
	js.Rewrite("$<.onreset")
	return e
}

func (*Document) SetOnreset(e *Event) {
	js.Rewrite("$<.onreset = $1", e)
}

func (*Document) GetOnscroll() (e *Event) {
	js.Rewrite("$<.onscroll")
	return e
}

func (*Document) SetOnscroll(e *Event) {
	js.Rewrite("$<.onscroll = $1", e)
}

func (*Document) GetOnseeked() (e *Event) {
	js.Rewrite("$<.onseeked")
	return e
}

func (*Document) SetOnseeked(e *Event) {
	js.Rewrite("$<.onseeked = $1", e)
}

func (*Document) GetOnseeking() (e *Event) {
	js.Rewrite("$<.onseeking")
	return e
}

func (*Document) SetOnseeking(e *Event) {
	js.Rewrite("$<.onseeking = $1", e)
}

func (*Document) GetOnselect() (e *Event) {
	js.Rewrite("$<.onselect")
	return e
}

func (*Document) SetOnselect(e *Event) {
	js.Rewrite("$<.onselect = $1", e)
}

func (*Document) GetOnselectionchange() (selectionchange *Event) {
	js.Rewrite("$<.onselectionchange")
	return selectionchange
}

func (*Document) SetOnselectionchange(selectionchange *Event) {
	js.Rewrite("$<.onselectionchange = $1", selectionchange)
}

func (*Document) GetOnselectstart() (e *Event) {
	js.Rewrite("$<.onselectstart")
	return e
}

func (*Document) SetOnselectstart(e *Event) {
	js.Rewrite("$<.onselectstart = $1", e)
}

func (*Document) GetOnstalled() (e *Event) {
	js.Rewrite("$<.onstalled")
	return e
}

func (*Document) SetOnstalled(e *Event) {
	js.Rewrite("$<.onstalled = $1", e)
}

func (*Document) GetOnstop() (stop *Event) {
	js.Rewrite("$<.onstop")
	return stop
}

func (*Document) SetOnstop(stop *Event) {
	js.Rewrite("$<.onstop = $1", stop)
}

func (*Document) GetOnsubmit() (e *Event) {
	js.Rewrite("$<.onsubmit")
	return e
}

func (*Document) SetOnsubmit(e *Event) {
	js.Rewrite("$<.onsubmit = $1", e)
}

func (*Document) GetOnsuspend() (e *Event) {
	js.Rewrite("$<.onsuspend")
	return e
}

func (*Document) SetOnsuspend(e *Event) {
	js.Rewrite("$<.onsuspend = $1", e)
}

func (*Document) GetOntimeupdate() (e *Event) {
	js.Rewrite("$<.ontimeupdate")
	return e
}

func (*Document) SetOntimeupdate(e *Event) {
	js.Rewrite("$<.ontimeupdate = $1", e)
}

func (*Document) GetOntouchcancel() (e *Event) {
	js.Rewrite("$<.ontouchcancel")
	return e
}

func (*Document) SetOntouchcancel(e *Event) {
	js.Rewrite("$<.ontouchcancel = $1", e)
}

func (*Document) GetOntouchend() (e *Event) {
	js.Rewrite("$<.ontouchend")
	return e
}

func (*Document) SetOntouchend(e *Event) {
	js.Rewrite("$<.ontouchend = $1", e)
}

func (*Document) GetOntouchmove() (e *Event) {
	js.Rewrite("$<.ontouchmove")
	return e
}

func (*Document) SetOntouchmove(e *Event) {
	js.Rewrite("$<.ontouchmove = $1", e)
}

func (*Document) GetOntouchstart() (e *Event) {
	js.Rewrite("$<.ontouchstart")
	return e
}

func (*Document) SetOntouchstart(e *Event) {
	js.Rewrite("$<.ontouchstart = $1", e)
}

func (*Document) GetOnvolumechange() (e *Event) {
	js.Rewrite("$<.onvolumechange")
	return e
}

func (*Document) SetOnvolumechange(e *Event) {
	js.Rewrite("$<.onvolumechange = $1", e)
}

func (*Document) GetOnwaiting() (e *Event) {
	js.Rewrite("$<.onwaiting")
	return e
}

func (*Document) SetOnwaiting(e *Event) {
	js.Rewrite("$<.onwaiting = $1", e)
}

func (*Document) GetOnwebkitfullscreenchange() (e *Event) {
	js.Rewrite("$<.onwebkitfullscreenchange")
	return e
}

func (*Document) SetOnwebkitfullscreenchange(e *Event) {
	js.Rewrite("$<.onwebkitfullscreenchange = $1", e)
}

func (*Document) GetOnwebkitfullscreenerror() (e *Event) {
	js.Rewrite("$<.onwebkitfullscreenerror")
	return e
}

func (*Document) SetOnwebkitfullscreenerror(e *Event) {
	js.Rewrite("$<.onwebkitfullscreenerror = $1", e)
}

func (*Document) GetPlugins() (plugins *HTMLCollection) {
	js.Rewrite("$<.plugins")
	return plugins
}

func (*Document) GetPointerLockElement() (pointerLockElement *Element) {
	js.Rewrite("$<.pointerLockElement")
	return pointerLockElement
}

func (*Document) GetReadyState() (readyState string) {
	js.Rewrite("$<.readyState")
	return readyState
}

func (*Document) GetReferrer() (referrer string) {
	js.Rewrite("$<.referrer")
	return referrer
}

func (*Document) GetRootElement() (rootElement *SVGSVGElement) {
	js.Rewrite("$<.rootElement")
	return rootElement
}

func (*Document) GetScripts() (scripts *HTMLCollection) {
	js.Rewrite("$<.scripts")
	return scripts
}

func (*Document) GetScrollingElement() (scrollingElement *Element) {
	js.Rewrite("$<.scrollingElement")
	return scrollingElement
}

func (*Document) GetStyleSheets() (styleSheets *StyleSheetList) {
	js.Rewrite("$<.styleSheets")
	return styleSheets
}

func (*Document) GetTitle() (title string) {
	js.Rewrite("$<.title")
	return title
}

func (*Document) SetTitle(title string) {
	js.Rewrite("$<.title = $1", title)
}

func (*Document) GetURL() (URL string) {
	js.Rewrite("$<.URL")
	return URL
}

func (*Document) GetURLUnencoded() (URLUnencoded string) {
	js.Rewrite("$<.URLUnencoded")
	return URLUnencoded
}

func (*Document) GetVisibilityState() (visibilityState *VisibilityState) {
	js.Rewrite("$<.visibilityState")
	return visibilityState
}

func (*Document) GetVlinkColor() (vlinkColor string) {
	js.Rewrite("$<.vlinkColor")
	return vlinkColor
}

func (*Document) SetVlinkColor(vlinkColor string) {
	js.Rewrite("$<.vlinkColor = $1", vlinkColor)
}

func (*Document) GetWebkitCurrentFullScreenElement() (webkitCurrentFullScreenElement *Element) {
	js.Rewrite("$<.webkitCurrentFullScreenElement")
	return webkitCurrentFullScreenElement
}

func (*Document) GetWebkitFullscreenElement() (webkitFullscreenElement *Element) {
	js.Rewrite("$<.webkitFullscreenElement")
	return webkitFullscreenElement
}

func (*Document) GetWebkitFullscreenEnabled() (webkitFullscreenEnabled bool) {
	js.Rewrite("$<.webkitFullscreenEnabled")
	return webkitFullscreenEnabled
}

func (*Document) GetWebkitIsFullScreen() (webkitIsFullScreen bool) {
	js.Rewrite("$<.webkitIsFullScreen")
	return webkitIsFullScreen
}

func (*Document) GetXMLEncoding() (xmlEncoding string) {
	js.Rewrite("$<.xmlEncoding")
	return xmlEncoding
}

func (*Document) GetXMLStandalone() (xmlStandalone bool) {
	js.Rewrite("$<.xmlStandalone")
	return xmlStandalone
}

func (*Document) SetXMLStandalone(xmlStandalone bool) {
	js.Rewrite("$<.xmlStandalone = $1", xmlStandalone)
}

func (*Document) GetXMLVersion() (xmlVersion string) {
	js.Rewrite("$<.xmlVersion")
	return xmlVersion
}

func (*Document) SetXMLVersion(xmlVersion string) {
	js.Rewrite("$<.xmlVersion = $1", xmlVersion)
}

type DocumentFragment struct {
	*Node
	*NodeSelector
}

type DocumentType struct {
	*Node
	*ChildNode
}

func (*DocumentType) GetEntities() (entities *NamedNodeMap) {
	js.Rewrite("$<.entities")
	return entities
}

func (*DocumentType) GetInternalSubset() (internalSubset string) {
	js.Rewrite("$<.internalSubset")
	return internalSubset
}

func (*DocumentType) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*DocumentType) GetNotations() (notations *NamedNodeMap) {
	js.Rewrite("$<.notations")
	return notations
}

func (*DocumentType) GetPublicID() (publicId string) {
	js.Rewrite("$<.publicId")
	return publicId
}

func (*DocumentType) GetSystemID() (systemId string) {
	js.Rewrite("$<.systemId")
	return systemId
}

type DOMError struct {
}

func (*DOMError) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

func (*DOMError) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

type DOMException struct {
}

func (*DOMException) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

func (*DOMException) GetCode() (code uint8) {
	js.Rewrite("$<.code")
	return code
}

func (*DOMException) GetMessage() (message string) {
	js.Rewrite("$<.message")
	return message
}

func (*DOMException) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

type DOMImplementation struct {
}

func (*DOMImplementation) CreateDocument(namespaceURI string, qualifiedName string, doctype *DocumentType) (d *Document) {
	js.Rewrite("$<.createDocument($1, $2, $3)", namespaceURI, qualifiedName, doctype)
	return d
}

func (*DOMImplementation) CreateDocumentType(qualifiedName string, publicId string, systemId string) (d *DocumentType) {
	js.Rewrite("$<.createDocumentType($1, $2, $3)", qualifiedName, publicId, systemId)
	return d
}

func (*DOMImplementation) CreateHTMLDocument(title string) (d *Document) {
	js.Rewrite("$<.createHTMLDocument($1)", title)
	return d
}

func (*DOMImplementation) HasFeature() (b bool) {
	js.Rewrite("$<.hasFeature()")
	return b
}

type DOMParser struct {
}

func (*DOMParser) ParseFromString(source string, mimeType string) (d *Document) {
	js.Rewrite("$<.parseFromString($1, $2)", source, mimeType)
	return d
}

type DOMSettableTokenList struct {
	*DOMTokenList
}

func (*DOMSettableTokenList) GetValue() (value string) {
	js.Rewrite("$<.value")
	return value
}

func (*DOMSettableTokenList) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}

type DOMStringList struct {
}

func (*DOMStringList) Contains(str string) (b bool) {
	js.Rewrite("$<.contains($1)", str)
	return b
}

func (*DOMStringList) Item(index uint) (s string) {
	js.Rewrite("$<.item($1)", index)
	return s
}

func (*DOMStringList) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

type DOMStringMap struct {
}

type DOMTokenList struct {
}

func (*DOMTokenList) Add(token string) {
	js.Rewrite("$<.add($1)", token)
}

func (*DOMTokenList) Contains(token string) (b bool) {
	js.Rewrite("$<.contains($1)", token)
	return b
}

func (*DOMTokenList) Item(index uint) (s string) {
	js.Rewrite("$<.item($1)", index)
	return s
}

func (*DOMTokenList) Remove(token string) {
	js.Rewrite("$<.remove($1)", token)
}

func (*DOMTokenList) Toggle(token string, force bool) (b bool) {
	js.Rewrite("$<.toggle($1, $2)", token, force)
	return b
}

func (*DOMTokenList) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

func (*DOMTokenList) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

type DragEvent struct {
	*MouseEvent
}

func (*DragEvent) InitDragEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, screenXArg int, screenYArg int, clientXArg int, clientYArg int, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg uint8, relatedTargetArg *EventTarget, dataTransferArg *DataTransfer) {
	js.Rewrite("$<.initDragEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, ctrlKeyArg, altKeyArg, shiftKeyArg, metaKeyArg, buttonArg, relatedTargetArg, dataTransferArg)
}

func (*DragEvent) MsConvertURL(file *File, targetType string, targetURL string) {
	js.Rewrite("$<.msConvertURL($1, $2, $3)", file, targetType, targetURL)
}

func (*DragEvent) GetDataTransfer() (dataTransfer *DataTransfer) {
	js.Rewrite("$<.dataTransfer")
	return dataTransfer
}

type DynamicsCompressorNode struct {
	*AudioNode
}

func (*DynamicsCompressorNode) GetAttack() (attack *AudioParam) {
	js.Rewrite("$<.attack")
	return attack
}

func (*DynamicsCompressorNode) GetKnee() (knee *AudioParam) {
	js.Rewrite("$<.knee")
	return knee
}

func (*DynamicsCompressorNode) GetRatio() (ratio *AudioParam) {
	js.Rewrite("$<.ratio")
	return ratio
}

func (*DynamicsCompressorNode) GetReduction() (reduction float32) {
	js.Rewrite("$<.reduction")
	return reduction
}

func (*DynamicsCompressorNode) GetRelease() (release *AudioParam) {
	js.Rewrite("$<.release")
	return release
}

func (*DynamicsCompressorNode) GetThreshold() (threshold *AudioParam) {
	js.Rewrite("$<.threshold")
	return threshold
}

type Element struct {
	*Node
	*GlobalEventHandlers
	*ElementTraversal
	*NodeSelector
	*ChildNode
}

func (*Element) GetAttribute(qualifiedName string) (s string) {
	js.Rewrite("$<.getAttribute($1)", qualifiedName)
	return s
}

func (*Element) GetAttributeNode(name string) (a *Attr) {
	js.Rewrite("$<.getAttributeNode($1)", name)
	return a
}

func (*Element) GetAttributeNodeNS(namespaceURI string, localName string) (a *Attr) {
	js.Rewrite("$<.getAttributeNodeNS($1, $2)", namespaceURI, localName)
	return a
}

func (*Element) GetAttributeNS(namespaceURI string, localName string) (s string) {
	js.Rewrite("$<.getAttributeNS($1, $2)", namespaceURI, localName)
	return s
}

func (*Element) GetBoundingClientRect() (c *ClientRect) {
	js.Rewrite("$<.getBoundingClientRect()")
	return c
}

func (*Element) GetClientRects() (c *ClientRectList) {
	js.Rewrite("$<.getClientRects()")
	return c
}

func (*Element) GetElementsByTagName(name string) (n *NodeList) {
	js.Rewrite("$<.getElementsByTagName($1)", name)
	return n
}

func (*Element) GetElementsByTagNameNS(namespaceURI string, localName string) (n *NodeList) {
	js.Rewrite("$<.getElementsByTagNameNS($1, $2)", namespaceURI, localName)
	return n
}

func (*Element) HasAttribute(name string) (b bool) {
	js.Rewrite("$<.hasAttribute($1)", name)
	return b
}

func (*Element) HasAttributeNS(namespaceURI string, localName string) (b bool) {
	js.Rewrite("$<.hasAttributeNS($1, $2)", namespaceURI, localName)
	return b
}

func (*Element) MsGetRegionContent() (m *MSRangeCollection) {
	js.Rewrite("$<.msGetRegionContent()")
	return m
}

func (*Element) MsGetUntransformedBounds() (c *ClientRect) {
	js.Rewrite("$<.msGetUntransformedBounds()")
	return c
}

func (*Element) MsMatchesSelector(selectors string) (b bool) {
	js.Rewrite("$<.msMatchesSelector($1)", selectors)
	return b
}

func (*Element) MsReleasePointerCapture(pointerId int) {
	js.Rewrite("$<.msReleasePointerCapture($1)", pointerId)
}

func (*Element) MsSetPointerCapture(pointerId int) {
	js.Rewrite("$<.msSetPointerCapture($1)", pointerId)
}

func (*Element) MsZoomTo(args *MsZoomToOptions) {
	js.Rewrite("$<.msZoomTo($1)", args)
}

func (*Element) ReleasePointerCapture(pointerId int) {
	js.Rewrite("$<.releasePointerCapture($1)", pointerId)
}

func (*Element) RemoveAttribute(qualifiedName string) {
	js.Rewrite("$<.removeAttribute($1)", qualifiedName)
}

func (*Element) RemoveAttributeNode(oldAttr *Attr) (a *Attr) {
	js.Rewrite("$<.removeAttributeNode($1)", oldAttr)
	return a
}

func (*Element) RemoveAttributeNS(namespaceURI string, localName string) {
	js.Rewrite("$<.removeAttributeNS($1, $2)", namespaceURI, localName)
}

func (*Element) RequestFullscreen() {
	js.Rewrite("$<.requestFullscreen()")
}

func (*Element) RequestPointerLock() {
	js.Rewrite("$<.requestPointerLock()")
}

func (*Element) SetAttribute(qualifiedName string, value string) {
	js.Rewrite("$<.setAttribute($1, $2)", qualifiedName, value)
}

func (*Element) SetAttributeNode(newAttr *Attr) (a *Attr) {
	js.Rewrite("$<.setAttributeNode($1)", newAttr)
	return a
}

func (*Element) SetAttributeNodeNS(newAttr *Attr) (a *Attr) {
	js.Rewrite("$<.setAttributeNodeNS($1)", newAttr)
	return a
}

func (*Element) SetAttributeNS(namespaceURI string, qualifiedName string, value string) {
	js.Rewrite("$<.setAttributeNS($1, $2, $3)", namespaceURI, qualifiedName, value)
}

func (*Element) SetPointerCapture(pointerId int) {
	js.Rewrite("$<.setPointerCapture($1)", pointerId)
}

func (*Element) WebkitMatchesSelector(selectors string) (b bool) {
	js.Rewrite("$<.webkitMatchesSelector($1)", selectors)
	return b
}

func (*Element) WebkitRequestFullscreen() {
	js.Rewrite("$<.webkitRequestFullscreen()")
}

func (*Element) WebkitRequestFullScreen() {
	js.Rewrite("$<.webkitRequestFullScreen()")
}

func (*Element) GetClassList() (classList *DOMTokenList) {
	js.Rewrite("$<.classList")
	return classList
}

func (*Element) GetClassName() (className string) {
	js.Rewrite("$<.className")
	return className
}

func (*Element) SetClassName(className string) {
	js.Rewrite("$<.className = $1", className)
}

func (*Element) GetClientHeight() (clientHeight int) {
	js.Rewrite("$<.clientHeight")
	return clientHeight
}

func (*Element) GetClientLeft() (clientLeft int) {
	js.Rewrite("$<.clientLeft")
	return clientLeft
}

func (*Element) GetClientTop() (clientTop int) {
	js.Rewrite("$<.clientTop")
	return clientTop
}

func (*Element) GetClientWidth() (clientWidth int) {
	js.Rewrite("$<.clientWidth")
	return clientWidth
}

func (*Element) GetID() (id string) {
	js.Rewrite("$<.id")
	return id
}

func (*Element) SetID(id string) {
	js.Rewrite("$<.id = $1", id)
}

func (*Element) GetInnerHTML() (innerHTML string) {
	js.Rewrite("$<.innerHTML")
	return innerHTML
}

func (*Element) SetInnerHTML(innerHTML string) {
	js.Rewrite("$<.innerHTML = $1", innerHTML)
}

func (*Element) GetMsContentZoomFactor() (msContentZoomFactor float32) {
	js.Rewrite("$<.msContentZoomFactor")
	return msContentZoomFactor
}

func (*Element) SetMsContentZoomFactor(msContentZoomFactor float32) {
	js.Rewrite("$<.msContentZoomFactor = $1", msContentZoomFactor)
}

func (*Element) GetMsRegionOverflow() (msRegionOverflow string) {
	js.Rewrite("$<.msRegionOverflow")
	return msRegionOverflow
}

func (*Element) GetOnariarequest() (e *Event) {
	js.Rewrite("$<.onariarequest")
	return e
}

func (*Element) SetOnariarequest(e *Event) {
	js.Rewrite("$<.onariarequest = $1", e)
}

func (*Element) GetOncommand() (e *Event) {
	js.Rewrite("$<.oncommand")
	return e
}

func (*Element) SetOncommand(e *Event) {
	js.Rewrite("$<.oncommand = $1", e)
}

func (*Element) GetOngotpointercapture() (gotpointercapture *PointerEvent) {
	js.Rewrite("$<.ongotpointercapture")
	return gotpointercapture
}

func (*Element) SetOngotpointercapture(gotpointercapture *PointerEvent) {
	js.Rewrite("$<.ongotpointercapture = $1", gotpointercapture)
}

func (*Element) GetOnlostpointercapture() (lostpointercapture *PointerEvent) {
	js.Rewrite("$<.onlostpointercapture")
	return lostpointercapture
}

func (*Element) SetOnlostpointercapture(lostpointercapture *PointerEvent) {
	js.Rewrite("$<.onlostpointercapture = $1", lostpointercapture)
}

func (*Element) GetOnmsgesturechange() (MSGestureChange *MSGestureEvent) {
	js.Rewrite("$<.onmsgesturechange")
	return MSGestureChange
}

func (*Element) SetOnmsgesturechange(MSGestureChange *MSGestureEvent) {
	js.Rewrite("$<.onmsgesturechange = $1", MSGestureChange)
}

func (*Element) GetOnmsgesturedoubletap() (MSGestureDoubleTap *MSGestureEvent) {
	js.Rewrite("$<.onmsgesturedoubletap")
	return MSGestureDoubleTap
}

func (*Element) SetOnmsgesturedoubletap(MSGestureDoubleTap *MSGestureEvent) {
	js.Rewrite("$<.onmsgesturedoubletap = $1", MSGestureDoubleTap)
}

func (*Element) GetOnmsgestureend() (MSGestureEnd *MSGestureEvent) {
	js.Rewrite("$<.onmsgestureend")
	return MSGestureEnd
}

func (*Element) SetOnmsgestureend(MSGestureEnd *MSGestureEvent) {
	js.Rewrite("$<.onmsgestureend = $1", MSGestureEnd)
}

func (*Element) GetOnmsgesturehold() (MSGestureHold *MSGestureEvent) {
	js.Rewrite("$<.onmsgesturehold")
	return MSGestureHold
}

func (*Element) SetOnmsgesturehold(MSGestureHold *MSGestureEvent) {
	js.Rewrite("$<.onmsgesturehold = $1", MSGestureHold)
}

func (*Element) GetOnmsgesturestart() (MSGestureStart *MSGestureEvent) {
	js.Rewrite("$<.onmsgesturestart")
	return MSGestureStart
}

func (*Element) SetOnmsgesturestart(MSGestureStart *MSGestureEvent) {
	js.Rewrite("$<.onmsgesturestart = $1", MSGestureStart)
}

func (*Element) GetOnmsgesturetap() (MSGestureTap *MSGestureEvent) {
	js.Rewrite("$<.onmsgesturetap")
	return MSGestureTap
}

func (*Element) SetOnmsgesturetap(MSGestureTap *MSGestureEvent) {
	js.Rewrite("$<.onmsgesturetap = $1", MSGestureTap)
}

func (*Element) GetOnmsgotpointercapture() (MSGotPointerCapture *MSPointerEvent) {
	js.Rewrite("$<.onmsgotpointercapture")
	return MSGotPointerCapture
}

func (*Element) SetOnmsgotpointercapture(MSGotPointerCapture *MSPointerEvent) {
	js.Rewrite("$<.onmsgotpointercapture = $1", MSGotPointerCapture)
}

func (*Element) GetOnmsinertiastart() (MSInertiaStart *MSGestureEvent) {
	js.Rewrite("$<.onmsinertiastart")
	return MSInertiaStart
}

func (*Element) SetOnmsinertiastart(MSInertiaStart *MSGestureEvent) {
	js.Rewrite("$<.onmsinertiastart = $1", MSInertiaStart)
}

func (*Element) GetOnmslostpointercapture() (MSLostPointerCapture *MSPointerEvent) {
	js.Rewrite("$<.onmslostpointercapture")
	return MSLostPointerCapture
}

func (*Element) SetOnmslostpointercapture(MSLostPointerCapture *MSPointerEvent) {
	js.Rewrite("$<.onmslostpointercapture = $1", MSLostPointerCapture)
}

func (*Element) GetOnmspointercancel() (MSPointerCancel *MSPointerEvent) {
	js.Rewrite("$<.onmspointercancel")
	return MSPointerCancel
}

func (*Element) SetOnmspointercancel(MSPointerCancel *MSPointerEvent) {
	js.Rewrite("$<.onmspointercancel = $1", MSPointerCancel)
}

func (*Element) GetOnmspointerdown() (MSPointerDown *MSPointerEvent) {
	js.Rewrite("$<.onmspointerdown")
	return MSPointerDown
}

func (*Element) SetOnmspointerdown(MSPointerDown *MSPointerEvent) {
	js.Rewrite("$<.onmspointerdown = $1", MSPointerDown)
}

func (*Element) GetOnmspointerenter() (MSPointerEnter *MSPointerEvent) {
	js.Rewrite("$<.onmspointerenter")
	return MSPointerEnter
}

func (*Element) SetOnmspointerenter(MSPointerEnter *MSPointerEvent) {
	js.Rewrite("$<.onmspointerenter = $1", MSPointerEnter)
}

func (*Element) GetOnmspointerleave() (MSPointerLeave *MSPointerEvent) {
	js.Rewrite("$<.onmspointerleave")
	return MSPointerLeave
}

func (*Element) SetOnmspointerleave(MSPointerLeave *MSPointerEvent) {
	js.Rewrite("$<.onmspointerleave = $1", MSPointerLeave)
}

func (*Element) GetOnmspointermove() (MSPointerMove *MSPointerEvent) {
	js.Rewrite("$<.onmspointermove")
	return MSPointerMove
}

func (*Element) SetOnmspointermove(MSPointerMove *MSPointerEvent) {
	js.Rewrite("$<.onmspointermove = $1", MSPointerMove)
}

func (*Element) GetOnmspointerout() (MSPointerOut *MSPointerEvent) {
	js.Rewrite("$<.onmspointerout")
	return MSPointerOut
}

func (*Element) SetOnmspointerout(MSPointerOut *MSPointerEvent) {
	js.Rewrite("$<.onmspointerout = $1", MSPointerOut)
}

func (*Element) GetOnmspointerover() (MSPointerOver *MSPointerEvent) {
	js.Rewrite("$<.onmspointerover")
	return MSPointerOver
}

func (*Element) SetOnmspointerover(MSPointerOver *MSPointerEvent) {
	js.Rewrite("$<.onmspointerover = $1", MSPointerOver)
}

func (*Element) GetOnmspointerup() (MSPointerUp *MSPointerEvent) {
	js.Rewrite("$<.onmspointerup")
	return MSPointerUp
}

func (*Element) SetOnmspointerup(MSPointerUp *MSPointerEvent) {
	js.Rewrite("$<.onmspointerup = $1", MSPointerUp)
}

func (*Element) GetOntouchcancel() (touchcancel *TouchEvent) {
	js.Rewrite("$<.ontouchcancel")
	return touchcancel
}

func (*Element) SetOntouchcancel(touchcancel *TouchEvent) {
	js.Rewrite("$<.ontouchcancel = $1", touchcancel)
}

func (*Element) GetOntouchend() (touchend *TouchEvent) {
	js.Rewrite("$<.ontouchend")
	return touchend
}

func (*Element) SetOntouchend(touchend *TouchEvent) {
	js.Rewrite("$<.ontouchend = $1", touchend)
}

func (*Element) GetOntouchmove() (touchmove *TouchEvent) {
	js.Rewrite("$<.ontouchmove")
	return touchmove
}

func (*Element) SetOntouchmove(touchmove *TouchEvent) {
	js.Rewrite("$<.ontouchmove = $1", touchmove)
}

func (*Element) GetOntouchstart() (touchstart *TouchEvent) {
	js.Rewrite("$<.ontouchstart")
	return touchstart
}

func (*Element) SetOntouchstart(touchstart *TouchEvent) {
	js.Rewrite("$<.ontouchstart = $1", touchstart)
}

func (*Element) GetOnwebkitfullscreenchange() (e *Event) {
	js.Rewrite("$<.onwebkitfullscreenchange")
	return e
}

func (*Element) SetOnwebkitfullscreenchange(e *Event) {
	js.Rewrite("$<.onwebkitfullscreenchange = $1", e)
}

func (*Element) GetOnwebkitfullscreenerror() (e *Event) {
	js.Rewrite("$<.onwebkitfullscreenerror")
	return e
}

func (*Element) SetOnwebkitfullscreenerror(e *Event) {
	js.Rewrite("$<.onwebkitfullscreenerror = $1", e)
}

func (*Element) GetOuterHTML() (outerHTML string) {
	js.Rewrite("$<.outerHTML")
	return outerHTML
}

func (*Element) SetOuterHTML(outerHTML string) {
	js.Rewrite("$<.outerHTML = $1", outerHTML)
}

func (*Element) GetPrefix() (prefix string) {
	js.Rewrite("$<.prefix")
	return prefix
}

func (*Element) GetScrollHeight() (scrollHeight int) {
	js.Rewrite("$<.scrollHeight")
	return scrollHeight
}

func (*Element) GetScrollLeft() (scrollLeft int) {
	js.Rewrite("$<.scrollLeft")
	return scrollLeft
}

func (*Element) SetScrollLeft(scrollLeft int) {
	js.Rewrite("$<.scrollLeft = $1", scrollLeft)
}

func (*Element) GetScrollTop() (scrollTop int) {
	js.Rewrite("$<.scrollTop")
	return scrollTop
}

func (*Element) SetScrollTop(scrollTop int) {
	js.Rewrite("$<.scrollTop = $1", scrollTop)
}

func (*Element) GetScrollWidth() (scrollWidth int) {
	js.Rewrite("$<.scrollWidth")
	return scrollWidth
}

func (*Element) GetTagName() (tagName string) {
	js.Rewrite("$<.tagName")
	return tagName
}

type ErrorEvent struct {
	*Event
}

func (*ErrorEvent) InitErrorEvent(typeArg string, canBubbleArg bool, cancelableArg bool, messageArg string, filenameArg string, linenoArg uint) {
	js.Rewrite("$<.initErrorEvent($1, $2, $3, $4, $5, $6)", typeArg, canBubbleArg, cancelableArg, messageArg, filenameArg, linenoArg)
}

func (*ErrorEvent) GetColno() (colno uint) {
	js.Rewrite("$<.colno")
	return colno
}

func (*ErrorEvent) GetError() (err interface{}) {
	js.Rewrite("$<.error")
	return err
}

func (*ErrorEvent) GetFilename() (filename string) {
	js.Rewrite("$<.filename")
	return filename
}

func (*ErrorEvent) GetLineno() (lineno uint) {
	js.Rewrite("$<.lineno")
	return lineno
}

func (*ErrorEvent) GetMessage() (message string) {
	js.Rewrite("$<.message")
	return message
}

type Event struct {
}

func (*Event) InitEvent(eventTypeArg string, canBubbleArg bool, cancelableArg bool) {
	js.Rewrite("$<.initEvent($1, $2, $3)", eventTypeArg, canBubbleArg, cancelableArg)
}

func (*Event) PreventDefault() {
	js.Rewrite("$<.preventDefault()")
}

func (*Event) StopImmediatePropagation() {
	js.Rewrite("$<.stopImmediatePropagation()")
}

func (*Event) StopPropagation() {
	js.Rewrite("$<.stopPropagation()")
}

func (*Event) GetBubbles() (bubbles bool) {
	js.Rewrite("$<.bubbles")
	return bubbles
}

func (*Event) GetCancelable() (cancelable bool) {
	js.Rewrite("$<.cancelable")
	return cancelable
}

func (*Event) GetCancelBubble() (cancelBubble bool) {
	js.Rewrite("$<.cancelBubble")
	return cancelBubble
}

func (*Event) SetCancelBubble(cancelBubble bool) {
	js.Rewrite("$<.cancelBubble = $1", cancelBubble)
}

func (*Event) GetCurrentTarget() (currentTarget *EventTarget) {
	js.Rewrite("$<.currentTarget")
	return currentTarget
}

func (*Event) GetDefaultPrevented() (defaultPrevented bool) {
	js.Rewrite("$<.defaultPrevented")
	return defaultPrevented
}

func (*Event) GetEventPhase() (eventPhase uint8) {
	js.Rewrite("$<.eventPhase")
	return eventPhase
}

func (*Event) GetIsTrusted() (isTrusted bool) {
	js.Rewrite("$<.isTrusted")
	return isTrusted
}

func (*Event) GetReturnValue() (returnValue bool) {
	js.Rewrite("$<.returnValue")
	return returnValue
}

func (*Event) SetReturnValue(returnValue bool) {
	js.Rewrite("$<.returnValue = $1", returnValue)
}

func (*Event) GetSrcElement() (srcElement *Element) {
	js.Rewrite("$<.srcElement")
	return srcElement
}

func (*Event) GetTarget() (target *EventTarget) {
	js.Rewrite("$<.target")
	return target
}

func (*Event) GetTimeStamp() (timeStamp uint64) {
	js.Rewrite("$<.timeStamp")
	return timeStamp
}

func (*Event) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

type EventTarget struct {
}

func (*EventTarget) AddEventListener(kind string, listener func(evt *Event), useCapture bool) {
	js.Rewrite("$<.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*EventTarget) DispatchEvent(evt *Event) (b bool) {
	js.Rewrite("$<.dispatchEvent($1)", evt)
	return b
}

func (*EventTarget) RemoveEventListener(kind string, listener func(evt *Event), useCapture bool) {
	js.Rewrite("$<.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

type EXT_frag_depth struct {
}

type EXT_texture_filter_anisotropic struct {
}

type ExtensionScriptApis struct {
}

func (*ExtensionScriptApis) ExtensionIDToShortID(extensionId string) (i int) {
	js.Rewrite("$<.extensionIdToShortId($1)", extensionId)
	return i
}

func (*ExtensionScriptApis) FireExtensionAPITelemetry(functionName string, isSucceeded bool, isSupported bool) {
	js.Rewrite("$<.fireExtensionApiTelemetry($1, $2, $3)", functionName, isSucceeded, isSupported)
}

func (*ExtensionScriptApis) GenericFunction(routerAddress interface{}, parameters string, callbackId int) {
	js.Rewrite("$<.genericFunction($1, $2, $3)", routerAddress, parameters, callbackId)
}

func (*ExtensionScriptApis) GenericSynchronousFunction(functionId int, parameters string) (s string) {
	js.Rewrite("$<.genericSynchronousFunction($1, $2)", functionId, parameters)
	return s
}

func (*ExtensionScriptApis) GetExtensionID() (s string) {
	js.Rewrite("$<.getExtensionId()")
	return s
}

func (*ExtensionScriptApis) RegisterGenericFunctionCallbackHandler(callbackHandler func()) {
	js.Rewrite("$<.registerGenericFunctionCallbackHandler($1)", callbackHandler)
}

func (*ExtensionScriptApis) RegisterGenericPersistentCallbackHandler(callbackHandler func()) {
	js.Rewrite("$<.registerGenericPersistentCallbackHandler($1)", callbackHandler)
}

type External struct {
}

type File struct {
	*Blob
}

func (*File) GetLastModifiedDate() (lastModifiedDate interface{}) {
	js.Rewrite("$<.lastModifiedDate")
	return lastModifiedDate
}

func (*File) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*File) GetWebkitRelativePath() (webkitRelativePath string) {
	js.Rewrite("$<.webkitRelativePath")
	return webkitRelativePath
}

type FileList struct {
}

func (*FileList) Item(index uint) (f *File) {
	js.Rewrite("$<.item($1)", index)
	return f
}

func (*FileList) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

type FileReader struct {
	*EventTarget
	*MSBaseReader
}

func (*FileReader) ReadAsArrayBuffer(blob *Blob) {
	js.Rewrite("$<.readAsArrayBuffer($1)", blob)
}

func (*FileReader) ReadAsBinaryString(blob *Blob) {
	js.Rewrite("$<.readAsBinaryString($1)", blob)
}

func (*FileReader) ReadAsDataURL(blob *Blob) {
	js.Rewrite("$<.readAsDataURL($1)", blob)
}

func (*FileReader) ReadAsText(blob *Blob, encoding string) {
	js.Rewrite("$<.readAsText($1, $2)", blob, encoding)
}

func (*FileReader) GetError() (err *DOMError) {
	js.Rewrite("$<.error")
	return err
}

type FocusEvent struct {
	*UIEvent
}

func (*FocusEvent) InitFocusEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, relatedTargetArg *EventTarget) {
	js.Rewrite("$<.initFocusEvent($1, $2, $3, $4, $5, $6)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, relatedTargetArg)
}

func (*FocusEvent) GetRelatedTarget() (relatedTarget *EventTarget) {
	js.Rewrite("$<.relatedTarget")
	return relatedTarget
}

type FocusNavigationEvent struct {
	*Event
}

func (*FocusNavigationEvent) RequestFocus() {
	js.Rewrite("$<.requestFocus()")
}

func (*FocusNavigationEvent) GetNavigationReason() (navigationReason *NavigationReason) {
	js.Rewrite("$<.navigationReason")
	return navigationReason
}

func (*FocusNavigationEvent) GetOriginHeight() (originHeight float32) {
	js.Rewrite("$<.originHeight")
	return originHeight
}

func (*FocusNavigationEvent) GetOriginLeft() (originLeft float32) {
	js.Rewrite("$<.originLeft")
	return originLeft
}

func (*FocusNavigationEvent) GetOriginTop() (originTop float32) {
	js.Rewrite("$<.originTop")
	return originTop
}

func (*FocusNavigationEvent) GetOriginWidth() (originWidth float32) {
	js.Rewrite("$<.originWidth")
	return originWidth
}

type FormData struct {
}

func (*FormData) Append(name interface{}, value interface{}, blobName string) {
	js.Rewrite("$<.append($1, $2, $3)", name, value, blobName)
}

type GainNode struct {
	*AudioNode
}

func (*GainNode) GetGain() (gain *AudioParam) {
	js.Rewrite("$<.gain")
	return gain
}

type Gamepad struct {
}

func (*Gamepad) GetAxes() (axes []float32) {
	js.Rewrite("$<.axes")
	return axes
}

func (*Gamepad) GetButtons() (buttons []*GamepadButton) {
	js.Rewrite("$<.buttons")
	return buttons
}

func (*Gamepad) GetConnected() (connected bool) {
	js.Rewrite("$<.connected")
	return connected
}

func (*Gamepad) GetID() (id string) {
	js.Rewrite("$<.id")
	return id
}

func (*Gamepad) GetIndex() (index int) {
	js.Rewrite("$<.index")
	return index
}

func (*Gamepad) GetMapping() (mapping string) {
	js.Rewrite("$<.mapping")
	return mapping
}

func (*Gamepad) GetTimestamp() (timestamp int) {
	js.Rewrite("$<.timestamp")
	return timestamp
}

type GamepadButton struct {
}

func (*GamepadButton) GetPressed() (pressed bool) {
	js.Rewrite("$<.pressed")
	return pressed
}

func (*GamepadButton) GetValue() (value float32) {
	js.Rewrite("$<.value")
	return value
}

type GamepadEvent struct {
	*Event
}

func (*GamepadEvent) GetGamepad() (gamepad *Gamepad) {
	js.Rewrite("$<.gamepad")
	return gamepad
}

type Geolocation struct {
}

func (*Geolocation) ClearWatch(watchId int) {
	js.Rewrite("$<.clearWatch($1)", watchId)
}

func (*Geolocation) GetCurrentPosition(successCallback func(position *Position), errorCallback func(err *PositionError), options *PositionOptions) {
	js.Rewrite("$<.getCurrentPosition($1, $2, $3)", successCallback, errorCallback, options)
}

func (*Geolocation) WatchPosition(successCallback func(position *Position), errorCallback func(err *PositionError), options *PositionOptions) (i int) {
	js.Rewrite("$<.watchPosition($1, $2, $3)", successCallback, errorCallback, options)
	return i
}

type HashChangeEvent struct {
	*Event
}

func (*HashChangeEvent) GetNewURL() (newURL string) {
	js.Rewrite("$<.newURL")
	return newURL
}

func (*HashChangeEvent) GetOldURL() (oldURL string) {
	js.Rewrite("$<.oldURL")
	return oldURL
}

type Headers struct {
}

func (*Headers) Append(name string, value string) {
	js.Rewrite("$<.append($1, $2)", name, value)
}

func (*Headers) Delete(name string) {
	js.Rewrite("$<.delete($1)", name)
}

func (*Headers) ForEach(callback func(keyId []byte, status *MediaKeyStatus)) {
	js.Rewrite("$<.forEach($1)", callback)
}

func (*Headers) Get(name string) (s string) {
	js.Rewrite("$<.get($1)", name)
	return s
}

func (*Headers) Has(name string) (b bool) {
	js.Rewrite("$<.has($1)", name)
	return b
}

func (*Headers) Set(name string, value string) {
	js.Rewrite("$<.set($1, $2)", name, value)
}

type History struct {
}

func (*History) Back(distance interface{}) {
	js.Rewrite("$<.back($1)", distance)
}

func (*History) Forward(distance interface{}) {
	js.Rewrite("$<.forward($1)", distance)
}

func (*History) Go(delta interface{}) {
	js.Rewrite("$<.go($1)", delta)
}

func (*History) PushState(statedata interface{}, title string, url string) {
	js.Rewrite("$<.pushState($1, $2, $3)", statedata, title, url)
}

func (*History) ReplaceState(statedata interface{}, title string, url string) {
	js.Rewrite("$<.replaceState($1, $2, $3)", statedata, title, url)
}

func (*History) GetLength() (length int) {
	js.Rewrite("$<.length")
	return length
}

func (*History) GetState() (state interface{}) {
	js.Rewrite("$<.state")
	return state
}

type HTMLAllCollection struct {
}

func (*HTMLAllCollection) Item(nameOrIndex string) (i interface{}) {
	js.Rewrite("$<.item($1)", nameOrIndex)
	return i
}

func (*HTMLAllCollection) NamedItem(name string) (i interface{}) {
	js.Rewrite("$<.namedItem($1)", name)
	return i
}

func (*HTMLAllCollection) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

type HTMLAnchorElement struct {
	*HTMLElement
}

func (*HTMLAnchorElement) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

func (*HTMLAnchorElement) GetCharset() (charset string) {
	js.Rewrite("$<.charset")
	return charset
}

func (*HTMLAnchorElement) SetCharset(charset string) {
	js.Rewrite("$<.charset = $1", charset)
}

func (*HTMLAnchorElement) GetCoords() (coords string) {
	js.Rewrite("$<.coords")
	return coords
}

func (*HTMLAnchorElement) SetCoords(coords string) {
	js.Rewrite("$<.coords = $1", coords)
}

func (*HTMLAnchorElement) GetDownload() (download string) {
	js.Rewrite("$<.download")
	return download
}

func (*HTMLAnchorElement) SetDownload(download string) {
	js.Rewrite("$<.download = $1", download)
}

func (*HTMLAnchorElement) GetHash() (hash string) {
	js.Rewrite("$<.hash")
	return hash
}

func (*HTMLAnchorElement) SetHash(hash string) {
	js.Rewrite("$<.hash = $1", hash)
}

func (*HTMLAnchorElement) GetHost() (host string) {
	js.Rewrite("$<.host")
	return host
}

func (*HTMLAnchorElement) SetHost(host string) {
	js.Rewrite("$<.host = $1", host)
}

func (*HTMLAnchorElement) GetHostname() (hostname string) {
	js.Rewrite("$<.hostname")
	return hostname
}

func (*HTMLAnchorElement) SetHostname(hostname string) {
	js.Rewrite("$<.hostname = $1", hostname)
}

func (*HTMLAnchorElement) GetHref() (href string) {
	js.Rewrite("$<.href")
	return href
}

func (*HTMLAnchorElement) SetHref(href string) {
	js.Rewrite("$<.href = $1", href)
}

func (*HTMLAnchorElement) GetHreflang() (hreflang string) {
	js.Rewrite("$<.hreflang")
	return hreflang
}

func (*HTMLAnchorElement) SetHreflang(hreflang string) {
	js.Rewrite("$<.hreflang = $1", hreflang)
}

func (*HTMLAnchorElement) GetMethods() (Methods string) {
	js.Rewrite("$<.Methods")
	return Methods
}

func (*HTMLAnchorElement) SetMethods(Methods string) {
	js.Rewrite("$<.Methods = $1", Methods)
}

func (*HTMLAnchorElement) GetMimeType() (mimeType string) {
	js.Rewrite("$<.mimeType")
	return mimeType
}

func (*HTMLAnchorElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLAnchorElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLAnchorElement) GetNameProp() (nameProp string) {
	js.Rewrite("$<.nameProp")
	return nameProp
}

func (*HTMLAnchorElement) GetPathname() (pathname string) {
	js.Rewrite("$<.pathname")
	return pathname
}

func (*HTMLAnchorElement) SetPathname(pathname string) {
	js.Rewrite("$<.pathname = $1", pathname)
}

func (*HTMLAnchorElement) GetPort() (port string) {
	js.Rewrite("$<.port")
	return port
}

func (*HTMLAnchorElement) SetPort(port string) {
	js.Rewrite("$<.port = $1", port)
}

func (*HTMLAnchorElement) GetProtocol() (protocol string) {
	js.Rewrite("$<.protocol")
	return protocol
}

func (*HTMLAnchorElement) SetProtocol(protocol string) {
	js.Rewrite("$<.protocol = $1", protocol)
}

func (*HTMLAnchorElement) GetProtocolLong() (protocolLong string) {
	js.Rewrite("$<.protocolLong")
	return protocolLong
}

func (*HTMLAnchorElement) GetRel() (rel string) {
	js.Rewrite("$<.rel")
	return rel
}

func (*HTMLAnchorElement) SetRel(rel string) {
	js.Rewrite("$<.rel = $1", rel)
}

func (*HTMLAnchorElement) GetRev() (rev string) {
	js.Rewrite("$<.rev")
	return rev
}

func (*HTMLAnchorElement) SetRev(rev string) {
	js.Rewrite("$<.rev = $1", rev)
}

func (*HTMLAnchorElement) GetSearch() (search string) {
	js.Rewrite("$<.search")
	return search
}

func (*HTMLAnchorElement) SetSearch(search string) {
	js.Rewrite("$<.search = $1", search)
}

func (*HTMLAnchorElement) GetShape() (shape string) {
	js.Rewrite("$<.shape")
	return shape
}

func (*HTMLAnchorElement) SetShape(shape string) {
	js.Rewrite("$<.shape = $1", shape)
}

func (*HTMLAnchorElement) GetTarget() (target string) {
	js.Rewrite("$<.target")
	return target
}

func (*HTMLAnchorElement) SetTarget(target string) {
	js.Rewrite("$<.target = $1", target)
}

func (*HTMLAnchorElement) GetText() (text string) {
	js.Rewrite("$<.text")
	return text
}

func (*HTMLAnchorElement) SetText(text string) {
	js.Rewrite("$<.text = $1", text)
}

func (*HTMLAnchorElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLAnchorElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}

func (*HTMLAnchorElement) GetUrn() (urn string) {
	js.Rewrite("$<.urn")
	return urn
}

func (*HTMLAnchorElement) SetUrn(urn string) {
	js.Rewrite("$<.urn = $1", urn)
}

type HTMLAppletElement struct {
	*HTMLElement
}

func (*HTMLAppletElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLAppletElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLAppletElement) GetAlt() (alt string) {
	js.Rewrite("$<.alt")
	return alt
}

func (*HTMLAppletElement) SetAlt(alt string) {
	js.Rewrite("$<.alt = $1", alt)
}

func (*HTMLAppletElement) GetAltHTML() (altHtml string) {
	js.Rewrite("$<.altHtml")
	return altHtml
}

func (*HTMLAppletElement) SetAltHTML(altHtml string) {
	js.Rewrite("$<.altHtml = $1", altHtml)
}

func (*HTMLAppletElement) GetArchive() (archive string) {
	js.Rewrite("$<.archive")
	return archive
}

func (*HTMLAppletElement) SetArchive(archive string) {
	js.Rewrite("$<.archive = $1", archive)
}

func (*HTMLAppletElement) GetBaseHref() (BaseHref string) {
	js.Rewrite("$<.BaseHref")
	return BaseHref
}

func (*HTMLAppletElement) GetBorder() (border string) {
	js.Rewrite("$<.border")
	return border
}

func (*HTMLAppletElement) SetBorder(border string) {
	js.Rewrite("$<.border = $1", border)
}

func (*HTMLAppletElement) GetCode() (code string) {
	js.Rewrite("$<.code")
	return code
}

func (*HTMLAppletElement) SetCode(code string) {
	js.Rewrite("$<.code = $1", code)
}

func (*HTMLAppletElement) GetCodeBase() (codeBase string) {
	js.Rewrite("$<.codeBase")
	return codeBase
}

func (*HTMLAppletElement) SetCodeBase(codeBase string) {
	js.Rewrite("$<.codeBase = $1", codeBase)
}

func (*HTMLAppletElement) GetCodeType() (codeType string) {
	js.Rewrite("$<.codeType")
	return codeType
}

func (*HTMLAppletElement) SetCodeType(codeType string) {
	js.Rewrite("$<.codeType = $1", codeType)
}

func (*HTMLAppletElement) GetContentDocument() (contentDocument *Document) {
	js.Rewrite("$<.contentDocument")
	return contentDocument
}

func (*HTMLAppletElement) GetData() (data string) {
	js.Rewrite("$<.data")
	return data
}

func (*HTMLAppletElement) SetData(data string) {
	js.Rewrite("$<.data = $1", data)
}

func (*HTMLAppletElement) GetDeclare() (declare bool) {
	js.Rewrite("$<.declare")
	return declare
}

func (*HTMLAppletElement) SetDeclare(declare bool) {
	js.Rewrite("$<.declare = $1", declare)
}

func (*HTMLAppletElement) GetForm() (form *HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

func (*HTMLAppletElement) GetHeight() (height string) {
	js.Rewrite("$<.height")
	return height
}

func (*HTMLAppletElement) SetHeight(height string) {
	js.Rewrite("$<.height = $1", height)
}

func (*HTMLAppletElement) GetHspace() (hspace int) {
	js.Rewrite("$<.hspace")
	return hspace
}

func (*HTMLAppletElement) SetHspace(hspace int) {
	js.Rewrite("$<.hspace = $1", hspace)
}

func (*HTMLAppletElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLAppletElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLAppletElement) GetObject() (object string) {
	js.Rewrite("$<.object")
	return object
}

func (*HTMLAppletElement) SetObject(object string) {
	js.Rewrite("$<.object = $1", object)
}

func (*HTMLAppletElement) GetStandby() (standby string) {
	js.Rewrite("$<.standby")
	return standby
}

func (*HTMLAppletElement) SetStandby(standby string) {
	js.Rewrite("$<.standby = $1", standby)
}

func (*HTMLAppletElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLAppletElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}

func (*HTMLAppletElement) GetUseMap() (useMap string) {
	js.Rewrite("$<.useMap")
	return useMap
}

func (*HTMLAppletElement) SetUseMap(useMap string) {
	js.Rewrite("$<.useMap = $1", useMap)
}

func (*HTMLAppletElement) GetVspace() (vspace int) {
	js.Rewrite("$<.vspace")
	return vspace
}

func (*HTMLAppletElement) SetVspace(vspace int) {
	js.Rewrite("$<.vspace = $1", vspace)
}

func (*HTMLAppletElement) GetWidth() (width int) {
	js.Rewrite("$<.width")
	return width
}

func (*HTMLAppletElement) SetWidth(width int) {
	js.Rewrite("$<.width = $1", width)
}

type HTMLAreaElement struct {
	*HTMLElement
}

func (*HTMLAreaElement) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

func (*HTMLAreaElement) GetAlt() (alt string) {
	js.Rewrite("$<.alt")
	return alt
}

func (*HTMLAreaElement) SetAlt(alt string) {
	js.Rewrite("$<.alt = $1", alt)
}

func (*HTMLAreaElement) GetCoords() (coords string) {
	js.Rewrite("$<.coords")
	return coords
}

func (*HTMLAreaElement) SetCoords(coords string) {
	js.Rewrite("$<.coords = $1", coords)
}

func (*HTMLAreaElement) GetDownload() (download string) {
	js.Rewrite("$<.download")
	return download
}

func (*HTMLAreaElement) SetDownload(download string) {
	js.Rewrite("$<.download = $1", download)
}

func (*HTMLAreaElement) GetHash() (hash string) {
	js.Rewrite("$<.hash")
	return hash
}

func (*HTMLAreaElement) SetHash(hash string) {
	js.Rewrite("$<.hash = $1", hash)
}

func (*HTMLAreaElement) GetHost() (host string) {
	js.Rewrite("$<.host")
	return host
}

func (*HTMLAreaElement) SetHost(host string) {
	js.Rewrite("$<.host = $1", host)
}

func (*HTMLAreaElement) GetHostname() (hostname string) {
	js.Rewrite("$<.hostname")
	return hostname
}

func (*HTMLAreaElement) SetHostname(hostname string) {
	js.Rewrite("$<.hostname = $1", hostname)
}

func (*HTMLAreaElement) GetHref() (href string) {
	js.Rewrite("$<.href")
	return href
}

func (*HTMLAreaElement) SetHref(href string) {
	js.Rewrite("$<.href = $1", href)
}

func (*HTMLAreaElement) GetNoHref() (noHref bool) {
	js.Rewrite("$<.noHref")
	return noHref
}

func (*HTMLAreaElement) SetNoHref(noHref bool) {
	js.Rewrite("$<.noHref = $1", noHref)
}

func (*HTMLAreaElement) GetPathname() (pathname string) {
	js.Rewrite("$<.pathname")
	return pathname
}

func (*HTMLAreaElement) SetPathname(pathname string) {
	js.Rewrite("$<.pathname = $1", pathname)
}

func (*HTMLAreaElement) GetPort() (port string) {
	js.Rewrite("$<.port")
	return port
}

func (*HTMLAreaElement) SetPort(port string) {
	js.Rewrite("$<.port = $1", port)
}

func (*HTMLAreaElement) GetProtocol() (protocol string) {
	js.Rewrite("$<.protocol")
	return protocol
}

func (*HTMLAreaElement) SetProtocol(protocol string) {
	js.Rewrite("$<.protocol = $1", protocol)
}

func (*HTMLAreaElement) GetRel() (rel string) {
	js.Rewrite("$<.rel")
	return rel
}

func (*HTMLAreaElement) SetRel(rel string) {
	js.Rewrite("$<.rel = $1", rel)
}

func (*HTMLAreaElement) GetSearch() (search string) {
	js.Rewrite("$<.search")
	return search
}

func (*HTMLAreaElement) SetSearch(search string) {
	js.Rewrite("$<.search = $1", search)
}

func (*HTMLAreaElement) GetShape() (shape string) {
	js.Rewrite("$<.shape")
	return shape
}

func (*HTMLAreaElement) SetShape(shape string) {
	js.Rewrite("$<.shape = $1", shape)
}

func (*HTMLAreaElement) GetTarget() (target string) {
	js.Rewrite("$<.target")
	return target
}

func (*HTMLAreaElement) SetTarget(target string) {
	js.Rewrite("$<.target = $1", target)
}

type HTMLAreasCollection struct {
	*HTMLCollection
}

type HTMLAudioElement struct {
	*HTMLMediaElement
}

type HTMLBaseElement struct {
	*HTMLElement
}

func (*HTMLBaseElement) GetHref() (href string) {
	js.Rewrite("$<.href")
	return href
}

func (*HTMLBaseElement) SetHref(href string) {
	js.Rewrite("$<.href = $1", href)
}

func (*HTMLBaseElement) GetTarget() (target string) {
	js.Rewrite("$<.target")
	return target
}

func (*HTMLBaseElement) SetTarget(target string) {
	js.Rewrite("$<.target = $1", target)
}

type HTMLBaseFontElement struct {
	*HTMLElement
	*DOML2DeprecatedColorProperty
}

func (*HTMLBaseFontElement) GetFace() (face string) {
	js.Rewrite("$<.face")
	return face
}

func (*HTMLBaseFontElement) SetFace(face string) {
	js.Rewrite("$<.face = $1", face)
}

func (*HTMLBaseFontElement) GetSize() (size int) {
	js.Rewrite("$<.size")
	return size
}

func (*HTMLBaseFontElement) SetSize(size int) {
	js.Rewrite("$<.size = $1", size)
}

type HTMLBodyElement struct {
	*HTMLElement
}

func (*HTMLBodyElement) GetALink() (aLink interface{}) {
	js.Rewrite("$<.aLink")
	return aLink
}

func (*HTMLBodyElement) SetALink(aLink interface{}) {
	js.Rewrite("$<.aLink = $1", aLink)
}

func (*HTMLBodyElement) GetBackground() (background string) {
	js.Rewrite("$<.background")
	return background
}

func (*HTMLBodyElement) SetBackground(background string) {
	js.Rewrite("$<.background = $1", background)
}

func (*HTMLBodyElement) GetBgColor() (bgColor interface{}) {
	js.Rewrite("$<.bgColor")
	return bgColor
}

func (*HTMLBodyElement) SetBgColor(bgColor interface{}) {
	js.Rewrite("$<.bgColor = $1", bgColor)
}

func (*HTMLBodyElement) GetBgProperties() (bgProperties string) {
	js.Rewrite("$<.bgProperties")
	return bgProperties
}

func (*HTMLBodyElement) SetBgProperties(bgProperties string) {
	js.Rewrite("$<.bgProperties = $1", bgProperties)
}

func (*HTMLBodyElement) GetLink() (link interface{}) {
	js.Rewrite("$<.link")
	return link
}

func (*HTMLBodyElement) SetLink(link interface{}) {
	js.Rewrite("$<.link = $1", link)
}

func (*HTMLBodyElement) GetNoWrap() (noWrap bool) {
	js.Rewrite("$<.noWrap")
	return noWrap
}

func (*HTMLBodyElement) SetNoWrap(noWrap bool) {
	js.Rewrite("$<.noWrap = $1", noWrap)
}

func (*HTMLBodyElement) GetOnafterprint() (e *Event) {
	js.Rewrite("$<.onafterprint")
	return e
}

func (*HTMLBodyElement) SetOnafterprint(e *Event) {
	js.Rewrite("$<.onafterprint = $1", e)
}

func (*HTMLBodyElement) GetOnbeforeprint() (e *Event) {
	js.Rewrite("$<.onbeforeprint")
	return e
}

func (*HTMLBodyElement) SetOnbeforeprint(e *Event) {
	js.Rewrite("$<.onbeforeprint = $1", e)
}

func (*HTMLBodyElement) GetOnbeforeunload() (e *Event) {
	js.Rewrite("$<.onbeforeunload")
	return e
}

func (*HTMLBodyElement) SetOnbeforeunload(e *Event) {
	js.Rewrite("$<.onbeforeunload = $1", e)
}

func (*HTMLBodyElement) GetOnblur() (blur *FocusEvent) {
	js.Rewrite("$<.onblur")
	return blur
}

func (*HTMLBodyElement) SetOnblur(blur *FocusEvent) {
	js.Rewrite("$<.onblur = $1", blur)
}

func (*HTMLBodyElement) GetOnerror() (e *Event) {
	js.Rewrite("$<.onerror")
	return e
}

func (*HTMLBodyElement) SetOnerror(e *Event) {
	js.Rewrite("$<.onerror = $1", e)
}

func (*HTMLBodyElement) GetOnfocus() (focus *FocusEvent) {
	js.Rewrite("$<.onfocus")
	return focus
}

func (*HTMLBodyElement) SetOnfocus(focus *FocusEvent) {
	js.Rewrite("$<.onfocus = $1", focus)
}

func (*HTMLBodyElement) GetOnhashchange() (e *Event) {
	js.Rewrite("$<.onhashchange")
	return e
}

func (*HTMLBodyElement) SetOnhashchange(e *Event) {
	js.Rewrite("$<.onhashchange = $1", e)
}

func (*HTMLBodyElement) GetOnload() (e *Event) {
	js.Rewrite("$<.onload")
	return e
}

func (*HTMLBodyElement) SetOnload(e *Event) {
	js.Rewrite("$<.onload = $1", e)
}

func (*HTMLBodyElement) GetOnmessage() (e *Event) {
	js.Rewrite("$<.onmessage")
	return e
}

func (*HTMLBodyElement) SetOnmessage(e *Event) {
	js.Rewrite("$<.onmessage = $1", e)
}

func (*HTMLBodyElement) GetOnoffline() (offline *Event) {
	js.Rewrite("$<.onoffline")
	return offline
}

func (*HTMLBodyElement) SetOnoffline(offline *Event) {
	js.Rewrite("$<.onoffline = $1", offline)
}

func (*HTMLBodyElement) GetOnonline() (online *Event) {
	js.Rewrite("$<.ononline")
	return online
}

func (*HTMLBodyElement) SetOnonline(online *Event) {
	js.Rewrite("$<.ononline = $1", online)
}

func (*HTMLBodyElement) GetOnorientationchange() (e *Event) {
	js.Rewrite("$<.onorientationchange")
	return e
}

func (*HTMLBodyElement) SetOnorientationchange(e *Event) {
	js.Rewrite("$<.onorientationchange = $1", e)
}

func (*HTMLBodyElement) GetOnpagehide() (e *Event) {
	js.Rewrite("$<.onpagehide")
	return e
}

func (*HTMLBodyElement) SetOnpagehide(e *Event) {
	js.Rewrite("$<.onpagehide = $1", e)
}

func (*HTMLBodyElement) GetOnpageshow() (e *Event) {
	js.Rewrite("$<.onpageshow")
	return e
}

func (*HTMLBodyElement) SetOnpageshow(e *Event) {
	js.Rewrite("$<.onpageshow = $1", e)
}

func (*HTMLBodyElement) GetOnpopstate() (e *Event) {
	js.Rewrite("$<.onpopstate")
	return e
}

func (*HTMLBodyElement) SetOnpopstate(e *Event) {
	js.Rewrite("$<.onpopstate = $1", e)
}

func (*HTMLBodyElement) GetOnresize() (e *Event) {
	js.Rewrite("$<.onresize")
	return e
}

func (*HTMLBodyElement) SetOnresize(e *Event) {
	js.Rewrite("$<.onresize = $1", e)
}

func (*HTMLBodyElement) GetOnscroll() (scroll *UIEvent) {
	js.Rewrite("$<.onscroll")
	return scroll
}

func (*HTMLBodyElement) SetOnscroll(scroll *UIEvent) {
	js.Rewrite("$<.onscroll = $1", scroll)
}

func (*HTMLBodyElement) GetOnstorage() (e *Event) {
	js.Rewrite("$<.onstorage")
	return e
}

func (*HTMLBodyElement) SetOnstorage(e *Event) {
	js.Rewrite("$<.onstorage = $1", e)
}

func (*HTMLBodyElement) GetOnunload() (e *Event) {
	js.Rewrite("$<.onunload")
	return e
}

func (*HTMLBodyElement) SetOnunload(e *Event) {
	js.Rewrite("$<.onunload = $1", e)
}

func (*HTMLBodyElement) GetText() (text interface{}) {
	js.Rewrite("$<.text")
	return text
}

func (*HTMLBodyElement) SetText(text interface{}) {
	js.Rewrite("$<.text = $1", text)
}

func (*HTMLBodyElement) GetVLink() (vLink interface{}) {
	js.Rewrite("$<.vLink")
	return vLink
}

func (*HTMLBodyElement) SetVLink(vLink interface{}) {
	js.Rewrite("$<.vLink = $1", vLink)
}

type HTMLBRElement struct {
	*HTMLElement
}

func (*HTMLBRElement) GetClear() (clear string) {
	js.Rewrite("$<.clear")
	return clear
}

func (*HTMLBRElement) SetClear(clear string) {
	js.Rewrite("$<.clear = $1", clear)
}

type HTMLButtonElement struct {
	*HTMLElement
}

func (*HTMLButtonElement) CheckValidity() (b bool) {
	js.Rewrite("$<.checkValidity()")
	return b
}

func (*HTMLButtonElement) SetCustomValidity(err string) {
	js.Rewrite("$<.setCustomValidity($1)", err)
}

func (*HTMLButtonElement) GetAutofocus() (autofocus bool) {
	js.Rewrite("$<.autofocus")
	return autofocus
}

func (*HTMLButtonElement) SetAutofocus(autofocus bool) {
	js.Rewrite("$<.autofocus = $1", autofocus)
}

func (*HTMLButtonElement) GetDisabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

func (*HTMLButtonElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

func (*HTMLButtonElement) GetForm() (form *HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

func (*HTMLButtonElement) GetFormAction() (formAction string) {
	js.Rewrite("$<.formAction")
	return formAction
}

func (*HTMLButtonElement) SetFormAction(formAction string) {
	js.Rewrite("$<.formAction = $1", formAction)
}

func (*HTMLButtonElement) GetFormEnctype() (formEnctype string) {
	js.Rewrite("$<.formEnctype")
	return formEnctype
}

func (*HTMLButtonElement) SetFormEnctype(formEnctype string) {
	js.Rewrite("$<.formEnctype = $1", formEnctype)
}

func (*HTMLButtonElement) GetFormMethod() (formMethod string) {
	js.Rewrite("$<.formMethod")
	return formMethod
}

func (*HTMLButtonElement) SetFormMethod(formMethod string) {
	js.Rewrite("$<.formMethod = $1", formMethod)
}

func (*HTMLButtonElement) GetFormNoValidate() (formNoValidate string) {
	js.Rewrite("$<.formNoValidate")
	return formNoValidate
}

func (*HTMLButtonElement) SetFormNoValidate(formNoValidate string) {
	js.Rewrite("$<.formNoValidate = $1", formNoValidate)
}

func (*HTMLButtonElement) GetFormTarget() (formTarget string) {
	js.Rewrite("$<.formTarget")
	return formTarget
}

func (*HTMLButtonElement) SetFormTarget(formTarget string) {
	js.Rewrite("$<.formTarget = $1", formTarget)
}

func (*HTMLButtonElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLButtonElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLButtonElement) GetStatus() (status interface{}) {
	js.Rewrite("$<.status")
	return status
}

func (*HTMLButtonElement) SetStatus(status interface{}) {
	js.Rewrite("$<.status = $1", status)
}

func (*HTMLButtonElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLButtonElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}

func (*HTMLButtonElement) GetValidationMessage() (validationMessage string) {
	js.Rewrite("$<.validationMessage")
	return validationMessage
}

func (*HTMLButtonElement) GetValidity() (validity *ValidityState) {
	js.Rewrite("$<.validity")
	return validity
}

func (*HTMLButtonElement) GetValue() (value string) {
	js.Rewrite("$<.value")
	return value
}

func (*HTMLButtonElement) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}

func (*HTMLButtonElement) GetWillValidate() (willValidate bool) {
	js.Rewrite("$<.willValidate")
	return willValidate
}

type HTMLCanvasElement struct {
	*HTMLElement
}

func (*HTMLCanvasElement) GetContext(contextId string, args interface{}) (i interface{}) {
	js.Rewrite("$<.getContext($1, $2)", contextId, args)
	return i
}

func (*HTMLCanvasElement) MsToBlob() (b *Blob) {
	js.Rewrite("$<.msToBlob()")
	return b
}

func (*HTMLCanvasElement) ToDataURL(kind string, args interface{}) (s string) {
	js.Rewrite("$<.toDataURL($1, $2)", kind, args)
	return s
}

func (*HTMLCanvasElement) GetHeight() (height uint) {
	js.Rewrite("$<.height")
	return height
}

func (*HTMLCanvasElement) SetHeight(height uint) {
	js.Rewrite("$<.height = $1", height)
}

func (*HTMLCanvasElement) GetWidth() (width uint) {
	js.Rewrite("$<.width")
	return width
}

func (*HTMLCanvasElement) SetWidth(width uint) {
	js.Rewrite("$<.width = $1", width)
}

type HTMLCollection struct {
}

func (*HTMLCollection) Item(index uint) (e *Element) {
	js.Rewrite("$<.item($1)", index)
	return e
}

func (*HTMLCollection) NamedItem(name string) (e *Element) {
	js.Rewrite("$<.namedItem($1)", name)
	return e
}

func (*HTMLCollection) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

type HTMLDataElement struct {
	*HTMLElement
}

func (*HTMLDataElement) GetValue() (value string) {
	js.Rewrite("$<.value")
	return value
}

func (*HTMLDataElement) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}

type HTMLDataListElement struct {
	*HTMLElement
}

func (*HTMLDataListElement) GetOptions() (options *HTMLCollection) {
	js.Rewrite("$<.options")
	return options
}

type HTMLDirectoryElement struct {
	*HTMLElement
}

func (*HTMLDirectoryElement) GetCompact() (compact bool) {
	js.Rewrite("$<.compact")
	return compact
}

func (*HTMLDirectoryElement) SetCompact(compact bool) {
	js.Rewrite("$<.compact = $1", compact)
}

type HTMLDivElement struct {
	*HTMLElement
}

func (*HTMLDivElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLDivElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLDivElement) GetNoWrap() (noWrap bool) {
	js.Rewrite("$<.noWrap")
	return noWrap
}

func (*HTMLDivElement) SetNoWrap(noWrap bool) {
	js.Rewrite("$<.noWrap = $1", noWrap)
}

type HTMLDListElement struct {
	*HTMLElement
}

func (*HTMLDListElement) GetCompact() (compact bool) {
	js.Rewrite("$<.compact")
	return compact
}

func (*HTMLDListElement) SetCompact(compact bool) {
	js.Rewrite("$<.compact = $1", compact)
}

type HTMLDocument struct {
	*Document
}

type HTMLElement struct {
	*Element
}

func (*HTMLElement) Blur() {
	js.Rewrite("$<.blur()")
}

func (*HTMLElement) Click() {
	js.Rewrite("$<.click()")
}

func (*HTMLElement) DragDrop() (b bool) {
	js.Rewrite("$<.dragDrop()")
	return b
}

func (*HTMLElement) Focus() {
	js.Rewrite("$<.focus()")
}

func (*HTMLElement) GetElementsByClassName(classNames string) (n *NodeList) {
	js.Rewrite("$<.getElementsByClassName($1)", classNames)
	return n
}

func (*HTMLElement) InsertAdjacentElement(position string, insertedElement *Element) (e *Element) {
	js.Rewrite("$<.insertAdjacentElement($1, $2)", position, insertedElement)
	return e
}

func (*HTMLElement) InsertAdjacentHTML(where string, html string) {
	js.Rewrite("$<.insertAdjacentHTML($1, $2)", where, html)
}

func (*HTMLElement) InsertAdjacentText(where string, text string) {
	js.Rewrite("$<.insertAdjacentText($1, $2)", where, text)
}

func (*HTMLElement) MsGetInputContext() (m *MSInputMethodContext) {
	js.Rewrite("$<.msGetInputContext()")
	return m
}

func (*HTMLElement) ScrollIntoView(top bool) {
	js.Rewrite("$<.scrollIntoView($1)", top)
}

func (*HTMLElement) GetAccessKey() (accessKey string) {
	js.Rewrite("$<.accessKey")
	return accessKey
}

func (*HTMLElement) SetAccessKey(accessKey string) {
	js.Rewrite("$<.accessKey = $1", accessKey)
}

func (*HTMLElement) GetChildren() (children *HTMLCollection) {
	js.Rewrite("$<.children")
	return children
}

func (*HTMLElement) GetContentEditable() (contentEditable string) {
	js.Rewrite("$<.contentEditable")
	return contentEditable
}

func (*HTMLElement) SetContentEditable(contentEditable string) {
	js.Rewrite("$<.contentEditable = $1", contentEditable)
}

func (*HTMLElement) GetDataset() (dataset *DOMStringMap) {
	js.Rewrite("$<.dataset")
	return dataset
}

func (*HTMLElement) GetDir() (dir string) {
	js.Rewrite("$<.dir")
	return dir
}

func (*HTMLElement) SetDir(dir string) {
	js.Rewrite("$<.dir = $1", dir)
}

func (*HTMLElement) GetDraggable() (draggable bool) {
	js.Rewrite("$<.draggable")
	return draggable
}

func (*HTMLElement) SetDraggable(draggable bool) {
	js.Rewrite("$<.draggable = $1", draggable)
}

func (*HTMLElement) GetHidden() (hidden bool) {
	js.Rewrite("$<.hidden")
	return hidden
}

func (*HTMLElement) SetHidden(hidden bool) {
	js.Rewrite("$<.hidden = $1", hidden)
}

func (*HTMLElement) GetHideFocus() (hideFocus bool) {
	js.Rewrite("$<.hideFocus")
	return hideFocus
}

func (*HTMLElement) SetHideFocus(hideFocus bool) {
	js.Rewrite("$<.hideFocus = $1", hideFocus)
}

func (*HTMLElement) GetInnerText() (innerText string) {
	js.Rewrite("$<.innerText")
	return innerText
}

func (*HTMLElement) SetInnerText(innerText string) {
	js.Rewrite("$<.innerText = $1", innerText)
}

func (*HTMLElement) GetIsContentEditable() (isContentEditable bool) {
	js.Rewrite("$<.isContentEditable")
	return isContentEditable
}

func (*HTMLElement) GetLang() (lang string) {
	js.Rewrite("$<.lang")
	return lang
}

func (*HTMLElement) SetLang(lang string) {
	js.Rewrite("$<.lang = $1", lang)
}

func (*HTMLElement) GetOffsetHeight() (offsetHeight int) {
	js.Rewrite("$<.offsetHeight")
	return offsetHeight
}

func (*HTMLElement) GetOffsetLeft() (offsetLeft int) {
	js.Rewrite("$<.offsetLeft")
	return offsetLeft
}

func (*HTMLElement) GetOffsetParent() (offsetParent *Element) {
	js.Rewrite("$<.offsetParent")
	return offsetParent
}

func (*HTMLElement) GetOffsetTop() (offsetTop int) {
	js.Rewrite("$<.offsetTop")
	return offsetTop
}

func (*HTMLElement) GetOffsetWidth() (offsetWidth int) {
	js.Rewrite("$<.offsetWidth")
	return offsetWidth
}

func (*HTMLElement) GetOnabort() (e *Event) {
	js.Rewrite("$<.onabort")
	return e
}

func (*HTMLElement) SetOnabort(e *Event) {
	js.Rewrite("$<.onabort = $1", e)
}

func (*HTMLElement) GetOnactivate() (activate *UIEvent) {
	js.Rewrite("$<.onactivate")
	return activate
}

func (*HTMLElement) SetOnactivate(activate *UIEvent) {
	js.Rewrite("$<.onactivate = $1", activate)
}

func (*HTMLElement) GetOnbeforeactivate() (beforeactivate *UIEvent) {
	js.Rewrite("$<.onbeforeactivate")
	return beforeactivate
}

func (*HTMLElement) SetOnbeforeactivate(beforeactivate *UIEvent) {
	js.Rewrite("$<.onbeforeactivate = $1", beforeactivate)
}

func (*HTMLElement) GetOnbeforecopy() (beforecopy *ClipboardEvent) {
	js.Rewrite("$<.onbeforecopy")
	return beforecopy
}

func (*HTMLElement) SetOnbeforecopy(beforecopy *ClipboardEvent) {
	js.Rewrite("$<.onbeforecopy = $1", beforecopy)
}

func (*HTMLElement) GetOnbeforecut() (beforecut *ClipboardEvent) {
	js.Rewrite("$<.onbeforecut")
	return beforecut
}

func (*HTMLElement) SetOnbeforecut(beforecut *ClipboardEvent) {
	js.Rewrite("$<.onbeforecut = $1", beforecut)
}

func (*HTMLElement) GetOnbeforedeactivate() (beforedeactivate *UIEvent) {
	js.Rewrite("$<.onbeforedeactivate")
	return beforedeactivate
}

func (*HTMLElement) SetOnbeforedeactivate(beforedeactivate *UIEvent) {
	js.Rewrite("$<.onbeforedeactivate = $1", beforedeactivate)
}

func (*HTMLElement) GetOnbeforepaste() (beforepaste *ClipboardEvent) {
	js.Rewrite("$<.onbeforepaste")
	return beforepaste
}

func (*HTMLElement) SetOnbeforepaste(beforepaste *ClipboardEvent) {
	js.Rewrite("$<.onbeforepaste = $1", beforepaste)
}

func (*HTMLElement) GetOnblur() (blur *FocusEvent) {
	js.Rewrite("$<.onblur")
	return blur
}

func (*HTMLElement) SetOnblur(blur *FocusEvent) {
	js.Rewrite("$<.onblur = $1", blur)
}

func (*HTMLElement) GetOncanplay() (e *Event) {
	js.Rewrite("$<.oncanplay")
	return e
}

func (*HTMLElement) SetOncanplay(e *Event) {
	js.Rewrite("$<.oncanplay = $1", e)
}

func (*HTMLElement) GetOncanplaythrough() (e *Event) {
	js.Rewrite("$<.oncanplaythrough")
	return e
}

func (*HTMLElement) SetOncanplaythrough(e *Event) {
	js.Rewrite("$<.oncanplaythrough = $1", e)
}

func (*HTMLElement) GetOnchange() (e *Event) {
	js.Rewrite("$<.onchange")
	return e
}

func (*HTMLElement) SetOnchange(e *Event) {
	js.Rewrite("$<.onchange = $1", e)
}

func (*HTMLElement) GetOnclick() (click *MouseEvent) {
	js.Rewrite("$<.onclick")
	return click
}

func (*HTMLElement) SetOnclick(click *MouseEvent) {
	js.Rewrite("$<.onclick = $1", click)
}

func (*HTMLElement) GetOncontextmenu() (contextmenu *PointerEvent) {
	js.Rewrite("$<.oncontextmenu")
	return contextmenu
}

func (*HTMLElement) SetOncontextmenu(contextmenu *PointerEvent) {
	js.Rewrite("$<.oncontextmenu = $1", contextmenu)
}

func (*HTMLElement) GetOncopy() (cpy *ClipboardEvent) {
	js.Rewrite("$<.oncopy")
	return cpy
}

func (*HTMLElement) SetOncopy(cpy *ClipboardEvent) {
	js.Rewrite("$<.oncopy = $1", cpy)
}

func (*HTMLElement) GetOncuechange() (e *Event) {
	js.Rewrite("$<.oncuechange")
	return e
}

func (*HTMLElement) SetOncuechange(e *Event) {
	js.Rewrite("$<.oncuechange = $1", e)
}

func (*HTMLElement) GetOncut() (cut *ClipboardEvent) {
	js.Rewrite("$<.oncut")
	return cut
}

func (*HTMLElement) SetOncut(cut *ClipboardEvent) {
	js.Rewrite("$<.oncut = $1", cut)
}

func (*HTMLElement) GetOndblclick() (dblclick *MouseEvent) {
	js.Rewrite("$<.ondblclick")
	return dblclick
}

func (*HTMLElement) SetOndblclick(dblclick *MouseEvent) {
	js.Rewrite("$<.ondblclick = $1", dblclick)
}

func (*HTMLElement) GetOndeactivate() (deactivate *UIEvent) {
	js.Rewrite("$<.ondeactivate")
	return deactivate
}

func (*HTMLElement) SetOndeactivate(deactivate *UIEvent) {
	js.Rewrite("$<.ondeactivate = $1", deactivate)
}

func (*HTMLElement) GetOndrag() (drag *DragEvent) {
	js.Rewrite("$<.ondrag")
	return drag
}

func (*HTMLElement) SetOndrag(drag *DragEvent) {
	js.Rewrite("$<.ondrag = $1", drag)
}

func (*HTMLElement) GetOndragend() (dragend *DragEvent) {
	js.Rewrite("$<.ondragend")
	return dragend
}

func (*HTMLElement) SetOndragend(dragend *DragEvent) {
	js.Rewrite("$<.ondragend = $1", dragend)
}

func (*HTMLElement) GetOndragenter() (dragenter *DragEvent) {
	js.Rewrite("$<.ondragenter")
	return dragenter
}

func (*HTMLElement) SetOndragenter(dragenter *DragEvent) {
	js.Rewrite("$<.ondragenter = $1", dragenter)
}

func (*HTMLElement) GetOndragleave() (dragleave *DragEvent) {
	js.Rewrite("$<.ondragleave")
	return dragleave
}

func (*HTMLElement) SetOndragleave(dragleave *DragEvent) {
	js.Rewrite("$<.ondragleave = $1", dragleave)
}

func (*HTMLElement) GetOndragover() (dragover *DragEvent) {
	js.Rewrite("$<.ondragover")
	return dragover
}

func (*HTMLElement) SetOndragover(dragover *DragEvent) {
	js.Rewrite("$<.ondragover = $1", dragover)
}

func (*HTMLElement) GetOndragstart() (dragstart *DragEvent) {
	js.Rewrite("$<.ondragstart")
	return dragstart
}

func (*HTMLElement) SetOndragstart(dragstart *DragEvent) {
	js.Rewrite("$<.ondragstart = $1", dragstart)
}

func (*HTMLElement) GetOndrop() (drop *DragEvent) {
	js.Rewrite("$<.ondrop")
	return drop
}

func (*HTMLElement) SetOndrop(drop *DragEvent) {
	js.Rewrite("$<.ondrop = $1", drop)
}

func (*HTMLElement) GetOndurationchange() (e *Event) {
	js.Rewrite("$<.ondurationchange")
	return e
}

func (*HTMLElement) SetOndurationchange(e *Event) {
	js.Rewrite("$<.ondurationchange = $1", e)
}

func (*HTMLElement) GetOnemptied() (e *Event) {
	js.Rewrite("$<.onemptied")
	return e
}

func (*HTMLElement) SetOnemptied(e *Event) {
	js.Rewrite("$<.onemptied = $1", e)
}

func (*HTMLElement) GetOnended() (e *Event) {
	js.Rewrite("$<.onended")
	return e
}

func (*HTMLElement) SetOnended(e *Event) {
	js.Rewrite("$<.onended = $1", e)
}

func (*HTMLElement) GetOnerror() (e *Event) {
	js.Rewrite("$<.onerror")
	return e
}

func (*HTMLElement) SetOnerror(e *Event) {
	js.Rewrite("$<.onerror = $1", e)
}

func (*HTMLElement) GetOnfocus() (focus *FocusEvent) {
	js.Rewrite("$<.onfocus")
	return focus
}

func (*HTMLElement) SetOnfocus(focus *FocusEvent) {
	js.Rewrite("$<.onfocus = $1", focus)
}

func (*HTMLElement) GetOninput() (e *Event) {
	js.Rewrite("$<.oninput")
	return e
}

func (*HTMLElement) SetOninput(e *Event) {
	js.Rewrite("$<.oninput = $1", e)
}

func (*HTMLElement) GetOninvalid() (e *Event) {
	js.Rewrite("$<.oninvalid")
	return e
}

func (*HTMLElement) SetOninvalid(e *Event) {
	js.Rewrite("$<.oninvalid = $1", e)
}

func (*HTMLElement) GetOnkeydown() (keydown *KeyboardEvent) {
	js.Rewrite("$<.onkeydown")
	return keydown
}

func (*HTMLElement) SetOnkeydown(keydown *KeyboardEvent) {
	js.Rewrite("$<.onkeydown = $1", keydown)
}

func (*HTMLElement) GetOnkeypress() (keypress *KeyboardEvent) {
	js.Rewrite("$<.onkeypress")
	return keypress
}

func (*HTMLElement) SetOnkeypress(keypress *KeyboardEvent) {
	js.Rewrite("$<.onkeypress = $1", keypress)
}

func (*HTMLElement) GetOnkeyup() (keyup *KeyboardEvent) {
	js.Rewrite("$<.onkeyup")
	return keyup
}

func (*HTMLElement) SetOnkeyup(keyup *KeyboardEvent) {
	js.Rewrite("$<.onkeyup = $1", keyup)
}

func (*HTMLElement) GetOnload() (e *Event) {
	js.Rewrite("$<.onload")
	return e
}

func (*HTMLElement) SetOnload(e *Event) {
	js.Rewrite("$<.onload = $1", e)
}

func (*HTMLElement) GetOnloadeddata() (e *Event) {
	js.Rewrite("$<.onloadeddata")
	return e
}

func (*HTMLElement) SetOnloadeddata(e *Event) {
	js.Rewrite("$<.onloadeddata = $1", e)
}

func (*HTMLElement) GetOnloadedmetadata() (e *Event) {
	js.Rewrite("$<.onloadedmetadata")
	return e
}

func (*HTMLElement) SetOnloadedmetadata(e *Event) {
	js.Rewrite("$<.onloadedmetadata = $1", e)
}

func (*HTMLElement) GetOnloadstart() (e *Event) {
	js.Rewrite("$<.onloadstart")
	return e
}

func (*HTMLElement) SetOnloadstart(e *Event) {
	js.Rewrite("$<.onloadstart = $1", e)
}

func (*HTMLElement) GetOnmousedown() (mousedown *MouseEvent) {
	js.Rewrite("$<.onmousedown")
	return mousedown
}

func (*HTMLElement) SetOnmousedown(mousedown *MouseEvent) {
	js.Rewrite("$<.onmousedown = $1", mousedown)
}

func (*HTMLElement) GetOnmouseenter() (mouseenter *MouseEvent) {
	js.Rewrite("$<.onmouseenter")
	return mouseenter
}

func (*HTMLElement) SetOnmouseenter(mouseenter *MouseEvent) {
	js.Rewrite("$<.onmouseenter = $1", mouseenter)
}

func (*HTMLElement) GetOnmouseleave() (mouseleave *MouseEvent) {
	js.Rewrite("$<.onmouseleave")
	return mouseleave
}

func (*HTMLElement) SetOnmouseleave(mouseleave *MouseEvent) {
	js.Rewrite("$<.onmouseleave = $1", mouseleave)
}

func (*HTMLElement) GetOnmousemove() (mousemove *MouseEvent) {
	js.Rewrite("$<.onmousemove")
	return mousemove
}

func (*HTMLElement) SetOnmousemove(mousemove *MouseEvent) {
	js.Rewrite("$<.onmousemove = $1", mousemove)
}

func (*HTMLElement) GetOnmouseout() (mouseout *MouseEvent) {
	js.Rewrite("$<.onmouseout")
	return mouseout
}

func (*HTMLElement) SetOnmouseout(mouseout *MouseEvent) {
	js.Rewrite("$<.onmouseout = $1", mouseout)
}

func (*HTMLElement) GetOnmouseover() (mouseover *MouseEvent) {
	js.Rewrite("$<.onmouseover")
	return mouseover
}

func (*HTMLElement) SetOnmouseover(mouseover *MouseEvent) {
	js.Rewrite("$<.onmouseover = $1", mouseover)
}

func (*HTMLElement) GetOnmouseup() (mouseup *MouseEvent) {
	js.Rewrite("$<.onmouseup")
	return mouseup
}

func (*HTMLElement) SetOnmouseup(mouseup *MouseEvent) {
	js.Rewrite("$<.onmouseup = $1", mouseup)
}

func (*HTMLElement) GetOnmousewheel() (mousewheel *WheelEvent) {
	js.Rewrite("$<.onmousewheel")
	return mousewheel
}

func (*HTMLElement) SetOnmousewheel(mousewheel *WheelEvent) {
	js.Rewrite("$<.onmousewheel = $1", mousewheel)
}

func (*HTMLElement) GetOnmscontentzoom() (MSContentZoom *UIEvent) {
	js.Rewrite("$<.onmscontentzoom")
	return MSContentZoom
}

func (*HTMLElement) SetOnmscontentzoom(MSContentZoom *UIEvent) {
	js.Rewrite("$<.onmscontentzoom = $1", MSContentZoom)
}

func (*HTMLElement) GetOnmsmanipulationstatechanged() (MSManipulationStateChanged *MSManipulationEvent) {
	js.Rewrite("$<.onmsmanipulationstatechanged")
	return MSManipulationStateChanged
}

func (*HTMLElement) SetOnmsmanipulationstatechanged(MSManipulationStateChanged *MSManipulationEvent) {
	js.Rewrite("$<.onmsmanipulationstatechanged = $1", MSManipulationStateChanged)
}

func (*HTMLElement) GetOnpaste() (paste *ClipboardEvent) {
	js.Rewrite("$<.onpaste")
	return paste
}

func (*HTMLElement) SetOnpaste(paste *ClipboardEvent) {
	js.Rewrite("$<.onpaste = $1", paste)
}

func (*HTMLElement) GetOnpause() (e *Event) {
	js.Rewrite("$<.onpause")
	return e
}

func (*HTMLElement) SetOnpause(e *Event) {
	js.Rewrite("$<.onpause = $1", e)
}

func (*HTMLElement) GetOnplay() (e *Event) {
	js.Rewrite("$<.onplay")
	return e
}

func (*HTMLElement) SetOnplay(e *Event) {
	js.Rewrite("$<.onplay = $1", e)
}

func (*HTMLElement) GetOnplaying() (e *Event) {
	js.Rewrite("$<.onplaying")
	return e
}

func (*HTMLElement) SetOnplaying(e *Event) {
	js.Rewrite("$<.onplaying = $1", e)
}

func (*HTMLElement) GetOnprogress() (e *Event) {
	js.Rewrite("$<.onprogress")
	return e
}

func (*HTMLElement) SetOnprogress(e *Event) {
	js.Rewrite("$<.onprogress = $1", e)
}

func (*HTMLElement) GetOnratechange() (e *Event) {
	js.Rewrite("$<.onratechange")
	return e
}

func (*HTMLElement) SetOnratechange(e *Event) {
	js.Rewrite("$<.onratechange = $1", e)
}

func (*HTMLElement) GetOnreset() (e *Event) {
	js.Rewrite("$<.onreset")
	return e
}

func (*HTMLElement) SetOnreset(e *Event) {
	js.Rewrite("$<.onreset = $1", e)
}

func (*HTMLElement) GetOnscroll() (scroll *UIEvent) {
	js.Rewrite("$<.onscroll")
	return scroll
}

func (*HTMLElement) SetOnscroll(scroll *UIEvent) {
	js.Rewrite("$<.onscroll = $1", scroll)
}

func (*HTMLElement) GetOnseeked() (e *Event) {
	js.Rewrite("$<.onseeked")
	return e
}

func (*HTMLElement) SetOnseeked(e *Event) {
	js.Rewrite("$<.onseeked = $1", e)
}

func (*HTMLElement) GetOnseeking() (e *Event) {
	js.Rewrite("$<.onseeking")
	return e
}

func (*HTMLElement) SetOnseeking(e *Event) {
	js.Rewrite("$<.onseeking = $1", e)
}

func (*HTMLElement) GetOnselect() (sel *UIEvent) {
	js.Rewrite("$<.onselect")
	return sel
}

func (*HTMLElement) SetOnselect(sel *UIEvent) {
	js.Rewrite("$<.onselect = $1", sel)
}

func (*HTMLElement) GetOnselectstart() (selectstart *Event) {
	js.Rewrite("$<.onselectstart")
	return selectstart
}

func (*HTMLElement) SetOnselectstart(selectstart *Event) {
	js.Rewrite("$<.onselectstart = $1", selectstart)
}

func (*HTMLElement) GetOnstalled() (e *Event) {
	js.Rewrite("$<.onstalled")
	return e
}

func (*HTMLElement) SetOnstalled(e *Event) {
	js.Rewrite("$<.onstalled = $1", e)
}

func (*HTMLElement) GetOnsubmit() (e *Event) {
	js.Rewrite("$<.onsubmit")
	return e
}

func (*HTMLElement) SetOnsubmit(e *Event) {
	js.Rewrite("$<.onsubmit = $1", e)
}

func (*HTMLElement) GetOnsuspend() (e *Event) {
	js.Rewrite("$<.onsuspend")
	return e
}

func (*HTMLElement) SetOnsuspend(e *Event) {
	js.Rewrite("$<.onsuspend = $1", e)
}

func (*HTMLElement) GetOntimeupdate() (e *Event) {
	js.Rewrite("$<.ontimeupdate")
	return e
}

func (*HTMLElement) SetOntimeupdate(e *Event) {
	js.Rewrite("$<.ontimeupdate = $1", e)
}

func (*HTMLElement) GetOnvolumechange() (e *Event) {
	js.Rewrite("$<.onvolumechange")
	return e
}

func (*HTMLElement) SetOnvolumechange(e *Event) {
	js.Rewrite("$<.onvolumechange = $1", e)
}

func (*HTMLElement) GetOnwaiting() (e *Event) {
	js.Rewrite("$<.onwaiting")
	return e
}

func (*HTMLElement) SetOnwaiting(e *Event) {
	js.Rewrite("$<.onwaiting = $1", e)
}

func (*HTMLElement) GetOuterText() (outerText string) {
	js.Rewrite("$<.outerText")
	return outerText
}

func (*HTMLElement) SetOuterText(outerText string) {
	js.Rewrite("$<.outerText = $1", outerText)
}

func (*HTMLElement) GetSpellcheck() (spellcheck bool) {
	js.Rewrite("$<.spellcheck")
	return spellcheck
}

func (*HTMLElement) SetSpellcheck(spellcheck bool) {
	js.Rewrite("$<.spellcheck = $1", spellcheck)
}

func (*HTMLElement) GetStyle() (style *CSSStyleDeclaration) {
	js.Rewrite("$<.style")
	return style
}

func (*HTMLElement) GetTabIndex() (tabIndex int8) {
	js.Rewrite("$<.tabIndex")
	return tabIndex
}

func (*HTMLElement) SetTabIndex(tabIndex int8) {
	js.Rewrite("$<.tabIndex = $1", tabIndex)
}

func (*HTMLElement) GetTitle() (title string) {
	js.Rewrite("$<.title")
	return title
}

func (*HTMLElement) SetTitle(title string) {
	js.Rewrite("$<.title = $1", title)
}

type HTMLEmbedElement struct {
	*HTMLElement
	*GetSVGDocument
}

func (*HTMLEmbedElement) GetHeight() (height string) {
	js.Rewrite("$<.height")
	return height
}

func (*HTMLEmbedElement) SetHeight(height string) {
	js.Rewrite("$<.height = $1", height)
}

func (*HTMLEmbedElement) GetHidden() (hidden string) {
	js.Rewrite("$<.hidden")
	return hidden
}

func (*HTMLEmbedElement) SetHidden(hidden string) {
	js.Rewrite("$<.hidden = $1", hidden)
}

func (*HTMLEmbedElement) GetMsPlayToDisabled() (msPlayToDisabled bool) {
	js.Rewrite("$<.msPlayToDisabled")
	return msPlayToDisabled
}

func (*HTMLEmbedElement) SetMsPlayToDisabled(msPlayToDisabled bool) {
	js.Rewrite("$<.msPlayToDisabled = $1", msPlayToDisabled)
}

func (*HTMLEmbedElement) GetMsPlayToPreferredSourceURI() (msPlayToPreferredSourceUri string) {
	js.Rewrite("$<.msPlayToPreferredSourceUri")
	return msPlayToPreferredSourceUri
}

func (*HTMLEmbedElement) SetMsPlayToPreferredSourceURI(msPlayToPreferredSourceUri string) {
	js.Rewrite("$<.msPlayToPreferredSourceUri = $1", msPlayToPreferredSourceUri)
}

func (*HTMLEmbedElement) GetMsPlayToPrimary() (msPlayToPrimary bool) {
	js.Rewrite("$<.msPlayToPrimary")
	return msPlayToPrimary
}

func (*HTMLEmbedElement) SetMsPlayToPrimary(msPlayToPrimary bool) {
	js.Rewrite("$<.msPlayToPrimary = $1", msPlayToPrimary)
}

func (*HTMLEmbedElement) GetMsPlayToSource() (msPlayToSource interface{}) {
	js.Rewrite("$<.msPlayToSource")
	return msPlayToSource
}

func (*HTMLEmbedElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLEmbedElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLEmbedElement) GetPalette() (palette string) {
	js.Rewrite("$<.palette")
	return palette
}

func (*HTMLEmbedElement) GetPluginspage() (pluginspage string) {
	js.Rewrite("$<.pluginspage")
	return pluginspage
}

func (*HTMLEmbedElement) GetReadyState() (readyState string) {
	js.Rewrite("$<.readyState")
	return readyState
}

func (*HTMLEmbedElement) GetSrc() (src string) {
	js.Rewrite("$<.src")
	return src
}

func (*HTMLEmbedElement) SetSrc(src string) {
	js.Rewrite("$<.src = $1", src)
}

func (*HTMLEmbedElement) GetUnits() (units string) {
	js.Rewrite("$<.units")
	return units
}

func (*HTMLEmbedElement) SetUnits(units string) {
	js.Rewrite("$<.units = $1", units)
}

func (*HTMLEmbedElement) GetWidth() (width string) {
	js.Rewrite("$<.width")
	return width
}

func (*HTMLEmbedElement) SetWidth(width string) {
	js.Rewrite("$<.width = $1", width)
}

type HTMLFieldSetElement struct {
	*HTMLElement
}

func (*HTMLFieldSetElement) CheckValidity() (b bool) {
	js.Rewrite("$<.checkValidity()")
	return b
}

func (*HTMLFieldSetElement) SetCustomValidity(err string) {
	js.Rewrite("$<.setCustomValidity($1)", err)
}

func (*HTMLFieldSetElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLFieldSetElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLFieldSetElement) GetDisabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

func (*HTMLFieldSetElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

func (*HTMLFieldSetElement) GetForm() (form *HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

func (*HTMLFieldSetElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLFieldSetElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLFieldSetElement) GetValidationMessage() (validationMessage string) {
	js.Rewrite("$<.validationMessage")
	return validationMessage
}

func (*HTMLFieldSetElement) GetValidity() (validity *ValidityState) {
	js.Rewrite("$<.validity")
	return validity
}

func (*HTMLFieldSetElement) GetWillValidate() (willValidate bool) {
	js.Rewrite("$<.willValidate")
	return willValidate
}

type HTMLFontElement struct {
	*HTMLElement
	*DOML2DeprecatedColorProperty
	*DOML2DeprecatedSizeProperty
}

func (*HTMLFontElement) GetFace() (face string) {
	js.Rewrite("$<.face")
	return face
}

func (*HTMLFontElement) SetFace(face string) {
	js.Rewrite("$<.face = $1", face)
}

type HTMLFormControlsCollection struct {
	*HTMLCollection
}

func (*HTMLFormControlsCollection) NamedItem(name string) (i interface{}) {
	js.Rewrite("$<.namedItem($1)", name)
	return i
}

type HTMLFormElement struct {
	*HTMLElement
}

func (*HTMLFormElement) CheckValidity() (b bool) {
	js.Rewrite("$<.checkValidity()")
	return b
}

func (*HTMLFormElement) Item(name interface{}, index interface{}) (i interface{}) {
	js.Rewrite("$<.item($1, $2)", name, index)
	return i
}

func (*HTMLFormElement) NamedItem(name string) (i interface{}) {
	js.Rewrite("$<.namedItem($1)", name)
	return i
}

func (*HTMLFormElement) Reset() {
	js.Rewrite("$<.reset()")
}

func (*HTMLFormElement) Submit() {
	js.Rewrite("$<.submit()")
}

func (*HTMLFormElement) GetAcceptCharset() (acceptCharset string) {
	js.Rewrite("$<.acceptCharset")
	return acceptCharset
}

func (*HTMLFormElement) SetAcceptCharset(acceptCharset string) {
	js.Rewrite("$<.acceptCharset = $1", acceptCharset)
}

func (*HTMLFormElement) GetAction() (action string) {
	js.Rewrite("$<.action")
	return action
}

func (*HTMLFormElement) SetAction(action string) {
	js.Rewrite("$<.action = $1", action)
}

func (*HTMLFormElement) GetAutocomplete() (autocomplete string) {
	js.Rewrite("$<.autocomplete")
	return autocomplete
}

func (*HTMLFormElement) SetAutocomplete(autocomplete string) {
	js.Rewrite("$<.autocomplete = $1", autocomplete)
}

func (*HTMLFormElement) GetElements() (elements *HTMLFormControlsCollection) {
	js.Rewrite("$<.elements")
	return elements
}

func (*HTMLFormElement) GetEncoding() (encoding string) {
	js.Rewrite("$<.encoding")
	return encoding
}

func (*HTMLFormElement) SetEncoding(encoding string) {
	js.Rewrite("$<.encoding = $1", encoding)
}

func (*HTMLFormElement) GetEnctype() (enctype string) {
	js.Rewrite("$<.enctype")
	return enctype
}

func (*HTMLFormElement) SetEnctype(enctype string) {
	js.Rewrite("$<.enctype = $1", enctype)
}

func (*HTMLFormElement) GetLength() (length int) {
	js.Rewrite("$<.length")
	return length
}

func (*HTMLFormElement) GetMethod() (method string) {
	js.Rewrite("$<.method")
	return method
}

func (*HTMLFormElement) SetMethod(method string) {
	js.Rewrite("$<.method = $1", method)
}

func (*HTMLFormElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLFormElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLFormElement) GetNoValidate() (noValidate bool) {
	js.Rewrite("$<.noValidate")
	return noValidate
}

func (*HTMLFormElement) SetNoValidate(noValidate bool) {
	js.Rewrite("$<.noValidate = $1", noValidate)
}

func (*HTMLFormElement) GetTarget() (target string) {
	js.Rewrite("$<.target")
	return target
}

func (*HTMLFormElement) SetTarget(target string) {
	js.Rewrite("$<.target = $1", target)
}

type HTMLFrameElement struct {
	*HTMLElement
	*GetSVGDocument
}

func (*HTMLFrameElement) GetBorder() (border string) {
	js.Rewrite("$<.border")
	return border
}

func (*HTMLFrameElement) SetBorder(border string) {
	js.Rewrite("$<.border = $1", border)
}

func (*HTMLFrameElement) GetBorderColor() (borderColor interface{}) {
	js.Rewrite("$<.borderColor")
	return borderColor
}

func (*HTMLFrameElement) SetBorderColor(borderColor interface{}) {
	js.Rewrite("$<.borderColor = $1", borderColor)
}

func (*HTMLFrameElement) GetContentDocument() (contentDocument *Document) {
	js.Rewrite("$<.contentDocument")
	return contentDocument
}

func (*HTMLFrameElement) GetContentWindow() (contentWindow *Window) {
	js.Rewrite("$<.contentWindow")
	return contentWindow
}

func (*HTMLFrameElement) GetFrameBorder() (frameBorder string) {
	js.Rewrite("$<.frameBorder")
	return frameBorder
}

func (*HTMLFrameElement) SetFrameBorder(frameBorder string) {
	js.Rewrite("$<.frameBorder = $1", frameBorder)
}

func (*HTMLFrameElement) GetFrameSpacing() (frameSpacing interface{}) {
	js.Rewrite("$<.frameSpacing")
	return frameSpacing
}

func (*HTMLFrameElement) SetFrameSpacing(frameSpacing interface{}) {
	js.Rewrite("$<.frameSpacing = $1", frameSpacing)
}

func (*HTMLFrameElement) GetHeight() (height interface{}) {
	js.Rewrite("$<.height")
	return height
}

func (*HTMLFrameElement) SetHeight(height interface{}) {
	js.Rewrite("$<.height = $1", height)
}

func (*HTMLFrameElement) GetLongDesc() (longDesc string) {
	js.Rewrite("$<.longDesc")
	return longDesc
}

func (*HTMLFrameElement) SetLongDesc(longDesc string) {
	js.Rewrite("$<.longDesc = $1", longDesc)
}

func (*HTMLFrameElement) GetMarginHeight() (marginHeight string) {
	js.Rewrite("$<.marginHeight")
	return marginHeight
}

func (*HTMLFrameElement) SetMarginHeight(marginHeight string) {
	js.Rewrite("$<.marginHeight = $1", marginHeight)
}

func (*HTMLFrameElement) GetMarginWidth() (marginWidth string) {
	js.Rewrite("$<.marginWidth")
	return marginWidth
}

func (*HTMLFrameElement) SetMarginWidth(marginWidth string) {
	js.Rewrite("$<.marginWidth = $1", marginWidth)
}

func (*HTMLFrameElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLFrameElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLFrameElement) GetNoResize() (noResize bool) {
	js.Rewrite("$<.noResize")
	return noResize
}

func (*HTMLFrameElement) SetNoResize(noResize bool) {
	js.Rewrite("$<.noResize = $1", noResize)
}

func (*HTMLFrameElement) GetOnload() (load *Event) {
	js.Rewrite("$<.onload")
	return load
}

func (*HTMLFrameElement) SetOnload(load *Event) {
	js.Rewrite("$<.onload = $1", load)
}

func (*HTMLFrameElement) GetScrolling() (scrolling string) {
	js.Rewrite("$<.scrolling")
	return scrolling
}

func (*HTMLFrameElement) SetScrolling(scrolling string) {
	js.Rewrite("$<.scrolling = $1", scrolling)
}

func (*HTMLFrameElement) GetSrc() (src string) {
	js.Rewrite("$<.src")
	return src
}

func (*HTMLFrameElement) SetSrc(src string) {
	js.Rewrite("$<.src = $1", src)
}

func (*HTMLFrameElement) GetWidth() (width interface{}) {
	js.Rewrite("$<.width")
	return width
}

func (*HTMLFrameElement) SetWidth(width interface{}) {
	js.Rewrite("$<.width = $1", width)
}

type HTMLFrameSetElement struct {
	*HTMLElement
}

func (*HTMLFrameSetElement) GetBorder() (border string) {
	js.Rewrite("$<.border")
	return border
}

func (*HTMLFrameSetElement) SetBorder(border string) {
	js.Rewrite("$<.border = $1", border)
}

func (*HTMLFrameSetElement) GetBorderColor() (borderColor interface{}) {
	js.Rewrite("$<.borderColor")
	return borderColor
}

func (*HTMLFrameSetElement) SetBorderColor(borderColor interface{}) {
	js.Rewrite("$<.borderColor = $1", borderColor)
}

func (*HTMLFrameSetElement) GetCols() (cols string) {
	js.Rewrite("$<.cols")
	return cols
}

func (*HTMLFrameSetElement) SetCols(cols string) {
	js.Rewrite("$<.cols = $1", cols)
}

func (*HTMLFrameSetElement) GetFrameBorder() (frameBorder string) {
	js.Rewrite("$<.frameBorder")
	return frameBorder
}

func (*HTMLFrameSetElement) SetFrameBorder(frameBorder string) {
	js.Rewrite("$<.frameBorder = $1", frameBorder)
}

func (*HTMLFrameSetElement) GetFrameSpacing() (frameSpacing interface{}) {
	js.Rewrite("$<.frameSpacing")
	return frameSpacing
}

func (*HTMLFrameSetElement) SetFrameSpacing(frameSpacing interface{}) {
	js.Rewrite("$<.frameSpacing = $1", frameSpacing)
}

func (*HTMLFrameSetElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLFrameSetElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLFrameSetElement) GetOnafterprint() (e *Event) {
	js.Rewrite("$<.onafterprint")
	return e
}

func (*HTMLFrameSetElement) SetOnafterprint(e *Event) {
	js.Rewrite("$<.onafterprint = $1", e)
}

func (*HTMLFrameSetElement) GetOnbeforeprint() (e *Event) {
	js.Rewrite("$<.onbeforeprint")
	return e
}

func (*HTMLFrameSetElement) SetOnbeforeprint(e *Event) {
	js.Rewrite("$<.onbeforeprint = $1", e)
}

func (*HTMLFrameSetElement) GetOnbeforeunload() (e *Event) {
	js.Rewrite("$<.onbeforeunload")
	return e
}

func (*HTMLFrameSetElement) SetOnbeforeunload(e *Event) {
	js.Rewrite("$<.onbeforeunload = $1", e)
}

func (*HTMLFrameSetElement) GetOnblur() (blur *FocusEvent) {
	js.Rewrite("$<.onblur")
	return blur
}

func (*HTMLFrameSetElement) SetOnblur(blur *FocusEvent) {
	js.Rewrite("$<.onblur = $1", blur)
}

func (*HTMLFrameSetElement) GetOnerror() (e *Event) {
	js.Rewrite("$<.onerror")
	return e
}

func (*HTMLFrameSetElement) SetOnerror(e *Event) {
	js.Rewrite("$<.onerror = $1", e)
}

func (*HTMLFrameSetElement) GetOnfocus() (focus *FocusEvent) {
	js.Rewrite("$<.onfocus")
	return focus
}

func (*HTMLFrameSetElement) SetOnfocus(focus *FocusEvent) {
	js.Rewrite("$<.onfocus = $1", focus)
}

func (*HTMLFrameSetElement) GetOnhashchange() (e *Event) {
	js.Rewrite("$<.onhashchange")
	return e
}

func (*HTMLFrameSetElement) SetOnhashchange(e *Event) {
	js.Rewrite("$<.onhashchange = $1", e)
}

func (*HTMLFrameSetElement) GetOnload() (e *Event) {
	js.Rewrite("$<.onload")
	return e
}

func (*HTMLFrameSetElement) SetOnload(e *Event) {
	js.Rewrite("$<.onload = $1", e)
}

func (*HTMLFrameSetElement) GetOnmessage() (e *Event) {
	js.Rewrite("$<.onmessage")
	return e
}

func (*HTMLFrameSetElement) SetOnmessage(e *Event) {
	js.Rewrite("$<.onmessage = $1", e)
}

func (*HTMLFrameSetElement) GetOnoffline() (e *Event) {
	js.Rewrite("$<.onoffline")
	return e
}

func (*HTMLFrameSetElement) SetOnoffline(e *Event) {
	js.Rewrite("$<.onoffline = $1", e)
}

func (*HTMLFrameSetElement) GetOnonline() (e *Event) {
	js.Rewrite("$<.ononline")
	return e
}

func (*HTMLFrameSetElement) SetOnonline(e *Event) {
	js.Rewrite("$<.ononline = $1", e)
}

func (*HTMLFrameSetElement) GetOnorientationchange() (e *Event) {
	js.Rewrite("$<.onorientationchange")
	return e
}

func (*HTMLFrameSetElement) SetOnorientationchange(e *Event) {
	js.Rewrite("$<.onorientationchange = $1", e)
}

func (*HTMLFrameSetElement) GetOnpagehide() (e *Event) {
	js.Rewrite("$<.onpagehide")
	return e
}

func (*HTMLFrameSetElement) SetOnpagehide(e *Event) {
	js.Rewrite("$<.onpagehide = $1", e)
}

func (*HTMLFrameSetElement) GetOnpageshow() (e *Event) {
	js.Rewrite("$<.onpageshow")
	return e
}

func (*HTMLFrameSetElement) SetOnpageshow(e *Event) {
	js.Rewrite("$<.onpageshow = $1", e)
}

func (*HTMLFrameSetElement) GetOnpopstate() (e *Event) {
	js.Rewrite("$<.onpopstate")
	return e
}

func (*HTMLFrameSetElement) SetOnpopstate(e *Event) {
	js.Rewrite("$<.onpopstate = $1", e)
}

func (*HTMLFrameSetElement) GetOnresize() (e *Event) {
	js.Rewrite("$<.onresize")
	return e
}

func (*HTMLFrameSetElement) SetOnresize(e *Event) {
	js.Rewrite("$<.onresize = $1", e)
}

func (*HTMLFrameSetElement) GetOnscroll() (scroll *UIEvent) {
	js.Rewrite("$<.onscroll")
	return scroll
}

func (*HTMLFrameSetElement) SetOnscroll(scroll *UIEvent) {
	js.Rewrite("$<.onscroll = $1", scroll)
}

func (*HTMLFrameSetElement) GetOnstorage() (e *Event) {
	js.Rewrite("$<.onstorage")
	return e
}

func (*HTMLFrameSetElement) SetOnstorage(e *Event) {
	js.Rewrite("$<.onstorage = $1", e)
}

func (*HTMLFrameSetElement) GetOnunload() (e *Event) {
	js.Rewrite("$<.onunload")
	return e
}

func (*HTMLFrameSetElement) SetOnunload(e *Event) {
	js.Rewrite("$<.onunload = $1", e)
}

func (*HTMLFrameSetElement) GetRows() (rows string) {
	js.Rewrite("$<.rows")
	return rows
}

func (*HTMLFrameSetElement) SetRows(rows string) {
	js.Rewrite("$<.rows = $1", rows)
}

type HTMLHeadElement struct {
	*HTMLElement
}

func (*HTMLHeadElement) GetProfile() (profile string) {
	js.Rewrite("$<.profile")
	return profile
}

func (*HTMLHeadElement) SetProfile(profile string) {
	js.Rewrite("$<.profile = $1", profile)
}

type HTMLHeadingElement struct {
	*HTMLElement
}

func (*HTMLHeadingElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLHeadingElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

type HTMLHRElement struct {
	*HTMLElement
	*DOML2DeprecatedColorProperty
	*DOML2DeprecatedSizeProperty
}

func (*HTMLHRElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLHRElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLHRElement) GetNoShade() (noShade bool) {
	js.Rewrite("$<.noShade")
	return noShade
}

func (*HTMLHRElement) SetNoShade(noShade bool) {
	js.Rewrite("$<.noShade = $1", noShade)
}

func (*HTMLHRElement) GetWidth() (width int) {
	js.Rewrite("$<.width")
	return width
}

func (*HTMLHRElement) SetWidth(width int) {
	js.Rewrite("$<.width = $1", width)
}

type HTMLHtmlElement struct {
	*HTMLElement
}

func (*HTMLHtmlElement) GetVersion() (version string) {
	js.Rewrite("$<.version")
	return version
}

func (*HTMLHtmlElement) SetVersion(version string) {
	js.Rewrite("$<.version = $1", version)
}

type HTMLIFrameElement struct {
	*HTMLElement
	*GetSVGDocument
}

func (*HTMLIFrameElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLIFrameElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLIFrameElement) GetAllowFullscreen() (allowFullscreen bool) {
	js.Rewrite("$<.allowFullscreen")
	return allowFullscreen
}

func (*HTMLIFrameElement) SetAllowFullscreen(allowFullscreen bool) {
	js.Rewrite("$<.allowFullscreen = $1", allowFullscreen)
}

func (*HTMLIFrameElement) GetAllowPaymentRequest() (allowPaymentRequest bool) {
	js.Rewrite("$<.allowPaymentRequest")
	return allowPaymentRequest
}

func (*HTMLIFrameElement) SetAllowPaymentRequest(allowPaymentRequest bool) {
	js.Rewrite("$<.allowPaymentRequest = $1", allowPaymentRequest)
}

func (*HTMLIFrameElement) GetBorder() (border string) {
	js.Rewrite("$<.border")
	return border
}

func (*HTMLIFrameElement) SetBorder(border string) {
	js.Rewrite("$<.border = $1", border)
}

func (*HTMLIFrameElement) GetContentDocument() (contentDocument *Document) {
	js.Rewrite("$<.contentDocument")
	return contentDocument
}

func (*HTMLIFrameElement) GetContentWindow() (contentWindow *Window) {
	js.Rewrite("$<.contentWindow")
	return contentWindow
}

func (*HTMLIFrameElement) GetFrameBorder() (frameBorder string) {
	js.Rewrite("$<.frameBorder")
	return frameBorder
}

func (*HTMLIFrameElement) SetFrameBorder(frameBorder string) {
	js.Rewrite("$<.frameBorder = $1", frameBorder)
}

func (*HTMLIFrameElement) GetFrameSpacing() (frameSpacing interface{}) {
	js.Rewrite("$<.frameSpacing")
	return frameSpacing
}

func (*HTMLIFrameElement) SetFrameSpacing(frameSpacing interface{}) {
	js.Rewrite("$<.frameSpacing = $1", frameSpacing)
}

func (*HTMLIFrameElement) GetHeight() (height string) {
	js.Rewrite("$<.height")
	return height
}

func (*HTMLIFrameElement) SetHeight(height string) {
	js.Rewrite("$<.height = $1", height)
}

func (*HTMLIFrameElement) GetHspace() (hspace int) {
	js.Rewrite("$<.hspace")
	return hspace
}

func (*HTMLIFrameElement) SetHspace(hspace int) {
	js.Rewrite("$<.hspace = $1", hspace)
}

func (*HTMLIFrameElement) GetLongDesc() (longDesc string) {
	js.Rewrite("$<.longDesc")
	return longDesc
}

func (*HTMLIFrameElement) SetLongDesc(longDesc string) {
	js.Rewrite("$<.longDesc = $1", longDesc)
}

func (*HTMLIFrameElement) GetMarginHeight() (marginHeight string) {
	js.Rewrite("$<.marginHeight")
	return marginHeight
}

func (*HTMLIFrameElement) SetMarginHeight(marginHeight string) {
	js.Rewrite("$<.marginHeight = $1", marginHeight)
}

func (*HTMLIFrameElement) GetMarginWidth() (marginWidth string) {
	js.Rewrite("$<.marginWidth")
	return marginWidth
}

func (*HTMLIFrameElement) SetMarginWidth(marginWidth string) {
	js.Rewrite("$<.marginWidth = $1", marginWidth)
}

func (*HTMLIFrameElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLIFrameElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLIFrameElement) GetNoResize() (noResize bool) {
	js.Rewrite("$<.noResize")
	return noResize
}

func (*HTMLIFrameElement) SetNoResize(noResize bool) {
	js.Rewrite("$<.noResize = $1", noResize)
}

func (*HTMLIFrameElement) GetOnload() (load *Event) {
	js.Rewrite("$<.onload")
	return load
}

func (*HTMLIFrameElement) SetOnload(load *Event) {
	js.Rewrite("$<.onload = $1", load)
}

func (*HTMLIFrameElement) GetSandbox() (sandbox *DOMSettableTokenList) {
	js.Rewrite("$<.sandbox")
	return sandbox
}

func (*HTMLIFrameElement) GetScrolling() (scrolling string) {
	js.Rewrite("$<.scrolling")
	return scrolling
}

func (*HTMLIFrameElement) SetScrolling(scrolling string) {
	js.Rewrite("$<.scrolling = $1", scrolling)
}

func (*HTMLIFrameElement) GetSrc() (src string) {
	js.Rewrite("$<.src")
	return src
}

func (*HTMLIFrameElement) SetSrc(src string) {
	js.Rewrite("$<.src = $1", src)
}

func (*HTMLIFrameElement) GetVspace() (vspace int) {
	js.Rewrite("$<.vspace")
	return vspace
}

func (*HTMLIFrameElement) SetVspace(vspace int) {
	js.Rewrite("$<.vspace = $1", vspace)
}

func (*HTMLIFrameElement) GetWidth() (width string) {
	js.Rewrite("$<.width")
	return width
}

func (*HTMLIFrameElement) SetWidth(width string) {
	js.Rewrite("$<.width = $1", width)
}

type HTMLImageElement struct {
	*HTMLElement
}

func (*HTMLImageElement) MsGetAsCastingSource() (i interface{}) {
	js.Rewrite("$<.msGetAsCastingSource()")
	return i
}

func (*HTMLImageElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLImageElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLImageElement) GetAlt() (alt string) {
	js.Rewrite("$<.alt")
	return alt
}

func (*HTMLImageElement) SetAlt(alt string) {
	js.Rewrite("$<.alt = $1", alt)
}

func (*HTMLImageElement) GetBorder() (border string) {
	js.Rewrite("$<.border")
	return border
}

func (*HTMLImageElement) SetBorder(border string) {
	js.Rewrite("$<.border = $1", border)
}

func (*HTMLImageElement) GetComplete() (complete bool) {
	js.Rewrite("$<.complete")
	return complete
}

func (*HTMLImageElement) GetCrossOrigin() (crossOrigin string) {
	js.Rewrite("$<.crossOrigin")
	return crossOrigin
}

func (*HTMLImageElement) SetCrossOrigin(crossOrigin string) {
	js.Rewrite("$<.crossOrigin = $1", crossOrigin)
}

func (*HTMLImageElement) GetCurrentSrc() (currentSrc string) {
	js.Rewrite("$<.currentSrc")
	return currentSrc
}

func (*HTMLImageElement) GetHeight() (height uint) {
	js.Rewrite("$<.height")
	return height
}

func (*HTMLImageElement) SetHeight(height uint) {
	js.Rewrite("$<.height = $1", height)
}

func (*HTMLImageElement) GetHspace() (hspace int) {
	js.Rewrite("$<.hspace")
	return hspace
}

func (*HTMLImageElement) SetHspace(hspace int) {
	js.Rewrite("$<.hspace = $1", hspace)
}

func (*HTMLImageElement) GetIsMap() (isMap bool) {
	js.Rewrite("$<.isMap")
	return isMap
}

func (*HTMLImageElement) SetIsMap(isMap bool) {
	js.Rewrite("$<.isMap = $1", isMap)
}

func (*HTMLImageElement) GetLongDesc() (longDesc string) {
	js.Rewrite("$<.longDesc")
	return longDesc
}

func (*HTMLImageElement) SetLongDesc(longDesc string) {
	js.Rewrite("$<.longDesc = $1", longDesc)
}

func (*HTMLImageElement) GetLowsrc() (lowsrc string) {
	js.Rewrite("$<.lowsrc")
	return lowsrc
}

func (*HTMLImageElement) SetLowsrc(lowsrc string) {
	js.Rewrite("$<.lowsrc = $1", lowsrc)
}

func (*HTMLImageElement) GetMsPlayToDisabled() (msPlayToDisabled bool) {
	js.Rewrite("$<.msPlayToDisabled")
	return msPlayToDisabled
}

func (*HTMLImageElement) SetMsPlayToDisabled(msPlayToDisabled bool) {
	js.Rewrite("$<.msPlayToDisabled = $1", msPlayToDisabled)
}

func (*HTMLImageElement) GetMsPlayToPreferredSourceURI() (msPlayToPreferredSourceUri string) {
	js.Rewrite("$<.msPlayToPreferredSourceUri")
	return msPlayToPreferredSourceUri
}

func (*HTMLImageElement) SetMsPlayToPreferredSourceURI(msPlayToPreferredSourceUri string) {
	js.Rewrite("$<.msPlayToPreferredSourceUri = $1", msPlayToPreferredSourceUri)
}

func (*HTMLImageElement) GetMsPlayToPrimary() (msPlayToPrimary bool) {
	js.Rewrite("$<.msPlayToPrimary")
	return msPlayToPrimary
}

func (*HTMLImageElement) SetMsPlayToPrimary(msPlayToPrimary bool) {
	js.Rewrite("$<.msPlayToPrimary = $1", msPlayToPrimary)
}

func (*HTMLImageElement) GetMsPlayToSource() (msPlayToSource interface{}) {
	js.Rewrite("$<.msPlayToSource")
	return msPlayToSource
}

func (*HTMLImageElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLImageElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLImageElement) GetNaturalHeight() (naturalHeight uint) {
	js.Rewrite("$<.naturalHeight")
	return naturalHeight
}

func (*HTMLImageElement) GetNaturalWidth() (naturalWidth uint) {
	js.Rewrite("$<.naturalWidth")
	return naturalWidth
}

func (*HTMLImageElement) GetSizes() (sizes string) {
	js.Rewrite("$<.sizes")
	return sizes
}

func (*HTMLImageElement) SetSizes(sizes string) {
	js.Rewrite("$<.sizes = $1", sizes)
}

func (*HTMLImageElement) GetSrc() (src string) {
	js.Rewrite("$<.src")
	return src
}

func (*HTMLImageElement) SetSrc(src string) {
	js.Rewrite("$<.src = $1", src)
}

func (*HTMLImageElement) GetSrcset() (srcset string) {
	js.Rewrite("$<.srcset")
	return srcset
}

func (*HTMLImageElement) SetSrcset(srcset string) {
	js.Rewrite("$<.srcset = $1", srcset)
}

func (*HTMLImageElement) GetUseMap() (useMap string) {
	js.Rewrite("$<.useMap")
	return useMap
}

func (*HTMLImageElement) SetUseMap(useMap string) {
	js.Rewrite("$<.useMap = $1", useMap)
}

func (*HTMLImageElement) GetVspace() (vspace int) {
	js.Rewrite("$<.vspace")
	return vspace
}

func (*HTMLImageElement) SetVspace(vspace int) {
	js.Rewrite("$<.vspace = $1", vspace)
}

func (*HTMLImageElement) GetWidth() (width uint) {
	js.Rewrite("$<.width")
	return width
}

func (*HTMLImageElement) SetWidth(width uint) {
	js.Rewrite("$<.width = $1", width)
}

func (*HTMLImageElement) GetX() (x int) {
	js.Rewrite("$<.x")
	return x
}

func (*HTMLImageElement) GetY() (y int) {
	js.Rewrite("$<.y")
	return y
}

type HTMLInputElement struct {
	*HTMLElement
}

func (*HTMLInputElement) CheckValidity() (b bool) {
	js.Rewrite("$<.checkValidity()")
	return b
}

func (*HTMLInputElement) Select() {
	js.Rewrite("$<.select()")
}

func (*HTMLInputElement) SetCustomValidity(err string) {
	js.Rewrite("$<.setCustomValidity($1)", err)
}

func (*HTMLInputElement) SetSelectionRange(start uint, end uint, direction string) {
	js.Rewrite("$<.setSelectionRange($1, $2, $3)", start, end, direction)
}

func (*HTMLInputElement) StepDown(n int) {
	js.Rewrite("$<.stepDown($1)", n)
}

func (*HTMLInputElement) StepUp(n int) {
	js.Rewrite("$<.stepUp($1)", n)
}

func (*HTMLInputElement) GetAccept() (accept string) {
	js.Rewrite("$<.accept")
	return accept
}

func (*HTMLInputElement) SetAccept(accept string) {
	js.Rewrite("$<.accept = $1", accept)
}

func (*HTMLInputElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLInputElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLInputElement) GetAlt() (alt string) {
	js.Rewrite("$<.alt")
	return alt
}

func (*HTMLInputElement) SetAlt(alt string) {
	js.Rewrite("$<.alt = $1", alt)
}

func (*HTMLInputElement) GetAutocomplete() (autocomplete string) {
	js.Rewrite("$<.autocomplete")
	return autocomplete
}

func (*HTMLInputElement) SetAutocomplete(autocomplete string) {
	js.Rewrite("$<.autocomplete = $1", autocomplete)
}

func (*HTMLInputElement) GetAutofocus() (autofocus bool) {
	js.Rewrite("$<.autofocus")
	return autofocus
}

func (*HTMLInputElement) SetAutofocus(autofocus bool) {
	js.Rewrite("$<.autofocus = $1", autofocus)
}

func (*HTMLInputElement) GetBorder() (border string) {
	js.Rewrite("$<.border")
	return border
}

func (*HTMLInputElement) SetBorder(border string) {
	js.Rewrite("$<.border = $1", border)
}

func (*HTMLInputElement) GetChecked() (checked bool) {
	js.Rewrite("$<.checked")
	return checked
}

func (*HTMLInputElement) SetChecked(checked bool) {
	js.Rewrite("$<.checked = $1", checked)
}

func (*HTMLInputElement) GetComplete() (complete bool) {
	js.Rewrite("$<.complete")
	return complete
}

func (*HTMLInputElement) GetDefaultChecked() (defaultChecked bool) {
	js.Rewrite("$<.defaultChecked")
	return defaultChecked
}

func (*HTMLInputElement) SetDefaultChecked(defaultChecked bool) {
	js.Rewrite("$<.defaultChecked = $1", defaultChecked)
}

func (*HTMLInputElement) GetDefaultValue() (defaultValue string) {
	js.Rewrite("$<.defaultValue")
	return defaultValue
}

func (*HTMLInputElement) SetDefaultValue(defaultValue string) {
	js.Rewrite("$<.defaultValue = $1", defaultValue)
}

func (*HTMLInputElement) GetDisabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

func (*HTMLInputElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

func (*HTMLInputElement) GetFiles() (files *FileList) {
	js.Rewrite("$<.files")
	return files
}

func (*HTMLInputElement) GetForm() (form *HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

func (*HTMLInputElement) GetFormAction() (formAction string) {
	js.Rewrite("$<.formAction")
	return formAction
}

func (*HTMLInputElement) SetFormAction(formAction string) {
	js.Rewrite("$<.formAction = $1", formAction)
}

func (*HTMLInputElement) GetFormEnctype() (formEnctype string) {
	js.Rewrite("$<.formEnctype")
	return formEnctype
}

func (*HTMLInputElement) SetFormEnctype(formEnctype string) {
	js.Rewrite("$<.formEnctype = $1", formEnctype)
}

func (*HTMLInputElement) GetFormMethod() (formMethod string) {
	js.Rewrite("$<.formMethod")
	return formMethod
}

func (*HTMLInputElement) SetFormMethod(formMethod string) {
	js.Rewrite("$<.formMethod = $1", formMethod)
}

func (*HTMLInputElement) GetFormNoValidate() (formNoValidate string) {
	js.Rewrite("$<.formNoValidate")
	return formNoValidate
}

func (*HTMLInputElement) SetFormNoValidate(formNoValidate string) {
	js.Rewrite("$<.formNoValidate = $1", formNoValidate)
}

func (*HTMLInputElement) GetFormTarget() (formTarget string) {
	js.Rewrite("$<.formTarget")
	return formTarget
}

func (*HTMLInputElement) SetFormTarget(formTarget string) {
	js.Rewrite("$<.formTarget = $1", formTarget)
}

func (*HTMLInputElement) GetHeight() (height string) {
	js.Rewrite("$<.height")
	return height
}

func (*HTMLInputElement) SetHeight(height string) {
	js.Rewrite("$<.height = $1", height)
}

func (*HTMLInputElement) GetHspace() (hspace int) {
	js.Rewrite("$<.hspace")
	return hspace
}

func (*HTMLInputElement) SetHspace(hspace int) {
	js.Rewrite("$<.hspace = $1", hspace)
}

func (*HTMLInputElement) GetIndeterminate() (indeterminate bool) {
	js.Rewrite("$<.indeterminate")
	return indeterminate
}

func (*HTMLInputElement) SetIndeterminate(indeterminate bool) {
	js.Rewrite("$<.indeterminate = $1", indeterminate)
}

func (*HTMLInputElement) GetList() (list *HTMLElement) {
	js.Rewrite("$<.list")
	return list
}

func (*HTMLInputElement) GetMax() (max string) {
	js.Rewrite("$<.max")
	return max
}

func (*HTMLInputElement) SetMax(max string) {
	js.Rewrite("$<.max = $1", max)
}

func (*HTMLInputElement) GetMaxLength() (maxLength int) {
	js.Rewrite("$<.maxLength")
	return maxLength
}

func (*HTMLInputElement) SetMaxLength(maxLength int) {
	js.Rewrite("$<.maxLength = $1", maxLength)
}

func (*HTMLInputElement) GetMin() (min string) {
	js.Rewrite("$<.min")
	return min
}

func (*HTMLInputElement) SetMin(min string) {
	js.Rewrite("$<.min = $1", min)
}

func (*HTMLInputElement) GetMultiple() (multiple bool) {
	js.Rewrite("$<.multiple")
	return multiple
}

func (*HTMLInputElement) SetMultiple(multiple bool) {
	js.Rewrite("$<.multiple = $1", multiple)
}

func (*HTMLInputElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLInputElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLInputElement) GetPattern() (pattern string) {
	js.Rewrite("$<.pattern")
	return pattern
}

func (*HTMLInputElement) SetPattern(pattern string) {
	js.Rewrite("$<.pattern = $1", pattern)
}

func (*HTMLInputElement) GetPlaceholder() (placeholder string) {
	js.Rewrite("$<.placeholder")
	return placeholder
}

func (*HTMLInputElement) SetPlaceholder(placeholder string) {
	js.Rewrite("$<.placeholder = $1", placeholder)
}

func (*HTMLInputElement) GetReadOnly() (readOnly bool) {
	js.Rewrite("$<.readOnly")
	return readOnly
}

func (*HTMLInputElement) SetReadOnly(readOnly bool) {
	js.Rewrite("$<.readOnly = $1", readOnly)
}

func (*HTMLInputElement) GetRequired() (required bool) {
	js.Rewrite("$<.required")
	return required
}

func (*HTMLInputElement) SetRequired(required bool) {
	js.Rewrite("$<.required = $1", required)
}

func (*HTMLInputElement) GetSelectionDirection() (selectionDirection string) {
	js.Rewrite("$<.selectionDirection")
	return selectionDirection
}

func (*HTMLInputElement) SetSelectionDirection(selectionDirection string) {
	js.Rewrite("$<.selectionDirection = $1", selectionDirection)
}

func (*HTMLInputElement) GetSelectionEnd() (selectionEnd uint) {
	js.Rewrite("$<.selectionEnd")
	return selectionEnd
}

func (*HTMLInputElement) SetSelectionEnd(selectionEnd uint) {
	js.Rewrite("$<.selectionEnd = $1", selectionEnd)
}

func (*HTMLInputElement) GetSelectionStart() (selectionStart uint) {
	js.Rewrite("$<.selectionStart")
	return selectionStart
}

func (*HTMLInputElement) SetSelectionStart(selectionStart uint) {
	js.Rewrite("$<.selectionStart = $1", selectionStart)
}

func (*HTMLInputElement) GetSize() (size uint) {
	js.Rewrite("$<.size")
	return size
}

func (*HTMLInputElement) SetSize(size uint) {
	js.Rewrite("$<.size = $1", size)
}

func (*HTMLInputElement) GetSrc() (src string) {
	js.Rewrite("$<.src")
	return src
}

func (*HTMLInputElement) SetSrc(src string) {
	js.Rewrite("$<.src = $1", src)
}

func (*HTMLInputElement) GetStatus() (status bool) {
	js.Rewrite("$<.status")
	return status
}

func (*HTMLInputElement) SetStatus(status bool) {
	js.Rewrite("$<.status = $1", status)
}

func (*HTMLInputElement) GetStep() (step string) {
	js.Rewrite("$<.step")
	return step
}

func (*HTMLInputElement) SetStep(step string) {
	js.Rewrite("$<.step = $1", step)
}

func (*HTMLInputElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLInputElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}

func (*HTMLInputElement) GetUseMap() (useMap string) {
	js.Rewrite("$<.useMap")
	return useMap
}

func (*HTMLInputElement) SetUseMap(useMap string) {
	js.Rewrite("$<.useMap = $1", useMap)
}

func (*HTMLInputElement) GetValidationMessage() (validationMessage string) {
	js.Rewrite("$<.validationMessage")
	return validationMessage
}

func (*HTMLInputElement) GetValidity() (validity *ValidityState) {
	js.Rewrite("$<.validity")
	return validity
}

func (*HTMLInputElement) GetValue() (value string) {
	js.Rewrite("$<.value")
	return value
}

func (*HTMLInputElement) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}

func (*HTMLInputElement) GetValueAsDate() (valueAsDate time.Time) {
	js.Rewrite("$<.valueAsDate")
	return valueAsDate
}

func (*HTMLInputElement) SetValueAsDate(valueAsDate time.Time) {
	js.Rewrite("$<.valueAsDate = $1", valueAsDate)
}

func (*HTMLInputElement) GetValueAsNumber() (valueAsNumber float32) {
	js.Rewrite("$<.valueAsNumber")
	return valueAsNumber
}

func (*HTMLInputElement) SetValueAsNumber(valueAsNumber float32) {
	js.Rewrite("$<.valueAsNumber = $1", valueAsNumber)
}

func (*HTMLInputElement) GetVspace() (vspace int) {
	js.Rewrite("$<.vspace")
	return vspace
}

func (*HTMLInputElement) SetVspace(vspace int) {
	js.Rewrite("$<.vspace = $1", vspace)
}

func (*HTMLInputElement) GetWebkitdirectory() (webkitdirectory bool) {
	js.Rewrite("$<.webkitdirectory")
	return webkitdirectory
}

func (*HTMLInputElement) SetWebkitdirectory(webkitdirectory bool) {
	js.Rewrite("$<.webkitdirectory = $1", webkitdirectory)
}

func (*HTMLInputElement) GetWidth() (width string) {
	js.Rewrite("$<.width")
	return width
}

func (*HTMLInputElement) SetWidth(width string) {
	js.Rewrite("$<.width = $1", width)
}

func (*HTMLInputElement) GetWillValidate() (willValidate bool) {
	js.Rewrite("$<.willValidate")
	return willValidate
}

type HTMLLabelElement struct {
	*HTMLElement
}

func (*HTMLLabelElement) GetForm() (form *HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

func (*HTMLLabelElement) GetHTMLFor() (htmlFor string) {
	js.Rewrite("$<.htmlFor")
	return htmlFor
}

func (*HTMLLabelElement) SetHTMLFor(htmlFor string) {
	js.Rewrite("$<.htmlFor = $1", htmlFor)
}

type HTMLLegendElement struct {
	*HTMLElement
}

func (*HTMLLegendElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLLegendElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLLegendElement) GetForm() (form *HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

type HTMLLIElement struct {
	*HTMLElement
}

func (*HTMLLIElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLLIElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}

func (*HTMLLIElement) GetValue() (value int) {
	js.Rewrite("$<.value")
	return value
}

func (*HTMLLIElement) SetValue(value int) {
	js.Rewrite("$<.value = $1", value)
}

type HTMLLinkElement struct {
	*HTMLElement
	*LinkStyle
}

func (*HTMLLinkElement) GetCharset() (charset string) {
	js.Rewrite("$<.charset")
	return charset
}

func (*HTMLLinkElement) SetCharset(charset string) {
	js.Rewrite("$<.charset = $1", charset)
}

func (*HTMLLinkElement) GetDisabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

func (*HTMLLinkElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

func (*HTMLLinkElement) GetHref() (href string) {
	js.Rewrite("$<.href")
	return href
}

func (*HTMLLinkElement) SetHref(href string) {
	js.Rewrite("$<.href = $1", href)
}

func (*HTMLLinkElement) GetHreflang() (hreflang string) {
	js.Rewrite("$<.hreflang")
	return hreflang
}

func (*HTMLLinkElement) SetHreflang(hreflang string) {
	js.Rewrite("$<.hreflang = $1", hreflang)
}

func (*HTMLLinkElement) GetMedia() (media string) {
	js.Rewrite("$<.media")
	return media
}

func (*HTMLLinkElement) SetMedia(media string) {
	js.Rewrite("$<.media = $1", media)
}

func (*HTMLLinkElement) GetRel() (rel string) {
	js.Rewrite("$<.rel")
	return rel
}

func (*HTMLLinkElement) SetRel(rel string) {
	js.Rewrite("$<.rel = $1", rel)
}

func (*HTMLLinkElement) GetRev() (rev string) {
	js.Rewrite("$<.rev")
	return rev
}

func (*HTMLLinkElement) SetRev(rev string) {
	js.Rewrite("$<.rev = $1", rev)
}

func (*HTMLLinkElement) GetTarget() (target string) {
	js.Rewrite("$<.target")
	return target
}

func (*HTMLLinkElement) SetTarget(target string) {
	js.Rewrite("$<.target = $1", target)
}

func (*HTMLLinkElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLLinkElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}

type HTMLMapElement struct {
	*HTMLElement
}

func (*HTMLMapElement) GetAreas() (areas *HTMLAreasCollection) {
	js.Rewrite("$<.areas")
	return areas
}

func (*HTMLMapElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLMapElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

type HTMLMarqueeElement struct {
	*HTMLElement
}

func (*HTMLMarqueeElement) Start() {
	js.Rewrite("$<.start()")
}

func (*HTMLMarqueeElement) Stop() {
	js.Rewrite("$<.stop()")
}

func (*HTMLMarqueeElement) GetBehavior() (behavior string) {
	js.Rewrite("$<.behavior")
	return behavior
}

func (*HTMLMarqueeElement) SetBehavior(behavior string) {
	js.Rewrite("$<.behavior = $1", behavior)
}

func (*HTMLMarqueeElement) GetBgColor() (bgColor interface{}) {
	js.Rewrite("$<.bgColor")
	return bgColor
}

func (*HTMLMarqueeElement) SetBgColor(bgColor interface{}) {
	js.Rewrite("$<.bgColor = $1", bgColor)
}

func (*HTMLMarqueeElement) GetDirection() (direction string) {
	js.Rewrite("$<.direction")
	return direction
}

func (*HTMLMarqueeElement) SetDirection(direction string) {
	js.Rewrite("$<.direction = $1", direction)
}

func (*HTMLMarqueeElement) GetHeight() (height string) {
	js.Rewrite("$<.height")
	return height
}

func (*HTMLMarqueeElement) SetHeight(height string) {
	js.Rewrite("$<.height = $1", height)
}

func (*HTMLMarqueeElement) GetHspace() (hspace int) {
	js.Rewrite("$<.hspace")
	return hspace
}

func (*HTMLMarqueeElement) SetHspace(hspace int) {
	js.Rewrite("$<.hspace = $1", hspace)
}

func (*HTMLMarqueeElement) GetLoop() (loop int) {
	js.Rewrite("$<.loop")
	return loop
}

func (*HTMLMarqueeElement) SetLoop(loop int) {
	js.Rewrite("$<.loop = $1", loop)
}

func (*HTMLMarqueeElement) GetOnbounce() (bounce *Event) {
	js.Rewrite("$<.onbounce")
	return bounce
}

func (*HTMLMarqueeElement) SetOnbounce(bounce *Event) {
	js.Rewrite("$<.onbounce = $1", bounce)
}

func (*HTMLMarqueeElement) GetOnfinish() (finish *Event) {
	js.Rewrite("$<.onfinish")
	return finish
}

func (*HTMLMarqueeElement) SetOnfinish(finish *Event) {
	js.Rewrite("$<.onfinish = $1", finish)
}

func (*HTMLMarqueeElement) GetOnstart() (start *Event) {
	js.Rewrite("$<.onstart")
	return start
}

func (*HTMLMarqueeElement) SetOnstart(start *Event) {
	js.Rewrite("$<.onstart = $1", start)
}

func (*HTMLMarqueeElement) GetScrollAmount() (scrollAmount uint) {
	js.Rewrite("$<.scrollAmount")
	return scrollAmount
}

func (*HTMLMarqueeElement) SetScrollAmount(scrollAmount uint) {
	js.Rewrite("$<.scrollAmount = $1", scrollAmount)
}

func (*HTMLMarqueeElement) GetScrollDelay() (scrollDelay uint) {
	js.Rewrite("$<.scrollDelay")
	return scrollDelay
}

func (*HTMLMarqueeElement) SetScrollDelay(scrollDelay uint) {
	js.Rewrite("$<.scrollDelay = $1", scrollDelay)
}

func (*HTMLMarqueeElement) GetTrueSpeed() (trueSpeed bool) {
	js.Rewrite("$<.trueSpeed")
	return trueSpeed
}

func (*HTMLMarqueeElement) SetTrueSpeed(trueSpeed bool) {
	js.Rewrite("$<.trueSpeed = $1", trueSpeed)
}

func (*HTMLMarqueeElement) GetVspace() (vspace int) {
	js.Rewrite("$<.vspace")
	return vspace
}

func (*HTMLMarqueeElement) SetVspace(vspace int) {
	js.Rewrite("$<.vspace = $1", vspace)
}

func (*HTMLMarqueeElement) GetWidth() (width string) {
	js.Rewrite("$<.width")
	return width
}

func (*HTMLMarqueeElement) SetWidth(width string) {
	js.Rewrite("$<.width = $1", width)
}

type HTMLMediaElement struct {
	*HTMLElement
}

func (*HTMLMediaElement) AddTextTrack(kind string, label string, language string) (t *TextTrack) {
	js.Rewrite("$<.addTextTrack($1, $2, $3)", kind, label, language)
	return t
}

func (*HTMLMediaElement) CanPlayType(kind string) (s string) {
	js.Rewrite("$<.canPlayType($1)", kind)
	return s
}

func (*HTMLMediaElement) Load() {
	js.Rewrite("$<.load()")
}

func (*HTMLMediaElement) MsClearEffects() {
	js.Rewrite("$<.msClearEffects()")
}

func (*HTMLMediaElement) MsGetAsCastingSource() (i interface{}) {
	js.Rewrite("$<.msGetAsCastingSource()")
	return i
}

func (*HTMLMediaElement) MsInsertAudioEffect(activatableClassId string, effectRequired bool, config interface{}) {
	js.Rewrite("$<.msInsertAudioEffect($1, $2, $3)", activatableClassId, effectRequired, config)
}

func (*HTMLMediaElement) MsSetMediaKeys(mediaKeys *MSMediaKeys) {
	js.Rewrite("$<.msSetMediaKeys($1)", mediaKeys)
}

func (*HTMLMediaElement) MsSetMediaProtectionManager(mediaProtectionManager interface{}) {
	js.Rewrite("$<.msSetMediaProtectionManager($1)", mediaProtectionManager)
}

func (*HTMLMediaElement) Pause() {
	js.Rewrite("$<.pause()")
}

func (*HTMLMediaElement) Play() {
	js.Rewrite("$<.play()")
}

func (*HTMLMediaElement) SetMediaKeys(mediaKeys *MediaKeys) {
	js.Rewrite("await $<.setMediaKeys($1)", mediaKeys)
}

func (*HTMLMediaElement) GetAudioTracks() (audioTracks *AudioTrackList) {
	js.Rewrite("$<.audioTracks")
	return audioTracks
}

func (*HTMLMediaElement) GetAutoplay() (autoplay bool) {
	js.Rewrite("$<.autoplay")
	return autoplay
}

func (*HTMLMediaElement) SetAutoplay(autoplay bool) {
	js.Rewrite("$<.autoplay = $1", autoplay)
}

func (*HTMLMediaElement) GetBuffered() (buffered *TimeRanges) {
	js.Rewrite("$<.buffered")
	return buffered
}

func (*HTMLMediaElement) GetControls() (controls bool) {
	js.Rewrite("$<.controls")
	return controls
}

func (*HTMLMediaElement) SetControls(controls bool) {
	js.Rewrite("$<.controls = $1", controls)
}

func (*HTMLMediaElement) GetCrossOrigin() (crossOrigin string) {
	js.Rewrite("$<.crossOrigin")
	return crossOrigin
}

func (*HTMLMediaElement) SetCrossOrigin(crossOrigin string) {
	js.Rewrite("$<.crossOrigin = $1", crossOrigin)
}

func (*HTMLMediaElement) GetCurrentSrc() (currentSrc string) {
	js.Rewrite("$<.currentSrc")
	return currentSrc
}

func (*HTMLMediaElement) GetCurrentTime() (currentTime float32) {
	js.Rewrite("$<.currentTime")
	return currentTime
}

func (*HTMLMediaElement) SetCurrentTime(currentTime float32) {
	js.Rewrite("$<.currentTime = $1", currentTime)
}

func (*HTMLMediaElement) GetDefaultMuted() (defaultMuted bool) {
	js.Rewrite("$<.defaultMuted")
	return defaultMuted
}

func (*HTMLMediaElement) SetDefaultMuted(defaultMuted bool) {
	js.Rewrite("$<.defaultMuted = $1", defaultMuted)
}

func (*HTMLMediaElement) GetDefaultPlaybackRate() (defaultPlaybackRate float32) {
	js.Rewrite("$<.defaultPlaybackRate")
	return defaultPlaybackRate
}

func (*HTMLMediaElement) SetDefaultPlaybackRate(defaultPlaybackRate float32) {
	js.Rewrite("$<.defaultPlaybackRate = $1", defaultPlaybackRate)
}

func (*HTMLMediaElement) GetDuration() (duration float32) {
	js.Rewrite("$<.duration")
	return duration
}

func (*HTMLMediaElement) GetEnded() (ended bool) {
	js.Rewrite("$<.ended")
	return ended
}

func (*HTMLMediaElement) GetError() (err *MediaError) {
	js.Rewrite("$<.error")
	return err
}

func (*HTMLMediaElement) GetLoop() (loop bool) {
	js.Rewrite("$<.loop")
	return loop
}

func (*HTMLMediaElement) SetLoop(loop bool) {
	js.Rewrite("$<.loop = $1", loop)
}

func (*HTMLMediaElement) GetMediaKeys() (mediaKeys *MediaKeys) {
	js.Rewrite("$<.mediaKeys")
	return mediaKeys
}

func (*HTMLMediaElement) GetMsAudioCategory() (msAudioCategory string) {
	js.Rewrite("$<.msAudioCategory")
	return msAudioCategory
}

func (*HTMLMediaElement) SetMsAudioCategory(msAudioCategory string) {
	js.Rewrite("$<.msAudioCategory = $1", msAudioCategory)
}

func (*HTMLMediaElement) GetMsAudioDeviceType() (msAudioDeviceType string) {
	js.Rewrite("$<.msAudioDeviceType")
	return msAudioDeviceType
}

func (*HTMLMediaElement) SetMsAudioDeviceType(msAudioDeviceType string) {
	js.Rewrite("$<.msAudioDeviceType = $1", msAudioDeviceType)
}

func (*HTMLMediaElement) GetMsGraphicsTrustStatus() (msGraphicsTrustStatus *MSGraphicsTrust) {
	js.Rewrite("$<.msGraphicsTrustStatus")
	return msGraphicsTrustStatus
}

func (*HTMLMediaElement) GetMsKeys() (msKeys *MSMediaKeys) {
	js.Rewrite("$<.msKeys")
	return msKeys
}

func (*HTMLMediaElement) GetMsPlayToDisabled() (msPlayToDisabled bool) {
	js.Rewrite("$<.msPlayToDisabled")
	return msPlayToDisabled
}

func (*HTMLMediaElement) SetMsPlayToDisabled(msPlayToDisabled bool) {
	js.Rewrite("$<.msPlayToDisabled = $1", msPlayToDisabled)
}

func (*HTMLMediaElement) GetMsPlayToPreferredSourceURI() (msPlayToPreferredSourceUri string) {
	js.Rewrite("$<.msPlayToPreferredSourceUri")
	return msPlayToPreferredSourceUri
}

func (*HTMLMediaElement) SetMsPlayToPreferredSourceURI(msPlayToPreferredSourceUri string) {
	js.Rewrite("$<.msPlayToPreferredSourceUri = $1", msPlayToPreferredSourceUri)
}

func (*HTMLMediaElement) GetMsPlayToPrimary() (msPlayToPrimary bool) {
	js.Rewrite("$<.msPlayToPrimary")
	return msPlayToPrimary
}

func (*HTMLMediaElement) SetMsPlayToPrimary(msPlayToPrimary bool) {
	js.Rewrite("$<.msPlayToPrimary = $1", msPlayToPrimary)
}

func (*HTMLMediaElement) GetMsPlayToSource() (msPlayToSource interface{}) {
	js.Rewrite("$<.msPlayToSource")
	return msPlayToSource
}

func (*HTMLMediaElement) GetMsRealTime() (msRealTime bool) {
	js.Rewrite("$<.msRealTime")
	return msRealTime
}

func (*HTMLMediaElement) SetMsRealTime(msRealTime bool) {
	js.Rewrite("$<.msRealTime = $1", msRealTime)
}

func (*HTMLMediaElement) GetMuted() (muted bool) {
	js.Rewrite("$<.muted")
	return muted
}

func (*HTMLMediaElement) SetMuted(muted bool) {
	js.Rewrite("$<.muted = $1", muted)
}

func (*HTMLMediaElement) GetNetworkState() (networkState uint8) {
	js.Rewrite("$<.networkState")
	return networkState
}

func (*HTMLMediaElement) GetOnencrypted() (encrypted *MediaEncryptedEvent) {
	js.Rewrite("$<.onencrypted")
	return encrypted
}

func (*HTMLMediaElement) SetOnencrypted(encrypted *MediaEncryptedEvent) {
	js.Rewrite("$<.onencrypted = $1", encrypted)
}

func (*HTMLMediaElement) GetOnmsneedkey() (msneedkey *MSMediaKeyNeededEvent) {
	js.Rewrite("$<.onmsneedkey")
	return msneedkey
}

func (*HTMLMediaElement) SetOnmsneedkey(msneedkey *MSMediaKeyNeededEvent) {
	js.Rewrite("$<.onmsneedkey = $1", msneedkey)
}

func (*HTMLMediaElement) GetPaused() (paused bool) {
	js.Rewrite("$<.paused")
	return paused
}

func (*HTMLMediaElement) GetPlaybackRate() (playbackRate float32) {
	js.Rewrite("$<.playbackRate")
	return playbackRate
}

func (*HTMLMediaElement) SetPlaybackRate(playbackRate float32) {
	js.Rewrite("$<.playbackRate = $1", playbackRate)
}

func (*HTMLMediaElement) GetPlayed() (played *TimeRanges) {
	js.Rewrite("$<.played")
	return played
}

func (*HTMLMediaElement) GetPreload() (preload string) {
	js.Rewrite("$<.preload")
	return preload
}

func (*HTMLMediaElement) SetPreload(preload string) {
	js.Rewrite("$<.preload = $1", preload)
}

func (*HTMLMediaElement) GetReadyState() (readyState interface{}) {
	js.Rewrite("$<.readyState")
	return readyState
}

func (*HTMLMediaElement) GetSeekable() (seekable *TimeRanges) {
	js.Rewrite("$<.seekable")
	return seekable
}

func (*HTMLMediaElement) GetSeeking() (seeking bool) {
	js.Rewrite("$<.seeking")
	return seeking
}

func (*HTMLMediaElement) GetSrc() (src string) {
	js.Rewrite("$<.src")
	return src
}

func (*HTMLMediaElement) SetSrc(src string) {
	js.Rewrite("$<.src = $1", src)
}

func (*HTMLMediaElement) GetSrcObject() (srcObject *MediaStream) {
	js.Rewrite("$<.srcObject")
	return srcObject
}

func (*HTMLMediaElement) SetSrcObject(srcObject *MediaStream) {
	js.Rewrite("$<.srcObject = $1", srcObject)
}

func (*HTMLMediaElement) GetTextTracks() (textTracks *TextTrackList) {
	js.Rewrite("$<.textTracks")
	return textTracks
}

func (*HTMLMediaElement) GetVideoTracks() (videoTracks *VideoTrackList) {
	js.Rewrite("$<.videoTracks")
	return videoTracks
}

func (*HTMLMediaElement) GetVolume() (volume float32) {
	js.Rewrite("$<.volume")
	return volume
}

func (*HTMLMediaElement) SetVolume(volume float32) {
	js.Rewrite("$<.volume = $1", volume)
}

type HTMLMenuElement struct {
	*HTMLElement
}

func (*HTMLMenuElement) GetCompact() (compact bool) {
	js.Rewrite("$<.compact")
	return compact
}

func (*HTMLMenuElement) SetCompact(compact bool) {
	js.Rewrite("$<.compact = $1", compact)
}

func (*HTMLMenuElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLMenuElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}

type HTMLMetaElement struct {
	*HTMLElement
}

func (*HTMLMetaElement) GetCharset() (charset string) {
	js.Rewrite("$<.charset")
	return charset
}

func (*HTMLMetaElement) SetCharset(charset string) {
	js.Rewrite("$<.charset = $1", charset)
}

func (*HTMLMetaElement) GetContent() (content string) {
	js.Rewrite("$<.content")
	return content
}

func (*HTMLMetaElement) SetContent(content string) {
	js.Rewrite("$<.content = $1", content)
}

func (*HTMLMetaElement) GetHTTPEquiv() (httpEquiv string) {
	js.Rewrite("$<.httpEquiv")
	return httpEquiv
}

func (*HTMLMetaElement) SetHTTPEquiv(httpEquiv string) {
	js.Rewrite("$<.httpEquiv = $1", httpEquiv)
}

func (*HTMLMetaElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLMetaElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLMetaElement) GetScheme() (scheme string) {
	js.Rewrite("$<.scheme")
	return scheme
}

func (*HTMLMetaElement) SetScheme(scheme string) {
	js.Rewrite("$<.scheme = $1", scheme)
}

func (*HTMLMetaElement) GetURL() (url string) {
	js.Rewrite("$<.url")
	return url
}

func (*HTMLMetaElement) SetURL(url string) {
	js.Rewrite("$<.url = $1", url)
}

type HTMLMeterElement struct {
	*HTMLElement
}

func (*HTMLMeterElement) GetHigh() (high float32) {
	js.Rewrite("$<.high")
	return high
}

func (*HTMLMeterElement) SetHigh(high float32) {
	js.Rewrite("$<.high = $1", high)
}

func (*HTMLMeterElement) GetLow() (low float32) {
	js.Rewrite("$<.low")
	return low
}

func (*HTMLMeterElement) SetLow(low float32) {
	js.Rewrite("$<.low = $1", low)
}

func (*HTMLMeterElement) GetMax() (max float32) {
	js.Rewrite("$<.max")
	return max
}

func (*HTMLMeterElement) SetMax(max float32) {
	js.Rewrite("$<.max = $1", max)
}

func (*HTMLMeterElement) GetMin() (min float32) {
	js.Rewrite("$<.min")
	return min
}

func (*HTMLMeterElement) SetMin(min float32) {
	js.Rewrite("$<.min = $1", min)
}

func (*HTMLMeterElement) GetOptimum() (optimum float32) {
	js.Rewrite("$<.optimum")
	return optimum
}

func (*HTMLMeterElement) SetOptimum(optimum float32) {
	js.Rewrite("$<.optimum = $1", optimum)
}

func (*HTMLMeterElement) GetValue() (value float32) {
	js.Rewrite("$<.value")
	return value
}

func (*HTMLMeterElement) SetValue(value float32) {
	js.Rewrite("$<.value = $1", value)
}

type HTMLModElement struct {
	*HTMLElement
}

func (*HTMLModElement) GetCite() (cite string) {
	js.Rewrite("$<.cite")
	return cite
}

func (*HTMLModElement) SetCite(cite string) {
	js.Rewrite("$<.cite = $1", cite)
}

func (*HTMLModElement) GetDateTime() (dateTime string) {
	js.Rewrite("$<.dateTime")
	return dateTime
}

func (*HTMLModElement) SetDateTime(dateTime string) {
	js.Rewrite("$<.dateTime = $1", dateTime)
}

type HTMLObjectElement struct {
	*HTMLElement
	*GetSVGDocument
}

func (*HTMLObjectElement) CheckValidity() (b bool) {
	js.Rewrite("$<.checkValidity()")
	return b
}

func (*HTMLObjectElement) SetCustomValidity(err string) {
	js.Rewrite("$<.setCustomValidity($1)", err)
}

func (*HTMLObjectElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLObjectElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLObjectElement) GetAlt() (alt string) {
	js.Rewrite("$<.alt")
	return alt
}

func (*HTMLObjectElement) SetAlt(alt string) {
	js.Rewrite("$<.alt = $1", alt)
}

func (*HTMLObjectElement) GetAltHTML() (altHtml string) {
	js.Rewrite("$<.altHtml")
	return altHtml
}

func (*HTMLObjectElement) SetAltHTML(altHtml string) {
	js.Rewrite("$<.altHtml = $1", altHtml)
}

func (*HTMLObjectElement) GetArchive() (archive string) {
	js.Rewrite("$<.archive")
	return archive
}

func (*HTMLObjectElement) SetArchive(archive string) {
	js.Rewrite("$<.archive = $1", archive)
}

func (*HTMLObjectElement) GetBaseHref() (BaseHref string) {
	js.Rewrite("$<.BaseHref")
	return BaseHref
}

func (*HTMLObjectElement) GetBorder() (border string) {
	js.Rewrite("$<.border")
	return border
}

func (*HTMLObjectElement) SetBorder(border string) {
	js.Rewrite("$<.border = $1", border)
}

func (*HTMLObjectElement) GetCode() (code string) {
	js.Rewrite("$<.code")
	return code
}

func (*HTMLObjectElement) SetCode(code string) {
	js.Rewrite("$<.code = $1", code)
}

func (*HTMLObjectElement) GetCodeBase() (codeBase string) {
	js.Rewrite("$<.codeBase")
	return codeBase
}

func (*HTMLObjectElement) SetCodeBase(codeBase string) {
	js.Rewrite("$<.codeBase = $1", codeBase)
}

func (*HTMLObjectElement) GetCodeType() (codeType string) {
	js.Rewrite("$<.codeType")
	return codeType
}

func (*HTMLObjectElement) SetCodeType(codeType string) {
	js.Rewrite("$<.codeType = $1", codeType)
}

func (*HTMLObjectElement) GetContentDocument() (contentDocument *Document) {
	js.Rewrite("$<.contentDocument")
	return contentDocument
}

func (*HTMLObjectElement) GetData() (data string) {
	js.Rewrite("$<.data")
	return data
}

func (*HTMLObjectElement) SetData(data string) {
	js.Rewrite("$<.data = $1", data)
}

func (*HTMLObjectElement) GetDeclare() (declare bool) {
	js.Rewrite("$<.declare")
	return declare
}

func (*HTMLObjectElement) SetDeclare(declare bool) {
	js.Rewrite("$<.declare = $1", declare)
}

func (*HTMLObjectElement) GetForm() (form *HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

func (*HTMLObjectElement) GetHeight() (height string) {
	js.Rewrite("$<.height")
	return height
}

func (*HTMLObjectElement) SetHeight(height string) {
	js.Rewrite("$<.height = $1", height)
}

func (*HTMLObjectElement) GetHspace() (hspace int) {
	js.Rewrite("$<.hspace")
	return hspace
}

func (*HTMLObjectElement) SetHspace(hspace int) {
	js.Rewrite("$<.hspace = $1", hspace)
}

func (*HTMLObjectElement) GetMsPlayToDisabled() (msPlayToDisabled bool) {
	js.Rewrite("$<.msPlayToDisabled")
	return msPlayToDisabled
}

func (*HTMLObjectElement) SetMsPlayToDisabled(msPlayToDisabled bool) {
	js.Rewrite("$<.msPlayToDisabled = $1", msPlayToDisabled)
}

func (*HTMLObjectElement) GetMsPlayToPreferredSourceURI() (msPlayToPreferredSourceUri string) {
	js.Rewrite("$<.msPlayToPreferredSourceUri")
	return msPlayToPreferredSourceUri
}

func (*HTMLObjectElement) SetMsPlayToPreferredSourceURI(msPlayToPreferredSourceUri string) {
	js.Rewrite("$<.msPlayToPreferredSourceUri = $1", msPlayToPreferredSourceUri)
}

func (*HTMLObjectElement) GetMsPlayToPrimary() (msPlayToPrimary bool) {
	js.Rewrite("$<.msPlayToPrimary")
	return msPlayToPrimary
}

func (*HTMLObjectElement) SetMsPlayToPrimary(msPlayToPrimary bool) {
	js.Rewrite("$<.msPlayToPrimary = $1", msPlayToPrimary)
}

func (*HTMLObjectElement) GetMsPlayToSource() (msPlayToSource interface{}) {
	js.Rewrite("$<.msPlayToSource")
	return msPlayToSource
}

func (*HTMLObjectElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLObjectElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLObjectElement) GetReadyState() (readyState uint8) {
	js.Rewrite("$<.readyState")
	return readyState
}

func (*HTMLObjectElement) GetStandby() (standby string) {
	js.Rewrite("$<.standby")
	return standby
}

func (*HTMLObjectElement) SetStandby(standby string) {
	js.Rewrite("$<.standby = $1", standby)
}

func (*HTMLObjectElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLObjectElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}

func (*HTMLObjectElement) GetUseMap() (useMap string) {
	js.Rewrite("$<.useMap")
	return useMap
}

func (*HTMLObjectElement) SetUseMap(useMap string) {
	js.Rewrite("$<.useMap = $1", useMap)
}

func (*HTMLObjectElement) GetValidationMessage() (validationMessage string) {
	js.Rewrite("$<.validationMessage")
	return validationMessage
}

func (*HTMLObjectElement) GetValidity() (validity *ValidityState) {
	js.Rewrite("$<.validity")
	return validity
}

func (*HTMLObjectElement) GetVspace() (vspace int) {
	js.Rewrite("$<.vspace")
	return vspace
}

func (*HTMLObjectElement) SetVspace(vspace int) {
	js.Rewrite("$<.vspace = $1", vspace)
}

func (*HTMLObjectElement) GetWidth() (width string) {
	js.Rewrite("$<.width")
	return width
}

func (*HTMLObjectElement) SetWidth(width string) {
	js.Rewrite("$<.width = $1", width)
}

func (*HTMLObjectElement) GetWillValidate() (willValidate bool) {
	js.Rewrite("$<.willValidate")
	return willValidate
}

type HTMLOListElement struct {
	*HTMLElement
}

func (*HTMLOListElement) GetCompact() (compact bool) {
	js.Rewrite("$<.compact")
	return compact
}

func (*HTMLOListElement) SetCompact(compact bool) {
	js.Rewrite("$<.compact = $1", compact)
}

func (*HTMLOListElement) GetStart() (start int) {
	js.Rewrite("$<.start")
	return start
}

func (*HTMLOListElement) SetStart(start int) {
	js.Rewrite("$<.start = $1", start)
}

func (*HTMLOListElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLOListElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}

type HTMLOptGroupElement struct {
	*HTMLElement
}

func (*HTMLOptGroupElement) GetDefaultSelected() (defaultSelected bool) {
	js.Rewrite("$<.defaultSelected")
	return defaultSelected
}

func (*HTMLOptGroupElement) SetDefaultSelected(defaultSelected bool) {
	js.Rewrite("$<.defaultSelected = $1", defaultSelected)
}

func (*HTMLOptGroupElement) GetDisabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

func (*HTMLOptGroupElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

func (*HTMLOptGroupElement) GetForm() (form *HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

func (*HTMLOptGroupElement) GetIndex() (index int) {
	js.Rewrite("$<.index")
	return index
}

func (*HTMLOptGroupElement) GetLabel() (label string) {
	js.Rewrite("$<.label")
	return label
}

func (*HTMLOptGroupElement) SetLabel(label string) {
	js.Rewrite("$<.label = $1", label)
}

func (*HTMLOptGroupElement) GetSelected() (selected bool) {
	js.Rewrite("$<.selected")
	return selected
}

func (*HTMLOptGroupElement) SetSelected(selected bool) {
	js.Rewrite("$<.selected = $1", selected)
}

func (*HTMLOptGroupElement) GetText() (text string) {
	js.Rewrite("$<.text")
	return text
}

func (*HTMLOptGroupElement) GetValue() (value string) {
	js.Rewrite("$<.value")
	return value
}

func (*HTMLOptGroupElement) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}

type HTMLOptionElement struct {
	*HTMLElement
}

func (*HTMLOptionElement) GetDefaultSelected() (defaultSelected bool) {
	js.Rewrite("$<.defaultSelected")
	return defaultSelected
}

func (*HTMLOptionElement) SetDefaultSelected(defaultSelected bool) {
	js.Rewrite("$<.defaultSelected = $1", defaultSelected)
}

func (*HTMLOptionElement) GetDisabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

func (*HTMLOptionElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

func (*HTMLOptionElement) GetForm() (form *HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

func (*HTMLOptionElement) GetIndex() (index int) {
	js.Rewrite("$<.index")
	return index
}

func (*HTMLOptionElement) GetLabel() (label string) {
	js.Rewrite("$<.label")
	return label
}

func (*HTMLOptionElement) SetLabel(label string) {
	js.Rewrite("$<.label = $1", label)
}

func (*HTMLOptionElement) GetSelected() (selected bool) {
	js.Rewrite("$<.selected")
	return selected
}

func (*HTMLOptionElement) SetSelected(selected bool) {
	js.Rewrite("$<.selected = $1", selected)
}

func (*HTMLOptionElement) GetText() (text string) {
	js.Rewrite("$<.text")
	return text
}

func (*HTMLOptionElement) SetText(text string) {
	js.Rewrite("$<.text = $1", text)
}

func (*HTMLOptionElement) GetValue() (value string) {
	js.Rewrite("$<.value")
	return value
}

func (*HTMLOptionElement) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}

type HTMLOptionsCollection struct {
	*HTMLCollection
}

func (*HTMLOptionsCollection) Add(element interface{}, before interface{}) {
	js.Rewrite("$<.add($1, $2)", element, before)
}

func (*HTMLOptionsCollection) Remove(index int) {
	js.Rewrite("$<.remove($1)", index)
}

func (*HTMLOptionsCollection) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

func (*HTMLOptionsCollection) SetLength(length uint) {
	js.Rewrite("$<.length = $1", length)
}

func (*HTMLOptionsCollection) GetSelectedIndex() (selectedIndex int) {
	js.Rewrite("$<.selectedIndex")
	return selectedIndex
}

func (*HTMLOptionsCollection) SetSelectedIndex(selectedIndex int) {
	js.Rewrite("$<.selectedIndex = $1", selectedIndex)
}

type HTMLOutputElement struct {
	*HTMLElement
}

func (*HTMLOutputElement) CheckValidity() (b bool) {
	js.Rewrite("$<.checkValidity()")
	return b
}

func (*HTMLOutputElement) ReportValidity() (b bool) {
	js.Rewrite("$<.reportValidity()")
	return b
}

func (*HTMLOutputElement) SetCustomValidity(err string) {
	js.Rewrite("$<.setCustomValidity($1)", err)
}

func (*HTMLOutputElement) GetDefaultValue() (defaultValue string) {
	js.Rewrite("$<.defaultValue")
	return defaultValue
}

func (*HTMLOutputElement) SetDefaultValue(defaultValue string) {
	js.Rewrite("$<.defaultValue = $1", defaultValue)
}

func (*HTMLOutputElement) GetForm() (form *HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

func (*HTMLOutputElement) GetHTMLFor() (htmlFor *DOMSettableTokenList) {
	js.Rewrite("$<.htmlFor")
	return htmlFor
}

func (*HTMLOutputElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLOutputElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLOutputElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLOutputElement) GetValidationMessage() (validationMessage string) {
	js.Rewrite("$<.validationMessage")
	return validationMessage
}

func (*HTMLOutputElement) GetValidity() (validity *ValidityState) {
	js.Rewrite("$<.validity")
	return validity
}

func (*HTMLOutputElement) GetValue() (value string) {
	js.Rewrite("$<.value")
	return value
}

func (*HTMLOutputElement) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}

func (*HTMLOutputElement) GetWillValidate() (willValidate bool) {
	js.Rewrite("$<.willValidate")
	return willValidate
}

type HTMLParagraphElement struct {
	*HTMLElement
}

func (*HTMLParagraphElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLParagraphElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLParagraphElement) GetClear() (clear string) {
	js.Rewrite("$<.clear")
	return clear
}

func (*HTMLParagraphElement) SetClear(clear string) {
	js.Rewrite("$<.clear = $1", clear)
}

type HTMLParamElement struct {
	*HTMLElement
}

func (*HTMLParamElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLParamElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLParamElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLParamElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}

func (*HTMLParamElement) GetValue() (value string) {
	js.Rewrite("$<.value")
	return value
}

func (*HTMLParamElement) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}

func (*HTMLParamElement) GetValueType() (valueType string) {
	js.Rewrite("$<.valueType")
	return valueType
}

func (*HTMLParamElement) SetValueType(valueType string) {
	js.Rewrite("$<.valueType = $1", valueType)
}

type HTMLPictureElement struct {
	*HTMLElement
}

type HTMLPreElement struct {
	*HTMLElement
}

func (*HTMLPreElement) GetWidth() (width int) {
	js.Rewrite("$<.width")
	return width
}

func (*HTMLPreElement) SetWidth(width int) {
	js.Rewrite("$<.width = $1", width)
}

type HTMLProgressElement struct {
	*HTMLElement
}

func (*HTMLProgressElement) GetForm() (form *HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

func (*HTMLProgressElement) GetMax() (max float32) {
	js.Rewrite("$<.max")
	return max
}

func (*HTMLProgressElement) SetMax(max float32) {
	js.Rewrite("$<.max = $1", max)
}

func (*HTMLProgressElement) GetPosition() (position float32) {
	js.Rewrite("$<.position")
	return position
}

func (*HTMLProgressElement) GetValue() (value float32) {
	js.Rewrite("$<.value")
	return value
}

func (*HTMLProgressElement) SetValue(value float32) {
	js.Rewrite("$<.value = $1", value)
}

type HTMLQuoteElement struct {
	*HTMLElement
}

func (*HTMLQuoteElement) GetCite() (cite string) {
	js.Rewrite("$<.cite")
	return cite
}

func (*HTMLQuoteElement) SetCite(cite string) {
	js.Rewrite("$<.cite = $1", cite)
}

type HTMLScriptElement struct {
	*HTMLElement
}

func (*HTMLScriptElement) GetAsync() (async bool) {
	js.Rewrite("$<.async")
	return async
}

func (*HTMLScriptElement) SetAsync(async bool) {
	js.Rewrite("$<.async = $1", async)
}

func (*HTMLScriptElement) GetCharset() (charset string) {
	js.Rewrite("$<.charset")
	return charset
}

func (*HTMLScriptElement) SetCharset(charset string) {
	js.Rewrite("$<.charset = $1", charset)
}

func (*HTMLScriptElement) GetCrossOrigin() (crossOrigin string) {
	js.Rewrite("$<.crossOrigin")
	return crossOrigin
}

func (*HTMLScriptElement) SetCrossOrigin(crossOrigin string) {
	js.Rewrite("$<.crossOrigin = $1", crossOrigin)
}

func (*HTMLScriptElement) GetDefer() (dfr bool) {
	js.Rewrite("$<.defer")
	return dfr
}

func (*HTMLScriptElement) SetDefer(dfr bool) {
	js.Rewrite("$<.defer = $1", dfr)
}

func (*HTMLScriptElement) GetEvent() (event string) {
	js.Rewrite("$<.event")
	return event
}

func (*HTMLScriptElement) SetEvent(event string) {
	js.Rewrite("$<.event = $1", event)
}

func (*HTMLScriptElement) GetHTMLFor() (htmlFor string) {
	js.Rewrite("$<.htmlFor")
	return htmlFor
}

func (*HTMLScriptElement) SetHTMLFor(htmlFor string) {
	js.Rewrite("$<.htmlFor = $1", htmlFor)
}

func (*HTMLScriptElement) GetSrc() (src string) {
	js.Rewrite("$<.src")
	return src
}

func (*HTMLScriptElement) SetSrc(src string) {
	js.Rewrite("$<.src = $1", src)
}

func (*HTMLScriptElement) GetText() (text string) {
	js.Rewrite("$<.text")
	return text
}

func (*HTMLScriptElement) SetText(text string) {
	js.Rewrite("$<.text = $1", text)
}

func (*HTMLScriptElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLScriptElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}

type HTMLSelectElement struct {
	*HTMLElement
}

func (*HTMLSelectElement) Add(element *HTMLElement, before interface{}) {
	js.Rewrite("$<.add($1, $2)", element, before)
}

func (*HTMLSelectElement) CheckValidity() (b bool) {
	js.Rewrite("$<.checkValidity()")
	return b
}

func (*HTMLSelectElement) Item(name interface{}, index interface{}) (i interface{}) {
	js.Rewrite("$<.item($1, $2)", name, index)
	return i
}

func (*HTMLSelectElement) NamedItem(name string) (i interface{}) {
	js.Rewrite("$<.namedItem($1)", name)
	return i
}

func (*HTMLSelectElement) Remove(index int) {
	js.Rewrite("$<.remove($1)", index)
}

func (*HTMLSelectElement) SetCustomValidity(err string) {
	js.Rewrite("$<.setCustomValidity($1)", err)
}

func (*HTMLSelectElement) GetAutofocus() (autofocus bool) {
	js.Rewrite("$<.autofocus")
	return autofocus
}

func (*HTMLSelectElement) SetAutofocus(autofocus bool) {
	js.Rewrite("$<.autofocus = $1", autofocus)
}

func (*HTMLSelectElement) GetDisabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

func (*HTMLSelectElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

func (*HTMLSelectElement) GetForm() (form *HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

func (*HTMLSelectElement) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

func (*HTMLSelectElement) SetLength(length uint) {
	js.Rewrite("$<.length = $1", length)
}

func (*HTMLSelectElement) GetMultiple() (multiple bool) {
	js.Rewrite("$<.multiple")
	return multiple
}

func (*HTMLSelectElement) SetMultiple(multiple bool) {
	js.Rewrite("$<.multiple = $1", multiple)
}

func (*HTMLSelectElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLSelectElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLSelectElement) GetOptions() (options *HTMLOptionsCollection) {
	js.Rewrite("$<.options")
	return options
}

func (*HTMLSelectElement) GetRequired() (required bool) {
	js.Rewrite("$<.required")
	return required
}

func (*HTMLSelectElement) SetRequired(required bool) {
	js.Rewrite("$<.required = $1", required)
}

func (*HTMLSelectElement) GetSelectedIndex() (selectedIndex int) {
	js.Rewrite("$<.selectedIndex")
	return selectedIndex
}

func (*HTMLSelectElement) SetSelectedIndex(selectedIndex int) {
	js.Rewrite("$<.selectedIndex = $1", selectedIndex)
}

func (*HTMLSelectElement) GetSelectedOptions() (selectedOptions *HTMLCollection) {
	js.Rewrite("$<.selectedOptions")
	return selectedOptions
}

func (*HTMLSelectElement) GetSize() (size uint) {
	js.Rewrite("$<.size")
	return size
}

func (*HTMLSelectElement) SetSize(size uint) {
	js.Rewrite("$<.size = $1", size)
}

func (*HTMLSelectElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLSelectElement) GetValidationMessage() (validationMessage string) {
	js.Rewrite("$<.validationMessage")
	return validationMessage
}

func (*HTMLSelectElement) GetValidity() (validity *ValidityState) {
	js.Rewrite("$<.validity")
	return validity
}

func (*HTMLSelectElement) GetValue() (value string) {
	js.Rewrite("$<.value")
	return value
}

func (*HTMLSelectElement) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}

func (*HTMLSelectElement) GetWillValidate() (willValidate bool) {
	js.Rewrite("$<.willValidate")
	return willValidate
}

type HTMLSourceElement struct {
	*HTMLElement
}

func (*HTMLSourceElement) GetMedia() (media string) {
	js.Rewrite("$<.media")
	return media
}

func (*HTMLSourceElement) SetMedia(media string) {
	js.Rewrite("$<.media = $1", media)
}

func (*HTMLSourceElement) GetMsKeySystem() (msKeySystem string) {
	js.Rewrite("$<.msKeySystem")
	return msKeySystem
}

func (*HTMLSourceElement) SetMsKeySystem(msKeySystem string) {
	js.Rewrite("$<.msKeySystem = $1", msKeySystem)
}

func (*HTMLSourceElement) GetSizes() (sizes string) {
	js.Rewrite("$<.sizes")
	return sizes
}

func (*HTMLSourceElement) SetSizes(sizes string) {
	js.Rewrite("$<.sizes = $1", sizes)
}

func (*HTMLSourceElement) GetSrc() (src string) {
	js.Rewrite("$<.src")
	return src
}

func (*HTMLSourceElement) SetSrc(src string) {
	js.Rewrite("$<.src = $1", src)
}

func (*HTMLSourceElement) GetSrcset() (srcset string) {
	js.Rewrite("$<.srcset")
	return srcset
}

func (*HTMLSourceElement) SetSrcset(srcset string) {
	js.Rewrite("$<.srcset = $1", srcset)
}

func (*HTMLSourceElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLSourceElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}

type HTMLSpanElement struct {
	*HTMLElement
}

type HTMLStyleElement struct {
	*HTMLElement
	*LinkStyle
}

func (*HTMLStyleElement) GetDisabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

func (*HTMLStyleElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

func (*HTMLStyleElement) GetMedia() (media string) {
	js.Rewrite("$<.media")
	return media
}

func (*HTMLStyleElement) SetMedia(media string) {
	js.Rewrite("$<.media = $1", media)
}

func (*HTMLStyleElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLStyleElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}

type HTMLTableCaptionElement struct {
	*HTMLElement
}

func (*HTMLTableCaptionElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLTableCaptionElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLTableCaptionElement) GetVAlign() (vAlign string) {
	js.Rewrite("$<.vAlign")
	return vAlign
}

func (*HTMLTableCaptionElement) SetVAlign(vAlign string) {
	js.Rewrite("$<.vAlign = $1", vAlign)
}

type HTMLTableCellElement struct {
	*HTMLElement
	*HTMLTableAlignment
}

func (*HTMLTableCellElement) GetAbbr() (abbr string) {
	js.Rewrite("$<.abbr")
	return abbr
}

func (*HTMLTableCellElement) SetAbbr(abbr string) {
	js.Rewrite("$<.abbr = $1", abbr)
}

func (*HTMLTableCellElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLTableCellElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLTableCellElement) GetAxis() (axis string) {
	js.Rewrite("$<.axis")
	return axis
}

func (*HTMLTableCellElement) SetAxis(axis string) {
	js.Rewrite("$<.axis = $1", axis)
}

func (*HTMLTableCellElement) GetBgColor() (bgColor interface{}) {
	js.Rewrite("$<.bgColor")
	return bgColor
}

func (*HTMLTableCellElement) SetBgColor(bgColor interface{}) {
	js.Rewrite("$<.bgColor = $1", bgColor)
}

func (*HTMLTableCellElement) GetCellIndex() (cellIndex int) {
	js.Rewrite("$<.cellIndex")
	return cellIndex
}

func (*HTMLTableCellElement) GetColSpan() (colSpan uint) {
	js.Rewrite("$<.colSpan")
	return colSpan
}

func (*HTMLTableCellElement) SetColSpan(colSpan uint) {
	js.Rewrite("$<.colSpan = $1", colSpan)
}

func (*HTMLTableCellElement) GetHeaders() (headers string) {
	js.Rewrite("$<.headers")
	return headers
}

func (*HTMLTableCellElement) SetHeaders(headers string) {
	js.Rewrite("$<.headers = $1", headers)
}

func (*HTMLTableCellElement) GetHeight() (height interface{}) {
	js.Rewrite("$<.height")
	return height
}

func (*HTMLTableCellElement) SetHeight(height interface{}) {
	js.Rewrite("$<.height = $1", height)
}

func (*HTMLTableCellElement) GetNoWrap() (noWrap bool) {
	js.Rewrite("$<.noWrap")
	return noWrap
}

func (*HTMLTableCellElement) SetNoWrap(noWrap bool) {
	js.Rewrite("$<.noWrap = $1", noWrap)
}

func (*HTMLTableCellElement) GetRowSpan() (rowSpan uint) {
	js.Rewrite("$<.rowSpan")
	return rowSpan
}

func (*HTMLTableCellElement) SetRowSpan(rowSpan uint) {
	js.Rewrite("$<.rowSpan = $1", rowSpan)
}

func (*HTMLTableCellElement) GetScope() (scope string) {
	js.Rewrite("$<.scope")
	return scope
}

func (*HTMLTableCellElement) SetScope(scope string) {
	js.Rewrite("$<.scope = $1", scope)
}

func (*HTMLTableCellElement) GetWidth() (width string) {
	js.Rewrite("$<.width")
	return width
}

func (*HTMLTableCellElement) SetWidth(width string) {
	js.Rewrite("$<.width = $1", width)
}

type HTMLTableColElement struct {
	*HTMLElement
	*HTMLTableAlignment
}

func (*HTMLTableColElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLTableColElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLTableColElement) GetSpan() (span uint) {
	js.Rewrite("$<.span")
	return span
}

func (*HTMLTableColElement) SetSpan(span uint) {
	js.Rewrite("$<.span = $1", span)
}

func (*HTMLTableColElement) GetWidth() (width interface{}) {
	js.Rewrite("$<.width")
	return width
}

func (*HTMLTableColElement) SetWidth(width interface{}) {
	js.Rewrite("$<.width = $1", width)
}

type HTMLTableDataCellElement struct {
	*HTMLTableCellElement
}

type HTMLTableElement struct {
	*HTMLElement
}

func (*HTMLTableElement) CreateCaption() (h *HTMLElement) {
	js.Rewrite("$<.createCaption()")
	return h
}

func (*HTMLTableElement) CreateTBody() (h *HTMLElement) {
	js.Rewrite("$<.createTBody()")
	return h
}

func (*HTMLTableElement) CreateTFoot() (h *HTMLElement) {
	js.Rewrite("$<.createTFoot()")
	return h
}

func (*HTMLTableElement) CreateTHead() (h *HTMLElement) {
	js.Rewrite("$<.createTHead()")
	return h
}

func (*HTMLTableElement) DeleteCaption() {
	js.Rewrite("$<.deleteCaption()")
}

func (*HTMLTableElement) DeleteRow(index int) {
	js.Rewrite("$<.deleteRow($1)", index)
}

func (*HTMLTableElement) DeleteTFoot() {
	js.Rewrite("$<.deleteTFoot()")
}

func (*HTMLTableElement) DeleteTHead() {
	js.Rewrite("$<.deleteTHead()")
}

func (*HTMLTableElement) InsertRow(index int) (h *HTMLElement) {
	js.Rewrite("$<.insertRow($1)", index)
	return h
}

func (*HTMLTableElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLTableElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLTableElement) GetBgColor() (bgColor interface{}) {
	js.Rewrite("$<.bgColor")
	return bgColor
}

func (*HTMLTableElement) SetBgColor(bgColor interface{}) {
	js.Rewrite("$<.bgColor = $1", bgColor)
}

func (*HTMLTableElement) GetBorder() (border string) {
	js.Rewrite("$<.border")
	return border
}

func (*HTMLTableElement) SetBorder(border string) {
	js.Rewrite("$<.border = $1", border)
}

func (*HTMLTableElement) GetBorderColor() (borderColor interface{}) {
	js.Rewrite("$<.borderColor")
	return borderColor
}

func (*HTMLTableElement) SetBorderColor(borderColor interface{}) {
	js.Rewrite("$<.borderColor = $1", borderColor)
}

func (*HTMLTableElement) GetCaption() (caption *HTMLTableCaptionElement) {
	js.Rewrite("$<.caption")
	return caption
}

func (*HTMLTableElement) SetCaption(caption *HTMLTableCaptionElement) {
	js.Rewrite("$<.caption = $1", caption)
}

func (*HTMLTableElement) GetCellPadding() (cellPadding string) {
	js.Rewrite("$<.cellPadding")
	return cellPadding
}

func (*HTMLTableElement) SetCellPadding(cellPadding string) {
	js.Rewrite("$<.cellPadding = $1", cellPadding)
}

func (*HTMLTableElement) GetCellSpacing() (cellSpacing string) {
	js.Rewrite("$<.cellSpacing")
	return cellSpacing
}

func (*HTMLTableElement) SetCellSpacing(cellSpacing string) {
	js.Rewrite("$<.cellSpacing = $1", cellSpacing)
}

func (*HTMLTableElement) GetCols() (cols int) {
	js.Rewrite("$<.cols")
	return cols
}

func (*HTMLTableElement) SetCols(cols int) {
	js.Rewrite("$<.cols = $1", cols)
}

func (*HTMLTableElement) GetFrame() (frame string) {
	js.Rewrite("$<.frame")
	return frame
}

func (*HTMLTableElement) SetFrame(frame string) {
	js.Rewrite("$<.frame = $1", frame)
}

func (*HTMLTableElement) GetHeight() (height interface{}) {
	js.Rewrite("$<.height")
	return height
}

func (*HTMLTableElement) SetHeight(height interface{}) {
	js.Rewrite("$<.height = $1", height)
}

func (*HTMLTableElement) GetRows() (rows *HTMLCollection) {
	js.Rewrite("$<.rows")
	return rows
}

func (*HTMLTableElement) GetRules() (rules string) {
	js.Rewrite("$<.rules")
	return rules
}

func (*HTMLTableElement) SetRules(rules string) {
	js.Rewrite("$<.rules = $1", rules)
}

func (*HTMLTableElement) GetSummary() (summary string) {
	js.Rewrite("$<.summary")
	return summary
}

func (*HTMLTableElement) SetSummary(summary string) {
	js.Rewrite("$<.summary = $1", summary)
}

func (*HTMLTableElement) GetTBodies() (tBodies *HTMLCollection) {
	js.Rewrite("$<.tBodies")
	return tBodies
}

func (*HTMLTableElement) GetTFoot() (tFoot *HTMLTableSectionElement) {
	js.Rewrite("$<.tFoot")
	return tFoot
}

func (*HTMLTableElement) SetTFoot(tFoot *HTMLTableSectionElement) {
	js.Rewrite("$<.tFoot = $1", tFoot)
}

func (*HTMLTableElement) GetTHead() (tHead *HTMLTableSectionElement) {
	js.Rewrite("$<.tHead")
	return tHead
}

func (*HTMLTableElement) SetTHead(tHead *HTMLTableSectionElement) {
	js.Rewrite("$<.tHead = $1", tHead)
}

func (*HTMLTableElement) GetWidth() (width string) {
	js.Rewrite("$<.width")
	return width
}

func (*HTMLTableElement) SetWidth(width string) {
	js.Rewrite("$<.width = $1", width)
}

type HTMLTableHeaderCellElement struct {
	*HTMLTableCellElement
}

func (*HTMLTableHeaderCellElement) GetScope() (scope string) {
	js.Rewrite("$<.scope")
	return scope
}

func (*HTMLTableHeaderCellElement) SetScope(scope string) {
	js.Rewrite("$<.scope = $1", scope)
}

type HTMLTableRowElement struct {
	*HTMLElement
	*HTMLTableAlignment
}

func (*HTMLTableRowElement) DeleteCell(index int) {
	js.Rewrite("$<.deleteCell($1)", index)
}

func (*HTMLTableRowElement) InsertCell(index int) (h *HTMLElement) {
	js.Rewrite("$<.insertCell($1)", index)
	return h
}

func (*HTMLTableRowElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLTableRowElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLTableRowElement) GetBgColor() (bgColor interface{}) {
	js.Rewrite("$<.bgColor")
	return bgColor
}

func (*HTMLTableRowElement) SetBgColor(bgColor interface{}) {
	js.Rewrite("$<.bgColor = $1", bgColor)
}

func (*HTMLTableRowElement) GetCells() (cells *HTMLCollection) {
	js.Rewrite("$<.cells")
	return cells
}

func (*HTMLTableRowElement) GetHeight() (height interface{}) {
	js.Rewrite("$<.height")
	return height
}

func (*HTMLTableRowElement) SetHeight(height interface{}) {
	js.Rewrite("$<.height = $1", height)
}

func (*HTMLTableRowElement) GetRowIndex() (rowIndex int) {
	js.Rewrite("$<.rowIndex")
	return rowIndex
}

func (*HTMLTableRowElement) GetSectionRowIndex() (sectionRowIndex int) {
	js.Rewrite("$<.sectionRowIndex")
	return sectionRowIndex
}

type HTMLTableSectionElement struct {
	*HTMLElement
	*HTMLTableAlignment
}

func (*HTMLTableSectionElement) DeleteRow(index int) {
	js.Rewrite("$<.deleteRow($1)", index)
}

func (*HTMLTableSectionElement) InsertRow(index int) (h *HTMLElement) {
	js.Rewrite("$<.insertRow($1)", index)
	return h
}

func (*HTMLTableSectionElement) GetAlign() (align string) {
	js.Rewrite("$<.align")
	return align
}

func (*HTMLTableSectionElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

func (*HTMLTableSectionElement) GetRows() (rows *HTMLCollection) {
	js.Rewrite("$<.rows")
	return rows
}

type HTMLTemplateElement struct {
	*HTMLElement
}

func (*HTMLTemplateElement) GetContent() (content *DocumentFragment) {
	js.Rewrite("$<.content")
	return content
}

type HTMLTextAreaElement struct {
	*HTMLElement
}

func (*HTMLTextAreaElement) CheckValidity() (b bool) {
	js.Rewrite("$<.checkValidity()")
	return b
}

func (*HTMLTextAreaElement) Select() {
	js.Rewrite("$<.select()")
}

func (*HTMLTextAreaElement) SetCustomValidity(err string) {
	js.Rewrite("$<.setCustomValidity($1)", err)
}

func (*HTMLTextAreaElement) SetSelectionRange(start uint, end uint) {
	js.Rewrite("$<.setSelectionRange($1, $2)", start, end)
}

func (*HTMLTextAreaElement) GetAutofocus() (autofocus bool) {
	js.Rewrite("$<.autofocus")
	return autofocus
}

func (*HTMLTextAreaElement) SetAutofocus(autofocus bool) {
	js.Rewrite("$<.autofocus = $1", autofocus)
}

func (*HTMLTextAreaElement) GetCols() (cols uint) {
	js.Rewrite("$<.cols")
	return cols
}

func (*HTMLTextAreaElement) SetCols(cols uint) {
	js.Rewrite("$<.cols = $1", cols)
}

func (*HTMLTextAreaElement) GetDefaultValue() (defaultValue string) {
	js.Rewrite("$<.defaultValue")
	return defaultValue
}

func (*HTMLTextAreaElement) SetDefaultValue(defaultValue string) {
	js.Rewrite("$<.defaultValue = $1", defaultValue)
}

func (*HTMLTextAreaElement) GetDisabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

func (*HTMLTextAreaElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

func (*HTMLTextAreaElement) GetForm() (form *HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

func (*HTMLTextAreaElement) GetMaxLength() (maxLength int) {
	js.Rewrite("$<.maxLength")
	return maxLength
}

func (*HTMLTextAreaElement) SetMaxLength(maxLength int) {
	js.Rewrite("$<.maxLength = $1", maxLength)
}

func (*HTMLTextAreaElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLTextAreaElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLTextAreaElement) GetPlaceholder() (placeholder string) {
	js.Rewrite("$<.placeholder")
	return placeholder
}

func (*HTMLTextAreaElement) SetPlaceholder(placeholder string) {
	js.Rewrite("$<.placeholder = $1", placeholder)
}

func (*HTMLTextAreaElement) GetReadOnly() (readOnly bool) {
	js.Rewrite("$<.readOnly")
	return readOnly
}

func (*HTMLTextAreaElement) SetReadOnly(readOnly bool) {
	js.Rewrite("$<.readOnly = $1", readOnly)
}

func (*HTMLTextAreaElement) GetRequired() (required bool) {
	js.Rewrite("$<.required")
	return required
}

func (*HTMLTextAreaElement) SetRequired(required bool) {
	js.Rewrite("$<.required = $1", required)
}

func (*HTMLTextAreaElement) GetRows() (rows uint) {
	js.Rewrite("$<.rows")
	return rows
}

func (*HTMLTextAreaElement) SetRows(rows uint) {
	js.Rewrite("$<.rows = $1", rows)
}

func (*HTMLTextAreaElement) GetSelectionEnd() (selectionEnd uint) {
	js.Rewrite("$<.selectionEnd")
	return selectionEnd
}

func (*HTMLTextAreaElement) SetSelectionEnd(selectionEnd uint) {
	js.Rewrite("$<.selectionEnd = $1", selectionEnd)
}

func (*HTMLTextAreaElement) GetSelectionStart() (selectionStart uint) {
	js.Rewrite("$<.selectionStart")
	return selectionStart
}

func (*HTMLTextAreaElement) SetSelectionStart(selectionStart uint) {
	js.Rewrite("$<.selectionStart = $1", selectionStart)
}

func (*HTMLTextAreaElement) GetStatus() (status interface{}) {
	js.Rewrite("$<.status")
	return status
}

func (*HTMLTextAreaElement) SetStatus(status interface{}) {
	js.Rewrite("$<.status = $1", status)
}

func (*HTMLTextAreaElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLTextAreaElement) GetValidationMessage() (validationMessage string) {
	js.Rewrite("$<.validationMessage")
	return validationMessage
}

func (*HTMLTextAreaElement) GetValidity() (validity *ValidityState) {
	js.Rewrite("$<.validity")
	return validity
}

func (*HTMLTextAreaElement) GetValue() (value string) {
	js.Rewrite("$<.value")
	return value
}

func (*HTMLTextAreaElement) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}

func (*HTMLTextAreaElement) GetWillValidate() (willValidate bool) {
	js.Rewrite("$<.willValidate")
	return willValidate
}

func (*HTMLTextAreaElement) GetWrap() (wrap string) {
	js.Rewrite("$<.wrap")
	return wrap
}

func (*HTMLTextAreaElement) SetWrap(wrap string) {
	js.Rewrite("$<.wrap = $1", wrap)
}

type HTMLTimeElement struct {
	*HTMLElement
}

func (*HTMLTimeElement) GetDateTime() (dateTime string) {
	js.Rewrite("$<.dateTime")
	return dateTime
}

func (*HTMLTimeElement) SetDateTime(dateTime string) {
	js.Rewrite("$<.dateTime = $1", dateTime)
}

type HTMLTitleElement struct {
	*HTMLElement
}

func (*HTMLTitleElement) GetText() (text string) {
	js.Rewrite("$<.text")
	return text
}

func (*HTMLTitleElement) SetText(text string) {
	js.Rewrite("$<.text = $1", text)
}

type HTMLTrackElement struct {
	*HTMLElement
}

func (*HTMLTrackElement) GetDefault() (def bool) {
	js.Rewrite("$<.default")
	return def
}

func (*HTMLTrackElement) SetDefault(def bool) {
	js.Rewrite("$<.default = $1", def)
}

func (*HTMLTrackElement) GetKind() (kind string) {
	js.Rewrite("$<.kind")
	return kind
}

func (*HTMLTrackElement) SetKind(kind string) {
	js.Rewrite("$<.kind = $1", kind)
}

func (*HTMLTrackElement) GetLabel() (label string) {
	js.Rewrite("$<.label")
	return label
}

func (*HTMLTrackElement) SetLabel(label string) {
	js.Rewrite("$<.label = $1", label)
}

func (*HTMLTrackElement) GetReadyState() (readyState uint8) {
	js.Rewrite("$<.readyState")
	return readyState
}

func (*HTMLTrackElement) GetSrc() (src string) {
	js.Rewrite("$<.src")
	return src
}

func (*HTMLTrackElement) SetSrc(src string) {
	js.Rewrite("$<.src = $1", src)
}

func (*HTMLTrackElement) GetSrclang() (srclang string) {
	js.Rewrite("$<.srclang")
	return srclang
}

func (*HTMLTrackElement) SetSrclang(srclang string) {
	js.Rewrite("$<.srclang = $1", srclang)
}

func (*HTMLTrackElement) GetTrack() (track *TextTrack) {
	js.Rewrite("$<.track")
	return track
}

type HTMLUListElement struct {
	*HTMLElement
}

func (*HTMLUListElement) GetCompact() (compact bool) {
	js.Rewrite("$<.compact")
	return compact
}

func (*HTMLUListElement) SetCompact(compact bool) {
	js.Rewrite("$<.compact = $1", compact)
}

func (*HTMLUListElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*HTMLUListElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}

type HTMLUnknownElement struct {
	*HTMLElement
}

type HTMLVideoElement struct {
	*HTMLMediaElement
}

func (*HTMLVideoElement) GetVideoPlaybackQuality() (v *VideoPlaybackQuality) {
	js.Rewrite("$<.getVideoPlaybackQuality()")
	return v
}

func (*HTMLVideoElement) MsFrameStep(forward bool) {
	js.Rewrite("$<.msFrameStep($1)", forward)
}

func (*HTMLVideoElement) MsInsertVideoEffect(activatableClassId string, effectRequired bool, config interface{}) {
	js.Rewrite("$<.msInsertVideoEffect($1, $2, $3)", activatableClassId, effectRequired, config)
}

func (*HTMLVideoElement) MsSetVideoRectangle(left float32, top float32, right float32, bottom float32) {
	js.Rewrite("$<.msSetVideoRectangle($1, $2, $3, $4)", left, top, right, bottom)
}

func (*HTMLVideoElement) WebkitEnterFullscreen() {
	js.Rewrite("$<.webkitEnterFullscreen()")
}

func (*HTMLVideoElement) WebkitEnterFullScreen() {
	js.Rewrite("$<.webkitEnterFullScreen()")
}

func (*HTMLVideoElement) WebkitExitFullscreen() {
	js.Rewrite("$<.webkitExitFullscreen()")
}

func (*HTMLVideoElement) WebkitExitFullScreen() {
	js.Rewrite("$<.webkitExitFullScreen()")
}

func (*HTMLVideoElement) GetHeight() (height uint) {
	js.Rewrite("$<.height")
	return height
}

func (*HTMLVideoElement) SetHeight(height uint) {
	js.Rewrite("$<.height = $1", height)
}

func (*HTMLVideoElement) GetMsHorizontalMirror() (msHorizontalMirror bool) {
	js.Rewrite("$<.msHorizontalMirror")
	return msHorizontalMirror
}

func (*HTMLVideoElement) SetMsHorizontalMirror(msHorizontalMirror bool) {
	js.Rewrite("$<.msHorizontalMirror = $1", msHorizontalMirror)
}

func (*HTMLVideoElement) GetMsIsLayoutOptimalForPlayback() (msIsLayoutOptimalForPlayback bool) {
	js.Rewrite("$<.msIsLayoutOptimalForPlayback")
	return msIsLayoutOptimalForPlayback
}

func (*HTMLVideoElement) GetMsIsStereo3d() (msIsStereo3D bool) {
	js.Rewrite("$<.msIsStereo3D")
	return msIsStereo3D
}

func (*HTMLVideoElement) GetMsStereo3dPackingMode() (msStereo3DPackingMode string) {
	js.Rewrite("$<.msStereo3DPackingMode")
	return msStereo3DPackingMode
}

func (*HTMLVideoElement) SetMsStereo3dPackingMode(msStereo3DPackingMode string) {
	js.Rewrite("$<.msStereo3DPackingMode = $1", msStereo3DPackingMode)
}

func (*HTMLVideoElement) GetMsStereo3dRenderMode() (msStereo3DRenderMode string) {
	js.Rewrite("$<.msStereo3DRenderMode")
	return msStereo3DRenderMode
}

func (*HTMLVideoElement) SetMsStereo3dRenderMode(msStereo3DRenderMode string) {
	js.Rewrite("$<.msStereo3DRenderMode = $1", msStereo3DRenderMode)
}

func (*HTMLVideoElement) GetMsZoom() (msZoom bool) {
	js.Rewrite("$<.msZoom")
	return msZoom
}

func (*HTMLVideoElement) SetMsZoom(msZoom bool) {
	js.Rewrite("$<.msZoom = $1", msZoom)
}

func (*HTMLVideoElement) GetOnMSVideoFormatChanged() (MSVideoFormatChanged *Event) {
	js.Rewrite("$<.onMSVideoFormatChanged")
	return MSVideoFormatChanged
}

func (*HTMLVideoElement) SetOnMSVideoFormatChanged(MSVideoFormatChanged *Event) {
	js.Rewrite("$<.onMSVideoFormatChanged = $1", MSVideoFormatChanged)
}

func (*HTMLVideoElement) GetOnMSVideoFrameStepCompleted() (MSVideoFrameStepCompleted *Event) {
	js.Rewrite("$<.onMSVideoFrameStepCompleted")
	return MSVideoFrameStepCompleted
}

func (*HTMLVideoElement) SetOnMSVideoFrameStepCompleted(MSVideoFrameStepCompleted *Event) {
	js.Rewrite("$<.onMSVideoFrameStepCompleted = $1", MSVideoFrameStepCompleted)
}

func (*HTMLVideoElement) GetOnMSVideoOptimalLayoutChanged() (MSVideoOptimalLayoutChanged *Event) {
	js.Rewrite("$<.onMSVideoOptimalLayoutChanged")
	return MSVideoOptimalLayoutChanged
}

func (*HTMLVideoElement) SetOnMSVideoOptimalLayoutChanged(MSVideoOptimalLayoutChanged *Event) {
	js.Rewrite("$<.onMSVideoOptimalLayoutChanged = $1", MSVideoOptimalLayoutChanged)
}

func (*HTMLVideoElement) GetPoster() (poster string) {
	js.Rewrite("$<.poster")
	return poster
}

func (*HTMLVideoElement) SetPoster(poster string) {
	js.Rewrite("$<.poster = $1", poster)
}

func (*HTMLVideoElement) GetVideoHeight() (videoHeight uint) {
	js.Rewrite("$<.videoHeight")
	return videoHeight
}

func (*HTMLVideoElement) GetVideoWidth() (videoWidth uint) {
	js.Rewrite("$<.videoWidth")
	return videoWidth
}

func (*HTMLVideoElement) GetWebkitDisplayingFullscreen() (webkitDisplayingFullscreen bool) {
	js.Rewrite("$<.webkitDisplayingFullscreen")
	return webkitDisplayingFullscreen
}

func (*HTMLVideoElement) GetWebkitSupportsFullscreen() (webkitSupportsFullscreen bool) {
	js.Rewrite("$<.webkitSupportsFullscreen")
	return webkitSupportsFullscreen
}

func (*HTMLVideoElement) GetWidth() (width uint) {
	js.Rewrite("$<.width")
	return width
}

func (*HTMLVideoElement) SetWidth(width uint) {
	js.Rewrite("$<.width = $1", width)
}

type IDBCursor struct {
}

func (*IDBCursor) Advance(count uint) {
	js.Rewrite("$<.advance($1)", count)
}

func (*IDBCursor) Continue(key interface{}) {
	js.Rewrite("$<.continue($1)", key)
}

func (*IDBCursor) Delete() (i *IDBRequest) {
	js.Rewrite("$<.delete()")
	return i
}

func (*IDBCursor) Update(value interface{}) (i *IDBRequest) {
	js.Rewrite("$<.update($1)", value)
	return i
}

func (*IDBCursor) GetDirection() (direction *IDBCursorDirection) {
	js.Rewrite("$<.direction")
	return direction
}

func (*IDBCursor) GetKey() (key interface{}) {
	js.Rewrite("$<.key")
	return key
}

func (*IDBCursor) GetPrimaryKey() (primaryKey interface{}) {
	js.Rewrite("$<.primaryKey")
	return primaryKey
}

func (*IDBCursor) GetSource() (source interface{}) {
	js.Rewrite("$<.source")
	return source
}

type IDBCursorWithValue struct {
	*IDBCursor
}

func (*IDBCursorWithValue) GetValue() (value interface{}) {
	js.Rewrite("$<.value")
	return value
}

type IDBDatabase struct {
	*EventTarget
}

func (*IDBDatabase) Close() {
	js.Rewrite("$<.close()")
}

func (*IDBDatabase) CreateObjectStore(name string, optionalParameters *IDBObjectStoreParameters) (i *IDBObjectStore) {
	js.Rewrite("$<.createObjectStore($1, $2)", name, optionalParameters)
	return i
}

func (*IDBDatabase) DeleteObjectStore(name string) {
	js.Rewrite("$<.deleteObjectStore($1)", name)
}

func (*IDBDatabase) Transaction(storeNames interface{}, mode *IDBTransactionMode) (i *IDBTransaction) {
	js.Rewrite("$<.transaction($1, $2)", storeNames, mode)
	return i
}

func (*IDBDatabase) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*IDBDatabase) GetObjectStoreNames() (objectStoreNames *DOMStringList) {
	js.Rewrite("$<.objectStoreNames")
	return objectStoreNames
}

func (*IDBDatabase) GetOnabort() (abort *Event) {
	js.Rewrite("$<.onabort")
	return abort
}

func (*IDBDatabase) SetOnabort(abort *Event) {
	js.Rewrite("$<.onabort = $1", abort)
}

func (*IDBDatabase) GetOnerror() (err *Event) {
	js.Rewrite("$<.onerror")
	return err
}

func (*IDBDatabase) SetOnerror(err *Event) {
	js.Rewrite("$<.onerror = $1", err)
}

func (*IDBDatabase) GetVersion() (version uint64) {
	js.Rewrite("$<.version")
	return version
}

type IDBFactory struct {
}

func (*IDBFactory) Cmp(first interface{}, second interface{}) (i int8) {
	js.Rewrite("$<.cmp($1, $2)", first, second)
	return i
}

func (*IDBFactory) DeleteDatabase(name string) (i *IDBOpenDBRequest) {
	js.Rewrite("$<.deleteDatabase($1)", name)
	return i
}

func (*IDBFactory) Open(name string, version uint64) (i *IDBOpenDBRequest) {
	js.Rewrite("$<.open($1, $2)", name, version)
	return i
}

type IDBIndex struct {
}

func (*IDBIndex) Count(key interface{}) (i *IDBRequest) {
	js.Rewrite("$<.count($1)", key)
	return i
}

func (*IDBIndex) Get(key interface{}) (i *IDBRequest) {
	js.Rewrite("$<.get($1)", key)
	return i
}

func (*IDBIndex) GetKey(key interface{}) (i *IDBRequest) {
	js.Rewrite("$<.getKey($1)", key)
	return i
}

func (*IDBIndex) OpenCursor(rng interface{}, direction *IDBCursorDirection) (i *IDBRequest) {
	js.Rewrite("$<.openCursor($1, $2)", rng, direction)
	return i
}

func (*IDBIndex) OpenKeyCursor(rng interface{}, direction *IDBCursorDirection) (i *IDBRequest) {
	js.Rewrite("$<.openKeyCursor($1, $2)", rng, direction)
	return i
}

func (*IDBIndex) GetKeyPath() (keyPath string) {
	js.Rewrite("$<.keyPath")
	return keyPath
}

func (*IDBIndex) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*IDBIndex) GetObjectStore() (objectStore *IDBObjectStore) {
	js.Rewrite("$<.objectStore")
	return objectStore
}

func (*IDBIndex) GetUnique() (unique bool) {
	js.Rewrite("$<.unique")
	return unique
}

type IDBKeyRange struct {
}

func (*IDBKeyRange) Bound(lower interface{}, upper interface{}, lowerOpen bool, upperOpen bool) (i *IDBKeyRange) {
	js.Rewrite("$<.bound($1, $2, $3, $4)", lower, upper, lowerOpen, upperOpen)
	return i
}

func (*IDBKeyRange) LowerBound(lower interface{}, open bool) (i *IDBKeyRange) {
	js.Rewrite("$<.lowerBound($1, $2)", lower, open)
	return i
}

func (*IDBKeyRange) Only(value interface{}) (i *IDBKeyRange) {
	js.Rewrite("$<.only($1)", value)
	return i
}

func (*IDBKeyRange) UpperBound(upper interface{}, open bool) (i *IDBKeyRange) {
	js.Rewrite("$<.upperBound($1, $2)", upper, open)
	return i
}

func (*IDBKeyRange) GetLower() (lower interface{}) {
	js.Rewrite("$<.lower")
	return lower
}

func (*IDBKeyRange) GetLowerOpen() (lowerOpen bool) {
	js.Rewrite("$<.lowerOpen")
	return lowerOpen
}

func (*IDBKeyRange) GetUpper() (upper interface{}) {
	js.Rewrite("$<.upper")
	return upper
}

func (*IDBKeyRange) GetUpperOpen() (upperOpen bool) {
	js.Rewrite("$<.upperOpen")
	return upperOpen
}

type IDBObjectStore struct {
}

func (*IDBObjectStore) Add(value interface{}, key interface{}) (i *IDBRequest) {
	js.Rewrite("$<.add($1, $2)", value, key)
	return i
}

func (*IDBObjectStore) Clear() (i *IDBRequest) {
	js.Rewrite("$<.clear()")
	return i
}

func (*IDBObjectStore) Count(key interface{}) (i *IDBRequest) {
	js.Rewrite("$<.count($1)", key)
	return i
}

func (*IDBObjectStore) CreateIndex(name string, keyPath string, optionalParameters *IDBIndexParameters) (i *IDBIndex) {
	js.Rewrite("$<.createIndex($1, $2, $3)", name, keyPath, optionalParameters)
	return i
}

func (*IDBObjectStore) Delete(key interface{}) (i *IDBRequest) {
	js.Rewrite("$<.delete($1)", key)
	return i
}

func (*IDBObjectStore) DeleteIndex(indexName string) {
	js.Rewrite("$<.deleteIndex($1)", indexName)
}

func (*IDBObjectStore) Get(key interface{}) (i *IDBRequest) {
	js.Rewrite("$<.get($1)", key)
	return i
}

func (*IDBObjectStore) Index(name string) (i *IDBIndex) {
	js.Rewrite("$<.index($1)", name)
	return i
}

func (*IDBObjectStore) OpenCursor(rng interface{}, direction *IDBCursorDirection) (i *IDBRequest) {
	js.Rewrite("$<.openCursor($1, $2)", rng, direction)
	return i
}

func (*IDBObjectStore) Put(value interface{}, key interface{}) (i *IDBRequest) {
	js.Rewrite("$<.put($1, $2)", value, key)
	return i
}

func (*IDBObjectStore) GetIndexNames() (indexNames *DOMStringList) {
	js.Rewrite("$<.indexNames")
	return indexNames
}

func (*IDBObjectStore) GetKeyPath() (keyPath string) {
	js.Rewrite("$<.keyPath")
	return keyPath
}

func (*IDBObjectStore) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*IDBObjectStore) GetTransaction() (transaction *IDBTransaction) {
	js.Rewrite("$<.transaction")
	return transaction
}

type IDBOpenDBRequest struct {
	*IDBRequest
}

func (*IDBOpenDBRequest) GetOnblocked() (blocked *Event) {
	js.Rewrite("$<.onblocked")
	return blocked
}

func (*IDBOpenDBRequest) SetOnblocked(blocked *Event) {
	js.Rewrite("$<.onblocked = $1", blocked)
}

func (*IDBOpenDBRequest) GetOnupgradeneeded() (upgradeneeded *IDBVersionChangeEvent) {
	js.Rewrite("$<.onupgradeneeded")
	return upgradeneeded
}

func (*IDBOpenDBRequest) SetOnupgradeneeded(upgradeneeded *IDBVersionChangeEvent) {
	js.Rewrite("$<.onupgradeneeded = $1", upgradeneeded)
}

type IDBRequest struct {
	*EventTarget
}

func (*IDBRequest) GetError() (err *DOMError) {
	js.Rewrite("$<.error")
	return err
}

func (*IDBRequest) GetOnerror() (err *Event) {
	js.Rewrite("$<.onerror")
	return err
}

func (*IDBRequest) SetOnerror(err *Event) {
	js.Rewrite("$<.onerror = $1", err)
}

func (*IDBRequest) GetOnsuccess() (success *Event) {
	js.Rewrite("$<.onsuccess")
	return success
}

func (*IDBRequest) SetOnsuccess(success *Event) {
	js.Rewrite("$<.onsuccess = $1", success)
}

func (*IDBRequest) GetReadyState() (readyState *IDBRequestReadyState) {
	js.Rewrite("$<.readyState")
	return readyState
}

func (*IDBRequest) GetResult() (result interface{}) {
	js.Rewrite("$<.result")
	return result
}

func (*IDBRequest) GetSource() (source interface{}) {
	js.Rewrite("$<.source")
	return source
}

func (*IDBRequest) GetTransaction() (transaction *IDBTransaction) {
	js.Rewrite("$<.transaction")
	return transaction
}

type IDBTransaction struct {
	*EventTarget
}

func (*IDBTransaction) Abort() {
	js.Rewrite("$<.abort()")
}

func (*IDBTransaction) ObjectStore(name string) (i *IDBObjectStore) {
	js.Rewrite("$<.objectStore($1)", name)
	return i
}

func (*IDBTransaction) GetDb() (db *IDBDatabase) {
	js.Rewrite("$<.db")
	return db
}

func (*IDBTransaction) GetError() (err *DOMError) {
	js.Rewrite("$<.error")
	return err
}

func (*IDBTransaction) GetMode() (mode *IDBTransactionMode) {
	js.Rewrite("$<.mode")
	return mode
}

func (*IDBTransaction) GetOnabort() (abort *Event) {
	js.Rewrite("$<.onabort")
	return abort
}

func (*IDBTransaction) SetOnabort(abort *Event) {
	js.Rewrite("$<.onabort = $1", abort)
}

func (*IDBTransaction) GetOncomplete() (complete *Event) {
	js.Rewrite("$<.oncomplete")
	return complete
}

func (*IDBTransaction) SetOncomplete(complete *Event) {
	js.Rewrite("$<.oncomplete = $1", complete)
}

func (*IDBTransaction) GetOnerror() (err *Event) {
	js.Rewrite("$<.onerror")
	return err
}

func (*IDBTransaction) SetOnerror(err *Event) {
	js.Rewrite("$<.onerror = $1", err)
}

type IDBVersionChangeEvent struct {
	*Event
}

func (*IDBVersionChangeEvent) GetNewVersion() (newVersion uint64) {
	js.Rewrite("$<.newVersion")
	return newVersion
}

func (*IDBVersionChangeEvent) GetOldVersion() (oldVersion uint64) {
	js.Rewrite("$<.oldVersion")
	return oldVersion
}

type IIRFilterNode struct {
	*AudioNode
}

func (*IIRFilterNode) GetFrequencyResponse(frequencyHz []float32, magResponse []float32, phaseResponse []float32) {
	js.Rewrite("$<.getFrequencyResponse($1, $2, $3)", frequencyHz, magResponse, phaseResponse)
}

type ImageData struct {
}

func (*ImageData) GetData() (data []uint8) {
	js.Rewrite("$<.data")
	return data
}

func (*ImageData) GetHeight() (height uint) {
	js.Rewrite("$<.height")
	return height
}

func (*ImageData) GetWidth() (width uint) {
	js.Rewrite("$<.width")
	return width
}

type IntersectionObserver struct {
}

func (*IntersectionObserver) Disconnect() {
	js.Rewrite("$<.disconnect()")
}

func (*IntersectionObserver) Observe(target *Element) {
	js.Rewrite("$<.observe($1)", target)
}

func (*IntersectionObserver) TakeRecords() (i []*IntersectionObserverEntry) {
	js.Rewrite("$<.takeRecords()")
	return i
}

func (*IntersectionObserver) Unobserve(target *Element) {
	js.Rewrite("$<.unobserve($1)", target)
}

func (*IntersectionObserver) GetRoot() (root *Element) {
	js.Rewrite("$<.root")
	return root
}

func (*IntersectionObserver) GetRootMargin() (rootMargin string) {
	js.Rewrite("$<.rootMargin")
	return rootMargin
}

func (*IntersectionObserver) GetThresholds() (thresholds []float32) {
	js.Rewrite("$<.thresholds")
	return thresholds
}

type IntersectionObserverEntry struct {
}

func (*IntersectionObserverEntry) GetBoundingClientRect() (boundingClientRect *ClientRect) {
	js.Rewrite("$<.boundingClientRect")
	return boundingClientRect
}

func (*IntersectionObserverEntry) GetIntersectionRatio() (intersectionRatio float32) {
	js.Rewrite("$<.intersectionRatio")
	return intersectionRatio
}

func (*IntersectionObserverEntry) GetIntersectionRect() (intersectionRect *ClientRect) {
	js.Rewrite("$<.intersectionRect")
	return intersectionRect
}

func (*IntersectionObserverEntry) GetRootBounds() (rootBounds *ClientRect) {
	js.Rewrite("$<.rootBounds")
	return rootBounds
}

func (*IntersectionObserverEntry) GetTarget() (target *Element) {
	js.Rewrite("$<.target")
	return target
}

func (*IntersectionObserverEntry) GetTime() (time int) {
	js.Rewrite("$<.time")
	return time
}

type KeyboardEvent struct {
	*UIEvent
}

func (*KeyboardEvent) GetModifierState(keyArg string) (b bool) {
	js.Rewrite("$<.getModifierState($1)", keyArg)
	return b
}

func (*KeyboardEvent) InitKeyboardEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, keyArg string, locationArg uint, modifiersListArg string, repeat bool, locale string) {
	js.Rewrite("$<.initKeyboardEvent($1, $2, $3, $4, $5, $6, $7, $8, $9)", typeArg, canBubbleArg, cancelableArg, viewArg, keyArg, locationArg, modifiersListArg, repeat, locale)
}

func (*KeyboardEvent) GetAltKey() (altKey bool) {
	js.Rewrite("$<.altKey")
	return altKey
}

func (*KeyboardEvent) GetChar() (char string) {
	js.Rewrite("$<.char")
	return char
}

func (*KeyboardEvent) GetCharCode() (charCode int8) {
	js.Rewrite("$<.charCode")
	return charCode
}

func (*KeyboardEvent) GetCtrlKey() (ctrlKey bool) {
	js.Rewrite("$<.ctrlKey")
	return ctrlKey
}

func (*KeyboardEvent) GetKey() (key string) {
	js.Rewrite("$<.key")
	return key
}

func (*KeyboardEvent) GetKeyCode() (keyCode int8) {
	js.Rewrite("$<.keyCode")
	return keyCode
}

func (*KeyboardEvent) GetLocale() (locale string) {
	js.Rewrite("$<.locale")
	return locale
}

func (*KeyboardEvent) GetLocation() (location uint) {
	js.Rewrite("$<.location")
	return location
}

func (*KeyboardEvent) GetMetaKey() (metaKey bool) {
	js.Rewrite("$<.metaKey")
	return metaKey
}

func (*KeyboardEvent) GetRepeat() (repeat bool) {
	js.Rewrite("$<.repeat")
	return repeat
}

func (*KeyboardEvent) GetShiftKey() (shiftKey bool) {
	js.Rewrite("$<.shiftKey")
	return shiftKey
}

func (*KeyboardEvent) GetWhich() (which int8) {
	js.Rewrite("$<.which")
	return which
}

type ListeningStateChangedEvent struct {
	*Event
}

func (*ListeningStateChangedEvent) GetLabel() (label string) {
	js.Rewrite("$<.label")
	return label
}

func (*ListeningStateChangedEvent) GetState() (state *ListeningState) {
	js.Rewrite("$<.state")
	return state
}

type Location struct {
}

func (*Location) Assign(url string) {
	js.Rewrite("$<.assign($1)", url)
}

func (*Location) Reload(forcedReload bool) {
	js.Rewrite("$<.reload($1)", forcedReload)
}

func (*Location) Replace(url string) {
	js.Rewrite("$<.replace($1)", url)
}

func (*Location) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

func (*Location) GetHash() (hash string) {
	js.Rewrite("$<.hash")
	return hash
}

func (*Location) SetHash(hash string) {
	js.Rewrite("$<.hash = $1", hash)
}

func (*Location) GetHost() (host string) {
	js.Rewrite("$<.host")
	return host
}

func (*Location) SetHost(host string) {
	js.Rewrite("$<.host = $1", host)
}

func (*Location) GetHostname() (hostname string) {
	js.Rewrite("$<.hostname")
	return hostname
}

func (*Location) SetHostname(hostname string) {
	js.Rewrite("$<.hostname = $1", hostname)
}

func (*Location) GetHref() (href string) {
	js.Rewrite("$<.href")
	return href
}

func (*Location) SetHref(href string) {
	js.Rewrite("$<.href = $1", href)
}

func (*Location) GetOrigin() (origin string) {
	js.Rewrite("$<.origin")
	return origin
}

func (*Location) GetPathname() (pathname string) {
	js.Rewrite("$<.pathname")
	return pathname
}

func (*Location) SetPathname(pathname string) {
	js.Rewrite("$<.pathname = $1", pathname)
}

func (*Location) GetPort() (port string) {
	js.Rewrite("$<.port")
	return port
}

func (*Location) SetPort(port string) {
	js.Rewrite("$<.port = $1", port)
}

func (*Location) GetProtocol() (protocol string) {
	js.Rewrite("$<.protocol")
	return protocol
}

func (*Location) SetProtocol(protocol string) {
	js.Rewrite("$<.protocol = $1", protocol)
}

func (*Location) GetSearch() (search string) {
	js.Rewrite("$<.search")
	return search
}

func (*Location) SetSearch(search string) {
	js.Rewrite("$<.search = $1", search)
}

type LongRunningScriptDetectedEvent struct {
	*Event
}

func (*LongRunningScriptDetectedEvent) GetExecutionTime() (executionTime int) {
	js.Rewrite("$<.executionTime")
	return executionTime
}

func (*LongRunningScriptDetectedEvent) GetStopPageScriptExecution() (stopPageScriptExecution bool) {
	js.Rewrite("$<.stopPageScriptExecution")
	return stopPageScriptExecution
}

func (*LongRunningScriptDetectedEvent) SetStopPageScriptExecution(stopPageScriptExecution bool) {
	js.Rewrite("$<.stopPageScriptExecution = $1", stopPageScriptExecution)
}

type MediaDeviceInfo struct {
}

func (*MediaDeviceInfo) GetDeviceID() (deviceId string) {
	js.Rewrite("$<.deviceId")
	return deviceId
}

func (*MediaDeviceInfo) GetGroupID() (groupId string) {
	js.Rewrite("$<.groupId")
	return groupId
}

func (*MediaDeviceInfo) GetKind() (kind *MediaDeviceKind) {
	js.Rewrite("$<.kind")
	return kind
}

func (*MediaDeviceInfo) GetLabel() (label string) {
	js.Rewrite("$<.label")
	return label
}

type MediaDevices struct {
	*EventTarget
}

func (*MediaDevices) EnumerateDevices() (m []*MediaDeviceInfo) {
	js.Rewrite("await $<.enumerateDevices()")
	return m
}

func (*MediaDevices) GetSupportedConstraints() (m *MediaTrackSupportedConstraints) {
	js.Rewrite("$<.getSupportedConstraints()")
	return m
}

func (*MediaDevices) GetUserMedia(constraints *MediaStreamConstraints) (m *MediaStream) {
	js.Rewrite("await $<.getUserMedia($1)", constraints)
	return m
}

func (*MediaDevices) GetOndevicechange() (devicechange *Event) {
	js.Rewrite("$<.ondevicechange")
	return devicechange
}

func (*MediaDevices) SetOndevicechange(devicechange *Event) {
	js.Rewrite("$<.ondevicechange = $1", devicechange)
}

type MediaElementAudioSourceNode struct {
	*AudioNode
}

type MediaEncryptedEvent struct {
	*Event
}

func (*MediaEncryptedEvent) GetInitData() (initData []byte) {
	js.Rewrite("$<.initData")
	return initData
}

func (*MediaEncryptedEvent) GetInitDataType() (initDataType string) {
	js.Rewrite("$<.initDataType")
	return initDataType
}

type MediaError struct {
}

func (*MediaError) GetCode() (code int8) {
	js.Rewrite("$<.code")
	return code
}

func (*MediaError) GetMsExtendedCode() (msExtendedCode int) {
	js.Rewrite("$<.msExtendedCode")
	return msExtendedCode
}

type MediaKeyMessageEvent struct {
	*Event
}

func (*MediaKeyMessageEvent) GetMessage() (message []byte) {
	js.Rewrite("$<.message")
	return message
}

func (*MediaKeyMessageEvent) GetMessageType() (messageType *MediaKeyMessageType) {
	js.Rewrite("$<.messageType")
	return messageType
}

type MediaKeys struct {
}

func (*MediaKeys) CreateSession(sessionType *MediaKeySessionType) (m *MediaKeySession) {
	js.Rewrite("$<.createSession($1)", sessionType)
	return m
}

func (*MediaKeys) SetServerCertificate(serverCertificate []byte) {
	js.Rewrite("await $<.setServerCertificate($1)", serverCertificate)
}

type MediaKeySession struct {
	*EventTarget
}

func (*MediaKeySession) Close() {
	js.Rewrite("await $<.close()")
}

func (*MediaKeySession) GenerateRequest(initDataType string, initData []byte) {
	js.Rewrite("await $<.generateRequest($1, $2)", initDataType, initData)
}

func (*MediaKeySession) Load(sessionId string) (b bool) {
	js.Rewrite("await $<.load($1)", sessionId)
	return b
}

func (*MediaKeySession) Remove() {
	js.Rewrite("await $<.remove()")
}

func (*MediaKeySession) Update(response []byte) {
	js.Rewrite("await $<.update($1)", response)
}

func (*MediaKeySession) GetClosed() {
	js.Rewrite("await $<.closed")
}

func (*MediaKeySession) GetExpiration() (expiration float32) {
	js.Rewrite("$<.expiration")
	return expiration
}

func (*MediaKeySession) GetKeyStatuses() (keyStatuses *MediaKeyStatusMap) {
	js.Rewrite("$<.keyStatuses")
	return keyStatuses
}

func (*MediaKeySession) GetSessionID() (sessionId string) {
	js.Rewrite("$<.sessionId")
	return sessionId
}

type MediaKeyStatusMap struct {
}

func (*MediaKeyStatusMap) ForEach(callback func(keyId []byte, status *MediaKeyStatus)) {
	js.Rewrite("$<.forEach($1)", callback)
}

func (*MediaKeyStatusMap) Get(keyId []byte) (m *MediaKeyStatus) {
	js.Rewrite("$<.get($1)", keyId)
	return m
}

func (*MediaKeyStatusMap) Has(keyId []byte) (b bool) {
	js.Rewrite("$<.has($1)", keyId)
	return b
}

func (*MediaKeyStatusMap) GetSize() (size uint) {
	js.Rewrite("$<.size")
	return size
}

type MediaKeySystemAccess struct {
}

func (*MediaKeySystemAccess) CreateMediaKeys() (m *MediaKeys) {
	js.Rewrite("await $<.createMediaKeys()")
	return m
}

func (*MediaKeySystemAccess) GetConfiguration() (m *MediaKeySystemConfiguration) {
	js.Rewrite("$<.getConfiguration()")
	return m
}

func (*MediaKeySystemAccess) GetKeySystem() (keySystem string) {
	js.Rewrite("$<.keySystem")
	return keySystem
}

type MediaList struct {
}

func (*MediaList) AppendMedium(newMedium string) {
	js.Rewrite("$<.appendMedium($1)", newMedium)
}

func (*MediaList) DeleteMedium(oldMedium string) {
	js.Rewrite("$<.deleteMedium($1)", oldMedium)
}

func (*MediaList) Item(index uint) (s string) {
	js.Rewrite("$<.item($1)", index)
	return s
}

func (*MediaList) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

func (*MediaList) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

func (*MediaList) GetMediaText() (mediaText string) {
	js.Rewrite("$<.mediaText")
	return mediaText
}

func (*MediaList) SetMediaText(mediaText string) {
	js.Rewrite("$<.mediaText = $1", mediaText)
}

type MediaQueryList struct {
}

func (*MediaQueryList) AddListener(listener func(mql *MediaQueryList)) {
	js.Rewrite("$<.addListener($1)", listener)
}

func (*MediaQueryList) RemoveListener(listener func(mql *MediaQueryList)) {
	js.Rewrite("$<.removeListener($1)", listener)
}

func (*MediaQueryList) GetMatches() (matches bool) {
	js.Rewrite("$<.matches")
	return matches
}

func (*MediaQueryList) GetMedia() (media string) {
	js.Rewrite("$<.media")
	return media
}

type MediaSource struct {
	*EventTarget
}

func (*MediaSource) AddSourceBuffer(kind string) (s *SourceBuffer) {
	js.Rewrite("$<.addSourceBuffer($1)", kind)
	return s
}

func (*MediaSource) EndOfStream(err string) {
	js.Rewrite("$<.endOfStream($1)", err)
}

func (*MediaSource) IsTypeSupported(kind string) (b bool) {
	js.Rewrite("$<.isTypeSupported($1)", kind)
	return b
}

func (*MediaSource) RemoveSourceBuffer(sourceBuffer *SourceBuffer) {
	js.Rewrite("$<.removeSourceBuffer($1)", sourceBuffer)
}

func (*MediaSource) GetActiveSourceBuffers() (activeSourceBuffers *SourceBufferList) {
	js.Rewrite("$<.activeSourceBuffers")
	return activeSourceBuffers
}

func (*MediaSource) GetDuration() (duration float32) {
	js.Rewrite("$<.duration")
	return duration
}

func (*MediaSource) SetDuration(duration float32) {
	js.Rewrite("$<.duration = $1", duration)
}

func (*MediaSource) GetReadyState() (readyState string) {
	js.Rewrite("$<.readyState")
	return readyState
}

func (*MediaSource) GetSourceBuffers() (sourceBuffers *SourceBufferList) {
	js.Rewrite("$<.sourceBuffers")
	return sourceBuffers
}

type MediaStream struct {
	*EventTarget
}

func (*MediaStream) AddTrack(track *MediaStreamTrack) {
	js.Rewrite("$<.addTrack($1)", track)
}

func (*MediaStream) Clone() (m *MediaStream) {
	js.Rewrite("$<.clone()")
	return m
}

func (*MediaStream) GetAudioTracks() (m []*MediaStreamTrack) {
	js.Rewrite("$<.getAudioTracks()")
	return m
}

func (*MediaStream) GetTrackByID(trackId string) (m *MediaStreamTrack) {
	js.Rewrite("$<.getTrackById($1)", trackId)
	return m
}

func (*MediaStream) GetTracks() (m []*MediaStreamTrack) {
	js.Rewrite("$<.getTracks()")
	return m
}

func (*MediaStream) GetVideoTracks() (m []*MediaStreamTrack) {
	js.Rewrite("$<.getVideoTracks()")
	return m
}

func (*MediaStream) RemoveTrack(track *MediaStreamTrack) {
	js.Rewrite("$<.removeTrack($1)", track)
}

func (*MediaStream) Stop() {
	js.Rewrite("$<.stop()")
}

func (*MediaStream) GetActive() (active bool) {
	js.Rewrite("$<.active")
	return active
}

func (*MediaStream) GetID() (id string) {
	js.Rewrite("$<.id")
	return id
}

func (*MediaStream) GetOnactive() (active *Event) {
	js.Rewrite("$<.onactive")
	return active
}

func (*MediaStream) SetOnactive(active *Event) {
	js.Rewrite("$<.onactive = $1", active)
}

func (*MediaStream) GetOnaddtrack() (addtrack *MediaStreamTrackEvent) {
	js.Rewrite("$<.onaddtrack")
	return addtrack
}

func (*MediaStream) SetOnaddtrack(addtrack *MediaStreamTrackEvent) {
	js.Rewrite("$<.onaddtrack = $1", addtrack)
}

func (*MediaStream) GetOninactive() (inactive *Event) {
	js.Rewrite("$<.oninactive")
	return inactive
}

func (*MediaStream) SetOninactive(inactive *Event) {
	js.Rewrite("$<.oninactive = $1", inactive)
}

func (*MediaStream) GetOnremovetrack() (removetrack *MediaStreamTrackEvent) {
	js.Rewrite("$<.onremovetrack")
	return removetrack
}

func (*MediaStream) SetOnremovetrack(removetrack *MediaStreamTrackEvent) {
	js.Rewrite("$<.onremovetrack = $1", removetrack)
}

type MediaStreamAudioSourceNode struct {
	*AudioNode
}

type MediaStreamError struct {
}

func (*MediaStreamError) GetConstraintName() (constraintName string) {
	js.Rewrite("$<.constraintName")
	return constraintName
}

func (*MediaStreamError) GetMessage() (message string) {
	js.Rewrite("$<.message")
	return message
}

func (*MediaStreamError) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

type MediaStreamErrorEvent struct {
	*Event
}

func (*MediaStreamErrorEvent) GetError() (err *MediaStreamError) {
	js.Rewrite("$<.error")
	return err
}

type MediaStreamEvent struct {
	*Event
}

func (*MediaStreamEvent) GetStream() (stream *MediaStream) {
	js.Rewrite("$<.stream")
	return stream
}

type MediaStreamTrack struct {
	*EventTarget
}

func (*MediaStreamTrack) ApplyConstraints(constraints *MediaTrackConstraints) {
	js.Rewrite("await $<.applyConstraints($1)", constraints)
}

func (*MediaStreamTrack) Clone() (m *MediaStreamTrack) {
	js.Rewrite("$<.clone()")
	return m
}

func (*MediaStreamTrack) GetCapabilities() (m *MediaTrackCapabilities) {
	js.Rewrite("$<.getCapabilities()")
	return m
}

func (*MediaStreamTrack) GetConstraints() (m *MediaTrackConstraints) {
	js.Rewrite("$<.getConstraints()")
	return m
}

func (*MediaStreamTrack) GetSettings() (m *MediaTrackSettings) {
	js.Rewrite("$<.getSettings()")
	return m
}

func (*MediaStreamTrack) Stop() {
	js.Rewrite("$<.stop()")
}

func (*MediaStreamTrack) GetEnabled() (enabled bool) {
	js.Rewrite("$<.enabled")
	return enabled
}

func (*MediaStreamTrack) SetEnabled(enabled bool) {
	js.Rewrite("$<.enabled = $1", enabled)
}

func (*MediaStreamTrack) GetID() (id string) {
	js.Rewrite("$<.id")
	return id
}

func (*MediaStreamTrack) GetKind() (kind string) {
	js.Rewrite("$<.kind")
	return kind
}

func (*MediaStreamTrack) GetLabel() (label string) {
	js.Rewrite("$<.label")
	return label
}

func (*MediaStreamTrack) GetMuted() (muted bool) {
	js.Rewrite("$<.muted")
	return muted
}

func (*MediaStreamTrack) GetOnended() (ended *MediaStreamErrorEvent) {
	js.Rewrite("$<.onended")
	return ended
}

func (*MediaStreamTrack) SetOnended(ended *MediaStreamErrorEvent) {
	js.Rewrite("$<.onended = $1", ended)
}

func (*MediaStreamTrack) GetOnmute() (mute *Event) {
	js.Rewrite("$<.onmute")
	return mute
}

func (*MediaStreamTrack) SetOnmute(mute *Event) {
	js.Rewrite("$<.onmute = $1", mute)
}

func (*MediaStreamTrack) GetOnoverconstrained() (overconstrained *MediaStreamErrorEvent) {
	js.Rewrite("$<.onoverconstrained")
	return overconstrained
}

func (*MediaStreamTrack) SetOnoverconstrained(overconstrained *MediaStreamErrorEvent) {
	js.Rewrite("$<.onoverconstrained = $1", overconstrained)
}

func (*MediaStreamTrack) GetOnunmute() (unmute *Event) {
	js.Rewrite("$<.onunmute")
	return unmute
}

func (*MediaStreamTrack) SetOnunmute(unmute *Event) {
	js.Rewrite("$<.onunmute = $1", unmute)
}

func (*MediaStreamTrack) GetReadonly() (readonly bool) {
	js.Rewrite("$<.readonly")
	return readonly
}

func (*MediaStreamTrack) GetReadyState() (readyState *MediaStreamTrackState) {
	js.Rewrite("$<.readyState")
	return readyState
}

func (*MediaStreamTrack) GetRemote() (remote bool) {
	js.Rewrite("$<.remote")
	return remote
}

type MediaStreamTrackEvent struct {
	*Event
}

func (*MediaStreamTrackEvent) GetTrack() (track *MediaStreamTrack) {
	js.Rewrite("$<.track")
	return track
}

type MessageChannel struct {
}

func (*MessageChannel) GetPort1() (port1 *MessagePort) {
	js.Rewrite("$<.port1")
	return port1
}

func (*MessageChannel) GetPort2() (port2 *MessagePort) {
	js.Rewrite("$<.port2")
	return port2
}

type MessageEvent struct {
	*Event
}

func (*MessageEvent) InitMessageEvent(typeArg string, canBubbleArg bool, cancelableArg bool, dataArg interface{}, originArg string, lastEventIdArg string, sourceArg *Window) {
	js.Rewrite("$<.initMessageEvent($1, $2, $3, $4, $5, $6, $7)", typeArg, canBubbleArg, cancelableArg, dataArg, originArg, lastEventIdArg, sourceArg)
}

func (*MessageEvent) GetData() (data interface{}) {
	js.Rewrite("$<.data")
	return data
}

func (*MessageEvent) GetOrigin() (origin string) {
	js.Rewrite("$<.origin")
	return origin
}

func (*MessageEvent) GetPorts() (ports interface{}) {
	js.Rewrite("$<.ports")
	return ports
}

func (*MessageEvent) GetSource() (source *Window) {
	js.Rewrite("$<.source")
	return source
}

type MessagePort struct {
	*EventTarget
}

func (*MessagePort) Close() {
	js.Rewrite("$<.close()")
}

func (*MessagePort) PostMessage(message interface{}, transfer []interface{}) {
	js.Rewrite("$<.postMessage($1, $2)", message, transfer)
}

func (*MessagePort) Start() {
	js.Rewrite("$<.start()")
}

func (*MessagePort) GetOnmessage() (message *MessageEvent) {
	js.Rewrite("$<.onmessage")
	return message
}

func (*MessagePort) SetOnmessage(message *MessageEvent) {
	js.Rewrite("$<.onmessage = $1", message)
}

type MimeType struct {
}

func (*MimeType) GetDescription() (description string) {
	js.Rewrite("$<.description")
	return description
}

func (*MimeType) GetEnabledPlugin() (enabledPlugin *Plugin) {
	js.Rewrite("$<.enabledPlugin")
	return enabledPlugin
}

func (*MimeType) GetSuffixes() (suffixes string) {
	js.Rewrite("$<.suffixes")
	return suffixes
}

func (*MimeType) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

type MimeTypeArray struct {
}

func (*MimeTypeArray) Item(index uint) (p *Plugin) {
	js.Rewrite("$<.item($1)", index)
	return p
}

func (*MimeTypeArray) NamedItem(kind string) (p *Plugin) {
	js.Rewrite("$<.namedItem($1)", kind)
	return p
}

func (*MimeTypeArray) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

type MouseEvent struct {
	*UIEvent
}

func (*MouseEvent) GetModifierState(keyArg string) (b bool) {
	js.Rewrite("$<.getModifierState($1)", keyArg)
	return b
}

func (*MouseEvent) InitMouseEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, screenXArg int, screenYArg int, clientXArg int, clientYArg int, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg uint8, relatedTargetArg *EventTarget) {
	js.Rewrite("$<.initMouseEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, ctrlKeyArg, altKeyArg, shiftKeyArg, metaKeyArg, buttonArg, relatedTargetArg)
}

func (*MouseEvent) GetAltKey() (altKey bool) {
	js.Rewrite("$<.altKey")
	return altKey
}

func (*MouseEvent) GetButton() (button uint8) {
	js.Rewrite("$<.button")
	return button
}

func (*MouseEvent) GetButtons() (buttons uint8) {
	js.Rewrite("$<.buttons")
	return buttons
}

func (*MouseEvent) GetClientX() (clientX int) {
	js.Rewrite("$<.clientX")
	return clientX
}

func (*MouseEvent) GetClientY() (clientY int) {
	js.Rewrite("$<.clientY")
	return clientY
}

func (*MouseEvent) GetCtrlKey() (ctrlKey bool) {
	js.Rewrite("$<.ctrlKey")
	return ctrlKey
}

func (*MouseEvent) GetFromElement() (fromElement *Element) {
	js.Rewrite("$<.fromElement")
	return fromElement
}

func (*MouseEvent) GetLayerX() (layerX int) {
	js.Rewrite("$<.layerX")
	return layerX
}

func (*MouseEvent) GetLayerY() (layerY int) {
	js.Rewrite("$<.layerY")
	return layerY
}

func (*MouseEvent) GetMetaKey() (metaKey bool) {
	js.Rewrite("$<.metaKey")
	return metaKey
}

func (*MouseEvent) GetMovementX() (movementX int) {
	js.Rewrite("$<.movementX")
	return movementX
}

func (*MouseEvent) GetMovementY() (movementY int) {
	js.Rewrite("$<.movementY")
	return movementY
}

func (*MouseEvent) GetOffsetX() (offsetX int) {
	js.Rewrite("$<.offsetX")
	return offsetX
}

func (*MouseEvent) GetOffsetY() (offsetY int) {
	js.Rewrite("$<.offsetY")
	return offsetY
}

func (*MouseEvent) GetPageX() (pageX int) {
	js.Rewrite("$<.pageX")
	return pageX
}

func (*MouseEvent) GetPageY() (pageY int) {
	js.Rewrite("$<.pageY")
	return pageY
}

func (*MouseEvent) GetRelatedTarget() (relatedTarget *EventTarget) {
	js.Rewrite("$<.relatedTarget")
	return relatedTarget
}

func (*MouseEvent) GetScreenX() (screenX int) {
	js.Rewrite("$<.screenX")
	return screenX
}

func (*MouseEvent) GetScreenY() (screenY int) {
	js.Rewrite("$<.screenY")
	return screenY
}

func (*MouseEvent) GetShiftKey() (shiftKey bool) {
	js.Rewrite("$<.shiftKey")
	return shiftKey
}

func (*MouseEvent) GetToElement() (toElement *Element) {
	js.Rewrite("$<.toElement")
	return toElement
}

func (*MouseEvent) GetWhich() (which uint8) {
	js.Rewrite("$<.which")
	return which
}

func (*MouseEvent) GetX() (x int) {
	js.Rewrite("$<.x")
	return x
}

func (*MouseEvent) GetY() (y int) {
	js.Rewrite("$<.y")
	return y
}

type MSApp struct {
}

func (*MSApp) ClearTemporaryWebDataAsync() (m *MSAppAsyncOperation) {
	js.Rewrite("$<.clearTemporaryWebDataAsync()")
	return m
}

func (*MSApp) CreateBlobFromRandomAccessStream(kind string, seeker interface{}) (b *Blob) {
	js.Rewrite("$<.createBlobFromRandomAccessStream($1, $2)", kind, seeker)
	return b
}

func (*MSApp) CreateDataPackage(object interface{}) (i interface{}) {
	js.Rewrite("$<.createDataPackage($1)", object)
	return i
}

func (*MSApp) CreateDataPackageFromSelection() (i interface{}) {
	js.Rewrite("$<.createDataPackageFromSelection()")
	return i
}

func (*MSApp) CreateFileFromStorageFile(storageFile interface{}) (f *File) {
	js.Rewrite("$<.createFileFromStorageFile($1)", storageFile)
	return f
}

func (*MSApp) CreateStreamFromInputStream(kind string, inputStream interface{}) (m *MSStream) {
	js.Rewrite("$<.createStreamFromInputStream($1, $2)", kind, inputStream)
	return m
}

func (*MSApp) ExecAsyncAtPriority(asynchronousCallback func(args interface{}) interface{}, priority string, args interface{}) {
	js.Rewrite("$<.execAsyncAtPriority($1, $2, $3)", asynchronousCallback, priority, args)
}

func (*MSApp) ExecAtPriority(synchronousCallback func(args interface{}) interface{}, priority string, args interface{}) (i interface{}) {
	js.Rewrite("$<.execAtPriority($1, $2, $3)", synchronousCallback, priority, args)
	return i
}

func (*MSApp) GetCurrentPriority() (s string) {
	js.Rewrite("$<.getCurrentPriority()")
	return s
}

func (*MSApp) GetHTMLPrintDocumentSourceAsync(htmlDoc interface{}) (i interface{}) {
	js.Rewrite("await $<.getHtmlPrintDocumentSourceAsync($1)", htmlDoc)
	return i
}

func (*MSApp) GetViewID(view interface{}) (i interface{}) {
	js.Rewrite("$<.getViewId($1)", view)
	return i
}

func (*MSApp) IsTaskScheduledAtPriorityOrHigher(priority string) (b bool) {
	js.Rewrite("$<.isTaskScheduledAtPriorityOrHigher($1)", priority)
	return b
}

func (*MSApp) PageHandlesAllApplicationActivations(enabled bool) {
	js.Rewrite("$<.pageHandlesAllApplicationActivations($1)", enabled)
}

func (*MSApp) SuppressSubdownloadCredentialPrompts(suppress bool) {
	js.Rewrite("$<.suppressSubdownloadCredentialPrompts($1)", suppress)
}

func (*MSApp) TerminateApp(exceptionObject interface{}) {
	js.Rewrite("$<.terminateApp($1)", exceptionObject)
}

type MSAppAsyncOperation struct {
	*EventTarget
}

func (*MSAppAsyncOperation) Start() {
	js.Rewrite("$<.start()")
}

func (*MSAppAsyncOperation) GetError() (err *DOMError) {
	js.Rewrite("$<.error")
	return err
}

func (*MSAppAsyncOperation) GetOncomplete() (complete *Event) {
	js.Rewrite("$<.oncomplete")
	return complete
}

func (*MSAppAsyncOperation) SetOncomplete(complete *Event) {
	js.Rewrite("$<.oncomplete = $1", complete)
}

func (*MSAppAsyncOperation) GetOnerror() (err *Event) {
	js.Rewrite("$<.onerror")
	return err
}

func (*MSAppAsyncOperation) SetOnerror(err *Event) {
	js.Rewrite("$<.onerror = $1", err)
}

func (*MSAppAsyncOperation) GetReadyState() (readyState uint8) {
	js.Rewrite("$<.readyState")
	return readyState
}

func (*MSAppAsyncOperation) GetResult() (result interface{}) {
	js.Rewrite("$<.result")
	return result
}

type MSAssertion struct {
}

func (*MSAssertion) GetID() (id string) {
	js.Rewrite("$<.id")
	return id
}

func (*MSAssertion) GetType() (kind *MSCredentialType) {
	js.Rewrite("$<.type")
	return kind
}

type MSBlobBuilder struct {
}

func (*MSBlobBuilder) Append(data interface{}, endings string) {
	js.Rewrite("$<.append($1, $2)", data, endings)
}

func (*MSBlobBuilder) GetBlob(contentType string) (b *Blob) {
	js.Rewrite("$<.getBlob($1)", contentType)
	return b
}

type MSCredentials struct {
}

func (*MSCredentials) GetAssertion(challenge string, filter *MSCredentialFilter, params *MSSignatureParameters) (m *MSAssertion) {
	js.Rewrite("await $<.getAssertion($1, $2, $3)", challenge, filter, params)
	return m
}

func (*MSCredentials) MakeCredential(accountInfo *MSAccountInfo, params []*MSCredentialParameters, challenge string) (m *MSAssertion) {
	js.Rewrite("await $<.makeCredential($1, $2, $3)", accountInfo, params, challenge)
	return m
}

type MSFIDOCredentialAssertion struct {
	*MSAssertion
}

func (*MSFIDOCredentialAssertion) GetAlgorithm() (algorithm interface{}) {
	js.Rewrite("$<.algorithm")
	return algorithm
}

func (*MSFIDOCredentialAssertion) GetAttestation() (attestation bool) {
	js.Rewrite("$<.attestation")
	return attestation
}

func (*MSFIDOCredentialAssertion) GetPublicKey() (publicKey string) {
	js.Rewrite("$<.publicKey")
	return publicKey
}

func (*MSFIDOCredentialAssertion) GetTransportHints() (transportHints []*MSTransportType) {
	js.Rewrite("$<.transportHints")
	return transportHints
}

type MSFIDOSignature struct {
}

func (*MSFIDOSignature) GetAuthnrData() (authnrData string) {
	js.Rewrite("$<.authnrData")
	return authnrData
}

func (*MSFIDOSignature) GetClientData() (clientData string) {
	js.Rewrite("$<.clientData")
	return clientData
}

func (*MSFIDOSignature) GetSignature() (signature string) {
	js.Rewrite("$<.signature")
	return signature
}

type MSFIDOSignatureAssertion struct {
	*MSAssertion
}

func (*MSFIDOSignatureAssertion) GetSignature() (signature *MSFIDOSignature) {
	js.Rewrite("$<.signature")
	return signature
}

type MSGesture struct {
}

func (*MSGesture) AddPointer(pointerId int) {
	js.Rewrite("$<.addPointer($1)", pointerId)
}

func (*MSGesture) Stop() {
	js.Rewrite("$<.stop()")
}

func (*MSGesture) GetTarget() (target *Element) {
	js.Rewrite("$<.target")
	return target
}

func (*MSGesture) SetTarget(target *Element) {
	js.Rewrite("$<.target = $1", target)
}

type MSGestureEvent struct {
	*UIEvent
}

func (*MSGestureEvent) InitGestureEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, screenXArg int, screenYArg int, clientXArg float32, clientYArg float32, offsetXArg float32, offsetYArg float32, translationXArg float32, translationYArg float32, scaleArg float32, expansionArg float32, rotationArg float32, velocityXArg float32, velocityYArg float32, velocityExpansionArg float32, velocityAngularArg float32, hwTimestampArg uint64) {
	js.Rewrite("$<.initGestureEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, offsetXArg, offsetYArg, translationXArg, translationYArg, scaleArg, expansionArg, rotationArg, velocityXArg, velocityYArg, velocityExpansionArg, velocityAngularArg, hwTimestampArg)
}

func (*MSGestureEvent) GetClientX() (clientX float32) {
	js.Rewrite("$<.clientX")
	return clientX
}

func (*MSGestureEvent) GetClientY() (clientY float32) {
	js.Rewrite("$<.clientY")
	return clientY
}

func (*MSGestureEvent) GetExpansion() (expansion float32) {
	js.Rewrite("$<.expansion")
	return expansion
}

func (*MSGestureEvent) GetGestureObject() (gestureObject interface{}) {
	js.Rewrite("$<.gestureObject")
	return gestureObject
}

func (*MSGestureEvent) GetHwTimestamp() (hwTimestamp uint64) {
	js.Rewrite("$<.hwTimestamp")
	return hwTimestamp
}

func (*MSGestureEvent) GetOffsetX() (offsetX float32) {
	js.Rewrite("$<.offsetX")
	return offsetX
}

func (*MSGestureEvent) GetOffsetY() (offsetY float32) {
	js.Rewrite("$<.offsetY")
	return offsetY
}

func (*MSGestureEvent) GetRotation() (rotation float32) {
	js.Rewrite("$<.rotation")
	return rotation
}

func (*MSGestureEvent) GetScale() (scale float32) {
	js.Rewrite("$<.scale")
	return scale
}

func (*MSGestureEvent) GetScreenX() (screenX int) {
	js.Rewrite("$<.screenX")
	return screenX
}

func (*MSGestureEvent) GetScreenY() (screenY int) {
	js.Rewrite("$<.screenY")
	return screenY
}

func (*MSGestureEvent) GetTranslationX() (translationX float32) {
	js.Rewrite("$<.translationX")
	return translationX
}

func (*MSGestureEvent) GetTranslationY() (translationY float32) {
	js.Rewrite("$<.translationY")
	return translationY
}

func (*MSGestureEvent) GetVelocityAngular() (velocityAngular float32) {
	js.Rewrite("$<.velocityAngular")
	return velocityAngular
}

func (*MSGestureEvent) GetVelocityExpansion() (velocityExpansion float32) {
	js.Rewrite("$<.velocityExpansion")
	return velocityExpansion
}

func (*MSGestureEvent) GetVelocityX() (velocityX float32) {
	js.Rewrite("$<.velocityX")
	return velocityX
}

func (*MSGestureEvent) GetVelocityY() (velocityY float32) {
	js.Rewrite("$<.velocityY")
	return velocityY
}

type MSGraphicsTrust struct {
}

func (*MSGraphicsTrust) GetConstrictionActive() (constrictionActive bool) {
	js.Rewrite("$<.constrictionActive")
	return constrictionActive
}

func (*MSGraphicsTrust) GetStatus() (status string) {
	js.Rewrite("$<.status")
	return status
}

type MSHTMLWebViewElement struct {
	*HTMLElement
}

func (*MSHTMLWebViewElement) AddWebAllowedObject(name string, applicationObject interface{}) {
	js.Rewrite("$<.addWebAllowedObject($1, $2)", name, applicationObject)
}

func (*MSHTMLWebViewElement) BuildLocalStreamURI(contentIdentifier string, relativePath string) (s string) {
	js.Rewrite("$<.buildLocalStreamUri($1, $2)", contentIdentifier, relativePath)
	return s
}

func (*MSHTMLWebViewElement) CapturePreviewToBlobAsync() (m *MSWebViewAsyncOperation) {
	js.Rewrite("$<.capturePreviewToBlobAsync()")
	return m
}

func (*MSHTMLWebViewElement) CaptureSelectedContentToDataPackageAsync() (m *MSWebViewAsyncOperation) {
	js.Rewrite("$<.captureSelectedContentToDataPackageAsync()")
	return m
}

func (*MSHTMLWebViewElement) GetDeferredPermissionRequestByID(id uint) (d *DeferredPermissionRequest) {
	js.Rewrite("$<.getDeferredPermissionRequestById($1)", id)
	return d
}

func (*MSHTMLWebViewElement) GetDeferredPermissionRequests() (d []*DeferredPermissionRequest) {
	js.Rewrite("$<.getDeferredPermissionRequests()")
	return d
}

func (*MSHTMLWebViewElement) GoBack() {
	js.Rewrite("$<.goBack()")
}

func (*MSHTMLWebViewElement) GoForward() {
	js.Rewrite("$<.goForward()")
}

func (*MSHTMLWebViewElement) InvokeScriptAsync(scriptName string, args interface{}) (m *MSWebViewAsyncOperation) {
	js.Rewrite("$<.invokeScriptAsync($1, $2)", scriptName, args)
	return m
}

func (*MSHTMLWebViewElement) Navigate(uri string) {
	js.Rewrite("$<.navigate($1)", uri)
}

func (*MSHTMLWebViewElement) NavigateFocus(navigationReason *NavigationReason, origin *FocusNavigationOrigin) {
	js.Rewrite("$<.navigateFocus($1, $2)", navigationReason, origin)
}

func (*MSHTMLWebViewElement) NavigateToLocalStreamURI(source string, streamResolver interface{}) {
	js.Rewrite("$<.navigateToLocalStreamUri($1, $2)", source, streamResolver)
}

func (*MSHTMLWebViewElement) NavigateToString(contents string) {
	js.Rewrite("$<.navigateToString($1)", contents)
}

func (*MSHTMLWebViewElement) NavigateWithHTTPRequestMessage(requestMessage interface{}) {
	js.Rewrite("$<.navigateWithHttpRequestMessage($1)", requestMessage)
}

func (*MSHTMLWebViewElement) Refresh() {
	js.Rewrite("$<.refresh()")
}

func (*MSHTMLWebViewElement) Stop() {
	js.Rewrite("$<.stop()")
}

func (*MSHTMLWebViewElement) GetCanGoBack() (canGoBack bool) {
	js.Rewrite("$<.canGoBack")
	return canGoBack
}

func (*MSHTMLWebViewElement) GetCanGoForward() (canGoForward bool) {
	js.Rewrite("$<.canGoForward")
	return canGoForward
}

func (*MSHTMLWebViewElement) GetContainsFullScreenElement() (containsFullScreenElement bool) {
	js.Rewrite("$<.containsFullScreenElement")
	return containsFullScreenElement
}

func (*MSHTMLWebViewElement) GetDocumentTitle() (documentTitle string) {
	js.Rewrite("$<.documentTitle")
	return documentTitle
}

func (*MSHTMLWebViewElement) GetHeight() (height uint) {
	js.Rewrite("$<.height")
	return height
}

func (*MSHTMLWebViewElement) SetHeight(height uint) {
	js.Rewrite("$<.height = $1", height)
}

func (*MSHTMLWebViewElement) GetSettings() (settings *MSWebViewSettings) {
	js.Rewrite("$<.settings")
	return settings
}

func (*MSHTMLWebViewElement) GetSrc() (src string) {
	js.Rewrite("$<.src")
	return src
}

func (*MSHTMLWebViewElement) SetSrc(src string) {
	js.Rewrite("$<.src = $1", src)
}

func (*MSHTMLWebViewElement) GetWidth() (width uint) {
	js.Rewrite("$<.width")
	return width
}

func (*MSHTMLWebViewElement) SetWidth(width uint) {
	js.Rewrite("$<.width = $1", width)
}

type MSInputMethodContext struct {
	*EventTarget
}

func (*MSInputMethodContext) GetCandidateWindowClientRect() (c *ClientRect) {
	js.Rewrite("$<.getCandidateWindowClientRect()")
	return c
}

func (*MSInputMethodContext) GetCompositionAlternatives() (s []string) {
	js.Rewrite("$<.getCompositionAlternatives()")
	return s
}

func (*MSInputMethodContext) HasComposition() (b bool) {
	js.Rewrite("$<.hasComposition()")
	return b
}

func (*MSInputMethodContext) IsCandidateWindowVisible() (b bool) {
	js.Rewrite("$<.isCandidateWindowVisible()")
	return b
}

func (*MSInputMethodContext) GetCompositionEndOffset() (compositionEndOffset uint) {
	js.Rewrite("$<.compositionEndOffset")
	return compositionEndOffset
}

func (*MSInputMethodContext) GetCompositionStartOffset() (compositionStartOffset uint) {
	js.Rewrite("$<.compositionStartOffset")
	return compositionStartOffset
}

func (*MSInputMethodContext) GetOncandidatewindowhide() (MSCandidateWindowHide *Event) {
	js.Rewrite("$<.oncandidatewindowhide")
	return MSCandidateWindowHide
}

func (*MSInputMethodContext) SetOncandidatewindowhide(MSCandidateWindowHide *Event) {
	js.Rewrite("$<.oncandidatewindowhide = $1", MSCandidateWindowHide)
}

func (*MSInputMethodContext) GetOncandidatewindowshow() (MSCandidateWindowShow *Event) {
	js.Rewrite("$<.oncandidatewindowshow")
	return MSCandidateWindowShow
}

func (*MSInputMethodContext) SetOncandidatewindowshow(MSCandidateWindowShow *Event) {
	js.Rewrite("$<.oncandidatewindowshow = $1", MSCandidateWindowShow)
}

func (*MSInputMethodContext) GetOncandidatewindowupdate() (MSCandidateWindowUpdate *Event) {
	js.Rewrite("$<.oncandidatewindowupdate")
	return MSCandidateWindowUpdate
}

func (*MSInputMethodContext) SetOncandidatewindowupdate(MSCandidateWindowUpdate *Event) {
	js.Rewrite("$<.oncandidatewindowupdate = $1", MSCandidateWindowUpdate)
}

func (*MSInputMethodContext) GetTarget() (target *HTMLElement) {
	js.Rewrite("$<.target")
	return target
}

type MSManipulationEvent struct {
	*UIEvent
}

func (*MSManipulationEvent) InitMSManipulationEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, lastState int, currentState int) {
	js.Rewrite("$<.initMSManipulationEvent($1, $2, $3, $4, $5, $6, $7)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, lastState, currentState)
}

func (*MSManipulationEvent) GetCurrentState() (currentState int) {
	js.Rewrite("$<.currentState")
	return currentState
}

func (*MSManipulationEvent) GetInertiaDestinationX() (inertiaDestinationX int) {
	js.Rewrite("$<.inertiaDestinationX")
	return inertiaDestinationX
}

func (*MSManipulationEvent) GetInertiaDestinationY() (inertiaDestinationY int) {
	js.Rewrite("$<.inertiaDestinationY")
	return inertiaDestinationY
}

func (*MSManipulationEvent) GetLastState() (lastState int) {
	js.Rewrite("$<.lastState")
	return lastState
}

type MSMediaKeyError struct {
}

func (*MSMediaKeyError) GetCode() (code uint8) {
	js.Rewrite("$<.code")
	return code
}

func (*MSMediaKeyError) GetSystemCode() (systemCode uint) {
	js.Rewrite("$<.systemCode")
	return systemCode
}

type MSMediaKeyMessageEvent struct {
	*Event
}

func (*MSMediaKeyMessageEvent) GetDestinationURL() (destinationURL string) {
	js.Rewrite("$<.destinationURL")
	return destinationURL
}

func (*MSMediaKeyMessageEvent) GetMessage() (message []uint8) {
	js.Rewrite("$<.message")
	return message
}

type MSMediaKeyNeededEvent struct {
	*Event
}

func (*MSMediaKeyNeededEvent) GetInitData() (initData []uint8) {
	js.Rewrite("$<.initData")
	return initData
}

type MSMediaKeys struct {
}

func (*MSMediaKeys) CreateSession(kind string, initData []uint8, cdmData []uint8) (m *MSMediaKeySession) {
	js.Rewrite("$<.createSession($1, $2, $3)", kind, initData, cdmData)
	return m
}

func (*MSMediaKeys) IsTypeSupported(keySystem string, kind string) (b bool) {
	js.Rewrite("$<.isTypeSupported($1, $2)", keySystem, kind)
	return b
}

func (*MSMediaKeys) IsTypeSupportedWithFeatures(keySystem string, kind string) (s string) {
	js.Rewrite("$<.isTypeSupportedWithFeatures($1, $2)", keySystem, kind)
	return s
}

func (*MSMediaKeys) GetKeySystem() (keySystem string) {
	js.Rewrite("$<.keySystem")
	return keySystem
}

type MSMediaKeySession struct {
	*EventTarget
}

func (*MSMediaKeySession) Close() {
	js.Rewrite("$<.close()")
}

func (*MSMediaKeySession) Update(key []uint8) {
	js.Rewrite("$<.update($1)", key)
}

func (*MSMediaKeySession) GetError() (err *MSMediaKeyError) {
	js.Rewrite("$<.error")
	return err
}

func (*MSMediaKeySession) GetKeySystem() (keySystem string) {
	js.Rewrite("$<.keySystem")
	return keySystem
}

func (*MSMediaKeySession) GetSessionID() (sessionId string) {
	js.Rewrite("$<.sessionId")
	return sessionId
}

type MSPointerEvent struct {
	*MouseEvent
}

func (*MSPointerEvent) GetCurrentPoint(element *Element) {
	js.Rewrite("$<.getCurrentPoint($1)", element)
}

func (*MSPointerEvent) GetIntermediatePoints(element *Element) {
	js.Rewrite("$<.getIntermediatePoints($1)", element)
}

func (*MSPointerEvent) InitPointerEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, screenXArg int, screenYArg int, clientXArg float32, clientYArg float32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg uint8, relatedTargetArg *EventTarget, offsetXArg float32, offsetYArg float32, widthArg int, heightArg int, pressure float32, rotation int, tiltX int, tiltY int, pointerIdArg int, pointerType interface{}, hwTimestampArg uint64, isPrimary bool) {
	js.Rewrite("$<.initPointerEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, ctrlKeyArg, altKeyArg, shiftKeyArg, metaKeyArg, buttonArg, relatedTargetArg, offsetXArg, offsetYArg, widthArg, heightArg, pressure, rotation, tiltX, tiltY, pointerIdArg, pointerType, hwTimestampArg, isPrimary)
}

func (*MSPointerEvent) GetHeight() (height int) {
	js.Rewrite("$<.height")
	return height
}

func (*MSPointerEvent) GetHwTimestamp() (hwTimestamp uint64) {
	js.Rewrite("$<.hwTimestamp")
	return hwTimestamp
}

func (*MSPointerEvent) GetIsPrimary() (isPrimary bool) {
	js.Rewrite("$<.isPrimary")
	return isPrimary
}

func (*MSPointerEvent) GetPointerID() (pointerId int) {
	js.Rewrite("$<.pointerId")
	return pointerId
}

func (*MSPointerEvent) GetPointerType() (pointerType interface{}) {
	js.Rewrite("$<.pointerType")
	return pointerType
}

func (*MSPointerEvent) GetPressure() (pressure float32) {
	js.Rewrite("$<.pressure")
	return pressure
}

func (*MSPointerEvent) GetRotation() (rotation int) {
	js.Rewrite("$<.rotation")
	return rotation
}

func (*MSPointerEvent) GetTiltX() (tiltX int) {
	js.Rewrite("$<.tiltX")
	return tiltX
}

func (*MSPointerEvent) GetTiltY() (tiltY int) {
	js.Rewrite("$<.tiltY")
	return tiltY
}

func (*MSPointerEvent) GetWidth() (width int) {
	js.Rewrite("$<.width")
	return width
}

type MSRangeCollection struct {
}

func (*MSRangeCollection) Item(index uint) (r *Range) {
	js.Rewrite("$<.item($1)", index)
	return r
}

func (*MSRangeCollection) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

type MSSiteModeEvent struct {
	*Event
}

func (*MSSiteModeEvent) GetActionURL() (actionURL string) {
	js.Rewrite("$<.actionURL")
	return actionURL
}

func (*MSSiteModeEvent) GetButtonID() (buttonID int) {
	js.Rewrite("$<.buttonID")
	return buttonID
}

type MSStream struct {
}

func (*MSStream) MsClose() {
	js.Rewrite("$<.msClose()")
}

func (*MSStream) MsDetachStream() (i interface{}) {
	js.Rewrite("$<.msDetachStream()")
	return i
}

func (*MSStream) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

type MSStreamReader struct {
	*EventTarget
	*MSBaseReader
}

func (*MSStreamReader) ReadAsArrayBuffer(stream *MSStream, size int64) {
	js.Rewrite("$<.readAsArrayBuffer($1, $2)", stream, size)
}

func (*MSStreamReader) ReadAsBinaryString(stream *MSStream, size int64) {
	js.Rewrite("$<.readAsBinaryString($1, $2)", stream, size)
}

func (*MSStreamReader) ReadAsBlob(stream *MSStream, size int64) {
	js.Rewrite("$<.readAsBlob($1, $2)", stream, size)
}

func (*MSStreamReader) ReadAsDataURL(stream *MSStream, size int64) {
	js.Rewrite("$<.readAsDataURL($1, $2)", stream, size)
}

func (*MSStreamReader) ReadAsText(stream *MSStream, encoding string, size int64) {
	js.Rewrite("$<.readAsText($1, $2, $3)", stream, encoding, size)
}

func (*MSStreamReader) GetError() (err *DOMError) {
	js.Rewrite("$<.error")
	return err
}

type MSWebViewAsyncOperation struct {
	*EventTarget
}

func (*MSWebViewAsyncOperation) Start() {
	js.Rewrite("$<.start()")
}

func (*MSWebViewAsyncOperation) GetError() (err *DOMError) {
	js.Rewrite("$<.error")
	return err
}

func (*MSWebViewAsyncOperation) GetOncomplete() (complete *Event) {
	js.Rewrite("$<.oncomplete")
	return complete
}

func (*MSWebViewAsyncOperation) SetOncomplete(complete *Event) {
	js.Rewrite("$<.oncomplete = $1", complete)
}

func (*MSWebViewAsyncOperation) GetOnerror() (err *Event) {
	js.Rewrite("$<.onerror")
	return err
}

func (*MSWebViewAsyncOperation) SetOnerror(err *Event) {
	js.Rewrite("$<.onerror = $1", err)
}

func (*MSWebViewAsyncOperation) GetReadyState() (readyState uint8) {
	js.Rewrite("$<.readyState")
	return readyState
}

func (*MSWebViewAsyncOperation) GetResult() (result interface{}) {
	js.Rewrite("$<.result")
	return result
}

func (*MSWebViewAsyncOperation) GetTarget() (target *MSHTMLWebViewElement) {
	js.Rewrite("$<.target")
	return target
}

func (*MSWebViewAsyncOperation) GetType() (kind uint8) {
	js.Rewrite("$<.type")
	return kind
}

type MSWebViewSettings struct {
}

func (*MSWebViewSettings) GetIsIndexedDBEnabled() (isIndexedDBEnabled bool) {
	js.Rewrite("$<.isIndexedDBEnabled")
	return isIndexedDBEnabled
}

func (*MSWebViewSettings) SetIsIndexedDBEnabled(isIndexedDBEnabled bool) {
	js.Rewrite("$<.isIndexedDBEnabled = $1", isIndexedDBEnabled)
}

func (*MSWebViewSettings) GetIsJavaScriptEnabled() (isJavaScriptEnabled bool) {
	js.Rewrite("$<.isJavaScriptEnabled")
	return isJavaScriptEnabled
}

func (*MSWebViewSettings) SetIsJavaScriptEnabled(isJavaScriptEnabled bool) {
	js.Rewrite("$<.isJavaScriptEnabled = $1", isJavaScriptEnabled)
}

type MutationEvent struct {
	*Event
}

func (*MutationEvent) InitMutationEvent(typeArg string, canBubbleArg bool, cancelableArg bool, relatedNodeArg *Node, prevValueArg string, newValueArg string, attrNameArg string, attrChangeArg uint8) {
	js.Rewrite("$<.initMutationEvent($1, $2, $3, $4, $5, $6, $7, $8)", typeArg, canBubbleArg, cancelableArg, relatedNodeArg, prevValueArg, newValueArg, attrNameArg, attrChangeArg)
}

func (*MutationEvent) GetAttrChange() (attrChange uint8) {
	js.Rewrite("$<.attrChange")
	return attrChange
}

func (*MutationEvent) GetAttrName() (attrName string) {
	js.Rewrite("$<.attrName")
	return attrName
}

func (*MutationEvent) GetNewValue() (newValue string) {
	js.Rewrite("$<.newValue")
	return newValue
}

func (*MutationEvent) GetPrevValue() (prevValue string) {
	js.Rewrite("$<.prevValue")
	return prevValue
}

func (*MutationEvent) GetRelatedNode() (relatedNode *Node) {
	js.Rewrite("$<.relatedNode")
	return relatedNode
}

type MutationObserver struct {
}

func (*MutationObserver) Disconnect() {
	js.Rewrite("$<.disconnect()")
}

func (*MutationObserver) Observe(target *Node, options *MutationObserverInit) {
	js.Rewrite("$<.observe($1, $2)", target, options)
}

func (*MutationObserver) TakeRecords() (m []*MutationRecord) {
	js.Rewrite("$<.takeRecords()")
	return m
}

type MutationRecord struct {
}

func (*MutationRecord) GetAddedNodes() (addedNodes *NodeList) {
	js.Rewrite("$<.addedNodes")
	return addedNodes
}

func (*MutationRecord) GetAttributeName() (attributeName string) {
	js.Rewrite("$<.attributeName")
	return attributeName
}

func (*MutationRecord) GetAttributeNamespace() (attributeNamespace string) {
	js.Rewrite("$<.attributeNamespace")
	return attributeNamespace
}

func (*MutationRecord) GetNextSibling() (nextSibling *Node) {
	js.Rewrite("$<.nextSibling")
	return nextSibling
}

func (*MutationRecord) GetOldValue() (oldValue string) {
	js.Rewrite("$<.oldValue")
	return oldValue
}

func (*MutationRecord) GetPreviousSibling() (previousSibling *Node) {
	js.Rewrite("$<.previousSibling")
	return previousSibling
}

func (*MutationRecord) GetRemovedNodes() (removedNodes *NodeList) {
	js.Rewrite("$<.removedNodes")
	return removedNodes
}

func (*MutationRecord) GetTarget() (target *Node) {
	js.Rewrite("$<.target")
	return target
}

func (*MutationRecord) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

type NamedNodeMap struct {
}

func (*NamedNodeMap) GetNamedItem(name string) (a *Attr) {
	js.Rewrite("$<.getNamedItem($1)", name)
	return a
}

func (*NamedNodeMap) GetNamedItemNS(namespaceURI string, localName string) (a *Attr) {
	js.Rewrite("$<.getNamedItemNS($1, $2)", namespaceURI, localName)
	return a
}

func (*NamedNodeMap) Item(index uint) (a *Attr) {
	js.Rewrite("$<.item($1)", index)
	return a
}

func (*NamedNodeMap) RemoveNamedItem(name string) (a *Attr) {
	js.Rewrite("$<.removeNamedItem($1)", name)
	return a
}

func (*NamedNodeMap) RemoveNamedItemNS(namespaceURI string, localName string) (a *Attr) {
	js.Rewrite("$<.removeNamedItemNS($1, $2)", namespaceURI, localName)
	return a
}

func (*NamedNodeMap) SetNamedItem(arg *Attr) (a *Attr) {
	js.Rewrite("$<.setNamedItem($1)", arg)
	return a
}

func (*NamedNodeMap) SetNamedItemNS(arg *Attr) (a *Attr) {
	js.Rewrite("$<.setNamedItemNS($1)", arg)
	return a
}

func (*NamedNodeMap) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

type NavigationCompletedEvent struct {
	*NavigationEvent
}

func (*NavigationCompletedEvent) GetIsSuccess() (isSuccess bool) {
	js.Rewrite("$<.isSuccess")
	return isSuccess
}

func (*NavigationCompletedEvent) GetWebErrorStatus() (webErrorStatus uint64) {
	js.Rewrite("$<.webErrorStatus")
	return webErrorStatus
}

type NavigationEvent struct {
	*Event
}

func (*NavigationEvent) GetURI() (uri string) {
	js.Rewrite("$<.uri")
	return uri
}

type NavigationEventWithReferrer struct {
	*NavigationEvent
}

func (*NavigationEventWithReferrer) GetReferer() (referer string) {
	js.Rewrite("$<.referer")
	return referer
}

type Navigator struct {
	*NavigatorID
	*NavigatorOnLine
	*NavigatorContentUtils
	*NavigatorStorageUtils
	*NavigatorGeolocation
	*MSNavigatorDoNotTrack
	*MSFileSaver
	*NavigatorBeacon
	*NavigatorConcurrentHardware
	*NavigatorUserMedia
}

func (*Navigator) GetGamepads() (g []*Gamepad) {
	js.Rewrite("$<.getGamepads()")
	return g
}

func (*Navigator) JavaEnabled() (b bool) {
	js.Rewrite("$<.javaEnabled()")
	return b
}

func (*Navigator) MsLaunchURI(uri string, successCallback func(), noHandlerCallback func()) {
	js.Rewrite("$<.msLaunchUri($1, $2, $3)", uri, successCallback, noHandlerCallback)
}

func (*Navigator) RequestMediaKeySystemAccess(keySystem string, supportedConfigurations []*MediaKeySystemConfiguration) (m *MediaKeySystemAccess) {
	js.Rewrite("await $<.requestMediaKeySystemAccess($1, $2)", keySystem, supportedConfigurations)
	return m
}

func (*Navigator) GetAuthentication() (authentication *WebAuthentication) {
	js.Rewrite("$<.authentication")
	return authentication
}

func (*Navigator) GetCookieEnabled() (cookieEnabled bool) {
	js.Rewrite("$<.cookieEnabled")
	return cookieEnabled
}

func (*Navigator) GetGamepadInputEmulation() (gamepadInputEmulation *GamepadInputEmulationType) {
	js.Rewrite("$<.gamepadInputEmulation")
	return gamepadInputEmulation
}

func (*Navigator) SetGamepadInputEmulation(gamepadInputEmulation *GamepadInputEmulationType) {
	js.Rewrite("$<.gamepadInputEmulation = $1", gamepadInputEmulation)
}

func (*Navigator) GetLanguage() (language string) {
	js.Rewrite("$<.language")
	return language
}

func (*Navigator) GetMaxTouchPoints() (maxTouchPoints int) {
	js.Rewrite("$<.maxTouchPoints")
	return maxTouchPoints
}

func (*Navigator) GetMimeTypes() (mimeTypes *MimeTypeArray) {
	js.Rewrite("$<.mimeTypes")
	return mimeTypes
}

func (*Navigator) GetMsManipulationViewsEnabled() (msManipulationViewsEnabled bool) {
	js.Rewrite("$<.msManipulationViewsEnabled")
	return msManipulationViewsEnabled
}

func (*Navigator) GetMsMaxTouchPoints() (msMaxTouchPoints int) {
	js.Rewrite("$<.msMaxTouchPoints")
	return msMaxTouchPoints
}

func (*Navigator) GetMsPointerEnabled() (msPointerEnabled bool) {
	js.Rewrite("$<.msPointerEnabled")
	return msPointerEnabled
}

func (*Navigator) GetPlugins() (plugins *PluginArray) {
	js.Rewrite("$<.plugins")
	return plugins
}

func (*Navigator) GetPointerEnabled() (pointerEnabled bool) {
	js.Rewrite("$<.pointerEnabled")
	return pointerEnabled
}

func (*Navigator) GetServiceWorker() (serviceWorker *ServiceWorkerContainer) {
	js.Rewrite("$<.serviceWorker")
	return serviceWorker
}

func (*Navigator) GetWebdriver() (webdriver bool) {
	js.Rewrite("$<.webdriver")
	return webdriver
}

type Node struct {
	*EventTarget
}

func (*Node) AppendChild(newChild *Node) (n *Node) {
	js.Rewrite("$<.appendChild($1)", newChild)
	return n
}

func (*Node) CloneNode(deep bool) (n *Node) {
	js.Rewrite("$<.cloneNode($1)", deep)
	return n
}

func (*Node) CompareDocumentPosition(other *Node) (u uint8) {
	js.Rewrite("$<.compareDocumentPosition($1)", other)
	return u
}

func (*Node) Contains(child *Node) (b bool) {
	js.Rewrite("$<.contains($1)", child)
	return b
}

func (*Node) HasAttributes() (b bool) {
	js.Rewrite("$<.hasAttributes()")
	return b
}

func (*Node) HasChildNodes() (b bool) {
	js.Rewrite("$<.hasChildNodes()")
	return b
}

func (*Node) InsertBefore(newChild *Node, refChild *Node) (n *Node) {
	js.Rewrite("$<.insertBefore($1, $2)", newChild, refChild)
	return n
}

func (*Node) IsDefaultNamespace(namespaceURI string) (b bool) {
	js.Rewrite("$<.isDefaultNamespace($1)", namespaceURI)
	return b
}

func (*Node) IsEqualNode(arg *Node) (b bool) {
	js.Rewrite("$<.isEqualNode($1)", arg)
	return b
}

func (*Node) IsSameNode(other *Node) (b bool) {
	js.Rewrite("$<.isSameNode($1)", other)
	return b
}

func (*Node) LookupNamespaceURI(prefix string) (s string) {
	js.Rewrite("$<.lookupNamespaceURI($1)", prefix)
	return s
}

func (*Node) LookupPrefix(namespaceURI string) (s string) {
	js.Rewrite("$<.lookupPrefix($1)", namespaceURI)
	return s
}

func (*Node) Normalize() {
	js.Rewrite("$<.normalize()")
}

func (*Node) RemoveChild(oldChild *Node) (n *Node) {
	js.Rewrite("$<.removeChild($1)", oldChild)
	return n
}

func (*Node) ReplaceChild(newChild *Node, oldChild *Node) (n *Node) {
	js.Rewrite("$<.replaceChild($1, $2)", newChild, oldChild)
	return n
}

func (*Node) GetAttributes() (attributes *NamedNodeMap) {
	js.Rewrite("$<.attributes")
	return attributes
}

func (*Node) GetBaseURI() (baseURI string) {
	js.Rewrite("$<.baseURI")
	return baseURI
}

func (*Node) GetChildNodes() (childNodes *NodeList) {
	js.Rewrite("$<.childNodes")
	return childNodes
}

func (*Node) GetFirstChild() (firstChild *Node) {
	js.Rewrite("$<.firstChild")
	return firstChild
}

func (*Node) GetLastChild() (lastChild *Node) {
	js.Rewrite("$<.lastChild")
	return lastChild
}

func (*Node) GetLocalName() (localName string) {
	js.Rewrite("$<.localName")
	return localName
}

func (*Node) GetNamespaceURI() (namespaceURI string) {
	js.Rewrite("$<.namespaceURI")
	return namespaceURI
}

func (*Node) GetNextSibling() (nextSibling *Node) {
	js.Rewrite("$<.nextSibling")
	return nextSibling
}

func (*Node) GetNodeName() (nodeName string) {
	js.Rewrite("$<.nodeName")
	return nodeName
}

func (*Node) GetNodeType() (nodeType uint8) {
	js.Rewrite("$<.nodeType")
	return nodeType
}

func (*Node) GetNodeValue() (nodeValue string) {
	js.Rewrite("$<.nodeValue")
	return nodeValue
}

func (*Node) SetNodeValue(nodeValue string) {
	js.Rewrite("$<.nodeValue = $1", nodeValue)
}

func (*Node) GetOwnerDocument() (ownerDocument *Document) {
	js.Rewrite("$<.ownerDocument")
	return ownerDocument
}

func (*Node) GetParentElement() (parentElement *HTMLElement) {
	js.Rewrite("$<.parentElement")
	return parentElement
}

func (*Node) GetParentNode() (parentNode *Node) {
	js.Rewrite("$<.parentNode")
	return parentNode
}

func (*Node) GetPreviousSibling() (previousSibling *Node) {
	js.Rewrite("$<.previousSibling")
	return previousSibling
}

func (*Node) GetTextContent() (textContent string) {
	js.Rewrite("$<.textContent")
	return textContent
}

func (*Node) SetTextContent(textContent string) {
	js.Rewrite("$<.textContent = $1", textContent)
}

type NodeFilter struct {
}

func (*NodeFilter) AcceptNode(n *Node) (i int8) {
	js.Rewrite("$<.acceptNode($1)", n)
	return i
}

type NodeIterator struct {
}

func (*NodeIterator) Detach() {
	js.Rewrite("$<.detach()")
}

func (*NodeIterator) NextNode() (n *Node) {
	js.Rewrite("$<.nextNode()")
	return n
}

func (*NodeIterator) PreviousNode() (n *Node) {
	js.Rewrite("$<.previousNode()")
	return n
}

func (*NodeIterator) GetExpandEntityReferences() (expandEntityReferences bool) {
	js.Rewrite("$<.expandEntityReferences")
	return expandEntityReferences
}

func (*NodeIterator) GetFilter() (filter *NodeFilter) {
	js.Rewrite("$<.filter")
	return filter
}

func (*NodeIterator) GetRoot() (root *Node) {
	js.Rewrite("$<.root")
	return root
}

func (*NodeIterator) GetWhatToShow() (whatToShow uint) {
	js.Rewrite("$<.whatToShow")
	return whatToShow
}

type NodeList struct {
}

func (*NodeList) Item(index uint) (n *Node) {
	js.Rewrite("$<.item($1)", index)
	return n
}

func (*NodeList) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

type Notification struct {
	*EventTarget
}

func (*Notification) Close() {
	js.Rewrite("$<.close()")
}

func (*Notification) RequestPermission(callback func(permission *NotificationPermission)) (n *NotificationPermission) {
	js.Rewrite("await $<.requestPermission($1)", callback)
	return n
}

func (*Notification) GetBody() (body string) {
	js.Rewrite("$<.body")
	return body
}

func (*Notification) GetDir() (dir *NotificationDirection) {
	js.Rewrite("$<.dir")
	return dir
}

func (*Notification) GetIcon() (icon string) {
	js.Rewrite("$<.icon")
	return icon
}

func (*Notification) GetLang() (lang string) {
	js.Rewrite("$<.lang")
	return lang
}

func (*Notification) GetOnclick() (click *Event) {
	js.Rewrite("$<.onclick")
	return click
}

func (*Notification) SetOnclick(click *Event) {
	js.Rewrite("$<.onclick = $1", click)
}

func (*Notification) GetOnclose() (cls *Event) {
	js.Rewrite("$<.onclose")
	return cls
}

func (*Notification) SetOnclose(cls *Event) {
	js.Rewrite("$<.onclose = $1", cls)
}

func (*Notification) GetOnerror() (err *Event) {
	js.Rewrite("$<.onerror")
	return err
}

func (*Notification) SetOnerror(err *Event) {
	js.Rewrite("$<.onerror = $1", err)
}

func (*Notification) GetOnshow() (show *Event) {
	js.Rewrite("$<.onshow")
	return show
}

func (*Notification) SetOnshow(show *Event) {
	js.Rewrite("$<.onshow = $1", show)
}

func (*Notification) GetPermission() (permission *NotificationPermission) {
	js.Rewrite("$<.permission")
	return permission
}

func (*Notification) GetTag() (tag string) {
	js.Rewrite("$<.tag")
	return tag
}

func (*Notification) GetTitle() (title string) {
	js.Rewrite("$<.title")
	return title
}

type OES_element_index_uint struct {
}

type OES_standard_derivatives struct {
}

type OES_texture_float struct {
}

type OES_texture_float_linear struct {
}

type OES_texture_half_float struct {
}

type OES_texture_half_float_linear struct {
}

type OfflineAudioCompletionEvent struct {
	*Event
}

func (*OfflineAudioCompletionEvent) GetRenderedBuffer() (renderedBuffer *AudioBuffer) {
	js.Rewrite("$<.renderedBuffer")
	return renderedBuffer
}

type OfflineAudioContext struct {
	*AudioContext
}

func (*OfflineAudioContext) StartRendering() (a *AudioBuffer) {
	js.Rewrite("await $<.startRendering()")
	return a
}

func (*OfflineAudioContext) Suspend(suspendTime float32) {
	js.Rewrite("await $<.suspend($1)", suspendTime)
}

func (*OfflineAudioContext) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

func (*OfflineAudioContext) GetOncomplete() (complete *OfflineAudioCompletionEvent) {
	js.Rewrite("$<.oncomplete")
	return complete
}

func (*OfflineAudioContext) SetOncomplete(complete *OfflineAudioCompletionEvent) {
	js.Rewrite("$<.oncomplete = $1", complete)
}

type OscillatorNode struct {
	*AudioNode
}

func (*OscillatorNode) SetPeriodicWave(periodicWave *PeriodicWave) {
	js.Rewrite("$<.setPeriodicWave($1)", periodicWave)
}

func (*OscillatorNode) Start(when float32) {
	js.Rewrite("$<.start($1)", when)
}

func (*OscillatorNode) Stop(when float32) {
	js.Rewrite("$<.stop($1)", when)
}

func (*OscillatorNode) GetDetune() (detune *AudioParam) {
	js.Rewrite("$<.detune")
	return detune
}

func (*OscillatorNode) GetFrequency() (frequency *AudioParam) {
	js.Rewrite("$<.frequency")
	return frequency
}

func (*OscillatorNode) GetOnended() (e *Event) {
	js.Rewrite("$<.onended")
	return e
}

func (*OscillatorNode) SetOnended(e *Event) {
	js.Rewrite("$<.onended = $1", e)
}

func (*OscillatorNode) GetType() (kind *OscillatorType) {
	js.Rewrite("$<.type")
	return kind
}

func (*OscillatorNode) SetType(kind *OscillatorType) {
	js.Rewrite("$<.type = $1", kind)
}

type OverflowEvent struct {
	*UIEvent
}

func (*OverflowEvent) GetHorizontalOverflow() (horizontalOverflow bool) {
	js.Rewrite("$<.horizontalOverflow")
	return horizontalOverflow
}

func (*OverflowEvent) GetOrient() (orient uint) {
	js.Rewrite("$<.orient")
	return orient
}

func (*OverflowEvent) GetVerticalOverflow() (verticalOverflow bool) {
	js.Rewrite("$<.verticalOverflow")
	return verticalOverflow
}

type PageTransitionEvent struct {
	*Event
}

func (*PageTransitionEvent) GetPersisted() (persisted bool) {
	js.Rewrite("$<.persisted")
	return persisted
}

type PannerNode struct {
	*AudioNode
}

func (*PannerNode) SetOrientation(x float32, y float32, z float32) {
	js.Rewrite("$<.setOrientation($1, $2, $3)", x, y, z)
}

func (*PannerNode) SetPosition(x float32, y float32, z float32) {
	js.Rewrite("$<.setPosition($1, $2, $3)", x, y, z)
}

func (*PannerNode) SetVelocity(x float32, y float32, z float32) {
	js.Rewrite("$<.setVelocity($1, $2, $3)", x, y, z)
}

func (*PannerNode) GetConeInnerAngle() (coneInnerAngle float32) {
	js.Rewrite("$<.coneInnerAngle")
	return coneInnerAngle
}

func (*PannerNode) SetConeInnerAngle(coneInnerAngle float32) {
	js.Rewrite("$<.coneInnerAngle = $1", coneInnerAngle)
}

func (*PannerNode) GetConeOuterAngle() (coneOuterAngle float32) {
	js.Rewrite("$<.coneOuterAngle")
	return coneOuterAngle
}

func (*PannerNode) SetConeOuterAngle(coneOuterAngle float32) {
	js.Rewrite("$<.coneOuterAngle = $1", coneOuterAngle)
}

func (*PannerNode) GetConeOuterGain() (coneOuterGain float32) {
	js.Rewrite("$<.coneOuterGain")
	return coneOuterGain
}

func (*PannerNode) SetConeOuterGain(coneOuterGain float32) {
	js.Rewrite("$<.coneOuterGain = $1", coneOuterGain)
}

func (*PannerNode) GetDistanceModel() (distanceModel *DistanceModelType) {
	js.Rewrite("$<.distanceModel")
	return distanceModel
}

func (*PannerNode) SetDistanceModel(distanceModel *DistanceModelType) {
	js.Rewrite("$<.distanceModel = $1", distanceModel)
}

func (*PannerNode) GetMaxDistance() (maxDistance float32) {
	js.Rewrite("$<.maxDistance")
	return maxDistance
}

func (*PannerNode) SetMaxDistance(maxDistance float32) {
	js.Rewrite("$<.maxDistance = $1", maxDistance)
}

func (*PannerNode) GetPanningModel() (panningModel *PanningModelType) {
	js.Rewrite("$<.panningModel")
	return panningModel
}

func (*PannerNode) SetPanningModel(panningModel *PanningModelType) {
	js.Rewrite("$<.panningModel = $1", panningModel)
}

func (*PannerNode) GetRefDistance() (refDistance float32) {
	js.Rewrite("$<.refDistance")
	return refDistance
}

func (*PannerNode) SetRefDistance(refDistance float32) {
	js.Rewrite("$<.refDistance = $1", refDistance)
}

func (*PannerNode) GetRolloffFactor() (rolloffFactor float32) {
	js.Rewrite("$<.rolloffFactor")
	return rolloffFactor
}

func (*PannerNode) SetRolloffFactor(rolloffFactor float32) {
	js.Rewrite("$<.rolloffFactor = $1", rolloffFactor)
}

type Path2D struct {
	*CanvasPathMethods
}

type PaymentAddress struct {
}

func (*PaymentAddress) ToJSON() (i interface{}) {
	js.Rewrite("$<.toJSON()")
	return i
}

func (*PaymentAddress) GetAddressLine() (addressLine []string) {
	js.Rewrite("$<.addressLine")
	return addressLine
}

func (*PaymentAddress) GetCity() (city string) {
	js.Rewrite("$<.city")
	return city
}

func (*PaymentAddress) GetCountry() (country string) {
	js.Rewrite("$<.country")
	return country
}

func (*PaymentAddress) GetDependentLocality() (dependentLocality string) {
	js.Rewrite("$<.dependentLocality")
	return dependentLocality
}

func (*PaymentAddress) GetLanguageCode() (languageCode string) {
	js.Rewrite("$<.languageCode")
	return languageCode
}

func (*PaymentAddress) GetOrganization() (organization string) {
	js.Rewrite("$<.organization")
	return organization
}

func (*PaymentAddress) GetPhone() (phone string) {
	js.Rewrite("$<.phone")
	return phone
}

func (*PaymentAddress) GetPostalCode() (postalCode string) {
	js.Rewrite("$<.postalCode")
	return postalCode
}

func (*PaymentAddress) GetRecipient() (recipient string) {
	js.Rewrite("$<.recipient")
	return recipient
}

func (*PaymentAddress) GetRegion() (region string) {
	js.Rewrite("$<.region")
	return region
}

func (*PaymentAddress) GetSortingCode() (sortingCode string) {
	js.Rewrite("$<.sortingCode")
	return sortingCode
}

type PaymentRequest struct {
	*EventTarget
}

func (*PaymentRequest) Abort() {
	js.Rewrite("await $<.abort()")
}

func (*PaymentRequest) Show() (p *PaymentResponse) {
	js.Rewrite("await $<.show()")
	return p
}

func (*PaymentRequest) GetOnshippingaddresschange() (shippingaddresschange *Event) {
	js.Rewrite("$<.onshippingaddresschange")
	return shippingaddresschange
}

func (*PaymentRequest) SetOnshippingaddresschange(shippingaddresschange *Event) {
	js.Rewrite("$<.onshippingaddresschange = $1", shippingaddresschange)
}

func (*PaymentRequest) GetOnshippingoptionchange() (shippingoptionchange *Event) {
	js.Rewrite("$<.onshippingoptionchange")
	return shippingoptionchange
}

func (*PaymentRequest) SetOnshippingoptionchange(shippingoptionchange *Event) {
	js.Rewrite("$<.onshippingoptionchange = $1", shippingoptionchange)
}

func (*PaymentRequest) GetShippingAddress() (shippingAddress *PaymentAddress) {
	js.Rewrite("$<.shippingAddress")
	return shippingAddress
}

func (*PaymentRequest) GetShippingOption() (shippingOption string) {
	js.Rewrite("$<.shippingOption")
	return shippingOption
}

func (*PaymentRequest) GetShippingType() (shippingType *PaymentShippingType) {
	js.Rewrite("$<.shippingType")
	return shippingType
}

type PaymentRequestUpdateEvent struct {
	*Event
}

func (*PaymentRequestUpdateEvent) UpdateWith(d *PaymentDetails) {
	js.Rewrite("$<.updateWith($1)", d)
}

type PaymentResponse struct {
}

func (*PaymentResponse) Complete(result *PaymentComplete) {
	js.Rewrite("await $<.complete($1)", result)
}

func (*PaymentResponse) ToJSON() (i interface{}) {
	js.Rewrite("$<.toJSON()")
	return i
}

func (*PaymentResponse) GetDetails() (details interface{}) {
	js.Rewrite("$<.details")
	return details
}

func (*PaymentResponse) GetMethodName() (methodName string) {
	js.Rewrite("$<.methodName")
	return methodName
}

func (*PaymentResponse) GetPayerEmail() (payerEmail string) {
	js.Rewrite("$<.payerEmail")
	return payerEmail
}

func (*PaymentResponse) GetPayerName() (payerName string) {
	js.Rewrite("$<.payerName")
	return payerName
}

func (*PaymentResponse) GetPayerPhone() (payerPhone string) {
	js.Rewrite("$<.payerPhone")
	return payerPhone
}

func (*PaymentResponse) GetShippingAddress() (shippingAddress *PaymentAddress) {
	js.Rewrite("$<.shippingAddress")
	return shippingAddress
}

func (*PaymentResponse) GetShippingOption() (shippingOption string) {
	js.Rewrite("$<.shippingOption")
	return shippingOption
}

type Performance struct {
}

func (*Performance) ClearMarks(markName string) {
	js.Rewrite("$<.clearMarks($1)", markName)
}

func (*Performance) ClearMeasures(measureName string) {
	js.Rewrite("$<.clearMeasures($1)", measureName)
}

func (*Performance) ClearResourceTimings() {
	js.Rewrite("$<.clearResourceTimings()")
}

func (*Performance) GetEntries() (i interface{}) {
	js.Rewrite("$<.getEntries()")
	return i
}

func (*Performance) GetEntriesByName(name string, entryType string) (i interface{}) {
	js.Rewrite("$<.getEntriesByName($1, $2)", name, entryType)
	return i
}

func (*Performance) GetEntriesByType(entryType string) (i interface{}) {
	js.Rewrite("$<.getEntriesByType($1)", entryType)
	return i
}

func (*Performance) GetMarks(markName string) (i interface{}) {
	js.Rewrite("$<.getMarks($1)", markName)
	return i
}

func (*Performance) GetMeasures(measureName string) (i interface{}) {
	js.Rewrite("$<.getMeasures($1)", measureName)
	return i
}

func (*Performance) Mark(markName string) {
	js.Rewrite("$<.mark($1)", markName)
}

func (*Performance) Measure(measureName string, startMarkName string, endMarkName string) {
	js.Rewrite("$<.measure($1, $2, $3)", measureName, startMarkName, endMarkName)
}

func (*Performance) Now() (i int) {
	js.Rewrite("$<.now()")
	return i
}

func (*Performance) SetResourceTimingBufferSize(maxSize uint) {
	js.Rewrite("$<.setResourceTimingBufferSize($1)", maxSize)
}

func (*Performance) ToJSON() (i interface{}) {
	js.Rewrite("$<.toJSON()")
	return i
}

func (*Performance) GetNavigation() (navigation *PerformanceNavigation) {
	js.Rewrite("$<.navigation")
	return navigation
}

func (*Performance) GetTiming() (timing *PerformanceTiming) {
	js.Rewrite("$<.timing")
	return timing
}

type PerformanceEntry struct {
}

func (*PerformanceEntry) GetDuration() (duration int) {
	js.Rewrite("$<.duration")
	return duration
}

func (*PerformanceEntry) GetEntryType() (entryType string) {
	js.Rewrite("$<.entryType")
	return entryType
}

func (*PerformanceEntry) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*PerformanceEntry) GetStartTime() (startTime int) {
	js.Rewrite("$<.startTime")
	return startTime
}

type PerformanceMark struct {
	*PerformanceEntry
}

type PerformanceMeasure struct {
	*PerformanceEntry
}

type PerformanceNavigation struct {
}

func (*PerformanceNavigation) ToJSON() (i interface{}) {
	js.Rewrite("$<.toJSON()")
	return i
}

func (*PerformanceNavigation) GetRedirectCount() (redirectCount uint) {
	js.Rewrite("$<.redirectCount")
	return redirectCount
}

func (*PerformanceNavigation) GetType() (kind uint) {
	js.Rewrite("$<.type")
	return kind
}

type PerformanceNavigationTiming struct {
	*PerformanceEntry
}

func (*PerformanceNavigationTiming) GetConnectEnd() (connectEnd int) {
	js.Rewrite("$<.connectEnd")
	return connectEnd
}

func (*PerformanceNavigationTiming) GetConnectStart() (connectStart int) {
	js.Rewrite("$<.connectStart")
	return connectStart
}

func (*PerformanceNavigationTiming) GetDomainLookupEnd() (domainLookupEnd int) {
	js.Rewrite("$<.domainLookupEnd")
	return domainLookupEnd
}

func (*PerformanceNavigationTiming) GetDomainLookupStart() (domainLookupStart int) {
	js.Rewrite("$<.domainLookupStart")
	return domainLookupStart
}

func (*PerformanceNavigationTiming) GetDomComplete() (domComplete int) {
	js.Rewrite("$<.domComplete")
	return domComplete
}

func (*PerformanceNavigationTiming) GetDomContentLoadedEventEnd() (domContentLoadedEventEnd int) {
	js.Rewrite("$<.domContentLoadedEventEnd")
	return domContentLoadedEventEnd
}

func (*PerformanceNavigationTiming) GetDomContentLoadedEventStart() (domContentLoadedEventStart int) {
	js.Rewrite("$<.domContentLoadedEventStart")
	return domContentLoadedEventStart
}

func (*PerformanceNavigationTiming) GetDomInteractive() (domInteractive int) {
	js.Rewrite("$<.domInteractive")
	return domInteractive
}

func (*PerformanceNavigationTiming) GetDomLoading() (domLoading int) {
	js.Rewrite("$<.domLoading")
	return domLoading
}

func (*PerformanceNavigationTiming) GetFetchStart() (fetchStart int) {
	js.Rewrite("$<.fetchStart")
	return fetchStart
}

func (*PerformanceNavigationTiming) GetLoadEventEnd() (loadEventEnd int) {
	js.Rewrite("$<.loadEventEnd")
	return loadEventEnd
}

func (*PerformanceNavigationTiming) GetLoadEventStart() (loadEventStart int) {
	js.Rewrite("$<.loadEventStart")
	return loadEventStart
}

func (*PerformanceNavigationTiming) GetNavigationStart() (navigationStart int) {
	js.Rewrite("$<.navigationStart")
	return navigationStart
}

func (*PerformanceNavigationTiming) GetRedirectCount() (redirectCount uint8) {
	js.Rewrite("$<.redirectCount")
	return redirectCount
}

func (*PerformanceNavigationTiming) GetRedirectEnd() (redirectEnd int) {
	js.Rewrite("$<.redirectEnd")
	return redirectEnd
}

func (*PerformanceNavigationTiming) GetRedirectStart() (redirectStart int) {
	js.Rewrite("$<.redirectStart")
	return redirectStart
}

func (*PerformanceNavigationTiming) GetRequestStart() (requestStart int) {
	js.Rewrite("$<.requestStart")
	return requestStart
}

func (*PerformanceNavigationTiming) GetResponseEnd() (responseEnd int) {
	js.Rewrite("$<.responseEnd")
	return responseEnd
}

func (*PerformanceNavigationTiming) GetResponseStart() (responseStart int) {
	js.Rewrite("$<.responseStart")
	return responseStart
}

func (*PerformanceNavigationTiming) GetType() (kind *NavigationType) {
	js.Rewrite("$<.type")
	return kind
}

func (*PerformanceNavigationTiming) GetUnloadEventEnd() (unloadEventEnd int) {
	js.Rewrite("$<.unloadEventEnd")
	return unloadEventEnd
}

func (*PerformanceNavigationTiming) GetUnloadEventStart() (unloadEventStart int) {
	js.Rewrite("$<.unloadEventStart")
	return unloadEventStart
}

type PerformanceResourceTiming struct {
	*PerformanceEntry
}

func (*PerformanceResourceTiming) GetConnectEnd() (connectEnd int) {
	js.Rewrite("$<.connectEnd")
	return connectEnd
}

func (*PerformanceResourceTiming) GetConnectStart() (connectStart int) {
	js.Rewrite("$<.connectStart")
	return connectStart
}

func (*PerformanceResourceTiming) GetDomainLookupEnd() (domainLookupEnd int) {
	js.Rewrite("$<.domainLookupEnd")
	return domainLookupEnd
}

func (*PerformanceResourceTiming) GetDomainLookupStart() (domainLookupStart int) {
	js.Rewrite("$<.domainLookupStart")
	return domainLookupStart
}

func (*PerformanceResourceTiming) GetFetchStart() (fetchStart int) {
	js.Rewrite("$<.fetchStart")
	return fetchStart
}

func (*PerformanceResourceTiming) GetInitiatorType() (initiatorType string) {
	js.Rewrite("$<.initiatorType")
	return initiatorType
}

func (*PerformanceResourceTiming) GetRedirectEnd() (redirectEnd int) {
	js.Rewrite("$<.redirectEnd")
	return redirectEnd
}

func (*PerformanceResourceTiming) GetRedirectStart() (redirectStart int) {
	js.Rewrite("$<.redirectStart")
	return redirectStart
}

func (*PerformanceResourceTiming) GetRequestStart() (requestStart int) {
	js.Rewrite("$<.requestStart")
	return requestStart
}

func (*PerformanceResourceTiming) GetResponseEnd() (responseEnd int) {
	js.Rewrite("$<.responseEnd")
	return responseEnd
}

func (*PerformanceResourceTiming) GetResponseStart() (responseStart int) {
	js.Rewrite("$<.responseStart")
	return responseStart
}

type PerformanceTiming struct {
}

func (*PerformanceTiming) ToJSON() (i interface{}) {
	js.Rewrite("$<.toJSON()")
	return i
}

func (*PerformanceTiming) GetConnectEnd() (connectEnd uint64) {
	js.Rewrite("$<.connectEnd")
	return connectEnd
}

func (*PerformanceTiming) GetConnectStart() (connectStart uint64) {
	js.Rewrite("$<.connectStart")
	return connectStart
}

func (*PerformanceTiming) GetDomainLookupEnd() (domainLookupEnd uint64) {
	js.Rewrite("$<.domainLookupEnd")
	return domainLookupEnd
}

func (*PerformanceTiming) GetDomainLookupStart() (domainLookupStart uint64) {
	js.Rewrite("$<.domainLookupStart")
	return domainLookupStart
}

func (*PerformanceTiming) GetDomComplete() (domComplete uint64) {
	js.Rewrite("$<.domComplete")
	return domComplete
}

func (*PerformanceTiming) GetDomContentLoadedEventEnd() (domContentLoadedEventEnd uint64) {
	js.Rewrite("$<.domContentLoadedEventEnd")
	return domContentLoadedEventEnd
}

func (*PerformanceTiming) GetDomContentLoadedEventStart() (domContentLoadedEventStart uint64) {
	js.Rewrite("$<.domContentLoadedEventStart")
	return domContentLoadedEventStart
}

func (*PerformanceTiming) GetDomInteractive() (domInteractive uint64) {
	js.Rewrite("$<.domInteractive")
	return domInteractive
}

func (*PerformanceTiming) GetDomLoading() (domLoading uint64) {
	js.Rewrite("$<.domLoading")
	return domLoading
}

func (*PerformanceTiming) GetFetchStart() (fetchStart uint64) {
	js.Rewrite("$<.fetchStart")
	return fetchStart
}

func (*PerformanceTiming) GetLoadEventEnd() (loadEventEnd uint64) {
	js.Rewrite("$<.loadEventEnd")
	return loadEventEnd
}

func (*PerformanceTiming) GetLoadEventStart() (loadEventStart uint64) {
	js.Rewrite("$<.loadEventStart")
	return loadEventStart
}

func (*PerformanceTiming) GetMsFirstPaint() (msFirstPaint uint64) {
	js.Rewrite("$<.msFirstPaint")
	return msFirstPaint
}

func (*PerformanceTiming) GetNavigationStart() (navigationStart uint64) {
	js.Rewrite("$<.navigationStart")
	return navigationStart
}

func (*PerformanceTiming) GetRedirectEnd() (redirectEnd uint64) {
	js.Rewrite("$<.redirectEnd")
	return redirectEnd
}

func (*PerformanceTiming) GetRedirectStart() (redirectStart uint64) {
	js.Rewrite("$<.redirectStart")
	return redirectStart
}

func (*PerformanceTiming) GetRequestStart() (requestStart uint64) {
	js.Rewrite("$<.requestStart")
	return requestStart
}

func (*PerformanceTiming) GetResponseEnd() (responseEnd uint64) {
	js.Rewrite("$<.responseEnd")
	return responseEnd
}

func (*PerformanceTiming) GetResponseStart() (responseStart uint64) {
	js.Rewrite("$<.responseStart")
	return responseStart
}

func (*PerformanceTiming) GetUnloadEventEnd() (unloadEventEnd uint64) {
	js.Rewrite("$<.unloadEventEnd")
	return unloadEventEnd
}

func (*PerformanceTiming) GetUnloadEventStart() (unloadEventStart uint64) {
	js.Rewrite("$<.unloadEventStart")
	return unloadEventStart
}

type PerfWidgetExternal struct {
}

func (*PerfWidgetExternal) AddEventListener(eventType string, callback func()) {
	js.Rewrite("$<.addEventListener($1, $2)", eventType, callback)
}

func (*PerfWidgetExternal) GetMemoryUsage() (u uint) {
	js.Rewrite("$<.getMemoryUsage()")
	return u
}

func (*PerfWidgetExternal) GetProcessCPUUsage() (u uint) {
	js.Rewrite("$<.getProcessCpuUsage()")
	return u
}

func (*PerfWidgetExternal) GetRecentCPUUsage(last uint64) (i interface{}) {
	js.Rewrite("$<.getRecentCpuUsage($1)", last)
	return i
}

func (*PerfWidgetExternal) GetRecentFrames(last uint64) (i interface{}) {
	js.Rewrite("$<.getRecentFrames($1)", last)
	return i
}

func (*PerfWidgetExternal) GetRecentMemoryUsage(last uint64) (i interface{}) {
	js.Rewrite("$<.getRecentMemoryUsage($1)", last)
	return i
}

func (*PerfWidgetExternal) GetRecentPaintRequests(last uint64) (i interface{}) {
	js.Rewrite("$<.getRecentPaintRequests($1)", last)
	return i
}

func (*PerfWidgetExternal) RemoveEventListener(eventType string, callback func()) {
	js.Rewrite("$<.removeEventListener($1, $2)", eventType, callback)
}

func (*PerfWidgetExternal) RepositionWindow(x int, y int) {
	js.Rewrite("$<.repositionWindow($1, $2)", x, y)
}

func (*PerfWidgetExternal) ResizeWindow(width uint, height uint) {
	js.Rewrite("$<.resizeWindow($1, $2)", width, height)
}

func (*PerfWidgetExternal) GetActiveNetworkRequestCount() (activeNetworkRequestCount uint) {
	js.Rewrite("$<.activeNetworkRequestCount")
	return activeNetworkRequestCount
}

func (*PerfWidgetExternal) GetAverageFrameTime() (averageFrameTime float32) {
	js.Rewrite("$<.averageFrameTime")
	return averageFrameTime
}

func (*PerfWidgetExternal) GetAveragePaintTime() (averagePaintTime float32) {
	js.Rewrite("$<.averagePaintTime")
	return averagePaintTime
}

func (*PerfWidgetExternal) GetExtraInformationEnabled() (extraInformationEnabled bool) {
	js.Rewrite("$<.extraInformationEnabled")
	return extraInformationEnabled
}

func (*PerfWidgetExternal) GetIndependentRenderingEnabled() (independentRenderingEnabled bool) {
	js.Rewrite("$<.independentRenderingEnabled")
	return independentRenderingEnabled
}

func (*PerfWidgetExternal) GetIrDisablingContentString() (irDisablingContentString string) {
	js.Rewrite("$<.irDisablingContentString")
	return irDisablingContentString
}

func (*PerfWidgetExternal) GetIrStatusAvailable() (irStatusAvailable bool) {
	js.Rewrite("$<.irStatusAvailable")
	return irStatusAvailable
}

func (*PerfWidgetExternal) GetMaxCPUSpeed() (maxCpuSpeed uint) {
	js.Rewrite("$<.maxCpuSpeed")
	return maxCpuSpeed
}

func (*PerfWidgetExternal) GetPaintRequestsPerSecond() (paintRequestsPerSecond uint) {
	js.Rewrite("$<.paintRequestsPerSecond")
	return paintRequestsPerSecond
}

func (*PerfWidgetExternal) GetPerformanceCounter() (performanceCounter uint64) {
	js.Rewrite("$<.performanceCounter")
	return performanceCounter
}

func (*PerfWidgetExternal) GetPerformanceCounterFrequency() (performanceCounterFrequency uint64) {
	js.Rewrite("$<.performanceCounterFrequency")
	return performanceCounterFrequency
}

type PeriodicWave struct {
}

type PermissionRequest struct {
	*DeferredPermissionRequest
}

func (*PermissionRequest) Defer() {
	js.Rewrite("$<.defer()")
}

func (*PermissionRequest) GetState() (state *MSWebViewPermissionState) {
	js.Rewrite("$<.state")
	return state
}

type PermissionRequestedEvent struct {
	*Event
}

func (*PermissionRequestedEvent) GetPermissionRequest() (permissionRequest *PermissionRequest) {
	js.Rewrite("$<.permissionRequest")
	return permissionRequest
}

type Plugin struct {
}

func (*Plugin) Item(index uint) (m *MimeType) {
	js.Rewrite("$<.item($1)", index)
	return m
}

func (*Plugin) NamedItem(kind string) (m *MimeType) {
	js.Rewrite("$<.namedItem($1)", kind)
	return m
}

func (*Plugin) GetDescription() (description string) {
	js.Rewrite("$<.description")
	return description
}

func (*Plugin) GetFilename() (filename string) {
	js.Rewrite("$<.filename")
	return filename
}

func (*Plugin) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

func (*Plugin) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*Plugin) GetVersion() (version string) {
	js.Rewrite("$<.version")
	return version
}

type PluginArray struct {
}

func (*PluginArray) Item(index uint) (p *Plugin) {
	js.Rewrite("$<.item($1)", index)
	return p
}

func (*PluginArray) NamedItem(name string) (p *Plugin) {
	js.Rewrite("$<.namedItem($1)", name)
	return p
}

func (*PluginArray) Refresh(reload bool) {
	js.Rewrite("$<.refresh($1)", reload)
}

func (*PluginArray) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

type PointerEvent struct {
	*MouseEvent
}

func (*PointerEvent) GetCurrentPoint(element *Element) {
	js.Rewrite("$<.getCurrentPoint($1)", element)
}

func (*PointerEvent) GetIntermediatePoints(element *Element) {
	js.Rewrite("$<.getIntermediatePoints($1)", element)
}

func (*PointerEvent) InitPointerEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, screenXArg int, screenYArg int, clientXArg float32, clientYArg float32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg uint8, relatedTargetArg *EventTarget, offsetXArg float32, offsetYArg float32, widthArg int, heightArg int, pressure float32, rotation int, tiltX int, tiltY int, pointerIdArg int, pointerType interface{}, hwTimestampArg uint64, isPrimary bool) {
	js.Rewrite("$<.initPointerEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, ctrlKeyArg, altKeyArg, shiftKeyArg, metaKeyArg, buttonArg, relatedTargetArg, offsetXArg, offsetYArg, widthArg, heightArg, pressure, rotation, tiltX, tiltY, pointerIdArg, pointerType, hwTimestampArg, isPrimary)
}

func (*PointerEvent) GetHeight() (height int) {
	js.Rewrite("$<.height")
	return height
}

func (*PointerEvent) GetHwTimestamp() (hwTimestamp uint64) {
	js.Rewrite("$<.hwTimestamp")
	return hwTimestamp
}

func (*PointerEvent) GetIsPrimary() (isPrimary bool) {
	js.Rewrite("$<.isPrimary")
	return isPrimary
}

func (*PointerEvent) GetPointerID() (pointerId int) {
	js.Rewrite("$<.pointerId")
	return pointerId
}

func (*PointerEvent) GetPointerType() (pointerType interface{}) {
	js.Rewrite("$<.pointerType")
	return pointerType
}

func (*PointerEvent) GetPressure() (pressure float32) {
	js.Rewrite("$<.pressure")
	return pressure
}

func (*PointerEvent) GetRotation() (rotation int) {
	js.Rewrite("$<.rotation")
	return rotation
}

func (*PointerEvent) GetTiltX() (tiltX int) {
	js.Rewrite("$<.tiltX")
	return tiltX
}

func (*PointerEvent) GetTiltY() (tiltY int) {
	js.Rewrite("$<.tiltY")
	return tiltY
}

func (*PointerEvent) GetWidth() (width int) {
	js.Rewrite("$<.width")
	return width
}

type PopStateEvent struct {
	*Event
}

func (*PopStateEvent) InitPopStateEvent(typeArg string, canBubbleArg bool, cancelableArg bool, stateArg interface{}) {
	js.Rewrite("$<.initPopStateEvent($1, $2, $3, $4)", typeArg, canBubbleArg, cancelableArg, stateArg)
}

func (*PopStateEvent) GetState() (state interface{}) {
	js.Rewrite("$<.state")
	return state
}

type Position struct {
}

func (*Position) GetCoords() (coords *Coordinates) {
	js.Rewrite("$<.coords")
	return coords
}

func (*Position) GetTimestamp() (timestamp int) {
	js.Rewrite("$<.timestamp")
	return timestamp
}

type PositionError struct {
}

func (*PositionError) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

func (*PositionError) GetCode() (code uint8) {
	js.Rewrite("$<.code")
	return code
}

func (*PositionError) GetMessage() (message string) {
	js.Rewrite("$<.message")
	return message
}

type ProcessingInstruction struct {
	*CharacterData
}

func (*ProcessingInstruction) GetTarget() (target string) {
	js.Rewrite("$<.target")
	return target
}

type ProgressEvent struct {
	*Event
}

func (*ProgressEvent) InitProgressEvent(typeArg string, canBubbleArg bool, cancelableArg bool, lengthComputableArg bool, loadedArg uint64, totalArg uint64) {
	js.Rewrite("$<.initProgressEvent($1, $2, $3, $4, $5, $6)", typeArg, canBubbleArg, cancelableArg, lengthComputableArg, loadedArg, totalArg)
}

func (*ProgressEvent) GetLengthComputable() (lengthComputable bool) {
	js.Rewrite("$<.lengthComputable")
	return lengthComputable
}

func (*ProgressEvent) GetLoaded() (loaded uint64) {
	js.Rewrite("$<.loaded")
	return loaded
}

func (*ProgressEvent) GetTotal() (total uint64) {
	js.Rewrite("$<.total")
	return total
}

type PushManager struct {
}

func (*PushManager) GetSubscription() (p *PushSubscription) {
	js.Rewrite("await $<.getSubscription()")
	return p
}

func (*PushManager) PermissionState(options *PushSubscriptionOptionsInit) (p *PushPermissionState) {
	js.Rewrite("await $<.permissionState($1)", options)
	return p
}

func (*PushManager) Subscribe(options *PushSubscriptionOptionsInit) (p *PushSubscription) {
	js.Rewrite("await $<.subscribe($1)", options)
	return p
}

type PushSubscription struct {
}

func (*PushSubscription) GetKey(name *PushEncryptionKeyName) (b []byte) {
	js.Rewrite("$<.getKey($1)", name)
	return b
}

func (*PushSubscription) ToJSON() (i interface{}) {
	js.Rewrite("$<.toJSON()")
	return i
}

func (*PushSubscription) Unsubscribe() (b bool) {
	js.Rewrite("await $<.unsubscribe()")
	return b
}

func (*PushSubscription) GetEndpoint() (endpoint string) {
	js.Rewrite("$<.endpoint")
	return endpoint
}

func (*PushSubscription) GetOptions() (options *PushSubscriptionOptions) {
	js.Rewrite("$<.options")
	return options
}

type PushSubscriptionOptions struct {
}

func (*PushSubscriptionOptions) GetApplicationServerKey() (applicationServerKey []byte) {
	js.Rewrite("$<.applicationServerKey")
	return applicationServerKey
}

func (*PushSubscriptionOptions) GetUserVisibleOnly() (userVisibleOnly bool) {
	js.Rewrite("$<.userVisibleOnly")
	return userVisibleOnly
}

type Range struct {
}

func (*Range) CloneContents() (d *DocumentFragment) {
	js.Rewrite("$<.cloneContents()")
	return d
}

func (*Range) CloneRange() (r *Range) {
	js.Rewrite("$<.cloneRange()")
	return r
}

func (*Range) Collapse(toStart bool) {
	js.Rewrite("$<.collapse($1)", toStart)
}

func (*Range) CompareBoundaryPoints(how uint8, sourceRange *Range) (i int8) {
	js.Rewrite("$<.compareBoundaryPoints($1, $2)", how, sourceRange)
	return i
}

func (*Range) CreateContextualFragment(fragment string) (d *DocumentFragment) {
	js.Rewrite("$<.createContextualFragment($1)", fragment)
	return d
}

func (*Range) DeleteContents() {
	js.Rewrite("$<.deleteContents()")
}

func (*Range) Detach() {
	js.Rewrite("$<.detach()")
}

func (*Range) Expand(Unit *ExpandGranularity) (b bool) {
	js.Rewrite("$<.expand($1)", Unit)
	return b
}

func (*Range) ExtractContents() (d *DocumentFragment) {
	js.Rewrite("$<.extractContents()")
	return d
}

func (*Range) GetBoundingClientRect() (c *ClientRect) {
	js.Rewrite("$<.getBoundingClientRect()")
	return c
}

func (*Range) GetClientRects() (c *ClientRectList) {
	js.Rewrite("$<.getClientRects()")
	return c
}

func (*Range) InsertNode(newNode *Node) {
	js.Rewrite("$<.insertNode($1)", newNode)
}

func (*Range) SelectNode(refNode *Node) {
	js.Rewrite("$<.selectNode($1)", refNode)
}

func (*Range) SelectNodeContents(refNode *Node) {
	js.Rewrite("$<.selectNodeContents($1)", refNode)
}

func (*Range) SetEnd(refNode *Node, offset int) {
	js.Rewrite("$<.setEnd($1, $2)", refNode, offset)
}

func (*Range) SetEndAfter(refNode *Node) {
	js.Rewrite("$<.setEndAfter($1)", refNode)
}

func (*Range) SetEndBefore(refNode *Node) {
	js.Rewrite("$<.setEndBefore($1)", refNode)
}

func (*Range) SetStart(refNode *Node, offset int) {
	js.Rewrite("$<.setStart($1, $2)", refNode, offset)
}

func (*Range) SetStartAfter(refNode *Node) {
	js.Rewrite("$<.setStartAfter($1)", refNode)
}

func (*Range) SetStartBefore(refNode *Node) {
	js.Rewrite("$<.setStartBefore($1)", refNode)
}

func (*Range) SurroundContents(newParent *Node) {
	js.Rewrite("$<.surroundContents($1)", newParent)
}

func (*Range) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

func (*Range) GetCollapsed() (collapsed bool) {
	js.Rewrite("$<.collapsed")
	return collapsed
}

func (*Range) GetCommonAncestorContainer() (commonAncestorContainer *Node) {
	js.Rewrite("$<.commonAncestorContainer")
	return commonAncestorContainer
}

func (*Range) GetEndContainer() (endContainer *Node) {
	js.Rewrite("$<.endContainer")
	return endContainer
}

func (*Range) GetEndOffset() (endOffset int) {
	js.Rewrite("$<.endOffset")
	return endOffset
}

func (*Range) GetStartContainer() (startContainer *Node) {
	js.Rewrite("$<.startContainer")
	return startContainer
}

func (*Range) GetStartOffset() (startOffset int) {
	js.Rewrite("$<.startOffset")
	return startOffset
}

type ReadableStream struct {
}

func (*ReadableStream) Cancel() {
	js.Rewrite("await $<.cancel()")
}

func (*ReadableStream) GetReader() (r *ReadableStreamReader) {
	js.Rewrite("$<.getReader()")
	return r
}

func (*ReadableStream) GetLocked() (locked bool) {
	js.Rewrite("$<.locked")
	return locked
}

type ReadableStreamReader struct {
}

func (*ReadableStreamReader) Cancel() {
	js.Rewrite("await $<.cancel()")
}

func (*ReadableStreamReader) Read() (i interface{}) {
	js.Rewrite("await $<.read()")
	return i
}

func (*ReadableStreamReader) ReleaseLock() {
	js.Rewrite("$<.releaseLock()")
}

type Request struct {
	*Body
}

func (*Request) Clone() (r *Request) {
	js.Rewrite("$<.clone()")
	return r
}

func (*Request) GetCache() (cache *RequestCache) {
	js.Rewrite("$<.cache")
	return cache
}

func (*Request) GetCredentials() (credentials *RequestCredentials) {
	js.Rewrite("$<.credentials")
	return credentials
}

func (*Request) GetDestination() (destination *RequestDestination) {
	js.Rewrite("$<.destination")
	return destination
}

func (*Request) GetHeaders() (headers *Headers) {
	js.Rewrite("$<.headers")
	return headers
}

func (*Request) GetIntegrity() (integrity string) {
	js.Rewrite("$<.integrity")
	return integrity
}

func (*Request) GetKeepalive() (keepalive bool) {
	js.Rewrite("$<.keepalive")
	return keepalive
}

func (*Request) GetMethod() (method string) {
	js.Rewrite("$<.method")
	return method
}

func (*Request) GetMode() (mode *RequestMode) {
	js.Rewrite("$<.mode")
	return mode
}

func (*Request) GetRedirect() (redirect *RequestRedirect) {
	js.Rewrite("$<.redirect")
	return redirect
}

func (*Request) GetReferrer() (referrer string) {
	js.Rewrite("$<.referrer")
	return referrer
}

func (*Request) GetReferrerPolicy() (referrerPolicy *ReferrerPolicy) {
	js.Rewrite("$<.referrerPolicy")
	return referrerPolicy
}

func (*Request) GetType() (kind *RequestType) {
	js.Rewrite("$<.type")
	return kind
}

func (*Request) GetURL() (url string) {
	js.Rewrite("$<.url")
	return url
}

type Response struct {
	*Body
}

func (*Response) Clone() (r *Response) {
	js.Rewrite("$<.clone()")
	return r
}

func (*Response) GetBody() (body *ReadableStream) {
	js.Rewrite("$<.body")
	return body
}

func (*Response) GetHeaders() (headers *Headers) {
	js.Rewrite("$<.headers")
	return headers
}

func (*Response) GetOk() (ok bool) {
	js.Rewrite("$<.ok")
	return ok
}

func (*Response) GetStatus() (status uint8) {
	js.Rewrite("$<.status")
	return status
}

func (*Response) GetStatusText() (statusText string) {
	js.Rewrite("$<.statusText")
	return statusText
}

func (*Response) GetType() (kind *ResponseType) {
	js.Rewrite("$<.type")
	return kind
}

func (*Response) GetURL() (url string) {
	js.Rewrite("$<.url")
	return url
}

type RTCDtlsTransport struct {
	*RTCStatsProvider
}

func (*RTCDtlsTransport) GetLocalParameters() (r *RTCDtlsParameters) {
	js.Rewrite("$<.getLocalParameters()")
	return r
}

func (*RTCDtlsTransport) GetRemoteCertificates() (b [][]byte) {
	js.Rewrite("$<.getRemoteCertificates()")
	return b
}

func (*RTCDtlsTransport) GetRemoteParameters() (r *RTCDtlsParameters) {
	js.Rewrite("$<.getRemoteParameters()")
	return r
}

func (*RTCDtlsTransport) Start(remoteParameters *RTCDtlsParameters) {
	js.Rewrite("$<.start($1)", remoteParameters)
}

func (*RTCDtlsTransport) Stop() {
	js.Rewrite("$<.stop()")
}

func (*RTCDtlsTransport) GetOndtlsstatechange() (dtlsstatechange *RTCDtlsTransportStateChangedEvent) {
	js.Rewrite("$<.ondtlsstatechange")
	return dtlsstatechange
}

func (*RTCDtlsTransport) SetOndtlsstatechange(dtlsstatechange *RTCDtlsTransportStateChangedEvent) {
	js.Rewrite("$<.ondtlsstatechange = $1", dtlsstatechange)
}

func (*RTCDtlsTransport) GetOnerror() (err *Event) {
	js.Rewrite("$<.onerror")
	return err
}

func (*RTCDtlsTransport) SetOnerror(err *Event) {
	js.Rewrite("$<.onerror = $1", err)
}

func (*RTCDtlsTransport) GetState() (state *RTCDtlsTransportState) {
	js.Rewrite("$<.state")
	return state
}

func (*RTCDtlsTransport) GetTransport() (transport *RTCIceTransport) {
	js.Rewrite("$<.transport")
	return transport
}

type RTCDtlsTransportStateChangedEvent struct {
	*Event
}

func (*RTCDtlsTransportStateChangedEvent) GetState() (state *RTCDtlsTransportState) {
	js.Rewrite("$<.state")
	return state
}

type RTCDtmfSender struct {
	*EventTarget
}

func (*RTCDtmfSender) InsertDTMF(tones string, duration int, interToneGap int) {
	js.Rewrite("$<.insertDTMF($1, $2, $3)", tones, duration, interToneGap)
}

func (*RTCDtmfSender) GetCanInsertDTMF() (canInsertDTMF bool) {
	js.Rewrite("$<.canInsertDTMF")
	return canInsertDTMF
}

func (*RTCDtmfSender) GetDuration() (duration int) {
	js.Rewrite("$<.duration")
	return duration
}

func (*RTCDtmfSender) GetInterToneGap() (interToneGap int) {
	js.Rewrite("$<.interToneGap")
	return interToneGap
}

func (*RTCDtmfSender) GetOntonechange() (tonechange *RTCDTMFToneChangeEvent) {
	js.Rewrite("$<.ontonechange")
	return tonechange
}

func (*RTCDtmfSender) SetOntonechange(tonechange *RTCDTMFToneChangeEvent) {
	js.Rewrite("$<.ontonechange = $1", tonechange)
}

func (*RTCDtmfSender) GetSender() (sender *RTCRtpSender) {
	js.Rewrite("$<.sender")
	return sender
}

func (*RTCDtmfSender) GetToneBuffer() (toneBuffer string) {
	js.Rewrite("$<.toneBuffer")
	return toneBuffer
}

type RTCDTMFToneChangeEvent struct {
	*Event
}

func (*RTCDTMFToneChangeEvent) GetTone() (tone string) {
	js.Rewrite("$<.tone")
	return tone
}

type RTCIceCandidate struct {
}

func (*RTCIceCandidate) ToJSON() (i interface{}) {
	js.Rewrite("$<.toJSON()")
	return i
}

func (*RTCIceCandidate) GetCandidate() (candidate string) {
	js.Rewrite("$<.candidate")
	return candidate
}

func (*RTCIceCandidate) SetCandidate(candidate string) {
	js.Rewrite("$<.candidate = $1", candidate)
}

func (*RTCIceCandidate) GetSdpMid() (sdpMid string) {
	js.Rewrite("$<.sdpMid")
	return sdpMid
}

func (*RTCIceCandidate) SetSdpMid(sdpMid string) {
	js.Rewrite("$<.sdpMid = $1", sdpMid)
}

func (*RTCIceCandidate) GetSdpMLineIndex() (sdpMLineIndex uint8) {
	js.Rewrite("$<.sdpMLineIndex")
	return sdpMLineIndex
}

func (*RTCIceCandidate) SetSdpMLineIndex(sdpMLineIndex uint8) {
	js.Rewrite("$<.sdpMLineIndex = $1", sdpMLineIndex)
}

type RTCIceCandidatePairChangedEvent struct {
	*Event
}

func (*RTCIceCandidatePairChangedEvent) GetPair() (pair *RTCIceCandidatePair) {
	js.Rewrite("$<.pair")
	return pair
}

type RTCIceGatherer struct {
	*RTCStatsProvider
}

func (*RTCIceGatherer) CreateAssociatedGatherer() (r *RTCIceGatherer) {
	js.Rewrite("$<.createAssociatedGatherer()")
	return r
}

func (*RTCIceGatherer) GetLocalCandidates() (r []*RTCIceCandidateDictionary) {
	js.Rewrite("$<.getLocalCandidates()")
	return r
}

func (*RTCIceGatherer) GetLocalParameters() (r *RTCIceParameters) {
	js.Rewrite("$<.getLocalParameters()")
	return r
}

func (*RTCIceGatherer) GetComponent() (component *RTCIceComponent) {
	js.Rewrite("$<.component")
	return component
}

func (*RTCIceGatherer) GetOnerror() (err *Event) {
	js.Rewrite("$<.onerror")
	return err
}

func (*RTCIceGatherer) SetOnerror(err *Event) {
	js.Rewrite("$<.onerror = $1", err)
}

func (*RTCIceGatherer) GetOnlocalcandidate() (localcandidate *RTCIceGathererEvent) {
	js.Rewrite("$<.onlocalcandidate")
	return localcandidate
}

func (*RTCIceGatherer) SetOnlocalcandidate(localcandidate *RTCIceGathererEvent) {
	js.Rewrite("$<.onlocalcandidate = $1", localcandidate)
}

type RTCIceGathererEvent struct {
	*Event
}

func (*RTCIceGathererEvent) GetCandidate() (candidate interface{}) {
	js.Rewrite("$<.candidate")
	return candidate
}

type RTCIceTransport struct {
	*RTCStatsProvider
}

func (*RTCIceTransport) AddRemoteCandidate(remoteCandidate interface{}) {
	js.Rewrite("$<.addRemoteCandidate($1)", remoteCandidate)
}

func (*RTCIceTransport) CreateAssociatedTransport() (r *RTCIceTransport) {
	js.Rewrite("$<.createAssociatedTransport()")
	return r
}

func (*RTCIceTransport) GetNominatedCandidatePair() (r *RTCIceCandidatePair) {
	js.Rewrite("$<.getNominatedCandidatePair()")
	return r
}

func (*RTCIceTransport) GetRemoteCandidates() (r []*RTCIceCandidateDictionary) {
	js.Rewrite("$<.getRemoteCandidates()")
	return r
}

func (*RTCIceTransport) GetRemoteParameters() (r *RTCIceParameters) {
	js.Rewrite("$<.getRemoteParameters()")
	return r
}

func (*RTCIceTransport) SetRemoteCandidates(remoteCandidates []*RTCIceCandidateDictionary) {
	js.Rewrite("$<.setRemoteCandidates($1)", remoteCandidates)
}

func (*RTCIceTransport) Start(gatherer *RTCIceGatherer, remoteParameters *RTCIceParameters, role *RTCIceRole) {
	js.Rewrite("$<.start($1, $2, $3)", gatherer, remoteParameters, role)
}

func (*RTCIceTransport) Stop() {
	js.Rewrite("$<.stop()")
}

func (*RTCIceTransport) GetComponent() (component *RTCIceComponent) {
	js.Rewrite("$<.component")
	return component
}

func (*RTCIceTransport) GetIceGatherer() (iceGatherer *RTCIceGatherer) {
	js.Rewrite("$<.iceGatherer")
	return iceGatherer
}

func (*RTCIceTransport) GetOncandidatepairchange() (candidatepairchange *RTCIceCandidatePairChangedEvent) {
	js.Rewrite("$<.oncandidatepairchange")
	return candidatepairchange
}

func (*RTCIceTransport) SetOncandidatepairchange(candidatepairchange *RTCIceCandidatePairChangedEvent) {
	js.Rewrite("$<.oncandidatepairchange = $1", candidatepairchange)
}

func (*RTCIceTransport) GetOnicestatechange() (icestatechange *RTCIceTransportStateChangedEvent) {
	js.Rewrite("$<.onicestatechange")
	return icestatechange
}

func (*RTCIceTransport) SetOnicestatechange(icestatechange *RTCIceTransportStateChangedEvent) {
	js.Rewrite("$<.onicestatechange = $1", icestatechange)
}

func (*RTCIceTransport) GetRole() (role *RTCIceRole) {
	js.Rewrite("$<.role")
	return role
}

func (*RTCIceTransport) GetState() (state *RTCIceTransportState) {
	js.Rewrite("$<.state")
	return state
}

type RTCIceTransportStateChangedEvent struct {
	*Event
}

func (*RTCIceTransportStateChangedEvent) GetState() (state *RTCIceTransportState) {
	js.Rewrite("$<.state")
	return state
}

type RTCPeerConnection struct {
	*EventTarget
}

func (*RTCPeerConnection) AddIceCandidate(candidate *RTCIceCandidate, successCallback func(), failureCallback func(err *DOMError)) {
	js.Rewrite("await $<.addIceCandidate($1, $2, $3)", candidate, successCallback, failureCallback)
}

func (*RTCPeerConnection) AddStream(stream *MediaStream) {
	js.Rewrite("$<.addStream($1)", stream)
}

func (*RTCPeerConnection) Close() {
	js.Rewrite("$<.close()")
}

func (*RTCPeerConnection) CreateAnswer(successCallback func(sdp *RTCSessionDescription), failureCallback func(err *DOMError)) (r *RTCSessionDescription) {
	js.Rewrite("await $<.createAnswer($1, $2)", successCallback, failureCallback)
	return r
}

func (*RTCPeerConnection) CreateOffer(successCallback func(sdp *RTCSessionDescription), failureCallback func(err *DOMError), options *RTCOfferOptions) (r *RTCSessionDescription) {
	js.Rewrite("await $<.createOffer($1, $2, $3)", successCallback, failureCallback, options)
	return r
}

func (*RTCPeerConnection) GetConfiguration() (r *RTCConfiguration) {
	js.Rewrite("$<.getConfiguration()")
	return r
}

func (*RTCPeerConnection) GetLocalStreams() (m []*MediaStream) {
	js.Rewrite("$<.getLocalStreams()")
	return m
}

func (*RTCPeerConnection) GetRemoteStreams() (m []*MediaStream) {
	js.Rewrite("$<.getRemoteStreams()")
	return m
}

func (*RTCPeerConnection) GetStats(selector *MediaStreamTrack, successCallback func(report *RTCStatsReport), failureCallback func(err *DOMError)) (r *RTCStatsReport) {
	js.Rewrite("await $<.getStats($1, $2, $3)", selector, successCallback, failureCallback)
	return r
}

func (*RTCPeerConnection) GetStreamByID(streamId string) (m *MediaStream) {
	js.Rewrite("$<.getStreamById($1)", streamId)
	return m
}

func (*RTCPeerConnection) RemoveStream(stream *MediaStream) {
	js.Rewrite("$<.removeStream($1)", stream)
}

func (*RTCPeerConnection) SetLocalDescription(description *RTCSessionDescription, successCallback func(), failureCallback func(err *DOMError)) {
	js.Rewrite("await $<.setLocalDescription($1, $2, $3)", description, successCallback, failureCallback)
}

func (*RTCPeerConnection) SetRemoteDescription(description *RTCSessionDescription, successCallback func(), failureCallback func(err *DOMError)) {
	js.Rewrite("await $<.setRemoteDescription($1, $2, $3)", description, successCallback, failureCallback)
}

func (*RTCPeerConnection) GetCanTrickleIceCandidates() (canTrickleIceCandidates bool) {
	js.Rewrite("$<.canTrickleIceCandidates")
	return canTrickleIceCandidates
}

func (*RTCPeerConnection) GetIceConnectionState() (iceConnectionState *RTCIceConnectionState) {
	js.Rewrite("$<.iceConnectionState")
	return iceConnectionState
}

func (*RTCPeerConnection) GetIceGatheringState() (iceGatheringState *RTCIceGatheringState) {
	js.Rewrite("$<.iceGatheringState")
	return iceGatheringState
}

func (*RTCPeerConnection) GetLocalDescription() (localDescription *RTCSessionDescription) {
	js.Rewrite("$<.localDescription")
	return localDescription
}

func (*RTCPeerConnection) GetOnaddstream() (addstream *MediaStreamEvent) {
	js.Rewrite("$<.onaddstream")
	return addstream
}

func (*RTCPeerConnection) SetOnaddstream(addstream *MediaStreamEvent) {
	js.Rewrite("$<.onaddstream = $1", addstream)
}

func (*RTCPeerConnection) GetOnicecandidate() (icecandidate *RTCPeerConnectionIceEvent) {
	js.Rewrite("$<.onicecandidate")
	return icecandidate
}

func (*RTCPeerConnection) SetOnicecandidate(icecandidate *RTCPeerConnectionIceEvent) {
	js.Rewrite("$<.onicecandidate = $1", icecandidate)
}

func (*RTCPeerConnection) GetOniceconnectionstatechange() (iceconnectionstatechange *Event) {
	js.Rewrite("$<.oniceconnectionstatechange")
	return iceconnectionstatechange
}

func (*RTCPeerConnection) SetOniceconnectionstatechange(iceconnectionstatechange *Event) {
	js.Rewrite("$<.oniceconnectionstatechange = $1", iceconnectionstatechange)
}

func (*RTCPeerConnection) GetOnicegatheringstatechange() (icegatheringstatechange *Event) {
	js.Rewrite("$<.onicegatheringstatechange")
	return icegatheringstatechange
}

func (*RTCPeerConnection) SetOnicegatheringstatechange(icegatheringstatechange *Event) {
	js.Rewrite("$<.onicegatheringstatechange = $1", icegatheringstatechange)
}

func (*RTCPeerConnection) GetOnnegotiationneeded() (negotiationneeded *Event) {
	js.Rewrite("$<.onnegotiationneeded")
	return negotiationneeded
}

func (*RTCPeerConnection) SetOnnegotiationneeded(negotiationneeded *Event) {
	js.Rewrite("$<.onnegotiationneeded = $1", negotiationneeded)
}

func (*RTCPeerConnection) GetOnremovestream() (removestream *MediaStreamEvent) {
	js.Rewrite("$<.onremovestream")
	return removestream
}

func (*RTCPeerConnection) SetOnremovestream(removestream *MediaStreamEvent) {
	js.Rewrite("$<.onremovestream = $1", removestream)
}

func (*RTCPeerConnection) GetOnsignalingstatechange() (signalingstatechange *Event) {
	js.Rewrite("$<.onsignalingstatechange")
	return signalingstatechange
}

func (*RTCPeerConnection) SetOnsignalingstatechange(signalingstatechange *Event) {
	js.Rewrite("$<.onsignalingstatechange = $1", signalingstatechange)
}

func (*RTCPeerConnection) GetRemoteDescription() (remoteDescription *RTCSessionDescription) {
	js.Rewrite("$<.remoteDescription")
	return remoteDescription
}

func (*RTCPeerConnection) GetSignalingState() (signalingState *RTCSignalingState) {
	js.Rewrite("$<.signalingState")
	return signalingState
}

type RTCPeerConnectionIceEvent struct {
	*Event
}

func (*RTCPeerConnectionIceEvent) GetCandidate() (candidate *RTCIceCandidate) {
	js.Rewrite("$<.candidate")
	return candidate
}

type RTCRtpReceiver struct {
	*RTCStatsProvider
}

func (*RTCRtpReceiver) GetCapabilities(kind string) (r *RTCRtpCapabilities) {
	js.Rewrite("$<.getCapabilities($1)", kind)
	return r
}

func (*RTCRtpReceiver) GetContributingSources() (r []*RTCRtpContributingSource) {
	js.Rewrite("$<.getContributingSources()")
	return r
}

func (*RTCRtpReceiver) Receive(parameters *RTCRtpParameters) {
	js.Rewrite("$<.receive($1)", parameters)
}

func (*RTCRtpReceiver) RequestSendCSRC(csrc uint) {
	js.Rewrite("$<.requestSendCSRC($1)", csrc)
}

func (*RTCRtpReceiver) SetTransport(transport interface{}, rtcpTransport *RTCDtlsTransport) {
	js.Rewrite("$<.setTransport($1, $2)", transport, rtcpTransport)
}

func (*RTCRtpReceiver) Stop() {
	js.Rewrite("$<.stop()")
}

func (*RTCRtpReceiver) GetOnerror() (err *Event) {
	js.Rewrite("$<.onerror")
	return err
}

func (*RTCRtpReceiver) SetOnerror(err *Event) {
	js.Rewrite("$<.onerror = $1", err)
}

func (*RTCRtpReceiver) GetRtcpTransport() (rtcpTransport *RTCDtlsTransport) {
	js.Rewrite("$<.rtcpTransport")
	return rtcpTransport
}

func (*RTCRtpReceiver) GetTrack() (track *MediaStreamTrack) {
	js.Rewrite("$<.track")
	return track
}

func (*RTCRtpReceiver) GetTransport() (transport interface{}) {
	js.Rewrite("$<.transport")
	return transport
}

type RTCRtpSender struct {
	*RTCStatsProvider
}

func (*RTCRtpSender) GetCapabilities(kind string) (r *RTCRtpCapabilities) {
	js.Rewrite("$<.getCapabilities($1)", kind)
	return r
}

func (*RTCRtpSender) Send(parameters *RTCRtpParameters) {
	js.Rewrite("$<.send($1)", parameters)
}

func (*RTCRtpSender) SetTrack(track *MediaStreamTrack) {
	js.Rewrite("$<.setTrack($1)", track)
}

func (*RTCRtpSender) SetTransport(transport interface{}, rtcpTransport *RTCDtlsTransport) {
	js.Rewrite("$<.setTransport($1, $2)", transport, rtcpTransport)
}

func (*RTCRtpSender) Stop() {
	js.Rewrite("$<.stop()")
}

func (*RTCRtpSender) GetOnerror() (err *Event) {
	js.Rewrite("$<.onerror")
	return err
}

func (*RTCRtpSender) SetOnerror(err *Event) {
	js.Rewrite("$<.onerror = $1", err)
}

func (*RTCRtpSender) GetOnssrcconflict() (ssrcconflict *RTCSsrcConflictEvent) {
	js.Rewrite("$<.onssrcconflict")
	return ssrcconflict
}

func (*RTCRtpSender) SetOnssrcconflict(ssrcconflict *RTCSsrcConflictEvent) {
	js.Rewrite("$<.onssrcconflict = $1", ssrcconflict)
}

func (*RTCRtpSender) GetRtcpTransport() (rtcpTransport *RTCDtlsTransport) {
	js.Rewrite("$<.rtcpTransport")
	return rtcpTransport
}

func (*RTCRtpSender) GetTrack() (track *MediaStreamTrack) {
	js.Rewrite("$<.track")
	return track
}

func (*RTCRtpSender) GetTransport() (transport interface{}) {
	js.Rewrite("$<.transport")
	return transport
}

type RTCSessionDescription struct {
}

func (*RTCSessionDescription) ToJSON() (i interface{}) {
	js.Rewrite("$<.toJSON()")
	return i
}

func (*RTCSessionDescription) GetSdp() (sdp string) {
	js.Rewrite("$<.sdp")
	return sdp
}

func (*RTCSessionDescription) SetSdp(sdp string) {
	js.Rewrite("$<.sdp = $1", sdp)
}

func (*RTCSessionDescription) GetType() (kind *RTCSdpType) {
	js.Rewrite("$<.type")
	return kind
}

func (*RTCSessionDescription) SetType(kind *RTCSdpType) {
	js.Rewrite("$<.type = $1", kind)
}

type RTCSrtpSdesTransport struct {
	*EventTarget
}

func (*RTCSrtpSdesTransport) GetLocalParameters() (r []*RTCSrtpSdesParameters) {
	js.Rewrite("$<.getLocalParameters()")
	return r
}

func (*RTCSrtpSdesTransport) GetOnerror() (err *Event) {
	js.Rewrite("$<.onerror")
	return err
}

func (*RTCSrtpSdesTransport) SetOnerror(err *Event) {
	js.Rewrite("$<.onerror = $1", err)
}

func (*RTCSrtpSdesTransport) GetTransport() (transport *RTCIceTransport) {
	js.Rewrite("$<.transport")
	return transport
}

type RTCSsrcConflictEvent struct {
	*Event
}

func (*RTCSsrcConflictEvent) GetSsrc() (ssrc uint) {
	js.Rewrite("$<.ssrc")
	return ssrc
}

type RTCStatsProvider struct {
	*EventTarget
}

func (*RTCStatsProvider) GetStats() (r *RTCStatsReport) {
	js.Rewrite("await $<.getStats()")
	return r
}

func (*RTCStatsProvider) MsGetStats() (r *RTCStatsReport) {
	js.Rewrite("await $<.msGetStats()")
	return r
}

type ScopedCredential struct {
}

func (*ScopedCredential) GetID() (id []byte) {
	js.Rewrite("$<.id")
	return id
}

func (*ScopedCredential) GetType() (kind *ScopedCredentialType) {
	js.Rewrite("$<.type")
	return kind
}

type ScopedCredentialInfo struct {
}

func (*ScopedCredentialInfo) GetCredential() (credential *ScopedCredential) {
	js.Rewrite("$<.credential")
	return credential
}

func (*ScopedCredentialInfo) GetPublicKey() (publicKey *CryptoKey) {
	js.Rewrite("$<.publicKey")
	return publicKey
}

type Screen struct {
	*EventTarget
}

func (*Screen) MsLockOrientation(orientations interface{}) (b bool) {
	js.Rewrite("$<.msLockOrientation($1)", orientations)
	return b
}

func (*Screen) MsUnlockOrientation() {
	js.Rewrite("$<.msUnlockOrientation()")
}

func (*Screen) GetAvailHeight() (availHeight uint) {
	js.Rewrite("$<.availHeight")
	return availHeight
}

func (*Screen) GetAvailWidth() (availWidth uint) {
	js.Rewrite("$<.availWidth")
	return availWidth
}

func (*Screen) GetBufferDepth() (bufferDepth int) {
	js.Rewrite("$<.bufferDepth")
	return bufferDepth
}

func (*Screen) SetBufferDepth(bufferDepth int) {
	js.Rewrite("$<.bufferDepth = $1", bufferDepth)
}

func (*Screen) GetColorDepth() (colorDepth uint) {
	js.Rewrite("$<.colorDepth")
	return colorDepth
}

func (*Screen) GetDeviceXDPI() (deviceXDPI int) {
	js.Rewrite("$<.deviceXDPI")
	return deviceXDPI
}

func (*Screen) GetDeviceYDPI() (deviceYDPI int) {
	js.Rewrite("$<.deviceYDPI")
	return deviceYDPI
}

func (*Screen) GetFontSmoothingEnabled() (fontSmoothingEnabled bool) {
	js.Rewrite("$<.fontSmoothingEnabled")
	return fontSmoothingEnabled
}

func (*Screen) GetHeight() (height uint) {
	js.Rewrite("$<.height")
	return height
}

func (*Screen) GetLogicalXDPI() (logicalXDPI int) {
	js.Rewrite("$<.logicalXDPI")
	return logicalXDPI
}

func (*Screen) GetLogicalYDPI() (logicalYDPI int) {
	js.Rewrite("$<.logicalYDPI")
	return logicalYDPI
}

func (*Screen) GetMsOrientation() (msOrientation string) {
	js.Rewrite("$<.msOrientation")
	return msOrientation
}

func (*Screen) GetOnmsorientationchange() (MSOrientationChange *Event) {
	js.Rewrite("$<.onmsorientationchange")
	return MSOrientationChange
}

func (*Screen) SetOnmsorientationchange(MSOrientationChange *Event) {
	js.Rewrite("$<.onmsorientationchange = $1", MSOrientationChange)
}

func (*Screen) GetPixelDepth() (pixelDepth uint) {
	js.Rewrite("$<.pixelDepth")
	return pixelDepth
}

func (*Screen) GetSystemXDPI() (systemXDPI int) {
	js.Rewrite("$<.systemXDPI")
	return systemXDPI
}

func (*Screen) GetSystemYDPI() (systemYDPI int) {
	js.Rewrite("$<.systemYDPI")
	return systemYDPI
}

func (*Screen) GetWidth() (width uint) {
	js.Rewrite("$<.width")
	return width
}

type ScriptNotifyEvent struct {
	*Event
}

func (*ScriptNotifyEvent) GetCallingURI() (callingUri string) {
	js.Rewrite("$<.callingUri")
	return callingUri
}

func (*ScriptNotifyEvent) GetValue() (value string) {
	js.Rewrite("$<.value")
	return value
}

type ScriptProcessorNode struct {
	*AudioNode
}

func (*ScriptProcessorNode) GetBufferSize() (bufferSize int) {
	js.Rewrite("$<.bufferSize")
	return bufferSize
}

func (*ScriptProcessorNode) GetOnaudioprocess() (audioprocess *AudioProcessingEvent) {
	js.Rewrite("$<.onaudioprocess")
	return audioprocess
}

func (*ScriptProcessorNode) SetOnaudioprocess(audioprocess *AudioProcessingEvent) {
	js.Rewrite("$<.onaudioprocess = $1", audioprocess)
}

type Selection struct {
}

func (*Selection) AddRange(rng *Range) {
	js.Rewrite("$<.addRange($1)", rng)
}

func (*Selection) Collapse(parentNode *Node, offset int) {
	js.Rewrite("$<.collapse($1, $2)", parentNode, offset)
}

func (*Selection) CollapseToEnd() {
	js.Rewrite("$<.collapseToEnd()")
}

func (*Selection) CollapseToStart() {
	js.Rewrite("$<.collapseToStart()")
}

func (*Selection) ContainsNode(node *Node, partlyContained bool) (b bool) {
	js.Rewrite("$<.containsNode($1, $2)", node, partlyContained)
	return b
}

func (*Selection) DeleteFromDocument() {
	js.Rewrite("$<.deleteFromDocument()")
}

func (*Selection) Empty() {
	js.Rewrite("$<.empty()")
}

func (*Selection) Extend(newNode *Node, offset int) {
	js.Rewrite("$<.extend($1, $2)", newNode, offset)
}

func (*Selection) GetRangeAt(index int) (r *Range) {
	js.Rewrite("$<.getRangeAt($1)", index)
	return r
}

func (*Selection) RemoveAllRanges() {
	js.Rewrite("$<.removeAllRanges()")
}

func (*Selection) RemoveRange(rng *Range) {
	js.Rewrite("$<.removeRange($1)", rng)
}

func (*Selection) SelectAllChildren(parentNode *Node) {
	js.Rewrite("$<.selectAllChildren($1)", parentNode)
}

func (*Selection) SetBaseAndExtent(baseNode *Node, baseOffset int, extentNode *Node, extentOffset int) {
	js.Rewrite("$<.setBaseAndExtent($1, $2, $3, $4)", baseNode, baseOffset, extentNode, extentOffset)
}

func (*Selection) SetPosition(parentNode *Node, offset int) {
	js.Rewrite("$<.setPosition($1, $2)", parentNode, offset)
}

func (*Selection) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

func (*Selection) GetAnchorNode() (anchorNode *Node) {
	js.Rewrite("$<.anchorNode")
	return anchorNode
}

func (*Selection) GetAnchorOffset() (anchorOffset int) {
	js.Rewrite("$<.anchorOffset")
	return anchorOffset
}

func (*Selection) GetBaseNode() (baseNode *Node) {
	js.Rewrite("$<.baseNode")
	return baseNode
}

func (*Selection) GetBaseOffset() (baseOffset int) {
	js.Rewrite("$<.baseOffset")
	return baseOffset
}

func (*Selection) GetExtentNode() (extentNode *Node) {
	js.Rewrite("$<.extentNode")
	return extentNode
}

func (*Selection) GetExtentOffset() (extentOffset int) {
	js.Rewrite("$<.extentOffset")
	return extentOffset
}

func (*Selection) GetFocusNode() (focusNode *Node) {
	js.Rewrite("$<.focusNode")
	return focusNode
}

func (*Selection) GetFocusOffset() (focusOffset int) {
	js.Rewrite("$<.focusOffset")
	return focusOffset
}

func (*Selection) GetIsCollapsed() (isCollapsed bool) {
	js.Rewrite("$<.isCollapsed")
	return isCollapsed
}

func (*Selection) GetRangeCount() (rangeCount int) {
	js.Rewrite("$<.rangeCount")
	return rangeCount
}

func (*Selection) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

type ServiceWorker struct {
	*EventTarget
	*AbstractWorker
}

func (*ServiceWorker) PostMessage(message interface{}, transfer []interface{}) {
	js.Rewrite("$<.postMessage($1, $2)", message, transfer)
}

func (*ServiceWorker) GetOnstatechange() (statechange *Event) {
	js.Rewrite("$<.onstatechange")
	return statechange
}

func (*ServiceWorker) SetOnstatechange(statechange *Event) {
	js.Rewrite("$<.onstatechange = $1", statechange)
}

func (*ServiceWorker) GetScriptURL() (scriptURL string) {
	js.Rewrite("$<.scriptURL")
	return scriptURL
}

func (*ServiceWorker) GetState() (state *ServiceWorkerState) {
	js.Rewrite("$<.state")
	return state
}

type ServiceWorkerContainer struct {
	*EventTarget
}

func (*ServiceWorkerContainer) GetRegistration(clientURL string) (i interface{}) {
	js.Rewrite("await $<.getRegistration($1)", clientURL)
	return i
}

func (*ServiceWorkerContainer) GetRegistrations() (s []*ServiceWorkerRegistration) {
	js.Rewrite("await $<.getRegistrations()")
	return s
}

func (*ServiceWorkerContainer) Register(scriptURL string, options *RegistrationOptions) (s *ServiceWorkerRegistration) {
	js.Rewrite("await $<.register($1, $2)", scriptURL, options)
	return s
}

func (*ServiceWorkerContainer) GetController() (controller *ServiceWorker) {
	js.Rewrite("$<.controller")
	return controller
}

func (*ServiceWorkerContainer) GetOncontrollerchange() (controllerchange *Event) {
	js.Rewrite("$<.oncontrollerchange")
	return controllerchange
}

func (*ServiceWorkerContainer) SetOncontrollerchange(controllerchange *Event) {
	js.Rewrite("$<.oncontrollerchange = $1", controllerchange)
}

func (*ServiceWorkerContainer) GetOnmessage() (message *ServiceWorkerMessageEvent) {
	js.Rewrite("$<.onmessage")
	return message
}

func (*ServiceWorkerContainer) SetOnmessage(message *ServiceWorkerMessageEvent) {
	js.Rewrite("$<.onmessage = $1", message)
}

func (*ServiceWorkerContainer) GetReady() (ready *ServiceWorkerRegistration) {
	js.Rewrite("$<.ready")
	return ready
}

type ServiceWorkerMessageEvent struct {
	*Event
}

func (*ServiceWorkerMessageEvent) GetData() (data interface{}) {
	js.Rewrite("$<.data")
	return data
}

func (*ServiceWorkerMessageEvent) GetLastEventID() (lastEventId string) {
	js.Rewrite("$<.lastEventId")
	return lastEventId
}

func (*ServiceWorkerMessageEvent) GetOrigin() (origin string) {
	js.Rewrite("$<.origin")
	return origin
}

func (*ServiceWorkerMessageEvent) GetPorts() (ports []*MessagePort) {
	js.Rewrite("$<.ports")
	return ports
}

func (*ServiceWorkerMessageEvent) GetSource() (source interface{}) {
	js.Rewrite("$<.source")
	return source
}

type ServiceWorkerRegistration struct {
	*EventTarget
}

func (*ServiceWorkerRegistration) GetNotifications(filter *GetNotificationOptions) (n []*Notification) {
	js.Rewrite("await $<.getNotifications($1)", filter)
	return n
}

func (*ServiceWorkerRegistration) ShowNotification(title string, options *NotificationOptions) {
	js.Rewrite("await $<.showNotification($1, $2)", title, options)
}

func (*ServiceWorkerRegistration) Unregister() (b bool) {
	js.Rewrite("await $<.unregister()")
	return b
}

func (*ServiceWorkerRegistration) Update() {
	js.Rewrite("await $<.update()")
}

func (*ServiceWorkerRegistration) GetActive() (active *ServiceWorker) {
	js.Rewrite("$<.active")
	return active
}

func (*ServiceWorkerRegistration) GetInstalling() (installing *ServiceWorker) {
	js.Rewrite("$<.installing")
	return installing
}

func (*ServiceWorkerRegistration) GetOnupdatefound() (updatefound *Event) {
	js.Rewrite("$<.onupdatefound")
	return updatefound
}

func (*ServiceWorkerRegistration) SetOnupdatefound(updatefound *Event) {
	js.Rewrite("$<.onupdatefound = $1", updatefound)
}

func (*ServiceWorkerRegistration) GetPushManager() (pushManager *PushManager) {
	js.Rewrite("$<.pushManager")
	return pushManager
}

func (*ServiceWorkerRegistration) GetScope() (scope string) {
	js.Rewrite("$<.scope")
	return scope
}

func (*ServiceWorkerRegistration) GetSync() (sync *SyncManager) {
	js.Rewrite("$<.sync")
	return sync
}

func (*ServiceWorkerRegistration) GetWaiting() (waiting *ServiceWorker) {
	js.Rewrite("$<.waiting")
	return waiting
}

type SourceBuffer struct {
	*EventTarget
}

func (*SourceBuffer) Abort() {
	js.Rewrite("$<.abort()")
}

func (*SourceBuffer) AppendBuffer(data interface{}) {
	js.Rewrite("$<.appendBuffer($1)", data)
}

func (*SourceBuffer) AppendStream(stream *MSStream, maxSize uint64) {
	js.Rewrite("$<.appendStream($1, $2)", stream, maxSize)
}

func (*SourceBuffer) Remove(start float32, end float32) {
	js.Rewrite("$<.remove($1, $2)", start, end)
}

func (*SourceBuffer) GetAppendWindowEnd() (appendWindowEnd float32) {
	js.Rewrite("$<.appendWindowEnd")
	return appendWindowEnd
}

func (*SourceBuffer) SetAppendWindowEnd(appendWindowEnd float32) {
	js.Rewrite("$<.appendWindowEnd = $1", appendWindowEnd)
}

func (*SourceBuffer) GetAppendWindowStart() (appendWindowStart float32) {
	js.Rewrite("$<.appendWindowStart")
	return appendWindowStart
}

func (*SourceBuffer) SetAppendWindowStart(appendWindowStart float32) {
	js.Rewrite("$<.appendWindowStart = $1", appendWindowStart)
}

func (*SourceBuffer) GetAudioTracks() (audioTracks *AudioTrackList) {
	js.Rewrite("$<.audioTracks")
	return audioTracks
}

func (*SourceBuffer) GetBuffered() (buffered *TimeRanges) {
	js.Rewrite("$<.buffered")
	return buffered
}

func (*SourceBuffer) GetMode() (mode *AppendMode) {
	js.Rewrite("$<.mode")
	return mode
}

func (*SourceBuffer) SetMode(mode *AppendMode) {
	js.Rewrite("$<.mode = $1", mode)
}

func (*SourceBuffer) GetTimestampOffset() (timestampOffset float32) {
	js.Rewrite("$<.timestampOffset")
	return timestampOffset
}

func (*SourceBuffer) SetTimestampOffset(timestampOffset float32) {
	js.Rewrite("$<.timestampOffset = $1", timestampOffset)
}

func (*SourceBuffer) GetUpdating() (updating bool) {
	js.Rewrite("$<.updating")
	return updating
}

func (*SourceBuffer) GetVideoTracks() (videoTracks *VideoTrackList) {
	js.Rewrite("$<.videoTracks")
	return videoTracks
}

type SourceBufferList struct {
	*EventTarget
}

func (*SourceBufferList) Item(index uint) (s *SourceBuffer) {
	js.Rewrite("$<.item($1)", index)
	return s
}

func (*SourceBufferList) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

type SpeechSynthesis struct {
	*EventTarget
}

func (*SpeechSynthesis) Cancel() {
	js.Rewrite("$<.cancel()")
}

func (*SpeechSynthesis) GetVoices() (s []*SpeechSynthesisVoice) {
	js.Rewrite("$<.getVoices()")
	return s
}

func (*SpeechSynthesis) Pause() {
	js.Rewrite("$<.pause()")
}

func (*SpeechSynthesis) Resume() {
	js.Rewrite("$<.resume()")
}

func (*SpeechSynthesis) Speak(utterance *SpeechSynthesisUtterance) {
	js.Rewrite("$<.speak($1)", utterance)
}

func (*SpeechSynthesis) GetOnvoiceschanged() (voiceschanged *Event) {
	js.Rewrite("$<.onvoiceschanged")
	return voiceschanged
}

func (*SpeechSynthesis) SetOnvoiceschanged(voiceschanged *Event) {
	js.Rewrite("$<.onvoiceschanged = $1", voiceschanged)
}

func (*SpeechSynthesis) GetPaused() (paused bool) {
	js.Rewrite("$<.paused")
	return paused
}

func (*SpeechSynthesis) GetPending() (pending bool) {
	js.Rewrite("$<.pending")
	return pending
}

func (*SpeechSynthesis) GetSpeaking() (speaking bool) {
	js.Rewrite("$<.speaking")
	return speaking
}

type SpeechSynthesisEvent struct {
	*Event
}

func (*SpeechSynthesisEvent) GetCharIndex() (charIndex uint) {
	js.Rewrite("$<.charIndex")
	return charIndex
}

func (*SpeechSynthesisEvent) GetElapsedTime() (elapsedTime float32) {
	js.Rewrite("$<.elapsedTime")
	return elapsedTime
}

func (*SpeechSynthesisEvent) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*SpeechSynthesisEvent) GetUtterance() (utterance *SpeechSynthesisUtterance) {
	js.Rewrite("$<.utterance")
	return utterance
}

type SpeechSynthesisUtterance struct {
	*EventTarget
}

func (*SpeechSynthesisUtterance) GetLang() (lang string) {
	js.Rewrite("$<.lang")
	return lang
}

func (*SpeechSynthesisUtterance) SetLang(lang string) {
	js.Rewrite("$<.lang = $1", lang)
}

func (*SpeechSynthesisUtterance) GetOnboundary() (boundary *Event) {
	js.Rewrite("$<.onboundary")
	return boundary
}

func (*SpeechSynthesisUtterance) SetOnboundary(boundary *Event) {
	js.Rewrite("$<.onboundary = $1", boundary)
}

func (*SpeechSynthesisUtterance) GetOnend() (end *Event) {
	js.Rewrite("$<.onend")
	return end
}

func (*SpeechSynthesisUtterance) SetOnend(end *Event) {
	js.Rewrite("$<.onend = $1", end)
}

func (*SpeechSynthesisUtterance) GetOnerror() (err *Event) {
	js.Rewrite("$<.onerror")
	return err
}

func (*SpeechSynthesisUtterance) SetOnerror(err *Event) {
	js.Rewrite("$<.onerror = $1", err)
}

func (*SpeechSynthesisUtterance) GetOnmark() (mark *Event) {
	js.Rewrite("$<.onmark")
	return mark
}

func (*SpeechSynthesisUtterance) SetOnmark(mark *Event) {
	js.Rewrite("$<.onmark = $1", mark)
}

func (*SpeechSynthesisUtterance) GetOnpause() (pause *Event) {
	js.Rewrite("$<.onpause")
	return pause
}

func (*SpeechSynthesisUtterance) SetOnpause(pause *Event) {
	js.Rewrite("$<.onpause = $1", pause)
}

func (*SpeechSynthesisUtterance) GetOnresume() (resume *Event) {
	js.Rewrite("$<.onresume")
	return resume
}

func (*SpeechSynthesisUtterance) SetOnresume(resume *Event) {
	js.Rewrite("$<.onresume = $1", resume)
}

func (*SpeechSynthesisUtterance) GetOnstart() (start *Event) {
	js.Rewrite("$<.onstart")
	return start
}

func (*SpeechSynthesisUtterance) SetOnstart(start *Event) {
	js.Rewrite("$<.onstart = $1", start)
}

func (*SpeechSynthesisUtterance) GetPitch() (pitch float32) {
	js.Rewrite("$<.pitch")
	return pitch
}

func (*SpeechSynthesisUtterance) SetPitch(pitch float32) {
	js.Rewrite("$<.pitch = $1", pitch)
}

func (*SpeechSynthesisUtterance) GetRate() (rate float32) {
	js.Rewrite("$<.rate")
	return rate
}

func (*SpeechSynthesisUtterance) SetRate(rate float32) {
	js.Rewrite("$<.rate = $1", rate)
}

func (*SpeechSynthesisUtterance) GetText() (text string) {
	js.Rewrite("$<.text")
	return text
}

func (*SpeechSynthesisUtterance) SetText(text string) {
	js.Rewrite("$<.text = $1", text)
}

func (*SpeechSynthesisUtterance) GetVoice() (voice *SpeechSynthesisVoice) {
	js.Rewrite("$<.voice")
	return voice
}

func (*SpeechSynthesisUtterance) SetVoice(voice *SpeechSynthesisVoice) {
	js.Rewrite("$<.voice = $1", voice)
}

func (*SpeechSynthesisUtterance) GetVolume() (volume float32) {
	js.Rewrite("$<.volume")
	return volume
}

func (*SpeechSynthesisUtterance) SetVolume(volume float32) {
	js.Rewrite("$<.volume = $1", volume)
}

type SpeechSynthesisVoice struct {
}

func (*SpeechSynthesisVoice) GetDefault() (def bool) {
	js.Rewrite("$<.default")
	return def
}

func (*SpeechSynthesisVoice) GetLang() (lang string) {
	js.Rewrite("$<.lang")
	return lang
}

func (*SpeechSynthesisVoice) GetLocalService() (localService bool) {
	js.Rewrite("$<.localService")
	return localService
}

func (*SpeechSynthesisVoice) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*SpeechSynthesisVoice) GetVoiceURI() (voiceURI string) {
	js.Rewrite("$<.voiceURI")
	return voiceURI
}

type StereoPannerNode struct {
	*AudioNode
}

func (*StereoPannerNode) GetPan() (pan *AudioParam) {
	js.Rewrite("$<.pan")
	return pan
}

type Storage struct {
}

func (*Storage) Clear() {
	js.Rewrite("$<.clear()")
}

func (*Storage) GetItem(key string) (i interface{}) {
	js.Rewrite("$<.getItem($1)", key)
	return i
}

func (*Storage) Key(index uint) (s string) {
	js.Rewrite("$<.key($1)", index)
	return s
}

func (*Storage) RemoveItem(key string) {
	js.Rewrite("$<.removeItem($1)", key)
}

func (*Storage) SetItem(key string, data string) {
	js.Rewrite("$<.setItem($1, $2)", key, data)
}

func (*Storage) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

type StorageEvent struct {
	*Event
}

func (*StorageEvent) InitStorageEvent(typeArg string, canBubbleArg bool, cancelableArg bool, keyArg string, oldValueArg interface{}, newValueArg interface{}, urlArg string, storageAreaArg *Storage) {
	js.Rewrite("$<.initStorageEvent($1, $2, $3, $4, $5, $6, $7, $8)", typeArg, canBubbleArg, cancelableArg, keyArg, oldValueArg, newValueArg, urlArg, storageAreaArg)
}

func (*StorageEvent) GetKey() (key string) {
	js.Rewrite("$<.key")
	return key
}

func (*StorageEvent) GetNewValue() (newValue interface{}) {
	js.Rewrite("$<.newValue")
	return newValue
}

func (*StorageEvent) GetOldValue() (oldValue interface{}) {
	js.Rewrite("$<.oldValue")
	return oldValue
}

func (*StorageEvent) GetStorageArea() (storageArea *Storage) {
	js.Rewrite("$<.storageArea")
	return storageArea
}

func (*StorageEvent) GetURL() (url string) {
	js.Rewrite("$<.url")
	return url
}

type StyleMedia struct {
}

func (*StyleMedia) MatchMedium(mediaquery string) (b bool) {
	js.Rewrite("$<.matchMedium($1)", mediaquery)
	return b
}

func (*StyleMedia) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

type StyleSheet struct {
}

func (*StyleSheet) GetDisabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

func (*StyleSheet) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

func (*StyleSheet) GetHref() (href string) {
	js.Rewrite("$<.href")
	return href
}

func (*StyleSheet) GetMedia() (media *MediaList) {
	js.Rewrite("$<.media")
	return media
}

func (*StyleSheet) GetOwnerNode() (ownerNode *Node) {
	js.Rewrite("$<.ownerNode")
	return ownerNode
}

func (*StyleSheet) GetParentStyleSheet() (parentStyleSheet *StyleSheet) {
	js.Rewrite("$<.parentStyleSheet")
	return parentStyleSheet
}

func (*StyleSheet) GetTitle() (title string) {
	js.Rewrite("$<.title")
	return title
}

func (*StyleSheet) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

type StyleSheetList struct {
}

func (*StyleSheetList) Item(index uint) (s *StyleSheet) {
	js.Rewrite("$<.item($1)", index)
	return s
}

func (*StyleSheetList) GetLength() (length int) {
	js.Rewrite("$<.length")
	return length
}

type StyleSheetPageList struct {
}

func (*StyleSheetPageList) Item(index uint) (c *CSSPageRule) {
	js.Rewrite("$<.item($1)", index)
	return c
}

func (*StyleSheetPageList) GetLength() (length int) {
	js.Rewrite("$<.length")
	return length
}

type SubtleCrypto struct {
}

func (*SubtleCrypto) Decrypt(algorithm interface{}, key *CryptoKey, data []byte) (i interface{}) {
	js.Rewrite("await $<.decrypt($1, $2, $3)", algorithm, key, data)
	return i
}

func (*SubtleCrypto) DeriveBits(algorithm interface{}, baseKey *CryptoKey, length uint) (i interface{}) {
	js.Rewrite("await $<.deriveBits($1, $2, $3)", algorithm, baseKey, length)
	return i
}

func (*SubtleCrypto) DeriveKey(algorithm interface{}, baseKey *CryptoKey, derivedKeyType interface{}, extractable bool, keyUsages []string) (i interface{}) {
	js.Rewrite("await $<.deriveKey($1, $2, $3, $4, $5)", algorithm, baseKey, derivedKeyType, extractable, keyUsages)
	return i
}

func (*SubtleCrypto) Digest(algorithm interface{}, data []byte) (i interface{}) {
	js.Rewrite("await $<.digest($1, $2)", algorithm, data)
	return i
}

func (*SubtleCrypto) Encrypt(algorithm interface{}, key *CryptoKey, data []byte) (i interface{}) {
	js.Rewrite("await $<.encrypt($1, $2, $3)", algorithm, key, data)
	return i
}

func (*SubtleCrypto) ExportKey(format string, key *CryptoKey) (i interface{}) {
	js.Rewrite("await $<.exportKey($1, $2)", format, key)
	return i
}

func (*SubtleCrypto) GenerateKey(algorithm interface{}, extractable bool, keyUsages []string) (i interface{}) {
	js.Rewrite("await $<.generateKey($1, $2, $3)", algorithm, extractable, keyUsages)
	return i
}

func (*SubtleCrypto) ImportKey(format string, keyData []byte, algorithm interface{}, extractable bool, keyUsages []string) (i interface{}) {
	js.Rewrite("await $<.importKey($1, $2, $3, $4, $5)", format, keyData, algorithm, extractable, keyUsages)
	return i
}

func (*SubtleCrypto) Sign(algorithm interface{}, key *CryptoKey, data []byte) (i interface{}) {
	js.Rewrite("await $<.sign($1, $2, $3)", algorithm, key, data)
	return i
}

func (*SubtleCrypto) UnwrapKey(format string, wrappedKey []byte, unwrappingKey *CryptoKey, unwrapAlgorithm interface{}, unwrappedKeyAlgorithm interface{}, extractable bool, keyUsages []string) (i interface{}) {
	js.Rewrite("await $<.unwrapKey($1, $2, $3, $4, $5, $6, $7)", format, wrappedKey, unwrappingKey, unwrapAlgorithm, unwrappedKeyAlgorithm, extractable, keyUsages)
	return i
}

func (*SubtleCrypto) Verify(algorithm interface{}, key *CryptoKey, signature []byte, data []byte) (i interface{}) {
	js.Rewrite("await $<.verify($1, $2, $3, $4)", algorithm, key, signature, data)
	return i
}

func (*SubtleCrypto) WrapKey(format string, key *CryptoKey, wrappingKey *CryptoKey, wrapAlgorithm interface{}) (i interface{}) {
	js.Rewrite("await $<.wrapKey($1, $2, $3, $4)", format, key, wrappingKey, wrapAlgorithm)
	return i
}

type SVGAElement struct {
	*SVGGraphicsElement
	*SVGURIReference
}

func (*SVGAElement) GetTarget() (target *SVGAnimatedString) {
	js.Rewrite("$<.target")
	return target
}

type SVGAngle struct {
}

func (*SVGAngle) ConvertToSpecifiedUnits(unitType uint8) {
	js.Rewrite("$<.convertToSpecifiedUnits($1)", unitType)
}

func (*SVGAngle) NewValueSpecifiedUnits(unitType uint8, valueInSpecifiedUnits float32) {
	js.Rewrite("$<.newValueSpecifiedUnits($1, $2)", unitType, valueInSpecifiedUnits)
}

func (*SVGAngle) GetUnitType() (unitType uint8) {
	js.Rewrite("$<.unitType")
	return unitType
}

func (*SVGAngle) GetValue() (value float32) {
	js.Rewrite("$<.value")
	return value
}

func (*SVGAngle) SetValue(value float32) {
	js.Rewrite("$<.value = $1", value)
}

func (*SVGAngle) GetValueAsString() (valueAsString string) {
	js.Rewrite("$<.valueAsString")
	return valueAsString
}

func (*SVGAngle) SetValueAsString(valueAsString string) {
	js.Rewrite("$<.valueAsString = $1", valueAsString)
}

func (*SVGAngle) GetValueInSpecifiedUnits() (valueInSpecifiedUnits float32) {
	js.Rewrite("$<.valueInSpecifiedUnits")
	return valueInSpecifiedUnits
}

func (*SVGAngle) SetValueInSpecifiedUnits(valueInSpecifiedUnits float32) {
	js.Rewrite("$<.valueInSpecifiedUnits = $1", valueInSpecifiedUnits)
}

type SVGAnimatedAngle struct {
}

func (*SVGAnimatedAngle) GetAnimVal() (animVal *SVGAngle) {
	js.Rewrite("$<.animVal")
	return animVal
}

func (*SVGAnimatedAngle) GetBaseVal() (baseVal *SVGAngle) {
	js.Rewrite("$<.baseVal")
	return baseVal
}

type SVGAnimatedBoolean struct {
}

func (*SVGAnimatedBoolean) GetAnimVal() (animVal bool) {
	js.Rewrite("$<.animVal")
	return animVal
}

func (*SVGAnimatedBoolean) GetBaseVal() (baseVal bool) {
	js.Rewrite("$<.baseVal")
	return baseVal
}

func (*SVGAnimatedBoolean) SetBaseVal(baseVal bool) {
	js.Rewrite("$<.baseVal = $1", baseVal)
}

type SVGAnimatedEnumeration struct {
}

func (*SVGAnimatedEnumeration) GetAnimVal() (animVal uint8) {
	js.Rewrite("$<.animVal")
	return animVal
}

func (*SVGAnimatedEnumeration) GetBaseVal() (baseVal uint8) {
	js.Rewrite("$<.baseVal")
	return baseVal
}

func (*SVGAnimatedEnumeration) SetBaseVal(baseVal uint8) {
	js.Rewrite("$<.baseVal = $1", baseVal)
}

type SVGAnimatedInteger struct {
}

func (*SVGAnimatedInteger) GetAnimVal() (animVal int) {
	js.Rewrite("$<.animVal")
	return animVal
}

func (*SVGAnimatedInteger) GetBaseVal() (baseVal int) {
	js.Rewrite("$<.baseVal")
	return baseVal
}

func (*SVGAnimatedInteger) SetBaseVal(baseVal int) {
	js.Rewrite("$<.baseVal = $1", baseVal)
}

type SVGAnimatedLength struct {
}

func (*SVGAnimatedLength) GetAnimVal() (animVal *SVGLength) {
	js.Rewrite("$<.animVal")
	return animVal
}

func (*SVGAnimatedLength) GetBaseVal() (baseVal *SVGLength) {
	js.Rewrite("$<.baseVal")
	return baseVal
}

type SVGAnimatedLengthList struct {
}

func (*SVGAnimatedLengthList) GetAnimVal() (animVal *SVGLengthList) {
	js.Rewrite("$<.animVal")
	return animVal
}

func (*SVGAnimatedLengthList) GetBaseVal() (baseVal *SVGLengthList) {
	js.Rewrite("$<.baseVal")
	return baseVal
}

type SVGAnimatedNumber struct {
}

func (*SVGAnimatedNumber) GetAnimVal() (animVal float32) {
	js.Rewrite("$<.animVal")
	return animVal
}

func (*SVGAnimatedNumber) GetBaseVal() (baseVal float32) {
	js.Rewrite("$<.baseVal")
	return baseVal
}

func (*SVGAnimatedNumber) SetBaseVal(baseVal float32) {
	js.Rewrite("$<.baseVal = $1", baseVal)
}

type SVGAnimatedNumberList struct {
}

func (*SVGAnimatedNumberList) GetAnimVal() (animVal *SVGNumberList) {
	js.Rewrite("$<.animVal")
	return animVal
}

func (*SVGAnimatedNumberList) GetBaseVal() (baseVal *SVGNumberList) {
	js.Rewrite("$<.baseVal")
	return baseVal
}

type SVGAnimatedPreserveAspectRatio struct {
}

func (*SVGAnimatedPreserveAspectRatio) GetAnimVal() (animVal *SVGPreserveAspectRatio) {
	js.Rewrite("$<.animVal")
	return animVal
}

func (*SVGAnimatedPreserveAspectRatio) GetBaseVal() (baseVal *SVGPreserveAspectRatio) {
	js.Rewrite("$<.baseVal")
	return baseVal
}

type SVGAnimatedRect struct {
}

func (*SVGAnimatedRect) GetAnimVal() (animVal *SVGRect) {
	js.Rewrite("$<.animVal")
	return animVal
}

func (*SVGAnimatedRect) GetBaseVal() (baseVal *SVGRect) {
	js.Rewrite("$<.baseVal")
	return baseVal
}

type SVGAnimatedString struct {
}

func (*SVGAnimatedString) GetAnimVal() (animVal string) {
	js.Rewrite("$<.animVal")
	return animVal
}

func (*SVGAnimatedString) GetBaseVal() (baseVal string) {
	js.Rewrite("$<.baseVal")
	return baseVal
}

func (*SVGAnimatedString) SetBaseVal(baseVal string) {
	js.Rewrite("$<.baseVal = $1", baseVal)
}

type SVGAnimatedTransformList struct {
}

func (*SVGAnimatedTransformList) GetAnimVal() (animVal *SVGTransformList) {
	js.Rewrite("$<.animVal")
	return animVal
}

func (*SVGAnimatedTransformList) GetBaseVal() (baseVal *SVGTransformList) {
	js.Rewrite("$<.baseVal")
	return baseVal
}

type SVGCircleElement struct {
	*SVGGraphicsElement
}

func (*SVGCircleElement) GetCx() (cx *SVGAnimatedLength) {
	js.Rewrite("$<.cx")
	return cx
}

func (*SVGCircleElement) GetCy() (cy *SVGAnimatedLength) {
	js.Rewrite("$<.cy")
	return cy
}

func (*SVGCircleElement) GetR() (r *SVGAnimatedLength) {
	js.Rewrite("$<.r")
	return r
}

type SVGClipPathElement struct {
	*SVGGraphicsElement
	*SVGUnitTypes
}

func (*SVGClipPathElement) GetClipPathUnits() (clipPathUnits *SVGAnimatedEnumeration) {
	js.Rewrite("$<.clipPathUnits")
	return clipPathUnits
}

type SVGComponentTransferFunctionElement struct {
	*SVGElement
}

func (*SVGComponentTransferFunctionElement) GetAmplitude() (amplitude *SVGAnimatedNumber) {
	js.Rewrite("$<.amplitude")
	return amplitude
}

func (*SVGComponentTransferFunctionElement) GetExponent() (exponent *SVGAnimatedNumber) {
	js.Rewrite("$<.exponent")
	return exponent
}

func (*SVGComponentTransferFunctionElement) GetIntercept() (intercept *SVGAnimatedNumber) {
	js.Rewrite("$<.intercept")
	return intercept
}

func (*SVGComponentTransferFunctionElement) GetOffset() (offset *SVGAnimatedNumber) {
	js.Rewrite("$<.offset")
	return offset
}

func (*SVGComponentTransferFunctionElement) GetSlope() (slope *SVGAnimatedNumber) {
	js.Rewrite("$<.slope")
	return slope
}

func (*SVGComponentTransferFunctionElement) GetTableValues() (tableValues *SVGAnimatedNumberList) {
	js.Rewrite("$<.tableValues")
	return tableValues
}

func (*SVGComponentTransferFunctionElement) GetType() (kind *SVGAnimatedEnumeration) {
	js.Rewrite("$<.type")
	return kind
}

type SVGDefsElement struct {
	*SVGGraphicsElement
}

type SVGDescElement struct {
	*SVGElement
}

type SVGElement struct {
	*Element
}

func (*SVGElement) GetClassName() (className *SVGAnimatedString) {
	js.Rewrite("$<.className")
	return className
}

func (*SVGElement) GetOnclick() (click *MouseEvent) {
	js.Rewrite("$<.onclick")
	return click
}

func (*SVGElement) SetOnclick(click *MouseEvent) {
	js.Rewrite("$<.onclick = $1", click)
}

func (*SVGElement) GetOndblclick() (dblclick *MouseEvent) {
	js.Rewrite("$<.ondblclick")
	return dblclick
}

func (*SVGElement) SetOndblclick(dblclick *MouseEvent) {
	js.Rewrite("$<.ondblclick = $1", dblclick)
}

func (*SVGElement) GetOnfocusin() (focusin *FocusEvent) {
	js.Rewrite("$<.onfocusin")
	return focusin
}

func (*SVGElement) SetOnfocusin(focusin *FocusEvent) {
	js.Rewrite("$<.onfocusin = $1", focusin)
}

func (*SVGElement) GetOnfocusout() (focusout *FocusEvent) {
	js.Rewrite("$<.onfocusout")
	return focusout
}

func (*SVGElement) SetOnfocusout(focusout *FocusEvent) {
	js.Rewrite("$<.onfocusout = $1", focusout)
}

func (*SVGElement) GetOnload() (e *Event) {
	js.Rewrite("$<.onload")
	return e
}

func (*SVGElement) SetOnload(e *Event) {
	js.Rewrite("$<.onload = $1", e)
}

func (*SVGElement) GetOnmousedown() (mousedown *MouseEvent) {
	js.Rewrite("$<.onmousedown")
	return mousedown
}

func (*SVGElement) SetOnmousedown(mousedown *MouseEvent) {
	js.Rewrite("$<.onmousedown = $1", mousedown)
}

func (*SVGElement) GetOnmousemove() (mousemove *MouseEvent) {
	js.Rewrite("$<.onmousemove")
	return mousemove
}

func (*SVGElement) SetOnmousemove(mousemove *MouseEvent) {
	js.Rewrite("$<.onmousemove = $1", mousemove)
}

func (*SVGElement) GetOnmouseout() (mouseout *MouseEvent) {
	js.Rewrite("$<.onmouseout")
	return mouseout
}

func (*SVGElement) SetOnmouseout(mouseout *MouseEvent) {
	js.Rewrite("$<.onmouseout = $1", mouseout)
}

func (*SVGElement) GetOnmouseover() (mouseover *MouseEvent) {
	js.Rewrite("$<.onmouseover")
	return mouseover
}

func (*SVGElement) SetOnmouseover(mouseover *MouseEvent) {
	js.Rewrite("$<.onmouseover = $1", mouseover)
}

func (*SVGElement) GetOnmouseup() (mouseup *MouseEvent) {
	js.Rewrite("$<.onmouseup")
	return mouseup
}

func (*SVGElement) SetOnmouseup(mouseup *MouseEvent) {
	js.Rewrite("$<.onmouseup = $1", mouseup)
}

func (*SVGElement) GetOwnerSVGElement() (ownerSVGElement *SVGSVGElement) {
	js.Rewrite("$<.ownerSVGElement")
	return ownerSVGElement
}

func (*SVGElement) GetStyle() (style *CSSStyleDeclaration) {
	js.Rewrite("$<.style")
	return style
}

func (*SVGElement) GetViewportElement() (viewportElement *SVGElement) {
	js.Rewrite("$<.viewportElement")
	return viewportElement
}

func (*SVGElement) GetXmlbase() (xmlbase string) {
	js.Rewrite("$<.xmlbase")
	return xmlbase
}

func (*SVGElement) SetXmlbase(xmlbase string) {
	js.Rewrite("$<.xmlbase = $1", xmlbase)
}

type SVGElementInstance struct {
	*EventTarget
}

func (*SVGElementInstance) GetChildNodes() (childNodes *SVGElementInstanceList) {
	js.Rewrite("$<.childNodes")
	return childNodes
}

func (*SVGElementInstance) GetCorrespondingElement() (correspondingElement *SVGElement) {
	js.Rewrite("$<.correspondingElement")
	return correspondingElement
}

func (*SVGElementInstance) GetCorrespondingUseElement() (correspondingUseElement *SVGUseElement) {
	js.Rewrite("$<.correspondingUseElement")
	return correspondingUseElement
}

func (*SVGElementInstance) GetFirstChild() (firstChild *SVGElementInstance) {
	js.Rewrite("$<.firstChild")
	return firstChild
}

func (*SVGElementInstance) GetLastChild() (lastChild *SVGElementInstance) {
	js.Rewrite("$<.lastChild")
	return lastChild
}

func (*SVGElementInstance) GetNextSibling() (nextSibling *SVGElementInstance) {
	js.Rewrite("$<.nextSibling")
	return nextSibling
}

func (*SVGElementInstance) GetParentNode() (parentNode *SVGElementInstance) {
	js.Rewrite("$<.parentNode")
	return parentNode
}

func (*SVGElementInstance) GetPreviousSibling() (previousSibling *SVGElementInstance) {
	js.Rewrite("$<.previousSibling")
	return previousSibling
}

type SVGElementInstanceList struct {
}

func (*SVGElementInstanceList) Item(index uint) (s *SVGElementInstance) {
	js.Rewrite("$<.item($1)", index)
	return s
}

func (*SVGElementInstanceList) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

type SVGEllipseElement struct {
	*SVGGraphicsElement
}

func (*SVGEllipseElement) GetCx() (cx *SVGAnimatedLength) {
	js.Rewrite("$<.cx")
	return cx
}

func (*SVGEllipseElement) GetCy() (cy *SVGAnimatedLength) {
	js.Rewrite("$<.cy")
	return cy
}

func (*SVGEllipseElement) GetRx() (rx *SVGAnimatedLength) {
	js.Rewrite("$<.rx")
	return rx
}

func (*SVGEllipseElement) GetRy() (ry *SVGAnimatedLength) {
	js.Rewrite("$<.ry")
	return ry
}

type SVGFEBlendElement struct {
	*SVGElement
	*SVGFilterPrimitiveStandardAttributes
}

func (*SVGFEBlendElement) GetIn1() (in1 *SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}

func (*SVGFEBlendElement) GetIn2() (in2 *SVGAnimatedString) {
	js.Rewrite("$<.in2")
	return in2
}

func (*SVGFEBlendElement) GetMode() (mode *SVGAnimatedEnumeration) {
	js.Rewrite("$<.mode")
	return mode
}

type SVGFEColorMatrixElement struct {
	*SVGElement
	*SVGFilterPrimitiveStandardAttributes
}

func (*SVGFEColorMatrixElement) GetIn1() (in1 *SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}

func (*SVGFEColorMatrixElement) GetType() (kind *SVGAnimatedEnumeration) {
	js.Rewrite("$<.type")
	return kind
}

func (*SVGFEColorMatrixElement) GetValues() (values *SVGAnimatedNumberList) {
	js.Rewrite("$<.values")
	return values
}

type SVGFEComponentTransferElement struct {
	*SVGElement
	*SVGFilterPrimitiveStandardAttributes
}

func (*SVGFEComponentTransferElement) GetIn1() (in1 *SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}

type SVGFECompositeElement struct {
	*SVGElement
	*SVGFilterPrimitiveStandardAttributes
}

func (*SVGFECompositeElement) GetIn1() (in1 *SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}

func (*SVGFECompositeElement) GetIn2() (in2 *SVGAnimatedString) {
	js.Rewrite("$<.in2")
	return in2
}

func (*SVGFECompositeElement) GetK1() (k1 *SVGAnimatedNumber) {
	js.Rewrite("$<.k1")
	return k1
}

func (*SVGFECompositeElement) GetK2() (k2 *SVGAnimatedNumber) {
	js.Rewrite("$<.k2")
	return k2
}

func (*SVGFECompositeElement) GetK3() (k3 *SVGAnimatedNumber) {
	js.Rewrite("$<.k3")
	return k3
}

func (*SVGFECompositeElement) GetK4() (k4 *SVGAnimatedNumber) {
	js.Rewrite("$<.k4")
	return k4
}

func (*SVGFECompositeElement) GetOperator() (operator *SVGAnimatedEnumeration) {
	js.Rewrite("$<.operator")
	return operator
}

type SVGFEConvolveMatrixElement struct {
	*SVGElement
	*SVGFilterPrimitiveStandardAttributes
}

func (*SVGFEConvolveMatrixElement) GetBias() (bias *SVGAnimatedNumber) {
	js.Rewrite("$<.bias")
	return bias
}

func (*SVGFEConvolveMatrixElement) GetDivisor() (divisor *SVGAnimatedNumber) {
	js.Rewrite("$<.divisor")
	return divisor
}

func (*SVGFEConvolveMatrixElement) GetEdgeMode() (edgeMode *SVGAnimatedEnumeration) {
	js.Rewrite("$<.edgeMode")
	return edgeMode
}

func (*SVGFEConvolveMatrixElement) GetIn1() (in1 *SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}

func (*SVGFEConvolveMatrixElement) GetKernelMatrix() (kernelMatrix *SVGAnimatedNumberList) {
	js.Rewrite("$<.kernelMatrix")
	return kernelMatrix
}

func (*SVGFEConvolveMatrixElement) GetKernelUnitLengthX() (kernelUnitLengthX *SVGAnimatedNumber) {
	js.Rewrite("$<.kernelUnitLengthX")
	return kernelUnitLengthX
}

func (*SVGFEConvolveMatrixElement) GetKernelUnitLengthY() (kernelUnitLengthY *SVGAnimatedNumber) {
	js.Rewrite("$<.kernelUnitLengthY")
	return kernelUnitLengthY
}

func (*SVGFEConvolveMatrixElement) GetOrderX() (orderX *SVGAnimatedInteger) {
	js.Rewrite("$<.orderX")
	return orderX
}

func (*SVGFEConvolveMatrixElement) GetOrderY() (orderY *SVGAnimatedInteger) {
	js.Rewrite("$<.orderY")
	return orderY
}

func (*SVGFEConvolveMatrixElement) GetPreserveAlpha() (preserveAlpha *SVGAnimatedBoolean) {
	js.Rewrite("$<.preserveAlpha")
	return preserveAlpha
}

func (*SVGFEConvolveMatrixElement) GetTargetX() (targetX *SVGAnimatedInteger) {
	js.Rewrite("$<.targetX")
	return targetX
}

func (*SVGFEConvolveMatrixElement) GetTargetY() (targetY *SVGAnimatedInteger) {
	js.Rewrite("$<.targetY")
	return targetY
}

type SVGFEDiffuseLightingElement struct {
	*SVGElement
	*SVGFilterPrimitiveStandardAttributes
}

func (*SVGFEDiffuseLightingElement) GetDiffuseConstant() (diffuseConstant *SVGAnimatedNumber) {
	js.Rewrite("$<.diffuseConstant")
	return diffuseConstant
}

func (*SVGFEDiffuseLightingElement) GetIn1() (in1 *SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}

func (*SVGFEDiffuseLightingElement) GetKernelUnitLengthX() (kernelUnitLengthX *SVGAnimatedNumber) {
	js.Rewrite("$<.kernelUnitLengthX")
	return kernelUnitLengthX
}

func (*SVGFEDiffuseLightingElement) GetKernelUnitLengthY() (kernelUnitLengthY *SVGAnimatedNumber) {
	js.Rewrite("$<.kernelUnitLengthY")
	return kernelUnitLengthY
}

func (*SVGFEDiffuseLightingElement) GetSurfaceScale() (surfaceScale *SVGAnimatedNumber) {
	js.Rewrite("$<.surfaceScale")
	return surfaceScale
}

type SVGFEDisplacementMapElement struct {
	*SVGElement
	*SVGFilterPrimitiveStandardAttributes
}

func (*SVGFEDisplacementMapElement) GetIn1() (in1 *SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}

func (*SVGFEDisplacementMapElement) GetIn2() (in2 *SVGAnimatedString) {
	js.Rewrite("$<.in2")
	return in2
}

func (*SVGFEDisplacementMapElement) GetScale() (scale *SVGAnimatedNumber) {
	js.Rewrite("$<.scale")
	return scale
}

func (*SVGFEDisplacementMapElement) GetXChannelSelector() (xChannelSelector *SVGAnimatedEnumeration) {
	js.Rewrite("$<.xChannelSelector")
	return xChannelSelector
}

func (*SVGFEDisplacementMapElement) GetYChannelSelector() (yChannelSelector *SVGAnimatedEnumeration) {
	js.Rewrite("$<.yChannelSelector")
	return yChannelSelector
}

type SVGFEDistantLightElement struct {
	*SVGElement
}

func (*SVGFEDistantLightElement) GetAzimuth() (azimuth *SVGAnimatedNumber) {
	js.Rewrite("$<.azimuth")
	return azimuth
}

func (*SVGFEDistantLightElement) GetElevation() (elevation *SVGAnimatedNumber) {
	js.Rewrite("$<.elevation")
	return elevation
}

type SVGFEFloodElement struct {
	*SVGElement
	*SVGFilterPrimitiveStandardAttributes
}

type SVGFEFuncAElement struct {
	*SVGComponentTransferFunctionElement
}

type SVGFEFuncBElement struct {
	*SVGComponentTransferFunctionElement
}

type SVGFEFuncGElement struct {
	*SVGComponentTransferFunctionElement
}

type SVGFEFuncRElement struct {
	*SVGComponentTransferFunctionElement
}

type SVGFEGaussianBlurElement struct {
	*SVGElement
	*SVGFilterPrimitiveStandardAttributes
}

func (*SVGFEGaussianBlurElement) SetStdDeviation(stdDeviationX float32, stdDeviationY float32) {
	js.Rewrite("$<.setStdDeviation($1, $2)", stdDeviationX, stdDeviationY)
}

func (*SVGFEGaussianBlurElement) GetIn1() (in1 *SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}

func (*SVGFEGaussianBlurElement) GetStdDeviationX() (stdDeviationX *SVGAnimatedNumber) {
	js.Rewrite("$<.stdDeviationX")
	return stdDeviationX
}

func (*SVGFEGaussianBlurElement) GetStdDeviationY() (stdDeviationY *SVGAnimatedNumber) {
	js.Rewrite("$<.stdDeviationY")
	return stdDeviationY
}

type SVGFEImageElement struct {
	*SVGElement
	*SVGFilterPrimitiveStandardAttributes
	*SVGURIReference
}

func (*SVGFEImageElement) GetPreserveAspectRatio() (preserveAspectRatio *SVGAnimatedPreserveAspectRatio) {
	js.Rewrite("$<.preserveAspectRatio")
	return preserveAspectRatio
}

type SVGFEMergeElement struct {
	*SVGElement
	*SVGFilterPrimitiveStandardAttributes
}

type SVGFEMergeNodeElement struct {
	*SVGElement
}

func (*SVGFEMergeNodeElement) GetIn1() (in1 *SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}

type SVGFEMorphologyElement struct {
	*SVGElement
	*SVGFilterPrimitiveStandardAttributes
}

func (*SVGFEMorphologyElement) GetIn1() (in1 *SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}

func (*SVGFEMorphologyElement) GetOperator() (operator *SVGAnimatedEnumeration) {
	js.Rewrite("$<.operator")
	return operator
}

func (*SVGFEMorphologyElement) GetRadiusX() (radiusX *SVGAnimatedNumber) {
	js.Rewrite("$<.radiusX")
	return radiusX
}

func (*SVGFEMorphologyElement) GetRadiusY() (radiusY *SVGAnimatedNumber) {
	js.Rewrite("$<.radiusY")
	return radiusY
}

type SVGFEOffsetElement struct {
	*SVGElement
	*SVGFilterPrimitiveStandardAttributes
}

func (*SVGFEOffsetElement) GetDx() (dx *SVGAnimatedNumber) {
	js.Rewrite("$<.dx")
	return dx
}

func (*SVGFEOffsetElement) GetDy() (dy *SVGAnimatedNumber) {
	js.Rewrite("$<.dy")
	return dy
}

func (*SVGFEOffsetElement) GetIn1() (in1 *SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}

type SVGFEPointLightElement struct {
	*SVGElement
}

func (*SVGFEPointLightElement) GetX() (x *SVGAnimatedNumber) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGFEPointLightElement) GetY() (y *SVGAnimatedNumber) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGFEPointLightElement) GetZ() (z *SVGAnimatedNumber) {
	js.Rewrite("$<.z")
	return z
}

type SVGFESpecularLightingElement struct {
	*SVGElement
	*SVGFilterPrimitiveStandardAttributes
}

func (*SVGFESpecularLightingElement) GetIn1() (in1 *SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}

func (*SVGFESpecularLightingElement) GetKernelUnitLengthX() (kernelUnitLengthX *SVGAnimatedNumber) {
	js.Rewrite("$<.kernelUnitLengthX")
	return kernelUnitLengthX
}

func (*SVGFESpecularLightingElement) GetKernelUnitLengthY() (kernelUnitLengthY *SVGAnimatedNumber) {
	js.Rewrite("$<.kernelUnitLengthY")
	return kernelUnitLengthY
}

func (*SVGFESpecularLightingElement) GetSpecularConstant() (specularConstant *SVGAnimatedNumber) {
	js.Rewrite("$<.specularConstant")
	return specularConstant
}

func (*SVGFESpecularLightingElement) GetSpecularExponent() (specularExponent *SVGAnimatedNumber) {
	js.Rewrite("$<.specularExponent")
	return specularExponent
}

func (*SVGFESpecularLightingElement) GetSurfaceScale() (surfaceScale *SVGAnimatedNumber) {
	js.Rewrite("$<.surfaceScale")
	return surfaceScale
}

type SVGFESpotLightElement struct {
	*SVGElement
}

func (*SVGFESpotLightElement) GetLimitingConeAngle() (limitingConeAngle *SVGAnimatedNumber) {
	js.Rewrite("$<.limitingConeAngle")
	return limitingConeAngle
}

func (*SVGFESpotLightElement) GetPointsAtX() (pointsAtX *SVGAnimatedNumber) {
	js.Rewrite("$<.pointsAtX")
	return pointsAtX
}

func (*SVGFESpotLightElement) GetPointsAtY() (pointsAtY *SVGAnimatedNumber) {
	js.Rewrite("$<.pointsAtY")
	return pointsAtY
}

func (*SVGFESpotLightElement) GetPointsAtZ() (pointsAtZ *SVGAnimatedNumber) {
	js.Rewrite("$<.pointsAtZ")
	return pointsAtZ
}

func (*SVGFESpotLightElement) GetSpecularExponent() (specularExponent *SVGAnimatedNumber) {
	js.Rewrite("$<.specularExponent")
	return specularExponent
}

func (*SVGFESpotLightElement) GetX() (x *SVGAnimatedNumber) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGFESpotLightElement) GetY() (y *SVGAnimatedNumber) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGFESpotLightElement) GetZ() (z *SVGAnimatedNumber) {
	js.Rewrite("$<.z")
	return z
}

type SVGFETileElement struct {
	*SVGElement
	*SVGFilterPrimitiveStandardAttributes
}

func (*SVGFETileElement) GetIn1() (in1 *SVGAnimatedString) {
	js.Rewrite("$<.in1")
	return in1
}

type SVGFETurbulenceElement struct {
	*SVGElement
	*SVGFilterPrimitiveStandardAttributes
}

func (*SVGFETurbulenceElement) GetBaseFrequencyX() (baseFrequencyX *SVGAnimatedNumber) {
	js.Rewrite("$<.baseFrequencyX")
	return baseFrequencyX
}

func (*SVGFETurbulenceElement) GetBaseFrequencyY() (baseFrequencyY *SVGAnimatedNumber) {
	js.Rewrite("$<.baseFrequencyY")
	return baseFrequencyY
}

func (*SVGFETurbulenceElement) GetNumOctaves() (numOctaves *SVGAnimatedInteger) {
	js.Rewrite("$<.numOctaves")
	return numOctaves
}

func (*SVGFETurbulenceElement) GetSeed() (seed *SVGAnimatedNumber) {
	js.Rewrite("$<.seed")
	return seed
}

func (*SVGFETurbulenceElement) GetStitchTiles() (stitchTiles *SVGAnimatedEnumeration) {
	js.Rewrite("$<.stitchTiles")
	return stitchTiles
}

func (*SVGFETurbulenceElement) GetType() (kind *SVGAnimatedEnumeration) {
	js.Rewrite("$<.type")
	return kind
}

type SVGFilterElement struct {
	*SVGElement
	*SVGUnitTypes
	*SVGURIReference
}

func (*SVGFilterElement) SetFilterRes(filterResX uint, filterResY uint) {
	js.Rewrite("$<.setFilterRes($1, $2)", filterResX, filterResY)
}

func (*SVGFilterElement) GetFilterResX() (filterResX *SVGAnimatedInteger) {
	js.Rewrite("$<.filterResX")
	return filterResX
}

func (*SVGFilterElement) GetFilterResY() (filterResY *SVGAnimatedInteger) {
	js.Rewrite("$<.filterResY")
	return filterResY
}

func (*SVGFilterElement) GetFilterUnits() (filterUnits *SVGAnimatedEnumeration) {
	js.Rewrite("$<.filterUnits")
	return filterUnits
}

func (*SVGFilterElement) GetHeight() (height *SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

func (*SVGFilterElement) GetPrimitiveUnits() (primitiveUnits *SVGAnimatedEnumeration) {
	js.Rewrite("$<.primitiveUnits")
	return primitiveUnits
}

func (*SVGFilterElement) GetWidth() (width *SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

func (*SVGFilterElement) GetX() (x *SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGFilterElement) GetY() (y *SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}

type SVGForeignObjectElement struct {
	*SVGGraphicsElement
}

func (*SVGForeignObjectElement) GetHeight() (height *SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

func (*SVGForeignObjectElement) GetWidth() (width *SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

func (*SVGForeignObjectElement) GetX() (x *SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGForeignObjectElement) GetY() (y *SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}

type SVGGElement struct {
	*SVGGraphicsElement
}

type SVGGradientElement struct {
	*SVGElement
	*SVGUnitTypes
	*SVGURIReference
}

func (*SVGGradientElement) GetGradientTransform() (gradientTransform *SVGAnimatedTransformList) {
	js.Rewrite("$<.gradientTransform")
	return gradientTransform
}

func (*SVGGradientElement) GetGradientUnits() (gradientUnits *SVGAnimatedEnumeration) {
	js.Rewrite("$<.gradientUnits")
	return gradientUnits
}

func (*SVGGradientElement) GetSpreadMethod() (spreadMethod *SVGAnimatedEnumeration) {
	js.Rewrite("$<.spreadMethod")
	return spreadMethod
}

type SVGGraphicsElement struct {
	*SVGElement
	*SVGTests
}

func (*SVGGraphicsElement) GetBBox() (s *SVGRect) {
	js.Rewrite("$<.getBBox()")
	return s
}

func (*SVGGraphicsElement) GetCTM() (s *SVGMatrix) {
	js.Rewrite("$<.getCTM()")
	return s
}

func (*SVGGraphicsElement) GetScreenCTM() (s *SVGMatrix) {
	js.Rewrite("$<.getScreenCTM()")
	return s
}

func (*SVGGraphicsElement) GetTransformToElement(element *SVGElement) (s *SVGMatrix) {
	js.Rewrite("$<.getTransformToElement($1)", element)
	return s
}

func (*SVGGraphicsElement) GetFarthestViewportElement() (farthestViewportElement *SVGElement) {
	js.Rewrite("$<.farthestViewportElement")
	return farthestViewportElement
}

func (*SVGGraphicsElement) GetNearestViewportElement() (nearestViewportElement *SVGElement) {
	js.Rewrite("$<.nearestViewportElement")
	return nearestViewportElement
}

func (*SVGGraphicsElement) GetTransform() (transform *SVGAnimatedTransformList) {
	js.Rewrite("$<.transform")
	return transform
}

type SVGImageElement struct {
	*SVGGraphicsElement
	*SVGURIReference
}

func (*SVGImageElement) GetHeight() (height *SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

func (*SVGImageElement) GetPreserveAspectRatio() (preserveAspectRatio *SVGAnimatedPreserveAspectRatio) {
	js.Rewrite("$<.preserveAspectRatio")
	return preserveAspectRatio
}

func (*SVGImageElement) GetWidth() (width *SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

func (*SVGImageElement) GetX() (x *SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGImageElement) GetY() (y *SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}

type SVGLength struct {
}

func (*SVGLength) ConvertToSpecifiedUnits(unitType uint8) {
	js.Rewrite("$<.convertToSpecifiedUnits($1)", unitType)
}

func (*SVGLength) NewValueSpecifiedUnits(unitType uint8, valueInSpecifiedUnits float32) {
	js.Rewrite("$<.newValueSpecifiedUnits($1, $2)", unitType, valueInSpecifiedUnits)
}

func (*SVGLength) GetUnitType() (unitType uint8) {
	js.Rewrite("$<.unitType")
	return unitType
}

func (*SVGLength) GetValue() (value float32) {
	js.Rewrite("$<.value")
	return value
}

func (*SVGLength) SetValue(value float32) {
	js.Rewrite("$<.value = $1", value)
}

func (*SVGLength) GetValueAsString() (valueAsString string) {
	js.Rewrite("$<.valueAsString")
	return valueAsString
}

func (*SVGLength) SetValueAsString(valueAsString string) {
	js.Rewrite("$<.valueAsString = $1", valueAsString)
}

func (*SVGLength) GetValueInSpecifiedUnits() (valueInSpecifiedUnits float32) {
	js.Rewrite("$<.valueInSpecifiedUnits")
	return valueInSpecifiedUnits
}

func (*SVGLength) SetValueInSpecifiedUnits(valueInSpecifiedUnits float32) {
	js.Rewrite("$<.valueInSpecifiedUnits = $1", valueInSpecifiedUnits)
}

type SVGLengthList struct {
}

func (*SVGLengthList) AppendItem(newItem *SVGLength) (s *SVGLength) {
	js.Rewrite("$<.appendItem($1)", newItem)
	return s
}

func (*SVGLengthList) Clear() {
	js.Rewrite("$<.clear()")
}

func (*SVGLengthList) GetItem(index uint) (s *SVGLength) {
	js.Rewrite("$<.getItem($1)", index)
	return s
}

func (*SVGLengthList) Initialize(newItem *SVGLength) (s *SVGLength) {
	js.Rewrite("$<.initialize($1)", newItem)
	return s
}

func (*SVGLengthList) InsertItemBefore(newItem *SVGLength, index uint) (s *SVGLength) {
	js.Rewrite("$<.insertItemBefore($1, $2)", newItem, index)
	return s
}

func (*SVGLengthList) RemoveItem(index uint) (s *SVGLength) {
	js.Rewrite("$<.removeItem($1)", index)
	return s
}

func (*SVGLengthList) ReplaceItem(newItem *SVGLength, index uint) (s *SVGLength) {
	js.Rewrite("$<.replaceItem($1, $2)", newItem, index)
	return s
}

func (*SVGLengthList) GetNumberOfItems() (numberOfItems uint) {
	js.Rewrite("$<.numberOfItems")
	return numberOfItems
}

type SVGLinearGradientElement struct {
	*SVGGradientElement
}

func (*SVGLinearGradientElement) GetX1() (x1 *SVGAnimatedLength) {
	js.Rewrite("$<.x1")
	return x1
}

func (*SVGLinearGradientElement) GetX2() (x2 *SVGAnimatedLength) {
	js.Rewrite("$<.x2")
	return x2
}

func (*SVGLinearGradientElement) GetY1() (y1 *SVGAnimatedLength) {
	js.Rewrite("$<.y1")
	return y1
}

func (*SVGLinearGradientElement) GetY2() (y2 *SVGAnimatedLength) {
	js.Rewrite("$<.y2")
	return y2
}

type SVGLineElement struct {
	*SVGGraphicsElement
}

func (*SVGLineElement) GetX1() (x1 *SVGAnimatedLength) {
	js.Rewrite("$<.x1")
	return x1
}

func (*SVGLineElement) GetX2() (x2 *SVGAnimatedLength) {
	js.Rewrite("$<.x2")
	return x2
}

func (*SVGLineElement) GetY1() (y1 *SVGAnimatedLength) {
	js.Rewrite("$<.y1")
	return y1
}

func (*SVGLineElement) GetY2() (y2 *SVGAnimatedLength) {
	js.Rewrite("$<.y2")
	return y2
}

type SVGMarkerElement struct {
	*SVGElement
	*SVGFitToViewBox
}

func (*SVGMarkerElement) SetOrientToAngle(angle *SVGAngle) {
	js.Rewrite("$<.setOrientToAngle($1)", angle)
}

func (*SVGMarkerElement) SetOrientToAuto() {
	js.Rewrite("$<.setOrientToAuto()")
}

func (*SVGMarkerElement) GetMarkerHeight() (markerHeight *SVGAnimatedLength) {
	js.Rewrite("$<.markerHeight")
	return markerHeight
}

func (*SVGMarkerElement) GetMarkerUnits() (markerUnits *SVGAnimatedEnumeration) {
	js.Rewrite("$<.markerUnits")
	return markerUnits
}

func (*SVGMarkerElement) GetMarkerWidth() (markerWidth *SVGAnimatedLength) {
	js.Rewrite("$<.markerWidth")
	return markerWidth
}

func (*SVGMarkerElement) GetOrientAngle() (orientAngle *SVGAnimatedAngle) {
	js.Rewrite("$<.orientAngle")
	return orientAngle
}

func (*SVGMarkerElement) GetOrientType() (orientType *SVGAnimatedEnumeration) {
	js.Rewrite("$<.orientType")
	return orientType
}

func (*SVGMarkerElement) GetRefX() (refX *SVGAnimatedLength) {
	js.Rewrite("$<.refX")
	return refX
}

func (*SVGMarkerElement) GetRefY() (refY *SVGAnimatedLength) {
	js.Rewrite("$<.refY")
	return refY
}

type SVGMaskElement struct {
	*SVGElement
	*SVGTests
	*SVGUnitTypes
}

func (*SVGMaskElement) GetHeight() (height *SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

func (*SVGMaskElement) GetMaskContentUnits() (maskContentUnits *SVGAnimatedEnumeration) {
	js.Rewrite("$<.maskContentUnits")
	return maskContentUnits
}

func (*SVGMaskElement) GetMaskUnits() (maskUnits *SVGAnimatedEnumeration) {
	js.Rewrite("$<.maskUnits")
	return maskUnits
}

func (*SVGMaskElement) GetWidth() (width *SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

func (*SVGMaskElement) GetX() (x *SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGMaskElement) GetY() (y *SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}

type SVGMatrix struct {
}

func (*SVGMatrix) FlipX() (s *SVGMatrix) {
	js.Rewrite("$<.flipX()")
	return s
}

func (*SVGMatrix) FlipY() (s *SVGMatrix) {
	js.Rewrite("$<.flipY()")
	return s
}

func (*SVGMatrix) Inverse() (s *SVGMatrix) {
	js.Rewrite("$<.inverse()")
	return s
}

func (*SVGMatrix) Multiply(secondMatrix *SVGMatrix) (s *SVGMatrix) {
	js.Rewrite("$<.multiply($1)", secondMatrix)
	return s
}

func (*SVGMatrix) Rotate(angle float32) (s *SVGMatrix) {
	js.Rewrite("$<.rotate($1)", angle)
	return s
}

func (*SVGMatrix) RotateFromVector(x float32, y float32) (s *SVGMatrix) {
	js.Rewrite("$<.rotateFromVector($1, $2)", x, y)
	return s
}

func (*SVGMatrix) Scale(scaleFactor float32) (s *SVGMatrix) {
	js.Rewrite("$<.scale($1)", scaleFactor)
	return s
}

func (*SVGMatrix) ScaleNonUniform(scaleFactorX float32, scaleFactorY float32) (s *SVGMatrix) {
	js.Rewrite("$<.scaleNonUniform($1, $2)", scaleFactorX, scaleFactorY)
	return s
}

func (*SVGMatrix) SkewX(angle float32) (s *SVGMatrix) {
	js.Rewrite("$<.skewX($1)", angle)
	return s
}

func (*SVGMatrix) SkewY(angle float32) (s *SVGMatrix) {
	js.Rewrite("$<.skewY($1)", angle)
	return s
}

func (*SVGMatrix) Translate(x float32, y float32) (s *SVGMatrix) {
	js.Rewrite("$<.translate($1, $2)", x, y)
	return s
}

func (*SVGMatrix) GetA() (a float32) {
	js.Rewrite("$<.a")
	return a
}

func (*SVGMatrix) SetA(a float32) {
	js.Rewrite("$<.a = $1", a)
}

func (*SVGMatrix) GetB() (b float32) {
	js.Rewrite("$<.b")
	return b
}

func (*SVGMatrix) SetB(b float32) {
	js.Rewrite("$<.b = $1", b)
}

func (*SVGMatrix) GetC() (c float32) {
	js.Rewrite("$<.c")
	return c
}

func (*SVGMatrix) SetC(c float32) {
	js.Rewrite("$<.c = $1", c)
}

func (*SVGMatrix) GetD() (d float32) {
	js.Rewrite("$<.d")
	return d
}

func (*SVGMatrix) SetD(d float32) {
	js.Rewrite("$<.d = $1", d)
}

func (*SVGMatrix) GetE() (e float32) {
	js.Rewrite("$<.e")
	return e
}

func (*SVGMatrix) SetE(e float32) {
	js.Rewrite("$<.e = $1", e)
}

func (*SVGMatrix) GetF() (f float32) {
	js.Rewrite("$<.f")
	return f
}

func (*SVGMatrix) SetF(f float32) {
	js.Rewrite("$<.f = $1", f)
}

type SVGMetadataElement struct {
	*SVGElement
}

type SVGNumber struct {
}

func (*SVGNumber) GetValue() (value float32) {
	js.Rewrite("$<.value")
	return value
}

func (*SVGNumber) SetValue(value float32) {
	js.Rewrite("$<.value = $1", value)
}

type SVGNumberList struct {
}

func (*SVGNumberList) AppendItem(newItem *SVGNumber) (s *SVGNumber) {
	js.Rewrite("$<.appendItem($1)", newItem)
	return s
}

func (*SVGNumberList) Clear() {
	js.Rewrite("$<.clear()")
}

func (*SVGNumberList) GetItem(index uint) (s *SVGNumber) {
	js.Rewrite("$<.getItem($1)", index)
	return s
}

func (*SVGNumberList) Initialize(newItem *SVGNumber) (s *SVGNumber) {
	js.Rewrite("$<.initialize($1)", newItem)
	return s
}

func (*SVGNumberList) InsertItemBefore(newItem *SVGNumber, index uint) (s *SVGNumber) {
	js.Rewrite("$<.insertItemBefore($1, $2)", newItem, index)
	return s
}

func (*SVGNumberList) RemoveItem(index uint) (s *SVGNumber) {
	js.Rewrite("$<.removeItem($1)", index)
	return s
}

func (*SVGNumberList) ReplaceItem(newItem *SVGNumber, index uint) (s *SVGNumber) {
	js.Rewrite("$<.replaceItem($1, $2)", newItem, index)
	return s
}

func (*SVGNumberList) GetNumberOfItems() (numberOfItems uint) {
	js.Rewrite("$<.numberOfItems")
	return numberOfItems
}

type SVGPathElement struct {
	*SVGGraphicsElement
}

func (*SVGPathElement) CreateSVGPathSegArcAbs(x float32, y float32, r1 float32, r2 float32, angle float32, largeArcFlag bool, sweepFlag bool) (s *SVGPathSegArcAbs) {
	js.Rewrite("$<.createSVGPathSegArcAbs($1, $2, $3, $4, $5, $6, $7)", x, y, r1, r2, angle, largeArcFlag, sweepFlag)
	return s
}

func (*SVGPathElement) CreateSVGPathSegArcRel(x float32, y float32, r1 float32, r2 float32, angle float32, largeArcFlag bool, sweepFlag bool) (s *SVGPathSegArcRel) {
	js.Rewrite("$<.createSVGPathSegArcRel($1, $2, $3, $4, $5, $6, $7)", x, y, r1, r2, angle, largeArcFlag, sweepFlag)
	return s
}

func (*SVGPathElement) CreateSVGPathSegClosePath() (s *SVGPathSegClosePath) {
	js.Rewrite("$<.createSVGPathSegClosePath()")
	return s
}

func (*SVGPathElement) CreateSVGPathSegCurvetoCubicAbs(x float32, y float32, x1 float32, y1 float32, x2 float32, y2 float32) (s *SVGPathSegCurvetoCubicAbs) {
	js.Rewrite("$<.createSVGPathSegCurvetoCubicAbs($1, $2, $3, $4, $5, $6)", x, y, x1, y1, x2, y2)
	return s
}

func (*SVGPathElement) CreateSVGPathSegCurvetoCubicRel(x float32, y float32, x1 float32, y1 float32, x2 float32, y2 float32) (s *SVGPathSegCurvetoCubicRel) {
	js.Rewrite("$<.createSVGPathSegCurvetoCubicRel($1, $2, $3, $4, $5, $6)", x, y, x1, y1, x2, y2)
	return s
}

func (*SVGPathElement) CreateSVGPathSegCurvetoCubicSmoothAbs(x float32, y float32, x2 float32, y2 float32) (s *SVGPathSegCurvetoCubicSmoothAbs) {
	js.Rewrite("$<.createSVGPathSegCurvetoCubicSmoothAbs($1, $2, $3, $4)", x, y, x2, y2)
	return s
}

func (*SVGPathElement) CreateSVGPathSegCurvetoCubicSmoothRel(x float32, y float32, x2 float32, y2 float32) (s *SVGPathSegCurvetoCubicSmoothRel) {
	js.Rewrite("$<.createSVGPathSegCurvetoCubicSmoothRel($1, $2, $3, $4)", x, y, x2, y2)
	return s
}

func (*SVGPathElement) CreateSVGPathSegCurvetoQuadraticAbs(x float32, y float32, x1 float32, y1 float32) (s *SVGPathSegCurvetoQuadraticAbs) {
	js.Rewrite("$<.createSVGPathSegCurvetoQuadraticAbs($1, $2, $3, $4)", x, y, x1, y1)
	return s
}

func (*SVGPathElement) CreateSVGPathSegCurvetoQuadraticRel(x float32, y float32, x1 float32, y1 float32) (s *SVGPathSegCurvetoQuadraticRel) {
	js.Rewrite("$<.createSVGPathSegCurvetoQuadraticRel($1, $2, $3, $4)", x, y, x1, y1)
	return s
}

func (*SVGPathElement) CreateSVGPathSegCurvetoQuadraticSmoothAbs(x float32, y float32) (s *SVGPathSegCurvetoQuadraticSmoothAbs) {
	js.Rewrite("$<.createSVGPathSegCurvetoQuadraticSmoothAbs($1, $2)", x, y)
	return s
}

func (*SVGPathElement) CreateSVGPathSegCurvetoQuadraticSmoothRel(x float32, y float32) (s *SVGPathSegCurvetoQuadraticSmoothRel) {
	js.Rewrite("$<.createSVGPathSegCurvetoQuadraticSmoothRel($1, $2)", x, y)
	return s
}

func (*SVGPathElement) CreateSVGPathSegLinetoAbs(x float32, y float32) (s *SVGPathSegLinetoAbs) {
	js.Rewrite("$<.createSVGPathSegLinetoAbs($1, $2)", x, y)
	return s
}

func (*SVGPathElement) CreateSVGPathSegLinetoHorizontalAbs(x float32) (s *SVGPathSegLinetoHorizontalAbs) {
	js.Rewrite("$<.createSVGPathSegLinetoHorizontalAbs($1)", x)
	return s
}

func (*SVGPathElement) CreateSVGPathSegLinetoHorizontalRel(x float32) (s *SVGPathSegLinetoHorizontalRel) {
	js.Rewrite("$<.createSVGPathSegLinetoHorizontalRel($1)", x)
	return s
}

func (*SVGPathElement) CreateSVGPathSegLinetoRel(x float32, y float32) (s *SVGPathSegLinetoRel) {
	js.Rewrite("$<.createSVGPathSegLinetoRel($1, $2)", x, y)
	return s
}

func (*SVGPathElement) CreateSVGPathSegLinetoVerticalAbs(y float32) (s *SVGPathSegLinetoVerticalAbs) {
	js.Rewrite("$<.createSVGPathSegLinetoVerticalAbs($1)", y)
	return s
}

func (*SVGPathElement) CreateSVGPathSegLinetoVerticalRel(y float32) (s *SVGPathSegLinetoVerticalRel) {
	js.Rewrite("$<.createSVGPathSegLinetoVerticalRel($1)", y)
	return s
}

func (*SVGPathElement) CreateSVGPathSegMovetoAbs(x float32, y float32) (s *SVGPathSegMovetoAbs) {
	js.Rewrite("$<.createSVGPathSegMovetoAbs($1, $2)", x, y)
	return s
}

func (*SVGPathElement) CreateSVGPathSegMovetoRel(x float32, y float32) (s *SVGPathSegMovetoRel) {
	js.Rewrite("$<.createSVGPathSegMovetoRel($1, $2)", x, y)
	return s
}

func (*SVGPathElement) GetPathSegAtLength(distance float32) (u uint) {
	js.Rewrite("$<.getPathSegAtLength($1)", distance)
	return u
}

func (*SVGPathElement) GetPointAtLength(distance float32) (s *SVGPoint) {
	js.Rewrite("$<.getPointAtLength($1)", distance)
	return s
}

func (*SVGPathElement) GetTotalLength() (f float32) {
	js.Rewrite("$<.getTotalLength()")
	return f
}

func (*SVGPathElement) GetPathSegList() (pathSegList *SVGPathSegList) {
	js.Rewrite("$<.pathSegList")
	return pathSegList
}

type SVGPathSeg struct {
}

func (*SVGPathSeg) GetPathSegType() (pathSegType uint8) {
	js.Rewrite("$<.pathSegType")
	return pathSegType
}

func (*SVGPathSeg) GetPathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	js.Rewrite("$<.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}

type SVGPathSegArcAbs struct {
	*SVGPathSeg
}

func (*SVGPathSegArcAbs) GetAngle() (angle float32) {
	js.Rewrite("$<.angle")
	return angle
}

func (*SVGPathSegArcAbs) SetAngle(angle float32) {
	js.Rewrite("$<.angle = $1", angle)
}

func (*SVGPathSegArcAbs) GetLargeArcFlag() (largeArcFlag bool) {
	js.Rewrite("$<.largeArcFlag")
	return largeArcFlag
}

func (*SVGPathSegArcAbs) SetLargeArcFlag(largeArcFlag bool) {
	js.Rewrite("$<.largeArcFlag = $1", largeArcFlag)
}

func (*SVGPathSegArcAbs) GetR1() (r1 float32) {
	js.Rewrite("$<.r1")
	return r1
}

func (*SVGPathSegArcAbs) SetR1(r1 float32) {
	js.Rewrite("$<.r1 = $1", r1)
}

func (*SVGPathSegArcAbs) GetR2() (r2 float32) {
	js.Rewrite("$<.r2")
	return r2
}

func (*SVGPathSegArcAbs) SetR2(r2 float32) {
	js.Rewrite("$<.r2 = $1", r2)
}

func (*SVGPathSegArcAbs) GetSweepFlag() (sweepFlag bool) {
	js.Rewrite("$<.sweepFlag")
	return sweepFlag
}

func (*SVGPathSegArcAbs) SetSweepFlag(sweepFlag bool) {
	js.Rewrite("$<.sweepFlag = $1", sweepFlag)
}

func (*SVGPathSegArcAbs) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegArcAbs) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGPathSegArcAbs) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegArcAbs) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}

type SVGPathSegArcRel struct {
	*SVGPathSeg
}

func (*SVGPathSegArcRel) GetAngle() (angle float32) {
	js.Rewrite("$<.angle")
	return angle
}

func (*SVGPathSegArcRel) SetAngle(angle float32) {
	js.Rewrite("$<.angle = $1", angle)
}

func (*SVGPathSegArcRel) GetLargeArcFlag() (largeArcFlag bool) {
	js.Rewrite("$<.largeArcFlag")
	return largeArcFlag
}

func (*SVGPathSegArcRel) SetLargeArcFlag(largeArcFlag bool) {
	js.Rewrite("$<.largeArcFlag = $1", largeArcFlag)
}

func (*SVGPathSegArcRel) GetR1() (r1 float32) {
	js.Rewrite("$<.r1")
	return r1
}

func (*SVGPathSegArcRel) SetR1(r1 float32) {
	js.Rewrite("$<.r1 = $1", r1)
}

func (*SVGPathSegArcRel) GetR2() (r2 float32) {
	js.Rewrite("$<.r2")
	return r2
}

func (*SVGPathSegArcRel) SetR2(r2 float32) {
	js.Rewrite("$<.r2 = $1", r2)
}

func (*SVGPathSegArcRel) GetSweepFlag() (sweepFlag bool) {
	js.Rewrite("$<.sweepFlag")
	return sweepFlag
}

func (*SVGPathSegArcRel) SetSweepFlag(sweepFlag bool) {
	js.Rewrite("$<.sweepFlag = $1", sweepFlag)
}

func (*SVGPathSegArcRel) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegArcRel) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGPathSegArcRel) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegArcRel) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}

type SVGPathSegClosePath struct {
	*SVGPathSeg
}

type SVGPathSegCurvetoCubicAbs struct {
	*SVGPathSeg
}

func (*SVGPathSegCurvetoCubicAbs) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegCurvetoCubicAbs) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGPathSegCurvetoCubicAbs) GetX1() (x1 float32) {
	js.Rewrite("$<.x1")
	return x1
}

func (*SVGPathSegCurvetoCubicAbs) SetX1(x1 float32) {
	js.Rewrite("$<.x1 = $1", x1)
}

func (*SVGPathSegCurvetoCubicAbs) GetX2() (x2 float32) {
	js.Rewrite("$<.x2")
	return x2
}

func (*SVGPathSegCurvetoCubicAbs) SetX2(x2 float32) {
	js.Rewrite("$<.x2 = $1", x2)
}

func (*SVGPathSegCurvetoCubicAbs) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegCurvetoCubicAbs) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}

func (*SVGPathSegCurvetoCubicAbs) GetY1() (y1 float32) {
	js.Rewrite("$<.y1")
	return y1
}

func (*SVGPathSegCurvetoCubicAbs) SetY1(y1 float32) {
	js.Rewrite("$<.y1 = $1", y1)
}

func (*SVGPathSegCurvetoCubicAbs) GetY2() (y2 float32) {
	js.Rewrite("$<.y2")
	return y2
}

func (*SVGPathSegCurvetoCubicAbs) SetY2(y2 float32) {
	js.Rewrite("$<.y2 = $1", y2)
}

type SVGPathSegCurvetoCubicRel struct {
	*SVGPathSeg
}

func (*SVGPathSegCurvetoCubicRel) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegCurvetoCubicRel) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGPathSegCurvetoCubicRel) GetX1() (x1 float32) {
	js.Rewrite("$<.x1")
	return x1
}

func (*SVGPathSegCurvetoCubicRel) SetX1(x1 float32) {
	js.Rewrite("$<.x1 = $1", x1)
}

func (*SVGPathSegCurvetoCubicRel) GetX2() (x2 float32) {
	js.Rewrite("$<.x2")
	return x2
}

func (*SVGPathSegCurvetoCubicRel) SetX2(x2 float32) {
	js.Rewrite("$<.x2 = $1", x2)
}

func (*SVGPathSegCurvetoCubicRel) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegCurvetoCubicRel) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}

func (*SVGPathSegCurvetoCubicRel) GetY1() (y1 float32) {
	js.Rewrite("$<.y1")
	return y1
}

func (*SVGPathSegCurvetoCubicRel) SetY1(y1 float32) {
	js.Rewrite("$<.y1 = $1", y1)
}

func (*SVGPathSegCurvetoCubicRel) GetY2() (y2 float32) {
	js.Rewrite("$<.y2")
	return y2
}

func (*SVGPathSegCurvetoCubicRel) SetY2(y2 float32) {
	js.Rewrite("$<.y2 = $1", y2)
}

type SVGPathSegCurvetoCubicSmoothAbs struct {
	*SVGPathSeg
}

func (*SVGPathSegCurvetoCubicSmoothAbs) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegCurvetoCubicSmoothAbs) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGPathSegCurvetoCubicSmoothAbs) GetX2() (x2 float32) {
	js.Rewrite("$<.x2")
	return x2
}

func (*SVGPathSegCurvetoCubicSmoothAbs) SetX2(x2 float32) {
	js.Rewrite("$<.x2 = $1", x2)
}

func (*SVGPathSegCurvetoCubicSmoothAbs) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegCurvetoCubicSmoothAbs) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}

func (*SVGPathSegCurvetoCubicSmoothAbs) GetY2() (y2 float32) {
	js.Rewrite("$<.y2")
	return y2
}

func (*SVGPathSegCurvetoCubicSmoothAbs) SetY2(y2 float32) {
	js.Rewrite("$<.y2 = $1", y2)
}

type SVGPathSegCurvetoCubicSmoothRel struct {
	*SVGPathSeg
}

func (*SVGPathSegCurvetoCubicSmoothRel) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegCurvetoCubicSmoothRel) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGPathSegCurvetoCubicSmoothRel) GetX2() (x2 float32) {
	js.Rewrite("$<.x2")
	return x2
}

func (*SVGPathSegCurvetoCubicSmoothRel) SetX2(x2 float32) {
	js.Rewrite("$<.x2 = $1", x2)
}

func (*SVGPathSegCurvetoCubicSmoothRel) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegCurvetoCubicSmoothRel) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}

func (*SVGPathSegCurvetoCubicSmoothRel) GetY2() (y2 float32) {
	js.Rewrite("$<.y2")
	return y2
}

func (*SVGPathSegCurvetoCubicSmoothRel) SetY2(y2 float32) {
	js.Rewrite("$<.y2 = $1", y2)
}

type SVGPathSegCurvetoQuadraticAbs struct {
	*SVGPathSeg
}

func (*SVGPathSegCurvetoQuadraticAbs) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegCurvetoQuadraticAbs) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGPathSegCurvetoQuadraticAbs) GetX1() (x1 float32) {
	js.Rewrite("$<.x1")
	return x1
}

func (*SVGPathSegCurvetoQuadraticAbs) SetX1(x1 float32) {
	js.Rewrite("$<.x1 = $1", x1)
}

func (*SVGPathSegCurvetoQuadraticAbs) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegCurvetoQuadraticAbs) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}

func (*SVGPathSegCurvetoQuadraticAbs) GetY1() (y1 float32) {
	js.Rewrite("$<.y1")
	return y1
}

func (*SVGPathSegCurvetoQuadraticAbs) SetY1(y1 float32) {
	js.Rewrite("$<.y1 = $1", y1)
}

type SVGPathSegCurvetoQuadraticRel struct {
	*SVGPathSeg
}

func (*SVGPathSegCurvetoQuadraticRel) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegCurvetoQuadraticRel) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGPathSegCurvetoQuadraticRel) GetX1() (x1 float32) {
	js.Rewrite("$<.x1")
	return x1
}

func (*SVGPathSegCurvetoQuadraticRel) SetX1(x1 float32) {
	js.Rewrite("$<.x1 = $1", x1)
}

func (*SVGPathSegCurvetoQuadraticRel) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegCurvetoQuadraticRel) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}

func (*SVGPathSegCurvetoQuadraticRel) GetY1() (y1 float32) {
	js.Rewrite("$<.y1")
	return y1
}

func (*SVGPathSegCurvetoQuadraticRel) SetY1(y1 float32) {
	js.Rewrite("$<.y1 = $1", y1)
}

type SVGPathSegCurvetoQuadraticSmoothAbs struct {
	*SVGPathSeg
}

func (*SVGPathSegCurvetoQuadraticSmoothAbs) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegCurvetoQuadraticSmoothAbs) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGPathSegCurvetoQuadraticSmoothAbs) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegCurvetoQuadraticSmoothAbs) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}

type SVGPathSegCurvetoQuadraticSmoothRel struct {
	*SVGPathSeg
}

func (*SVGPathSegCurvetoQuadraticSmoothRel) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegCurvetoQuadraticSmoothRel) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGPathSegCurvetoQuadraticSmoothRel) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegCurvetoQuadraticSmoothRel) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}

type SVGPathSegLinetoAbs struct {
	*SVGPathSeg
}

func (*SVGPathSegLinetoAbs) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegLinetoAbs) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGPathSegLinetoAbs) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegLinetoAbs) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}

type SVGPathSegLinetoHorizontalAbs struct {
	*SVGPathSeg
}

func (*SVGPathSegLinetoHorizontalAbs) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegLinetoHorizontalAbs) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

type SVGPathSegLinetoHorizontalRel struct {
	*SVGPathSeg
}

func (*SVGPathSegLinetoHorizontalRel) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegLinetoHorizontalRel) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

type SVGPathSegLinetoRel struct {
	*SVGPathSeg
}

func (*SVGPathSegLinetoRel) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegLinetoRel) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGPathSegLinetoRel) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegLinetoRel) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}

type SVGPathSegLinetoVerticalAbs struct {
	*SVGPathSeg
}

func (*SVGPathSegLinetoVerticalAbs) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegLinetoVerticalAbs) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}

type SVGPathSegLinetoVerticalRel struct {
	*SVGPathSeg
}

func (*SVGPathSegLinetoVerticalRel) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegLinetoVerticalRel) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}

type SVGPathSegList struct {
}

func (*SVGPathSegList) AppendItem(newItem *SVGPathSeg) (s *SVGPathSeg) {
	js.Rewrite("$<.appendItem($1)", newItem)
	return s
}

func (*SVGPathSegList) Clear() {
	js.Rewrite("$<.clear()")
}

func (*SVGPathSegList) GetItem(index uint) (s *SVGPathSeg) {
	js.Rewrite("$<.getItem($1)", index)
	return s
}

func (*SVGPathSegList) Initialize(newItem *SVGPathSeg) (s *SVGPathSeg) {
	js.Rewrite("$<.initialize($1)", newItem)
	return s
}

func (*SVGPathSegList) InsertItemBefore(newItem *SVGPathSeg, index uint) (s *SVGPathSeg) {
	js.Rewrite("$<.insertItemBefore($1, $2)", newItem, index)
	return s
}

func (*SVGPathSegList) RemoveItem(index uint) (s *SVGPathSeg) {
	js.Rewrite("$<.removeItem($1)", index)
	return s
}

func (*SVGPathSegList) ReplaceItem(newItem *SVGPathSeg, index uint) (s *SVGPathSeg) {
	js.Rewrite("$<.replaceItem($1, $2)", newItem, index)
	return s
}

func (*SVGPathSegList) GetNumberOfItems() (numberOfItems uint) {
	js.Rewrite("$<.numberOfItems")
	return numberOfItems
}

type SVGPathSegMovetoAbs struct {
	*SVGPathSeg
}

func (*SVGPathSegMovetoAbs) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegMovetoAbs) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGPathSegMovetoAbs) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegMovetoAbs) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}

type SVGPathSegMovetoRel struct {
	*SVGPathSeg
}

func (*SVGPathSegMovetoRel) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPathSegMovetoRel) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGPathSegMovetoRel) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPathSegMovetoRel) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}

type SVGPatternElement struct {
	*SVGElement
	*SVGTests
	*SVGUnitTypes
	*SVGFitToViewBox
	*SVGURIReference
}

func (*SVGPatternElement) GetHeight() (height *SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

func (*SVGPatternElement) GetPatternContentUnits() (patternContentUnits *SVGAnimatedEnumeration) {
	js.Rewrite("$<.patternContentUnits")
	return patternContentUnits
}

func (*SVGPatternElement) GetPatternTransform() (patternTransform *SVGAnimatedTransformList) {
	js.Rewrite("$<.patternTransform")
	return patternTransform
}

func (*SVGPatternElement) GetPatternUnits() (patternUnits *SVGAnimatedEnumeration) {
	js.Rewrite("$<.patternUnits")
	return patternUnits
}

func (*SVGPatternElement) GetWidth() (width *SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

func (*SVGPatternElement) GetX() (x *SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPatternElement) GetY() (y *SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}

type SVGPoint struct {
}

func (*SVGPoint) MatrixTransform(matrix *SVGMatrix) (s *SVGPoint) {
	js.Rewrite("$<.matrixTransform($1)", matrix)
	return s
}

func (*SVGPoint) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGPoint) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGPoint) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGPoint) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}

type SVGPointList struct {
}

func (*SVGPointList) AppendItem(newItem *SVGPoint) (s *SVGPoint) {
	js.Rewrite("$<.appendItem($1)", newItem)
	return s
}

func (*SVGPointList) Clear() {
	js.Rewrite("$<.clear()")
}

func (*SVGPointList) GetItem(index uint) (s *SVGPoint) {
	js.Rewrite("$<.getItem($1)", index)
	return s
}

func (*SVGPointList) Initialize(newItem *SVGPoint) (s *SVGPoint) {
	js.Rewrite("$<.initialize($1)", newItem)
	return s
}

func (*SVGPointList) InsertItemBefore(newItem *SVGPoint, index uint) (s *SVGPoint) {
	js.Rewrite("$<.insertItemBefore($1, $2)", newItem, index)
	return s
}

func (*SVGPointList) RemoveItem(index uint) (s *SVGPoint) {
	js.Rewrite("$<.removeItem($1)", index)
	return s
}

func (*SVGPointList) ReplaceItem(newItem *SVGPoint, index uint) (s *SVGPoint) {
	js.Rewrite("$<.replaceItem($1, $2)", newItem, index)
	return s
}

func (*SVGPointList) GetNumberOfItems() (numberOfItems uint) {
	js.Rewrite("$<.numberOfItems")
	return numberOfItems
}

type SVGPolygonElement struct {
	*SVGGraphicsElement
	*SVGAnimatedPoints
}

type SVGPolylineElement struct {
	*SVGGraphicsElement
	*SVGAnimatedPoints
}

type SVGPreserveAspectRatio struct {
}

func (*SVGPreserveAspectRatio) GetAlign() (align uint8) {
	js.Rewrite("$<.align")
	return align
}

func (*SVGPreserveAspectRatio) SetAlign(align uint8) {
	js.Rewrite("$<.align = $1", align)
}

func (*SVGPreserveAspectRatio) GetMeetOrSlice() (meetOrSlice uint8) {
	js.Rewrite("$<.meetOrSlice")
	return meetOrSlice
}

func (*SVGPreserveAspectRatio) SetMeetOrSlice(meetOrSlice uint8) {
	js.Rewrite("$<.meetOrSlice = $1", meetOrSlice)
}

type SVGRadialGradientElement struct {
	*SVGGradientElement
}

func (*SVGRadialGradientElement) GetCx() (cx *SVGAnimatedLength) {
	js.Rewrite("$<.cx")
	return cx
}

func (*SVGRadialGradientElement) GetCy() (cy *SVGAnimatedLength) {
	js.Rewrite("$<.cy")
	return cy
}

func (*SVGRadialGradientElement) GetFx() (fx *SVGAnimatedLength) {
	js.Rewrite("$<.fx")
	return fx
}

func (*SVGRadialGradientElement) GetFy() (fy *SVGAnimatedLength) {
	js.Rewrite("$<.fy")
	return fy
}

func (*SVGRadialGradientElement) GetR() (r *SVGAnimatedLength) {
	js.Rewrite("$<.r")
	return r
}

type SVGRect struct {
}

func (*SVGRect) GetHeight() (height float32) {
	js.Rewrite("$<.height")
	return height
}

func (*SVGRect) SetHeight(height float32) {
	js.Rewrite("$<.height = $1", height)
}

func (*SVGRect) GetWidth() (width float32) {
	js.Rewrite("$<.width")
	return width
}

func (*SVGRect) SetWidth(width float32) {
	js.Rewrite("$<.width = $1", width)
}

func (*SVGRect) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGRect) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*SVGRect) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*SVGRect) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}

type SVGRectElement struct {
	*SVGGraphicsElement
}

func (*SVGRectElement) GetHeight() (height *SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

func (*SVGRectElement) GetRx() (rx *SVGAnimatedLength) {
	js.Rewrite("$<.rx")
	return rx
}

func (*SVGRectElement) GetRy() (ry *SVGAnimatedLength) {
	js.Rewrite("$<.ry")
	return ry
}

func (*SVGRectElement) GetWidth() (width *SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

func (*SVGRectElement) GetX() (x *SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGRectElement) GetY() (y *SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}

type SVGScriptElement struct {
	*SVGElement
	*SVGURIReference
}

func (*SVGScriptElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*SVGScriptElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}

type SVGStopElement struct {
	*SVGElement
}

func (*SVGStopElement) GetOffset() (offset *SVGAnimatedNumber) {
	js.Rewrite("$<.offset")
	return offset
}

type SVGStringList struct {
}

func (*SVGStringList) AppendItem(newItem string) (s string) {
	js.Rewrite("$<.appendItem($1)", newItem)
	return s
}

func (*SVGStringList) Clear() {
	js.Rewrite("$<.clear()")
}

func (*SVGStringList) GetItem(index uint) (s string) {
	js.Rewrite("$<.getItem($1)", index)
	return s
}

func (*SVGStringList) Initialize(newItem string) (s string) {
	js.Rewrite("$<.initialize($1)", newItem)
	return s
}

func (*SVGStringList) InsertItemBefore(newItem string, index uint) (s string) {
	js.Rewrite("$<.insertItemBefore($1, $2)", newItem, index)
	return s
}

func (*SVGStringList) RemoveItem(index uint) (s string) {
	js.Rewrite("$<.removeItem($1)", index)
	return s
}

func (*SVGStringList) ReplaceItem(newItem string, index uint) (s string) {
	js.Rewrite("$<.replaceItem($1, $2)", newItem, index)
	return s
}

func (*SVGStringList) GetNumberOfItems() (numberOfItems uint) {
	js.Rewrite("$<.numberOfItems")
	return numberOfItems
}

type SVGStyleElement struct {
	*SVGElement
}

func (*SVGStyleElement) GetDisabled() (disabled bool) {
	js.Rewrite("$<.disabled")
	return disabled
}

func (*SVGStyleElement) SetDisabled(disabled bool) {
	js.Rewrite("$<.disabled = $1", disabled)
}

func (*SVGStyleElement) GetMedia() (media string) {
	js.Rewrite("$<.media")
	return media
}

func (*SVGStyleElement) SetMedia(media string) {
	js.Rewrite("$<.media = $1", media)
}

func (*SVGStyleElement) GetTitle() (title string) {
	js.Rewrite("$<.title")
	return title
}

func (*SVGStyleElement) SetTitle(title string) {
	js.Rewrite("$<.title = $1", title)
}

func (*SVGStyleElement) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*SVGStyleElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}

type SVGSVGElement struct {
	*SVGGraphicsElement
	*DocumentEvent
	*SVGFitToViewBox
	*SVGZoomAndPan
}

func (*SVGSVGElement) CheckEnclosure(element *SVGElement, rect *SVGRect) (b bool) {
	js.Rewrite("$<.checkEnclosure($1, $2)", element, rect)
	return b
}

func (*SVGSVGElement) CheckIntersection(element *SVGElement, rect *SVGRect) (b bool) {
	js.Rewrite("$<.checkIntersection($1, $2)", element, rect)
	return b
}

func (*SVGSVGElement) CreateSVGAngle() (s *SVGAngle) {
	js.Rewrite("$<.createSVGAngle()")
	return s
}

func (*SVGSVGElement) CreateSVGLength() (s *SVGLength) {
	js.Rewrite("$<.createSVGLength()")
	return s
}

func (*SVGSVGElement) CreateSVGMatrix() (s *SVGMatrix) {
	js.Rewrite("$<.createSVGMatrix()")
	return s
}

func (*SVGSVGElement) CreateSVGNumber() (s *SVGNumber) {
	js.Rewrite("$<.createSVGNumber()")
	return s
}

func (*SVGSVGElement) CreateSVGPoint() (s *SVGPoint) {
	js.Rewrite("$<.createSVGPoint()")
	return s
}

func (*SVGSVGElement) CreateSVGRect() (s *SVGRect) {
	js.Rewrite("$<.createSVGRect()")
	return s
}

func (*SVGSVGElement) CreateSVGTransform() (s *SVGTransform) {
	js.Rewrite("$<.createSVGTransform()")
	return s
}

func (*SVGSVGElement) CreateSVGTransformFromMatrix(matrix *SVGMatrix) (s *SVGTransform) {
	js.Rewrite("$<.createSVGTransformFromMatrix($1)", matrix)
	return s
}

func (*SVGSVGElement) DeselectAll() {
	js.Rewrite("$<.deselectAll()")
}

func (*SVGSVGElement) ForceRedraw() {
	js.Rewrite("$<.forceRedraw()")
}

func (*SVGSVGElement) GetComputedStyle(elt *Element, pseudoElt string) (c *CSSStyleDeclaration) {
	js.Rewrite("$<.getComputedStyle($1, $2)", elt, pseudoElt)
	return c
}

func (*SVGSVGElement) GetCurrentTime() (f float32) {
	js.Rewrite("$<.getCurrentTime()")
	return f
}

func (*SVGSVGElement) GetElementByID(elementId string) (e *Element) {
	js.Rewrite("$<.getElementById($1)", elementId)
	return e
}

func (*SVGSVGElement) GetEnclosureList(rect *SVGRect, referenceElement *SVGElement) (n *NodeList) {
	js.Rewrite("$<.getEnclosureList($1, $2)", rect, referenceElement)
	return n
}

func (*SVGSVGElement) GetIntersectionList(rect *SVGRect, referenceElement *SVGElement) (n *NodeList) {
	js.Rewrite("$<.getIntersectionList($1, $2)", rect, referenceElement)
	return n
}

func (*SVGSVGElement) PauseAnimations() {
	js.Rewrite("$<.pauseAnimations()")
}

func (*SVGSVGElement) SetCurrentTime(seconds float32) {
	js.Rewrite("$<.setCurrentTime($1)", seconds)
}

func (*SVGSVGElement) SuspendRedraw(maxWaitMilliseconds uint) (u uint) {
	js.Rewrite("$<.suspendRedraw($1)", maxWaitMilliseconds)
	return u
}

func (*SVGSVGElement) UnpauseAnimations() {
	js.Rewrite("$<.unpauseAnimations()")
}

func (*SVGSVGElement) UnsuspendRedraw(suspendHandleID uint) {
	js.Rewrite("$<.unsuspendRedraw($1)", suspendHandleID)
}

func (*SVGSVGElement) UnsuspendRedrawAll() {
	js.Rewrite("$<.unsuspendRedrawAll()")
}

func (*SVGSVGElement) GetContentScriptType() (contentScriptType string) {
	js.Rewrite("$<.contentScriptType")
	return contentScriptType
}

func (*SVGSVGElement) SetContentScriptType(contentScriptType string) {
	js.Rewrite("$<.contentScriptType = $1", contentScriptType)
}

func (*SVGSVGElement) GetContentStyleType() (contentStyleType string) {
	js.Rewrite("$<.contentStyleType")
	return contentStyleType
}

func (*SVGSVGElement) SetContentStyleType(contentStyleType string) {
	js.Rewrite("$<.contentStyleType = $1", contentStyleType)
}

func (*SVGSVGElement) GetCurrentScale() (currentScale float32) {
	js.Rewrite("$<.currentScale")
	return currentScale
}

func (*SVGSVGElement) SetCurrentScale(currentScale float32) {
	js.Rewrite("$<.currentScale = $1", currentScale)
}

func (*SVGSVGElement) GetCurrentTranslate() (currentTranslate *SVGPoint) {
	js.Rewrite("$<.currentTranslate")
	return currentTranslate
}

func (*SVGSVGElement) GetHeight() (height *SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

func (*SVGSVGElement) GetOnabort() (e *Event) {
	js.Rewrite("$<.onabort")
	return e
}

func (*SVGSVGElement) SetOnabort(e *Event) {
	js.Rewrite("$<.onabort = $1", e)
}

func (*SVGSVGElement) GetOnerror() (e *Event) {
	js.Rewrite("$<.onerror")
	return e
}

func (*SVGSVGElement) SetOnerror(e *Event) {
	js.Rewrite("$<.onerror = $1", e)
}

func (*SVGSVGElement) GetOnresize() (e *Event) {
	js.Rewrite("$<.onresize")
	return e
}

func (*SVGSVGElement) SetOnresize(e *Event) {
	js.Rewrite("$<.onresize = $1", e)
}

func (*SVGSVGElement) GetOnscroll() (scroll *UIEvent) {
	js.Rewrite("$<.onscroll")
	return scroll
}

func (*SVGSVGElement) SetOnscroll(scroll *UIEvent) {
	js.Rewrite("$<.onscroll = $1", scroll)
}

func (*SVGSVGElement) GetOnunload() (e *Event) {
	js.Rewrite("$<.onunload")
	return e
}

func (*SVGSVGElement) SetOnunload(e *Event) {
	js.Rewrite("$<.onunload = $1", e)
}

func (*SVGSVGElement) GetOnzoom() (SVGZoom *SVGZoomEvent) {
	js.Rewrite("$<.onzoom")
	return SVGZoom
}

func (*SVGSVGElement) SetOnzoom(SVGZoom *SVGZoomEvent) {
	js.Rewrite("$<.onzoom = $1", SVGZoom)
}

func (*SVGSVGElement) GetPixelUnitToMillimeterX() (pixelUnitToMillimeterX float32) {
	js.Rewrite("$<.pixelUnitToMillimeterX")
	return pixelUnitToMillimeterX
}

func (*SVGSVGElement) GetPixelUnitToMillimeterY() (pixelUnitToMillimeterY float32) {
	js.Rewrite("$<.pixelUnitToMillimeterY")
	return pixelUnitToMillimeterY
}

func (*SVGSVGElement) GetScreenPixelToMillimeterX() (screenPixelToMillimeterX float32) {
	js.Rewrite("$<.screenPixelToMillimeterX")
	return screenPixelToMillimeterX
}

func (*SVGSVGElement) GetScreenPixelToMillimeterY() (screenPixelToMillimeterY float32) {
	js.Rewrite("$<.screenPixelToMillimeterY")
	return screenPixelToMillimeterY
}

func (*SVGSVGElement) GetViewport() (viewport *SVGRect) {
	js.Rewrite("$<.viewport")
	return viewport
}

func (*SVGSVGElement) GetWidth() (width *SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

func (*SVGSVGElement) GetX() (x *SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGSVGElement) GetY() (y *SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}

type SVGSwitchElement struct {
	*SVGGraphicsElement
}

type SVGSymbolElement struct {
	*SVGElement
	*SVGFitToViewBox
}

type SVGTextContentElement struct {
	*SVGGraphicsElement
}

func (*SVGTextContentElement) GetCharNumAtPosition(point *SVGPoint) (i int) {
	js.Rewrite("$<.getCharNumAtPosition($1)", point)
	return i
}

func (*SVGTextContentElement) GetComputedTextLength() (f float32) {
	js.Rewrite("$<.getComputedTextLength()")
	return f
}

func (*SVGTextContentElement) GetEndPositionOfChar(charnum uint) (s *SVGPoint) {
	js.Rewrite("$<.getEndPositionOfChar($1)", charnum)
	return s
}

func (*SVGTextContentElement) GetExtentOfChar(charnum uint) (s *SVGRect) {
	js.Rewrite("$<.getExtentOfChar($1)", charnum)
	return s
}

func (*SVGTextContentElement) GetNumberOfChars() (i int) {
	js.Rewrite("$<.getNumberOfChars()")
	return i
}

func (*SVGTextContentElement) GetRotationOfChar(charnum uint) (f float32) {
	js.Rewrite("$<.getRotationOfChar($1)", charnum)
	return f
}

func (*SVGTextContentElement) GetStartPositionOfChar(charnum uint) (s *SVGPoint) {
	js.Rewrite("$<.getStartPositionOfChar($1)", charnum)
	return s
}

func (*SVGTextContentElement) GetSubStringLength(charnum uint, nchars uint) (f float32) {
	js.Rewrite("$<.getSubStringLength($1, $2)", charnum, nchars)
	return f
}

func (*SVGTextContentElement) SelectSubString(charnum uint, nchars uint) {
	js.Rewrite("$<.selectSubString($1, $2)", charnum, nchars)
}

func (*SVGTextContentElement) GetLengthAdjust() (lengthAdjust *SVGAnimatedEnumeration) {
	js.Rewrite("$<.lengthAdjust")
	return lengthAdjust
}

func (*SVGTextContentElement) GetTextLength() (textLength *SVGAnimatedLength) {
	js.Rewrite("$<.textLength")
	return textLength
}

type SVGTextElement struct {
	*SVGTextPositioningElement
}

type SVGTextPathElement struct {
	*SVGTextContentElement
	*SVGURIReference
}

func (*SVGTextPathElement) GetMethod() (method *SVGAnimatedEnumeration) {
	js.Rewrite("$<.method")
	return method
}

func (*SVGTextPathElement) GetSpacing() (spacing *SVGAnimatedEnumeration) {
	js.Rewrite("$<.spacing")
	return spacing
}

func (*SVGTextPathElement) GetStartOffset() (startOffset *SVGAnimatedLength) {
	js.Rewrite("$<.startOffset")
	return startOffset
}

type SVGTextPositioningElement struct {
	*SVGTextContentElement
}

func (*SVGTextPositioningElement) GetDx() (dx *SVGAnimatedLengthList) {
	js.Rewrite("$<.dx")
	return dx
}

func (*SVGTextPositioningElement) GetDy() (dy *SVGAnimatedLengthList) {
	js.Rewrite("$<.dy")
	return dy
}

func (*SVGTextPositioningElement) GetRotate() (rotate *SVGAnimatedNumberList) {
	js.Rewrite("$<.rotate")
	return rotate
}

func (*SVGTextPositioningElement) GetX() (x *SVGAnimatedLengthList) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGTextPositioningElement) GetY() (y *SVGAnimatedLengthList) {
	js.Rewrite("$<.y")
	return y
}

type SVGTitleElement struct {
	*SVGElement
}

type SVGTransform struct {
}

func (*SVGTransform) SetMatrix(matrix *SVGMatrix) {
	js.Rewrite("$<.setMatrix($1)", matrix)
}

func (*SVGTransform) SetRotate(angle float32, cx float32, cy float32) {
	js.Rewrite("$<.setRotate($1, $2, $3)", angle, cx, cy)
}

func (*SVGTransform) SetScale(sx float32, sy float32) {
	js.Rewrite("$<.setScale($1, $2)", sx, sy)
}

func (*SVGTransform) SetSkewX(angle float32) {
	js.Rewrite("$<.setSkewX($1)", angle)
}

func (*SVGTransform) SetSkewY(angle float32) {
	js.Rewrite("$<.setSkewY($1)", angle)
}

func (*SVGTransform) SetTranslate(tx float32, ty float32) {
	js.Rewrite("$<.setTranslate($1, $2)", tx, ty)
}

func (*SVGTransform) GetAngle() (angle float32) {
	js.Rewrite("$<.angle")
	return angle
}

func (*SVGTransform) GetMatrix() (matrix *SVGMatrix) {
	js.Rewrite("$<.matrix")
	return matrix
}

func (*SVGTransform) GetType() (kind uint8) {
	js.Rewrite("$<.type")
	return kind
}

type SVGTransformList struct {
}

func (*SVGTransformList) AppendItem(newItem *SVGTransform) (s *SVGTransform) {
	js.Rewrite("$<.appendItem($1)", newItem)
	return s
}

func (*SVGTransformList) Clear() {
	js.Rewrite("$<.clear()")
}

func (*SVGTransformList) Consolidate() (s *SVGTransform) {
	js.Rewrite("$<.consolidate()")
	return s
}

func (*SVGTransformList) CreateSVGTransformFromMatrix(matrix *SVGMatrix) (s *SVGTransform) {
	js.Rewrite("$<.createSVGTransformFromMatrix($1)", matrix)
	return s
}

func (*SVGTransformList) GetItem(index uint) (s *SVGTransform) {
	js.Rewrite("$<.getItem($1)", index)
	return s
}

func (*SVGTransformList) Initialize(newItem *SVGTransform) (s *SVGTransform) {
	js.Rewrite("$<.initialize($1)", newItem)
	return s
}

func (*SVGTransformList) InsertItemBefore(newItem *SVGTransform, index uint) (s *SVGTransform) {
	js.Rewrite("$<.insertItemBefore($1, $2)", newItem, index)
	return s
}

func (*SVGTransformList) RemoveItem(index uint) (s *SVGTransform) {
	js.Rewrite("$<.removeItem($1)", index)
	return s
}

func (*SVGTransformList) ReplaceItem(newItem *SVGTransform, index uint) (s *SVGTransform) {
	js.Rewrite("$<.replaceItem($1, $2)", newItem, index)
	return s
}

func (*SVGTransformList) GetNumberOfItems() (numberOfItems uint) {
	js.Rewrite("$<.numberOfItems")
	return numberOfItems
}

type SVGTSpanElement struct {
	*SVGTextPositioningElement
}

type SVGUnitTypes struct {
}

type SVGUseElement struct {
	*SVGGraphicsElement
	*SVGURIReference
}

func (*SVGUseElement) GetAnimatedInstanceRoot() (animatedInstanceRoot *SVGElementInstance) {
	js.Rewrite("$<.animatedInstanceRoot")
	return animatedInstanceRoot
}

func (*SVGUseElement) GetHeight() (height *SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

func (*SVGUseElement) GetInstanceRoot() (instanceRoot *SVGElementInstance) {
	js.Rewrite("$<.instanceRoot")
	return instanceRoot
}

func (*SVGUseElement) GetWidth() (width *SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

func (*SVGUseElement) GetX() (x *SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGUseElement) GetY() (y *SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}

type SVGViewElement struct {
	*SVGElement
	*SVGZoomAndPan
	*SVGFitToViewBox
}

func (*SVGViewElement) GetViewTarget() (viewTarget *SVGStringList) {
	js.Rewrite("$<.viewTarget")
	return viewTarget
}

type SVGZoomAndPan struct {
}

func (*SVGZoomAndPan) GetZoomAndPan() (zoomAndPan uint8) {
	js.Rewrite("$<.zoomAndPan")
	return zoomAndPan
}

type SVGZoomEvent struct {
	*UIEvent
}

func (*SVGZoomEvent) GetNewScale() (newScale float32) {
	js.Rewrite("$<.newScale")
	return newScale
}

func (*SVGZoomEvent) GetNewTranslate() (newTranslate *SVGPoint) {
	js.Rewrite("$<.newTranslate")
	return newTranslate
}

func (*SVGZoomEvent) GetPreviousScale() (previousScale float32) {
	js.Rewrite("$<.previousScale")
	return previousScale
}

func (*SVGZoomEvent) GetPreviousTranslate() (previousTranslate *SVGPoint) {
	js.Rewrite("$<.previousTranslate")
	return previousTranslate
}

func (*SVGZoomEvent) GetZoomRectScreen() (zoomRectScreen *SVGRect) {
	js.Rewrite("$<.zoomRectScreen")
	return zoomRectScreen
}

type SyncManager struct {
}

func (*SyncManager) GetTags() (s []string) {
	js.Rewrite("await $<.getTags()")
	return s
}

func (*SyncManager) Register(tag string) {
	js.Rewrite("await $<.register($1)", tag)
}

type Text struct {
	*CharacterData
}

func (*Text) SplitText(offset uint) (t *Text) {
	js.Rewrite("$<.splitText($1)", offset)
	return t
}

func (*Text) GetWholeText() (wholeText string) {
	js.Rewrite("$<.wholeText")
	return wholeText
}

type TextEvent struct {
	*UIEvent
}

func (*TextEvent) InitTextEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, dataArg string, inputMethod uint, locale string) {
	js.Rewrite("$<.initTextEvent($1, $2, $3, $4, $5, $6, $7)", typeArg, canBubbleArg, cancelableArg, viewArg, dataArg, inputMethod, locale)
}

func (*TextEvent) GetData() (data string) {
	js.Rewrite("$<.data")
	return data
}

func (*TextEvent) GetInputMethod() (inputMethod uint) {
	js.Rewrite("$<.inputMethod")
	return inputMethod
}

func (*TextEvent) GetLocale() (locale string) {
	js.Rewrite("$<.locale")
	return locale
}

type TextMetrics struct {
}

func (*TextMetrics) GetWidth() (width float32) {
	js.Rewrite("$<.width")
	return width
}

type TextTrack struct {
	*EventTarget
}

func (*TextTrack) AddCue(cue *TextTrackCue) {
	js.Rewrite("$<.addCue($1)", cue)
}

func (*TextTrack) RemoveCue(cue *TextTrackCue) {
	js.Rewrite("$<.removeCue($1)", cue)
}

func (*TextTrack) GetActiveCues() (activeCues *TextTrackCueList) {
	js.Rewrite("$<.activeCues")
	return activeCues
}

func (*TextTrack) GetCues() (cues *TextTrackCueList) {
	js.Rewrite("$<.cues")
	return cues
}

func (*TextTrack) GetInBandMetadataTrackDispatchType() (inBandMetadataTrackDispatchType string) {
	js.Rewrite("$<.inBandMetadataTrackDispatchType")
	return inBandMetadataTrackDispatchType
}

func (*TextTrack) GetKind() (kind string) {
	js.Rewrite("$<.kind")
	return kind
}

func (*TextTrack) GetLabel() (label string) {
	js.Rewrite("$<.label")
	return label
}

func (*TextTrack) GetLanguage() (language string) {
	js.Rewrite("$<.language")
	return language
}

func (*TextTrack) GetMode() (mode interface{}) {
	js.Rewrite("$<.mode")
	return mode
}

func (*TextTrack) SetMode(mode interface{}) {
	js.Rewrite("$<.mode = $1", mode)
}

func (*TextTrack) GetOncuechange() (cuechange *Event) {
	js.Rewrite("$<.oncuechange")
	return cuechange
}

func (*TextTrack) SetOncuechange(cuechange *Event) {
	js.Rewrite("$<.oncuechange = $1", cuechange)
}

func (*TextTrack) GetOnerror() (err *Event) {
	js.Rewrite("$<.onerror")
	return err
}

func (*TextTrack) SetOnerror(err *Event) {
	js.Rewrite("$<.onerror = $1", err)
}

func (*TextTrack) GetOnload() (load *Event) {
	js.Rewrite("$<.onload")
	return load
}

func (*TextTrack) SetOnload(load *Event) {
	js.Rewrite("$<.onload = $1", load)
}

func (*TextTrack) GetReadyState() (readyState uint8) {
	js.Rewrite("$<.readyState")
	return readyState
}

type TextTrackCue struct {
	*EventTarget
}

func (*TextTrackCue) GetCueAsHTML() (d *DocumentFragment) {
	js.Rewrite("$<.getCueAsHTML()")
	return d
}

func (*TextTrackCue) GetEndTime() (endTime float32) {
	js.Rewrite("$<.endTime")
	return endTime
}

func (*TextTrackCue) SetEndTime(endTime float32) {
	js.Rewrite("$<.endTime = $1", endTime)
}

func (*TextTrackCue) GetID() (id string) {
	js.Rewrite("$<.id")
	return id
}

func (*TextTrackCue) SetID(id string) {
	js.Rewrite("$<.id = $1", id)
}

func (*TextTrackCue) GetOnenter() (enter *Event) {
	js.Rewrite("$<.onenter")
	return enter
}

func (*TextTrackCue) SetOnenter(enter *Event) {
	js.Rewrite("$<.onenter = $1", enter)
}

func (*TextTrackCue) GetOnexit() (exit *Event) {
	js.Rewrite("$<.onexit")
	return exit
}

func (*TextTrackCue) SetOnexit(exit *Event) {
	js.Rewrite("$<.onexit = $1", exit)
}

func (*TextTrackCue) GetPauseOnExit() (pauseOnExit bool) {
	js.Rewrite("$<.pauseOnExit")
	return pauseOnExit
}

func (*TextTrackCue) SetPauseOnExit(pauseOnExit bool) {
	js.Rewrite("$<.pauseOnExit = $1", pauseOnExit)
}

func (*TextTrackCue) GetStartTime() (startTime float32) {
	js.Rewrite("$<.startTime")
	return startTime
}

func (*TextTrackCue) SetStartTime(startTime float32) {
	js.Rewrite("$<.startTime = $1", startTime)
}

func (*TextTrackCue) GetText() (text string) {
	js.Rewrite("$<.text")
	return text
}

func (*TextTrackCue) SetText(text string) {
	js.Rewrite("$<.text = $1", text)
}

func (*TextTrackCue) GetTrack() (track *TextTrack) {
	js.Rewrite("$<.track")
	return track
}

type TextTrackCueList struct {
}

func (*TextTrackCueList) GetCueByID(id string) (t *TextTrackCue) {
	js.Rewrite("$<.getCueById($1)", id)
	return t
}

func (*TextTrackCueList) Item(index uint) (t *TextTrackCue) {
	js.Rewrite("$<.item($1)", index)
	return t
}

func (*TextTrackCueList) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

type TextTrackList struct {
	*EventTarget
}

func (*TextTrackList) Item(index uint) (t *TextTrack) {
	js.Rewrite("$<.item($1)", index)
	return t
}

func (*TextTrackList) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

func (*TextTrackList) GetOnaddtrack() (addtrack *TrackEvent) {
	js.Rewrite("$<.onaddtrack")
	return addtrack
}

func (*TextTrackList) SetOnaddtrack(addtrack *TrackEvent) {
	js.Rewrite("$<.onaddtrack = $1", addtrack)
}

type TimeRanges struct {
}

func (*TimeRanges) End(index uint) (f float32) {
	js.Rewrite("$<.end($1)", index)
	return f
}

func (*TimeRanges) Start(index uint) (f float32) {
	js.Rewrite("$<.start($1)", index)
	return f
}

func (*TimeRanges) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

type Touch struct {
}

func (*Touch) GetClientX() (clientX int) {
	js.Rewrite("$<.clientX")
	return clientX
}

func (*Touch) GetClientY() (clientY int) {
	js.Rewrite("$<.clientY")
	return clientY
}

func (*Touch) GetIdentifier() (identifier int) {
	js.Rewrite("$<.identifier")
	return identifier
}

func (*Touch) GetPageX() (pageX int) {
	js.Rewrite("$<.pageX")
	return pageX
}

func (*Touch) GetPageY() (pageY int) {
	js.Rewrite("$<.pageY")
	return pageY
}

func (*Touch) GetScreenX() (screenX int) {
	js.Rewrite("$<.screenX")
	return screenX
}

func (*Touch) GetScreenY() (screenY int) {
	js.Rewrite("$<.screenY")
	return screenY
}

func (*Touch) GetTarget() (target *EventTarget) {
	js.Rewrite("$<.target")
	return target
}

type TouchEvent struct {
	*UIEvent
}

func (*TouchEvent) GetAltKey() (altKey bool) {
	js.Rewrite("$<.altKey")
	return altKey
}

func (*TouchEvent) GetChangedTouches() (changedTouches *TouchList) {
	js.Rewrite("$<.changedTouches")
	return changedTouches
}

func (*TouchEvent) GetCharCode() (charCode int8) {
	js.Rewrite("$<.charCode")
	return charCode
}

func (*TouchEvent) GetCtrlKey() (ctrlKey bool) {
	js.Rewrite("$<.ctrlKey")
	return ctrlKey
}

func (*TouchEvent) GetKeyCode() (keyCode int8) {
	js.Rewrite("$<.keyCode")
	return keyCode
}

func (*TouchEvent) GetMetaKey() (metaKey bool) {
	js.Rewrite("$<.metaKey")
	return metaKey
}

func (*TouchEvent) GetShiftKey() (shiftKey bool) {
	js.Rewrite("$<.shiftKey")
	return shiftKey
}

func (*TouchEvent) GetTargetTouches() (targetTouches *TouchList) {
	js.Rewrite("$<.targetTouches")
	return targetTouches
}

func (*TouchEvent) GetTouches() (touches *TouchList) {
	js.Rewrite("$<.touches")
	return touches
}

func (*TouchEvent) GetWhich() (which uint8) {
	js.Rewrite("$<.which")
	return which
}

type TouchList struct {
}

func (*TouchList) Item(index uint) (t *Touch) {
	js.Rewrite("$<.item($1)", index)
	return t
}

func (*TouchList) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

type TrackEvent struct {
	*Event
}

func (*TrackEvent) GetTrack() (track interface{}) {
	js.Rewrite("$<.track")
	return track
}

type TransitionEvent struct {
	*Event
}

func (*TransitionEvent) InitTransitionEvent(typeArg string, canBubbleArg bool, cancelableArg bool, propertyNameArg string, elapsedTimeArg float32) {
	js.Rewrite("$<.initTransitionEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, propertyNameArg, elapsedTimeArg)
}

func (*TransitionEvent) GetElapsedTime() (elapsedTime float32) {
	js.Rewrite("$<.elapsedTime")
	return elapsedTime
}

func (*TransitionEvent) GetPropertyName() (propertyName string) {
	js.Rewrite("$<.propertyName")
	return propertyName
}

type TreeWalker struct {
}

func (*TreeWalker) FirstChild() (n *Node) {
	js.Rewrite("$<.firstChild()")
	return n
}

func (*TreeWalker) LastChild() (n *Node) {
	js.Rewrite("$<.lastChild()")
	return n
}

func (*TreeWalker) NextNode() (n *Node) {
	js.Rewrite("$<.nextNode()")
	return n
}

func (*TreeWalker) NextSibling() (n *Node) {
	js.Rewrite("$<.nextSibling()")
	return n
}

func (*TreeWalker) ParentNode() (n *Node) {
	js.Rewrite("$<.parentNode()")
	return n
}

func (*TreeWalker) PreviousNode() (n *Node) {
	js.Rewrite("$<.previousNode()")
	return n
}

func (*TreeWalker) PreviousSibling() (n *Node) {
	js.Rewrite("$<.previousSibling()")
	return n
}

func (*TreeWalker) GetCurrentNode() (currentNode *Node) {
	js.Rewrite("$<.currentNode")
	return currentNode
}

func (*TreeWalker) SetCurrentNode(currentNode *Node) {
	js.Rewrite("$<.currentNode = $1", currentNode)
}

func (*TreeWalker) GetExpandEntityReferences() (expandEntityReferences bool) {
	js.Rewrite("$<.expandEntityReferences")
	return expandEntityReferences
}

func (*TreeWalker) GetFilter() (filter *NodeFilter) {
	js.Rewrite("$<.filter")
	return filter
}

func (*TreeWalker) GetRoot() (root *Node) {
	js.Rewrite("$<.root")
	return root
}

func (*TreeWalker) GetWhatToShow() (whatToShow uint) {
	js.Rewrite("$<.whatToShow")
	return whatToShow
}

type UIEvent struct {
	*Event
}

func (*UIEvent) InitUIEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int) {
	js.Rewrite("$<.initUIEvent($1, $2, $3, $4, $5)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg)
}

func (*UIEvent) GetDetail() (detail int) {
	js.Rewrite("$<.detail")
	return detail
}

func (*UIEvent) GetView() (view *Window) {
	js.Rewrite("$<.view")
	return view
}

type UnviewableContentIdentifiedEvent struct {
	*NavigationEventWithReferrer
}

func (*UnviewableContentIdentifiedEvent) GetMediaType() (mediaType string) {
	js.Rewrite("$<.mediaType")
	return mediaType
}

type URL struct {
}

func (*URL) CreateObjectURL(object interface{}, options *ObjectURLOptions) (s string) {
	js.Rewrite("$<.createObjectURL($1, $2)", object, options)
	return s
}

func (*URL) RevokeObjectURL(url string) {
	js.Rewrite("$<.revokeObjectURL($1)", url)
}

func (*URL) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

func (*URL) GetHash() (hash string) {
	js.Rewrite("$<.hash")
	return hash
}

func (*URL) SetHash(hash string) {
	js.Rewrite("$<.hash = $1", hash)
}

func (*URL) GetHost() (host string) {
	js.Rewrite("$<.host")
	return host
}

func (*URL) SetHost(host string) {
	js.Rewrite("$<.host = $1", host)
}

func (*URL) GetHostname() (hostname string) {
	js.Rewrite("$<.hostname")
	return hostname
}

func (*URL) SetHostname(hostname string) {
	js.Rewrite("$<.hostname = $1", hostname)
}

func (*URL) GetHref() (href string) {
	js.Rewrite("$<.href")
	return href
}

func (*URL) SetHref(href string) {
	js.Rewrite("$<.href = $1", href)
}

func (*URL) GetOrigin() (origin string) {
	js.Rewrite("$<.origin")
	return origin
}

func (*URL) GetPassword() (password string) {
	js.Rewrite("$<.password")
	return password
}

func (*URL) SetPassword(password string) {
	js.Rewrite("$<.password = $1", password)
}

func (*URL) GetPathname() (pathname string) {
	js.Rewrite("$<.pathname")
	return pathname
}

func (*URL) SetPathname(pathname string) {
	js.Rewrite("$<.pathname = $1", pathname)
}

func (*URL) GetPort() (port string) {
	js.Rewrite("$<.port")
	return port
}

func (*URL) SetPort(port string) {
	js.Rewrite("$<.port = $1", port)
}

func (*URL) GetProtocol() (protocol string) {
	js.Rewrite("$<.protocol")
	return protocol
}

func (*URL) SetProtocol(protocol string) {
	js.Rewrite("$<.protocol = $1", protocol)
}

func (*URL) GetSearch() (search string) {
	js.Rewrite("$<.search")
	return search
}

func (*URL) SetSearch(search string) {
	js.Rewrite("$<.search = $1", search)
}

func (*URL) GetUsername() (username string) {
	js.Rewrite("$<.username")
	return username
}

func (*URL) SetUsername(username string) {
	js.Rewrite("$<.username = $1", username)
}

type ValidityState struct {
}

func (*ValidityState) GetBadInput() (badInput bool) {
	js.Rewrite("$<.badInput")
	return badInput
}

func (*ValidityState) GetCustomError() (customError bool) {
	js.Rewrite("$<.customError")
	return customError
}

func (*ValidityState) GetPatternMismatch() (patternMismatch bool) {
	js.Rewrite("$<.patternMismatch")
	return patternMismatch
}

func (*ValidityState) GetRangeOverflow() (rangeOverflow bool) {
	js.Rewrite("$<.rangeOverflow")
	return rangeOverflow
}

func (*ValidityState) GetRangeUnderflow() (rangeUnderflow bool) {
	js.Rewrite("$<.rangeUnderflow")
	return rangeUnderflow
}

func (*ValidityState) GetStepMismatch() (stepMismatch bool) {
	js.Rewrite("$<.stepMismatch")
	return stepMismatch
}

func (*ValidityState) GetTooLong() (tooLong bool) {
	js.Rewrite("$<.tooLong")
	return tooLong
}

func (*ValidityState) GetTypeMismatch() (typeMismatch bool) {
	js.Rewrite("$<.typeMismatch")
	return typeMismatch
}

func (*ValidityState) GetValid() (valid bool) {
	js.Rewrite("$<.valid")
	return valid
}

func (*ValidityState) GetValueMissing() (valueMissing bool) {
	js.Rewrite("$<.valueMissing")
	return valueMissing
}

type VideoPlaybackQuality struct {
}

func (*VideoPlaybackQuality) GetCorruptedVideoFrames() (corruptedVideoFrames uint) {
	js.Rewrite("$<.corruptedVideoFrames")
	return corruptedVideoFrames
}

func (*VideoPlaybackQuality) GetCreationTime() (creationTime int) {
	js.Rewrite("$<.creationTime")
	return creationTime
}

func (*VideoPlaybackQuality) GetDroppedVideoFrames() (droppedVideoFrames uint) {
	js.Rewrite("$<.droppedVideoFrames")
	return droppedVideoFrames
}

func (*VideoPlaybackQuality) GetTotalFrameDelay() (totalFrameDelay float32) {
	js.Rewrite("$<.totalFrameDelay")
	return totalFrameDelay
}

func (*VideoPlaybackQuality) GetTotalVideoFrames() (totalVideoFrames uint) {
	js.Rewrite("$<.totalVideoFrames")
	return totalVideoFrames
}

type VideoTrack struct {
}

func (*VideoTrack) GetID() (id string) {
	js.Rewrite("$<.id")
	return id
}

func (*VideoTrack) GetKind() (kind string) {
	js.Rewrite("$<.kind")
	return kind
}

func (*VideoTrack) SetKind(kind string) {
	js.Rewrite("$<.kind = $1", kind)
}

func (*VideoTrack) GetLabel() (label string) {
	js.Rewrite("$<.label")
	return label
}

func (*VideoTrack) GetLanguage() (language string) {
	js.Rewrite("$<.language")
	return language
}

func (*VideoTrack) SetLanguage(language string) {
	js.Rewrite("$<.language = $1", language)
}

func (*VideoTrack) GetSelected() (selected bool) {
	js.Rewrite("$<.selected")
	return selected
}

func (*VideoTrack) SetSelected(selected bool) {
	js.Rewrite("$<.selected = $1", selected)
}

func (*VideoTrack) GetSourceBuffer() (sourceBuffer *SourceBuffer) {
	js.Rewrite("$<.sourceBuffer")
	return sourceBuffer
}

type VideoTrackList struct {
	*EventTarget
}

func (*VideoTrackList) GetTrackByID(id string) (v *VideoTrack) {
	js.Rewrite("$<.getTrackById($1)", id)
	return v
}

func (*VideoTrackList) Item(index uint) (v *VideoTrack) {
	js.Rewrite("$<.item($1)", index)
	return v
}

func (*VideoTrackList) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

func (*VideoTrackList) GetOnaddtrack() (addtrack *TrackEvent) {
	js.Rewrite("$<.onaddtrack")
	return addtrack
}

func (*VideoTrackList) SetOnaddtrack(addtrack *TrackEvent) {
	js.Rewrite("$<.onaddtrack = $1", addtrack)
}

func (*VideoTrackList) GetOnchange() (change *Event) {
	js.Rewrite("$<.onchange")
	return change
}

func (*VideoTrackList) SetOnchange(change *Event) {
	js.Rewrite("$<.onchange = $1", change)
}

func (*VideoTrackList) GetOnremovetrack() (removetrack *TrackEvent) {
	js.Rewrite("$<.onremovetrack")
	return removetrack
}

func (*VideoTrackList) SetOnremovetrack(removetrack *TrackEvent) {
	js.Rewrite("$<.onremovetrack = $1", removetrack)
}

func (*VideoTrackList) GetSelectedIndex() (selectedIndex int) {
	js.Rewrite("$<.selectedIndex")
	return selectedIndex
}

type WaveShaperNode struct {
	*AudioNode
}

func (*WaveShaperNode) GetCurve() (curve []float32) {
	js.Rewrite("$<.curve")
	return curve
}

func (*WaveShaperNode) SetCurve(curve []float32) {
	js.Rewrite("$<.curve = $1", curve)
}

func (*WaveShaperNode) GetOversample() (oversample *OverSampleType) {
	js.Rewrite("$<.oversample")
	return oversample
}

func (*WaveShaperNode) SetOversample(oversample *OverSampleType) {
	js.Rewrite("$<.oversample = $1", oversample)
}

type WebAuthentication struct {
}

func (*WebAuthentication) GetAssertion(assertionChallenge []byte, options *AssertionOptions) (w *WebAuthnAssertion) {
	js.Rewrite("await $<.getAssertion($1, $2)", assertionChallenge, options)
	return w
}

func (*WebAuthentication) MakeCredential(accountInformation *Account, cryptoParameters []*ScopedCredentialParameters, attestationChallenge []byte, options *ScopedCredentialOptions) (s *ScopedCredentialInfo) {
	js.Rewrite("await $<.makeCredential($1, $2, $3, $4)", accountInformation, cryptoParameters, attestationChallenge, options)
	return s
}

type WebAuthnAssertion struct {
}

func (*WebAuthnAssertion) GetAuthenticatorData() (authenticatorData []byte) {
	js.Rewrite("$<.authenticatorData")
	return authenticatorData
}

func (*WebAuthnAssertion) GetClientData() (clientData []byte) {
	js.Rewrite("$<.clientData")
	return clientData
}

func (*WebAuthnAssertion) GetCredential() (credential *ScopedCredential) {
	js.Rewrite("$<.credential")
	return credential
}

func (*WebAuthnAssertion) GetSignature() (signature []byte) {
	js.Rewrite("$<.signature")
	return signature
}

type WEBGL_compressed_texture_s3tc struct {
}

type WEBGL_debug_renderer_info struct {
}

type WEBGL_depth_texture struct {
}

type WebGLActiveInfo struct {
}

func (*WebGLActiveInfo) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*WebGLActiveInfo) GetSize() (size int) {
	js.Rewrite("$<.size")
	return size
}

func (*WebGLActiveInfo) GetType() (kind uint) {
	js.Rewrite("$<.type")
	return kind
}

type WebGLBuffer struct {
	*WebGLObject
}

type WebGLContextEvent struct {
	*Event
}

func (*WebGLContextEvent) GetStatusMessage() (statusMessage string) {
	js.Rewrite("$<.statusMessage")
	return statusMessage
}

type WebGLFramebuffer struct {
	*WebGLObject
}

type WebGLObject struct {
}

type WebGLProgram struct {
	*WebGLObject
}

type WebGLRenderbuffer struct {
	*WebGLObject
}

type WebGLRenderingContext struct {
}

func (*WebGLRenderingContext) ActiveTexture(texture uint) {
	js.Rewrite("$<.activeTexture($1)", texture)
}

func (*WebGLRenderingContext) AttachShader(program *WebGLProgram, shader *WebGLShader) {
	js.Rewrite("$<.attachShader($1, $2)", program, shader)
}

func (*WebGLRenderingContext) BindAttribLocation(program *WebGLProgram, index uint, name string) {
	js.Rewrite("$<.bindAttribLocation($1, $2, $3)", program, index, name)
}

func (*WebGLRenderingContext) BindBuffer(target uint, buffer *WebGLBuffer) {
	js.Rewrite("$<.bindBuffer($1, $2)", target, buffer)
}

func (*WebGLRenderingContext) BindFramebuffer(target uint, framebuffer *WebGLFramebuffer) {
	js.Rewrite("$<.bindFramebuffer($1, $2)", target, framebuffer)
}

func (*WebGLRenderingContext) BindRenderbuffer(target uint, renderbuffer *WebGLRenderbuffer) {
	js.Rewrite("$<.bindRenderbuffer($1, $2)", target, renderbuffer)
}

func (*WebGLRenderingContext) BindTexture(target uint, texture *WebGLTexture) {
	js.Rewrite("$<.bindTexture($1, $2)", target, texture)
}

func (*WebGLRenderingContext) BlendColor(red float32, green float32, blue float32, alpha float32) {
	js.Rewrite("$<.blendColor($1, $2, $3, $4)", red, green, blue, alpha)
}

func (*WebGLRenderingContext) BlendEquation(mode uint) {
	js.Rewrite("$<.blendEquation($1)", mode)
}

func (*WebGLRenderingContext) BlendEquationSeparate(modeRGB uint, modeAlpha uint) {
	js.Rewrite("$<.blendEquationSeparate($1, $2)", modeRGB, modeAlpha)
}

func (*WebGLRenderingContext) BlendFunc(sfactor uint, dfactor uint) {
	js.Rewrite("$<.blendFunc($1, $2)", sfactor, dfactor)
}

func (*WebGLRenderingContext) BlendFuncSeparate(srcRGB uint, dstRGB uint, srcAlpha uint, dstAlpha uint) {
	js.Rewrite("$<.blendFuncSeparate($1, $2, $3, $4)", srcRGB, dstRGB, srcAlpha, dstAlpha)
}

func (*WebGLRenderingContext) BufferData(target uint, size interface{}, usage uint) {
	js.Rewrite("$<.bufferData($1, $2, $3)", target, size, usage)
}

func (*WebGLRenderingContext) BufferSubData(target uint, offset int64, data interface{}) {
	js.Rewrite("$<.bufferSubData($1, $2, $3)", target, offset, data)
}

func (*WebGLRenderingContext) CheckFramebufferStatus(target uint) (u uint) {
	js.Rewrite("$<.checkFramebufferStatus($1)", target)
	return u
}

func (*WebGLRenderingContext) Clear(mask uint) {
	js.Rewrite("$<.clear($1)", mask)
}

func (*WebGLRenderingContext) ClearColor(red float32, green float32, blue float32, alpha float32) {
	js.Rewrite("$<.clearColor($1, $2, $3, $4)", red, green, blue, alpha)
}

func (*WebGLRenderingContext) ClearDepth(depth float32) {
	js.Rewrite("$<.clearDepth($1)", depth)
}

func (*WebGLRenderingContext) ClearStencil(s int) {
	js.Rewrite("$<.clearStencil($1)", s)
}

func (*WebGLRenderingContext) ColorMask(red bool, green bool, blue bool, alpha bool) {
	js.Rewrite("$<.colorMask($1, $2, $3, $4)", red, green, blue, alpha)
}

func (*WebGLRenderingContext) CompileShader(shader *WebGLShader) {
	js.Rewrite("$<.compileShader($1)", shader)
}

func (*WebGLRenderingContext) CompressedTexImage2d(target uint, level int, internalformat uint, width int, height int, border int, data []byte) {
	js.Rewrite("$<.compressedTexImage2D($1, $2, $3, $4, $5, $6, $7)", target, level, internalformat, width, height, border, data)
}

func (*WebGLRenderingContext) CompressedTexSubImage2d(target uint, level int, xoffset int, yoffset int, width int, height int, format uint, data []byte) {
	js.Rewrite("$<.compressedTexSubImage2D($1, $2, $3, $4, $5, $6, $7, $8)", target, level, xoffset, yoffset, width, height, format, data)
}

func (*WebGLRenderingContext) CopyTexImage2d(target uint, level int, internalformat uint, x int, y int, width int, height int, border int) {
	js.Rewrite("$<.copyTexImage2D($1, $2, $3, $4, $5, $6, $7, $8)", target, level, internalformat, x, y, width, height, border)
}

func (*WebGLRenderingContext) CopyTexSubImage2d(target uint, level int, xoffset int, yoffset int, x int, y int, width int, height int) {
	js.Rewrite("$<.copyTexSubImage2D($1, $2, $3, $4, $5, $6, $7, $8)", target, level, xoffset, yoffset, x, y, width, height)
}

func (*WebGLRenderingContext) CreateBuffer() (w *WebGLBuffer) {
	js.Rewrite("$<.createBuffer()")
	return w
}

func (*WebGLRenderingContext) CreateFramebuffer() (w *WebGLFramebuffer) {
	js.Rewrite("$<.createFramebuffer()")
	return w
}

func (*WebGLRenderingContext) CreateProgram() (w *WebGLProgram) {
	js.Rewrite("$<.createProgram()")
	return w
}

func (*WebGLRenderingContext) CreateRenderbuffer() (w *WebGLRenderbuffer) {
	js.Rewrite("$<.createRenderbuffer()")
	return w
}

func (*WebGLRenderingContext) CreateShader(kind uint) (w *WebGLShader) {
	js.Rewrite("$<.createShader($1)", kind)
	return w
}

func (*WebGLRenderingContext) CreateTexture() (w *WebGLTexture) {
	js.Rewrite("$<.createTexture()")
	return w
}

func (*WebGLRenderingContext) CullFace(mode uint) {
	js.Rewrite("$<.cullFace($1)", mode)
}

func (*WebGLRenderingContext) DeleteBuffer(buffer *WebGLBuffer) {
	js.Rewrite("$<.deleteBuffer($1)", buffer)
}

func (*WebGLRenderingContext) DeleteFramebuffer(framebuffer *WebGLFramebuffer) {
	js.Rewrite("$<.deleteFramebuffer($1)", framebuffer)
}

func (*WebGLRenderingContext) DeleteProgram(program *WebGLProgram) {
	js.Rewrite("$<.deleteProgram($1)", program)
}

func (*WebGLRenderingContext) DeleteRenderbuffer(renderbuffer *WebGLRenderbuffer) {
	js.Rewrite("$<.deleteRenderbuffer($1)", renderbuffer)
}

func (*WebGLRenderingContext) DeleteShader(shader *WebGLShader) {
	js.Rewrite("$<.deleteShader($1)", shader)
}

func (*WebGLRenderingContext) DeleteTexture(texture *WebGLTexture) {
	js.Rewrite("$<.deleteTexture($1)", texture)
}

func (*WebGLRenderingContext) DepthFunc(fn uint) {
	js.Rewrite("$<.depthFunc($1)", fn)
}

func (*WebGLRenderingContext) DepthMask(flag bool) {
	js.Rewrite("$<.depthMask($1)", flag)
}

func (*WebGLRenderingContext) DepthRange(zNear float32, zFar float32) {
	js.Rewrite("$<.depthRange($1, $2)", zNear, zFar)
}

func (*WebGLRenderingContext) DetachShader(program *WebGLProgram, shader *WebGLShader) {
	js.Rewrite("$<.detachShader($1, $2)", program, shader)
}

func (*WebGLRenderingContext) Disable(capacity uint) {
	js.Rewrite("$<.disable($1)", capacity)
}

func (*WebGLRenderingContext) DisableVertexAttribArray(index uint) {
	js.Rewrite("$<.disableVertexAttribArray($1)", index)
}

func (*WebGLRenderingContext) DrawArrays(mode uint, first int, count int) {
	js.Rewrite("$<.drawArrays($1, $2, $3)", mode, first, count)
}

func (*WebGLRenderingContext) DrawElements(mode uint, count int, kind uint, offset int64) {
	js.Rewrite("$<.drawElements($1, $2, $3, $4)", mode, count, kind, offset)
}

func (*WebGLRenderingContext) Enable(capacity uint) {
	js.Rewrite("$<.enable($1)", capacity)
}

func (*WebGLRenderingContext) EnableVertexAttribArray(index uint) {
	js.Rewrite("$<.enableVertexAttribArray($1)", index)
}

func (*WebGLRenderingContext) Finish() {
	js.Rewrite("$<.finish()")
}

func (*WebGLRenderingContext) Flush() {
	js.Rewrite("$<.flush()")
}

func (*WebGLRenderingContext) FramebufferRenderbuffer(target uint, attachment uint, renderbuffertarget uint, renderbuffer *WebGLRenderbuffer) {
	js.Rewrite("$<.framebufferRenderbuffer($1, $2, $3, $4)", target, attachment, renderbuffertarget, renderbuffer)
}

func (*WebGLRenderingContext) FramebufferTexture2d(target uint, attachment uint, textarget uint, texture *WebGLTexture, level int) {
	js.Rewrite("$<.framebufferTexture2D($1, $2, $3, $4, $5)", target, attachment, textarget, texture, level)
}

func (*WebGLRenderingContext) FrontFace(mode uint) {
	js.Rewrite("$<.frontFace($1)", mode)
}

func (*WebGLRenderingContext) GenerateMipmap(target uint) {
	js.Rewrite("$<.generateMipmap($1)", target)
}

func (*WebGLRenderingContext) GetActiveAttrib(program *WebGLProgram, index uint) (w *WebGLActiveInfo) {
	js.Rewrite("$<.getActiveAttrib($1, $2)", program, index)
	return w
}

func (*WebGLRenderingContext) GetActiveUniform(program *WebGLProgram, index uint) (w *WebGLActiveInfo) {
	js.Rewrite("$<.getActiveUniform($1, $2)", program, index)
	return w
}

func (*WebGLRenderingContext) GetAttachedShaders(program *WebGLProgram) (w []*WebGLShader) {
	js.Rewrite("$<.getAttachedShaders($1)", program)
	return w
}

func (*WebGLRenderingContext) GetAttribLocation(program *WebGLProgram, name string) (i int) {
	js.Rewrite("$<.getAttribLocation($1, $2)", program, name)
	return i
}

func (*WebGLRenderingContext) GetBufferParameter(target uint, pname uint) (i interface{}) {
	js.Rewrite("$<.getBufferParameter($1, $2)", target, pname)
	return i
}

func (*WebGLRenderingContext) GetContextAttributes() (w *WebGLContextAttributes) {
	js.Rewrite("$<.getContextAttributes()")
	return w
}

func (*WebGLRenderingContext) GetError() (u uint) {
	js.Rewrite("$<.getError()")
	return u
}

func (*WebGLRenderingContext) GetExtension(name string) (i interface{}) {
	js.Rewrite("$<.getExtension($1)", name)
	return i
}

func (*WebGLRenderingContext) GetFramebufferAttachmentParameter(target uint, attachment uint, pname uint) (i interface{}) {
	js.Rewrite("$<.getFramebufferAttachmentParameter($1, $2, $3)", target, attachment, pname)
	return i
}

func (*WebGLRenderingContext) GetParameter(pname uint) (i interface{}) {
	js.Rewrite("$<.getParameter($1)", pname)
	return i
}

func (*WebGLRenderingContext) GetProgramInfoLog(program *WebGLProgram) (s string) {
	js.Rewrite("$<.getProgramInfoLog($1)", program)
	return s
}

func (*WebGLRenderingContext) GetProgramParameter(program *WebGLProgram, pname uint) (i interface{}) {
	js.Rewrite("$<.getProgramParameter($1, $2)", program, pname)
	return i
}

func (*WebGLRenderingContext) GetRenderbufferParameter(target uint, pname uint) (i interface{}) {
	js.Rewrite("$<.getRenderbufferParameter($1, $2)", target, pname)
	return i
}

func (*WebGLRenderingContext) GetShaderInfoLog(shader *WebGLShader) (s string) {
	js.Rewrite("$<.getShaderInfoLog($1)", shader)
	return s
}

func (*WebGLRenderingContext) GetShaderParameter(shader *WebGLShader, pname uint) (i interface{}) {
	js.Rewrite("$<.getShaderParameter($1, $2)", shader, pname)
	return i
}

func (*WebGLRenderingContext) GetShaderPrecisionFormat(shadertype uint, precisiontype uint) (w *WebGLShaderPrecisionFormat) {
	js.Rewrite("$<.getShaderPrecisionFormat($1, $2)", shadertype, precisiontype)
	return w
}

func (*WebGLRenderingContext) GetShaderSource(shader *WebGLShader) (s string) {
	js.Rewrite("$<.getShaderSource($1)", shader)
	return s
}

func (*WebGLRenderingContext) GetSupportedExtensions() (s []string) {
	js.Rewrite("$<.getSupportedExtensions()")
	return s
}

func (*WebGLRenderingContext) GetTexParameter(target uint, pname uint) (i interface{}) {
	js.Rewrite("$<.getTexParameter($1, $2)", target, pname)
	return i
}

func (*WebGLRenderingContext) GetUniform(program *WebGLProgram, location *WebGLUniformLocation) (i interface{}) {
	js.Rewrite("$<.getUniform($1, $2)", program, location)
	return i
}

func (*WebGLRenderingContext) GetUniformLocation(program *WebGLProgram, name string) (w *WebGLUniformLocation) {
	js.Rewrite("$<.getUniformLocation($1, $2)", program, name)
	return w
}

func (*WebGLRenderingContext) GetVertexAttrib(index uint, pname uint) (i interface{}) {
	js.Rewrite("$<.getVertexAttrib($1, $2)", index, pname)
	return i
}

func (*WebGLRenderingContext) GetVertexAttribOffset(index uint, pname uint) (i int64) {
	js.Rewrite("$<.getVertexAttribOffset($1, $2)", index, pname)
	return i
}

func (*WebGLRenderingContext) Hint(target uint, mode uint) {
	js.Rewrite("$<.hint($1, $2)", target, mode)
}

func (*WebGLRenderingContext) IsBuffer(buffer *WebGLBuffer) (b bool) {
	js.Rewrite("$<.isBuffer($1)", buffer)
	return b
}

func (*WebGLRenderingContext) IsContextLost() (b bool) {
	js.Rewrite("$<.isContextLost()")
	return b
}

func (*WebGLRenderingContext) IsEnabled(capacity uint) (b bool) {
	js.Rewrite("$<.isEnabled($1)", capacity)
	return b
}

func (*WebGLRenderingContext) IsFramebuffer(framebuffer *WebGLFramebuffer) (b bool) {
	js.Rewrite("$<.isFramebuffer($1)", framebuffer)
	return b
}

func (*WebGLRenderingContext) IsProgram(program *WebGLProgram) (b bool) {
	js.Rewrite("$<.isProgram($1)", program)
	return b
}

func (*WebGLRenderingContext) IsRenderbuffer(renderbuffer *WebGLRenderbuffer) (b bool) {
	js.Rewrite("$<.isRenderbuffer($1)", renderbuffer)
	return b
}

func (*WebGLRenderingContext) IsShader(shader *WebGLShader) (b bool) {
	js.Rewrite("$<.isShader($1)", shader)
	return b
}

func (*WebGLRenderingContext) IsTexture(texture *WebGLTexture) (b bool) {
	js.Rewrite("$<.isTexture($1)", texture)
	return b
}

func (*WebGLRenderingContext) LineWidth(width float32) {
	js.Rewrite("$<.lineWidth($1)", width)
}

func (*WebGLRenderingContext) LinkProgram(program *WebGLProgram) {
	js.Rewrite("$<.linkProgram($1)", program)
}

func (*WebGLRenderingContext) PixelStorei(pname uint, param int) {
	js.Rewrite("$<.pixelStorei($1, $2)", pname, param)
}

func (*WebGLRenderingContext) PolygonOffset(factor float32, units float32) {
	js.Rewrite("$<.polygonOffset($1, $2)", factor, units)
}

func (*WebGLRenderingContext) ReadPixels(x int, y int, width int, height int, format uint, kind uint, pixels []byte) {
	js.Rewrite("$<.readPixels($1, $2, $3, $4, $5, $6, $7)", x, y, width, height, format, kind, pixels)
}

func (*WebGLRenderingContext) RenderbufferStorage(target uint, internalformat uint, width int, height int) {
	js.Rewrite("$<.renderbufferStorage($1, $2, $3, $4)", target, internalformat, width, height)
}

func (*WebGLRenderingContext) SampleCoverage(value float32, invert bool) {
	js.Rewrite("$<.sampleCoverage($1, $2)", value, invert)
}

func (*WebGLRenderingContext) Scissor(x int, y int, width int, height int) {
	js.Rewrite("$<.scissor($1, $2, $3, $4)", x, y, width, height)
}

func (*WebGLRenderingContext) ShaderSource(shader *WebGLShader, source string) {
	js.Rewrite("$<.shaderSource($1, $2)", shader, source)
}

func (*WebGLRenderingContext) StencilFunc(fn uint, ref int, mask uint) {
	js.Rewrite("$<.stencilFunc($1, $2, $3)", fn, ref, mask)
}

func (*WebGLRenderingContext) StencilFuncSeparate(face uint, fn uint, ref int, mask uint) {
	js.Rewrite("$<.stencilFuncSeparate($1, $2, $3, $4)", face, fn, ref, mask)
}

func (*WebGLRenderingContext) StencilMask(mask uint) {
	js.Rewrite("$<.stencilMask($1)", mask)
}

func (*WebGLRenderingContext) StencilMaskSeparate(face uint, mask uint) {
	js.Rewrite("$<.stencilMaskSeparate($1, $2)", face, mask)
}

func (*WebGLRenderingContext) StencilOp(fail uint, zfail uint, zpass uint) {
	js.Rewrite("$<.stencilOp($1, $2, $3)", fail, zfail, zpass)
}

func (*WebGLRenderingContext) StencilOpSeparate(face uint, fail uint, zfail uint, zpass uint) {
	js.Rewrite("$<.stencilOpSeparate($1, $2, $3, $4)", face, fail, zfail, zpass)
}

func (*WebGLRenderingContext) TexImage2d(target uint, level int, internalformat uint, format uint, kind uint, pixels *ImageData) {
	js.Rewrite("$<.texImage2D($1, $2, $3, $4, $5, $6)", target, level, internalformat, format, kind, pixels)
}

func (*WebGLRenderingContext) TexParameterf(target uint, pname uint, param float32) {
	js.Rewrite("$<.texParameterf($1, $2, $3)", target, pname, param)
}

func (*WebGLRenderingContext) TexParameteri(target uint, pname uint, param int) {
	js.Rewrite("$<.texParameteri($1, $2, $3)", target, pname, param)
}

func (*WebGLRenderingContext) TexSubImage2d(target uint, level int, xoffset int, yoffset int, format uint, kind uint, pixels *ImageData) {
	js.Rewrite("$<.texSubImage2D($1, $2, $3, $4, $5, $6, $7)", target, level, xoffset, yoffset, format, kind, pixels)
}

func (*WebGLRenderingContext) Uniform1f(location *WebGLUniformLocation, x float32) {
	js.Rewrite("$<.uniform1f($1, $2)", location, x)
}

func (*WebGLRenderingContext) Uniform1fv(location *WebGLUniformLocation, v []float32) {
	js.Rewrite("$<.uniform1fv($1, $2)", location, v)
}

func (*WebGLRenderingContext) Uniform1i(location *WebGLUniformLocation, x int) {
	js.Rewrite("$<.uniform1i($1, $2)", location, x)
}

func (*WebGLRenderingContext) Uniform1iv(location *WebGLUniformLocation, v []int32) {
	js.Rewrite("$<.uniform1iv($1, $2)", location, v)
}

func (*WebGLRenderingContext) Uniform2f(location *WebGLUniformLocation, x float32, y float32) {
	js.Rewrite("$<.uniform2f($1, $2, $3)", location, x, y)
}

func (*WebGLRenderingContext) Uniform2fv(location *WebGLUniformLocation, v []float32) {
	js.Rewrite("$<.uniform2fv($1, $2)", location, v)
}

func (*WebGLRenderingContext) Uniform2i(location *WebGLUniformLocation, x int, y int) {
	js.Rewrite("$<.uniform2i($1, $2, $3)", location, x, y)
}

func (*WebGLRenderingContext) Uniform2iv(location *WebGLUniformLocation, v []int32) {
	js.Rewrite("$<.uniform2iv($1, $2)", location, v)
}

func (*WebGLRenderingContext) Uniform3f(location *WebGLUniformLocation, x float32, y float32, z float32) {
	js.Rewrite("$<.uniform3f($1, $2, $3, $4)", location, x, y, z)
}

func (*WebGLRenderingContext) Uniform3fv(location *WebGLUniformLocation, v []float32) {
	js.Rewrite("$<.uniform3fv($1, $2)", location, v)
}

func (*WebGLRenderingContext) Uniform3i(location *WebGLUniformLocation, x int, y int, z int) {
	js.Rewrite("$<.uniform3i($1, $2, $3, $4)", location, x, y, z)
}

func (*WebGLRenderingContext) Uniform3iv(location *WebGLUniformLocation, v []int32) {
	js.Rewrite("$<.uniform3iv($1, $2)", location, v)
}

func (*WebGLRenderingContext) Uniform4f(location *WebGLUniformLocation, x float32, y float32, z float32, w float32) {
	js.Rewrite("$<.uniform4f($1, $2, $3, $4, $5)", location, x, y, z, w)
}

func (*WebGLRenderingContext) Uniform4fv(location *WebGLUniformLocation, v []float32) {
	js.Rewrite("$<.uniform4fv($1, $2)", location, v)
}

func (*WebGLRenderingContext) Uniform4i(location *WebGLUniformLocation, x int, y int, z int, w int) {
	js.Rewrite("$<.uniform4i($1, $2, $3, $4, $5)", location, x, y, z, w)
}

func (*WebGLRenderingContext) Uniform4iv(location *WebGLUniformLocation, v []int32) {
	js.Rewrite("$<.uniform4iv($1, $2)", location, v)
}

func (*WebGLRenderingContext) UniformMatrix2fv(location *WebGLUniformLocation, transpose bool, value []float32) {
	js.Rewrite("$<.uniformMatrix2fv($1, $2, $3)", location, transpose, value)
}

func (*WebGLRenderingContext) UniformMatrix3fv(location *WebGLUniformLocation, transpose bool, value []float32) {
	js.Rewrite("$<.uniformMatrix3fv($1, $2, $3)", location, transpose, value)
}

func (*WebGLRenderingContext) UniformMatrix4fv(location *WebGLUniformLocation, transpose bool, value []float32) {
	js.Rewrite("$<.uniformMatrix4fv($1, $2, $3)", location, transpose, value)
}

func (*WebGLRenderingContext) UseProgram(program *WebGLProgram) {
	js.Rewrite("$<.useProgram($1)", program)
}

func (*WebGLRenderingContext) ValidateProgram(program *WebGLProgram) {
	js.Rewrite("$<.validateProgram($1)", program)
}

func (*WebGLRenderingContext) VertexAttrib1f(indx uint, x float32) {
	js.Rewrite("$<.vertexAttrib1f($1, $2)", indx, x)
}

func (*WebGLRenderingContext) VertexAttrib1fv(indx uint, values []float32) {
	js.Rewrite("$<.vertexAttrib1fv($1, $2)", indx, values)
}

func (*WebGLRenderingContext) VertexAttrib2f(indx uint, x float32, y float32) {
	js.Rewrite("$<.vertexAttrib2f($1, $2, $3)", indx, x, y)
}

func (*WebGLRenderingContext) VertexAttrib2fv(indx uint, values []float32) {
	js.Rewrite("$<.vertexAttrib2fv($1, $2)", indx, values)
}

func (*WebGLRenderingContext) VertexAttrib3f(indx uint, x float32, y float32, z float32) {
	js.Rewrite("$<.vertexAttrib3f($1, $2, $3, $4)", indx, x, y, z)
}

func (*WebGLRenderingContext) VertexAttrib3fv(indx uint, values []float32) {
	js.Rewrite("$<.vertexAttrib3fv($1, $2)", indx, values)
}

func (*WebGLRenderingContext) VertexAttrib4f(indx uint, x float32, y float32, z float32, w float32) {
	js.Rewrite("$<.vertexAttrib4f($1, $2, $3, $4, $5)", indx, x, y, z, w)
}

func (*WebGLRenderingContext) VertexAttrib4fv(indx uint, values []float32) {
	js.Rewrite("$<.vertexAttrib4fv($1, $2)", indx, values)
}

func (*WebGLRenderingContext) VertexAttribPointer(indx uint, size int, kind uint, normalized bool, stride int, offset int64) {
	js.Rewrite("$<.vertexAttribPointer($1, $2, $3, $4, $5, $6)", indx, size, kind, normalized, stride, offset)
}

func (*WebGLRenderingContext) Viewport(x int, y int, width int, height int) {
	js.Rewrite("$<.viewport($1, $2, $3, $4)", x, y, width, height)
}

func (*WebGLRenderingContext) GetCanvas() (canvas *HTMLCanvasElement) {
	js.Rewrite("$<.canvas")
	return canvas
}

func (*WebGLRenderingContext) GetDrawingBufferHeight() (drawingBufferHeight int) {
	js.Rewrite("$<.drawingBufferHeight")
	return drawingBufferHeight
}

func (*WebGLRenderingContext) GetDrawingBufferWidth() (drawingBufferWidth int) {
	js.Rewrite("$<.drawingBufferWidth")
	return drawingBufferWidth
}

type WebGLShader struct {
	*WebGLObject
}

type WebGLShaderPrecisionFormat struct {
}

func (*WebGLShaderPrecisionFormat) GetPrecision() (precision int) {
	js.Rewrite("$<.precision")
	return precision
}

func (*WebGLShaderPrecisionFormat) GetRangeMax() (rangeMax int) {
	js.Rewrite("$<.rangeMax")
	return rangeMax
}

func (*WebGLShaderPrecisionFormat) GetRangeMin() (rangeMin int) {
	js.Rewrite("$<.rangeMin")
	return rangeMin
}

type WebGLTexture struct {
	*WebGLObject
}

type WebGLUniformLocation struct {
}

type WebKitCSSMatrix struct {
}

func (*WebKitCSSMatrix) Inverse() (w *WebKitCSSMatrix) {
	js.Rewrite("$<.inverse()")
	return w
}

func (*WebKitCSSMatrix) Multiply(secondMatrix *WebKitCSSMatrix) (w *WebKitCSSMatrix) {
	js.Rewrite("$<.multiply($1)", secondMatrix)
	return w
}

func (*WebKitCSSMatrix) Rotate(angleX float32, angleY float32, angleZ float32) (w *WebKitCSSMatrix) {
	js.Rewrite("$<.rotate($1, $2, $3)", angleX, angleY, angleZ)
	return w
}

func (*WebKitCSSMatrix) RotateAxisAngle(x float32, y float32, z float32, angle float32) (w *WebKitCSSMatrix) {
	js.Rewrite("$<.rotateAxisAngle($1, $2, $3, $4)", x, y, z, angle)
	return w
}

func (*WebKitCSSMatrix) Scale(scaleX float32, scaleY float32, scaleZ float32) (w *WebKitCSSMatrix) {
	js.Rewrite("$<.scale($1, $2, $3)", scaleX, scaleY, scaleZ)
	return w
}

func (*WebKitCSSMatrix) SetMatrixValue(value string) {
	js.Rewrite("$<.setMatrixValue($1)", value)
}

func (*WebKitCSSMatrix) SkewX(angle float32) (w *WebKitCSSMatrix) {
	js.Rewrite("$<.skewX($1)", angle)
	return w
}

func (*WebKitCSSMatrix) SkewY(angle float32) (w *WebKitCSSMatrix) {
	js.Rewrite("$<.skewY($1)", angle)
	return w
}

func (*WebKitCSSMatrix) ToString() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

func (*WebKitCSSMatrix) Translate(x float32, y float32, z float32) (w *WebKitCSSMatrix) {
	js.Rewrite("$<.translate($1, $2, $3)", x, y, z)
	return w
}

func (*WebKitCSSMatrix) GetA() (a float32) {
	js.Rewrite("$<.a")
	return a
}

func (*WebKitCSSMatrix) SetA(a float32) {
	js.Rewrite("$<.a = $1", a)
}

func (*WebKitCSSMatrix) GetB() (b float32) {
	js.Rewrite("$<.b")
	return b
}

func (*WebKitCSSMatrix) SetB(b float32) {
	js.Rewrite("$<.b = $1", b)
}

func (*WebKitCSSMatrix) GetC() (c float32) {
	js.Rewrite("$<.c")
	return c
}

func (*WebKitCSSMatrix) SetC(c float32) {
	js.Rewrite("$<.c = $1", c)
}

func (*WebKitCSSMatrix) GetD() (d float32) {
	js.Rewrite("$<.d")
	return d
}

func (*WebKitCSSMatrix) SetD(d float32) {
	js.Rewrite("$<.d = $1", d)
}

func (*WebKitCSSMatrix) GetE() (e float32) {
	js.Rewrite("$<.e")
	return e
}

func (*WebKitCSSMatrix) SetE(e float32) {
	js.Rewrite("$<.e = $1", e)
}

func (*WebKitCSSMatrix) GetF() (f float32) {
	js.Rewrite("$<.f")
	return f
}

func (*WebKitCSSMatrix) SetF(f float32) {
	js.Rewrite("$<.f = $1", f)
}

func (*WebKitCSSMatrix) GetM11() (m11 float32) {
	js.Rewrite("$<.m11")
	return m11
}

func (*WebKitCSSMatrix) SetM11(m11 float32) {
	js.Rewrite("$<.m11 = $1", m11)
}

func (*WebKitCSSMatrix) GetM12() (m12 float32) {
	js.Rewrite("$<.m12")
	return m12
}

func (*WebKitCSSMatrix) SetM12(m12 float32) {
	js.Rewrite("$<.m12 = $1", m12)
}

func (*WebKitCSSMatrix) GetM13() (m13 float32) {
	js.Rewrite("$<.m13")
	return m13
}

func (*WebKitCSSMatrix) SetM13(m13 float32) {
	js.Rewrite("$<.m13 = $1", m13)
}

func (*WebKitCSSMatrix) GetM14() (m14 float32) {
	js.Rewrite("$<.m14")
	return m14
}

func (*WebKitCSSMatrix) SetM14(m14 float32) {
	js.Rewrite("$<.m14 = $1", m14)
}

func (*WebKitCSSMatrix) GetM21() (m21 float32) {
	js.Rewrite("$<.m21")
	return m21
}

func (*WebKitCSSMatrix) SetM21(m21 float32) {
	js.Rewrite("$<.m21 = $1", m21)
}

func (*WebKitCSSMatrix) GetM22() (m22 float32) {
	js.Rewrite("$<.m22")
	return m22
}

func (*WebKitCSSMatrix) SetM22(m22 float32) {
	js.Rewrite("$<.m22 = $1", m22)
}

func (*WebKitCSSMatrix) GetM23() (m23 float32) {
	js.Rewrite("$<.m23")
	return m23
}

func (*WebKitCSSMatrix) SetM23(m23 float32) {
	js.Rewrite("$<.m23 = $1", m23)
}

func (*WebKitCSSMatrix) GetM24() (m24 float32) {
	js.Rewrite("$<.m24")
	return m24
}

func (*WebKitCSSMatrix) SetM24(m24 float32) {
	js.Rewrite("$<.m24 = $1", m24)
}

func (*WebKitCSSMatrix) GetM31() (m31 float32) {
	js.Rewrite("$<.m31")
	return m31
}

func (*WebKitCSSMatrix) SetM31(m31 float32) {
	js.Rewrite("$<.m31 = $1", m31)
}

func (*WebKitCSSMatrix) GetM32() (m32 float32) {
	js.Rewrite("$<.m32")
	return m32
}

func (*WebKitCSSMatrix) SetM32(m32 float32) {
	js.Rewrite("$<.m32 = $1", m32)
}

func (*WebKitCSSMatrix) GetM33() (m33 float32) {
	js.Rewrite("$<.m33")
	return m33
}

func (*WebKitCSSMatrix) SetM33(m33 float32) {
	js.Rewrite("$<.m33 = $1", m33)
}

func (*WebKitCSSMatrix) GetM34() (m34 float32) {
	js.Rewrite("$<.m34")
	return m34
}

func (*WebKitCSSMatrix) SetM34(m34 float32) {
	js.Rewrite("$<.m34 = $1", m34)
}

func (*WebKitCSSMatrix) GetM41() (m41 float32) {
	js.Rewrite("$<.m41")
	return m41
}

func (*WebKitCSSMatrix) SetM41(m41 float32) {
	js.Rewrite("$<.m41 = $1", m41)
}

func (*WebKitCSSMatrix) GetM42() (m42 float32) {
	js.Rewrite("$<.m42")
	return m42
}

func (*WebKitCSSMatrix) SetM42(m42 float32) {
	js.Rewrite("$<.m42 = $1", m42)
}

func (*WebKitCSSMatrix) GetM43() (m43 float32) {
	js.Rewrite("$<.m43")
	return m43
}

func (*WebKitCSSMatrix) SetM43(m43 float32) {
	js.Rewrite("$<.m43 = $1", m43)
}

func (*WebKitCSSMatrix) GetM44() (m44 float32) {
	js.Rewrite("$<.m44")
	return m44
}

func (*WebKitCSSMatrix) SetM44(m44 float32) {
	js.Rewrite("$<.m44 = $1", m44)
}

type WebKitDirectoryEntry struct {
	*WebKitEntry
}

func (*WebKitDirectoryEntry) CreateReader() (w *WebKitDirectoryReader) {
	js.Rewrite("$<.createReader()")
	return w
}

type WebKitDirectoryReader struct {
}

func (*WebKitDirectoryReader) ReadEntries(successCallback func(entries []*WebKitEntry), errorCallback func(err *DOMError)) {
	js.Rewrite("$<.readEntries($1, $2)", successCallback, errorCallback)
}

type WebKitEntry struct {
}

func (*WebKitEntry) GetFilesystem() (filesystem *WebKitFileSystem) {
	js.Rewrite("$<.filesystem")
	return filesystem
}

func (*WebKitEntry) GetFullPath() (fullPath string) {
	js.Rewrite("$<.fullPath")
	return fullPath
}

func (*WebKitEntry) GetIsDirectory() (isDirectory bool) {
	js.Rewrite("$<.isDirectory")
	return isDirectory
}

func (*WebKitEntry) GetIsFile() (isFile bool) {
	js.Rewrite("$<.isFile")
	return isFile
}

func (*WebKitEntry) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

type WebKitFileEntry struct {
	*WebKitEntry
}

func (*WebKitFileEntry) File(successCallback func(file *File), errorCallback func(err *DOMError)) {
	js.Rewrite("$<.file($1, $2)", successCallback, errorCallback)
}

type WebKitFileSystem struct {
}

func (*WebKitFileSystem) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*WebKitFileSystem) GetRoot() (root *WebKitDirectoryEntry) {
	js.Rewrite("$<.root")
	return root
}

type WebKitPoint struct {
}

func (*WebKitPoint) GetX() (x float32) {
	js.Rewrite("$<.x")
	return x
}

func (*WebKitPoint) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

func (*WebKitPoint) GetY() (y float32) {
	js.Rewrite("$<.y")
	return y
}

func (*WebKitPoint) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}

type webkitRTCPeerConnection struct {
	*RTCPeerConnection
}

type WebSocket struct {
	*EventTarget
}

func (*WebSocket) Close(code uint8, reason string) {
	js.Rewrite("$<.close($1, $2)", code, reason)
}

func (*WebSocket) Send(data interface{}) {
	js.Rewrite("$<.send($1)", data)
}

func (*WebSocket) GetBinaryType() (binaryType string) {
	js.Rewrite("$<.binaryType")
	return binaryType
}

func (*WebSocket) SetBinaryType(binaryType string) {
	js.Rewrite("$<.binaryType = $1", binaryType)
}

func (*WebSocket) GetBufferedAmount() (bufferedAmount uint) {
	js.Rewrite("$<.bufferedAmount")
	return bufferedAmount
}

func (*WebSocket) GetExtensions() (extensions string) {
	js.Rewrite("$<.extensions")
	return extensions
}

func (*WebSocket) GetOnclose() (cls *CloseEvent) {
	js.Rewrite("$<.onclose")
	return cls
}

func (*WebSocket) SetOnclose(cls *CloseEvent) {
	js.Rewrite("$<.onclose = $1", cls)
}

func (*WebSocket) GetOnerror() (err *Event) {
	js.Rewrite("$<.onerror")
	return err
}

func (*WebSocket) SetOnerror(err *Event) {
	js.Rewrite("$<.onerror = $1", err)
}

func (*WebSocket) GetOnmessage() (message *MessageEvent) {
	js.Rewrite("$<.onmessage")
	return message
}

func (*WebSocket) SetOnmessage(message *MessageEvent) {
	js.Rewrite("$<.onmessage = $1", message)
}

func (*WebSocket) GetOnopen() (open *Event) {
	js.Rewrite("$<.onopen")
	return open
}

func (*WebSocket) SetOnopen(open *Event) {
	js.Rewrite("$<.onopen = $1", open)
}

func (*WebSocket) GetProtocol() (protocol string) {
	js.Rewrite("$<.protocol")
	return protocol
}

func (*WebSocket) GetReadyState() (readyState uint8) {
	js.Rewrite("$<.readyState")
	return readyState
}

func (*WebSocket) GetURL() (url string) {
	js.Rewrite("$<.url")
	return url
}

type WheelEvent struct {
	*MouseEvent
}

func (*WheelEvent) GetCurrentPoint(element *Element) {
	js.Rewrite("$<.getCurrentPoint($1)", element)
}

func (*WheelEvent) InitWheelEvent(typeArg string, canBubbleArg bool, cancelableArg bool, viewArg *Window, detailArg int, screenXArg int, screenYArg int, clientXArg int, clientYArg int, buttonArg uint8, relatedTargetArg *EventTarget, modifiersListArg string, deltaXArg int, deltaYArg int, deltaZArg int, deltaMode uint) {
	js.Rewrite("$<.initWheelEvent($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)", typeArg, canBubbleArg, cancelableArg, viewArg, detailArg, screenXArg, screenYArg, clientXArg, clientYArg, buttonArg, relatedTargetArg, modifiersListArg, deltaXArg, deltaYArg, deltaZArg, deltaMode)
}

func (*WheelEvent) GetDeltaMode() (deltaMode uint) {
	js.Rewrite("$<.deltaMode")
	return deltaMode
}

func (*WheelEvent) GetDeltaX() (deltaX int) {
	js.Rewrite("$<.deltaX")
	return deltaX
}

func (*WheelEvent) GetDeltaY() (deltaY int) {
	js.Rewrite("$<.deltaY")
	return deltaY
}

func (*WheelEvent) GetDeltaZ() (deltaZ int) {
	js.Rewrite("$<.deltaZ")
	return deltaZ
}

func (*WheelEvent) GetWheelDelta() (wheelDelta int) {
	js.Rewrite("$<.wheelDelta")
	return wheelDelta
}

func (*WheelEvent) GetWheelDeltaX() (wheelDeltaX int) {
	js.Rewrite("$<.wheelDeltaX")
	return wheelDeltaX
}

func (*WheelEvent) GetWheelDeltaY() (wheelDeltaY int) {
	js.Rewrite("$<.wheelDeltaY")
	return wheelDeltaY
}

type Window struct {
	*EventTarget
	*WindowTimers
	*WindowSessionStorage
	*WindowLocalStorage
	*WindowConsole
	*GlobalEventHandlers
	*IDBEnvironment
	*WindowBase64
	*GlobalFetch
}

func (*Window) Alert(message string) {
	js.Rewrite("$<.alert($1)", message)
}

func (*Window) Blur() {
	js.Rewrite("$<.blur()")
}

func (*Window) CancelAnimationFrame(handle int) {
	js.Rewrite("$<.cancelAnimationFrame($1)", handle)
}

func (*Window) CaptureEvents() {
	js.Rewrite("$<.captureEvents()")
}

func (*Window) Close() {
	js.Rewrite("$<.close()")
}

func (*Window) Confirm(message string) (b bool) {
	js.Rewrite("$<.confirm($1)", message)
	return b
}

func (*Window) DepartFocus(navigationReason *NavigationReason, origin *FocusNavigationOrigin) {
	js.Rewrite("$<.departFocus($1, $2)", navigationReason, origin)
}

func (*Window) Focus() {
	js.Rewrite("$<.focus()")
}

func (*Window) GetComputedStyle(elt *Element, pseudoElt string) (c *CSSStyleDeclaration) {
	js.Rewrite("$<.getComputedStyle($1, $2)", elt, pseudoElt)
	return c
}

func (*Window) GetMatchedCSSRules(elt *Element, pseudoElt string) (c *CSSRuleList) {
	js.Rewrite("$<.getMatchedCSSRules($1, $2)", elt, pseudoElt)
	return c
}

func (*Window) GetSelection() (s *Selection) {
	js.Rewrite("$<.getSelection()")
	return s
}

func (*Window) MatchMedia(mediaQuery string) (m *MediaQueryList) {
	js.Rewrite("$<.matchMedia($1)", mediaQuery)
	return m
}

func (*Window) MoveBy(x int, y int) {
	js.Rewrite("$<.moveBy($1, $2)", x, y)
}

func (*Window) MoveTo(x int, y int) {
	js.Rewrite("$<.moveTo($1, $2)", x, y)
}

func (*Window) MsWriteProfilerMark(profilerMarkName string) {
	js.Rewrite("$<.msWriteProfilerMark($1)", profilerMarkName)
}

func (*Window) Open(url string, target string, features string, replace bool) (w *Window) {
	js.Rewrite("$<.open($1, $2, $3, $4)", url, target, features, replace)
	return w
}

func (*Window) PostMessage(message interface{}, targetOrigin string, transfer []interface{}) {
	js.Rewrite("$<.postMessage($1, $2, $3)", message, targetOrigin, transfer)
}

func (*Window) Print() {
	js.Rewrite("$<.print()")
}

func (*Window) Prompt(message string, def string) (s string) {
	js.Rewrite("$<.prompt($1, $2)", message, def)
	return s
}

func (*Window) ReleaseEvents() {
	js.Rewrite("$<.releaseEvents()")
}

func (*Window) RequestAnimationFrame(callback func(time int)) (i int) {
	js.Rewrite("$<.requestAnimationFrame($1)", callback)
	return i
}

func (*Window) ResizeBy(x int, y int) {
	js.Rewrite("$<.resizeBy($1, $2)", x, y)
}

func (*Window) ResizeTo(x int, y int) {
	js.Rewrite("$<.resizeTo($1, $2)", x, y)
}

func (*Window) Scroll(x int, y int) {
	js.Rewrite("$<.scroll($1, $2)", x, y)
}

func (*Window) ScrollBy(x int, y int) {
	js.Rewrite("$<.scrollBy($1, $2)", x, y)
}

func (*Window) ScrollTo(x int, y int) {
	js.Rewrite("$<.scrollTo($1, $2)", x, y)
}

func (*Window) Stop() {
	js.Rewrite("$<.stop()")
}

func (*Window) WebkitCancelAnimationFrame(handle int) {
	js.Rewrite("$<.webkitCancelAnimationFrame($1)", handle)
}

func (*Window) WebkitConvertPointFromNodeToPage(node *Node, pt *WebKitPoint) (w *WebKitPoint) {
	js.Rewrite("$<.webkitConvertPointFromNodeToPage($1, $2)", node, pt)
	return w
}

func (*Window) WebkitConvertPointFromPageToNode(node *Node, pt *WebKitPoint) (w *WebKitPoint) {
	js.Rewrite("$<.webkitConvertPointFromPageToNode($1, $2)", node, pt)
	return w
}

func (*Window) WebkitRequestAnimationFrame(callback func(time int)) (i int) {
	js.Rewrite("$<.webkitRequestAnimationFrame($1)", callback)
	return i
}

func (*Window) GetApplicationCache() (applicationCache *ApplicationCache) {
	js.Rewrite("$<.applicationCache")
	return applicationCache
}

func (*Window) GetCaches() (caches *CacheStorage) {
	js.Rewrite("$<.caches")
	return caches
}

func (*Window) GetClientInformation() (clientInformation *Navigator) {
	js.Rewrite("$<.clientInformation")
	return clientInformation
}

func (*Window) GetClosed() (closed bool) {
	js.Rewrite("$<.closed")
	return closed
}

func (*Window) GetCrypto() (crypto *Crypto) {
	js.Rewrite("$<.crypto")
	return crypto
}

func (*Window) GetDefaultStatus() (defaultStatus string) {
	js.Rewrite("$<.defaultStatus")
	return defaultStatus
}

func (*Window) SetDefaultStatus(defaultStatus string) {
	js.Rewrite("$<.defaultStatus = $1", defaultStatus)
}

func (*Window) GetDevicePixelRatio() (devicePixelRatio float32) {
	js.Rewrite("$<.devicePixelRatio")
	return devicePixelRatio
}

func (*Window) GetDocument() (document *Document) {
	js.Rewrite("$<.document")
	return document
}

func (*Window) GetDoNotTrack() (doNotTrack string) {
	js.Rewrite("$<.doNotTrack")
	return doNotTrack
}

func (*Window) GetEvent() (event *Event) {
	js.Rewrite("$<.event")
	return event
}

func (*Window) SetEvent(event *Event) {
	js.Rewrite("$<.event = $1", event)
}

func (*Window) GetExternal() (external *External) {
	js.Rewrite("$<.external")
	return external
}

func (*Window) GetFrameElement() (frameElement *Element) {
	js.Rewrite("$<.frameElement")
	return frameElement
}

func (*Window) GetFrames() (frames *Window) {
	js.Rewrite("$<.frames")
	return frames
}

func (*Window) GetHistory() (history *History) {
	js.Rewrite("$<.history")
	return history
}

func (*Window) GetInnerHeight() (innerHeight int) {
	js.Rewrite("$<.innerHeight")
	return innerHeight
}

func (*Window) GetInnerWidth() (innerWidth int) {
	js.Rewrite("$<.innerWidth")
	return innerWidth
}

func (*Window) GetIsSecureContext() (isSecureContext bool) {
	js.Rewrite("$<.isSecureContext")
	return isSecureContext
}

func (*Window) GetLength() (length uint) {
	js.Rewrite("$<.length")
	return length
}

func (*Window) GetLocation() (location *Location) {
	js.Rewrite("$<.location")
	return location
}

func (*Window) GetLocationbar() (locationbar *BarProp) {
	js.Rewrite("$<.locationbar")
	return locationbar
}

func (*Window) GetMenubar() (menubar *BarProp) {
	js.Rewrite("$<.menubar")
	return menubar
}

func (*Window) GetMsContentScript() (msContentScript *ExtensionScriptApis) {
	js.Rewrite("$<.msContentScript")
	return msContentScript
}

func (*Window) GetMsCredentials() (msCredentials *MSCredentials) {
	js.Rewrite("$<.msCredentials")
	return msCredentials
}

func (*Window) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*Window) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*Window) GetNavigator() (navigator *Navigator) {
	js.Rewrite("$<.navigator")
	return navigator
}

func (*Window) GetOffscreenBuffering() (offscreenBuffering interface{}) {
	js.Rewrite("$<.offscreenBuffering")
	return offscreenBuffering
}

func (*Window) SetOffscreenBuffering(offscreenBuffering interface{}) {
	js.Rewrite("$<.offscreenBuffering = $1", offscreenBuffering)
}

func (*Window) GetOnabort() (e *Event) {
	js.Rewrite("$<.onabort")
	return e
}

func (*Window) SetOnabort(e *Event) {
	js.Rewrite("$<.onabort = $1", e)
}

func (*Window) GetOnafterprint() (afterprint *Event) {
	js.Rewrite("$<.onafterprint")
	return afterprint
}

func (*Window) SetOnafterprint(afterprint *Event) {
	js.Rewrite("$<.onafterprint = $1", afterprint)
}

func (*Window) GetOnbeforeprint() (beforeprint *Event) {
	js.Rewrite("$<.onbeforeprint")
	return beforeprint
}

func (*Window) SetOnbeforeprint(beforeprint *Event) {
	js.Rewrite("$<.onbeforeprint = $1", beforeprint)
}

func (*Window) GetOnbeforeunload() (beforeunload *BeforeUnloadEvent) {
	js.Rewrite("$<.onbeforeunload")
	return beforeunload
}

func (*Window) SetOnbeforeunload(beforeunload *BeforeUnloadEvent) {
	js.Rewrite("$<.onbeforeunload = $1", beforeunload)
}

func (*Window) GetOnblur() (blur *FocusEvent) {
	js.Rewrite("$<.onblur")
	return blur
}

func (*Window) SetOnblur(blur *FocusEvent) {
	js.Rewrite("$<.onblur = $1", blur)
}

func (*Window) GetOncanplay() (e *Event) {
	js.Rewrite("$<.oncanplay")
	return e
}

func (*Window) SetOncanplay(e *Event) {
	js.Rewrite("$<.oncanplay = $1", e)
}

func (*Window) GetOncanplaythrough() (e *Event) {
	js.Rewrite("$<.oncanplaythrough")
	return e
}

func (*Window) SetOncanplaythrough(e *Event) {
	js.Rewrite("$<.oncanplaythrough = $1", e)
}

func (*Window) GetOnchange() (e *Event) {
	js.Rewrite("$<.onchange")
	return e
}

func (*Window) SetOnchange(e *Event) {
	js.Rewrite("$<.onchange = $1", e)
}

func (*Window) GetOnclick() (e *Event) {
	js.Rewrite("$<.onclick")
	return e
}

func (*Window) SetOnclick(e *Event) {
	js.Rewrite("$<.onclick = $1", e)
}

func (*Window) GetOncompassneedscalibration() (compassneedscalibration *Event) {
	js.Rewrite("$<.oncompassneedscalibration")
	return compassneedscalibration
}

func (*Window) SetOncompassneedscalibration(compassneedscalibration *Event) {
	js.Rewrite("$<.oncompassneedscalibration = $1", compassneedscalibration)
}

func (*Window) GetOncontextmenu() (e *Event) {
	js.Rewrite("$<.oncontextmenu")
	return e
}

func (*Window) SetOncontextmenu(e *Event) {
	js.Rewrite("$<.oncontextmenu = $1", e)
}

func (*Window) GetOndblclick() (e *Event) {
	js.Rewrite("$<.ondblclick")
	return e
}

func (*Window) SetOndblclick(e *Event) {
	js.Rewrite("$<.ondblclick = $1", e)
}

func (*Window) GetOndevicelight() (devicelight *DeviceLightEvent) {
	js.Rewrite("$<.ondevicelight")
	return devicelight
}

func (*Window) SetOndevicelight(devicelight *DeviceLightEvent) {
	js.Rewrite("$<.ondevicelight = $1", devicelight)
}

func (*Window) GetOndevicemotion() (devicemotion *DeviceMotionEvent) {
	js.Rewrite("$<.ondevicemotion")
	return devicemotion
}

func (*Window) SetOndevicemotion(devicemotion *DeviceMotionEvent) {
	js.Rewrite("$<.ondevicemotion = $1", devicemotion)
}

func (*Window) GetOndeviceorientation() (deviceorientation *DeviceOrientationEvent) {
	js.Rewrite("$<.ondeviceorientation")
	return deviceorientation
}

func (*Window) SetOndeviceorientation(deviceorientation *DeviceOrientationEvent) {
	js.Rewrite("$<.ondeviceorientation = $1", deviceorientation)
}

func (*Window) GetOndrag() (e *Event) {
	js.Rewrite("$<.ondrag")
	return e
}

func (*Window) SetOndrag(e *Event) {
	js.Rewrite("$<.ondrag = $1", e)
}

func (*Window) GetOndragend() (e *Event) {
	js.Rewrite("$<.ondragend")
	return e
}

func (*Window) SetOndragend(e *Event) {
	js.Rewrite("$<.ondragend = $1", e)
}

func (*Window) GetOndragenter() (e *Event) {
	js.Rewrite("$<.ondragenter")
	return e
}

func (*Window) SetOndragenter(e *Event) {
	js.Rewrite("$<.ondragenter = $1", e)
}

func (*Window) GetOndragleave() (e *Event) {
	js.Rewrite("$<.ondragleave")
	return e
}

func (*Window) SetOndragleave(e *Event) {
	js.Rewrite("$<.ondragleave = $1", e)
}

func (*Window) GetOndragover() (e *Event) {
	js.Rewrite("$<.ondragover")
	return e
}

func (*Window) SetOndragover(e *Event) {
	js.Rewrite("$<.ondragover = $1", e)
}

func (*Window) GetOndragstart() (e *Event) {
	js.Rewrite("$<.ondragstart")
	return e
}

func (*Window) SetOndragstart(e *Event) {
	js.Rewrite("$<.ondragstart = $1", e)
}

func (*Window) GetOndrop() (e *Event) {
	js.Rewrite("$<.ondrop")
	return e
}

func (*Window) SetOndrop(e *Event) {
	js.Rewrite("$<.ondrop = $1", e)
}

func (*Window) GetOndurationchange() (e *Event) {
	js.Rewrite("$<.ondurationchange")
	return e
}

func (*Window) SetOndurationchange(e *Event) {
	js.Rewrite("$<.ondurationchange = $1", e)
}

func (*Window) GetOnemptied() (e *Event) {
	js.Rewrite("$<.onemptied")
	return e
}

func (*Window) SetOnemptied(e *Event) {
	js.Rewrite("$<.onemptied = $1", e)
}

func (*Window) GetOnended() (e *Event) {
	js.Rewrite("$<.onended")
	return e
}

func (*Window) SetOnended(e *Event) {
	js.Rewrite("$<.onended = $1", e)
}

func (*Window) GetOnerror() (ErrorEventHandler func(columnNumber uint, event interface{}, fileno uint, source string)) {
	js.Rewrite("$<.onerror")
	return ErrorEventHandler
}

func (*Window) SetOnerror(ErrorEventHandler func(columnNumber uint, event interface{}, fileno uint, source string)) {
	js.Rewrite("$<.onerror = $1", ErrorEventHandler)
}

func (*Window) GetOnfocus() (focus *FocusEvent) {
	js.Rewrite("$<.onfocus")
	return focus
}

func (*Window) SetOnfocus(focus *FocusEvent) {
	js.Rewrite("$<.onfocus = $1", focus)
}

func (*Window) GetOnhashchange() (hashchange *HashChangeEvent) {
	js.Rewrite("$<.onhashchange")
	return hashchange
}

func (*Window) SetOnhashchange(hashchange *HashChangeEvent) {
	js.Rewrite("$<.onhashchange = $1", hashchange)
}

func (*Window) GetOninput() (e *Event) {
	js.Rewrite("$<.oninput")
	return e
}

func (*Window) SetOninput(e *Event) {
	js.Rewrite("$<.oninput = $1", e)
}

func (*Window) GetOninvalid() (e *Event) {
	js.Rewrite("$<.oninvalid")
	return e
}

func (*Window) SetOninvalid(e *Event) {
	js.Rewrite("$<.oninvalid = $1", e)
}

func (*Window) GetOnkeydown() (e *Event) {
	js.Rewrite("$<.onkeydown")
	return e
}

func (*Window) SetOnkeydown(e *Event) {
	js.Rewrite("$<.onkeydown = $1", e)
}

func (*Window) GetOnkeypress() (e *Event) {
	js.Rewrite("$<.onkeypress")
	return e
}

func (*Window) SetOnkeypress(e *Event) {
	js.Rewrite("$<.onkeypress = $1", e)
}

func (*Window) GetOnkeyup() (e *Event) {
	js.Rewrite("$<.onkeyup")
	return e
}

func (*Window) SetOnkeyup(e *Event) {
	js.Rewrite("$<.onkeyup = $1", e)
}

func (*Window) GetOnload() (load *Event) {
	js.Rewrite("$<.onload")
	return load
}

func (*Window) SetOnload(load *Event) {
	js.Rewrite("$<.onload = $1", load)
}

func (*Window) GetOnloadeddata() (e *Event) {
	js.Rewrite("$<.onloadeddata")
	return e
}

func (*Window) SetOnloadeddata(e *Event) {
	js.Rewrite("$<.onloadeddata = $1", e)
}

func (*Window) GetOnloadedmetadata() (e *Event) {
	js.Rewrite("$<.onloadedmetadata")
	return e
}

func (*Window) SetOnloadedmetadata(e *Event) {
	js.Rewrite("$<.onloadedmetadata = $1", e)
}

func (*Window) GetOnloadstart() (e *Event) {
	js.Rewrite("$<.onloadstart")
	return e
}

func (*Window) SetOnloadstart(e *Event) {
	js.Rewrite("$<.onloadstart = $1", e)
}

func (*Window) GetOnmessage() (message *MessageEvent) {
	js.Rewrite("$<.onmessage")
	return message
}

func (*Window) SetOnmessage(message *MessageEvent) {
	js.Rewrite("$<.onmessage = $1", message)
}

func (*Window) GetOnmousedown() (e *Event) {
	js.Rewrite("$<.onmousedown")
	return e
}

func (*Window) SetOnmousedown(e *Event) {
	js.Rewrite("$<.onmousedown = $1", e)
}

func (*Window) GetOnmouseenter() (e *Event) {
	js.Rewrite("$<.onmouseenter")
	return e
}

func (*Window) SetOnmouseenter(e *Event) {
	js.Rewrite("$<.onmouseenter = $1", e)
}

func (*Window) GetOnmouseleave() (e *Event) {
	js.Rewrite("$<.onmouseleave")
	return e
}

func (*Window) SetOnmouseleave(e *Event) {
	js.Rewrite("$<.onmouseleave = $1", e)
}

func (*Window) GetOnmousemove() (e *Event) {
	js.Rewrite("$<.onmousemove")
	return e
}

func (*Window) SetOnmousemove(e *Event) {
	js.Rewrite("$<.onmousemove = $1", e)
}

func (*Window) GetOnmouseout() (e *Event) {
	js.Rewrite("$<.onmouseout")
	return e
}

func (*Window) SetOnmouseout(e *Event) {
	js.Rewrite("$<.onmouseout = $1", e)
}

func (*Window) GetOnmouseover() (e *Event) {
	js.Rewrite("$<.onmouseover")
	return e
}

func (*Window) SetOnmouseover(e *Event) {
	js.Rewrite("$<.onmouseover = $1", e)
}

func (*Window) GetOnmouseup() (e *Event) {
	js.Rewrite("$<.onmouseup")
	return e
}

func (*Window) SetOnmouseup(e *Event) {
	js.Rewrite("$<.onmouseup = $1", e)
}

func (*Window) GetOnmousewheel() (e *Event) {
	js.Rewrite("$<.onmousewheel")
	return e
}

func (*Window) SetOnmousewheel(e *Event) {
	js.Rewrite("$<.onmousewheel = $1", e)
}

func (*Window) GetOnmsgesturechange() (e *Event) {
	js.Rewrite("$<.onmsgesturechange")
	return e
}

func (*Window) SetOnmsgesturechange(e *Event) {
	js.Rewrite("$<.onmsgesturechange = $1", e)
}

func (*Window) GetOnmsgesturedoubletap() (e *Event) {
	js.Rewrite("$<.onmsgesturedoubletap")
	return e
}

func (*Window) SetOnmsgesturedoubletap(e *Event) {
	js.Rewrite("$<.onmsgesturedoubletap = $1", e)
}

func (*Window) GetOnmsgestureend() (e *Event) {
	js.Rewrite("$<.onmsgestureend")
	return e
}

func (*Window) SetOnmsgestureend(e *Event) {
	js.Rewrite("$<.onmsgestureend = $1", e)
}

func (*Window) GetOnmsgesturehold() (e *Event) {
	js.Rewrite("$<.onmsgesturehold")
	return e
}

func (*Window) SetOnmsgesturehold(e *Event) {
	js.Rewrite("$<.onmsgesturehold = $1", e)
}

func (*Window) GetOnmsgesturestart() (e *Event) {
	js.Rewrite("$<.onmsgesturestart")
	return e
}

func (*Window) SetOnmsgesturestart(e *Event) {
	js.Rewrite("$<.onmsgesturestart = $1", e)
}

func (*Window) GetOnmsgesturetap() (e *Event) {
	js.Rewrite("$<.onmsgesturetap")
	return e
}

func (*Window) SetOnmsgesturetap(e *Event) {
	js.Rewrite("$<.onmsgesturetap = $1", e)
}

func (*Window) GetOnmsinertiastart() (e *Event) {
	js.Rewrite("$<.onmsinertiastart")
	return e
}

func (*Window) SetOnmsinertiastart(e *Event) {
	js.Rewrite("$<.onmsinertiastart = $1", e)
}

func (*Window) GetOnmspointercancel() (e *Event) {
	js.Rewrite("$<.onmspointercancel")
	return e
}

func (*Window) SetOnmspointercancel(e *Event) {
	js.Rewrite("$<.onmspointercancel = $1", e)
}

func (*Window) GetOnmspointerdown() (e *Event) {
	js.Rewrite("$<.onmspointerdown")
	return e
}

func (*Window) SetOnmspointerdown(e *Event) {
	js.Rewrite("$<.onmspointerdown = $1", e)
}

func (*Window) GetOnmspointerenter() (e *Event) {
	js.Rewrite("$<.onmspointerenter")
	return e
}

func (*Window) SetOnmspointerenter(e *Event) {
	js.Rewrite("$<.onmspointerenter = $1", e)
}

func (*Window) GetOnmspointerleave() (e *Event) {
	js.Rewrite("$<.onmspointerleave")
	return e
}

func (*Window) SetOnmspointerleave(e *Event) {
	js.Rewrite("$<.onmspointerleave = $1", e)
}

func (*Window) GetOnmspointermove() (e *Event) {
	js.Rewrite("$<.onmspointermove")
	return e
}

func (*Window) SetOnmspointermove(e *Event) {
	js.Rewrite("$<.onmspointermove = $1", e)
}

func (*Window) GetOnmspointerout() (e *Event) {
	js.Rewrite("$<.onmspointerout")
	return e
}

func (*Window) SetOnmspointerout(e *Event) {
	js.Rewrite("$<.onmspointerout = $1", e)
}

func (*Window) GetOnmspointerover() (e *Event) {
	js.Rewrite("$<.onmspointerover")
	return e
}

func (*Window) SetOnmspointerover(e *Event) {
	js.Rewrite("$<.onmspointerover = $1", e)
}

func (*Window) GetOnmspointerup() (e *Event) {
	js.Rewrite("$<.onmspointerup")
	return e
}

func (*Window) SetOnmspointerup(e *Event) {
	js.Rewrite("$<.onmspointerup = $1", e)
}

func (*Window) GetOnoffline() (e *Event) {
	js.Rewrite("$<.onoffline")
	return e
}

func (*Window) SetOnoffline(e *Event) {
	js.Rewrite("$<.onoffline = $1", e)
}

func (*Window) GetOnonline() (e *Event) {
	js.Rewrite("$<.ononline")
	return e
}

func (*Window) SetOnonline(e *Event) {
	js.Rewrite("$<.ononline = $1", e)
}

func (*Window) GetOnorientationchange() (orientationchange *Event) {
	js.Rewrite("$<.onorientationchange")
	return orientationchange
}

func (*Window) SetOnorientationchange(orientationchange *Event) {
	js.Rewrite("$<.onorientationchange = $1", orientationchange)
}

func (*Window) GetOnpagehide() (pagehide *PageTransitionEvent) {
	js.Rewrite("$<.onpagehide")
	return pagehide
}

func (*Window) SetOnpagehide(pagehide *PageTransitionEvent) {
	js.Rewrite("$<.onpagehide = $1", pagehide)
}

func (*Window) GetOnpageshow() (pageshow *PageTransitionEvent) {
	js.Rewrite("$<.onpageshow")
	return pageshow
}

func (*Window) SetOnpageshow(pageshow *PageTransitionEvent) {
	js.Rewrite("$<.onpageshow = $1", pageshow)
}

func (*Window) GetOnpause() (e *Event) {
	js.Rewrite("$<.onpause")
	return e
}

func (*Window) SetOnpause(e *Event) {
	js.Rewrite("$<.onpause = $1", e)
}

func (*Window) GetOnplay() (e *Event) {
	js.Rewrite("$<.onplay")
	return e
}

func (*Window) SetOnplay(e *Event) {
	js.Rewrite("$<.onplay = $1", e)
}

func (*Window) GetOnplaying() (e *Event) {
	js.Rewrite("$<.onplaying")
	return e
}

func (*Window) SetOnplaying(e *Event) {
	js.Rewrite("$<.onplaying = $1", e)
}

func (*Window) GetOnpopstate() (popstate *PopStateEvent) {
	js.Rewrite("$<.onpopstate")
	return popstate
}

func (*Window) SetOnpopstate(popstate *PopStateEvent) {
	js.Rewrite("$<.onpopstate = $1", popstate)
}

func (*Window) GetOnprogress() (e *Event) {
	js.Rewrite("$<.onprogress")
	return e
}

func (*Window) SetOnprogress(e *Event) {
	js.Rewrite("$<.onprogress = $1", e)
}

func (*Window) GetOnratechange() (e *Event) {
	js.Rewrite("$<.onratechange")
	return e
}

func (*Window) SetOnratechange(e *Event) {
	js.Rewrite("$<.onratechange = $1", e)
}

func (*Window) GetOnreadystatechange() (e *Event) {
	js.Rewrite("$<.onreadystatechange")
	return e
}

func (*Window) SetOnreadystatechange(e *Event) {
	js.Rewrite("$<.onreadystatechange = $1", e)
}

func (*Window) GetOnreset() (e *Event) {
	js.Rewrite("$<.onreset")
	return e
}

func (*Window) SetOnreset(e *Event) {
	js.Rewrite("$<.onreset = $1", e)
}

func (*Window) GetOnresize() (resize *UIEvent) {
	js.Rewrite("$<.onresize")
	return resize
}

func (*Window) SetOnresize(resize *UIEvent) {
	js.Rewrite("$<.onresize = $1", resize)
}

func (*Window) GetOnscroll() (e *Event) {
	js.Rewrite("$<.onscroll")
	return e
}

func (*Window) SetOnscroll(e *Event) {
	js.Rewrite("$<.onscroll = $1", e)
}

func (*Window) GetOnseeked() (e *Event) {
	js.Rewrite("$<.onseeked")
	return e
}

func (*Window) SetOnseeked(e *Event) {
	js.Rewrite("$<.onseeked = $1", e)
}

func (*Window) GetOnseeking() (e *Event) {
	js.Rewrite("$<.onseeking")
	return e
}

func (*Window) SetOnseeking(e *Event) {
	js.Rewrite("$<.onseeking = $1", e)
}

func (*Window) GetOnselect() (e *Event) {
	js.Rewrite("$<.onselect")
	return e
}

func (*Window) SetOnselect(e *Event) {
	js.Rewrite("$<.onselect = $1", e)
}

func (*Window) GetOnstalled() (e *Event) {
	js.Rewrite("$<.onstalled")
	return e
}

func (*Window) SetOnstalled(e *Event) {
	js.Rewrite("$<.onstalled = $1", e)
}

func (*Window) GetOnstorage() (storage *StorageEvent) {
	js.Rewrite("$<.onstorage")
	return storage
}

func (*Window) SetOnstorage(storage *StorageEvent) {
	js.Rewrite("$<.onstorage = $1", storage)
}

func (*Window) GetOnsubmit() (e *Event) {
	js.Rewrite("$<.onsubmit")
	return e
}

func (*Window) SetOnsubmit(e *Event) {
	js.Rewrite("$<.onsubmit = $1", e)
}

func (*Window) GetOnsuspend() (e *Event) {
	js.Rewrite("$<.onsuspend")
	return e
}

func (*Window) SetOnsuspend(e *Event) {
	js.Rewrite("$<.onsuspend = $1", e)
}

func (*Window) GetOntimeupdate() (e *Event) {
	js.Rewrite("$<.ontimeupdate")
	return e
}

func (*Window) SetOntimeupdate(e *Event) {
	js.Rewrite("$<.ontimeupdate = $1", e)
}

func (*Window) GetOntouchcancel() (ontouchcancel interface{}) {
	js.Rewrite("$<.ontouchcancel")
	return ontouchcancel
}

func (*Window) SetOntouchcancel(ontouchcancel interface{}) {
	js.Rewrite("$<.ontouchcancel = $1", ontouchcancel)
}

func (*Window) GetOntouchend() (ontouchend interface{}) {
	js.Rewrite("$<.ontouchend")
	return ontouchend
}

func (*Window) SetOntouchend(ontouchend interface{}) {
	js.Rewrite("$<.ontouchend = $1", ontouchend)
}

func (*Window) GetOntouchmove() (ontouchmove interface{}) {
	js.Rewrite("$<.ontouchmove")
	return ontouchmove
}

func (*Window) SetOntouchmove(ontouchmove interface{}) {
	js.Rewrite("$<.ontouchmove = $1", ontouchmove)
}

func (*Window) GetOntouchstart() (ontouchstart interface{}) {
	js.Rewrite("$<.ontouchstart")
	return ontouchstart
}

func (*Window) SetOntouchstart(ontouchstart interface{}) {
	js.Rewrite("$<.ontouchstart = $1", ontouchstart)
}

func (*Window) GetOnunload() (unload *Event) {
	js.Rewrite("$<.onunload")
	return unload
}

func (*Window) SetOnunload(unload *Event) {
	js.Rewrite("$<.onunload = $1", unload)
}

func (*Window) GetOnvolumechange() (e *Event) {
	js.Rewrite("$<.onvolumechange")
	return e
}

func (*Window) SetOnvolumechange(e *Event) {
	js.Rewrite("$<.onvolumechange = $1", e)
}

func (*Window) GetOnwaiting() (e *Event) {
	js.Rewrite("$<.onwaiting")
	return e
}

func (*Window) SetOnwaiting(e *Event) {
	js.Rewrite("$<.onwaiting = $1", e)
}

func (*Window) GetOpener() (opener *Window) {
	js.Rewrite("$<.opener")
	return opener
}

func (*Window) GetOrientation() (orientation string) {
	js.Rewrite("$<.orientation")
	return orientation
}

func (*Window) GetOuterHeight() (outerHeight int) {
	js.Rewrite("$<.outerHeight")
	return outerHeight
}

func (*Window) GetOuterWidth() (outerWidth int) {
	js.Rewrite("$<.outerWidth")
	return outerWidth
}

func (*Window) GetPageXOffset() (pageXOffset int) {
	js.Rewrite("$<.pageXOffset")
	return pageXOffset
}

func (*Window) GetPageYOffset() (pageYOffset int) {
	js.Rewrite("$<.pageYOffset")
	return pageYOffset
}

func (*Window) GetParent() (parent *Window) {
	js.Rewrite("$<.parent")
	return parent
}

func (*Window) GetPerformance() (performance *Performance) {
	js.Rewrite("$<.performance")
	return performance
}

func (*Window) GetPersonalbar() (personalbar *BarProp) {
	js.Rewrite("$<.personalbar")
	return personalbar
}

func (*Window) GetScreen() (screen *Screen) {
	js.Rewrite("$<.screen")
	return screen
}

func (*Window) GetScreenLeft() (screenLeft int) {
	js.Rewrite("$<.screenLeft")
	return screenLeft
}

func (*Window) GetScreenTop() (screenTop int) {
	js.Rewrite("$<.screenTop")
	return screenTop
}

func (*Window) GetScreenX() (screenX int) {
	js.Rewrite("$<.screenX")
	return screenX
}

func (*Window) GetScreenY() (screenY int) {
	js.Rewrite("$<.screenY")
	return screenY
}

func (*Window) GetScrollbars() (scrollbars *BarProp) {
	js.Rewrite("$<.scrollbars")
	return scrollbars
}

func (*Window) GetScrollX() (scrollX int) {
	js.Rewrite("$<.scrollX")
	return scrollX
}

func (*Window) GetScrollY() (scrollY int) {
	js.Rewrite("$<.scrollY")
	return scrollY
}

func (*Window) GetSelf() (self *Window) {
	js.Rewrite("$<.self")
	return self
}

func (*Window) GetSpeechSynthesis() (speechSynthesis *SpeechSynthesis) {
	js.Rewrite("$<.speechSynthesis")
	return speechSynthesis
}

func (*Window) GetStatus() (status string) {
	js.Rewrite("$<.status")
	return status
}

func (*Window) SetStatus(status string) {
	js.Rewrite("$<.status = $1", status)
}

func (*Window) GetStatusbar() (statusbar *BarProp) {
	js.Rewrite("$<.statusbar")
	return statusbar
}

func (*Window) GetStyleMedia() (styleMedia *StyleMedia) {
	js.Rewrite("$<.styleMedia")
	return styleMedia
}

func (*Window) GetToolbar() (toolbar *BarProp) {
	js.Rewrite("$<.toolbar")
	return toolbar
}

func (*Window) GetTop() (top *Window) {
	js.Rewrite("$<.top")
	return top
}

func (*Window) GetWindow() (window *Window) {
	js.Rewrite("$<.window")
	return window
}

type Worker struct {
	*EventTarget
	*AbstractWorker
}

func (*Worker) PostMessage(message interface{}, transfer []interface{}) {
	js.Rewrite("$<.postMessage($1, $2)", message, transfer)
}

func (*Worker) Terminate() {
	js.Rewrite("$<.terminate()")
}

func (*Worker) GetOnmessage() (message *MessageEvent) {
	js.Rewrite("$<.onmessage")
	return message
}

func (*Worker) SetOnmessage(message *MessageEvent) {
	js.Rewrite("$<.onmessage = $1", message)
}

type XMLDocument struct {
	*Document
}

type XMLHttpRequest struct {
	*EventTarget
	*XMLHttpRequestEventTarget
}

func (*XMLHttpRequest) Abort() {
	js.Rewrite("$<.abort()")
}

func (*XMLHttpRequest) GetAllResponseHeaders() (s string) {
	js.Rewrite("$<.getAllResponseHeaders()")
	return s
}

func (*XMLHttpRequest) GetResponseHeader(header string) (s string) {
	js.Rewrite("$<.getResponseHeader($1)", header)
	return s
}

func (*XMLHttpRequest) MsCachingEnabled() (b bool) {
	js.Rewrite("$<.msCachingEnabled()")
	return b
}

func (*XMLHttpRequest) Open(method string, url string, async bool, user string, password string) {
	js.Rewrite("$<.open($1, $2, $3, $4, $5)", method, url, async, user, password)
}

func (*XMLHttpRequest) OverrideMimeType(mime string) {
	js.Rewrite("$<.overrideMimeType($1)", mime)
}

func (*XMLHttpRequest) Send(data interface{}) {
	js.Rewrite("$<.send($1)", data)
}

func (*XMLHttpRequest) SetRequestHeader(header string, value string) {
	js.Rewrite("$<.setRequestHeader($1, $2)", header, value)
}

func (*XMLHttpRequest) GetMsCaching() (msCaching string) {
	js.Rewrite("$<.msCaching")
	return msCaching
}

func (*XMLHttpRequest) SetMsCaching(msCaching string) {
	js.Rewrite("$<.msCaching = $1", msCaching)
}

func (*XMLHttpRequest) GetOnreadystatechange() (readystatechange *Event) {
	js.Rewrite("$<.onreadystatechange")
	return readystatechange
}

func (*XMLHttpRequest) SetOnreadystatechange(readystatechange *Event) {
	js.Rewrite("$<.onreadystatechange = $1", readystatechange)
}

func (*XMLHttpRequest) GetReadyState() (readyState uint8) {
	js.Rewrite("$<.readyState")
	return readyState
}

func (*XMLHttpRequest) GetResponse() (response interface{}) {
	js.Rewrite("$<.response")
	return response
}

func (*XMLHttpRequest) GetResponseText() (responseText string) {
	js.Rewrite("$<.responseText")
	return responseText
}

func (*XMLHttpRequest) GetResponseType() (responseType *XMLHttpRequestResponseType) {
	js.Rewrite("$<.responseType")
	return responseType
}

func (*XMLHttpRequest) SetResponseType(responseType *XMLHttpRequestResponseType) {
	js.Rewrite("$<.responseType = $1", responseType)
}

func (*XMLHttpRequest) GetResponseURL() (responseURL string) {
	js.Rewrite("$<.responseURL")
	return responseURL
}

func (*XMLHttpRequest) GetResponseXML() (responseXML *Document) {
	js.Rewrite("$<.responseXML")
	return responseXML
}

func (*XMLHttpRequest) GetStatus() (status uint8) {
	js.Rewrite("$<.status")
	return status
}

func (*XMLHttpRequest) GetStatusText() (statusText string) {
	js.Rewrite("$<.statusText")
	return statusText
}

func (*XMLHttpRequest) GetTimeout() (timeout uint) {
	js.Rewrite("$<.timeout")
	return timeout
}

func (*XMLHttpRequest) SetTimeout(timeout uint) {
	js.Rewrite("$<.timeout = $1", timeout)
}

func (*XMLHttpRequest) GetUpload() (upload *XMLHttpRequestUpload) {
	js.Rewrite("$<.upload")
	return upload
}

func (*XMLHttpRequest) GetWithCredentials() (withCredentials bool) {
	js.Rewrite("$<.withCredentials")
	return withCredentials
}

func (*XMLHttpRequest) SetWithCredentials(withCredentials bool) {
	js.Rewrite("$<.withCredentials = $1", withCredentials)
}

type XMLHttpRequestUpload struct {
	*EventTarget
	*XMLHttpRequestEventTarget
}

type XMLSerializer struct {
}

func (*XMLSerializer) SerializeToString(target *Node) (s string) {
	js.Rewrite("$<.serializeToString($1)", target)
	return s
}

type XPathEvaluator struct {
}

func (*XPathEvaluator) CreateExpression(expression string, resolver *XPathNSResolver) (x *XPathExpression) {
	js.Rewrite("$<.createExpression($1, $2)", expression, resolver)
	return x
}

func (*XPathEvaluator) CreateNSResolver(nodeResolver *Node) (x *XPathNSResolver) {
	js.Rewrite("$<.createNSResolver($1)", nodeResolver)
	return x
}

func (*XPathEvaluator) Evaluate(expression string, contextNode *Node, resolver *XPathNSResolver, kind uint8, result *XPathResult) (x *XPathResult) {
	js.Rewrite("$<.evaluate($1, $2, $3, $4, $5)", expression, contextNode, resolver, kind, result)
	return x
}

type XPathExpression struct {
}

func (*XPathExpression) Evaluate(contextNode *Node, kind uint8, result *XPathResult) (x *XPathExpression) {
	js.Rewrite("$<.evaluate($1, $2, $3)", contextNode, kind, result)
	return x
}

type XPathNSResolver struct {
}

func (*XPathNSResolver) LookupNamespaceURI(prefix string) (s string) {
	js.Rewrite("$<.lookupNamespaceURI($1)", prefix)
	return s
}

type XPathResult struct {
}

func (*XPathResult) IterateNext() (n *Node) {
	js.Rewrite("$<.iterateNext()")
	return n
}

func (*XPathResult) SnapshotItem(index uint) (n *Node) {
	js.Rewrite("$<.snapshotItem($1)", index)
	return n
}

func (*XPathResult) GetBooleanValue() (booleanValue bool) {
	js.Rewrite("$<.booleanValue")
	return booleanValue
}

func (*XPathResult) GetInvalidIteratorState() (invalidIteratorState bool) {
	js.Rewrite("$<.invalidIteratorState")
	return invalidIteratorState
}

func (*XPathResult) GetNumberValue() (numberValue float32) {
	js.Rewrite("$<.numberValue")
	return numberValue
}

func (*XPathResult) GetResultType() (resultType uint8) {
	js.Rewrite("$<.resultType")
	return resultType
}

func (*XPathResult) GetSingleNodeValue() (singleNodeValue *Node) {
	js.Rewrite("$<.singleNodeValue")
	return singleNodeValue
}

func (*XPathResult) GetSnapshotLength() (snapshotLength uint) {
	js.Rewrite("$<.snapshotLength")
	return snapshotLength
}

func (*XPathResult) GetStringValue() (stringValue string) {
	js.Rewrite("$<.stringValue")
	return stringValue
}

type XSLTProcessor struct {
}

func (*XSLTProcessor) ClearParameters() {
	js.Rewrite("$<.clearParameters()")
}

func (*XSLTProcessor) GetParameter(namespaceURI string, localName string) (i interface{}) {
	js.Rewrite("$<.getParameter($1, $2)", namespaceURI, localName)
	return i
}

func (*XSLTProcessor) ImportStylesheet(style *Node) {
	js.Rewrite("$<.importStylesheet($1)", style)
}

func (*XSLTProcessor) RemoveParameter(namespaceURI string, localName string) {
	js.Rewrite("$<.removeParameter($1, $2)", namespaceURI, localName)
}

func (*XSLTProcessor) Reset() {
	js.Rewrite("$<.reset()")
}

func (*XSLTProcessor) SetParameter(namespaceURI string, localName string, value interface{}) {
	js.Rewrite("$<.setParameter($1, $2, $3)", namespaceURI, localName, value)
}

func (*XSLTProcessor) TransformToDocument(source *Node) (d *Document) {
	js.Rewrite("$<.transformToDocument($1)", source)
	return d
}

func (*XSLTProcessor) TransformToFragment(source *Node, document *Document) (d *DocumentFragment) {
	js.Rewrite("$<.transformToFragment($1, $2)", source, document)
	return d
}

type AbstractWorker struct {
}

func (*AbstractWorker) GetOnerror() (e *Event) {
	js.Rewrite("$<.onerror")
	return e
}

func (*AbstractWorker) SetOnerror(e *Event) {
	js.Rewrite("$<.onerror = $1", e)
}

type Body struct {
}

func (*Body) ArrayBuffer() (b []byte) {
	js.Rewrite("await $<.arrayBuffer()")
	return b
}

func (*Body) Blob() (b *Blob) {
	js.Rewrite("await $<.blob()")
	return b
}

func (*Body) JSON() (i interface{}) {
	js.Rewrite("await $<.json()")
	return i
}

func (*Body) Text() (s string) {
	js.Rewrite("await $<.text()")
	return s
}

func (*Body) GetBodyUsed() (bodyUsed bool) {
	js.Rewrite("$<.bodyUsed")
	return bodyUsed
}

type CanvasPathMethods struct {
}

func (*CanvasPathMethods) Arc(x float32, y float32, radius float32, startAngle float32, endAngle float32, anticlockwise bool) {
	js.Rewrite("$<.arc($1, $2, $3, $4, $5, $6)", x, y, radius, startAngle, endAngle, anticlockwise)
}

func (*CanvasPathMethods) ArcTo(x1 float32, y1 float32, x2 float32, y2 float32, radius float32) {
	js.Rewrite("$<.arcTo($1, $2, $3, $4, $5)", x1, y1, x2, y2, radius)
}

func (*CanvasPathMethods) BezierCurveTo(cp1x float32, cp1y float32, cp2x float32, cp2y float32, x float32, y float32) {
	js.Rewrite("$<.bezierCurveTo($1, $2, $3, $4, $5, $6)", cp1x, cp1y, cp2x, cp2y, x, y)
}

func (*CanvasPathMethods) ClosePath() {
	js.Rewrite("$<.closePath()")
}

func (*CanvasPathMethods) Ellipse(x float32, y float32, radiusX float32, radiusY float32, rotation float32, startAngle float32, endAngle float32, anticlockwise bool) {
	js.Rewrite("$<.ellipse($1, $2, $3, $4, $5, $6, $7, $8)", x, y, radiusX, radiusY, rotation, startAngle, endAngle, anticlockwise)
}

func (*CanvasPathMethods) LineTo(x float32, y float32) {
	js.Rewrite("$<.lineTo($1, $2)", x, y)
}

func (*CanvasPathMethods) MoveTo(x float32, y float32) {
	js.Rewrite("$<.moveTo($1, $2)", x, y)
}

func (*CanvasPathMethods) QuadraticCurveTo(cpx float32, cpy float32, x float32, y float32) {
	js.Rewrite("$<.quadraticCurveTo($1, $2, $3, $4)", cpx, cpy, x, y)
}

func (*CanvasPathMethods) Rect(x float32, y float32, w float32, h float32) {
	js.Rewrite("$<.rect($1, $2, $3, $4)", x, y, w, h)
}

type ChildNode struct {
}

func (*ChildNode) Remove() {
	js.Rewrite("$<.remove()")
}

type DocumentEvent struct {
}

func (*DocumentEvent) CreateEvent(eventInterface string) (e *Event) {
	js.Rewrite("$<.createEvent($1)", eventInterface)
	return e
}

type DOML2DeprecatedColorProperty struct {
}

func (*DOML2DeprecatedColorProperty) GetColor() (color string) {
	js.Rewrite("$<.color")
	return color
}

func (*DOML2DeprecatedColorProperty) SetColor(color string) {
	js.Rewrite("$<.color = $1", color)
}

type DOML2DeprecatedSizeProperty struct {
}

func (*DOML2DeprecatedSizeProperty) GetSize() (size int) {
	js.Rewrite("$<.size")
	return size
}

func (*DOML2DeprecatedSizeProperty) SetSize(size int) {
	js.Rewrite("$<.size = $1", size)
}

type ElementTraversal struct {
}

func (*ElementTraversal) GetChildElementCount() (childElementCount uint) {
	js.Rewrite("$<.childElementCount")
	return childElementCount
}

func (*ElementTraversal) GetFirstElementChild() (firstElementChild *Element) {
	js.Rewrite("$<.firstElementChild")
	return firstElementChild
}

func (*ElementTraversal) GetLastElementChild() (lastElementChild *Element) {
	js.Rewrite("$<.lastElementChild")
	return lastElementChild
}

func (*ElementTraversal) GetNextElementSibling() (nextElementSibling *Element) {
	js.Rewrite("$<.nextElementSibling")
	return nextElementSibling
}

func (*ElementTraversal) GetPreviousElementSibling() (previousElementSibling *Element) {
	js.Rewrite("$<.previousElementSibling")
	return previousElementSibling
}

type GetSVGDocument struct {
}

func (*GetSVGDocument) GetSVGDocument() (d *Document) {
	js.Rewrite("$<.getSVGDocument()")
	return d
}

type GlobalEventHandlers struct {
}

func (*GlobalEventHandlers) GetOnpointercancel() (e *Event) {
	js.Rewrite("$<.onpointercancel")
	return e
}

func (*GlobalEventHandlers) SetOnpointercancel(e *Event) {
	js.Rewrite("$<.onpointercancel = $1", e)
}

func (*GlobalEventHandlers) GetOnpointerdown() (e *Event) {
	js.Rewrite("$<.onpointerdown")
	return e
}

func (*GlobalEventHandlers) SetOnpointerdown(e *Event) {
	js.Rewrite("$<.onpointerdown = $1", e)
}

func (*GlobalEventHandlers) GetOnpointerenter() (e *Event) {
	js.Rewrite("$<.onpointerenter")
	return e
}

func (*GlobalEventHandlers) SetOnpointerenter(e *Event) {
	js.Rewrite("$<.onpointerenter = $1", e)
}

func (*GlobalEventHandlers) GetOnpointerleave() (e *Event) {
	js.Rewrite("$<.onpointerleave")
	return e
}

func (*GlobalEventHandlers) SetOnpointerleave(e *Event) {
	js.Rewrite("$<.onpointerleave = $1", e)
}

func (*GlobalEventHandlers) GetOnpointermove() (e *Event) {
	js.Rewrite("$<.onpointermove")
	return e
}

func (*GlobalEventHandlers) SetOnpointermove(e *Event) {
	js.Rewrite("$<.onpointermove = $1", e)
}

func (*GlobalEventHandlers) GetOnpointerout() (e *Event) {
	js.Rewrite("$<.onpointerout")
	return e
}

func (*GlobalEventHandlers) SetOnpointerout(e *Event) {
	js.Rewrite("$<.onpointerout = $1", e)
}

func (*GlobalEventHandlers) GetOnpointerover() (e *Event) {
	js.Rewrite("$<.onpointerover")
	return e
}

func (*GlobalEventHandlers) SetOnpointerover(e *Event) {
	js.Rewrite("$<.onpointerover = $1", e)
}

func (*GlobalEventHandlers) GetOnpointerup() (e *Event) {
	js.Rewrite("$<.onpointerup")
	return e
}

func (*GlobalEventHandlers) SetOnpointerup(e *Event) {
	js.Rewrite("$<.onpointerup = $1", e)
}

func (*GlobalEventHandlers) GetOnwheel() (e *Event) {
	js.Rewrite("$<.onwheel")
	return e
}

func (*GlobalEventHandlers) SetOnwheel(e *Event) {
	js.Rewrite("$<.onwheel = $1", e)
}

type GlobalFetch struct {
}

func (*GlobalFetch) Fetch(input *Request, init *RequestInit) (r *Response) {
	js.Rewrite("await $<.fetch($1, $2)", input, init)
	return r
}

type HTMLTableAlignment struct {
}

func (*HTMLTableAlignment) GetCh() (ch string) {
	js.Rewrite("$<.ch")
	return ch
}

func (*HTMLTableAlignment) SetCh(ch string) {
	js.Rewrite("$<.ch = $1", ch)
}

func (*HTMLTableAlignment) GetChOff() (chOff string) {
	js.Rewrite("$<.chOff")
	return chOff
}

func (*HTMLTableAlignment) SetChOff(chOff string) {
	js.Rewrite("$<.chOff = $1", chOff)
}

func (*HTMLTableAlignment) GetVAlign() (vAlign string) {
	js.Rewrite("$<.vAlign")
	return vAlign
}

func (*HTMLTableAlignment) SetVAlign(vAlign string) {
	js.Rewrite("$<.vAlign = $1", vAlign)
}

type IDBEnvironment struct {
}

func (*IDBEnvironment) GetIndexedDB() (indexedDB *IDBFactory) {
	js.Rewrite("$<.indexedDB")
	return indexedDB
}

type LinkStyle struct {
}

func (*LinkStyle) GetSheet() (sheet *StyleSheet) {
	js.Rewrite("$<.sheet")
	return sheet
}

type MSBaseReader struct {
}

func (*MSBaseReader) Abort() {
	js.Rewrite("$<.abort()")
}

func (*MSBaseReader) GetOnabort() (e *Event) {
	js.Rewrite("$<.onabort")
	return e
}

func (*MSBaseReader) SetOnabort(e *Event) {
	js.Rewrite("$<.onabort = $1", e)
}

func (*MSBaseReader) GetOnerror() (e *Event) {
	js.Rewrite("$<.onerror")
	return e
}

func (*MSBaseReader) SetOnerror(e *Event) {
	js.Rewrite("$<.onerror = $1", e)
}

func (*MSBaseReader) GetOnload() (e *Event) {
	js.Rewrite("$<.onload")
	return e
}

func (*MSBaseReader) SetOnload(e *Event) {
	js.Rewrite("$<.onload = $1", e)
}

func (*MSBaseReader) GetOnloadend() (e *Event) {
	js.Rewrite("$<.onloadend")
	return e
}

func (*MSBaseReader) SetOnloadend(e *Event) {
	js.Rewrite("$<.onloadend = $1", e)
}

func (*MSBaseReader) GetOnloadstart() (e *Event) {
	js.Rewrite("$<.onloadstart")
	return e
}

func (*MSBaseReader) SetOnloadstart(e *Event) {
	js.Rewrite("$<.onloadstart = $1", e)
}

func (*MSBaseReader) GetOnprogress() (e *Event) {
	js.Rewrite("$<.onprogress")
	return e
}

func (*MSBaseReader) SetOnprogress(e *Event) {
	js.Rewrite("$<.onprogress = $1", e)
}

func (*MSBaseReader) GetReadyState() (readyState uint8) {
	js.Rewrite("$<.readyState")
	return readyState
}

func (*MSBaseReader) GetResult() (result interface{}) {
	js.Rewrite("$<.result")
	return result
}

type MSFileSaver struct {
}

func (*MSFileSaver) MsSaveBlob(blob interface{}, defaultName string) (b bool) {
	js.Rewrite("$<.msSaveBlob($1, $2)", blob, defaultName)
	return b
}

func (*MSFileSaver) MsSaveOrOpenBlob(blob interface{}, defaultName string) (b bool) {
	js.Rewrite("$<.msSaveOrOpenBlob($1, $2)", blob, defaultName)
	return b
}

type MSNavigatorDoNotTrack struct {
}

func (*MSNavigatorDoNotTrack) ConfirmSiteSpecificTrackingException(args *ConfirmSiteSpecificExceptionsInformation) (b bool) {
	js.Rewrite("$<.confirmSiteSpecificTrackingException($1)", args)
	return b
}

func (*MSNavigatorDoNotTrack) ConfirmWebWideTrackingException(args *ExceptionInformation) (b bool) {
	js.Rewrite("$<.confirmWebWideTrackingException($1)", args)
	return b
}

func (*MSNavigatorDoNotTrack) RemoveSiteSpecificTrackingException(args *ExceptionInformation) {
	js.Rewrite("$<.removeSiteSpecificTrackingException($1)", args)
}

func (*MSNavigatorDoNotTrack) RemoveWebWideTrackingException(args *ExceptionInformation) {
	js.Rewrite("$<.removeWebWideTrackingException($1)", args)
}

func (*MSNavigatorDoNotTrack) StoreSiteSpecificTrackingException(args *StoreSiteSpecificExceptionsInformation) {
	js.Rewrite("$<.storeSiteSpecificTrackingException($1)", args)
}

func (*MSNavigatorDoNotTrack) StoreWebWideTrackingException(args *StoreExceptionsInformation) {
	js.Rewrite("$<.storeWebWideTrackingException($1)", args)
}

type NavigatorBeacon struct {
}

func (*NavigatorBeacon) SendBeacon(url string, data interface{}) (b bool) {
	js.Rewrite("$<.sendBeacon($1, $2)", url, data)
	return b
}

type NavigatorConcurrentHardware struct {
}

func (*NavigatorConcurrentHardware) GetHardwareConcurrency() (hardwareConcurrency uint64) {
	js.Rewrite("$<.hardwareConcurrency")
	return hardwareConcurrency
}

type NavigatorContentUtils struct {
}

type NavigatorGeolocation struct {
}

func (*NavigatorGeolocation) GetGeolocation() (geolocation *Geolocation) {
	js.Rewrite("$<.geolocation")
	return geolocation
}

type NavigatorID struct {
}

func (*NavigatorID) GetAppCodeName() (appCodeName string) {
	js.Rewrite("$<.appCodeName")
	return appCodeName
}

func (*NavigatorID) GetAppName() (appName string) {
	js.Rewrite("$<.appName")
	return appName
}

func (*NavigatorID) GetAppVersion() (appVersion string) {
	js.Rewrite("$<.appVersion")
	return appVersion
}

func (*NavigatorID) GetPlatform() (platform string) {
	js.Rewrite("$<.platform")
	return platform
}

func (*NavigatorID) GetProduct() (product string) {
	js.Rewrite("$<.product")
	return product
}

func (*NavigatorID) GetProductSub() (productSub string) {
	js.Rewrite("$<.productSub")
	return productSub
}

func (*NavigatorID) GetUserAgent() (userAgent string) {
	js.Rewrite("$<.userAgent")
	return userAgent
}

func (*NavigatorID) GetVendor() (vendor string) {
	js.Rewrite("$<.vendor")
	return vendor
}

func (*NavigatorID) GetVendorSub() (vendorSub string) {
	js.Rewrite("$<.vendorSub")
	return vendorSub
}

type NavigatorOnLine struct {
}

func (*NavigatorOnLine) GetOnLine() (onLine bool) {
	js.Rewrite("$<.onLine")
	return onLine
}

type NavigatorStorageUtils struct {
}

type NavigatorUserMedia struct {
}

func (*NavigatorUserMedia) GetUserMedia(constraints *MediaStreamConstraints, successCallback func(stream *MediaStream), errorCallback func(err *MediaStreamError)) {
	js.Rewrite("$<.getUserMedia($1, $2, $3)", constraints, successCallback, errorCallback)
}

func (*NavigatorUserMedia) GetMediaDevices() (mediaDevices *MediaDevices) {
	js.Rewrite("$<.mediaDevices")
	return mediaDevices
}

type NodeSelector struct {
}

func (*NodeSelector) QuerySelector(selectors string) (e *Element) {
	js.Rewrite("$<.querySelector($1)", selectors)
	return e
}

func (*NodeSelector) QuerySelectorAll(selectors string) (n *NodeList) {
	js.Rewrite("$<.querySelectorAll($1)", selectors)
	return n
}

type RandomSource struct {
}

func (*RandomSource) GetRandomValues(array []byte) (b []byte) {
	js.Rewrite("$<.getRandomValues($1)", array)
	return b
}

type SVGAnimatedPoints struct {
}

func (*SVGAnimatedPoints) GetAnimatedPoints() (animatedPoints *SVGPointList) {
	js.Rewrite("$<.animatedPoints")
	return animatedPoints
}

func (*SVGAnimatedPoints) GetPoints() (points *SVGPointList) {
	js.Rewrite("$<.points")
	return points
}

type SVGFilterPrimitiveStandardAttributes struct {
}

func (*SVGFilterPrimitiveStandardAttributes) GetHeight() (height *SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

func (*SVGFilterPrimitiveStandardAttributes) GetResult() (result *SVGAnimatedString) {
	js.Rewrite("$<.result")
	return result
}

func (*SVGFilterPrimitiveStandardAttributes) GetWidth() (width *SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

func (*SVGFilterPrimitiveStandardAttributes) GetX() (x *SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

func (*SVGFilterPrimitiveStandardAttributes) GetY() (y *SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}

type SVGFitToViewBox struct {
}

func (*SVGFitToViewBox) GetPreserveAspectRatio() (preserveAspectRatio *SVGAnimatedPreserveAspectRatio) {
	js.Rewrite("$<.preserveAspectRatio")
	return preserveAspectRatio
}

func (*SVGFitToViewBox) GetViewBox() (viewBox *SVGAnimatedRect) {
	js.Rewrite("$<.viewBox")
	return viewBox
}

type SVGTests struct {
}

func (*SVGTests) HasExtension(extension string) (b bool) {
	js.Rewrite("$<.hasExtension($1)", extension)
	return b
}

func (*SVGTests) GetRequiredExtensions() (requiredExtensions *SVGStringList) {
	js.Rewrite("$<.requiredExtensions")
	return requiredExtensions
}

func (*SVGTests) GetRequiredFeatures() (requiredFeatures *SVGStringList) {
	js.Rewrite("$<.requiredFeatures")
	return requiredFeatures
}

func (*SVGTests) GetSystemLanguage() (systemLanguage *SVGStringList) {
	js.Rewrite("$<.systemLanguage")
	return systemLanguage
}

type SVGURIReference struct {
}

func (*SVGURIReference) GetHref() (href *SVGAnimatedString) {
	js.Rewrite("$<.href")
	return href
}

type WindowBase64 struct {
}

func (*WindowBase64) Atob(encodedString string) (s string) {
	js.Rewrite("$<.atob($1)", encodedString)
	return s
}

func (*WindowBase64) Btoa(rawString string) (s string) {
	js.Rewrite("$<.btoa($1)", rawString)
	return s
}

type WindowConsole struct {
}

func (*WindowConsole) GetConsole() (console *Console) {
	js.Rewrite("$<.console")
	return console
}

type WindowLocalStorage struct {
}

func (*WindowLocalStorage) GetLocalStorage() (localStorage *Storage) {
	js.Rewrite("$<.localStorage")
	return localStorage
}

type WindowSessionStorage struct {
}

func (*WindowSessionStorage) GetSessionStorage() (sessionStorage *Storage) {
	js.Rewrite("$<.sessionStorage")
	return sessionStorage
}

type WindowTimers struct {
	*WindowTimersExtension
}

func (*WindowTimers) ClearInterval(handle int) {
	js.Rewrite("$<.clearInterval($1)", handle)
}

func (*WindowTimers) ClearTimeout(handle int) {
	js.Rewrite("$<.clearTimeout($1)", handle)
}

func (*WindowTimers) SetInterval(handler interface{}, timeout interface{}, args interface{}) (i int) {
	js.Rewrite("$<.setInterval($1, $2, $3)", handler, timeout, args)
	return i
}

func (*WindowTimers) SetTimeout(handler interface{}, timeout interface{}, args interface{}) (i int) {
	js.Rewrite("$<.setTimeout($1, $2, $3)", handler, timeout, args)
	return i
}

type WindowTimersExtension struct {
}

func (*WindowTimersExtension) ClearImmediate(handle int) {
	js.Rewrite("$<.clearImmediate($1)", handle)
}

func (*WindowTimersExtension) SetImmediate(expression interface{}, args interface{}) (i int) {
	js.Rewrite("$<.setImmediate($1, $2)", expression, args)
	return i
}

type XMLHttpRequestEventTarget struct {
}

func (*XMLHttpRequestEventTarget) GetOnabort() (e *Event) {
	js.Rewrite("$<.onabort")
	return e
}

func (*XMLHttpRequestEventTarget) SetOnabort(e *Event) {
	js.Rewrite("$<.onabort = $1", e)
}

func (*XMLHttpRequestEventTarget) GetOnerror() (e *Event) {
	js.Rewrite("$<.onerror")
	return e
}

func (*XMLHttpRequestEventTarget) SetOnerror(e *Event) {
	js.Rewrite("$<.onerror = $1", e)
}

func (*XMLHttpRequestEventTarget) GetOnload() (e *Event) {
	js.Rewrite("$<.onload")
	return e
}

func (*XMLHttpRequestEventTarget) SetOnload(e *Event) {
	js.Rewrite("$<.onload = $1", e)
}

func (*XMLHttpRequestEventTarget) GetOnloadend() (e *Event) {
	js.Rewrite("$<.onloadend")
	return e
}

func (*XMLHttpRequestEventTarget) SetOnloadend(e *Event) {
	js.Rewrite("$<.onloadend = $1", e)
}

func (*XMLHttpRequestEventTarget) GetOnloadstart() (e *Event) {
	js.Rewrite("$<.onloadstart")
	return e
}

func (*XMLHttpRequestEventTarget) SetOnloadstart(e *Event) {
	js.Rewrite("$<.onloadstart = $1", e)
}

func (*XMLHttpRequestEventTarget) GetOnprogress() (e *Event) {
	js.Rewrite("$<.onprogress")
	return e
}

func (*XMLHttpRequestEventTarget) SetOnprogress(e *Event) {
	js.Rewrite("$<.onprogress = $1", e)
}

func (*XMLHttpRequestEventTarget) GetOntimeout() (e *Event) {
	js.Rewrite("$<.ontimeout")
	return e
}

func (*XMLHttpRequestEventTarget) SetOntimeout(e *Event) {
	js.Rewrite("$<.ontimeout = $1", e)
}
