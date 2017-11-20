package messageeventinit

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/messageport"
	"github.com/matthewmueller/golly/internal/gendom/dom/window"
)

type MessageEventInit struct {
	*EventInit

	data   *interface{}
	origin *string
	ports  *[]*messageport.MessagePort
	source *window.Window
}
