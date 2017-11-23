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

// Cite Sets or retrieves reference information about the object.
func (*HTMLModElement) Cite() (cite string) {
	js.Rewrite("$<.Cite")
	return cite
}

// Cite Sets or retrieves reference information about the object.
func (*HTMLModElement) SetCite(cite string) {
	js.Rewrite("$<.Cite = $1", cite)
}

// DateTime Sets or retrieves the date and time of a modification to the object.
func (*HTMLModElement) DateTime() (dateTime string) {
	js.Rewrite("$<.DateTime")
	return dateTime
}

// DateTime Sets or retrieves the date and time of a modification to the object.
func (*HTMLModElement) SetDateTime(dateTime string) {
	js.Rewrite("$<.DateTime = $1", dateTime)
}
