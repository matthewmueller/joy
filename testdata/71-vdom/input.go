package main

import (
	"github.com/matthewmueller/joy/dom/htmlbodyelement"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/vdom"
	"github.com/matthewmueller/joy/vdom/body"
	"github.com/matthewmueller/joy/vdom/head"
	"github.com/matthewmueller/joy/vdom/html"
)

func main() {
	vdom.Use("preact.h", "./preact.js")

	hdr := html.New(html.Manifest("hi").Class("cool"),
		head.New(nil),
		body.New(nil, vdom.S("hi world")),
	)

	w := window.New()
	document := w.Document()
	vdom.Render(hdr, document.Body().(*htmlbodyelement.HTMLBodyElement), nil)
	println(document.Body().InnerHTML())
}
