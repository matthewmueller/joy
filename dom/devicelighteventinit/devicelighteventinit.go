package devicelighteventinit

import "github.com/matthewmueller/golly/dom/eventinit"

type DeviceLightEventInit struct {
	*eventinit.EventInit

	value *float32
}
