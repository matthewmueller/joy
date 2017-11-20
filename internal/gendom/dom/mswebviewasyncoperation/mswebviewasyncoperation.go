package mswebviewasyncoperation

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/domerror"
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/mshtmlwebviewelement"
	"github.com/matthewmueller/golly/js"
)

type MSWebViewAsyncOperation struct {
	*eventtarget.EventTarget
}

func (*MSWebViewAsyncOperation) Start() {
	js.Rewrite("$<.start()")
}

func (*MSWebViewAsyncOperation) GetError() (err *domerror.DOMError) {
	js.Rewrite("$<.error")
	return err
}

func (*MSWebViewAsyncOperation) GetOncomplete() (complete *event.Event) {
	js.Rewrite("$<.oncomplete")
	return complete
}

func (*MSWebViewAsyncOperation) SetOncomplete(complete *event.Event) {
	js.Rewrite("$<.oncomplete = $1", complete)
}

func (*MSWebViewAsyncOperation) GetOnerror() (err *event.Event) {
	js.Rewrite("$<.onerror")
	return err
}

func (*MSWebViewAsyncOperation) SetOnerror(err *event.Event) {
	js.Rewrite("$<.onerror = $1", err)
}

func (*MSWebViewAsyncOperation) GetReadyState() (readyState uint8) {
	js.Rewrite("$<.readyState")
	return readyState
}

func (*MSWebViewAsyncOperation) GetResult() (result interface{}) {
	js.Rewrite("$<.result")
	return result
}

func (*MSWebViewAsyncOperation) GetTarget() (target *mshtmlwebviewelement.MSHTMLWebViewElement) {
	js.Rewrite("$<.target")
	return target
}

func (*MSWebViewAsyncOperation) GetType() (kind uint8) {
	js.Rewrite("$<.type")
	return kind
}
