package windowtimersextension

// WindowTimersExtension interface
// js:"WindowTimersExtension"
type WindowTimersExtension interface {

	// ClearImmediate
	// js:"clearImmediate"
	// jsrewrite:"$_.clearImmediate($1)"
	ClearImmediate(handle int)

	// SetImmediate
	// js:"setImmediate"
	// jsrewrite:"$_.setImmediate($1, $2)"
	SetImmediate(expression interface{}, args interface{}) (i int)
}
