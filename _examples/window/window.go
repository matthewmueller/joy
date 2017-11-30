package main

import (
	dom "github.com/matthewmueller/golly/dom"
)

func main() {
	w := dom.New()
	doc := w.Document()
	println(doc.NodeName())

	html := doc.DocumentElement().(dom.HTMLElement)

	html.AddEventListener("click", func(evt dom.Event) {
		evt.PreventDefault()
		evt.StopImmediatePropagation()
		println(evt.Type())
	}, false)

	q := html.QuerySelector("a")
	if q != nil {
		println(q)
	}
	a := q.(*dom.HTMLAnchorElement)

	println(a.Href())
	// for _, node := range a.ChildNodes() {
	// 	println("value %s", node.NodeValue())
	// }
}
