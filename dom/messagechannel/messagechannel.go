package messagechannel

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New() *MessageChannel {
	js.Rewrite("MessageChannel")
	return &MessageChannel{}
}

// MessageChannel struct
// js:"MessageChannel,omit"
type MessageChannel struct {
}

// Port1 prop
func (*MessageChannel) Port1() (port1 *window.MessagePort) {
	js.Rewrite("$<.port1")
	return port1
}

// Port2 prop
func (*MessageChannel) Port2() (port2 *window.MessagePort) {
	js.Rewrite("$<.port2")
	return port2
}
