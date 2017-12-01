package popstateeventinit

import "github.com/matthewmueller/golly/dom/eventinit"

type PopStateEventInit struct {
	*eventinit.EventInit

	state *interface{}
}
