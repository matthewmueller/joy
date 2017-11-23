package svgnumber

import "github.com/matthewmueller/golly/js"

// SVGNumber struct
// js:"SVGNumber,omit"
type SVGNumber struct {
}

// Value prop
func (*SVGNumber) Value() (value float32) {
	js.Rewrite("$<.value")
	return value
}

// Value prop
func (*SVGNumber) SetValue(value float32) {
	js.Rewrite("$<.value = $1", value)
}
