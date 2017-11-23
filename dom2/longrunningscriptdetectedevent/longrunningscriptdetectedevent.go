package longrunningscriptdetectedevent

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// LongRunningScriptDetectedEvent struct
// js:"LongRunningScriptDetectedEvent,omit"
type LongRunningScriptDetectedEvent struct {
	window.Event
}

// ExecutionTime
func (*LongRunningScriptDetectedEvent) ExecutionTime() (executionTime int) {
	js.Rewrite("$<.ExecutionTime")
	return executionTime
}

// StopPageScriptExecution
func (*LongRunningScriptDetectedEvent) StopPageScriptExecution() (stopPageScriptExecution bool) {
	js.Rewrite("$<.StopPageScriptExecution")
	return stopPageScriptExecution
}

// StopPageScriptExecution
func (*LongRunningScriptDetectedEvent) SetStopPageScriptExecution(stopPageScriptExecution bool) {
	js.Rewrite("$<.StopPageScriptExecution = $1", stopPageScriptExecution)
}
