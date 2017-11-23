package webkitfilesytem

import (
	"github.com/matthewmueller/golly/dom2/webkitdirectoryreader"
	"github.com/matthewmueller/golly/js"
)

// WebKitDirectoryEntry struct
// js:"WebKitDirectoryEntry,omit"
type WebKitDirectoryEntry struct {
	WebKitEntry
}

// CreateReader
func (*WebKitDirectoryEntry) CreateReader() (w *webkitdirectoryreader.WebKitDirectoryReader) {
	js.Rewrite("$<.CreateReader()")
	return w
}
