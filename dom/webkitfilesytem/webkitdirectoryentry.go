package webkitfilesytem

import (
	"github.com/matthewmueller/golly/dom/webkitdirectoryreader"
	"github.com/matthewmueller/golly/js"
)

// WebKitDirectoryEntry struct
// js:"WebKitDirectoryEntry,omit"
type WebKitDirectoryEntry struct {
	WebKitEntry
}

// CreateReader fn
func (*WebKitDirectoryEntry) CreateReader() (w *webkitdirectoryreader.WebKitDirectoryReader) {
	js.Rewrite("$<.createReader()")
	return w
}
