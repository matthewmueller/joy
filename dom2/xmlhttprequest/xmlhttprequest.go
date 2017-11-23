package xmlhttprequest

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/dom2/xmlhttprequesteventtarget"
	"github.com/matthewmueller/golly/dom2/xmlhttprequestresponsetype"
	"github.com/matthewmueller/golly/js"
)

// js:"XMLHTTPRequest,omit"
type XMLHTTPRequest struct {
	window.EventTarget
	xmlhttprequesteventtarget.XMLHTTPRequestEventTarget
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
