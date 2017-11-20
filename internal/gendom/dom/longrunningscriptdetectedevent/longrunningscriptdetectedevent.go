package longrunningscriptdetectedevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type LongRunningScriptDetectedEvent struct {
	*event.Event
}

func (*LongRunningScriptDetectedEvent) GetExecutionTime() (executionTime int) {
	js.Rewrite("$<.executionTime")
	return executionTime
}

func (*LongRunningScriptDetectedEvent) GetStopPageScriptExecution() (stopPageScriptExecution bool) {
	js.Rewrite("$<.stopPageScriptExecution")
	return stopPageScriptExecution
}

func (*LongRunningScriptDetectedEvent) SetStopPageScriptExecution(stopPageScriptExecution bool) {
	js.Rewrite("$<.stopPageScriptExecution = $1", stopPageScriptExecution)
}
