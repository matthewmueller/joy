package devicelighteventinit

import "github.com/matthewmueller/golly/dom2/eventinit"

type DeviceLightEventInit struct {
	*eventinit.EventInit

	value *float32
}
