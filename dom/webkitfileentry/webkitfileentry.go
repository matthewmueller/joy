package webkitfileentry

import (
	"github.com/matthewmueller/golly/dom/domerror"
	"github.com/matthewmueller/golly/dom/file"
	"github.com/matthewmueller/golly/dom/webkitfilesytem"
	"github.com/matthewmueller/golly/js"
)

var _ webkitfilesytem.WebKitEntry = (*WebKitFileEntry)(nil)

// WebKitFileEntry struct
// js:"WebKitFileEntry,omit"
type WebKitFileEntry struct {
}

// File fn
// js:"file"
func (*WebKitFileEntry) File(successCallback func(file *file.File), errorCallback *func(err *domerror.DOMError)) {
	js.Rewrite("$_.file($1, $2)", successCallback, errorCallback)
}

// Filesystem prop
// js:"filesystem"
func (*WebKitFileEntry) Filesystem() (filesystem *webkitfilesytem.WebKitFileSystem) {
	js.Rewrite("$_.filesystem")
	return filesystem
}

// FullPath prop
// js:"fullPath"
func (*WebKitFileEntry) FullPath() (fullPath string) {
	js.Rewrite("$_.fullPath")
	return fullPath
}

// IsDirectory prop
// js:"isDirectory"
func (*WebKitFileEntry) IsDirectory() (isDirectory bool) {
	js.Rewrite("$_.isDirectory")
	return isDirectory
}

// IsFile prop
// js:"isFile"
func (*WebKitFileEntry) IsFile() (isFile bool) {
	js.Rewrite("$_.isFile")
	return isFile
}

// Name prop
// js:"name"
func (*WebKitFileEntry) Name() (name string) {
	js.Rewrite("$_.name")
	return name
}
