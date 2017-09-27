package main

import (
	"github.com/matthewmueller/golly/dom"
)

var document = dom.Document{
	Body: &dom.Node{
		NodeName: "body",
	},
}

func main() {
	// document.
	document.Body.InnerHTML = "matt"
	// fmt.Printf(document)
}
