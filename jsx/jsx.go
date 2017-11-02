package jsx

import (
	"strings"

	"github.com/matthewmueller/golly/js"
)

// Component interface
type Component interface {
	Render() JSX
}

// JSX interface
type JSX interface {
	String() string
}

// Text struct
// js:"Text,omit"
type Text struct {
	Value string
}

// Render text
func (t *Text) Render() JSX {
	return t
}

func (t *Text) String() string {
	return t.Value
}

// Element struct
// js:"Element,omit"
type Element struct {
	NodeName   string            `js:"nodeName"`
	Attributes map[string]string `js:"attributes"`
	Children   []Component       `js:"children"`
}

// Render element
func (e *Element) Render() JSX {
	// another hack to make sure preact is present
	// when we do the transforms
	var preact = js.RawFile("./preact.js")
	_ = preact

	// TODO: do not remove me...
	var children []string
	for _, child := range e.Children {
		children = append(children, child.Render().String())
	}
	// ... this is a hack that tricks the dead-code elimination
	// to include all the component's render functions.
	//
	// we'll need to find a better way to do this because this
	// function is omitted from the build, so technically we
	// should also be emitting all it's dependencies.
	return e
}

func (e *Element) String() string {
	var attrs []string
	// for key, val := range e.Attributes {
	// 	attrs = append(attrs, key+"=\""+val+"\"")
	// 	// // TODO: switch statements
	// 	// switch t := val.(type) {
	// 	// case string:
	// 	// }
	// }

	var children []string
	for _, child := range e.Children {
		children = append(children, child.Render().String())
	}

	if len(attrs) > 0 {
		return "<" + e.NodeName + " " + strings.Join(attrs, " ") + ">" + strings.Join(children, "") + "</" + e.NodeName + ">"
	}

	return "<" + e.NodeName + ">" + strings.Join(children, "") + "</" + e.NodeName + ">"
}
