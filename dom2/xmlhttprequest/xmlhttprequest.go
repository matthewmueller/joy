package xmlhttprequest

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/dom2/xmlhttprequestresponsetype"
	"github.com/matthewmueller/golly/dom2/xmlhttprequestupload"
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

// Abort
func (*XMLHttpRequest) Abort() {
	js.Rewrite("$<.Abort()")
}

// GetAllResponseHeaders
func (*XMLHttpRequest) GetAllResponseHeaders() (s string) {
	js.Rewrite("$<.GetAllResponseHeaders()")
	return s
}

// GetResponseHeader
func (*XMLHttpRequest) GetResponseHeader(header string) (s string) {
	js.Rewrite("$<.GetResponseHeader($1)", header)
	return s
}

// MsCachingEnabled
func (*XMLHttpRequest) MsCachingEnabled() (b bool) {
	js.Rewrite("$<.MsCachingEnabled()")
	return b
}

// Open
func (*XMLHttpRequest) Open(method string, url string, async *bool, user *string, password *string) {
	js.Rewrite("$<.Open($1, $2, $3, $4, $5)", method, url, async, user, password)
}

// OverrideMimeType
func (*XMLHttpRequest) OverrideMimeType(mime string) {
	js.Rewrite("$<.OverrideMimeType($1)", mime)
}

// Send
func (*XMLHttpRequest) Send(data *interface{}) {
	js.Rewrite("$<.Send($1)", data)
}

// SetRequestHeader
func (*XMLHttpRequest) SetRequestHeader(header string, value string) {
	js.Rewrite("$<.SetRequestHeader($1, $2)", header, value)
}

// Onabort
func (*XMLHttpRequest) Onabort() (onabort func(window.Event)) {
	js.Rewrite("$<.Onabort")
	return onabort
}

// Onabort
func (*XMLHttpRequest) SetOnabort(onabort func(window.Event)) {
	js.Rewrite("$<.Onabort = $1", onabort)
}

// Onerror
func (*XMLHttpRequest) Onerror() (onerror func(window.Event)) {
	js.Rewrite("$<.Onerror")
	return onerror
}

// Onerror
func (*XMLHttpRequest) SetOnerror(onerror func(window.Event)) {
	js.Rewrite("$<.Onerror = $1", onerror)
}

// Onload
func (*XMLHttpRequest) Onload() (onload func(window.Event)) {
	js.Rewrite("$<.Onload")
	return onload
}

// Onload
func (*XMLHttpRequest) SetOnload(onload func(window.Event)) {
	js.Rewrite("$<.Onload = $1", onload)
}

// Onloadend
func (*XMLHttpRequest) Onloadend() (onloadend func(window.Event)) {
	js.Rewrite("$<.Onloadend")
	return onloadend
}

// Onloadend
func (*XMLHttpRequest) SetOnloadend(onloadend func(window.Event)) {
	js.Rewrite("$<.Onloadend = $1", onloadend)
}

// Onloadstart
func (*XMLHttpRequest) Onloadstart() (onloadstart func(window.Event)) {
	js.Rewrite("$<.Onloadstart")
	return onloadstart
}

// Onloadstart
func (*XMLHttpRequest) SetOnloadstart(onloadstart func(window.Event)) {
	js.Rewrite("$<.Onloadstart = $1", onloadstart)
}

// Onprogress
func (*XMLHttpRequest) Onprogress() (onprogress func(window.Event)) {
	js.Rewrite("$<.Onprogress")
	return onprogress
}

// Onprogress
func (*XMLHttpRequest) SetOnprogress(onprogress func(window.Event)) {
	js.Rewrite("$<.Onprogress = $1", onprogress)
}

// Ontimeout
func (*XMLHttpRequest) Ontimeout() (ontimeout func(window.Event)) {
	js.Rewrite("$<.Ontimeout")
	return ontimeout
}

// Ontimeout
func (*XMLHttpRequest) SetOntimeout(ontimeout func(window.Event)) {
	js.Rewrite("$<.Ontimeout = $1", ontimeout)
}

// MsCaching
func (*XMLHttpRequest) MsCaching() (msCaching string) {
	js.Rewrite("$<.MsCaching")
	return msCaching
}

// MsCaching
func (*XMLHttpRequest) SetMsCaching(msCaching string) {
	js.Rewrite("$<.MsCaching = $1", msCaching)
}

// Onreadystatechange
func (*XMLHttpRequest) Onreadystatechange() (onreadystatechange func(window.Event)) {
	js.Rewrite("$<.Onreadystatechange")
	return onreadystatechange
}

// Onreadystatechange
func (*XMLHttpRequest) SetOnreadystatechange(onreadystatechange func(window.Event)) {
	js.Rewrite("$<.Onreadystatechange = $1", onreadystatechange)
}

// ReadyState
func (*XMLHttpRequest) ReadyState() (readyState uint8) {
	js.Rewrite("$<.ReadyState")
	return readyState
}

// Response
func (*XMLHttpRequest) Response() (response interface{}) {
	js.Rewrite("$<.Response")
	return response
}

// ResponseText
func (*XMLHttpRequest) ResponseText() (responseText string) {
	js.Rewrite("$<.ResponseText")
	return responseText
}

// ResponseType
func (*XMLHttpRequest) ResponseType() (responseType *xmlhttprequestresponsetype.XMLHTTPRequestResponseType) {
	js.Rewrite("$<.ResponseType")
	return responseType
}

// ResponseType
func (*XMLHttpRequest) SetResponseType(responseType *xmlhttprequestresponsetype.XMLHTTPRequestResponseType) {
	js.Rewrite("$<.ResponseType = $1", responseType)
}

// ResponseURL
func (*XMLHttpRequest) ResponseURL() (responseURL string) {
	js.Rewrite("$<.ResponseURL")
	return responseURL
}

// ResponseXML
func (*XMLHttpRequest) ResponseXML() (responseXML window.Document) {
	js.Rewrite("$<.ResponseXML")
	return responseXML
}

// Status
func (*XMLHttpRequest) Status() (status uint8) {
	js.Rewrite("$<.Status")
	return status
}

// StatusText
func (*XMLHttpRequest) StatusText() (statusText string) {
	js.Rewrite("$<.StatusText")
	return statusText
}

// Timeout
func (*XMLHttpRequest) Timeout() (timeout uint) {
	js.Rewrite("$<.Timeout")
	return timeout
}

// Timeout
func (*XMLHttpRequest) SetTimeout(timeout uint) {
	js.Rewrite("$<.Timeout = $1", timeout)
}

// Upload
func (*XMLHttpRequest) Upload() (upload *xmlhttprequestupload.XMLHTTPRequestUpload) {
	js.Rewrite("$<.Upload")
	return upload
}

// WithCredentials
func (*XMLHttpRequest) WithCredentials() (withCredentials bool) {
	js.Rewrite("$<.WithCredentials")
	return withCredentials
}

// WithCredentials
func (*XMLHttpRequest) SetWithCredentials(withCredentials bool) {
	js.Rewrite("$<.WithCredentials = $1", withCredentials)
}
