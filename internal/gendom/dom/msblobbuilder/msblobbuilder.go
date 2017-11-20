package msblobbuilder

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/blob"
	"github.com/matthewmueller/golly/js"
)

type MSBlobBuilder struct {
}

func (*MSBlobBuilder) Append(data interface{}, endings string) {
	js.Rewrite("$<.append($1, $2)", data, endings)
}

func (*MSBlobBuilder) GetBlob(contentType string) (b *blob.Blob) {
	js.Rewrite("$<.getBlob($1)", contentType)
	return b
}
