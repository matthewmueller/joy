package htmltablecellelement

import "github.com/matthewmueller/golly/dom/window"

// HTMLTableCellElement interface
// js:"HTMLTableCellElement"
type HTMLTableCellElement interface {
	window.HTMLElement

	// Ch prop
	// js:"ch"
	Ch() (ch string)

	// SetCh prop
	// js:"ch"
	SetCh(ch string)

	// ChOff prop
	// js:"chOff"
	ChOff() (chOff string)

	// SetChOff prop
	// js:"chOff"
	SetChOff(chOff string)

	// VAlign prop
	// js:"vAlign"
	VAlign() (vAlign string)

	// SetVAlign prop
	// js:"vAlign"
	SetVAlign(vAlign string)

	// Abbr prop Sets or retrieves abbreviated text for the object.
	// js:"abbr"
	Abbr() (abbr string)

	// SetAbbr prop Sets or retrieves abbreviated text for the object.
	// js:"abbr"
	SetAbbr(abbr string)

	// Align prop Sets or retrieves how the object is aligned with adjacent text.
	// js:"align"
	Align() (align string)

	// SetAlign prop Sets or retrieves how the object is aligned with adjacent text.
	// js:"align"
	SetAlign(align string)

	// Axis prop Sets or retrieves a comma-delimited list of conceptual categories associated with the object.
	// js:"axis"
	Axis() (axis string)

	// SetAxis prop Sets or retrieves a comma-delimited list of conceptual categories associated with the object.
	// js:"axis"
	SetAxis(axis string)

	// BgColor prop
	// js:"bgColor"
	BgColor() (bgColor interface{})

	// SetBgColor prop
	// js:"bgColor"
	SetBgColor(bgColor interface{})

	// CellIndex prop Retrieves the position of the object in the cells collection of a row.
	// js:"cellIndex"
	CellIndex() (cellIndex int)

	// ColSpan prop Sets or retrieves the number columns in the table that the object should span.
	// js:"colSpan"
	ColSpan() (colSpan uint)

	// SetColSpan prop Sets or retrieves the number columns in the table that the object should span.
	// js:"colSpan"
	SetColSpan(colSpan uint)

	// Headers prop Sets or retrieves a list of header cells that provide information for the object.
	// js:"headers"
	Headers() (headers string)

	// SetHeaders prop Sets or retrieves a list of header cells that provide information for the object.
	// js:"headers"
	SetHeaders(headers string)

	// Height prop Sets or retrieves the height of the object.
	// js:"height"
	Height() (height interface{})

	// SetHeight prop Sets or retrieves the height of the object.
	// js:"height"
	SetHeight(height interface{})

	// NoWrap prop Sets or retrieves whether the browser automatically performs wordwrap.
	// js:"noWrap"
	NoWrap() (noWrap bool)

	// SetNoWrap prop Sets or retrieves whether the browser automatically performs wordwrap.
	// js:"noWrap"
	SetNoWrap(noWrap bool)

	// RowSpan prop Sets or retrieves how many rows in a table the cell should span.
	// js:"rowSpan"
	RowSpan() (rowSpan uint)

	// SetRowSpan prop Sets or retrieves how many rows in a table the cell should span.
	// js:"rowSpan"
	SetRowSpan(rowSpan uint)

	// Scope prop Sets or retrieves the group of cells in a table to which the object's information applies.
	// js:"scope"
	Scope() (scope string)

	// SetScope prop Sets or retrieves the group of cells in a table to which the object's information applies.
	// js:"scope"
	SetScope(scope string)

	// Width prop Sets or retrieves the width of the object.
	// js:"width"
	Width() (width string)

	// SetWidth prop Sets or retrieves the width of the object.
	// js:"width"
	SetWidth(width string)
}
