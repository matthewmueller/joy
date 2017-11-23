package main

import "github.com/matthewmueller/golly/js"

// Document struct
type Document struct {
}

// NodeName fn
// js:"nodeName"
func (d *Document) NodeName() string {
	return ""
}

// Window struct
type Window struct {
}

// Document fn
func (w *Window) Document() *Document {
	js.Rewrite("document")
	return &Document{}
}

func main() {
	w := Window{}
	doc := w.Document()
	println(doc.NodeName())
}
