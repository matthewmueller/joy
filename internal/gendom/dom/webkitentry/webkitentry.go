package webkitentry

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/webkitfilesystem"
	"github.com/matthewmueller/golly/js"
)

type WebKitEntry struct {
}

func (*WebKitEntry) GetFilesystem() (filesystem *webkitfilesystem.WebKitFileSystem) {
	js.Rewrite("$<.filesystem")
	return filesystem
}

func (*WebKitEntry) GetFullPath() (fullPath string) {
	js.Rewrite("$<.fullPath")
	return fullPath
}

func (*WebKitEntry) GetIsDirectory() (isDirectory bool) {
	js.Rewrite("$<.isDirectory")
	return isDirectory
}

func (*WebKitEntry) GetIsFile() (isFile bool) {
	js.Rewrite("$<.isFile")
	return isFile
}

func (*WebKitEntry) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}
