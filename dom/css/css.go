package css

import "github.com/matthewmueller/joy/js"

// CSS struct
// js:"CSS,omit"
type CSS struct {
}

// Supports fn
// js:"supports"
func (*CSS) Supports(property string, value *string) (b bool) {
	js.Rewrite("$_.supports($1, $2)", property, value)
	return b
}
