package doml2deprecatedcolorproperty

import "github.com/matthewmueller/golly/js"

// DOML2deprecatedColorProperty struct
// js:"DOML2deprecatedColorProperty,omit"
type DOML2deprecatedColorProperty struct {
}

// Color prop
func (*DOML2DeprecatedColorProperty) Color() (color string) {
	js.Rewrite("$<.color")
	return color
}

// Color prop
func (*DOML2DeprecatedColorProperty) SetColor(color string) {
	js.Rewrite("$<.color = $1", color)
}
