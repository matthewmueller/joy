package webkitdirectoryreader

import (
	"github.com/matthewmueller/golly/dom/domerror"
	"github.com/matthewmueller/golly/js"
)

// WebKitDirectoryReader struct
// js:"WebKitDirectoryReader,omit"
type WebKitDirectoryReader struct {
}

// ReadEntries fn
func (*WebKitDirectoryReader) ReadEntries(successCallback func(entries []*WebKitEntry), errorCallback *func(err *domerror.DOMError)) {
	js.Rewrite("$<.readEntries($1, $2)", successCallback, errorCallback)
}
