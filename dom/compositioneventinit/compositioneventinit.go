package compositioneventinit

import "github.com/matthewmueller/golly/dom/window"

type CompositionEventInit struct {
	*window.UIEventInit

	data *string
}
