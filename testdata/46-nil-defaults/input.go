package main

import (
	"github.com/matthewmueller/joy/testdata/46-nil-defaults/element"
)

// var _ element.Element = (*Page)(nil)

// func New() element.Element {

// }

// type page struct {
// 	value string
// }

// func (p *page) Render() string {
// 	return element.New("header", element.New("strong", nil)).Render()
// }

func main() {
	el := element.New("header", element.New("strong", nil))
	println(el.Render())
	println(el.Attrs())
}
