package main

import (
	"github.com/matthewmueller/golly/dom"
)

func main() {
	w := &dom.Window{}
	w.AddEventListener("click", func(evt *dom.Event) {
		t := evt.GetType()
		println(t)
	}, false)
	w.Alert("amazing!!!!!!")
}
