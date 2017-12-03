package mediastreameventinit

import (
	"github.com/matthewmueller/joy/dom/eventinit"
	"github.com/matthewmueller/joy/dom/window"
)

type MediaStreamEventInit struct {
	*eventinit.EventInit

	stream *window.MediaStream
}
