package filereader

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/blob"
	"github.com/matthewmueller/golly/internal/gendom/dom/domerror"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/msbasereader"
	"github.com/matthewmueller/golly/js"
)

type FileReader struct {
	*eventtarget.EventTarget
	*msbasereader.MSBaseReader
}

func (*FileReader) ReadAsArrayBuffer(blob *blob.Blob) {
	js.Rewrite("$<.readAsArrayBuffer($1)", blob)
}

func (*FileReader) ReadAsBinaryString(blob *blob.Blob) {
	js.Rewrite("$<.readAsBinaryString($1)", blob)
}

func (*FileReader) ReadAsDataURL(blob *blob.Blob) {
	js.Rewrite("$<.readAsDataURL($1)", blob)
}

func (*FileReader) ReadAsText(blob *blob.Blob, encoding string) {
	js.Rewrite("$<.readAsText($1, $2)", blob, encoding)
}

func (*FileReader) GetError() (err *domerror.DOMError) {
	js.Rewrite("$<.error")
	return err
}
