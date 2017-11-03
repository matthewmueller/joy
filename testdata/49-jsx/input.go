package main

import (
	"github.com/matthewmueller/golly/js"
	"github.com/matthewmueller/golly/jsx"
	"github.com/matthewmueller/golly/testdata/49-jsx/header"
)

func main() {
	hdr := jsx.H("h2", map[string]interface{}{"class": "hi"},
		&jsx.Text{Value: "yo!"},
		header.New("lol", &jsx.Text{Value: "hi!"}),
	)

	js.Raw("preact.render(hdr, document.body)", hdr.Render)
	println(js.Raw("document.body.innerHTML"))
	// println(hdr.String())
}
