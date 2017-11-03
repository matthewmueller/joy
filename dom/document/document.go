package document

import "github.com/matthewmueller/golly/js"

// Element struct
// js:"element,omit"
type Element struct {
	NodeType int `js:"nodeType"`
}

// Body of the document
// js:"body,omit"
var Body = Raw(js.Raw("body"))

// Raw element
func Raw(element interface{}) *Element {
	js.Rewrite("$1", element)
	return nil
}

// InnerHTML fn
// func (e *Element) InnerHTML() string {
// 	js.Rewrite("$1.innerHTML", e)
// 	return ""
// }
