package window

import "github.com/matthewmueller/golly/js"

// Console struct
// js:"Console,omit"
type Console struct {
}

// Assert fn
func (*Console) Assert(test *bool, message *string, optionalParams interface{}) {
	js.Rewrite("$<.assert($1, $2, $3)", test, message, optionalParams)
}

// Clear fn
func (*Console) Clear() {
	js.Rewrite("$<.clear()")
}

// Count fn
func (*Console) Count(countTitle *string) {
	js.Rewrite("$<.count($1)", countTitle)
}

// Debug fn
func (*Console) Debug(message *string, optionalParams interface{}) {
	js.Rewrite("$<.debug($1, $2)", message, optionalParams)
}

// Dir fn
func (*Console) Dir(value *interface{}, optionalParams interface{}) {
	js.Rewrite("$<.dir($1, $2)", value, optionalParams)
}

// Dirxml fn
func (*Console) Dirxml(value interface{}) {
	js.Rewrite("$<.dirxml($1)", value)
}

// Error fn
func (*Console) Error(message *string, optionalParams interface{}) {
	js.Rewrite("$<.error($1, $2)", message, optionalParams)
}

// Exception fn
func (*Console) Exception(message *string, optionalParams interface{}) {
	js.Rewrite("$<.exception($1, $2)", message, optionalParams)
}

// Group fn
func (*Console) Group(groupTitle *string) {
	js.Rewrite("$<.group($1)", groupTitle)
}

// GroupCollapsed fn
func (*Console) GroupCollapsed(groupTitle *string) {
	js.Rewrite("$<.groupCollapsed($1)", groupTitle)
}

// GroupEnd fn
func (*Console) GroupEnd() {
	js.Rewrite("$<.groupEnd()")
}

// Info fn
func (*Console) Info(message *string, optionalParams interface{}) {
	js.Rewrite("$<.info($1, $2)", message, optionalParams)
}

// Log fn
func (*Console) Log(message *string, optionalParams interface{}) {
	js.Rewrite("$<.log($1, $2)", message, optionalParams)
}

// MsIsIndependentlyComposed fn
func (*Console) MsIsIndependentlyComposed(element Element) (b bool) {
	js.Rewrite("$<.msIsIndependentlyComposed($1)", element)
	return b
}

// Profile fn
func (*Console) Profile(reportName *string) {
	js.Rewrite("$<.profile($1)", reportName)
}

// ProfileEnd fn
func (*Console) ProfileEnd() {
	js.Rewrite("$<.profileEnd()")
}

// Select fn
func (*Console) Select(element Element) {
	js.Rewrite("$<.select($1)", element)
}

// Table fn
func (*Console) Table(data interface{}) {
	js.Rewrite("$<.table($1)", data)
}

// Time fn
func (*Console) Time(timerName *string) {
	js.Rewrite("$<.time($1)", timerName)
}

// TimeEnd fn
func (*Console) TimeEnd(timerName *string) {
	js.Rewrite("$<.timeEnd($1)", timerName)
}

// Trace fn
func (*Console) Trace() {
	js.Rewrite("$<.trace()")
}

// Warn fn
func (*Console) Warn(message *string, optionalParams interface{}) {
	js.Rewrite("$<.warn($1, $2)", message, optionalParams)
}
