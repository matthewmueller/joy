package main

import "github.com/matthewmueller/joy/dom/window"

func main() {
	w := window.New()
	d := w.Document()
	html := d.DocumentElement().(window.HTMLElement)

	html.AddEventListener("click", func(evt window.Event) {
	}, false)
}
