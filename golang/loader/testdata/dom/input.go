package main

import "github.com/matthewmueller/golly/golang/loader/testdata/dom/window"

func Element() window.Element {
	return nil
}

func main() {
	html := Element().(window.HTMLElement)
	html.DispatchEvent()
}
