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
	body := document.Body
	body.InnerHTML = "matt"
	// document.Body.InnerHTML = "matt"
	// google := test()
	// println(google)
	// fmt.Printf(document)
}

// func test() string {
// 	return "test"
// }
