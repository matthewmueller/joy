package serviceworkermessageeventinit

import "github.com/matthewmueller/golly/internal/gendom/dom/messageport"

type ServiceWorkerMessageEventInit struct {
	*EventInit

	data        *interface{}
	lastEventId *string
	origin      *string
	ports       *[]*messageport.MessagePort
	source      *interface{}
}
