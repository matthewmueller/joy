package dom

import (
	"strings"

	"github.com/matthewmueller/golly/js"
)

// Body is a reference to the document body
var Body = func() *Element {
	js.Rewrite("document.body")
	return &Element{
		nodeType: 1,
		nodeName: "BODY",
	}
}()

// Hello string
var Hello = "hi!"

// Element struct
// js:"element,omit"
type Element struct {
	nodeType int
	nodeName string
}

// NodeName gets the node name
func (e *Element) NodeName() string {
	js.Rewrite("$<.nodeName", e)
	return e.nodeName
}

// OuterHTML fn
func (e *Element) OuterHTML() string {
	js.Rewrite("$<.outerHTML")
	lower := strings.ToLower(e.nodeName)
	return "<" + lower + ">" + "</" + lower + ">"
}

// CreateElement creates an element
func CreateElement(element string) *Element {
	js.Rewrite("document.createElement($1)", element)
	return &Element{
		nodeType: 1,
		nodeName: strings.ToUpper(element),
	}
}
