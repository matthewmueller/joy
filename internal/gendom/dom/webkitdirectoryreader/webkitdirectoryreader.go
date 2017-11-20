package webkitdirectoryreader

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/domerror"
	"github.com/matthewmueller/golly/js"
)

type WebKitDirectoryReader struct {
}

func (*WebKitDirectoryReader) ReadEntries(successCallback func(entries []*WebKitEntry), errorCallback func(err *domerror.DOMError)) {
	js.Rewrite("$<.readEntries($1, $2)", successCallback, errorCallback)
}
