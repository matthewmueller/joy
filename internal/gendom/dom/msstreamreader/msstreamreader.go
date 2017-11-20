package msstreamreader

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/domerror"
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/msbasereader"
	"github.com/matthewmueller/golly/internal/gendom/dom/msstream"
	"github.com/matthewmueller/golly/js"
)

type MSStreamReader struct {
	*eventtarget.EventTarget
	*msbasereader.MSBaseReader
}

func (*MSStreamReader) ReadAsArrayBuffer(stream *msstream.MSStream, size int64) {
	js.Rewrite("$<.readAsArrayBuffer($1, $2)", stream, size)
}

func (*MSStreamReader) ReadAsBinaryString(stream *msstream.MSStream, size int64) {
	js.Rewrite("$<.readAsBinaryString($1, $2)", stream, size)
}

func (*MSStreamReader) ReadAsBlob(stream *msstream.MSStream, size int64) {
	js.Rewrite("$<.readAsBlob($1, $2)", stream, size)
}

func (*MSStreamReader) ReadAsDataURL(stream *msstream.MSStream, size int64) {
	js.Rewrite("$<.readAsDataURL($1, $2)", stream, size)
}

func (*MSStreamReader) ReadAsText(stream *msstream.MSStream, encoding string, size int64) {
	js.Rewrite("$<.readAsText($1, $2, $3)", stream, encoding, size)
}

func (*MSStreamReader) GetError() (err *domerror.DOMError) {
	js.Rewrite("$<.error")
	return err
}
