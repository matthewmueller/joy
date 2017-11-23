package main

import (
	"github.com/matthewmueller/golly/dom2/window"
)

func main() {

	w := window.Window{}
	w.AddEventListener("click", func(evt window.Event) {
		// evt.
	})
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
