package mediastreamerroreventinit

import (
	"github.com/matthewmueller/golly/dom2/eventinit"
	"github.com/matthewmueller/golly/dom2/mediastreamerror"
)

type MediaStreamErrorEventInit struct {
	*eventinit.EventInit

	err *mediastreamerror.MediaStreamError
}
