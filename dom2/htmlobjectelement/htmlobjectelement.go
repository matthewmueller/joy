package htmlobjectelement

import (
	"github.com/matthewmueller/golly/dom2/htmlformelement"
	"github.com/matthewmueller/golly/dom2/validitystate"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLObjectElement struct
// js:"HTMLObjectElement,omit"
type HTMLObjectElement struct {
	window.HTMLElement
}

// GetSVGDocument
func (*HTMLObjectElement) GetSVGDocument() (w window.Document) {
	js.Rewrite("$<.GetSVGDocument()")
	return w
}

// CheckValidity Returns whether a form will validate when it is submitted, without having to submit it.
func (*HTMLObjectElement) CheckValidity() (b bool) {
	js.Rewrite("$<.CheckValidity()")
	return b
}

// SetCustomValidity Sets a custom error message that is displayed when a form is submitted.
//     * @param error Sets a custom error message that is displayed when a form is submitted.
func (*HTMLObjectElement) SetCustomValidity(err string) {
	js.Rewrite("$<.SetCustomValidity($1)", err)
}

// Align
func (*HTMLObjectElement) Align() (align string) {
	js.Rewrite("$<.Align")
	return align
}

// Align
func (*HTMLObjectElement) SetAlign(align string) {
	js.Rewrite("$<.Align = $1", align)
}

// Alt Sets or retrieves a text alternative to the graphic.
func (*HTMLObjectElement) Alt() (alt string) {
	js.Rewrite("$<.Alt")
	return alt
}

// Alt Sets or retrieves a text alternative to the graphic.
func (*HTMLObjectElement) SetAlt(alt string) {
	js.Rewrite("$<.Alt = $1", alt)
}

// AltHTML Gets or sets the optional alternative HTML script to execute if the object fails to load.
func (*HTMLObjectElement) AltHTML() (altHtml string) {
	js.Rewrite("$<.AltHTML")
	return altHtml
}

// AltHTML Gets or sets the optional alternative HTML script to execute if the object fails to load.
func (*HTMLObjectElement) SetAltHTML(altHtml string) {
	js.Rewrite("$<.AltHTML = $1", altHtml)
}

// Archive Sets or retrieves a character string that can be used to implement your own archive functionality for the object.
func (*HTMLObjectElement) Archive() (archive string) {
	js.Rewrite("$<.Archive")
	return archive
}

// Archive Sets or retrieves a character string that can be used to implement your own archive functionality for the object.
func (*HTMLObjectElement) SetArchive(archive string) {
	js.Rewrite("$<.Archive = $1", archive)
}

// BaseHref Retrieves a string of the URL where the object tag can be found. This is often the href of the document that the object is in, or the value set by a base element.
func (*HTMLObjectElement) BaseHref() (BaseHref string) {
	js.Rewrite("$<.BaseHref")
	return BaseHref
}

// Border
func (*HTMLObjectElement) Border() (border string) {
	js.Rewrite("$<.Border")
	return border
}

// Border
func (*HTMLObjectElement) SetBorder(border string) {
	js.Rewrite("$<.Border = $1", border)
}

// Code Sets or retrieves the URL of the file containing the compiled Java class.
func (*HTMLObjectElement) Code() (code string) {
	js.Rewrite("$<.Code")
	return code
}

// Code Sets or retrieves the URL of the file containing the compiled Java class.
func (*HTMLObjectElement) SetCode(code string) {
	js.Rewrite("$<.Code = $1", code)
}

// CodeBase Sets or retrieves the URL of the component.
func (*HTMLObjectElement) CodeBase() (codeBase string) {
	js.Rewrite("$<.CodeBase")
	return codeBase
}

// CodeBase Sets or retrieves the URL of the component.
func (*HTMLObjectElement) SetCodeBase(codeBase string) {
	js.Rewrite("$<.CodeBase = $1", codeBase)
}

// CodeType Sets or retrieves the Internet media type for the code associated with the object.
func (*HTMLObjectElement) CodeType() (codeType string) {
	js.Rewrite("$<.CodeType")
	return codeType
}

// CodeType Sets or retrieves the Internet media type for the code associated with the object.
func (*HTMLObjectElement) SetCodeType(codeType string) {
	js.Rewrite("$<.CodeType = $1", codeType)
}

// ContentDocument Retrieves the document object of the page or frame.
func (*HTMLObjectElement) ContentDocument() (contentDocument window.Document) {
	js.Rewrite("$<.ContentDocument")
	return contentDocument
}

// Data Sets or retrieves the URL that references the data of the object.
func (*HTMLObjectElement) Data() (data string) {
	js.Rewrite("$<.Data")
	return data
}

// Data Sets or retrieves the URL that references the data of the object.
func (*HTMLObjectElement) SetData(data string) {
	js.Rewrite("$<.Data = $1", data)
}

// Declare
func (*HTMLObjectElement) Declare() (declare bool) {
	js.Rewrite("$<.Declare")
	return declare
}

// Declare
func (*HTMLObjectElement) SetDeclare(declare bool) {
	js.Rewrite("$<.Declare = $1", declare)
}

// Form Retrieves a reference to the form that the object is embedded in.
func (*HTMLObjectElement) Form() (form *htmlformelement.HTMLFormElement) {
	js.Rewrite("$<.Form")
	return form
}

// Height Sets or retrieves the height of the object.
func (*HTMLObjectElement) Height() (height string) {
	js.Rewrite("$<.Height")
	return height
}

// Height Sets or retrieves the height of the object.
func (*HTMLObjectElement) SetHeight(height string) {
	js.Rewrite("$<.Height = $1", height)
}

// Hspace
func (*HTMLObjectElement) Hspace() (hspace int) {
	js.Rewrite("$<.Hspace")
	return hspace
}

// Hspace
func (*HTMLObjectElement) SetHspace(hspace int) {
	js.Rewrite("$<.Hspace = $1", hspace)
}

// MsPlayToDisabled Gets or sets whether the DLNA PlayTo device is available.
func (*HTMLObjectElement) MsPlayToDisabled() (msPlayToDisabled bool) {
	js.Rewrite("$<.MsPlayToDisabled")
	return msPlayToDisabled
}

// MsPlayToDisabled Gets or sets whether the DLNA PlayTo device is available.
func (*HTMLObjectElement) SetMsPlayToDisabled(msPlayToDisabled bool) {
	js.Rewrite("$<.MsPlayToDisabled = $1", msPlayToDisabled)
}

// MsPlayToPreferredSourceURI Gets or sets the path to the preferred media source. This enables the Play To target device to stream the media content, which can be DRM protected, from a different location, such as a cloud media server.
func (*HTMLObjectElement) MsPlayToPreferredSourceURI() (msPlayToPreferredSourceUri string) {
	js.Rewrite("$<.MsPlayToPreferredSourceURI")
	return msPlayToPreferredSourceUri
}

// MsPlayToPreferredSourceURI Gets or sets the path to the preferred media source. This enables the Play To target device to stream the media content, which can be DRM protected, from a different location, such as a cloud media server.
func (*HTMLObjectElement) SetMsPlayToPreferredSourceURI(msPlayToPreferredSourceUri string) {
	js.Rewrite("$<.MsPlayToPreferredSourceURI = $1", msPlayToPreferredSourceUri)
}

// MsPlayToPrimary Gets or sets the primary DLNA PlayTo device.
func (*HTMLObjectElement) MsPlayToPrimary() (msPlayToPrimary bool) {
	js.Rewrite("$<.MsPlayToPrimary")
	return msPlayToPrimary
}

// MsPlayToPrimary Gets or sets the primary DLNA PlayTo device.
func (*HTMLObjectElement) SetMsPlayToPrimary(msPlayToPrimary bool) {
	js.Rewrite("$<.MsPlayToPrimary = $1", msPlayToPrimary)
}

// MsPlayToSource Gets the source associated with the media element for use by the PlayToManager.
func (*HTMLObjectElement) MsPlayToSource() (msPlayToSource interface{}) {
	js.Rewrite("$<.MsPlayToSource")
	return msPlayToSource
}

// Name Sets or retrieves the name of the object.
func (*HTMLObjectElement) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}

// Name Sets or retrieves the name of the object.
func (*HTMLObjectElement) SetName(name string) {
	js.Rewrite("$<.Name = $1", name)
}

// ReadyState
func (*HTMLObjectElement) ReadyState() (readyState uint8) {
	js.Rewrite("$<.ReadyState")
	return readyState
}

// Standby Sets or retrieves a message to be displayed while an object is loading.
func (*HTMLObjectElement) Standby() (standby string) {
	js.Rewrite("$<.Standby")
	return standby
}

// Standby Sets or retrieves a message to be displayed while an object is loading.
func (*HTMLObjectElement) SetStandby(standby string) {
	js.Rewrite("$<.Standby = $1", standby)
}

// Type Sets or retrieves the MIME type of the object.
func (*HTMLObjectElement) Type() (kind string) {
	js.Rewrite("$<.Type")
	return kind
}

// Type Sets or retrieves the MIME type of the object.
func (*HTMLObjectElement) SetType(kind string) {
	js.Rewrite("$<.Type = $1", kind)
}

// UseMap Sets or retrieves the URL, often with a bookmark extension (#name), to use as a client-side image map.
func (*HTMLObjectElement) UseMap() (useMap string) {
	js.Rewrite("$<.UseMap")
	return useMap
}

// UseMap Sets or retrieves the URL, often with a bookmark extension (#name), to use as a client-side image map.
func (*HTMLObjectElement) SetUseMap(useMap string) {
	js.Rewrite("$<.UseMap = $1", useMap)
}

// ValidationMessage Returns the error message that would be displayed if the user submits the form, or an empty string if no error message. It also triggers the standard error message, such as "this is a required field". The result is that the user sees validation messages without actually submitting.
func (*HTMLObjectElement) ValidationMessage() (validationMessage string) {
	js.Rewrite("$<.ValidationMessage")
	return validationMessage
}

// Validity Returns a  ValidityState object that represents the validity states of an element.
func (*HTMLObjectElement) Validity() (validity *validitystate.ValidityState) {
	js.Rewrite("$<.Validity")
	return validity
}

// Vspace
func (*HTMLObjectElement) Vspace() (vspace int) {
	js.Rewrite("$<.Vspace")
	return vspace
}

// Vspace
func (*HTMLObjectElement) SetVspace(vspace int) {
	js.Rewrite("$<.Vspace = $1", vspace)
}

// Width Sets or retrieves the width of the object.
func (*HTMLObjectElement) Width() (width string) {
	js.Rewrite("$<.Width")
	return width
}

// Width Sets or retrieves the width of the object.
func (*HTMLObjectElement) SetWidth(width string) {
	js.Rewrite("$<.Width = $1", width)
}

// WillValidate Returns whether an element will successfully validate based on forms validation rules and constraints.
func (*HTMLObjectElement) WillValidate() (willValidate bool) {
	js.Rewrite("$<.WillValidate")
	return willValidate
}
