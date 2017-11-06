package main

import (
	"github.com/matthewmueller/golly/js"
	"github.com/matthewmueller/golly/jsx"
	"github.com/matthewmueller/golly/testdata/49-jsx/header"
)

func main() {
	jsx.Use("preact.h", "./preact/preact.js")

	hdr := jsx.H("h2", map[string]interface{}{"class": "hi"},
		&jsx.Text{Value: "yo!"},
		header.New("lol", &jsx.Text{Value: "hi!"}),
	)

	// preact.Render(hdr, js.Raw("document.body"))
	// println(document.Body.InnerHTML())
	js.Raw("preact.render(hdr, document.body)", hdr.Render)
	println(js.Raw("document.body.innerHTML"))
	// println(hdr.String())
}
