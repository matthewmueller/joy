package topitems

import (
	"github.com/matthewmueller/golly/jsx"
	"github.com/matthewmueller/golly/testdata/56-hn-preact/item"
	"github.com/matthewmueller/golly/testdata/56-hn-preact/preact"
)

type component struct {
	preact.Component

	props props
}

// Props struct
type props struct {
	items []item.Item
}

// New topitems component
func New(items []item.Item) jsx.Node {
	return &component{
		props: props{
			items: items,
		},
	}
}

func (c *component) Render() jsx.JSX {
	return jsx.H("div", nil)
}
