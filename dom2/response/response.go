package response

import (
	"github.com/matthewmueller/golly/dom2/blob"
	"github.com/matthewmueller/golly/dom2/headers"
	"github.com/matthewmueller/golly/dom2/readablestream"
	"github.com/matthewmueller/golly/dom2/responsetype"
	"github.com/matthewmueller/golly/js"
)

// js:"Response,omit"
type Response struct {
}

// ArrayBuffer
func (*Response) ArrayBuffer() (b []byte) {
	js.Rewrite("await $<.ArrayBuffer()")
	return b
}

// Blob
func (*Response) Blob() (b blob.Blob) {
	js.Rewrite("await $<.Blob()")
	return b
}

// JSON
func (*Response) JSON() (i interface{}) {
	js.Rewrite("await $<.JSON()")
	return i
}

// Text
func (*Response) Text() (s string) {
	js.Rewrite("await $<.Text()")
	return s
}

// Clone
func (*Response) Clone() (r *Response) {
	js.Rewrite("$<.Clone()")
	return r
}

// BodyUsed
func (*Response) BodyUsed() (bodyUsed bool) {
	js.Rewrite("$<.BodyUsed")
	return bodyUsed
}

// Body
func (*Response) Body() (body *readablestream.ReadableStream) {
	js.Rewrite("$<.Body")
	return body
}

// Headers
func (*Response) Headers() (headers *headers.Headers) {
	js.Rewrite("$<.Headers")
	return headers
}

// Ok
func (*Response) Ok() (ok bool) {
	js.Rewrite("$<.Ok")
	return ok
}

// Status
func (*Response) Status() (status uint8) {
	js.Rewrite("$<.Status")
	return status
}

// StatusText
func (*Response) StatusText() (statusText string) {
	js.Rewrite("$<.StatusText")
	return statusText
}

// Type
func (*Response) Type() (kind *responsetype.ResponseType) {
	js.Rewrite("$<.Type")
	return kind
}

// URL
func (*Response) URL() (url string) {
	js.Rewrite("$<.URL")
	return url
}
