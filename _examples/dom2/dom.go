package main

import (
	"github.com/matthewmueller/golly/dom"
	"github.com/matthewmueller/golly/dom/window"
)

func main() {

	w := window.New()

	w.AddEventListener("click", func(evt *dom.Event) {
		evt.PreventDefault()
		evt.StopImmediatePropagation()

	}, false)
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
