package window

import "github.com/matthewmueller/golly/js"

// js:"Console,omit"
type Console struct {
}

// Assert
func (*Console) Assert(test *bool, message *string, optionalParams interface{}) {
	js.Rewrite("$<.Assert($1, $2, $3)", test, message, optionalParams)
}

// Clear
func (*Console) Clear() {
	js.Rewrite("$<.Clear()")
}

// Count
func (*Console) Count(countTitle *string) {
	js.Rewrite("$<.Count($1)", countTitle)
}

// Debug
func (*Console) Debug(message *string, optionalParams interface{}) {
	js.Rewrite("$<.Debug($1, $2)", message, optionalParams)
}

// Dir
func (*Console) Dir(value *interface{}, optionalParams interface{}) {
	js.Rewrite("$<.Dir($1, $2)", value, optionalParams)
}

// Dirxml
func (*Console) Dirxml(value interface{}) {
	js.Rewrite("$<.Dirxml($1)", value)
}

// Error
func (*Console) Error(message *string, optionalParams interface{}) {
	js.Rewrite("$<.Error($1, $2)", message, optionalParams)
}

// Exception
func (*Console) Exception(message *string, optionalParams interface{}) {
	js.Rewrite("$<.Exception($1, $2)", message, optionalParams)
}

// Group
func (*Console) Group(groupTitle *string) {
	js.Rewrite("$<.Group($1)", groupTitle)
}

// GroupCollapsed
func (*Console) GroupCollapsed(groupTitle *string) {
	js.Rewrite("$<.GroupCollapsed($1)", groupTitle)
}

// GroupEnd
func (*Console) GroupEnd() {
	js.Rewrite("$<.GroupEnd()")
}

// Info
func (*Console) Info(message *string, optionalParams interface{}) {
	js.Rewrite("$<.Info($1, $2)", message, optionalParams)
}

// Log
func (*Console) Log(message *string, optionalParams interface{}) {
	js.Rewrite("$<.Log($1, $2)", message, optionalParams)
}

// MsIsIndependentlyComposed
func (*Console) MsIsIndependentlyComposed(element Element) (b bool) {
	js.Rewrite("$<.MsIsIndependentlyComposed($1)", element)
	return b
}

// Profile
func (*Console) Profile(reportName *string) {
	js.Rewrite("$<.Profile($1)", reportName)
}

// ProfileEnd
func (*Console) ProfileEnd() {
	js.Rewrite("$<.ProfileEnd()")
}

// Select
func (*Console) Select(element Element) {
	js.Rewrite("$<.Select($1)", element)
}

// Table
func (*Console) Table(data interface{}) {
	js.Rewrite("$<.Table($1)", data)
}

// Time
func (*Console) Time(timerName *string) {
	js.Rewrite("$<.Time($1)", timerName)
}

// TimeEnd
func (*Console) TimeEnd(timerName *string) {
	js.Rewrite("$<.TimeEnd($1)", timerName)
}

// Trace
func (*Console) Trace() {
	js.Rewrite("$<.Trace()")
}

// Warn
func (*Console) Warn(message *string, optionalParams interface{}) {
	js.Rewrite("$<.Warn($1, $2)", message, optionalParams)
}
