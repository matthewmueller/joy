package main

import (
	"github.com/matthewmueller/joy/dom/htmlbodyelement"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/testdata/49-jsx/header"
	"github.com/matthewmueller/joy/testdata/49-jsx/preact"
	"github.com/matthewmueller/joy/vdom"
	"github.com/matthewmueller/joy/vdom/h2"
)

func main() {
	vdom.Use("preact.h", "./preact/preact.js")

	hdr := h2.New(h2.Class("hi"),
		vdom.S("yo!"),
		header.New("lol", vdom.S("hi!")),
	)

	w := window.New()
	document := w.Document()
	preact.Render(hdr, document.Body().(*htmlbodyelement.HTMLBodyElement))
	println(document.Body().InnerHTML())
}
