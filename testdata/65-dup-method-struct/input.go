package main

import "github.com/matthewmueller/golly/js"

// Document struct
// js:"document,omit"
type Document struct {
}

// NodeName fn
// js:"nodeName"
func (d *Document) NodeName() string {
	js.Rewrite("$<.nodeName")
	return ""
}

// New fn
func New() *Window {
	js.Rewrite("window")
	return &Window{}
}

// Window struct
// js:"window,omit"
type Window struct {
}

// Document fn
func (w *Window) Document() *Document {
	js.Rewrite("document")
	return &Document{}
}

func main() {
	w := New()
	doc := w.Document()
	println(doc.NodeName())
}
