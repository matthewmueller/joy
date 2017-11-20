package mouseeventinit

import "github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"

type MouseEventInit struct {
	*EventModifierInit

	button        *int8
	buttons       *uint8
	clientX       *int
	clientY       *int
	relatedTarget *eventtarget.EventTarget
	screenX       *int
	screenY       *int
}
