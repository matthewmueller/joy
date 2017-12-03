package popstateeventinit

import "github.com/matthewmueller/joy/dom/eventinit"

type PopStateEventInit struct {
	*eventinit.EventInit

	state *interface{}
}
