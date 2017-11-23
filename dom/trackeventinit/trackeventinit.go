package trackeventinit

import "github.com/matthewmueller/golly/dom/eventinit"

type TrackEventInit struct {
	*eventinit.EventInit

	track *interface{}
}
