package webkitdirectoryreader

import (
	"github.com/matthewmueller/golly/dom2/domerror"
	"github.com/matthewmueller/golly/js"
)

// WebKitDirectoryReader struct
// js:"WebKitDirectoryReader,omit"
type WebKitDirectoryReader struct {
}

// ReadEntries
func (*WebKitDirectoryReader) ReadEntries(successCallback func(entries []*WebKitEntry), errorCallback *func(err *domerror.DOMError)) {
	js.Rewrite("$<.ReadEntries($1, $2)", successCallback, errorCallback)
}
