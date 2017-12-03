package main

import "github.com/matthewmueller/joy/macro"

// Body var
// js:"body"
var Body = macro.Raw("document.body.nodeName")

func main() {
	println(Body)
}
