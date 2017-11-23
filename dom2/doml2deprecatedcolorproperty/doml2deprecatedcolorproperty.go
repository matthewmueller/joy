package doml2deprecatedcolorproperty

import "github.com/matthewmueller/golly/js"

// DOML2deprecatedColorProperty struct
// js:"DOML2deprecatedColorProperty,omit"
type DOML2deprecatedColorProperty struct {
}

// Color
func (*DOML2DeprecatedColorProperty) Color() (color string) {
	js.Rewrite("$<.Color")
	return color
}

// Color
func (*DOML2DeprecatedColorProperty) SetColor(color string) {
	js.Rewrite("$<.Color = $1", color)
}
