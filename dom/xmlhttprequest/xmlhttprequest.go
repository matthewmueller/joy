package xmlhttprequest

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/dom/xmlhttprequestresponsetype"
	"github.com/matthewmueller/golly/dom/xmlhttprequestupload"
	"github.com/matthewmueller/golly/js"
)

var _ window.EventTarget = (*XMLHTTPRequest)(nil)

// XMLHTTPRequest struct
// js:"XMLHTTPRequest,omit"
type XMLHTTPRequest struct {
}

// Abort fn
// js:"abort"
func (*XMLHTTPRequest) Abort() {
	js.Rewrite("$_.abort()")
}

// GetAllResponseHeaders fn
// js:"getAllResponseHeaders"
func (*XMLHTTPRequest) GetAllResponseHeaders() (s string) {
	js.Rewrite("$_.getAllResponseHeaders()")
	return s
}

// GetResponseHeader fn
// js:"getResponseHeader"
func (*XMLHTTPRequest) GetResponseHeader(header string) (s string) {
	js.Rewrite("$_.getResponseHeader($1)", header)
	return s
}

// MsCachingEnabled fn
// js:"msCachingEnabled"
func (*XMLHTTPRequest) MsCachingEnabled() (b bool) {
	js.Rewrite("$_.msCachingEnabled()")
	return b
}

// Open fn
// js:"open"
func (*XMLHTTPRequest) Open(method string, url string, async *bool, user *string, password *string) {
	js.Rewrite("$_.open($1, $2, $3, $4, $5)", method, url, async, user, password)
}

// OverrideMimeType fn
// js:"overrideMimeType"
func (*XMLHTTPRequest) OverrideMimeType(mime string) {
	js.Rewrite("$_.overrideMimeType($1)", mime)
}

// Send fn
// js:"send"
func (*XMLHTTPRequest) Send(data *interface{}) {
	js.Rewrite("$_.send($1)", data)
}

// SetRequestHeader fn
// js:"setRequestHeader"
func (*XMLHTTPRequest) SetRequestHeader(header string, value string) {
	js.Rewrite("$_.setRequestHeader($1, $2)", header, value)
}

// AddEventListener fn
// js:"addEventListener"
func (*XMLHTTPRequest) AddEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*XMLHTTPRequest) DispatchEvent(evt window.Event) (b bool) {
	js.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*XMLHTTPRequest) RemoveEventListener(kind string, listener func(evt window.Event), useCapture bool) {
	js.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// MsCaching prop
// js:"msCaching"
func (*XMLHTTPRequest) MsCaching() (msCaching string) {
	js.Rewrite("$_.msCaching")
	return msCaching
}

// SetMsCaching prop
// js:"msCaching"
func (*XMLHTTPRequest) SetMsCaching(msCaching string) {
	js.Rewrite("$_.msCaching = $1", msCaching)
}

// Onreadystatechange prop
// js:"onreadystatechange"
func (*XMLHTTPRequest) Onreadystatechange() (onreadystatechange func(window.Event)) {
	js.Rewrite("$_.onreadystatechange")
	return onreadystatechange
}

// SetOnreadystatechange prop
// js:"onreadystatechange"
func (*XMLHTTPRequest) SetOnreadystatechange(onreadystatechange func(window.Event)) {
	js.Rewrite("$_.onreadystatechange = $1", onreadystatechange)
}

// ReadyState prop
// js:"readyState"
func (*XMLHTTPRequest) ReadyState() (readyState uint8) {
	js.Rewrite("$_.readyState")
	return readyState
}

// Response prop
// js:"response"
func (*XMLHTTPRequest) Response() (response interface{}) {
	js.Rewrite("$_.response")
	return response
}

// ResponseText prop
// js:"responseText"
func (*XMLHTTPRequest) ResponseText() (responseText string) {
	js.Rewrite("$_.responseText")
	return responseText
}

// ResponseType prop
// js:"responseType"
func (*XMLHTTPRequest) ResponseType() (responseType *xmlhttprequestresponsetype.XMLHTTPRequestResponseType) {
	js.Rewrite("$_.responseType")
	return responseType
}

// SetResponseType prop
// js:"responseType"
func (*XMLHTTPRequest) SetResponseType(responseType *xmlhttprequestresponsetype.XMLHTTPRequestResponseType) {
	js.Rewrite("$_.responseType = $1", responseType)
}

// ResponseURL prop
// js:"responseURL"
func (*XMLHTTPRequest) ResponseURL() (responseURL string) {
	js.Rewrite("$_.responseURL")
	return responseURL
}

// ResponseXML prop
// js:"responseXML"
func (*XMLHTTPRequest) ResponseXML() (responseXML window.Document) {
	js.Rewrite("$_.responseXML")
	return responseXML
}

// Status prop
// js:"status"
func (*XMLHTTPRequest) Status() (status uint8) {
	js.Rewrite("$_.status")
	return status
}

// StatusText prop
// js:"statusText"
func (*XMLHTTPRequest) StatusText() (statusText string) {
	js.Rewrite("$_.statusText")
	return statusText
}

// Timeout prop
// js:"timeout"
func (*XMLHTTPRequest) Timeout() (timeout uint) {
	js.Rewrite("$_.timeout")
	return timeout
}

// SetTimeout prop
// js:"timeout"
func (*XMLHTTPRequest) SetTimeout(timeout uint) {
	js.Rewrite("$_.timeout = $1", timeout)
}

// Upload prop
// js:"upload"
func (*XMLHTTPRequest) Upload() (upload *xmlhttprequestupload.XMLHTTPRequestUpload) {
	js.Rewrite("$_.upload")
	return upload
}

// WithCredentials prop
// js:"withCredentials"
func (*XMLHTTPRequest) WithCredentials() (withCredentials bool) {
	js.Rewrite("$_.withCredentials")
	return withCredentials
}

// SetWithCredentials prop
// js:"withCredentials"
func (*XMLHTTPRequest) SetWithCredentials(withCredentials bool) {
	js.Rewrite("$_.withCredentials = $1", withCredentials)
}

// Onabort prop
// js:"onabort"
func (*XMLHTTPRequest) Onabort() (onabort func(window.Event)) {
	js.Rewrite("$_.onabort")
	return onabort
}

// SetOnabort prop
// js:"onabort"
func (*XMLHTTPRequest) SetOnabort(onabort func(window.Event)) {
	js.Rewrite("$_.onabort = $1", onabort)
}

// Onerror prop
// js:"onerror"
func (*XMLHTTPRequest) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$_.onerror")
	return onerror
}

// SetOnerror prop
// js:"onerror"
func (*XMLHTTPRequest) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$_.onerror = $1", onerror)
}

// Onload prop
// js:"onload"
func (*XMLHTTPRequest) Onload() (onload func(window.Event)) {
	js.Rewrite("$_.onload")
	return onload
}

// SetOnload prop
// js:"onload"
func (*XMLHTTPRequest) SetOnload(onload func(window.Event)) {
	js.Rewrite("$_.onload = $1", onload)
}

// Onloadend prop
// js:"onloadend"
func (*XMLHTTPRequest) Onloadend() (onloadend func(window.Event)) {
	js.Rewrite("$_.onloadend")
	return onloadend
}

// SetOnloadend prop
// js:"onloadend"
func (*XMLHTTPRequest) SetOnloadend(onloadend func(window.Event)) {
	js.Rewrite("$_.onloadend = $1", onloadend)
}

// Onloadstart prop
// js:"onloadstart"
func (*XMLHTTPRequest) Onloadstart() (onloadstart func(window.Event)) {
	js.Rewrite("$_.onloadstart")
	return onloadstart
}

// SetOnloadstart prop
// js:"onloadstart"
func (*XMLHTTPRequest) SetOnloadstart(onloadstart func(window.Event)) {
	js.Rewrite("$_.onloadstart = $1", onloadstart)
}

// Onprogress prop
// js:"onprogress"
func (*XMLHTTPRequest) Onprogress() (onprogress func(window.Event)) {
	js.Rewrite("$_.onprogress")
	return onprogress
}

// SetOnprogress prop
// js:"onprogress"
func (*XMLHTTPRequest) SetOnprogress(onprogress func(window.Event)) {
	js.Rewrite("$_.onprogress = $1", onprogress)
}

// Ontimeout prop
// js:"ontimeout"
func (*XMLHTTPRequest) Ontimeout() (ontimeout func(window.Event)) {
	js.Rewrite("$_.ontimeout")
	return ontimeout
}

// SetOntimeout prop
// js:"ontimeout"
func (*XMLHTTPRequest) SetOntimeout(ontimeout func(window.Event)) {
	js.Rewrite("$_.ontimeout = $1", ontimeout)
}
