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

// InitMessageEvent
func (*MessageEvent) InitMessageEvent(typeArg string, canBubbleArg bool, cancelableArg bool, dataArg interface{}, originArg string, lastEventIdArg string, sourceArg *Window) {
	js.Rewrite("$<.InitMessageEvent($1, $2, $3, $4, $5, $6, $7)", typeArg, canBubbleArg, cancelableArg, dataArg, originArg, lastEventIdArg, sourceArg)
}

// Data
func (*MessageEvent) Data() (data interface{}) {
	js.Rewrite("$<.Data")
	return data
}

// Origin
func (*MessageEvent) Origin() (origin string) {
	js.Rewrite("$<.Origin")
	return origin
}

// Ports
func (*MessageEvent) Ports() (ports interface{}) {
	js.Rewrite("$<.Ports")
	return ports
}

// Source
func (*MessageEvent) Source() (source *Window) {
	js.Rewrite("$<.Source")
	return source
}
