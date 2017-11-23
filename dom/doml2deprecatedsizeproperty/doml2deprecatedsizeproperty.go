package doml2deprecatedsizeproperty

import "github.com/matthewmueller/golly/js"

// DOML2deprecatedSizeProperty struct
// js:"DOML2deprecatedSizeProperty,omit"
type DOML2deprecatedSizeProperty struct {
}

// Size prop
func (*DOML2DeprecatedSizeProperty) Size() (size int) {
	js.Rewrite("$<.size")
	return size
}

// Size prop
func (*DOML2DeprecatedSizeProperty) SetSize(size int) {
	js.Rewrite("$<.size = $1", size)
}
