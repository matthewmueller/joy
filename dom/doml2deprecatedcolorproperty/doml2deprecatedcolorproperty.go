package doml2deprecatedcolorproperty

// DOML2deprecatedColorProperty interface
// js:"DOML2deprecatedColorProperty"
type DOML2deprecatedColorProperty interface {

	// Color prop
	// js:"color"
	// jsrewrite:"$_.color"
	Color() (color string)

	// SetColor prop
	// js:"color"
	// jsrewrite:"$_.color = $1"
	SetColor(color string)
}
