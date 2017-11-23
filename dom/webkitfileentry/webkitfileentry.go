package webkitfileentry

import (
	"github.com/matthewmueller/golly/dom/domerror"
	"github.com/matthewmueller/golly/dom/file"
	"github.com/matthewmueller/golly/dom/webkitfilesytem"
	"github.com/matthewmueller/golly/js"
)

// WebKitFileEntry struct
// js:"WebKitFileEntry,omit"
type WebKitFileEntry struct {
	webkitfilesytem.WebKitEntry
}

// File fn
func (*WebKitFileEntry) File(successCallback func(file *file.File), errorCallback *func(err *domerror.DOMError)) {
	js.Rewrite("$<.file($1, $2)", successCallback, errorCallback)
}
