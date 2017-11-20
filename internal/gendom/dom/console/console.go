package console

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/element"
	"github.com/matthewmueller/golly/js"
)

type Console struct {
}

func (*Console) Assert(test bool, message string, optionalParams interface{}) {
	js.Rewrite("$<.assert($1, $2, $3)", test, message, optionalParams)
}

func (*Console) Clear() {
	js.Rewrite("$<.clear()")
}

func (*Console) Count(countTitle string) {
	js.Rewrite("$<.count($1)", countTitle)
}

func (*Console) Debug(message string, optionalParams interface{}) {
	js.Rewrite("$<.debug($1, $2)", message, optionalParams)
}

func (*Console) Dir(value interface{}, optionalParams interface{}) {
	js.Rewrite("$<.dir($1, $2)", value, optionalParams)
}

func (*Console) Dirxml(value interface{}) {
	js.Rewrite("$<.dirxml($1)", value)
}

func (*Console) Error(message string, optionalParams interface{}) {
	js.Rewrite("$<.error($1, $2)", message, optionalParams)
}

func (*Console) Exception(message string, optionalParams interface{}) {
	js.Rewrite("$<.exception($1, $2)", message, optionalParams)
}

func (*Console) Group(groupTitle string) {
	js.Rewrite("$<.group($1)", groupTitle)
}

func (*Console) GroupCollapsed(groupTitle string) {
	js.Rewrite("$<.groupCollapsed($1)", groupTitle)
}

func (*Console) GroupEnd() {
	js.Rewrite("$<.groupEnd()")
}

func (*Console) Info(message string, optionalParams interface{}) {
	js.Rewrite("$<.info($1, $2)", message, optionalParams)
}

func (*Console) Log(message string, optionalParams interface{}) {
	js.Rewrite("$<.log($1, $2)", message, optionalParams)
}

func (*Console) MsIsIndependentlyComposed(element *element.Element) (b bool) {
	js.Rewrite("$<.msIsIndependentlyComposed($1)", element)
	return b
}

func (*Console) Profile(reportName string) {
	js.Rewrite("$<.profile($1)", reportName)
}

func (*Console) ProfileEnd() {
	js.Rewrite("$<.profileEnd()")
}

func (*Console) Select(element *element.Element) {
	js.Rewrite("$<.select($1)", element)
}

func (*Console) Table(data interface{}) {
	js.Rewrite("$<.table($1)", data)
}

func (*Console) Time(timerName string) {
	js.Rewrite("$<.time($1)", timerName)
}

func (*Console) TimeEnd(timerName string) {
	js.Rewrite("$<.timeEnd($1)", timerName)
}

func (*Console) Trace() {
	js.Rewrite("$<.trace()")
}

func (*Console) Warn(message string, optionalParams interface{}) {
	js.Rewrite("$<.warn($1, $2)", message, optionalParams)
}
