package header

import (
	"github.com/matthewmueller/golly/jsx"
	"github.com/matthewmueller/golly/testdata/49-jsx/preact"
)

// Header struct
type Header struct {
	preact.Component

	props props
	state state
}

// Props for the header
type props struct {
	title    string
	bats     string
	children []jsx.Node
}

type state struct {
	count int
}

// - type-safe props + state
// - no generation step so it has IDE proper support
// - go-compatible for SSR
// - not too much boilerplate to create components
// - works with other types of like html / text nodes
// - works in all enumerations of code

// New header
func New(title string, children ...jsx.Node) jsx.Node {
	return &Header{
		props: props{
			title:    title,
			children: children,
		},
		state: state{
			count: 10,
		},
	}
}

// ComponentWillMount fn
// js:"componentWillMount,keep"
func (d *Header) ComponentWillMount() {
	d.state.count = 5
}

func (d *Header) onClick(e interface{}) {
	d.SetState(&state{
		count: d.state.count + 1,
	})
}

// Render header
// js:"render"
func (d *Header) Render() jsx.JSX {
	// children := append(d.props.children, &jsx.Text{Value: strconv.Itoa(d.state.count)})
	return jsx.H("h3", map[string]interface{}{"class": d.props.title, "onClick": d.onClick, "count": d.state.count}, d.props.children...)
}
