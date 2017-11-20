package webkitfilesystem

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/webkitdirectoryentry"
	"github.com/matthewmueller/golly/js"
)

type WebKitFileSystem struct {
}

func (*WebKitFileSystem) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*WebKitFileSystem) GetRoot() (root *webkitdirectoryentry.WebKitDirectoryEntry) {
	js.Rewrite("$<.root")
	return root
}
