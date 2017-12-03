package devicelighteventinit

import "github.com/matthewmueller/joy/dom/eventinit"

type DeviceLightEventInit struct {
	*eventinit.EventInit

	value *float32
}
