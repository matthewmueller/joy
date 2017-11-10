package main

import (
	"github.com/matthewmueller/golly/dom/document"
	"github.com/matthewmueller/golly/jsx"
	"github.com/matthewmueller/golly/testdata/56-hn-preact/app"
	"github.com/matthewmueller/golly/testdata/56-hn-preact/preact"
)

func main() {
	jsx.Use("preact.h", "./preact/preact.js")
	preact.Render(app.New(), document.Body)
	// time.Sleep(10 * time.Second)
}
