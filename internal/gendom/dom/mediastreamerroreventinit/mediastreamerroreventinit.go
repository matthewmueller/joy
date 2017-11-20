package mediastreamerroreventinit

import "github.com/matthewmueller/golly/internal/gendom/dom/mediastreamerror"

type MediaStreamErrorEventInit struct {
	*EventInit

	err *mediastreamerror.MediaStreamError
}
