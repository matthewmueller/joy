package doml2deprecatedsizeproperty

import "github.com/matthewmueller/golly/js"

type DOML2DeprecatedSizeProperty struct {
}

func (*DOML2DeprecatedSizeProperty) GetSize() (size int) {
	js.Rewrite("$<.size")
	return size
}

func (*DOML2DeprecatedSizeProperty) SetSize(size int) {
	js.Rewrite("$<.size = $1", size)
}
