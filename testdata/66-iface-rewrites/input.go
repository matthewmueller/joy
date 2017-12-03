package main

import (
	"github.com/matthewmueller/joy/macro"
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
	macro.Rewrite("$_.nodeName")
	return ""
}

func (d *document) TextContext(text string) {
	macro.Rewrite("$_.textContent = $1")
}

type window struct {
}

// js:"nodeName"
func (d *window) NodeName() string {
	macro.Rewrite("$_.nodeName")
	return ""
}

func (d *window) TextContext(text string) {
	macro.Rewrite("$_.textContent = $1")
}

// New fn
func New() (doc Document) {
	macro.Rewrite("document")
	return &window{}
}

func main() {
	doc := New()
	println(doc.NodeName())
	doc.TextContext("hi")
}
