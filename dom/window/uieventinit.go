package window

import "github.com/matthewmueller/golly/dom/eventinit"

type UIEventInit struct {
	*eventinit.EventInit

	detail *int
	view   *Window
}
