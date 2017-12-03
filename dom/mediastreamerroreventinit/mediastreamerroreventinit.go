package mediastreamerroreventinit

import (
	"github.com/matthewmueller/joy/dom/eventinit"
	"github.com/matthewmueller/joy/dom/mediastreamerror"
)

type MediaStreamErrorEventInit struct {
	*eventinit.EventInit

	err *mediastreamerror.MediaStreamError
}
