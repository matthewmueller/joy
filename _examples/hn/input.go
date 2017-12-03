package main

import (
	"github.com/matthewmueller/joy/dom/document"
	"github.com/matthewmueller/joy/jsx"
	"github.com/matthewmueller/joy/_examples/hn/app"
	"github.com/matthewmueller/joy/_examples/hn/preact"
)

func main() {
	jsx.Use("preact.h", "./preact/preact.js")
	preact.Render(app.New(), document.Body)
}
