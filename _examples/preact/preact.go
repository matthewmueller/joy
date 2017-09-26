package main

import (
	"github.com/matthewmueller/golly/modules/preact"
)

func main() {
	component := preact.New(NewPage(Props{}))
	preact.Render()
	// preact.Render(preact.H(component, ))
	// log.Infof("got component", component)
}

// Props struct
type Props struct {
}

// NewPage creates a new <Page />
func NewPage(props Props) preact.Component {
	return &Page{
		props: props,
	}
}

// Page component
type Page struct {
	props Props
}

// Render the page
func (p *Page) Render() *preact.VNode {
	return preact.Div("hi world!")
}
