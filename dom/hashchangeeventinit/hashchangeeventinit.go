package hashchangeeventinit

import "github.com/matthewmueller/golly/dom/eventinit"

type HashChangeEventInit struct {
	*eventinit.EventInit

	newURL *string
	oldURL *string
}
