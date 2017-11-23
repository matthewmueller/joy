package webkitfileentry

import (
	"github.com/matthewmueller/golly/dom2/domerror"
	"github.com/matthewmueller/golly/dom2/file"
	"github.com/matthewmueller/golly/dom2/webkitfilesytem"
	"github.com/matthewmueller/golly/js"
)

// WebKitFileEntry struct
// js:"WebKitFileEntry,omit"
type WebKitFileEntry struct {
	webkitfilesytem.WebKitEntry
}

// File
func (*WebKitFileEntry) File(successCallback func(file *file.File), errorCallback *func(err *domerror.DOMError)) {
	js.Rewrite("$<.File($1, $2)", successCallback, errorCallback)
}
