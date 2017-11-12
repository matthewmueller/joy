package header

import (
	"strconv"

	"github.com/matthewmueller/golly/jsx"
	"github.com/matthewmueller/golly/testdata/60-chaining/strong"
)

// import "github.com/matthewmueller/golly/js"

// header struct
type header struct {
	props *props
	state *state
}

type props struct {
	title    string
	body     string
	children []jsx.Node
}

type state struct {
	id string
}

// New fn
func New(title string, body string, children ...jsx.Node) *header {
	return &header{
		props: &props{
			title:    title,
			body:     body,
			children: children,
		},
	}
}

// Render function
func (h *header) Render() jsx.JSX {
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
