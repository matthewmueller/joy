package main

import (
	"github.com/matthewmueller/golly/dom2/window"
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
	w := window.Window{}

	// s := Sub{}
	// s.Go()

	w.AddEventListener("click", func(evt window.Event) {
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
