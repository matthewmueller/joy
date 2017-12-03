package main

import "github.com/matthewmueller/joy/macro"

// Document struct
// js:"document,omit"
type Document struct {
}

// NodeName fn
// js:"nodeName"
func (d *Document) NodeName() string {
	macro.Rewrite("$_.nodeName")
	return ""
}

// New fn
func New() *Window {
	macro.Rewrite("window")
	return &Window{}
}

// Window struct
// js:"window,omit"
type Window struct {
}

// Document fn
func (w *Window) Document() *Document {
	macro.Rewrite("document")
	return &Document{}
}

func main() {
	w := New()
	doc := w.Document()
	println(doc.NodeName())
}
