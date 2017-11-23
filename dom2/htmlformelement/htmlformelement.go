package htmlformelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLFormElement,omit"
type HTMLFormElement struct {
	window.HTMLElement
}

// CheckValidity Returns whether a form will validate when it is submitted, without having to submit it.
func (*HTMLFormElement) CheckValidity() (b bool) {
	js.Rewrite("$<.CheckValidity()")
	return b
}

// Item Retrieves a form object or an object from an elements collection.
//     * @param name Variant of type Number or String that specifies the object or collection to retrieve. If this parameter is a Number, it is the zero-based index of the object. If this parameter is a string, all objects with matching name or id properties are retrieved, and a collection is returned if more than one match is made.
//     * @param index Variant of type Number that specifies the zero-based index of the object to retrieve when a collection is returned.
func (*HTMLFormElement) Item(name *interface{}, index *interface{}) (i interface{}) {
	js.Rewrite("$<.Item($1, $2)", name, index)
	return i
}

// NamedItem Retrieves a form object or an object from an elements collection.
func (*HTMLFormElement) NamedItem(name string) (i interface{}) {
	js.Rewrite("$<.NamedItem($1)", name)
	return i
}

// Reset Fires when the user resets a form.
func (*HTMLFormElement) Reset() {
	js.Rewrite("$<.Reset()")
}

// Submit Fires when a FORM is about to be submitted.
func (*HTMLFormElement) Submit() {
	js.Rewrite("$<.Submit()")
}

// AcceptCharset Sets or retrieves a list of character encodings for input data that must be accepted by the server processing the form.
func (*HTMLFormElement) AcceptCharset() (acceptCharset string) {
	js.Rewrite("$<.AcceptCharset")
	return acceptCharset
}

// AcceptCharset Sets or retrieves a list of character encodings for input data that must be accepted by the server processing the form.
func (*HTMLFormElement) SetAcceptCharset(acceptCharset string) {
	js.Rewrite("$<.AcceptCharset = $1", acceptCharset)
}

// Action Sets or retrieves the URL to which the form content is sent for processing.
func (*HTMLFormElement) Action() (action string) {
	js.Rewrite("$<.Action")
	return action
}

// Action Sets or retrieves the URL to which the form content is sent for processing.
func (*HTMLFormElement) SetAction(action string) {
	js.Rewrite("$<.Action = $1", action)
}

// Autocomplete Specifies whether autocomplete is applied to an editable text field.
func (*HTMLFormElement) Autocomplete() (autocomplete string) {
	js.Rewrite("$<.Autocomplete")
	return autocomplete
}

// Autocomplete Specifies whether autocomplete is applied to an editable text field.
func (*HTMLFormElement) SetAutocomplete(autocomplete string) {
	js.Rewrite("$<.Autocomplete = $1", autocomplete)
}

// Elements Retrieves a collection, in source order, of all controls in a given form.
func (*HTMLFormElement) Elements() (elements *htmlformcontrolscollection.HTMLFormControlsCollection) {
	js.Rewrite("$<.Elements")
	return elements
}

// Encoding Sets or retrieves the MIME encoding for the form.
func (*HTMLFormElement) Encoding() (encoding string) {
	js.Rewrite("$<.Encoding")
	return encoding
}

// Encoding Sets or retrieves the MIME encoding for the form.
func (*HTMLFormElement) SetEncoding(encoding string) {
	js.Rewrite("$<.Encoding = $1", encoding)
}

// Enctype Sets or retrieves the encoding type for the form.
func (*HTMLFormElement) Enctype() (enctype string) {
	js.Rewrite("$<.Enctype")
	return enctype
}

// Enctype Sets or retrieves the encoding type for the form.
func (*HTMLFormElement) SetEnctype(enctype string) {
	js.Rewrite("$<.Enctype = $1", enctype)
}

// Length Sets or retrieves the number of objects in a collection.
func (*HTMLFormElement) Length() (length int) {
	js.Rewrite("$<.Length")
	return length
}

// Method Sets or retrieves how to send the form data to the server.
func (*HTMLFormElement) Method() (method string) {
	js.Rewrite("$<.Method")
	return method
}

// Method Sets or retrieves how to send the form data to the server.
func (*HTMLFormElement) SetMethod(method string) {
	js.Rewrite("$<.Method = $1", method)
}

// Name Sets or retrieves the name of the object.
func (*HTMLFormElement) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}

// Name Sets or retrieves the name of the object.
func (*HTMLFormElement) SetName(name string) {
	js.Rewrite("$<.Name = $1", name)
}

// NoValidate Designates a form that is not validated when submitted.
func (*HTMLFormElement) NoValidate() (noValidate bool) {
	js.Rewrite("$<.NoValidate")
	return noValidate
}

// NoValidate Designates a form that is not validated when submitted.
func (*HTMLFormElement) SetNoValidate(noValidate bool) {
	js.Rewrite("$<.NoValidate = $1", noValidate)
}

// Target Sets or retrieves the window or frame at which to target content.
func (*HTMLFormElement) Target() (target string) {
	js.Rewrite("$<.Target")
	return target
}

// Target Sets or retrieves the window or frame at which to target content.
func (*HTMLFormElement) SetTarget(target string) {
	js.Rewrite("$<.Target = $1", target)
}
