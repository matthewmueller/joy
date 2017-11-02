package main

import (
	"github.com/matthewmueller/golly/js"
	"github.com/matthewmueller/golly/jsx"
	"github.com/matthewmueller/golly/testdata/49-jsx/header"
)

func main() {
	hdr := jsx.Element{NodeName: "h2", Attributes: map[string]string{"class": "hi"}, Children: []jsx.Component{
		&jsx.Text{Value: "yo!"},
		&header.Header{Props: &header.Props{Title: "lol", Children: []jsx.Component{
			&jsx.Text{Value: "hi!"},
		}}},
	}}

	// var header = h("h2", { class: "hi" }, [
	// 	"yo!",
	// 	h(Header, { title: "lol" }, [
	// 		"hi!"
	// 	])
	// ])

	js.Raw("preact.render(hdr, document.body)", hdr.Render)
	println(js.Raw("document.body.innerHTML"))
}

// 1. test if composite literal's struct implements jsx.Component
