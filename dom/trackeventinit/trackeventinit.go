package trackeventinit

import "github.com/matthewmueller/joy/dom/eventinit"

type TrackEventInit struct {
	*eventinit.EventInit

	track *interface{}
}
