package svgnumber

import "github.com/matthewmueller/golly/js"

// js:"SVGNumber,omit"
type SVGNumber struct {
}

// Value
func (*SVGNumber) Value() (value float32) {
	js.Rewrite("$<.Value")
	return value
}

// Value
func (*SVGNumber) SetValue(value float32) {
	js.Rewrite("$<.Value = $1", value)
}
