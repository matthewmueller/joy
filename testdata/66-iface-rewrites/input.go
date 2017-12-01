package main

import (
	"github.com/matthewmueller/golly/js"
)

// Document interface
type Document interface {
	// js:"nodeName"
	// jsrewrite:"$_.nodeName"
	NodeName() string
	// jsrewrite:"$_.textContent = $1"
	TextContext(text string)
}

type document struct {
}

// js:"nodeName"
func (d *document) NodeName() string {
	js.Rewrite("$_.nodeName")
	return ""
}

func (d *document) TextContext(text string) {
	js.Rewrite("$_.textContent = $1")
}

type window struct {
}

// js:"nodeName"
func (d *window) NodeName() string {
	js.Rewrite("$_.nodeName")
	return ""
}

func (d *window) TextContext(text string) {
	js.Rewrite("$_.textContent = $1")
}

// New fn
func New() (doc Document) {
	js.Rewrite("document")
	return &window{}
}

func main() {
	doc := New()
	println(doc.NodeName())
	doc.TextContext("hi")
}
