package doml2deprecatedsizeproperty

import "github.com/matthewmueller/golly/js"

// DOML2deprecatedSizeProperty struct
// js:"DOML2deprecatedSizeProperty,omit"
type DOML2deprecatedSizeProperty struct {
}

// Size
func (*DOML2DeprecatedSizeProperty) Size() (size int) {
	js.Rewrite("$<.Size")
	return size
}

// Size
func (*DOML2DeprecatedSizeProperty) SetSize(size int) {
	js.Rewrite("$<.Size = $1", size)
}
