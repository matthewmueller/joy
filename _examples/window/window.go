package main

import (
	"github.com/matthewmueller/golly/dom/htmlanchorelement"
	"github.com/matthewmueller/golly/dom/htmlhtmlelement"
	"github.com/matthewmueller/golly/dom/websocket"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/dom/xmlhttprequest"
)

func main() {
	w := window.Window{}
	doc := w.Document()
	println(doc.NodeName())

	html := doc.DocumentElement().(*htmlhtmlelement.HTMLHTMLElement)

	html.AddEventListener("click", func(evt window.Event) {
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

	x := xmlhttprequest.XMLHTTPRequest{}
	x.Open("get", "http://google.com", nil, nil, nil)
	x.Send(nil)

	ws := &websocket.WebSocket{}

	ws.SetOnopen(func(e window.Event) {
		println(e.Bubbles())
	})

	ws.SetOnmessage(func(m *window.MessageEvent) {
		data := m.Data()
		println(data)
	})
	// for _, node := range a.ChildNodes() {
	// 	println("value %s", node.NodeValue())
	// }
}
