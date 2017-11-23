package window

import "github.com/matthewmueller/golly/dom2/eventinit"

type UIEventInit struct {
	*eventinit.EventInit

	detail *int
	view   *Window
}
