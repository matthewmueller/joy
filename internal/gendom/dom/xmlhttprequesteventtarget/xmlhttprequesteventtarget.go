package xmlhttprequesteventtarget

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type XMLHttpRequestEventTarget struct {
}

func (*XMLHttpRequestEventTarget) GetOnabort() (e *event.Event) {
	js.Rewrite("$<.onabort")
	return e
}

func (*XMLHttpRequestEventTarget) SetOnabort(e *event.Event) {
	js.Rewrite("$<.onabort = $1", e)
}

func (*XMLHttpRequestEventTarget) GetOnerror() (e *event.Event) {
	js.Rewrite("$<.onerror")
	return e
}

func (*XMLHttpRequestEventTarget) SetOnerror(e *event.Event) {
	js.Rewrite("$<.onerror = $1", e)
}

func (*XMLHttpRequestEventTarget) GetOnload() (e *event.Event) {
	js.Rewrite("$<.onload")
	return e
}

func (*XMLHttpRequestEventTarget) SetOnload(e *event.Event) {
	js.Rewrite("$<.onload = $1", e)
}

func (*XMLHttpRequestEventTarget) GetOnloadend() (e *event.Event) {
	js.Rewrite("$<.onloadend")
	return e
}

func (*XMLHttpRequestEventTarget) SetOnloadend(e *event.Event) {
	js.Rewrite("$<.onloadend = $1", e)
}

func (*XMLHttpRequestEventTarget) GetOnloadstart() (e *event.Event) {
	js.Rewrite("$<.onloadstart")
	return e
}

func (*XMLHttpRequestEventTarget) SetOnloadstart(e *event.Event) {
	js.Rewrite("$<.onloadstart = $1", e)
}

func (*XMLHttpRequestEventTarget) GetOnprogress() (e *event.Event) {
	js.Rewrite("$<.onprogress")
	return e
}

func (*XMLHttpRequestEventTarget) SetOnprogress(e *event.Event) {
	js.Rewrite("$<.onprogress = $1", e)
}

func (*XMLHttpRequestEventTarget) GetOntimeout() (e *event.Event) {
	js.Rewrite("$<.ontimeout")
	return e
}

func (*XMLHttpRequestEventTarget) SetOntimeout(e *event.Event) {
	js.Rewrite("$<.ontimeout = $1", e)
}
