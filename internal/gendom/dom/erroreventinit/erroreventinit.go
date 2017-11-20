package erroreventinit

type ErrorEventInit struct {
	*EventInit

	colno    *uint
	err      *interface{}
	filename *string
	lineno   *uint
	message  *string
}
