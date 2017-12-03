package webkitfileentry

import (
	"github.com/matthewmueller/joy/dom/domerror"
	"github.com/matthewmueller/joy/dom/file"
	"github.com/matthewmueller/joy/dom/webkitfilesytem"
	"github.com/matthewmueller/joy/macro"
)

var _ webkitfilesytem.WebKitEntry = (*WebKitFileEntry)(nil)

// WebKitFileEntry struct
// js:"WebKitFileEntry,omit"
type WebKitFileEntry struct {
}

// File fn
// js:"file"
func (*WebKitFileEntry) File(successCallback func(file *file.File), errorCallback *func(err *domerror.DOMError)) {
	macro.Rewrite("$_.file($1, $2)", successCallback, errorCallback)
}

// Filesystem prop
// js:"filesystem"
func (*WebKitFileEntry) Filesystem() (filesystem *webkitfilesytem.WebKitFileSystem) {
	macro.Rewrite("$_.filesystem")
	return filesystem
}

// FullPath prop
// js:"fullPath"
func (*WebKitFileEntry) FullPath() (fullPath string) {
	macro.Rewrite("$_.fullPath")
	return fullPath
}

// IsDirectory prop
// js:"isDirectory"
func (*WebKitFileEntry) IsDirectory() (isDirectory bool) {
	macro.Rewrite("$_.isDirectory")
	return isDirectory
}

// IsFile prop
// js:"isFile"
func (*WebKitFileEntry) IsFile() (isFile bool) {
	macro.Rewrite("$_.isFile")
	return isFile
}

// Name prop
// js:"name"
func (*WebKitFileEntry) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}
