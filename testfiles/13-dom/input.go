package main

func main() {
	document := Document{}
	div := document.CreateElement("div")
	strong := document.CreateElement("strong")
	strong.TextContent("nice!")
	div.AppendChild(strong)
	println(getText(div))
}

func getText(div *Node) string {
	return div.children[0].children[0].nodeValue
}

// Document struct
type Document struct {
}

// CreateElement struct
func (d *Document) CreateElement(nodeName string) *Node {
	return &Node{
		nodeType: 1,
		nodeName: nodeName,
		children: []*Node{},
	}
}

// Node struct
type Node struct {
	nodeType  int
	nodeName  string
	nodeValue string
	children  []*Node
}

// TextContent fn
func (n *Node) TextContent(text string) {
	n.children = append(n, &Node{
		nodeType:  3,
		nodeName:  "#text",
		nodeValue: text,
	})
}

// AppendChild fn
func (n *Node) AppendChild(child *Node) *Node {
	n.children = append(n.children, child)
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
