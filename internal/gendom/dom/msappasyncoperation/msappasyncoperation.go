package msappasyncoperation

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/domerror"
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/js"
)

type MSAppAsyncOperation struct {
	*eventtarget.EventTarget
}

func (*MSAppAsyncOperation) Start() {
	js.Rewrite("$<.start()")
}

func (*MSAppAsyncOperation) GetError() (err *domerror.DOMError) {
	js.Rewrite("$<.error")
	return err
}

func (*MSAppAsyncOperation) GetOncomplete() (complete *event.Event) {
	js.Rewrite("$<.oncomplete")
	return complete
}

func (*MSAppAsyncOperation) SetOncomplete(complete *event.Event) {
	js.Rewrite("$<.oncomplete = $1", complete)
}

func (*MSAppAsyncOperation) GetOnerror() (err *event.Event) {
	js.Rewrite("$<.onerror")
	return err
}

func (*MSAppAsyncOperation) SetOnerror(err *event.Event) {
	js.Rewrite("$<.onerror = $1", err)
}

func (*MSAppAsyncOperation) GetReadyState() (readyState uint8) {
	js.Rewrite("$<.readyState")
	return readyState
}

func (*MSAppAsyncOperation) GetResult() (result interface{}) {
	js.Rewrite("$<.result")
	return result
}
