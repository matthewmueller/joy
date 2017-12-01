package mediastreameventinit

import (
	"github.com/matthewmueller/golly/dom/eventinit"
	"github.com/matthewmueller/golly/dom/window"
)

type MediaStreamEventInit struct {
	*eventinit.EventInit

	stream *window.MediaStream
}
