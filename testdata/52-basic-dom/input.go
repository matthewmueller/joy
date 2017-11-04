package main

import "github.com/matthewmueller/golly/testdata/52-basic-dom/dom"

func main() {
	a := dom.CreateElement("a")
	println(a.NodeName())
	println(dom.CreateElement("strong").OuterHTML())
	println(dom.Body.OuterHTML())
	println(dom.Body.OuterHTML)
	println(dom.Hello)
}
