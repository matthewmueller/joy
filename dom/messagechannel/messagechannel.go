package messagechannel

import (
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

// New fn
func New() *MessageChannel {
	macro.Rewrite("MessageChannel")
	return &MessageChannel{}
}

// MessageChannel struct
// js:"MessageChannel,omit"
type MessageChannel struct {
}

// Port1 prop
// js:"port1"
func (*MessageChannel) Port1() (port1 *window.MessagePort) {
	macro.Rewrite("$_.port1")
	return port1
}

// Port2 prop
// js:"port2"
func (*MessageChannel) Port2() (port2 *window.MessagePort) {
	macro.Rewrite("$_.port2")
	return port2
}
