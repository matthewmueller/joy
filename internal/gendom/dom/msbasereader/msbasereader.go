package msbasereader

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type MSBaseReader struct {
}

func (*MSBaseReader) Abort() {
	js.Rewrite("$<.abort()")
}

func (*MSBaseReader) GetOnabort() (e *event.Event) {
	js.Rewrite("$<.onabort")
	return e
}

func (*MSBaseReader) SetOnabort(e *event.Event) {
	js.Rewrite("$<.onabort = $1", e)
}

func (*MSBaseReader) GetOnerror() (e *event.Event) {
	js.Rewrite("$<.onerror")
	return e
}

func (*MSBaseReader) SetOnerror(e *event.Event) {
	js.Rewrite("$<.onerror = $1", e)
}

func (*MSBaseReader) GetOnload() (e *event.Event) {
	js.Rewrite("$<.onload")
	return e
}

func (*MSBaseReader) SetOnload(e *event.Event) {
	js.Rewrite("$<.onload = $1", e)
}

func (*MSBaseReader) GetOnloadend() (e *event.Event) {
	js.Rewrite("$<.onloadend")
	return e
}

func (*MSBaseReader) SetOnloadend(e *event.Event) {
	js.Rewrite("$<.onloadend = $1", e)
}

func (*MSBaseReader) GetOnloadstart() (e *event.Event) {
	js.Rewrite("$<.onloadstart")
	return e
}

func (*MSBaseReader) SetOnloadstart(e *event.Event) {
	js.Rewrite("$<.onloadstart = $1", e)
}

func (*MSBaseReader) GetOnprogress() (e *event.Event) {
	js.Rewrite("$<.onprogress")
	return e
}

func (*MSBaseReader) SetOnprogress(e *event.Event) {
	js.Rewrite("$<.onprogress = $1", e)
}

func (*MSBaseReader) GetReadyState() (readyState uint8) {
	js.Rewrite("$<.readyState")
	return readyState
}

func (*MSBaseReader) GetResult() (result interface{}) {
	js.Rewrite("$<.result")
	return result
}
