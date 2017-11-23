package erroreventinit

import "github.com/matthewmueller/golly/dom/eventinit"

type ErrorEventInit struct {
	*eventinit.EventInit

	colno    *uint
	err      *interface{}
	filename *string
	lineno   *uint
	message  *string
}
