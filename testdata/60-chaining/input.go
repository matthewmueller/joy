package main

import (
	"github.com/matthewmueller/golly/dom/document"
	"github.com/matthewmueller/golly/testdata/60-chaining/header"
	"github.com/matthewmueller/golly/testdata/60-chaining/preact"
	"github.com/matthewmueller/golly/vdom"
	"github.com/matthewmueller/golly/vdom/h/strong"
)

// - type-safe props + state
// - no generation step so it has IDE proper support
// - go-compatible for SSR
// - not too much boilerplate to create components
// - works with other types of like html / text nodes
// - works in all enumerations of code
// - works with other virtual dom libraries (particularly preact)
// - can have multiple elements in a single file

func main() {
	vdom.Use("preact.h", "./preact/preact.js")

	v := header.New("some title", "some body",
		strong.New(strong.Class("some class").ID("some id"),
			vdom.S("some text"),
			header.New("subtitle", "subbody", vdom.S("subbody")),
		),
	)

	preact.Render(v, document.Body)
	println(document.Body.InnerHTML())
}
