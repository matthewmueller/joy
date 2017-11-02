package jsx

import "strings"

// Component interface
type Component interface {
	Render() JSX
}

// JSX interface
type JSX interface {
	String() string
}

// Text struct
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
type Element struct {
	NodeName   string
	Attributes map[string]string
	Children   []Component
}

// Render element
func (e *Element) Render() JSX {
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
