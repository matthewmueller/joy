package element

import (
	"strings"

	"github.com/matthewmueller/golly/testdata/40-vnodes/vnode"
)

// enforce VNode type
var _ vnode.VNode = (*Element)(nil)
var _ VElement = (*Element)(nil)

// New element node
func New(tagName string, attributes map[string]string, children ...vnode.VNode) *Element {
	return &Element{
		name:       tagName,
		kind:       1,
		attributes: attributes,
		children:   children,
	}
}

// VElement interface
type VElement interface {
	vnode.VNode
	Attributes() map[string]string
	Children() []vnode.VNode
}

// Element node
type Element struct {
	name       string
	kind       int
	attributes map[string]string
	children   []vnode.VNode
}

// Name fn
func (e *Element) Name() string {
	return ""
}

// Type fn
func (e *Element) Type() int {
	return e.kind
}

// Attributes fn
func (e *Element) Attributes() map[string]string {
	return e.attributes
}

// Children fn
func (e *Element) Children() []vnode.VNode {
	return e.children
}

// String fn
func (e *Element) String() string {
	var children []string
	for _, child := range e.children {
		children = append(children, child.String())
	}
	return "<" + e.name + ">" + strings.Join(children, "") + "</" + e.name + ">"
}
