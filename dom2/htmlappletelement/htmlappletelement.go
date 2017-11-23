package htmlappletelement

import (
	"github.com/matthewmueller/golly/dom2/htmlformelement"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLAppletElement,omit"
type HTMLAppletElement struct {
	window.HTMLElement
}

// Align
func (*HTMLAppletElement) Align() (align string) {
	js.Rewrite("$<.Align")
	return align
}

// Align
func (*HTMLAppletElement) SetAlign(align string) {
	js.Rewrite("$<.Align = $1", align)
}

// Alt Sets or retrieves a text alternative to the graphic.
func (*HTMLAppletElement) Alt() (alt string) {
	js.Rewrite("$<.Alt")
	return alt
}

// Alt Sets or retrieves a text alternative to the graphic.
func (*HTMLAppletElement) SetAlt(alt string) {
	js.Rewrite("$<.Alt = $1", alt)
}

// AltHTML Gets or sets the optional alternative HTML script to execute if the object fails to load.
func (*HTMLAppletElement) AltHTML() (altHtml string) {
	js.Rewrite("$<.AltHTML")
	return altHtml
}

// AltHTML Gets or sets the optional alternative HTML script to execute if the object fails to load.
func (*HTMLAppletElement) SetAltHTML(altHtml string) {
	js.Rewrite("$<.AltHTML = $1", altHtml)
}

// Archive Sets or retrieves a character string that can be used to implement your own archive functionality for the object.
func (*HTMLAppletElement) Archive() (archive string) {
	js.Rewrite("$<.Archive")
	return archive
}

// Archive Sets or retrieves a character string that can be used to implement your own archive functionality for the object.
func (*HTMLAppletElement) SetArchive(archive string) {
	js.Rewrite("$<.Archive = $1", archive)
}

// BaseHref Retrieves a string of the URL where the object tag can be found. This is often the href of the document that the object is in, or the value set by a base element.
func (*HTMLAppletElement) BaseHref() (BaseHref string) {
	js.Rewrite("$<.BaseHref")
	return BaseHref
}

// Border
func (*HTMLAppletElement) Border() (border string) {
	js.Rewrite("$<.Border")
	return border
}

// Border
func (*HTMLAppletElement) SetBorder(border string) {
	js.Rewrite("$<.Border = $1", border)
}

// Code
func (*HTMLAppletElement) Code() (code string) {
	js.Rewrite("$<.Code")
	return code
}

// Code
func (*HTMLAppletElement) SetCode(code string) {
	js.Rewrite("$<.Code = $1", code)
}

// CodeBase Sets or retrieves the URL of the component.
func (*HTMLAppletElement) CodeBase() (codeBase string) {
	js.Rewrite("$<.CodeBase")
	return codeBase
}

// CodeBase Sets or retrieves the URL of the component.
func (*HTMLAppletElement) SetCodeBase(codeBase string) {
	js.Rewrite("$<.CodeBase = $1", codeBase)
}

// CodeType Sets or retrieves the Internet media type for the code associated with the object.
func (*HTMLAppletElement) CodeType() (codeType string) {
	js.Rewrite("$<.CodeType")
	return codeType
}

// CodeType Sets or retrieves the Internet media type for the code associated with the object.
func (*HTMLAppletElement) SetCodeType(codeType string) {
	js.Rewrite("$<.CodeType = $1", codeType)
}

// ContentDocument Address of a pointer to the document this page or frame contains. If there is no document, then null will be returned.
func (*HTMLAppletElement) ContentDocument() (contentDocument window.Document) {
	js.Rewrite("$<.ContentDocument")
	return contentDocument
}

// Data Sets or retrieves the URL that references the data of the object.
func (*HTMLAppletElement) Data() (data string) {
	js.Rewrite("$<.Data")
	return data
}

// Data Sets or retrieves the URL that references the data of the object.
func (*HTMLAppletElement) SetData(data string) {
	js.Rewrite("$<.Data = $1", data)
}

// Declare Sets or retrieves a character string that can be used to implement your own declare functionality for the object.
func (*HTMLAppletElement) Declare() (declare bool) {
	js.Rewrite("$<.Declare")
	return declare
}

// Declare Sets or retrieves a character string that can be used to implement your own declare functionality for the object.
func (*HTMLAppletElement) SetDeclare(declare bool) {
	js.Rewrite("$<.Declare = $1", declare)
}

// Form
func (*HTMLAppletElement) Form() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.Form")
	return form
}

// Height Sets or retrieves the height of the object.
func (*HTMLAppletElement) Height() (height string) {
	js.Rewrite("$<.Height")
	return height
}

// Height Sets or retrieves the height of the object.
func (*HTMLAppletElement) SetHeight(height string) {
	js.Rewrite("$<.Height = $1", height)
}

// Hspace
func (*HTMLAppletElement) Hspace() (hspace int) {
	js.Rewrite("$<.Hspace")
	return hspace
}

// Hspace
func (*HTMLAppletElement) SetHspace(hspace int) {
	js.Rewrite("$<.Hspace = $1", hspace)
}

// Name Sets or retrieves the shape of the object.
func (*HTMLAppletElement) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}

// Name Sets or retrieves the shape of the object.
func (*HTMLAppletElement) SetName(name string) {
	js.Rewrite("$<.Name = $1", name)
}

// Object
func (*HTMLAppletElement) Object() (object string) {
	js.Rewrite("$<.Object")
	return object
}

// Object
func (*HTMLAppletElement) SetObject(object string) {
	js.Rewrite("$<.Object = $1", object)
}

// Standby Sets or retrieves a message to be displayed while an object is loading.
func (*HTMLAppletElement) Standby() (standby string) {
	js.Rewrite("$<.Standby")
	return standby
}

// Standby Sets or retrieves a message to be displayed while an object is loading.
func (*HTMLAppletElement) SetStandby(standby string) {
	js.Rewrite("$<.Standby = $1", standby)
}

// Type Returns the content type of the object.
func (*HTMLAppletElement) Type() (kind string) {
	js.Rewrite("$<.Type")
	return kind
}

// Type Returns the content type of the object.
func (*HTMLAppletElement) SetType(kind string) {
	js.Rewrite("$<.Type = $1", kind)
}

// UseMap Sets or retrieves the URL, often with a bookmark extension (#name), to use as a client-side image map.
func (*HTMLAppletElement) UseMap() (useMap string) {
	js.Rewrite("$<.UseMap")
	return useMap
}

// UseMap Sets or retrieves the URL, often with a bookmark extension (#name), to use as a client-side image map.
func (*HTMLAppletElement) SetUseMap(useMap string) {
	js.Rewrite("$<.UseMap = $1", useMap)
}

// Vspace
func (*HTMLAppletElement) Vspace() (vspace int) {
	js.Rewrite("$<.Vspace")
	return vspace
}

// Vspace
func (*HTMLAppletElement) SetVspace(vspace int) {
	js.Rewrite("$<.Vspace = $1", vspace)
}

// Width
func (*HTMLAppletElement) Width() (width int) {
	js.Rewrite("$<.Width")
	return width
}

// Width
func (*HTMLAppletElement) SetWidth(width int) {
	js.Rewrite("$<.Width = $1", width)
}
