package main

// Document interface
type Document interface {
	// js:"NodeName,rewrite=joyNodeName"
	NodeName() string
}

type document struct {
}

func (d *document) NodeName() string {
	return ""
}

// New fn
func New() (doc Document) {
	return doc
}

func main() {
	doc := New()
	println(doc.NodeName())
}
