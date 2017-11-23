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

// Append fn
func (*MSBlobBuilder) Append(data interface{}, endings *string) {
	js.Rewrite("$<.append($1, $2)", data, endings)
}

// GetBlob fn
func (*MSBlobBuilder) GetBlob(contentType *string) (b blob.Blob) {
	js.Rewrite("$<.getBlob($1)", contentType)
	return b
}
