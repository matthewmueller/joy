package htmltablecellelement

import "github.com/matthewmueller/golly/dom2/window"

// js:"HTMLTableCellElement,omit"
type HTMLTableCellElement interface {
	window.HTMLElement

	// Abbr prop Sets or retrieves abbreviated text for the object.
	// js:"abbr"
	Abbr() (abbr string)

	// Abbr prop Sets or retrieves abbreviated text for the object.
	SetAbbr(abbr string)

	// Align prop Sets or retrieves how the object is aligned with adjacent text.
	// js:"align"
	Align() (align string)

	// Align prop Sets or retrieves how the object is aligned with adjacent text.
	SetAlign(align string)

	// Axis prop Sets or retrieves a comma-delimited list of conceptual categories associated with the object.
	// js:"axis"
	Axis() (axis string)

	// Axis prop Sets or retrieves a comma-delimited list of conceptual categories associated with the object.
	SetAxis(axis string)

	// BgColor prop
	// js:"bgColor"
	BgColor() (bgColor interface{})

	// BgColor prop
	SetBgColor(bgColor interface{})

	// CellIndex prop Retrieves the position of the object in the cells collection of a row.
	// js:"cellIndex"
	CellIndex() (cellIndex int)

	// ColSpan prop Sets or retrieves the number columns in the table that the object should span.
	// js:"colSpan"
	ColSpan() (colSpan uint)

	// ColSpan prop Sets or retrieves the number columns in the table that the object should span.
	SetColSpan(colSpan uint)

	// Headers prop Sets or retrieves a list of header cells that provide information for the object.
	// js:"headers"
	Headers() (headers string)

	// Headers prop Sets or retrieves a list of header cells that provide information for the object.
	SetHeaders(headers string)

	// Height prop Sets or retrieves the height of the object.
	// js:"height"
	Height() (height interface{})

	// Height prop Sets or retrieves the height of the object.
	SetHeight(height interface{})

	// NoWrap prop Sets or retrieves whether the browser automatically performs wordwrap.
	// js:"noWrap"
	NoWrap() (noWrap bool)

	// NoWrap prop Sets or retrieves whether the browser automatically performs wordwrap.
	SetNoWrap(noWrap bool)

	// RowSpan prop Sets or retrieves how many rows in a table the cell should span.
	// js:"rowSpan"
	RowSpan() (rowSpan uint)

	// RowSpan prop Sets or retrieves how many rows in a table the cell should span.
	SetRowSpan(rowSpan uint)

	// Scope prop Sets or retrieves the group of cells in a table to which the object's information applies.
	// js:"scope"
	Scope() (scope string)

	// Scope prop Sets or retrieves the group of cells in a table to which the object's information applies.
	SetScope(scope string)

	// Width prop Sets or retrieves the width of the object.
	// js:"width"
	Width() (width string)

	// Width prop Sets or retrieves the width of the object.
	SetWidth(width string)
}
