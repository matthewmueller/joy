package window

import "github.com/matthewmueller/joy/dom/eventinit"

type UIEventInit struct {
	*eventinit.EventInit

	detail *int
	view   *Window
}
