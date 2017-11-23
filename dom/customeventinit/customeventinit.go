package customeventinit

import "github.com/matthewmueller/golly/dom2/eventinit"

type CustomEventInit struct {
	*eventinit.EventInit

	detail *interface{}
}
