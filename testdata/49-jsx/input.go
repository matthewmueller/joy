package main

import (
	"github.com/matthewmueller/golly/jsx"
	"github.com/matthewmueller/golly/testdata/49-jsx/h2"
	"github.com/matthewmueller/golly/testdata/49-jsx/header"
)

func main() {
	header := &h2.H2{Class: "hi", Children: []jsx.Component{
		&jsx.Text{Value: "yo!"},
		&header.Header{Title: "lol", Children: []jsx.Component{
			&jsx.Text{Value: "hi!"},
		}},
	}}

	// var header = h("h2", { class: "hi" }, [
	// 	"yo!",
	// 	h(Header, { title: "lol" }, [
	// 		"hi!"
	// 	])
	// ])

	println(header.Render().String())
}

// 1. test if composite literal's struct implements jsx.Component
