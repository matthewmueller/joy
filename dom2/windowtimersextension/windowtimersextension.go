package windowtimersextension

// js:"WindowTimersExtension,omit"
type WindowTimersExtension interface {

	// ClearImmediate
	ClearImmediate(handle int)

	// SetImmediate
	SetImmediate(expression interface{}, args interface{}) (i int)
}
