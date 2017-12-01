package response

import (
	"github.com/matthewmueller/golly/dom/blob"
	"github.com/matthewmueller/golly/dom/headers"
	"github.com/matthewmueller/golly/dom/readablestream"
	"github.com/matthewmueller/golly/dom/responseinit"
	"github.com/matthewmueller/golly/dom/responsetype"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(body *interface{}, init *responseinit.ResponseInit) *Response {
	js.Rewrite("Response")
	return &Response{}
}

// Response struct
// js:"Response,omit"
type Response struct {
}

// Clone fn
// js:"clone"
func (*Response) Clone() (r *Response) {
	js.Rewrite("$_.clone()")
	return r
}

// ArrayBuffer fn
// js:"arrayBuffer"
func (*Response) ArrayBuffer() (b []byte) {
	js.Rewrite("await $_.arrayBuffer()")
	return b
}

// Blob fn
// js:"blob"
func (*Response) Blob() (b blob.Blob) {
	js.Rewrite("await $_.blob()")
	return b
}

// JSON fn
// js:"json"
func (*Response) JSON() (i interface{}) {
	js.Rewrite("await $_.json()")
	return i
}

// Text fn
// js:"text"
func (*Response) Text() (s string) {
	js.Rewrite("await $_.text()")
	return s
}

// Body prop
// js:"body"
func (*Response) Body() (body *readablestream.ReadableStream) {
	js.Rewrite("$_.body")
	return body
}

// Headers prop
// js:"headers"
func (*Response) Headers() (headers *headers.Headers) {
	js.Rewrite("$_.headers")
	return headers
}

// Ok prop
// js:"ok"
func (*Response) Ok() (ok bool) {
	js.Rewrite("$_.ok")
	return ok
}

// Status prop
// js:"status"
func (*Response) Status() (status uint8) {
	js.Rewrite("$_.status")
	return status
}

// StatusText prop
// js:"statusText"
func (*Response) StatusText() (statusText string) {
	js.Rewrite("$_.statusText")
	return statusText
}

// Type prop
// js:"type"
func (*Response) Type() (kind *responsetype.ResponseType) {
	js.Rewrite("$_.type")
	return kind
}

// URL prop
// js:"url"
func (*Response) URL() (url string) {
	js.Rewrite("$_.url")
	return url
}

// BodyUsed prop
// js:"bodyUsed"
func (*Response) BodyUsed() (bodyUsed bool) {
	js.Rewrite("$_.bodyUsed")
	return bodyUsed
}
