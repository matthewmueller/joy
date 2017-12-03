package webkitdirectoryreader

import (
	"github.com/matthewmueller/joy/dom/domerror"
	"github.com/matthewmueller/joy/macro"
)

// WebKitDirectoryReader struct
// js:"WebKitDirectoryReader,omit"
type WebKitDirectoryReader struct {
}

// ReadEntries fn
// js:"readEntries"
func (*WebKitDirectoryReader) ReadEntries(successCallback func(entries []*WebKitEntry), errorCallback *func(err *domerror.DOMError)) {
	macro.Rewrite("$_.readEntries($1, $2)", successCallback, errorCallback)
}
