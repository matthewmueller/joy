package hashchangeeventinit

import "github.com/matthewmueller/joy/dom/eventinit"

type HashChangeEventInit struct {
	*eventinit.EventInit

	newURL *string
	oldURL *string
}
