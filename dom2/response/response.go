package response

import (
	"github.com/matthewmueller/golly/dom2/headers"
	"github.com/matthewmueller/golly/dom2/readablestream"
	"github.com/matthewmueller/golly/dom2/responsetype"
	"github.com/matthewmueller/golly/js"
	"github.com/matthewmueller/golly/vdom/h/body"
)

// js:"Response,omit"
type Response struct {
	body.Body
}

// Clone
func (*Response) Clone() (r *Response) {
	js.Rewrite("$<.Clone()")
	return r
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
