package webkitfilesytem

import (
	"github.com/matthewmueller/joy/dom/webkitdirectoryreader"
	"github.com/matthewmueller/joy/js"
)

var _ WebKitEntry = (*WebKitDirectoryEntry)(nil)

// WebKitDirectoryEntry struct
// js:"WebKitDirectoryEntry,omit"
type WebKitDirectoryEntry struct {
}

// CreateReader fn
// js:"createReader"
func (*WebKitDirectoryEntry) CreateReader() (w *webkitdirectoryreader.WebKitDirectoryReader) {
	js.Rewrite("$_.createReader()")
	return w
}

// Filesystem prop
// js:"filesystem"
func (*WebKitDirectoryEntry) Filesystem() (filesystem *WebKitFileSystem) {
	js.Rewrite("$_.filesystem")
	return filesystem
}

// FullPath prop
// js:"fullPath"
func (*WebKitDirectoryEntry) FullPath() (fullPath string) {
	js.Rewrite("$_.fullPath")
	return fullPath
}

// IsDirectory prop
// js:"isDirectory"
func (*WebKitDirectoryEntry) IsDirectory() (isDirectory bool) {
	js.Rewrite("$_.isDirectory")
	return isDirectory
}

// IsFile prop
// js:"isFile"
func (*WebKitDirectoryEntry) IsFile() (isFile bool) {
	js.Rewrite("$_.isFile")
	return isFile
}

// Name prop
// js:"name"
func (*WebKitDirectoryEntry) Name() (name string) {
	js.Rewrite("$_.name")
	return name
}
