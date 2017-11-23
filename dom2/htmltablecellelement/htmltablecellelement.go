package htmltablecellelement

import (
	"github.com/matthewmueller/golly/dom2/htmltablealignment"
	"github.com/matthewmueller/golly/dom2/window"
)

// js:"HTMLTableCellElement,omit"
type HTMLTableCellElement interface {
	window.HTMLElement
	htmltablealignment.HTMLTableAlignment

	// Abbr Sets or retrieves abbreviated text for the object.
	Abbr() (abbr string)

	// Abbr Sets or retrieves abbreviated text for the object.
	SetAbbr(abbr string)

	// Align Sets or retrieves how the object is aligned with adjacent text.
	Align() (align string)

	// Align Sets or retrieves how the object is aligned with adjacent text.
	SetAlign(align string)

	// Axis Sets or retrieves a comma-delimited list of conceptual categories associated with the object.
	Axis() (axis string)

	// Axis Sets or retrieves a comma-delimited list of conceptual categories associated with the object.
	SetAxis(axis string)

	// BgColor
	BgColor() (bgColor interface{})

	// BgColor
	SetBgColor(bgColor interface{})

	// CellIndex Retrieves the position of the object in the cells collection of a row.
	CellIndex() (cellIndex int)

	// ColSpan Sets or retrieves the number columns in the table that the object should span.
	ColSpan() (colSpan uint)

	// ColSpan Sets or retrieves the number columns in the table that the object should span.
	SetColSpan(colSpan uint)

	// Headers Sets or retrieves a list of header cells that provide information for the object.
	Headers() (headers string)

	// Headers Sets or retrieves a list of header cells that provide information for the object.
	SetHeaders(headers string)

	// Height Sets or retrieves the height of the object.
	Height() (height interface{})

	// Height Sets or retrieves the height of the object.
	SetHeight(height interface{})

	// NoWrap Sets or retrieves whether the browser automatically performs wordwrap.
	NoWrap() (noWrap bool)

	// NoWrap Sets or retrieves whether the browser automatically performs wordwrap.
	SetNoWrap(noWrap bool)

	// RowSpan Sets or retrieves how many rows in a table the cell should span.
	RowSpan() (rowSpan uint)

	// RowSpan Sets or retrieves how many rows in a table the cell should span.
	SetRowSpan(rowSpan uint)

	// Scope Sets or retrieves the group of cells in a table to which the object's information applies.
	Scope() (scope string)

	// Scope Sets or retrieves the group of cells in a table to which the object's information applies.
	SetScope(scope string)

	// Width Sets or retrieves the width of the object.
	Width() (width string)

	// Width Sets or retrieves the width of the object.
	SetWidth(width string)
}
