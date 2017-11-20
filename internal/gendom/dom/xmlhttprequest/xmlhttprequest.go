package xmlhttprequest

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/document"
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/xmlhttprequesteventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/xmlhttprequestresponsetype"
	"github.com/matthewmueller/golly/internal/gendom/dom/xmlhttprequestupload"
	"github.com/matthewmueller/golly/js"
)

type XMLHttpRequest struct {
	*eventtarget.EventTarget
	*xmlhttprequesteventtarget.XMLHttpRequestEventTarget
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

func (*XMLHttpRequest) GetOnreadystatechange() (readystatechange *event.Event) {
	js.Rewrite("$<.onreadystatechange")
	return readystatechange
}

func (*XMLHttpRequest) SetOnreadystatechange(readystatechange *event.Event) {
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

func (*XMLHttpRequest) GetResponseType() (responseType *xmlhttprequestresponsetype.XMLHttpRequestResponseType) {
	js.Rewrite("$<.responseType")
	return responseType
}

func (*XMLHttpRequest) SetResponseType(responseType *xmlhttprequestresponsetype.XMLHttpRequestResponseType) {
	js.Rewrite("$<.responseType = $1", responseType)
}

func (*XMLHttpRequest) GetResponseURL() (responseURL string) {
	js.Rewrite("$<.responseURL")
	return responseURL
}

func (*XMLHttpRequest) GetResponseXML() (responseXML *document.Document) {
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

func (*XMLHttpRequest) GetUpload() (upload *xmlhttprequestupload.XMLHttpRequestUpload) {
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
