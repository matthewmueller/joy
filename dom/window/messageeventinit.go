package window

import "github.com/matthewmueller/golly/dom/eventinit"

type MessageEventInit struct {
	*eventinit.EventInit

	data   *interface{}
	origin *string
	ports  *[]*MessagePort
	source *Window
}
