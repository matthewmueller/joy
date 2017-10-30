package main

import (
	"github.com/matthewmueller/golly/testdata/47-circular/element"
)

type page struct {
}

func (p *page) Render() string {
	return element.New("header", element.New("strong", nil)).Render()
}

func main() {
	p := page{}
	println(p.Render())
}
