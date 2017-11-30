package window

import "github.com/matthewmueller/golly/dom/eventinit"

type ServiceWorkerMessageEventInit struct {
	*eventinit.EventInit

	data        *interface{}
	lastEventId *string
	origin      *string
	ports       *[]*MessagePort
	source      *interface{}
}
