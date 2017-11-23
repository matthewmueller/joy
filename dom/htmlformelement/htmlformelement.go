package htmlformelement

import (
	"github.com/matthewmueller/golly/dom2/htmlformcontrolscollection"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLFormElement struct
// js:"HTMLFormElement,omit"
type HTMLFormElement struct {
	window.HTMLElement
}

// CheckValidity fn Returns whether a form will validate when it is submitted, without having to submit it.
func (*HTMLFormElement) CheckValidity() (b bool) {
	js.Rewrite("$<.checkValidity()")
	return b
}

// Item fn Retrieves a form object or an object from an elements collection.
//     * @param name Variant of type Number or String that specifies the object or collection to retrieve. If this parameter is a Number, it is the zero-based index of the object. If this parameter is a string, all objects with matching name or id properties are retrieved, and a collection is returned if more than one match is made.
//     * @param index Variant of type Number that specifies the zero-based index of the object to retrieve when a collection is returned.
func (*HTMLFormElement) Item(name *interface{}, index *interface{}) (i interface{}) {
	js.Rewrite("$<.item($1, $2)", name, index)
	return i
}

// NamedItem fn Retrieves a form object or an object from an elements collection.
func (*HTMLFormElement) NamedItem(name string) (i interface{}) {
	js.Rewrite("$<.namedItem($1)", name)
	return i
}

// Reset fn Fires when the user resets a form.
func (*HTMLFormElement) Reset() {
	js.Rewrite("$<.reset()")
}

// Submit fn Fires when a FORM is about to be submitted.
func (*HTMLFormElement) Submit() {
	js.Rewrite("$<.submit()")
}

// AcceptCharset prop Sets or retrieves a list of character encodings for input data that must be accepted by the server processing the form.
func (*HTMLFormElement) AcceptCharset() (acceptCharset string) {
	js.Rewrite("$<.acceptCharset")
	return acceptCharset
}

// AcceptCharset prop Sets or retrieves a list of character encodings for input data that must be accepted by the server processing the form.
func (*HTMLFormElement) SetAcceptCharset(acceptCharset string) {
	js.Rewrite("$<.acceptCharset = $1", acceptCharset)
}

// Action prop Sets or retrieves the URL to which the form content is sent for processing.
func (*HTMLFormElement) Action() (action string) {
	js.Rewrite("$<.action")
	return action
}

// Action prop Sets or retrieves the URL to which the form content is sent for processing.
func (*HTMLFormElement) SetAction(action string) {
	js.Rewrite("$<.action = $1", action)
}

// Autocomplete prop Specifies whether autocomplete is applied to an editable text field.
func (*HTMLFormElement) Autocomplete() (autocomplete string) {
	js.Rewrite("$<.autocomplete")
	return autocomplete
}

// Autocomplete prop Specifies whether autocomplete is applied to an editable text field.
func (*HTMLFormElement) SetAutocomplete(autocomplete string) {
	js.Rewrite("$<.autocomplete = $1", autocomplete)
}

// Elements prop Retrieves a collection, in source order, of all controls in a given form.
func (*HTMLFormElement) Elements() (elements *htmlformcontrolscollection.HTMLFormControlsCollection) {
	js.Rewrite("$<.elements")
	return elements
}

// Encoding prop Sets or retrieves the MIME encoding for the form.
func (*HTMLFormElement) Encoding() (encoding string) {
	js.Rewrite("$<.encoding")
	return encoding
}

// Encoding prop Sets or retrieves the MIME encoding for the form.
func (*HTMLFormElement) SetEncoding(encoding string) {
	js.Rewrite("$<.encoding = $1", encoding)
}

// Enctype prop Sets or retrieves the encoding type for the form.
func (*HTMLFormElement) Enctype() (enctype string) {
	js.Rewrite("$<.enctype")
	return enctype
}

// Enctype prop Sets or retrieves the encoding type for the form.
func (*HTMLFormElement) SetEnctype(enctype string) {
	js.Rewrite("$<.enctype = $1", enctype)
}

// Length prop Sets or retrieves the number of objects in a collection.
func (*HTMLFormElement) Length() (length int) {
	js.Rewrite("$<.length")
	return length
}

// Method prop Sets or retrieves how to send the form data to the server.
func (*HTMLFormElement) Method() (method string) {
	js.Rewrite("$<.method")
	return method
}

// Method prop Sets or retrieves how to send the form data to the server.
func (*HTMLFormElement) SetMethod(method string) {
	js.Rewrite("$<.method = $1", method)
}

// Name prop Sets or retrieves the name of the object.
func (*HTMLFormElement) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// Name prop Sets or retrieves the name of the object.
func (*HTMLFormElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

// NoValidate prop Designates a form that is not validated when submitted.
func (*HTMLFormElement) NoValidate() (noValidate bool) {
	js.Rewrite("$<.noValidate")
	return noValidate
}

// NoValidate prop Designates a form that is not validated when submitted.
func (*HTMLFormElement) SetNoValidate(noValidate bool) {
	js.Rewrite("$<.noValidate = $1", noValidate)
}

// Target prop Sets or retrieves the window or frame at which to target content.
func (*HTMLFormElement) Target() (target string) {
	js.Rewrite("$<.target")
	return target
}

// Target prop Sets or retrieves the window or frame at which to target content.
func (*HTMLFormElement) SetTarget(target string) {
	js.Rewrite("$<.target = $1", target)
}
