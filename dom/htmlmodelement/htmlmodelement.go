package htmlmodelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLModElement struct
// js:"HTMLModElement,omit"
type HTMLModElement struct {
	window.HTMLElement
}

// Cite prop Sets or retrieves reference information about the object.
func (*HTMLModElement) Cite() (cite string) {
	js.Rewrite("$<.cite")
	return cite
}

// Cite prop Sets or retrieves reference information about the object.
func (*HTMLModElement) SetCite(cite string) {
	js.Rewrite("$<.cite = $1", cite)
}

// DateTime prop Sets or retrieves the date and time of a modification to the object.
func (*HTMLModElement) DateTime() (dateTime string) {
	js.Rewrite("$<.dateTime")
	return dateTime
}

// DateTime prop Sets or retrieves the date and time of a modification to the object.
func (*HTMLModElement) SetDateTime(dateTime string) {
	js.Rewrite("$<.dateTime = $1", dateTime)
}
