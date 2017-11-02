package text

import "github.com/matthewmueller/golly/testdata/49-jsx/jsx"

// Text struct
type Text struct {
	Value string
}

// Render fn
func (t *Text) Render() jsx.Node {
	return t
}

func (t *Text) String() string {
	return t.Value
}
