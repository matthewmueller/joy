package main

import (
	"github.com/matthewmueller/golly/dom"
)

func main() {
	var document = dom.Document{}
	body := document.Body
	body.InnerHTML = "matt"

	div := document.CreateElement("div")
	div.InnerHTML = "some title"

	div.AddEventListener("click", func(e dom.Event) {
		println(e.Type)
	})
}
