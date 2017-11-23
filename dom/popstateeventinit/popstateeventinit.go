package popstateeventinit

import "github.com/matthewmueller/golly/dom2/eventinit"

type PopStateEventInit struct {
	*eventinit.EventInit

	state *interface{}
}
