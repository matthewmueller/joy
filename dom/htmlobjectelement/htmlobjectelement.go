package htmlobjectelement

import (
	"github.com/matthewmueller/golly/dom/htmlformelement"
	"github.com/matthewmueller/golly/dom/validitystate"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLObjectElement struct
// js:"HTMLObjectElement,omit"
type HTMLObjectElement struct {
	window.HTMLElement
}

// GetSVGDocument fn
func (*HTMLObjectElement) GetSVGDocument() (w window.Document) {
	js.Rewrite("$<.getSVGDocument()")
	return w
}

// CheckValidity fn Returns whether a form will validate when it is submitted, without having to submit it.
func (*HTMLObjectElement) CheckValidity() (b bool) {
	js.Rewrite("$<.checkValidity()")
	return b
}

// SetCustomValidity fn Sets a custom error message that is displayed when a form is submitted.
//     * @param error Sets a custom error message that is displayed when a form is submitted.
func (*HTMLObjectElement) SetCustomValidity(err string) {
	js.Rewrite("$<.setCustomValidity($1)", err)
}

// Align prop
func (*HTMLObjectElement) Align() (align string) {
	js.Rewrite("$<.align")
	return align
}

// Align prop
func (*HTMLObjectElement) SetAlign(align string) {
	js.Rewrite("$<.align = $1", align)
}

// Alt prop Sets or retrieves a text alternative to the graphic.
func (*HTMLObjectElement) Alt() (alt string) {
	js.Rewrite("$<.alt")
	return alt
}

// Alt prop Sets or retrieves a text alternative to the graphic.
func (*HTMLObjectElement) SetAlt(alt string) {
	js.Rewrite("$<.alt = $1", alt)
}

// AltHTML prop Gets or sets the optional alternative HTML script to execute if the object fails to load.
func (*HTMLObjectElement) AltHTML() (altHtml string) {
	js.Rewrite("$<.altHtml")
	return altHtml
}

// AltHTML prop Gets or sets the optional alternative HTML script to execute if the object fails to load.
func (*HTMLObjectElement) SetAltHTML(altHtml string) {
	js.Rewrite("$<.altHtml = $1", altHtml)
}

// Archive prop Sets or retrieves a character string that can be used to implement your own archive functionality for the object.
func (*HTMLObjectElement) Archive() (archive string) {
	js.Rewrite("$<.archive")
	return archive
}

// Archive prop Sets or retrieves a character string that can be used to implement your own archive functionality for the object.
func (*HTMLObjectElement) SetArchive(archive string) {
	js.Rewrite("$<.archive = $1", archive)
}

// BaseHref prop Retrieves a string of the URL where the object tag can be found. This is often the href of the document that the object is in, or the value set by a base element.
func (*HTMLObjectElement) BaseHref() (BaseHref string) {
	js.Rewrite("$<.BaseHref")
	return BaseHref
}

// Border prop
func (*HTMLObjectElement) Border() (border string) {
	js.Rewrite("$<.border")
	return border
}

// Border prop
func (*HTMLObjectElement) SetBorder(border string) {
	js.Rewrite("$<.border = $1", border)
}

// Code prop Sets or retrieves the URL of the file containing the compiled Java class.
func (*HTMLObjectElement) Code() (code string) {
	js.Rewrite("$<.code")
	return code
}

// Code prop Sets or retrieves the URL of the file containing the compiled Java class.
func (*HTMLObjectElement) SetCode(code string) {
	js.Rewrite("$<.code = $1", code)
}

// CodeBase prop Sets or retrieves the URL of the component.
func (*HTMLObjectElement) CodeBase() (codeBase string) {
	js.Rewrite("$<.codeBase")
	return codeBase
}

// CodeBase prop Sets or retrieves the URL of the component.
func (*HTMLObjectElement) SetCodeBase(codeBase string) {
	js.Rewrite("$<.codeBase = $1", codeBase)
}

// CodeType prop Sets or retrieves the Internet media type for the code associated with the object.
func (*HTMLObjectElement) CodeType() (codeType string) {
	js.Rewrite("$<.codeType")
	return codeType
}

// CodeType prop Sets or retrieves the Internet media type for the code associated with the object.
func (*HTMLObjectElement) SetCodeType(codeType string) {
	js.Rewrite("$<.codeType = $1", codeType)
}

// ContentDocument prop Retrieves the document object of the page or frame.
func (*HTMLObjectElement) ContentDocument() (contentDocument window.Document) {
	js.Rewrite("$<.contentDocument")
	return contentDocument
}

// Data prop Sets or retrieves the URL that references the data of the object.
func (*HTMLObjectElement) Data() (data string) {
	js.Rewrite("$<.data")
	return data
}

// Data prop Sets or retrieves the URL that references the data of the object.
func (*HTMLObjectElement) SetData(data string) {
	js.Rewrite("$<.data = $1", data)
}

// Declare prop
func (*HTMLObjectElement) Declare() (declare bool) {
	js.Rewrite("$<.declare")
	return declare
}

// Declare prop
func (*HTMLObjectElement) SetDeclare(declare bool) {
	js.Rewrite("$<.declare = $1", declare)
}

// Form prop Retrieves a reference to the form that the object is embedded in.
func (*HTMLObjectElement) Form() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.form")
	return form
}

// Height prop Sets or retrieves the height of the object.
func (*HTMLObjectElement) Height() (height string) {
	js.Rewrite("$<.height")
	return height
}

// Height prop Sets or retrieves the height of the object.
func (*HTMLObjectElement) SetHeight(height string) {
	js.Rewrite("$<.height = $1", height)
}

// Hspace prop
func (*HTMLObjectElement) Hspace() (hspace int) {
	js.Rewrite("$<.hspace")
	return hspace
}

// Hspace prop
func (*HTMLObjectElement) SetHspace(hspace int) {
	js.Rewrite("$<.hspace = $1", hspace)
}

// MsPlayToDisabled prop Gets or sets whether the DLNA PlayTo device is available.
func (*HTMLObjectElement) MsPlayToDisabled() (msPlayToDisabled bool) {
	js.Rewrite("$<.msPlayToDisabled")
	return msPlayToDisabled
}

// MsPlayToDisabled prop Gets or sets whether the DLNA PlayTo device is available.
func (*HTMLObjectElement) SetMsPlayToDisabled(msPlayToDisabled bool) {
	js.Rewrite("$<.msPlayToDisabled = $1", msPlayToDisabled)
}

// MsPlayToPreferredSourceURI prop Gets or sets the path to the preferred media source. This enables the Play To target device to stream the media content, which can be DRM protected, from a different location, such as a cloud media server.
func (*HTMLObjectElement) MsPlayToPreferredSourceURI() (msPlayToPreferredSourceUri string) {
	js.Rewrite("$<.msPlayToPreferredSourceUri")
	return msPlayToPreferredSourceUri
}

// MsPlayToPreferredSourceURI prop Gets or sets the path to the preferred media source. This enables the Play To target device to stream the media content, which can be DRM protected, from a different location, such as a cloud media server.
func (*HTMLObjectElement) SetMsPlayToPreferredSourceURI(msPlayToPreferredSourceUri string) {
	js.Rewrite("$<.msPlayToPreferredSourceUri = $1", msPlayToPreferredSourceUri)
}

// MsPlayToPrimary prop Gets or sets the primary DLNA PlayTo device.
func (*HTMLObjectElement) MsPlayToPrimary() (msPlayToPrimary bool) {
	js.Rewrite("$<.msPlayToPrimary")
	return msPlayToPrimary
}

// MsPlayToPrimary prop Gets or sets the primary DLNA PlayTo device.
func (*HTMLObjectElement) SetMsPlayToPrimary(msPlayToPrimary bool) {
	js.Rewrite("$<.msPlayToPrimary = $1", msPlayToPrimary)
}

// MsPlayToSource prop Gets the source associated with the media element for use by the PlayToManager.
func (*HTMLObjectElement) MsPlayToSource() (msPlayToSource interface{}) {
	js.Rewrite("$<.msPlayToSource")
	return msPlayToSource
}

// Name prop Sets or retrieves the name of the object.
func (*HTMLObjectElement) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// Name prop Sets or retrieves the name of the object.
func (*HTMLObjectElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

// ReadyState prop
func (*HTMLObjectElement) ReadyState() (readyState uint8) {
	js.Rewrite("$<.readyState")
	return readyState
}

// Standby prop Sets or retrieves a message to be displayed while an object is loading.
func (*HTMLObjectElement) Standby() (standby string) {
	js.Rewrite("$<.standby")
	return standby
}

// Standby prop Sets or retrieves a message to be displayed while an object is loading.
func (*HTMLObjectElement) SetStandby(standby string) {
	js.Rewrite("$<.standby = $1", standby)
}

// Type prop Sets or retrieves the MIME type of the object.
func (*HTMLObjectElement) Type() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

// Type prop Sets or retrieves the MIME type of the object.
func (*HTMLObjectElement) SetType(kind string) {
	js.Rewrite("$<.type = $1", kind)
}

// UseMap prop Sets or retrieves the URL, often with a bookmark extension (#name), to use as a client-side image map.
func (*HTMLObjectElement) UseMap() (useMap string) {
	js.Rewrite("$<.useMap")
	return useMap
}

// UseMap prop Sets or retrieves the URL, often with a bookmark extension (#name), to use as a client-side image map.
func (*HTMLObjectElement) SetUseMap(useMap string) {
	js.Rewrite("$<.useMap = $1", useMap)
}

// ValidationMessage prop Returns the error message that would be displayed if the user submits the form, or an empty string if no error message. It also triggers the standard error message, such as "this is a required field". The result is that the user sees validation messages without actually submitting.
func (*HTMLObjectElement) ValidationMessage() (validationMessage string) {
	js.Rewrite("$<.validationMessage")
	return validationMessage
}

// Validity prop Returns a  ValidityState object that represents the validity states of an element.
func (*HTMLObjectElement) Validity() (validity *validitystate.ValidityState) {
	js.Rewrite("$<.validity")
	return validity
}

// Vspace prop
func (*HTMLObjectElement) Vspace() (vspace int) {
	js.Rewrite("$<.vspace")
	return vspace
}

// Vspace prop
func (*HTMLObjectElement) SetVspace(vspace int) {
	js.Rewrite("$<.vspace = $1", vspace)
}

// Width prop Sets or retrieves the width of the object.
func (*HTMLObjectElement) Width() (width string) {
	js.Rewrite("$<.width")
	return width
}

// Width prop Sets or retrieves the width of the object.
func (*HTMLObjectElement) SetWidth(width string) {
	js.Rewrite("$<.width = $1", width)
}

// WillValidate prop Returns whether an element will successfully validate based on forms validation rules and constraints.
func (*HTMLObjectElement) WillValidate() (willValidate bool) {
	js.Rewrite("$<.willValidate")
	return willValidate
}
