package longrunningscriptdetectedevent

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// LongRunningScriptDetectedEvent struct
// js:"LongRunningScriptDetectedEvent,omit"
type LongRunningScriptDetectedEvent struct {
	window.Event
}

// ExecutionTime prop
func (*LongRunningScriptDetectedEvent) ExecutionTime() (executionTime int) {
	js.Rewrite("$<.executionTime")
	return executionTime
}

// StopPageScriptExecution prop
func (*LongRunningScriptDetectedEvent) StopPageScriptExecution() (stopPageScriptExecution bool) {
	js.Rewrite("$<.stopPageScriptExecution")
	return stopPageScriptExecution
}

// StopPageScriptExecution prop
func (*LongRunningScriptDetectedEvent) SetStopPageScriptExecution(stopPageScriptExecution bool) {
	js.Rewrite("$<.stopPageScriptExecution = $1", stopPageScriptExecution)
}
