package windowtimers

// WindowTimers interface
// js:"WindowTimers"
type WindowTimers interface {

	// ClearImmediate
	// js:"clearImmediate"
	// jsrewrite:"$_.clearImmediate($1)"
	ClearImmediate(handle int)

	// SetImmediate
	// js:"setImmediate"
	// jsrewrite:"$_.setImmediate($1, $2)"
	SetImmediate(expression interface{}, args interface{}) (i int)

	// ClearInterval
	// js:"clearInterval"
	// jsrewrite:"$_.clearInterval($1)"
	ClearInterval(handle int)

	// ClearTimeout
	// js:"clearTimeout"
	// jsrewrite:"$_.clearTimeout($1)"
	ClearTimeout(handle int)

	// SetInterval
	// js:"setInterval"
	// jsrewrite:"$_.setInterval($1, $2, $3)"
	SetInterval(handler interface{}, timeout *interface{}, args interface{}) (i int)

	// SetTimeout
	// js:"setTimeout"
	// jsrewrite:"$_.setTimeout($1, $2, $3)"
	SetTimeout(handler interface{}, timeout *interface{}, args interface{}) (i int)
}
