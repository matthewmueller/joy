package webkitdirectoryreader

import "github.com/matthewmueller/golly/js"

// js:"WebKitDirectoryReader,omit"
type WebKitDirectoryReader struct {
}

// ReadEntries
func (*WebKitDirectoryReader) ReadEntries(successCallback func(entries []*WebKitEntry) handleEvent, errorCallback *func(err *domerror.DOMError) handleEvent) {
	js.Rewrite("$<.ReadEntries($1, $2)", successCallback, errorCallback)
}
