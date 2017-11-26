package main

import "github.com/matthewmueller/golly/js"

// Document interface
type Document interface {
	NodeName() string
}

type document struct {
}

func (d *document) NodeName() string {
	js.Rewrite("$_.nodeName")
	return ""
}

type window struct {
}

func (d *window) NodeName() string {
	js.Rewrite("$_.nodeName")
	return ""
}

// New fn
func New() (doc Document) {
	js.Rewrite("document")
	return doc
}

func main() {
	doc := New()
	println(doc.NodeName())
}
