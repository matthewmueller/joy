package mediastreameventinit

import (
	"github.com/matthewmueller/golly/dom2/eventinit"
	"github.com/matthewmueller/golly/dom2/window"
)

type MediaStreamEventInit struct {
	*eventinit.EventInit

	stream *window.MediaStream
}
