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

	h := html.New(html.Manifest("hi").Class("cool"),
		head.New(nil),
		body.New(nil),
	)

	document := window.New().Document()
	vdom.Render(h, document.Body().(*htmlbodyelement.HTMLBodyElement), nil)
}
