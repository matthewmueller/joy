package item

import "github.com/matthewmueller/golly/jsx"

// Item from hacker news
type Item struct {
	ID          string `json:"id,omitempty" js:"id"`
	Score       int    `json:"score,omitempty" js:"score"`
	Descendants int    `json:"descendants,omitempty" js:"descendants"`
	Title       string `json:"title,omitempty" js:"title"`
	URL         string `json:"url,omitempty" js:"url"`
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
