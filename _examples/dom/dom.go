package main

import (
	"github.com/apex/log"
	"github.com/matthewmueller/golly/dom"
	"github.com/matthewmueller/golly/dom/internal/tester"
)

func main() {
	w := &dom.Window{}
	w.AddEventListener("click", func(evt *dom.Event) {
		t := evt.GetType()
		println(t)
	}, false)
	w.Alert("amazing!!!!!!")
}

func test() {
	result := dom.Test()
	r := tester.Tester
	log.Infof("result %s", result)
	log.Infof("result %s", r)
}
