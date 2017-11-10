package header

import "github.com/matthewmueller/golly/testdata/60-chaining/strong"
import "github.com/matthewmueller/golly/js"
import "github.com/matthewmueller/golly/jsx"

type header struct {
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
func New(title string, body string, children ...jsx.Node) jsx.Node {
	js.Rewrite("preact.h(header, { title: $1, body: $2 }, $3)",title, body, children)
	return &header{
		props: &props{
			title: title,
			body: body,
			children: children,
		},
	}
}

func (h *header) Render() jsx.JSX {
	return strong.New(strong.Class(h.props.title))
}