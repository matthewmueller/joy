package main

import (
	"github.com/matthewmueller/golly/dom3/window"
)

func main() {
	w := window.New()
	doc := w.Document
	println(doc.NodeName)

	html, ok := doc.DocumentElement.(window.HTMLElement)
	if !ok {
		return
	}

	a := html.QuerySelector("a")
	if a != nil {
		println(a)
	}

	html.AddEventListener("click", func(evt window.Event) {
		evt.PreventDefault()
		evt.StopImmediatePropagation()
	}, false)
}
