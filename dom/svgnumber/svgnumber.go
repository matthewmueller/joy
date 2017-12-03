package svgnumber

import "github.com/matthewmueller/joy/js"

// SVGNumber struct
// js:"SVGNumber,omit"
type SVGNumber struct {
}

// Value prop
// js:"value"
func (*SVGNumber) Value() (value float32) {
	js.Rewrite("$_.value")
	return value
}

// SetValue prop
// js:"value"
func (*SVGNumber) SetValue(value float32) {
	js.Rewrite("$_.value = $1", value)
}
