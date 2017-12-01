package main

import (
	"github.com/matthewmueller/golly/dom/htmlbodyelement"
	"github.com/matthewmueller/golly/dom/window"
)

func main() {
	document := window.New().Document()

	a := document.CreateElement("a")
	println(a.NodeName())
	strong := document.CreateElement("strong")
	println(document.CreateElement("strong").OuterHTML())
	a.AppendChild(strong)

	strong.SetTextContent("hi world!")

	body := document.Body().(*htmlbodyelement.HTMLBodyElement)
	body.AppendChild(a)
	println(document.Body().OuterHTML())
}
