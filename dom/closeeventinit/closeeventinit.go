package closeeventinit

import "github.com/matthewmueller/golly/dom/eventinit"

type CloseEventInit struct {
	*eventinit.EventInit

	code     *uint8
	reason   *string
	wasClean *bool
}
