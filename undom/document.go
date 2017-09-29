package undom

// Node interface
type Node interface {
	Parent() (parent Node)
	SetParent(parent Node)
	AppendChild(child Node)
	InsertBefore(child, ref Node)
	ReplaceChild(child, ref Node)
	RemoveChild(child Node)
	NextSibling() (sibling Node)
	PreviousSibling() (sibling Node)
	FirstChild() (child Node)
	LastChild() (child Node)
	Children() (children []Node)
	Remove()
}

type node struct {
	parentNode Node
	childNodes []Node
}

// AppendChild fn
func (n *node) AppendChild(child Node) {
	n.InsertBefore(child, nil)
}

// InsertBefore fn
func (n *node) InsertBefore(child Node, ref Node) {
	child.Remove()
	child.SetParent(n)
	if ref == nil {
		n.childNodes = append(n.childNodes, child)
		return
	}

	// splice
	for i, ch := range n.childNodes {
		if ch == ref {
			first := append(n.childNodes[0:i], ref)
			n.childNodes = append(first, n.childNodes[i+1:]...)
		}
	}
}

// ReplaceChild fn
func (n *node) ReplaceChild(child Node, ref Node) {
	if ref.Parent() == n {
		n.InsertBefore(child, ref)
		ref.Remove()
	}
}

// RemoveChild fn
func (n *node) RemoveChild(child Node) {
	// splice
	for i, ch := range n.childNodes {
		if ch == n {
			n.childNodes = append(n.childNodes[0:i], n.childNodes[i+1:]...)
		}
	}
}

// Remove fn
func (n *node) Remove() {
	if n.parentNode == nil {
		return
	}
	n.parentNode.RemoveChild(n)
}

// SetParent fn
func (n *node) SetParent(parent Node) {
	n.parentNode = parent
}

// Parent fn
func (n *node) Parent() (parent Node) {
	return n.parentNode
}

// Children fn
func (n *node) Children() (children []Node) {
	return n.childNodes
}

// NextSibling fn
func (n *node) NextSibling() (sibling Node) {
	parent := n.parentNode
	if parent == nil {
		return nil
	}

	children := parent.Children()
	count := len(children)
	for i := 0; i < count; i++ {
		if children[i] == n && count > i+1 {
			return children[i+1]
		}
	}

	return nil
}

// PreviousSibling fn
func (n *node) PreviousSibling() (sibling Node) {
	parent := n.parentNode
	if parent == nil {
		return nil
	}

	children := parent.Children()
	count := len(children)
	for i := count - 1; i >= 0; i-- {
		if children[i] == n && i-1 >= 0 {
			return children[i-1]
		}
	}

	return nil
}

// FirstChild fn
func (n *node) FirstChild() (child Node) {
	if len(n.childNodes) == 0 {
		return nil
	}
	return n.childNodes[0]
}

// LastChild fn
func (n *node) LastChild() (child Node) {
	l := len(n.childNodes)
	if l == 0 {
		return nil
	}
	return n.childNodes[l-1]
}

// // IEvent interface
// type IEvent interface {
// 	StopPropagation()
// 	StopImmediatePropagation()
// 	PreventDefault()
// }

// // IElement interface
// type IElement interface {
// 	Node
// 	SetAttribute(key, value string)
// 	GetAttribute(key string)
// 	RemoveAttribute(key string)
// 	SetAttributeNS(ns, name, value string)
// 	GetAttributeNS(ns, name string)
// 	RemoveAttributeNS(ns, name string)
// 	AddEventListener(kind string, handler func(e IEvent))
// 	RemoveEventListener(kind string, handler func(e IEvent))
// 	DispatchEvent(e IEvent)
// }

// type Document struct {
// }

// func createDocument() {

// }
