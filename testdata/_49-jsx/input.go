package main

import (
	"github.com/matthewmueller/golly/testdata/49-jsx/header"
	"github.com/matthewmueller/golly/testdata/49-jsx/preact"
	"github.com/matthewmueller/golly/testdata/49-jsx/window"
	"github.com/matthewmueller/golly/vdom"
	"github.com/matthewmueller/golly/vdom/h/h2"
)

func main() {
	vdom.Use("preact.h", "./preact/preact.js")

	hdr := h2.New(h2.Class("hi"),
		vdom.S("yo!"),
		header.New("lol", vdom.S("hi!")),
	)

	w := window.New()
	document := w.Document()
	preact.Render(hdr, document.Body())
	println(document.Body().InnerHTML())
}
