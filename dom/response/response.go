package response

import (
	"github.com/matthewmueller/golly/dom2/blob"
	"github.com/matthewmueller/golly/dom2/headers"
	"github.com/matthewmueller/golly/dom2/readablestream"
	"github.com/matthewmueller/golly/dom2/responseinit"
	"github.com/matthewmueller/golly/dom2/responsetype"
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

// ArrayBuffer fn
func (*Response) ArrayBuffer() (b []byte) {
	js.Rewrite("await $<.arrayBuffer()")
	return b
}

// Blob fn
func (*Response) Blob() (b blob.Blob) {
	js.Rewrite("await $<.blob()")
	return b
}

// JSON fn
func (*Response) JSON() (i interface{}) {
	js.Rewrite("await $<.json()")
	return i
}

// Text fn
func (*Response) Text() (s string) {
	js.Rewrite("await $<.text()")
	return s
}

// Clone fn
func (*Response) Clone() (r *Response) {
	js.Rewrite("$<.clone()")
	return r
}

// BodyUsed prop
func (*Response) BodyUsed() (bodyUsed bool) {
	js.Rewrite("$<.bodyUsed")
	return bodyUsed
}

// Body prop
func (*Response) Body() (body *readablestream.ReadableStream) {
	js.Rewrite("$<.body")
	return body
}

// Headers prop
func (*Response) Headers() (headers *headers.Headers) {
	js.Rewrite("$<.headers")
	return headers
}

// Ok prop
func (*Response) Ok() (ok bool) {
	js.Rewrite("$<.ok")
	return ok
}

// Status prop
func (*Response) Status() (status uint8) {
	js.Rewrite("$<.status")
	return status
}

// StatusText prop
func (*Response) StatusText() (statusText string) {
	js.Rewrite("$<.statusText")
	return statusText
}

// Type prop
func (*Response) Type() (kind *responsetype.ResponseType) {
	js.Rewrite("$<.type")
	return kind
}

// URL prop
func (*Response) URL() (url string) {
	js.Rewrite("$<.url")
	return url
}
