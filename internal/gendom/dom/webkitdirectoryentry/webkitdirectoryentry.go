package webkitdirectoryentry

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/webkitdirectoryreader"
	"github.com/matthewmueller/golly/internal/gendom/dom/webkitentry"
	"github.com/matthewmueller/golly/js"
)

type WebKitDirectoryEntry struct {
	*webkitentry.WebKitEntry
}

func (*WebKitDirectoryEntry) CreateReader() (w *webkitdirectoryreader.WebKitDirectoryReader) {
	js.Rewrite("$<.createReader()")
	return w
}
