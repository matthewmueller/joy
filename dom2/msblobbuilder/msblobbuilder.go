package msblobbuilder

import (
	"github.com/matthewmueller/golly/dom2/blob"
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

// Append
func (*MSBlobBuilder) Append(data interface{}, endings *string) {
	js.Rewrite("$<.Append($1, $2)", data, endings)
}

// GetBlob
func (*MSBlobBuilder) GetBlob(contentType *string) (b blob.Blob) {
	js.Rewrite("$<.GetBlob($1)", contentType)
	return b
}
