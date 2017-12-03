package focusnavigationeventinit

import "github.com/matthewmueller/joy/dom/eventinit"

type FocusNavigationEventInit struct {
	*eventinit.EventInit

	navigationReason *string
	originHeight     *float32
	originLeft       *float32
	originTop        *float32
	originWidth      *float32
}
