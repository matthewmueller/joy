package webkitfilesytem

import "github.com/matthewmueller/golly/js"

// WebKitFileSystem struct
// js:"WebKitFileSystem,omit"
type WebKitFileSystem struct {
}

// Name prop
func (*WebKitFileSystem) Name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// Root prop
func (*WebKitFileSystem) Root() (root *WebKitDirectoryEntry) {
	js.Rewrite("$<.root")
	return root
}
