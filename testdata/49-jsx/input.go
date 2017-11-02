package main

import "github.com/matthewmueller/golly/testdata/49-jsx/h2"
import "github.com/matthewmueller/golly/testdata/49-jsx/jsx"
import "github.com/matthewmueller/golly/testdata/49-jsx/text"

func main() {

	header := &h2.H2{Class: "hi", Children: []jsx.Component{
		&text.Text{Value: "yo!"},
		&h2.H2{Class: "lol", Children: []jsx.Component{
			&text.Text{Value: "hi!"},
		}},
	}}

	println(header.String())
}
