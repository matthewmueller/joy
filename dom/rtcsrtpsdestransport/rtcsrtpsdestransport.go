package rtcsrtpsdestransport

import (
	"github.com/matthewmueller/golly/dom/rtcicetransport"
	"github.com/matthewmueller/golly/dom/rtcsrtpsdesparameters"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(transport *rtcicetransport.RTCIceTransport, encryptparameters *rtcsrtpsdesparameters.RTCSrtpSdesParameters, decryptparameters *rtcsrtpsdesparameters.RTCSrtpSdesParameters) *RTCSrtpSdesTransport {
	js.Rewrite("RTCSrtpSdesTransport")
	return &RTCSrtpSdesTransport{}
}

// RTCSrtpSdesTransport struct
// js:"RTCSrtpSdesTransport,omit"
type RTCSrtpSdesTransport struct {
	window.EventTarget
}

// GetLocalParameters fn
func (*RTCSrtpSdesTransport) GetLocalParameters() (r []*rtcsrtpsdesparameters.RTCSrtpSdesParameters) {
	js.Rewrite("$<.getLocalParameters()")
	return r
}

// Onerror prop
func (*RTCSrtpSdesTransport) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.onerror")
	return onerror
}

// Onerror prop
func (*RTCSrtpSdesTransport) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.onerror = $1", onerror)
}

// Transport prop
func (*RTCSrtpSdesTransport) Transport() (transport *rtcicetransport.RTCIceTransport) {
	js.Rewrite("$<.transport")
	return transport
}
