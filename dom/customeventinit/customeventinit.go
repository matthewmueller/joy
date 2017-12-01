package customeventinit

import "github.com/matthewmueller/golly/dom/eventinit"

type CustomEventInit struct {
	*eventinit.EventInit

	detail *interface{}
}
