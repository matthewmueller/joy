package windowtimersextension

// WindowTimersExtension interface
// js:"WindowTimersExtension"
type WindowTimersExtension interface {

	// ClearImmediate
	// js:"clearImmediate"
	ClearImmediate(handle int)

	// SetImmediate
	// js:"setImmediate"
	SetImmediate(expression interface{}, args interface{}) (i int)
}
