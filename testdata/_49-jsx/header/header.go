package header

import (
	"github.com/matthewmueller/golly/vdom"
	"github.com/matthewmueller/golly/vdom/h/h3"
)

// Header struct
type Header struct {
	vdom.Component

	props props
	state state
}

// Props for the header
type props struct {
	title    string
	bats     string
	children []vdom.Child
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
func New(title string, children ...vdom.Child) vdom.Component {
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
func (d *Header) Render() vdom.Node {
	return h3.New(h3.Class(d.props.title).Attr("count", d.state.count), d.props.children...)
}
