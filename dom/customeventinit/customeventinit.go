package customeventinit

import "github.com/matthewmueller/joy/dom/eventinit"

type CustomEventInit struct {
	*eventinit.EventInit

	detail *interface{}
}
