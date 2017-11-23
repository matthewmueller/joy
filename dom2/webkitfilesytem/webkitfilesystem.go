package webkitfilesytem

import "github.com/matthewmueller/golly/js"

// js:"WebKitFileSystem,omit"
type WebKitFileSystem struct {
}

// Name
func (*WebKitFileSystem) Name() (name string) {
	js.Rewrite("$<.Name")
	return name
}

// Root
func (*WebKitFileSystem) Root() (root *WebKitDirectoryEntry) {
	js.Rewrite("$<.Root")
	return root
}
