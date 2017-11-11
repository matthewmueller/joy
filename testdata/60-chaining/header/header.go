package header

import "github.com/matthewmueller/golly/testdata/60-chaining/strong"
// import "github.com/matthewmueller/golly/js"
import "github.com/matthewmueller/golly/jsx"

// Header struct
type Header struct {
	props *props
	state *state
}

type props struct{
	title string
	body string
	children []jsx.Node
}

type state struct{
	
}

// New fn
func New(title string, body string, children ...jsx.Node) *Header {
	// js.Rewrite("$1.h(header.Header, { title: $2, body: $3 }, $4)", js.RawFile("../preact/preact.js"), title, body, children)
	return &Header{
		props: &props{
			title: title,
			body: body,
			children: children,
		},
	}
}

// Render function
func (h *Header) Render() jsx.JSX {
	return strong.New(strong.Class(h.props.title))
}