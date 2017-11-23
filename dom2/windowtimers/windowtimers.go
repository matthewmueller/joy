package windowtimers

import "github.com/matthewmueller/golly/dom2/windowtimersextension"

// js:"WindowTimers,omit"
type WindowTimers interface {
	windowtimersextension.WindowTimersExtension

	// ClearInterval
	ClearInterval(handle int)

	// ClearTimeout
	ClearTimeout(handle int)

	// SetInterval
	SetInterval(handler interface{}, timeout *interface{}, args interface{}) (i int)

	// SetTimeout
	SetTimeout(handler interface{}, timeout *interface{}, args interface{}) (i int)
}
