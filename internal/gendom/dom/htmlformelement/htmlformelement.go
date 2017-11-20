package htmlformelement

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlelement"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlformcontrolscollection"
	"github.com/matthewmueller/golly/js"
)

type HTMLFormElement struct {
	*htmlelement.HTMLElement
}

func (*HTMLFormElement) CheckValidity() (b bool) {
	js.Rewrite("$<.checkValidity()")
	return b
}

func (*HTMLFormElement) Item(name interface{}, index interface{}) (i interface{}) {
	js.Rewrite("$<.item($1, $2)", name, index)
	return i
}

func (*HTMLFormElement) NamedItem(name string) (i interface{}) {
	js.Rewrite("$<.namedItem($1)", name)
	return i
}

func (*HTMLFormElement) Reset() {
	js.Rewrite("$<.reset()")
}

func (*HTMLFormElement) Submit() {
	js.Rewrite("$<.submit()")
}

func (*HTMLFormElement) GetAcceptCharset() (acceptCharset string) {
	js.Rewrite("$<.acceptCharset")
	return acceptCharset
}

func (*HTMLFormElement) SetAcceptCharset(acceptCharset string) {
	js.Rewrite("$<.acceptCharset = $1", acceptCharset)
}

func (*HTMLFormElement) GetAction() (action string) {
	js.Rewrite("$<.action")
	return action
}

func (*HTMLFormElement) SetAction(action string) {
	js.Rewrite("$<.action = $1", action)
}

func (*HTMLFormElement) GetAutocomplete() (autocomplete string) {
	js.Rewrite("$<.autocomplete")
	return autocomplete
}

func (*HTMLFormElement) SetAutocomplete(autocomplete string) {
	js.Rewrite("$<.autocomplete = $1", autocomplete)
}

func (*HTMLFormElement) GetElements() (elements *htmlformcontrolscollection.HTMLFormControlsCollection) {
	js.Rewrite("$<.elements")
	return elements
}

func (*HTMLFormElement) GetEncoding() (encoding string) {
	js.Rewrite("$<.encoding")
	return encoding
}

func (*HTMLFormElement) SetEncoding(encoding string) {
	js.Rewrite("$<.encoding = $1", encoding)
}

func (*HTMLFormElement) GetEnctype() (enctype string) {
	js.Rewrite("$<.enctype")
	return enctype
}

func (*HTMLFormElement) SetEnctype(enctype string) {
	js.Rewrite("$<.enctype = $1", enctype)
}

func (*HTMLFormElement) GetLength() (length int) {
	js.Rewrite("$<.length")
	return length
}

func (*HTMLFormElement) GetMethod() (method string) {
	js.Rewrite("$<.method")
	return method
}

func (*HTMLFormElement) SetMethod(method string) {
	js.Rewrite("$<.method = $1", method)
}

func (*HTMLFormElement) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*HTMLFormElement) SetName(name string) {
	js.Rewrite("$<.name = $1", name)
}

func (*HTMLFormElement) GetNoValidate() (noValidate bool) {
	js.Rewrite("$<.noValidate")
	return noValidate
}

func (*HTMLFormElement) SetNoValidate(noValidate bool) {
	js.Rewrite("$<.noValidate = $1", noValidate)
}

func (*HTMLFormElement) GetTarget() (target string) {
	js.Rewrite("$<.target")
	return target
}

func (*HTMLFormElement) SetTarget(target string) {
	js.Rewrite("$<.target = $1", target)
}
