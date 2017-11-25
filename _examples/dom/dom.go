package main

import (
	"github.com/matthewmueller/golly/dom/window"
)

// type Test interface {
// 	Go()
// }

// type Sub struct {
// 	Test
// }

// func (t *Sub) Go() {

// }

func main() {
	w := window.New()
	doc := w.Document()
	println(doc.NodeName())
	// html, ok := doc.DocumentElement().(window.HTMLElement)
	// if !ok {
	// 	return
	// }
	// html.EscapeString(s)
	// w.Alert("hi")
	// println(w.DefaultStatus())
	// // s := Sub{}
	// // s.Go()

	// a := html.QuerySelector("a")
	// if a != nil {
	// 	println(a)
	// }
	// doc.QueryCommandEnabled(commandId)

	// html.AddEventListener("click", func(evt window.Event) {
	// 	evt.PreventDefault()
	// 	evt.StopImmediatePropagation()
	// }, false)

	// w.AddEventListener("click", func(evt *dom.Event) {
	// 	t := evt.GetType()
	// 	println(t)
	// }, false)
	// w.Alert("amazing!!!!!!")

	// a := w.GetDocument().QuerySelector("a")
	// hostname := a.(dom.HTMLAnchorElement).GetHostname()
	// println(hostname)
}

// func test() {
// 	result := dom.Test()
// 	r := tester.Tester
// 	log.Infof("result %s", result)
// 	log.Infof("result %s", r)
// }
