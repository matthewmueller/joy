package focuseventinit

import "github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"

type FocusEventInit struct {
	*UIEventInit

	relatedTarget *eventtarget.EventTarget
}
