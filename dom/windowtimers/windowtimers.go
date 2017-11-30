package windowtimers

// WindowTimers interface
// js:"WindowTimers"
type WindowTimers interface {

	// ClearImmediate
	// js:"clearImmediate"
	ClearImmediate(handle int)

	// SetImmediate
	// js:"setImmediate"
	SetImmediate(expression interface{}, args interface{}) (i int)

	// ClearInterval
	// js:"clearInterval"
	ClearInterval(handle int)

	// ClearTimeout
	// js:"clearTimeout"
	ClearTimeout(handle int)

	// SetInterval
	// js:"setInterval"
	SetInterval(handler interface{}, timeout *interface{}, args interface{}) (i int)

	// SetTimeout
	// js:"setTimeout"
	SetTimeout(handler interface{}, timeout *interface{}, args interface{}) (i int)
}
