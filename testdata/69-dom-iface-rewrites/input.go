package main

import "github.com/matthewmueller/golly/dom/window"

func main() {
	w := window.New()
	doc := w.Document()
	html := doc.DocumentElement()
	println(html.TagName())
}
