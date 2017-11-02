package strong

import (
	"github.com/matthewmueller/golly/jsx"
)

// Strong struct
type Strong struct {
	Class    string
	Children []jsx.Component
}

// Render strong
func (s *Strong) Render() jsx.JSX {
	// js.Rewrite("preact.h('strong', $2, $3)", map[string]string{
	// 	"class": s.Class,
	// }, s.Children[0])

	return jsx.H("strong", map[string]interface{}{
		"class": s.Class,
	}, s.Children[0])
}
