package index

import (
	"github.com/matthewmueller/golly/testdata/40-vnodes/vnode"
	"github.com/matthewmueller/golly/testdata/40-vnodes/vnode/component"
	"github.com/matthewmueller/golly/testdata/40-vnodes/vnode/element"
)

// Props struct
type Props struct {
	Location string
}

// Page struct
type Page struct {
	props *Props
}

// New fn
func New(p *Props) vnode.VNode {
	return component.New(&Page{p})
}

// ComponentWillMount fn
func (p *Page) ComponentWillMount() {
	p.props.Location = "earth"
}

// Render the index page
func (p *Page) Render() vnode.VNode {
	return element.New("div", nil)
}
