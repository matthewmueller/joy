package htmlappletelement

import (
	"github.com/matthewmueller/golly/dom2/htmlformelement"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLAppletElement struct
// js:"HTMLAppletElement,omit"
type HTMLAppletElement struct {
	window.HTMLElement
}

// Align prop
func (*HTMLAppletElement) Align() (align string) {
	js.Rewrite("$<.align")
	return align
}

// Align prop
func (*HTMLAppletElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

// Alt prop Sets or retrieves a text alternative to the graphic.
func (*HTMLAppletElement) Alt() (alt string) {
	js.Rewrite("$<.alt")
	return alt
}

// Alt prop Sets or retrieves a text alternative to the graphic.
func (*HTMLAppletElement) SetAlt(alt string) {
	js.Rewrite("$<.alt = $1", alt)
}

// AltHTML prop Gets or sets the optional alternative HTML script to execute if the object fails to load.
func (*HTMLAppletElement) AltHTML() (altHtml string) {
	js.Rewrite("$<.altHtml")
	return altHtml
}

// AltHTML prop Gets or sets the optional alternative HTML script to execute if the object fails to load.
func (*HTMLAppletElement) SetAltHTML(altHtml string) {
	js.Rewrite("$<.altHtml = $1", altHtml)
}

// Archive prop Sets or retrieves a character string that can be used to implement your own archive functionality for the object.
func (*HTMLAppletElement) Archive() (archive string) {
	js.Rewrite("$<.archive")
	return archive
}

// Archive prop Sets or retrieves a character string that can be used to implement your own archive functionality for the object.
func (*HTMLAppletElement) SetArchive(archive string) {
	js.Rewrite("$<.archive = $1", archive)
}

// BaseHref prop Retrieves a string of the URL where the object tag can be found. This is often the href of the document that the object is in, or the value set by a base element.
func (*HTMLAppletElement) BaseHref() (BaseHref string) {
	js.Rewrite("$<.BaseHref")
	return BaseHref
}

// Border prop
func (*HTMLAppletElement) Border() (border string) {
	js.Rewrite("$<.border")
	return border
}

// Border prop
func (*HTMLAppletElement) SetBorder(border string) {
	js.Rewrite("$<.border = $1", border)
}

// Code prop
func (*HTMLAppletElement) Code() (code string) {
	js.Rewrite("$<.code")
	return code
}

// Code prop
func (*HTMLAppletElement) SetCode(code string) {
	js.Rewrite("$<.code = $1", code)
}

// CodeBase prop Sets or retrieves the URL of the component.
func (*HTMLAppletElement) CodeBase() (codeBase string) {
	js.Rewrite("$<.codeBase")
	return codeBase
}

// CodeBase prop Sets or retrieves the URL of the component.
func (*HTMLAppletElement) SetCodeBase(codeBase string) {
	js.Rewrite("$<.codeBase = $1", codeBase)
}

// CodeType prop Sets or retrieves the Internet media type for the code associated with the object.
func (*HTMLAppletElement) CodeType() (codeType string) {
	js.Rewrite("$<.codeType")
	return codeType
}

// CodeType prop Sets or retrieves the Internet media type for the code associated with the object.
func (*HTMLAppletElement) SetCodeType(codeType string) {
	js.Rewrite("$<.codeType = $1", codeType)
}

// ContentDocument prop Address of a pointer to the document this page or frame contains. If there is no document, then null will be returned.
func (*HTMLAppletElement) ContentDocument() (contentDocument window.Document) {
	js.Rewrite("$<.contentDocument")
	return contentDocument
}

// Data prop Sets or retrieves the URL that references the data of the object.
func (*HTMLAppletElement) Data() (data string) {
	js.Rewrite("$<.data")
	return data
}

// Data prop Sets or retrieves the URL that references the data of the object.
func (*HTMLAppletElement) SetData(data string) {
	js.Rewrite("$<.data = $1", data)
}

// Declare prop Sets or retrieves a character string that can be used to implement your own declare functionality for the object.
func (*HTMLAppletElement) Declare() (declare bool) {
	js.Rewrite("$<.declare")
	return declare
}

// Declare prop Sets or retrieves a character string that can be used to implement your own declare functionality for the object.
func (*HTMLAppletElement) SetDeclare(declare bool) {
	js.Rewrite("$<.declare = $1", declare)
}

// Form prop
func (*HTMLAppletElement) Form() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

// Height prop Sets or retrieves the height of the object.
func (*HTMLAppletElement) Height() (height string) {
	js.Rewrite("$<.height")
	return height
}

// Height prop Sets or retrieves the height of the object.
func (*HTMLAppletElement) SetHeight(height string) {
	js.Rewrite("$<.height = $1", height)
}

// Hspace prop
func (*HTMLAppletElement) Hspace() (hspace int) {
	js.Rewrite("$<.hspace")
	return hspace
}

// Hspace prop
func (*HTMLAppletElement) SetHspace(hspace int) {
	js.Rewrite("$<.hspace = $1", hspace)
}

// Name prop Sets or retrieves the shape of the object.
func (*HTMLAppletElement) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// Name prop Sets or retrieves the shape of the object.
func (*HTMLAppletElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

// Object prop
func (*HTMLAppletElement) Object() (object string) {
	js.Rewrite("$<.object")
	return object
}

// Object prop
func (*HTMLAppletElement) SetObject(object string) {
	js.Rewrite("$<.object = $1", object)
}

// Standby prop Sets or retrieves a message to be displayed while an object is loading.
func (*HTMLAppletElement) Standby() (standby string) {
	js.Rewrite("$<.standby")
	return standby
}

// Standby prop Sets or retrieves a message to be displayed while an object is loading.
func (*HTMLAppletElement) SetStandby(standby string) {
	js.Rewrite("$<.standby = $1", standby)
}

// Type prop Returns the content type of the object.
func (*HTMLAppletElement) Type() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

// Type prop Returns the content type of the object.
func (*HTMLAppletElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}

// UseMap prop Sets or retrieves the URL, often with a bookmark extension (#name), to use as a client-side image map.
func (*HTMLAppletElement) UseMap() (useMap string) {
	js.Rewrite("$<.useMap")
	return useMap
}

// UseMap prop Sets or retrieves the URL, often with a bookmark extension (#name), to use as a client-side image map.
func (*HTMLAppletElement) SetUseMap(useMap string) {
	js.Rewrite("$<.useMap = $1", useMap)
}

// Vspace prop
func (*HTMLAppletElement) Vspace() (vspace int) {
	js.Rewrite("$<.vspace")
	return vspace
}

// Vspace prop
func (*HTMLAppletElement) SetVspace(vspace int) {
	js.Rewrite("$<.vspace = $1", vspace)
}

// Width prop
func (*HTMLAppletElement) Width() (width int) {
	js.Rewrite("$<.width")
	return width
}

// Width prop
func (*HTMLAppletElement) SetWidth(width int) {
	js.Rewrite("$<.width = $1", width)
}
