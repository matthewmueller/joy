package messagechannel

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"MessageChannel,omit"
type MessageChannel struct {
}

// Port1
func (*MessageChannel) Port1() (port1 *window.MessagePort) {
	js.Rewrite("$<.Port1")
	return port1
}

// Port2
func (*MessageChannel) Port2() (port2 *window.MessagePort) {
	js.Rewrite("$<.Port2")
	return port2
}
