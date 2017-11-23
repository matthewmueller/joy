package mediastreamerroreventinit

import (
	"github.com/matthewmueller/golly/dom/eventinit"
	"github.com/matthewmueller/golly/dom/mediastreamerror"
)

type MediaStreamErrorEventInit struct {
	*eventinit.EventInit

	err *mediastreamerror.MediaStreamError
}
