package htmltablecellelement

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLTableCellElement,omit"
type HTMLTableCellElement interface {
	window.HTMLElement

	// Ch prop
	// js:"ch,rewrite=ch"
	Ch() (ch string)

	// Ch prop
	// js:"setch,rewrite=setch"
	SetCh(ch string)

	// ChOff prop
	// js:"chOff,rewrite=choff"
	ChOff() (chOff string)

	// ChOff prop
	// js:"setchOff,rewrite=setchoff"
	SetChOff(chOff string)

	// VAlign prop
	// js:"vAlign,rewrite=valign"
	VAlign() (vAlign string)

	// VAlign prop
	// js:"setvAlign,rewrite=setvalign"
	SetVAlign(vAlign string)

	// Abbr prop Sets or retrieves abbreviated text for the object.
	// js:"abbr,rewrite=abbr"
	Abbr() (abbr string)

	// Abbr prop Sets or retrieves abbreviated text for the object.
	// js:"setabbr,rewrite=setabbr"
	SetAbbr(abbr string)

	// Align prop Sets or retrieves how the object is aligned with adjacent text.
	// js:"align,rewrite=align"
	Align() (align string)

	// Align prop Sets or retrieves how the object is aligned with adjacent text.
	// js:"setalign,rewrite=setalign"
	SetAlign(align string)

	// Axis prop Sets or retrieves a comma-delimited list of conceptual categories associated with the object.
	// js:"axis,rewrite=axis"
	Axis() (axis string)

	// Axis prop Sets or retrieves a comma-delimited list of conceptual categories associated with the object.
	// js:"setaxis,rewrite=setaxis"
	SetAxis(axis string)

	// BgColor prop
	// js:"bgColor,rewrite=bgcolor"
	BgColor() (bgColor interface{})

	// BgColor prop
	// js:"setbgColor,rewrite=setbgcolor"
	SetBgColor(bgColor interface{})

	// CellIndex prop Retrieves the position of the object in the cells collection of a row.
	// js:"cellIndex,rewrite=cellindex"
	CellIndex() (cellIndex int)

	// ColSpan prop Sets or retrieves the number columns in the table that the object should span.
	// js:"colSpan,rewrite=colspan"
	ColSpan() (colSpan uint)

	// ColSpan prop Sets or retrieves the number columns in the table that the object should span.
	// js:"setcolSpan,rewrite=setcolspan"
	SetColSpan(colSpan uint)

	// Headers prop Sets or retrieves a list of header cells that provide information for the object.
	// js:"headers,rewrite=headers"
	Headers() (headers string)

	// Headers prop Sets or retrieves a list of header cells that provide information for the object.
	// js:"setheaders,rewrite=setheaders"
	SetHeaders(headers string)

	// Height prop Sets or retrieves the height of the object.
	// js:"height,rewrite=height"
	Height() (height interface{})

	// Height prop Sets or retrieves the height of the object.
	// js:"setheight,rewrite=setheight"
	SetHeight(height interface{})

	// NoWrap prop Sets or retrieves whether the browser automatically performs wordwrap.
	// js:"noWrap,rewrite=nowrap"
	NoWrap() (noWrap bool)

	// NoWrap prop Sets or retrieves whether the browser automatically performs wordwrap.
	// js:"setnoWrap,rewrite=setnowrap"
	SetNoWrap(noWrap bool)

	// RowSpan prop Sets or retrieves how many rows in a table the cell should span.
	// js:"rowSpan,rewrite=rowspan"
	RowSpan() (rowSpan uint)

	// RowSpan prop Sets or retrieves how many rows in a table the cell should span.
	// js:"setrowSpan,rewrite=setrowspan"
	SetRowSpan(rowSpan uint)

	// Scope prop Sets or retrieves the group of cells in a table to which the object's information applies.
	// js:"scope,rewrite=scope"
	Scope() (scope string)

	// Scope prop Sets or retrieves the group of cells in a table to which the object's information applies.
	// js:"setscope,rewrite=setscope"
	SetScope(scope string)

	// Width prop Sets or retrieves the width of the object.
	// js:"width,rewrite=width"
	Width() (width string)

	// Width prop Sets or retrieves the width of the object.
	// js:"setwidth,rewrite=setwidth"
	SetWidth(width string)
}

// abbr prop Sets or retrieves abbreviated text for the object.
func abbr() (abbr string) {
	js.Rewrite("$<.abbr")
	return abbr
}

// setabbr prop Sets or retrieves abbreviated text for the object.
func setabbr(abbr string) {
	js.Rewrite("$<.abbr = abbr")
}

// align prop Sets or retrieves how the object is aligned with adjacent text.
func align() (align string) {
	js.Rewrite("$<.align")
	return align
}

// setalign prop Sets or retrieves how the object is aligned with adjacent text.
func setalign(align string) {
	js.Rewrite("$<.align = align")
}

// axis prop Sets or retrieves a comma-delimited list of conceptual categories associated with the object.
func axis() (axis string) {
	js.Rewrite("$<.axis")
	return axis
}

// setaxis prop Sets or retrieves a comma-delimited list of conceptual categories associated with the object.
func setaxis(axis string) {
	js.Rewrite("$<.axis = axis")
}

// bgcolor prop
func bgcolor() (bgColor interface{}) {
	js.Rewrite("$<.bgColor")
	return bgColor
}

// setbgcolor prop
func setbgcolor(bgColor interface{}) {
	js.Rewrite("$<.bgColor = bgColor")
}

// cellindex prop Retrieves the position of the object in the cells collection of a row.
func cellindex() (cellIndex int) {
	js.Rewrite("$<.cellIndex")
	return cellIndex
}

// colspan prop Sets or retrieves the number columns in the table that the object should span.
func colspan() (colSpan uint) {
	js.Rewrite("$<.colSpan")
	return colSpan
}

// setcolspan prop Sets or retrieves the number columns in the table that the object should span.
func setcolspan(colSpan uint) {
	js.Rewrite("$<.colSpan = colSpan")
}

// headers prop Sets or retrieves a list of header cells that provide information for the object.
func headers() (headers string) {
	js.Rewrite("$<.headers")
	return headers
}

// setheaders prop Sets or retrieves a list of header cells that provide information for the object.
func setheaders(headers string) {
	js.Rewrite("$<.headers = headers")
}

// height prop Sets or retrieves the height of the object.
func height() (height interface{}) {
	js.Rewrite("$<.height")
	return height
}

// setheight prop Sets or retrieves the height of the object.
func setheight(height interface{}) {
	js.Rewrite("$<.height = height")
}

// nowrap prop Sets or retrieves whether the browser automatically performs wordwrap.
func nowrap() (noWrap bool) {
	js.Rewrite("$<.noWrap")
	return noWrap
}

// setnowrap prop Sets or retrieves whether the browser automatically performs wordwrap.
func setnowrap(noWrap bool) {
	js.Rewrite("$<.noWrap = noWrap")
}

// rowspan prop Sets or retrieves how many rows in a table the cell should span.
func rowspan() (rowSpan uint) {
	js.Rewrite("$<.rowSpan")
	return rowSpan
}

// setrowspan prop Sets or retrieves how many rows in a table the cell should span.
func setrowspan(rowSpan uint) {
	js.Rewrite("$<.rowSpan = rowSpan")
}

// scope prop Sets or retrieves the group of cells in a table to which the object's information applies.
func scope() (scope string) {
	js.Rewrite("$<.scope")
	return scope
}

// setscope prop Sets or retrieves the group of cells in a table to which the object's information applies.
func setscope(scope string) {
	js.Rewrite("$<.scope = scope")
}

// width prop Sets or retrieves the width of the object.
func width() (width string) {
	js.Rewrite("$<.width")
	return width
}

// setwidth prop Sets or retrieves the width of the object.
func setwidth(width string) {
	js.Rewrite("$<.width = width")
}
