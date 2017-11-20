package uieventinit

import "github.com/matthewmueller/golly/internal/gendom/dom/window"

type UIEventInit struct {
	*EventInit

	detail *int
	view   *window.Window
}
