package htmltableheadercellelement

import "github.com/matthewmueller/golly/js"

// js:"HTMLTableHeaderCellElement,omit"
type HTMLTableHeaderCellElement struct {
	htmltablecellelement.HTMLTableCellElement
}

// Scope Sets or retrieves the group of cells in a table to which the object's information applies.
func (*HTMLTableHeaderCellElement) Scope() (scope string) {
	js.Rewrite("$<.Scope")
	return scope
}

// Scope Sets or retrieves the group of cells in a table to which the object's information applies.
func (*HTMLTableHeaderCellElement) SetScope(scope string) {
	js.Rewrite("$<.Scope = $1", scope)
}
