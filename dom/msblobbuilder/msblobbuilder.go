package msblobbuilder

import (
	"github.com/matthewmueller/joy/dom/blob"
	"github.com/matthewmueller/joy/macro"
)

// New fn
func New() *MSBlobBuilder {
	macro.Rewrite("MSBlobBuilder")
	return &MSBlobBuilder{}
}

// MSBlobBuilder struct
// js:"MSBlobBuilder,omit"
type MSBlobBuilder struct {
}

// Append fn
// js:"append"
func (*MSBlobBuilder) Append(data interface{}, endings *string) {
	macro.Rewrite("$_.append($1, $2)", data, endings)
}

// GetBlob fn
// js:"getBlob"
func (*MSBlobBuilder) GetBlob(contentType *string) (b blob.Blob) {
	macro.Rewrite("$_.getBlob($1)", contentType)
	return b
}
