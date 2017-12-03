package header

import (
	"strconv"

	"github.com/matthewmueller/joy/vdom"
	"github.com/matthewmueller/joy/vdom/h/strong"
)

// import "github.com/matthewmueller/joy/macro"

// header struct
type header struct {
	vdom.Component

	props *props
	state *state
}

type props struct {
	title    string
	body     string
	children []vdom.Child
}

type state struct {
	id string
}

// New fn
func New(title string, body string, children ...vdom.Child) vdom.Component {
	return &header{
		props: &props{
			title:    title,
			body:     body,
			children: children,
		},
	}
}

// Render function
func (h *header) Render() vdom.Node {
	return strong.New(strong.Class(h.props.title).ID(h.state.id), h.props.children...)
}

// ComponentWillMount function
func (h *header) ComponentWillMount() {
	h.state.id = strconv.Itoa(5)
}

// ComponentDidMount function
func (h *header) ComponentDidMount() {
	println("did mount")
}
