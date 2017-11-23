package msstreamreader

import (
	"github.com/matthewmueller/golly/dom2/domerror"
	"github.com/matthewmueller/golly/js"
)

// js:"MSStreamReader,omit"
type MSStreamReader struct {
	window.EventTarget
	msbasereader.MSBaseReader
}

// ReadAsArrayBuffer
func (*MSStreamReader) ReadAsArrayBuffer(stream *msstream.MSStream, size *int64) {
	js.Rewrite("$<.ReadAsArrayBuffer($1, $2)", stream, size)
}

// ReadAsBinaryString
func (*MSStreamReader) ReadAsBinaryString(stream *msstream.MSStream, size *int64) {
	js.Rewrite("$<.ReadAsBinaryString($1, $2)", stream, size)
}

// ReadAsBlob
func (*MSStreamReader) ReadAsBlob(stream *msstream.MSStream, size *int64) {
	js.Rewrite("$<.ReadAsBlob($1, $2)", stream, size)
}

// ReadAsDataURL
func (*MSStreamReader) ReadAsDataURL(stream *msstream.MSStream, size *int64) {
	js.Rewrite("$<.ReadAsDataURL($1, $2)", stream, size)
}

// ReadAsText
func (*MSStreamReader) ReadAsText(stream *msstream.MSStream, encoding *string, size *int64) {
	js.Rewrite("$<.ReadAsText($1, $2, $3)", stream, encoding, size)
}

// Error
func (*MSStreamReader) Error() (err *domerror.DOMError) {
	js.Rewrite("$<.Error")
	return err
}
