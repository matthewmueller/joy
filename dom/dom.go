// Package dom is the public-facing DOM library
//
// TODO: autogenerate with typescript scripts
package dom

// Document struct
type Document struct {
	Body *Node `js:"body"`
}

// CreateElement struct
func (d *Document) CreateElement(nodeName string) *Node {
	return &Node{
		NodeType: 1,
		NodeName: nodeName,
		Children: []*Node{},
	}
}

// Node struct
type Node struct {
	NodeType  int
	NodeName  string
	NodeValue string
	Children  []*Node
}

// TextContent fn
func (n *Node) TextContent(text string) {
	n.Children = append(n.Children, &Node{
		NodeType:  3,
		NodeName:  "#text",
		NodeValue: text,
	})
}

// AppendChild fn
func (n *Node) AppendChild(child *Node) *Node {
	n.Children = append(n.Children, child)
	return n
}
