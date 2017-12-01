package htmltablecellelement

import "github.com/matthewmueller/golly/dom/window"

// HTMLTableCellElement interface
// js:"HTMLTableCellElement"
type HTMLTableCellElement interface {
	window.HTMLElement

	// Ch prop
	// js:"ch"
	// jsrewrite:"$_.ch"
	Ch() (ch string)

	// SetCh prop
	// js:"ch"
	// jsrewrite:"$_.ch = $1"
	SetCh(ch string)

	// ChOff prop
	// js:"chOff"
	// jsrewrite:"$_.chOff"
	ChOff() (chOff string)

	// SetChOff prop
	// js:"chOff"
	// jsrewrite:"$_.chOff = $1"
	SetChOff(chOff string)

	// VAlign prop
	// js:"vAlign"
	// jsrewrite:"$_.vAlign"
	VAlign() (vAlign string)

	// SetVAlign prop
	// js:"vAlign"
	// jsrewrite:"$_.vAlign = $1"
	SetVAlign(vAlign string)

	// Abbr prop Sets or retrieves abbreviated text for the object.
	// js:"abbr"
	// jsrewrite:"$_.abbr"
	Abbr() (abbr string)

	// SetAbbr prop Sets or retrieves abbreviated text for the object.
	// js:"abbr"
	// jsrewrite:"$_.abbr = $1"
	SetAbbr(abbr string)

	// Align prop Sets or retrieves how the object is aligned with adjacent text.
	// js:"align"
	// jsrewrite:"$_.align"
	Align() (align string)

	// SetAlign prop Sets or retrieves how the object is aligned with adjacent text.
	// js:"align"
	// jsrewrite:"$_.align = $1"
	SetAlign(align string)

	// Axis prop Sets or retrieves a comma-delimited list of conceptual categories associated with the object.
	// js:"axis"
	// jsrewrite:"$_.axis"
	Axis() (axis string)

	// SetAxis prop Sets or retrieves a comma-delimited list of conceptual categories associated with the object.
	// js:"axis"
	// jsrewrite:"$_.axis = $1"
	SetAxis(axis string)

	// BgColor prop
	// js:"bgColor"
	// jsrewrite:"$_.bgColor"
	BgColor() (bgColor interface{})

	// SetBgColor prop
	// js:"bgColor"
	// jsrewrite:"$_.bgColor = $1"
	SetBgColor(bgColor interface{})

	// CellIndex prop Retrieves the position of the object in the cells collection of a row.
	// js:"cellIndex"
	// jsrewrite:"$_.cellIndex"
	CellIndex() (cellIndex int)

	// ColSpan prop Sets or retrieves the number columns in the table that the object should span.
	// js:"colSpan"
	// jsrewrite:"$_.colSpan"
	ColSpan() (colSpan uint)

	// SetColSpan prop Sets or retrieves the number columns in the table that the object should span.
	// js:"colSpan"
	// jsrewrite:"$_.colSpan = $1"
	SetColSpan(colSpan uint)

	// Headers prop Sets or retrieves a list of header cells that provide information for the object.
	// js:"headers"
	// jsrewrite:"$_.headers"
	Headers() (headers string)

	// SetHeaders prop Sets or retrieves a list of header cells that provide information for the object.
	// js:"headers"
	// jsrewrite:"$_.headers = $1"
	SetHeaders(headers string)

	// Height prop Sets or retrieves the height of the object.
	// js:"height"
	// jsrewrite:"$_.height"
	Height() (height interface{})

	// SetHeight prop Sets or retrieves the height of the object.
	// js:"height"
	// jsrewrite:"$_.height = $1"
	SetHeight(height interface{})

	// NoWrap prop Sets or retrieves whether the browser automatically performs wordwrap.
	// js:"noWrap"
	// jsrewrite:"$_.noWrap"
	NoWrap() (noWrap bool)

	// SetNoWrap prop Sets or retrieves whether the browser automatically performs wordwrap.
	// js:"noWrap"
	// jsrewrite:"$_.noWrap = $1"
	SetNoWrap(noWrap bool)

	// RowSpan prop Sets or retrieves how many rows in a table the cell should span.
	// js:"rowSpan"
	// jsrewrite:"$_.rowSpan"
	RowSpan() (rowSpan uint)

	// SetRowSpan prop Sets or retrieves how many rows in a table the cell should span.
	// js:"rowSpan"
	// jsrewrite:"$_.rowSpan = $1"
	SetRowSpan(rowSpan uint)

	// Scope prop Sets or retrieves the group of cells in a table to which the object's information applies.
	// js:"scope"
	// jsrewrite:"$_.scope"
	Scope() (scope string)

	// SetScope prop Sets or retrieves the group of cells in a table to which the object's information applies.
	// js:"scope"
	// jsrewrite:"$_.scope = $1"
	SetScope(scope string)

	// Width prop Sets or retrieves the width of the object.
	// js:"width"
	// jsrewrite:"$_.width"
	Width() (width string)

	// SetWidth prop Sets or retrieves the width of the object.
	// js:"width"
	// jsrewrite:"$_.width = $1"
	SetWidth(width string)
}
