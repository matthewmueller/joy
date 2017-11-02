package jsx

import "strings"

// Component interface
type Component interface {
	Render() Node
}

// type vnode struct {
// 	NodeName   string
// 	Attributes map[string]interface{}
// }

// Node interface
type Node interface {
	String() string
}

// Element struct
type Element struct {
	NodeName   string
	Attributes map[string]string
	Children   []Component
}

func (e *Element) String() string {
	var attrs []string
	for key, val := range e.Attributes {
		attrs = append(attrs, key+"=\""+val+"\"")
		// // TODO: switch statements
		// switch t := val.(type) {
		// case string:
		// }
	}

	var children []string
	for _, child := range e.Children {
		children = append(children, child.Render().String())
	}

	if len(attrs) > 0 {
		return "<" + e.NodeName + " " + strings.Join(attrs, " ") + ">" + strings.Join(children, "") + "</" + e.NodeName + ">"
	}

	return "<" + e.NodeName + ">" + strings.Join(children, "") + "</" + e.NodeName + ">"
}
