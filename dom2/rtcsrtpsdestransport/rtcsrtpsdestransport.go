package rtcsrtpsdestransport

import (
	"github.com/matthewmueller/golly/dom2/rtcicetransport"
	"github.com/matthewmueller/golly/dom2/rtcsrtpsdesparameters"
	"github.com/matthewmueller/golly/dom2/window"
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

// GetLocalParameters
func (*RTCSrtpSdesTransport) GetLocalParameters() (r []*rtcsrtpsdesparameters.RTCSrtpSdesParameters) {
	js.Rewrite("$<.GetLocalParameters()")
	return r
}

// Onerror
func (*RTCSrtpSdesTransport) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.Onerror")
	return onerror
}

// Onerror
func (*RTCSrtpSdesTransport) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.Onerror = $1", onerror)
}

// Transport
func (*RTCSrtpSdesTransport) Transport() (transport *rtcicetransport.RTCIceTransport) {
	js.Rewrite("$<.Transport")
	return transport
}
