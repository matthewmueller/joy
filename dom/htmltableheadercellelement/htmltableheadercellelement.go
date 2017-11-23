package htmltableheadercellelement

import (
	"github.com/matthewmueller/golly/dom2/htmltablecellelement"
	"github.com/matthewmueller/golly/js"
)

// HTMLTableHeaderCellElement struct
// js:"HTMLTableHeaderCellElement,omit"
type HTMLTableHeaderCellElement struct {
	htmltablecellelement.HTMLTableCellElement
}

// Scope prop Sets or retrieves the group of cells in a table to which the object's information applies.
func (*HTMLTableHeaderCellElement) Scope() (scope string) {
	js.Rewrite("$<.scope")
	return scope
}

// Scope prop Sets or retrieves the group of cells in a table to which the object's information applies.
func (*HTMLTableHeaderCellElement) SetScope(scope string) {
	js.Rewrite("$<.scope = $1", scope)
}
