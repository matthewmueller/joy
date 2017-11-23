package rtcpeerconnection

import (
	"github.com/matthewmueller/golly/dom2/domerror"
	"github.com/matthewmueller/golly/dom2/mediastreamevent"
	"github.com/matthewmueller/golly/dom2/rtcconfiguration"
	"github.com/matthewmueller/golly/dom2/rtcicecandidate"
	"github.com/matthewmueller/golly/dom2/rtciceconnectionstate"
	"github.com/matthewmueller/golly/dom2/rtcicegatheringstate"
	"github.com/matthewmueller/golly/dom2/rtcofferoptions"
	"github.com/matthewmueller/golly/dom2/rtcpeerconnectioniceevent"
	"github.com/matthewmueller/golly/dom2/rtcsessiondescription"
	"github.com/matthewmueller/golly/dom2/rtcsignalingstate"
	"github.com/matthewmueller/golly/dom2/rtcstatsreport"
	"github.com/matthewmueller/golly/dom2/window"
)

// js:"RTCPeerConnection,omit"
type RTCPeerConnection interface {
	window.EventTarget

	// AddIceCandidate
	AddIceCandidate(candidate *rtcicecandidate.RTCIceCandidate, successCallback *func(), failureCallback *func(err *domerror.DOMError))

	// AddStream
	AddStream(stream *window.MediaStream)

	// Close
	Close()

	// CreateAnswer
	CreateAnswer(successCallback *func(sdp *rtcsessiondescription.RTCSessionDescription), failureCallback *func(err *domerror.DOMError)) (r *rtcsessiondescription.RTCSessionDescription)

	// CreateOffer
	CreateOffer(successCallback *func(sdp *rtcsessiondescription.RTCSessionDescription), failureCallback *func(err *domerror.DOMError), options *rtcofferoptions.RTCOfferOptions) (r *rtcsessiondescription.RTCSessionDescription)

	// GetConfiguration
	GetConfiguration() (r *rtcconfiguration.RTCConfiguration)

	// GetLocalStreams
	GetLocalStreams() (w []*window.MediaStream)

	// GetRemoteStreams
	GetRemoteStreams() (w []*window.MediaStream)

	// GetStats
	GetStats(selector *window.MediaStreamTrack, successCallback *func(report *rtcstatsreport.RTCStatsReport), failureCallback *func(err *domerror.DOMError)) (r *rtcstatsreport.RTCStatsReport)

	// GetStreamByID
	GetStreamByID(streamId string) (w *window.MediaStream)

	// RemoveStream
	RemoveStream(stream *window.MediaStream)

	// SetLocalDescription
	SetLocalDescription(description *rtcsessiondescription.RTCSessionDescription, successCallback *func(), failureCallback *func(err *domerror.DOMError))

	// SetRemoteDescription
	SetRemoteDescription(description *rtcsessiondescription.RTCSessionDescription, successCallback *func(), failureCallback *func(err *domerror.DOMError))

	// CanTrickleIceCandidates
	CanTrickleIceCandidates() (canTrickleIceCandidates bool)

	// IceConnectionState
	IceConnectionState() (iceConnectionState *rtciceconnectionstate.RTCIceConnectionState)

	// IceGatheringState
	IceGatheringState() (iceGatheringState *rtcicegatheringstate.RTCIceGatheringState)

	// LocalDescription
	LocalDescription() (localDescription *rtcsessiondescription.RTCSessionDescription)

	// Onaddstream
	Onaddstream() (onaddstream func(*mediastreamevent.MediaStreamEvent))

	// Onaddstream
	SetOnaddstream(onaddstream func(*mediastreamevent.MediaStreamEvent))

	// Onicecandidate
	Onicecandidate() (onicecandidate func(*rtcpeerconnectioniceevent.RTCPeerConnectionIceEvent))

	// Onicecandidate
	SetOnicecandidate(onicecandidate func(*rtcpeerconnectioniceevent.RTCPeerConnectionIceEvent))

	// Oniceconnectionstatechange
	Oniceconnectionstatechange() (oniceconnectionstatechange func(window.Event))

	// Oniceconnectionstatechange
	SetOniceconnectionstatechange(oniceconnectionstatechange func(window.Event))

	// Onicegatheringstatechange
	Onicegatheringstatechange() (onicegatheringstatechange func(window.Event))

	// Onicegatheringstatechange
	SetOnicegatheringstatechange(onicegatheringstatechange func(window.Event))

	// Onnegotiationneeded
	Onnegotiationneeded() (onnegotiationneeded func(window.Event))

	// Onnegotiationneeded
	SetOnnegotiationneeded(onnegotiationneeded func(window.Event))

	// Onremovestream
	Onremovestream() (onremovestream func(*mediastreamevent.MediaStreamEvent))

	// Onremovestream
	SetOnremovestream(onremovestream func(*mediastreamevent.MediaStreamEvent))

	// Onsignalingstatechange
	Onsignalingstatechange() (onsignalingstatechange func(window.Event))

	// Onsignalingstatechange
	SetOnsignalingstatechange(onsignalingstatechange func(window.Event))

	// RemoteDescription
	RemoteDescription() (remoteDescription *rtcsessiondescription.RTCSessionDescription)

	// SignalingState
	SignalingState() (signalingState *rtcsignalingstate.RTCSignalingState)
}
