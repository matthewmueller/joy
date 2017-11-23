package compositioneventinit

import "github.com/matthewmueller/golly/dom2/window"

type CompositionEventInit struct {
	*window.UIEventInit

	data *string
}
