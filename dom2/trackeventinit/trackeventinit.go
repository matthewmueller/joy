package trackeventinit

import "github.com/matthewmueller/golly/dom2/eventinit"

type TrackEventInit struct {
	*eventinit.EventInit

	track *interface{}
}
