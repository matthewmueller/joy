package response

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/body"
	"github.com/matthewmueller/golly/internal/gendom/dom/headers"
	"github.com/matthewmueller/golly/internal/gendom/dom/readablestream"
	"github.com/matthewmueller/golly/internal/gendom/dom/responsetype"
	"github.com/matthewmueller/golly/js"
)

type Response struct {
	*body.Body
}

func (*Response) Clone() (r *Response) {
	js.Rewrite("$<.clone()")
	return r
}

func (*Response) GetBody() (body *readablestream.ReadableStream) {
	js.Rewrite("$<.body")
	return body
}

func (*Response) GetHeaders() (headers *headers.Headers) {
	js.Rewrite("$<.headers")
	return headers
}

func (*Response) GetOk() (ok bool) {
	js.Rewrite("$<.ok")
	return ok
}

func (*Response) GetStatus() (status uint8) {
	js.Rewrite("$<.status")
	return status
}

func (*Response) GetStatusText() (statusText string) {
	js.Rewrite("$<.statusText")
	return statusText
}

func (*Response) GetType() (kind *responsetype.ResponseType) {
	js.Rewrite("$<.type")
	return kind
}

func (*Response) GetURL() (url string) {
	js.Rewrite("$<.url")
	return url
}
