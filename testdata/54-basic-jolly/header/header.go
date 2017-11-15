package header

import (
	"strconv"

	"github.com/matthewmueller/golly/js"
	"github.com/matthewmueller/golly/vdom"
	"github.com/matthewmueller/golly/vdom/h/h3"
)

// Header struct
type Header struct {
	vdom.Component

	// props
	props props

	// state
	state state
}

type props struct {
	title    string
	children []vdom.Child
	bats     string
}

// State struct
type state struct {
	count int
}

// New Header
func New(title string, children ...vdom.Child) vdom.Component {
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
func (d *Header) OnClick(e interface{}) {
	println(js.Raw("e.type"))
	d.SetState(&state{
		count: d.state.count + 1,
	})
}

// Render header
// js:"render"
func (d *Header) Render() vdom.Node {
	children := append(d.props.children, vdom.S(strconv.Itoa(d.state.count)))
	return h3.New(h3.Class(d.props.title).Attr("count", d.state.count), children...)
}
