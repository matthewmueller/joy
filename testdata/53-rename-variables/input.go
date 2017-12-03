package main

import "github.com/matthewmueller/joy/js"

// Body var
// js:"body"
var Body = js.Raw("document.body.nodeName")

func main() {
	println(Body)
}
