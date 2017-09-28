package main

import (
	"github.com/matthewmueller/golly/dom"
)

func main() {
	document := dom.Document{}
	body := document.Body
	body.InnerHTML = "matt"

	div := document.CreateElement("div")
	div.InnerHTML = "some title"

	div.AddEventListener("click", onClick)
}

func onClick(e dom.Event) {
	println(e.Type)
}
