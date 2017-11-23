package window

import "github.com/matthewmueller/golly/dom2/eventinit"

type ServiceWorkerMessageEventInit struct {
	*eventinit.EventInit

	data        *interface{}
	lastEventId *string
	origin      *string
	ports       *[]*MessagePort
	source      *interface{}
}
