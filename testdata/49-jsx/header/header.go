package header

import (
	"strconv"

	"github.com/matthewmueller/golly/jsx"
	"github.com/matthewmueller/golly/testdata/49-jsx/preact"
)

// Header struct
type header struct {
	preact.Component

	props props
	state state
}

type props struct {
	title    string
	children []jsx.Node
	bats     string
}

type state struct {
	count int
}

// New Header
func New(title string, children ...jsx.Node) jsx.Node {
	return &header{
		props: props{
			title:    title,
			children: children,
			bats:     "are crazy",
		},
		state: state{
			count: 0,
		},
	}
}

func (d *header) onClick(e interface{}) {
	d.SetState(&state{
		count: d.state.count + 1,
	})
}

// Render header
// js:"render"
func (d *header) Render() jsx.JSX {
	children := append(d.props.children, &jsx.Text{Value: strconv.Itoa(d.state.count)})
	return jsx.H("h3", map[string]interface{}{"class": d.props.title, "onClick": d.onClick}, children...)
}
