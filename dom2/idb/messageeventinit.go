package idb

import "github.com/matthewmueller/golly/dom2/eventinit"

type MessageEventInit struct {
	*eventinit.EventInit

	data   *interface{}
	origin *string
	ports  *[]*MessagePort
	source *Window
}
