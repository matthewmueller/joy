package xmlhttprequest

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/dom/xmlhttprequestresponsetype"
	"github.com/matthewmueller/golly/dom/xmlhttprequestupload"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New() *XMLHTTPRequest {
	js.Rewrite("XMLHttpRequest")
	return &XMLHTTPRequest{}
}

// XMLHTTPRequest struct
// js:"XMLHTTPRequest,omit"
type XMLHTTPRequest struct {
	window.EventTarget
}

// Abort fn
func (*XMLHttpRequest) Abort() {
	js.Rewrite("$<.abort()")
}

// GetAllResponseHeaders fn
func (*XMLHttpRequest) GetAllResponseHeaders() (s string) {
	js.Rewrite("$<.getAllResponseHeaders()")
	return s
}

// GetResponseHeader fn
func (*XMLHttpRequest) GetResponseHeader(header string) (s string) {
	js.Rewrite("$<.getResponseHeader($1)", header)
	return s
}

// MsCachingEnabled fn
func (*XMLHttpRequest) MsCachingEnabled() (b bool) {
	js.Rewrite("$<.msCachingEnabled()")
	return b
}

// Open fn
func (*XMLHttpRequest) Open(method string, url string, async *bool, user *string, password *string) {
	js.Rewrite("$<.open($1, $2, $3, $4, $5)", method, url, async, user, password)
}

// OverrideMimeType fn
func (*XMLHttpRequest) OverrideMimeType(mime string) {
	js.Rewrite("$<.overrideMimeType($1)", mime)
}

// Send fn
func (*XMLHttpRequest) Send(data *interface{}) {
	js.Rewrite("$<.send($1)", data)
}

// SetRequestHeader fn
func (*XMLHttpRequest) SetRequestHeader(header string, value string) {
	js.Rewrite("$<.setRequestHeader($1, $2)", header, value)
}

// Onabort prop
func (*XMLHttpRequest) Onabort() (onabort func(window.Event)) {
	js.Rewrite("$<.onabort")
	return onabort
}

// Onabort prop
func (*XMLHttpRequest) SetOnabort(onabort func(window.Event)) {
	js.Rewrite("$<.onabort = $1", onabort)
}

// Onerror prop
func (*XMLHttpRequest) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.onerror")
	return onerror
}

// Onerror prop
func (*XMLHttpRequest) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.onerror = $1", onerror)
}

// Onload prop
func (*XMLHttpRequest) Onload() (onload func(window.Event)) {
	js.Rewrite("$<.onload")
	return onload
}

// Onload prop
func (*XMLHttpRequest) SetOnload(onload func(window.Event)) {
	js.Rewrite("$<.onload = $1", onload)
}

// Onloadend prop
func (*XMLHttpRequest) Onloadend() (onloadend func(window.Event)) {
	js.Rewrite("$<.onloadend")
	return onloadend
}

// Onloadend prop
func (*XMLHttpRequest) SetOnloadend(onloadend func(window.Event)) {
	js.Rewrite("$<.onloadend = $1", onloadend)
}

// Onloadstart prop
func (*XMLHttpRequest) Onloadstart() (onloadstart func(window.Event)) {
	js.Rewrite("$<.onloadstart")
	return onloadstart
}

// Onloadstart prop
func (*XMLHttpRequest) SetOnloadstart(onloadstart func(window.Event)) {
	js.Rewrite("$<.onloadstart = $1", onloadstart)
}

// Onprogress prop
func (*XMLHttpRequest) Onprogress() (onprogress func(window.Event)) {
	js.Rewrite("$<.onprogress")
	return onprogress
}

// Onprogress prop
func (*XMLHttpRequest) SetOnprogress(onprogress func(window.Event)) {
	js.Rewrite("$<.onprogress = $1", onprogress)
}

// Ontimeout prop
func (*XMLHttpRequest) Ontimeout() (ontimeout func(window.Event)) {
	js.Rewrite("$<.ontimeout")
	return ontimeout
}

// Ontimeout prop
func (*XMLHttpRequest) SetOntimeout(ontimeout func(window.Event)) {
	js.Rewrite("$<.ontimeout = $1", ontimeout)
}

// MsCaching prop
func (*XMLHttpRequest) MsCaching() (msCaching string) {
	js.Rewrite("$<.msCaching")
	return msCaching
}

// MsCaching prop
func (*XMLHttpRequest) SetMsCaching(msCaching string) {
	js.Rewrite("$<.msCaching = $1", msCaching)
}

// Onreadystatechange prop
func (*XMLHttpRequest) Onreadystatechange() (onreadystatechange func(window.Event)) {
	js.Rewrite("$<.onreadystatechange")
	return onreadystatechange
}

// Onreadystatechange prop
func (*XMLHttpRequest) SetOnreadystatechange(onreadystatechange func(window.Event)) {
	js.Rewrite("$<.onreadystatechange = $1", onreadystatechange)
}

// ReadyState prop
func (*XMLHttpRequest) ReadyState() (readyState uint8) {
	js.Rewrite("$<.readyState")
	return readyState
}

// Response prop
func (*XMLHttpRequest) Response() (response interface{}) {
	js.Rewrite("$<.response")
	return response
}

// ResponseText prop
func (*XMLHttpRequest) ResponseText() (responseText string) {
	js.Rewrite("$<.responseText")
	return responseText
}

// ResponseType prop
func (*XMLHttpRequest) ResponseType() (responseType *xmlhttprequestresponsetype.XMLHTTPRequestResponseType) {
	js.Rewrite("$<.responseType")
	return responseType
}

// ResponseType prop
func (*XMLHttpRequest) SetResponseType(responseType *xmlhttprequestresponsetype.XMLHTTPRequestResponseType) {
	js.Rewrite("$<.responseType = $1", responseType)
}

// ResponseURL prop
func (*XMLHttpRequest) ResponseURL() (responseURL string) {
	js.Rewrite("$<.responseURL")
	return responseURL
}

// ResponseXML prop
func (*XMLHttpRequest) ResponseXML() (responseXML window.Document) {
	js.Rewrite("$<.responseXML")
	return responseXML
}

// Status prop
func (*XMLHttpRequest) Status() (status uint8) {
	js.Rewrite("$<.status")
	return status
}

// StatusText prop
func (*XMLHttpRequest) StatusText() (statusText string) {
	js.Rewrite("$<.statusText")
	return statusText
}

// Timeout prop
func (*XMLHttpRequest) Timeout() (timeout uint) {
	js.Rewrite("$<.timeout")
	return timeout
}

// Timeout prop
func (*XMLHttpRequest) SetTimeout(timeout uint) {
	js.Rewrite("$<.timeout = $1", timeout)
}

// Upload prop
func (*XMLHttpRequest) Upload() (upload *xmlhttprequestupload.XMLHTTPRequestUpload) {
	js.Rewrite("$<.upload")
	return upload
}

// WithCredentials prop
func (*XMLHttpRequest) WithCredentials() (withCredentials bool) {
	js.Rewrite("$<.withCredentials")
	return withCredentials
}

// WithCredentials prop
func (*XMLHttpRequest) SetWithCredentials(withCredentials bool) {
	js.Rewrite("$<.withCredentials = $1", withCredentials)
}
