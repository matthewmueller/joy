package main


import "github.com/matthewmueller/golly/testdata/60-chaining/vdom"
import "github.com/matthewmueller/golly/testdata/60-chaining/header"
import "github.com/matthewmueller/golly/testdata/60-chaining/strong"
import "github.com/matthewmueller/golly/testdata/60-chaining/preact"
import "github.com/matthewmueller/golly/dom/document"
import "github.com/matthewmueller/golly/jsx"

// - type-safe props + state
// - no generation step so it has IDE proper support
// - go-compatible for SSR
// - not too much boilerplate to create components
// - works with other types of like html / text nodes
// - works in all enumerations of code
// - works with other virtual dom libraries (particularly preact)
// - can have multiple elements in a single file

func main() {
	jsx.Use("preact.h", "./preact/preact.js")
	// props := strong.Class("hi").ID("cool")

	v := header.New("some title", "some body",
		strong.New(strong.Class("some class").ID("some id"), 
			vdom.T("some text"),
		),
	)

	println(document.Body)
	preact.Render(v, document.Body)
	println(document.Body.InnerHTML())	
	// println(v.Render().String())
	// h := header.New("welcome", "world", strong.New(strong.Class("lol...").Key("child")))
	// s := strong.New(props.Key("lol"), h)
	// println(s)
	// println(h)
}

