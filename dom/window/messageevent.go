package window

import "github.com/matthewmueller/golly/js"

// NewMessageEvent fn
func NewMessageEvent(typearg string, eventinitdict *MessageEventInit) *MessageEvent {
	js.Rewrite("MessageEvent")
	return &MessageEvent{}
}

// MessageEvent struct
// js:"MessageEvent,omit"
type MessageEvent struct {
	Event
}

// InitMessageEvent fn
func (*MessageEvent) InitMessageEvent(typeArg string, canBubbleArg bool, cancelableArg bool, dataArg interface{}, originArg string, lastEventIdArg string, sourceArg *Window) {
	js.Rewrite("$<.initMessageEvent($1, $2, $3, $4, $5, $6, $7)", typeArg, canBubbleArg, cancelableArg, dataArg, originArg, lastEventIdArg, sourceArg)
}

// Data prop
func (*MessageEvent) Data() (data interface{}) {
	js.Rewrite("$<.data")
	return data
}

// Origin prop
func (*MessageEvent) Origin() (origin string) {
	js.Rewrite("$<.origin")
	return origin
}

// Ports prop
func (*MessageEvent) Ports() (ports interface{}) {
	js.Rewrite("$<.ports")
	return ports
}

// Source prop
func (*MessageEvent) Source() (source *Window) {
	js.Rewrite("$<.source")
	return source
}
