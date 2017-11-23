package css

import "github.com/matthewmueller/golly/js"

// CSS struct
// js:"CSS,omit"
type CSS struct {
}

// Supports
func (*CSS) Supports(property string, value *string) (b bool) {
	js.Rewrite("$<.Supports($1, $2)", property, value)
	return b
}
