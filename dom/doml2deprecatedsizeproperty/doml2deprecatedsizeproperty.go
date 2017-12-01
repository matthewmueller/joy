package doml2deprecatedsizeproperty

// DOML2deprecatedSizeProperty interface
// js:"DOML2deprecatedSizeProperty"
type DOML2deprecatedSizeProperty interface {

	// Size prop
	// js:"size"
	// jsrewrite:"$_.size"
	Size() (size int)

	// SetSize prop
	// js:"size"
	// jsrewrite:"$_.size = $1"
	SetSize(size int)
}
