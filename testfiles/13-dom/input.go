package main

import "github.com/matthewmueller/golly/js"

func main() {
	document := New()
	div := document.CreateElement("div")
	strong := document.CreateElement("strong")
	strong.TextContent("nice!")
	div.AppendChild(strong)
	println(getText(div))
}

func getText(div *Node) string {
	return div.Children[0].Children[0].NodeValue
}

// New Document
func New() Document {
	return Document{}
}

// Document struct
type Document struct {
	Body *Node `js:"body"`
}

// golly map
var golly = map[string]func(js.INode) (js.INode, error){
	"Document": func(existing js.INode) (js.INode, error) {
		return nil, nil
	},
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

// func NewWindow() *Window {
// 	return &Window{
// 		Document
// 	}
// }

// // Window fn
// func Window() *Document {
// 	doc, _ := html.Parse(bytes.NewBuffer([]byte("<html><body>hi world!</body></html>")))

// 	// TODO: more robust
// 	body := doc.FirstChild

// 	return &Document{
// 		root: doc,
// 		body: &Node{
// 			nodeName:   body.Data,
// 			nodeType:   int(body.Type),
// 			attributes: body.Attr,
// 		},
// 	}
// }

// // Document struct
// type Document struct {
// 	root *html.Node
// 	body *Node
// }

// // Node struct
// type Node struct {
// 	nodeName   string
// 	nodeType   int
// 	attributes []html.Attribute
// 	children   []Node
// }

// // GetElementByID lookup element by it's id
// func (*Document) GetElementByID(id string) string {
// 	return "#" + id
// }

// // InnerHTML str
// func (*Node) InnerHTML() string {
// 	return "hello!"
// }
