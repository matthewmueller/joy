package main

import (
	"github.com/matthewmueller/golly/dom/event"
	"github.com/matthewmueller/golly/dom/htmlanchorelement"
	"github.com/matthewmueller/golly/dom/htmlhtmlelement"
	"github.com/matthewmueller/golly/dom/window"
)

func main() {
	w := window.Window{}
	doc := w.Document()
	println(doc.NodeName())

	html := doc.DocumentElement().(*htmlhtmlelement.HTMLHTMLElement)

	html.AddEventListener("click", func(evt event.Event) {
		evt.PreventDefault()
		evt.StopImmediatePropagation()
		println(evt.Type())
	}, false)

	q := html.QuerySelector("a")
	if q != nil {
		println(q)
	}
	a := q.(*htmlanchorelement.HTMLAnchorElement)

	println(a.Href())
	// for _, node := range a.ChildNodes() {
	// 	println("value %s", node.NodeValue())
	// }
}
