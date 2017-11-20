package messageevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/internal/gendom/dom/window"
	"github.com/matthewmueller/golly/js"
)

type MessageEvent struct {
	*event.Event
}

func (*MessageEvent) InitMessageEvent(typeArg string, canBubbleArg bool, cancelableArg bool, dataArg interface{}, originArg string, lastEventIdArg string, sourceArg *window.Window) {
	js.Rewrite("$<.initMessageEvent($1, $2, $3, $4, $5, $6, $7)", typeArg, canBubbleArg, cancelableArg, dataArg, originArg, lastEventIdArg, sourceArg)
}

func (*MessageEvent) GetData() (data interface{}) {
	js.Rewrite("$<.data")
	return data
}

func (*MessageEvent) GetOrigin() (origin string) {
	js.Rewrite("$<.origin")
	return origin
}

func (*MessageEvent) GetPorts() (ports interface{}) {
	js.Rewrite("$<.ports")
	return ports
}

func (*MessageEvent) GetSource() (source *window.Window) {
	js.Rewrite("$<.source")
	return source
}
