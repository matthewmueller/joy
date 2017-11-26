package main

import "github.com/matthewmueller/golly/internal/compiler/loader/testdata/dom/window"

func Element() window.Element {
	return nil
}

func main() {
	html := Element().(window.HTMLElement)
	html.DispatchEvent()
}
