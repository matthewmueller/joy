package jsx

import (
	"strings"

	"github.com/matthewmueller/golly/js"
)

// Use fn
// js:"use,omit"
func Use(pragma, filepath string) {
}

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

// element struct
// js:"element,omit"
type element struct {
	NodeName   string                 `js:"nodeName"`
	Attributes map[string]interface{} `js:"attributes"`
	Children   []JSX                  `js:"children"`
}

// Element interface
type Element interface {
	Component
	JSX
}

// H creates an HTML element
func H(name string, attrs map[string]interface{}, children ...Component) Element {
	js.Rewrite("preact.h($1, $2, $3)", name, attrs, children)

	var chs []JSX
	for _, child := range children {
		chs = append(chs, child.Render())
	}

	// Hack to make sure preact is present
	// when we do the transforms
	// var preact = js.RawFile("./preact.js")
	// _ = preact

	return &element{
		NodeName:   name,
		Attributes: attrs,
		Children:   chs,
	}
}

// Render element
func (e *element) Render() JSX {
	return e
}

func (e *element) String() string {
	var attrs []string
	for key, val := range e.Attributes {
		// TODO: switch statements
		switch t := val.(type) {
		case string:
			attrs = append(attrs, key+"=\""+t+"\"")
		}
	}

	var children []string
	for _, child := range e.Children {
		children = append(children, child.String())
	}

	if len(attrs) > 0 {
		return "<" + e.NodeName + " " + strings.Join(attrs, " ") + ">" + strings.Join(children, "") + "</" + e.NodeName + ">"
	}

	return "<" + e.NodeName + ">" + strings.Join(children, "") + "</" + e.NodeName + ">"
}
