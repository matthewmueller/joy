package xmlhttprequest

import (
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/dom/xmlhttprequestresponsetype"
	"github.com/matthewmueller/joy/dom/xmlhttprequestupload"
	"github.com/matthewmueller/joy/macro"
)

var _ window.EventTarget = (*XMLHTTPRequest)(nil)

// New fn
func New() *XMLHTTPRequest {
	macro.Rewrite("XMLHttpRequest")
	return &XMLHTTPRequest{}
}

// XMLHTTPRequest struct
// js:"XMLHTTPRequest,omit"
type XMLHTTPRequest struct {
}

// Abort fn
// js:"abort"
func (*XMLHTTPRequest) Abort() {
	macro.Rewrite("$_.abort()")
}

// GetAllResponseHeaders fn
// js:"getAllResponseHeaders"
func (*XMLHTTPRequest) GetAllResponseHeaders() (s string) {
	macro.Rewrite("$_.getAllResponseHeaders()")
	return s
}

// GetResponseHeader fn
// js:"getResponseHeader"
func (*XMLHTTPRequest) GetResponseHeader(header string) (s string) {
	macro.Rewrite("$_.getResponseHeader($1)", header)
	return s
}

// MsCachingEnabled fn
// js:"msCachingEnabled"
func (*XMLHTTPRequest) MsCachingEnabled() (b bool) {
	macro.Rewrite("$_.msCachingEnabled()")
	return b
}

// Open fn
// js:"open"
func (*XMLHTTPRequest) Open(method string, url string, async *bool, user *string, password *string) {
	macro.Rewrite("$_.open($1, $2, $3, $4, $5)", method, url, async, user, password)
}

// OverrideMimeType fn
// js:"overrideMimeType"
func (*XMLHTTPRequest) OverrideMimeType(mime string) {
	macro.Rewrite("$_.overrideMimeType($1)", mime)
}

// Send fn
// js:"send"
func (*XMLHTTPRequest) Send(data *interface{}) {
	macro.Rewrite("$_.send($1)", data)
}

// SetRequestHeader fn
// js:"setRequestHeader"
func (*XMLHTTPRequest) SetRequestHeader(header string, value string) {
	macro.Rewrite("$_.setRequestHeader($1, $2)", header, value)
}

// AddEventListener fn
// js:"addEventListener"
func (*XMLHTTPRequest) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*XMLHTTPRequest) DispatchEvent(evt window.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*XMLHTTPRequest) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// MsCaching prop
// js:"msCaching"
func (*XMLHTTPRequest) MsCaching() (msCaching string) {
	macro.Rewrite("$_.msCaching")
	return msCaching
}

// SetMsCaching prop
// js:"msCaching"
func (*XMLHTTPRequest) SetMsCaching(msCaching string) {
	macro.Rewrite("$_.msCaching = $1", msCaching)
}

// Onreadystatechange prop
// js:"onreadystatechange"
func (*XMLHTTPRequest) Onreadystatechange() (onreadystatechange func(window.Event)) {
	macro.Rewrite("$_.onreadystatechange")
	return onreadystatechange
}

// SetOnreadystatechange prop
// js:"onreadystatechange"
func (*XMLHTTPRequest) SetOnreadystatechange(onreadystatechange func(window.Event)) {
	macro.Rewrite("$_.onreadystatechange = $1", onreadystatechange)
}

// ReadyState prop
// js:"readyState"
func (*XMLHTTPRequest) ReadyState() (readyState uint8) {
	macro.Rewrite("$_.readyState")
	return readyState
}

// Response prop
// js:"response"
func (*XMLHTTPRequest) Response() (response interface{}) {
	macro.Rewrite("$_.response")
	return response
}

// ResponseText prop
// js:"responseText"
func (*XMLHTTPRequest) ResponseText() (responseText string) {
	macro.Rewrite("$_.responseText")
	return responseText
}

// ResponseType prop
// js:"responseType"
func (*XMLHTTPRequest) ResponseType() (responseType *xmlhttprequestresponsetype.XMLHTTPRequestResponseType) {
	macro.Rewrite("$_.responseType")
	return responseType
}

// SetResponseType prop
// js:"responseType"
func (*XMLHTTPRequest) SetResponseType(responseType *xmlhttprequestresponsetype.XMLHTTPRequestResponseType) {
	macro.Rewrite("$_.responseType = $1", responseType)
}

// ResponseURL prop
// js:"responseURL"
func (*XMLHTTPRequest) ResponseURL() (responseURL string) {
	macro.Rewrite("$_.responseURL")
	return responseURL
}

// ResponseXML prop
// js:"responseXML"
func (*XMLHTTPRequest) ResponseXML() (responseXML window.Document) {
	macro.Rewrite("$_.responseXML")
	return responseXML
}

// Status prop
// js:"status"
func (*XMLHTTPRequest) Status() (status uint8) {
	macro.Rewrite("$_.status")
	return status
}

// StatusText prop
// js:"statusText"
func (*XMLHTTPRequest) StatusText() (statusText string) {
	macro.Rewrite("$_.statusText")
	return statusText
}

// Timeout prop
// js:"timeout"
func (*XMLHTTPRequest) Timeout() (timeout uint) {
	macro.Rewrite("$_.timeout")
	return timeout
}

// SetTimeout prop
// js:"timeout"
func (*XMLHTTPRequest) SetTimeout(timeout uint) {
	macro.Rewrite("$_.timeout = $1", timeout)
}

// Upload prop
// js:"upload"
func (*XMLHTTPRequest) Upload() (upload *xmlhttprequestupload.XMLHTTPRequestUpload) {
	macro.Rewrite("$_.upload")
	return upload
}

// WithCredentials prop
// js:"withCredentials"
func (*XMLHTTPRequest) WithCredentials() (withCredentials bool) {
	macro.Rewrite("$_.withCredentials")
	return withCredentials
}

// SetWithCredentials prop
// js:"withCredentials"
func (*XMLHTTPRequest) SetWithCredentials(withCredentials bool) {
	macro.Rewrite("$_.withCredentials = $1", withCredentials)
}

// Onabort prop
// js:"onabort"
func (*XMLHTTPRequest) Onabort() (onabort func(window.Event)) {
	macro.Rewrite("$_.onabort")
	return onabort
}

// SetOnabort prop
// js:"onabort"
func (*XMLHTTPRequest) SetOnabort(onabort func(window.Event)) {
	macro.Rewrite("$_.onabort = $1", onabort)
}

// Onerror prop
// js:"onerror"
func (*XMLHTTPRequest) Onerror() (onerror func(window.Event)) {
	macro.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*XMLHTTPRequest) SetOnerror(onerror func(window.Event)) {
	macro.Rewrite("$_.onerror = $1", onerror)
}

// Onload prop
// js:"onload"
func (*XMLHTTPRequest) Onload() (onload func(window.Event)) {
	macro.Rewrite("$_.onload")
	return onload
}

// SetOnload prop
// js:"onload"
func (*XMLHTTPRequest) SetOnload(onload func(window.Event)) {
	macro.Rewrite("$_.onload = $1", onload)
}

// Onloadend prop
// js:"onloadend"
func (*XMLHTTPRequest) Onloadend() (onloadend func(window.Event)) {
	macro.Rewrite("$_.onloadend")
	return onloadend
}

// SetOnloadend prop
// js:"onloadend"
func (*XMLHTTPRequest) SetOnloadend(onloadend func(window.Event)) {
	macro.Rewrite("$_.onloadend = $1", onloadend)
}

// Onloadstart prop
// js:"onloadstart"
func (*XMLHTTPRequest) Onloadstart() (onloadstart func(window.Event)) {
	macro.Rewrite("$_.onloadstart")
	return onloadstart
}

// SetOnloadstart prop
// js:"onloadstart"
func (*XMLHTTPRequest) SetOnloadstart(onloadstart func(window.Event)) {
	macro.Rewrite("$_.onloadstart = $1", onloadstart)
}

// Onprogress prop
// js:"onprogress"
func (*XMLHTTPRequest) Onprogress() (onprogress func(window.Event)) {
	macro.Rewrite("$_.onprogress")
	return onprogress
}

// SetOnprogress prop
// js:"onprogress"
func (*XMLHTTPRequest) SetOnprogress(onprogress func(window.Event)) {
	macro.Rewrite("$_.onprogress = $1", onprogress)
}

// Ontimeout prop
// js:"ontimeout"
func (*XMLHTTPRequest) Ontimeout() (ontimeout func(window.Event)) {
	macro.Rewrite("$_.ontimeout")
	return ontimeout
}

// SetOntimeout prop
// js:"ontimeout"
func (*XMLHTTPRequest) SetOntimeout(ontimeout func(window.Event)) {
	macro.Rewrite("$_.ontimeout = $1", ontimeout)
}
