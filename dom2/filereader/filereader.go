package filereader

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"FileReader,omit"
type FileReader struct {
	window.EventTarget
	msbasereader.MSBaseReader
}

// ReadAsArrayBuffer
func (*FileReader) ReadAsArrayBuffer(blob blob.Blob) {
	js.Rewrite("$<.ReadAsArrayBuffer($1)", blob)
}

// ReadAsBinaryString
func (*FileReader) ReadAsBinaryString(blob blob.Blob) {
	js.Rewrite("$<.ReadAsBinaryString($1)", blob)
}

// ReadAsDataURL
func (*FileReader) ReadAsDataURL(blob blob.Blob) {
	js.Rewrite("$<.ReadAsDataURL($1)", blob)
}

// ReadAsText
func (*FileReader) ReadAsText(blob blob.Blob, encoding *string) {
	js.Rewrite("$<.ReadAsText($1, $2)", blob, encoding)
}

// Error
func (*FileReader) Error() (err *domerror.DOMError) {
	js.Rewrite("$<.Error")
	return err
}
