package compositioneventinit

import "github.com/matthewmueller/joy/dom/window"

type CompositionEventInit struct {
	*window.UIEventInit

	data *string
}
