package main

import (
	"github.com/matthewmueller/golly/window"
)

func main() {
	w := window.New()
	doc := w.Document()
	println(doc.NodeName())

	html := doc.DocumentElement().(window.HTMLElement)

	html.AddEventListener("click", func(evt window.Event) {
		evt.PreventDefault()
		evt.StopImmediatePropagation()
		println(evt.Type())
	}, false)

	q := html.QuerySelector("a")
	if q != nil {
		println(q)
	}
	a := q.(*window.HTMLAnchorElement)

	println(a.Href())
}
