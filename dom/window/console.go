package window

import "github.com/matthewmueller/golly/js"

// Console struct
// js:"Console,omit"
type Console struct {
}

// Assert fn
// js:"assert"
func (*Console) Assert(test *bool, message *string, optionalParams interface{}) {
	js.Rewrite("$_.assert($1, $2, $3)", test, message, optionalParams)
}

// Clear fn
// js:"clear"
func (*Console) Clear() {
	js.Rewrite("$_.clear()")
}

// Count fn
// js:"count"
func (*Console) Count(countTitle *string) {
	js.Rewrite("$_.count($1)", countTitle)
}

// Debug fn
// js:"debug"
func (*Console) Debug(message *string, optionalParams interface{}) {
	js.Rewrite("$_.debug($1, $2)", message, optionalParams)
}

// Dir fn
// js:"dir"
func (*Console) Dir(value *interface{}, optionalParams interface{}) {
	js.Rewrite("$_.dir($1, $2)", value, optionalParams)
}

// Dirxml fn
// js:"dirxml"
func (*Console) Dirxml(value interface{}) {
	js.Rewrite("$_.dirxml($1)", value)
}

// Error fn
// js:"error"
func (*Console) Error(message *string, optionalParams interface{}) {
	js.Rewrite("$_.error($1, $2)", message, optionalParams)
}

// Exception fn
// js:"exception"
func (*Console) Exception(message *string, optionalParams interface{}) {
	js.Rewrite("$_.exception($1, $2)", message, optionalParams)
}

// Group fn
// js:"group"
func (*Console) Group(groupTitle *string) {
	js.Rewrite("$_.group($1)", groupTitle)
}

// GroupCollapsed fn
// js:"groupCollapsed"
func (*Console) GroupCollapsed(groupTitle *string) {
	js.Rewrite("$_.groupCollapsed($1)", groupTitle)
}

// GroupEnd fn
// js:"groupEnd"
func (*Console) GroupEnd() {
	js.Rewrite("$_.groupEnd()")
}

// Info fn
// js:"info"
func (*Console) Info(message *string, optionalParams interface{}) {
	js.Rewrite("$_.info($1, $2)", message, optionalParams)
}

// Log fn
// js:"log"
func (*Console) Log(message *string, optionalParams interface{}) {
	js.Rewrite("$_.log($1, $2)", message, optionalParams)
}

// MsIsIndependentlyComposed fn
// js:"msIsIndependentlyComposed"
func (*Console) MsIsIndependentlyComposed(element Element) (b bool) {
	js.Rewrite("$_.msIsIndependentlyComposed($1)", element)
	return b
}

// Profile fn
// js:"profile"
func (*Console) Profile(reportName *string) {
	js.Rewrite("$_.profile($1)", reportName)
}

// ProfileEnd fn
// js:"profileEnd"
func (*Console) ProfileEnd() {
	js.Rewrite("$_.profileEnd()")
}

// Select fn
// js:"select"
func (*Console) Select(element Element) {
	js.Rewrite("$_.select($1)", element)
}

// Table fn
// js:"table"
func (*Console) Table(data interface{}) {
	js.Rewrite("$_.table($1)", data)
}

// Time fn
// js:"time"
func (*Console) Time(timerName *string) {
	js.Rewrite("$_.time($1)", timerName)
}

// TimeEnd fn
// js:"timeEnd"
func (*Console) TimeEnd(timerName *string) {
	js.Rewrite("$_.timeEnd($1)", timerName)
}

// Trace fn
// js:"trace"
func (*Console) Trace() {
	js.Rewrite("$_.trace()")
}

// Warn fn
// js:"warn"
func (*Console) Warn(message *string, optionalParams interface{}) {
	js.Rewrite("$_.warn($1, $2)", message, optionalParams)
}
