package msblobbuilder

import (
	"github.com/matthewmueller/golly/dom/blob"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New() *MSBlobBuilder {
	js.Rewrite("MSBlobBuilder")
	return &MSBlobBuilder{}
}

// MSBlobBuilder struct
// js:"MSBlobBuilder,omit"
type MSBlobBuilder struct {
}

// Append fn
// js:"append"
func (*MSBlobBuilder) Append(data interface{}, endings *string) {
	js.Rewrite("$_.append($1, $2)", data, endings)
}

// GetBlob fn
// js:"getBlob"
func (*MSBlobBuilder) GetBlob(contentType *string) (b blob.Blob) {
	js.Rewrite("$_.getBlob($1)", contentType)
	return b
}
