package hashchangeeventinit

import "github.com/matthewmueller/golly/dom2/eventinit"

type HashChangeEventInit struct {
	*eventinit.EventInit

	newURL *string
	oldURL *string
}
