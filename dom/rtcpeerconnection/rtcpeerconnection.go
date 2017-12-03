package rtcpeerconnection

import (
	"github.com/matthewmueller/joy/dom/domerror"
	"github.com/matthewmueller/joy/dom/mediastreamevent"
	"github.com/matthewmueller/joy/dom/rtcconfiguration"
	"github.com/matthewmueller/joy/dom/rtcicecandidate"
	"github.com/matthewmueller/joy/dom/rtciceconnectionstate"
	"github.com/matthewmueller/joy/dom/rtcicegatheringstate"
	"github.com/matthewmueller/joy/dom/rtcofferoptions"
	"github.com/matthewmueller/joy/dom/rtcpeerconnectioniceevent"
	"github.com/matthewmueller/joy/dom/rtcsessiondescription"
	"github.com/matthewmueller/joy/dom/rtcsignalingstate"
	"github.com/matthewmueller/joy/dom/rtcstatsreport"
	"github.com/matthewmueller/joy/dom/window"
)

// RTCPeerConnection interface
// js:"RTCPeerConnection"
type RTCPeerConnection interface {
	window.EventTarget

	// AddIceCandidate
	// js:"addIceCandidate"
	// jsrewrite:"await $_.addIceCandidate($1, $2, $3)"
	AddIceCandidate(candidate *rtcicecandidate.RTCIceCandidate, successCallback *func(), failureCallback *func(err *domerror.DOMError))

	// AddStream
	// js:"addStream"
	// jsrewrite:"$_.addStream($1)"
	AddStream(stream *window.MediaStream)

	// Close
	// js:"close"
	// jsrewrite:"$_.close()"
	Close()

	// CreateAnswer
	// js:"createAnswer"
	// jsrewrite:"await $_.createAnswer($1, $2)"
	CreateAnswer(successCallback *func(sdp *rtcsessiondescription.RTCSessionDescription), failureCallback *func(err *domerror.DOMError)) (r *rtcsessiondescription.RTCSessionDescription)

	// CreateOffer
	// js:"createOffer"
	// jsrewrite:"await $_.createOffer($1, $2, $3)"
	CreateOffer(successCallback *func(sdp *rtcsessiondescription.RTCSessionDescription), failureCallback *func(err *domerror.DOMError), options *rtcofferoptions.RTCOfferOptions) (r *rtcsessiondescription.RTCSessionDescription)

	// GetConfiguration
	// js:"getConfiguration"
	// jsrewrite:"$_.getConfiguration()"
	GetConfiguration() (r *rtcconfiguration.RTCConfiguration)

	// GetLocalStreams
	// js:"getLocalStreams"
	// jsrewrite:"$_.getLocalStreams()"
	GetLocalStreams() (w []*window.MediaStream)

	// GetRemoteStreams
	// js:"getRemoteStreams"
	// jsrewrite:"$_.getRemoteStreams()"
	GetRemoteStreams() (w []*window.MediaStream)

	// GetStats
	// js:"getStats"
	// jsrewrite:"await $_.getStats($1, $2, $3)"
	GetStats(selector *window.MediaStreamTrack, successCallback *func(report *rtcstatsreport.RTCStatsReport), failureCallback *func(err *domerror.DOMError)) (r *rtcstatsreport.RTCStatsReport)

	// GetStreamByID
	// js:"getStreamById"
	// jsrewrite:"$_.getStreamById($1)"
	GetStreamByID(streamId string) (w *window.MediaStream)

	// RemoveStream
	// js:"removeStream"
	// jsrewrite:"$_.removeStream($1)"
	RemoveStream(stream *window.MediaStream)

	// SetLocalDescription
	// js:"setLocalDescription"
	// jsrewrite:"await $_.setLocalDescription($1, $2, $3)"
	SetLocalDescription(description *rtcsessiondescription.RTCSessionDescription, successCallback *func(), failureCallback *func(err *domerror.DOMError))

	// SetRemoteDescription
	// js:"setRemoteDescription"
	// jsrewrite:"await $_.setRemoteDescription($1, $2, $3)"
	SetRemoteDescription(description *rtcsessiondescription.RTCSessionDescription, successCallback *func(), failureCallback *func(err *domerror.DOMError))

	// CanTrickleIceCandidates prop
	// js:"canTrickleIceCandidates"
	// jsrewrite:"$_.canTrickleIceCandidates"
	CanTrickleIceCandidates() (canTrickleIceCandidates bool)

	// IceConnectionState prop
	// js:"iceConnectionState"
	// jsrewrite:"$_.iceConnectionState"
	IceConnectionState() (iceConnectionState *rtciceconnectionstate.RTCIceConnectionState)

	// IceGatheringState prop
	// js:"iceGatheringState"
	// jsrewrite:"$_.iceGatheringState"
	IceGatheringState() (iceGatheringState *rtcicegatheringstate.RTCIceGatheringState)

	// LocalDescription prop
	// js:"localDescription"
	// jsrewrite:"$_.localDescription"
	LocalDescription() (localDescription *rtcsessiondescription.RTCSessionDescription)

	// Onaddstream prop
	// js:"onaddstream"
	// jsrewrite:"$_.onaddstream"
	Onaddstream() (onaddstream func(*mediastreamevent.MediaStreamEvent))

	// SetOnaddstream prop
	// js:"onaddstream"
	// jsrewrite:"$_.onaddstream = $1"
	SetOnaddstream(onaddstream func(*mediastreamevent.MediaStreamEvent))

	// Onicecandidate prop
	// js:"onicecandidate"
	// jsrewrite:"$_.onicecandidate"
	Onicecandidate() (onicecandidate func(*rtcpeerconnectioniceevent.RTCPeerConnectionIceEvent))

	// SetOnicecandidate prop
	// js:"onicecandidate"
	// jsrewrite:"$_.onicecandidate = $1"
	SetOnicecandidate(onicecandidate func(*rtcpeerconnectioniceevent.RTCPeerConnectionIceEvent))

	// Oniceconnectionstatechange prop
	// js:"oniceconnectionstatechange"
	// jsrewrite:"$_.oniceconnectionstatechange"
	Oniceconnectionstatechange() (oniceconnectionstatechange func(window.Event))

	// SetOniceconnectionstatechange prop
	// js:"oniceconnectionstatechange"
	// jsrewrite:"$_.oniceconnectionstatechange = $1"
	SetOniceconnectionstatechange(oniceconnectionstatechange func(window.Event))

	// Onicegatheringstatechange prop
	// js:"onicegatheringstatechange"
	// jsrewrite:"$_.onicegatheringstatechange"
	Onicegatheringstatechange() (onicegatheringstatechange func(window.Event))

	// SetOnicegatheringstatechange prop
	// js:"onicegatheringstatechange"
	// jsrewrite:"$_.onicegatheringstatechange = $1"
	SetOnicegatheringstatechange(onicegatheringstatechange func(window.Event))

	// Onnegotiationneeded prop
	// js:"onnegotiationneeded"
	// jsrewrite:"$_.onnegotiationneeded"
	Onnegotiationneeded() (onnegotiationneeded func(window.Event))

	// SetOnnegotiationneeded prop
	// js:"onnegotiationneeded"
	// jsrewrite:"$_.onnegotiationneeded = $1"
	SetOnnegotiationneeded(onnegotiationneeded func(window.Event))

	// Onremovestream prop
	// js:"onremovestream"
	// jsrewrite:"$_.onremovestream"
	Onremovestream() (onremovestream func(*mediastreamevent.MediaStreamEvent))

	// SetOnremovestream prop
	// js:"onremovestream"
	// jsrewrite:"$_.onremovestream = $1"
	SetOnremovestream(onremovestream func(*mediastreamevent.MediaStreamEvent))

	// Onsignalingstatechange prop
	// js:"onsignalingstatechange"
	// jsrewrite:"$_.onsignalingstatechange"
	Onsignalingstatechange() (onsignalingstatechange func(window.Event))

	// SetOnsignalingstatechange prop
	// js:"onsignalingstatechange"
	// jsrewrite:"$_.onsignalingstatechange = $1"
	SetOnsignalingstatechange(onsignalingstatechange func(window.Event))

	// RemoteDescription prop
	// js:"remoteDescription"
	// jsrewrite:"$_.remoteDescription"
	RemoteDescription() (remoteDescription *rtcsessiondescription.RTCSessionDescription)

	// SignalingState prop
	// js:"signalingState"
	// jsrewrite:"$_.signalingState"
	SignalingState() (signalingState *rtcsignalingstate.RTCSignalingState)
}
