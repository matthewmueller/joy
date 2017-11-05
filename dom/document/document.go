package document

import (
	"strings"

	"github.com/matthewmueller/golly/js"
)

// Node struct
// js:"element,omit"
type Node struct {
	nodeType  int
	nodeName  string
	NodeValue string
	Children  []Node
}

// Body is a reference to the document body
var Body = js.Rewrite("document.body").(*Node)

// NodeName gets the node name
func (e *Node) NodeName() string {
	js.Rewrite("$<.nodeName", e)
	return e.nodeName
}

// InnerHTML fn
// TODO: go implementation
func (e *Node) InnerHTML() string {
	js.Rewrite("$<.innerHTML")
	// lower := strings.ToLower(e.nodeName)
	return ""
}

// OuterHTML fn
// TODO: go implementation
func (e *Node) OuterHTML() string {
	js.Rewrite("$<.outerHTML")
	lower := strings.ToLower(e.nodeName)
	return "<" + lower + ">" + "</" + lower + ">"
}

// CreateElement creates an element
func CreateElement(element string) *Node {
	js.Rewrite("document.createElement($1)", element)
	return &Node{
		nodeType: 1,
		nodeName: strings.ToUpper(element),
	}
}

// Event struct
// js:"event,omits"
type Event struct {
	Type string `js:"type"`
}

// AddEventListener fn
// js:"addEventListener"
func (e *Node) AddEventListener(event string, fn func(e Event)) {

}
