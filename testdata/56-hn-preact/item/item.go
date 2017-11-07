package item

import "github.com/matthewmueller/golly/jsx"

// Item from hacker news
type Item struct {
	ID          string
	Score       int
	Descendants int
	Title       string
	URL         string
}

type component struct {
	props props
}

// Props for the item component
type props struct {
	item *Item
}

// New itemview
func New(item *Item) jsx.Node {
	return &component{
		props: props{
			item: item,
		},
	}
}

func (c *component) Render() jsx.JSX {
	return jsx.H("div", nil)
}
