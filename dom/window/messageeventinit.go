package window

import "github.com/matthewmueller/joy/dom/eventinit"

type MessageEventInit struct {
	*eventinit.EventInit

	data   *interface{}
	origin *string
	ports  *[]*MessagePort
	source *Window
}
