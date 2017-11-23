package erroreventinit

type ErrorEventInit struct {
	*eventinit.EventInit

	colno    *uint
	err      *interface{}
	filename *string
	lineno   *uint
	message  *string
}
