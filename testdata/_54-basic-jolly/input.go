package main

import (
	"github.com/matthewmueller/golly/dom/document"
	"github.com/matthewmueller/golly/testdata/54-basic-jolly/header"
	"github.com/matthewmueller/golly/testdata/54-basic-jolly/preact"
	"github.com/matthewmueller/golly/vdom"
	"github.com/matthewmueller/golly/vdom/h/h2"
)

func main() {
	vdom.Use("preact.h", "./preact/preact.js")

	hdr := h2.New(h2.Class("hi"),
		vdom.S("yo!"),
		header.New("lol", vdom.S("hi!")),
	)

	preact.Render(hdr, document.Body)
	println(document.Body.InnerHTML())
}
