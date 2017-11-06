package header

import (
	"strconv"

	"github.com/matthewmueller/golly/jsx"
	"github.com/matthewmueller/golly/testdata/54-basic-jolly/preact"
)

// Header struct
type Header struct {
	preact.Component

	// props
	props props

	// state
	state state
}

type props struct {
	title    string
	children []jsx.Node
	bats     string
}

// State struct
type state struct {
	count int
}

// New Header
func New(title string, children ...jsx.Node) jsx.Node {
	return &Header{
		props: props{
			title:    title,
			children: children,
			bats:     "are crazy",
		},
	}
}

// MouseEvent struct
// js:"mouseevent,omit"
type MouseEvent struct {
	Type string `js:"type"`
}

// OnClick fn
func (d *Header) OnClick(e MouseEvent) {
	println(e.Type)
	d.SetState(&state{
		count: d.state.count + 1,
	})
}

// Render header
// js:"render"
func (d *Header) Render() jsx.JSX {
	children := append(d.props.children, &jsx.Text{Value: strconv.Itoa(d.state.count)})
	return jsx.H("h3", map[string]interface{}{"class": d.props.title, "count": d.state.count, "onClick": d.OnClick}, children...)
}
