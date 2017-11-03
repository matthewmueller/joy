package main

import (
	"github.com/matthewmueller/golly/dom/document"
	"github.com/matthewmueller/golly/jsx"
	"github.com/matthewmueller/golly/testdata/52-basic-jolly/header"
	"github.com/matthewmueller/golly/testdata/52-basic-jolly/preact"
)

func main() {
	hdr := jsx.H("h2", map[string]interface{}{"class": "hi"},
		&jsx.Text{Value: "yo!"},
		header.New("lol", &jsx.Text{Value: "hi!"}),
	)

	preact.Render(hdr, document.Body)
	// println(document.Body.InnerHTML())
	// println(hdr.String())
}
