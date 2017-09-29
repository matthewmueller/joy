// Package dom is the public-facing DOM library
//
// TODO: autogenerate with typescript scripts
package dom

// Document struct
//
// js:"document,global"
type Document struct {
	Body       *Node `json:"body" js:"body"`
	FirstChild *Node `json:"first_child"`
}

// CreateElement struct
// js:"createElement"
func (d *Document) CreateElement(nodeName string) Node {
	return Node{
		NodeType: 1,
		NodeName: nodeName,
		Children: []Node{},
	}
}

// Node struct
// js:"node,global"
type Node struct {
	NodeType  int
	NodeName  string
	NodeValue string
	Children  []Node

	// functions
	InnerHTML string `js:"innerHTML"`
}

// Event struct
// js:"event,global"
type Event struct {
	Type string `js:"type"`
}

// AddEventListener fn
// js:"addEventListener"
func (n *Node) AddEventListener(event string, fn func(e Event)) {

}
