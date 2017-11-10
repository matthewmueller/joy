package main

import (
	"github.com/matthewmueller/golly/dom/document"
	"github.com/matthewmueller/golly/jsx"
	"github.com/matthewmueller/golly/_examples/hn/app"
	"github.com/matthewmueller/golly/_examples/hn/preact"
)

func main() {
	jsx.Use("preact.h", "./preact/preact.js")
	preact.Render(app.New(), document.Body)
}
