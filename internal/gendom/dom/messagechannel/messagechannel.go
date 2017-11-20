package messagechannel

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/messageport"
	"github.com/matthewmueller/golly/js"
)

type MessageChannel struct {
}

func (*MessageChannel) GetPort1() (port1 *messageport.MessagePort) {
	js.Rewrite("$<.port1")
	return port1
}

func (*MessageChannel) GetPort2() (port2 *messageport.MessagePort) {
	js.Rewrite("$<.port2")
	return port2
}
